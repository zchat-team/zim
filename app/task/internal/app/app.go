package app

import (
	"flag"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/zmicro-team/zim/app/task/internal/server"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
	"github.com/zmicro-team/zmicro/core/util/env"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "config", "config.yaml", "config file")
}

type App struct {
	opts   Options
	zc     *zconfig
	server *server.Server
}

type zconfig struct {
	App struct {
		Mode string
		Name string
	}
	Logger struct {
		Level      string `json:"level"`
		Filename   string `json:"filename"`
		MaxSize    int    `json:"maxSize"`
		MaxBackups int    `json:"maxBackups"`
		MaxAge     int    `json:"maxAge"`
		Compress   bool   `json:"compress"`
	}
	Nats struct {
		Addr string
	}
}

func New(opts ...Option) *App {
	options := newOptions(opts...)
	flag.Parse()
	_, err := os.Stat(cfgFile)
	if os.IsNotExist(err) {
		log.Fatal("config file not exists")
	}

	c := config.New(config.Path(cfgFile))
	config.ResetDefault(c)

	zc := &zconfig{}
	if err = config.Unmarshal(zc); err != nil {
		log.Fatal(err.Error())
	}

	if zc.App.Name == "" {
		log.Fatal("配置项app.name不能为空")
	}

	level, err := zapcore.ParseLevel(zc.Logger.Level)
	if err != nil {
		level = log.InfoLevel
	}
	if env.IsDevelop() {
		w := &lumberjack.Logger{
			Filename:   zc.Logger.Filename,
			MaxSize:    zc.Logger.MaxSize,
			MaxBackups: zc.Logger.MaxBackups,
			MaxAge:     zc.Logger.MaxAge,
			Compress:   zc.Logger.Compress,
		}
		l := log.NewTee([]io.Writer{os.Stderr, w}, level, log.WithCaller(true))
		log.ResetDefault(l)
	} else {
		w := &lumberjack.Logger{
			Filename:   zc.Logger.Filename,
			MaxSize:    zc.Logger.MaxSize,
			MaxBackups: zc.Logger.MaxBackups,
			MaxAge:     zc.Logger.MaxAge,
			Compress:   zc.Logger.Compress,
		}
		l := log.New(w, level, log.WithCaller(true))
		log.ResetDefault(l)
	}

	app := &App{
		opts: options,
		zc:   zc,
	}

	app.server = server.NewServer()
	return app
}

func (a *App) Run() error {
	if a.opts.Before != nil {
		if err := a.opts.Before(); err != nil {
			return err
		}
	}
	if err := a.server.Start(); err != nil {
		return err
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	log.Infof("received signal %s", <-ch)

	err := a.server.Stop()

	return err
}
