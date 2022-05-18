package server

import "sync"

type ClientManager struct {
	sync.Mutex
	devices map[string]*Client
	clients map[*Client]bool
}

func NewClientManager() *ClientManager {
	cm := new(ClientManager)
	cm.devices = make(map[string]*Client)
	cm.clients = make(map[*Client]bool)
	return cm
}

func (cm *ClientManager) Add(client *Client) {
	cm.Lock()
	cm.devices[client.DeviceId] = client
	cm.clients[client] = true
	cm.Unlock()
}

func (cm *ClientManager) Get(id string) *Client {
	cm.Lock()
	defer cm.Unlock()

	client := cm.devices[id]
	return client
}

func (cm *ClientManager) Remove(client *Client) {
	cm.Lock()
	defer cm.Unlock()

	if ok := cm.clients[client]; ok {
		c := cm.devices[client.DeviceId]
		if c == client {
			delete(cm.devices, c.DeviceId)
		}
	}

	return
}
