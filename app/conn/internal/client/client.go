package client

import (
	"github.com/zmicro-team/zim/proto/chat"
	"github.com/zmicro-team/zim/proto/sess"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
	"github.com/zmicro-team/zmicro/core/transport/rpc/client"
)

var (
	chatClient *chat.ChatClient
	convClient *chat.ConvClient
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

func GetChatClient() *chat.ChatClient {
	if chatClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("GetChatClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Chat"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("getChatClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		chatClient = chat.NewChatClient(cli)
	}
	return chatClient
}

func GetConvClient() *chat.ConvClient {
	if convClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("getConvClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Conv"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("getConvClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		convClient = chat.NewConvClient(cli)
	}
	return convClient
}
