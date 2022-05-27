package client

import (
	"github.com/zchat-team/zim/proto/chat"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/transport/rpc/client"
)

var (
	chatClient *chat.ChatClient
)

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func GetChatClient() *chat.ChatClient {
	if chatClient == nil {
		r := &Registry{}
		config.Scan("registry", &r)

		// TODO: 优化
		cc, _ := client.NewClient(
			client.WithServiceName("Chat"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		cli := cc.GetXClient()
		chatClient = chat.NewChatClient(cli)
	}

	return chatClient
}
