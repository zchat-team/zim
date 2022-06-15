package client

import (
	"github.com/zchat-team/zim/proto/rpc/sess"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
	"github.com/zmicro-team/zmicro/core/transport/rpc/client"
)

var (
	sessClient *sess.SessClient
)

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func GetSessClient() *sess.SessClient {
	if sessClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("getSessClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Sess"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("getSessClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		sessClient = sess.NewSessClient(cli)
	}
	return sessClient
}
