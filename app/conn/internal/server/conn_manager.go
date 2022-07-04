package server

import "sync"

type ConnManager struct {
	sync.Mutex
	//devices map[string]*Connection
	//clients map[*Connection]bool

	conns map[string]*Connection
}

func NewConnManager() *ConnManager {
	cm := new(ConnManager)
	//cm.devices = make(map[string]*Connection)
	//cm.clients = make(map[*Connection]bool)

	cm.conns = make(map[string]*Connection)
	return cm
}

func (cm *ConnManager) Add(c *Connection) {
	cm.Lock()
	//cm.devices[c.DeviceId] = c
	//cm.clients[c] = true

	cm.conns[c.ID] = c
	cm.Unlock()
}

func (cm *ConnManager) Get(id string) *Connection {
	cm.Lock()
	defer cm.Unlock()

	//c := cm.devices[id]

	c := cm.conns[id]
	return c
}

func (cm *ConnManager) Remove(c *Connection) {
	cm.Lock()
	defer cm.Unlock()

	//if ok := cm.clients[c]; ok {
	//	d := cm.devices[c.DeviceId]
	//	if d == c {
	//		delete(cm.devices, d.DeviceId)
	//	}
	//}

	delete(cm.conns, c.ID)

	return
}
