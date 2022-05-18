package app

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zmicro-team/zim/app/conn/internal/server"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
)

var cfgFile string

func init() {
	flag.StringVar(&cfgFile, "config", "config.yaml", "config file")
}

type App struct {
	opts   Options
	conf   *zconfig
	server *server.Server
}

type zconfig struct {
	App struct {
		Name    string
		TcpAddr string
		WsAddr  string
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

	conf := &zconfig{}
	if err = config.Scan("app", &conf.App); err != nil {
		log.Fatal(err)
	}
	options.Name = conf.App.Name

	app := &App{
		opts: options,
		conf: conf,
	}

	if err = config.Scan("nats", &conf.Nats); err != nil {
		log.Fatal(err)
	}
	app.server = server.NewServer(
		server.TcpAddr(conf.App.TcpAddr),
		server.WsAddr(conf.App.WsAddr),
		server.NatsAddr(conf.Nats.Addr),
	)
	return app
}

func (a *App) Run() error {
	if err := a.server.Start(); err != nil {
		return err
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	log.Infof("received signal %s", <-ch)

	err := a.server.Stop()

	return err
}
