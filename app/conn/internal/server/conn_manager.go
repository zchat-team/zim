package server

import "sync"

type ConnManager struct {
	sync.Mutex
	devices map[string]*Connection
	conns   map[*Connection]bool
}

func NewConnManager() *ConnManager {
	cm := new(ConnManager)
	cm.devices = make(map[string]*Connection)
	cm.conns = make(map[*Connection]bool)
	return cm
}

func (cm *ConnManager) Add(c *Connection) {
	cm.Lock()
	cm.devices[c.DeviceId] = c
	cm.conns[c] = true
	cm.Unlock()
}

func (cm *ConnManager) Get(id string) *Connection {
	cm.Lock()
	defer cm.Unlock()

	c := cm.devices[id]
	return c
}

func (cm *ConnManager) Remove(c *Connection) {
	cm.Lock()
	defer cm.Unlock()

	if ok := cm.conns[c]; ok {
		d := cm.devices[c.DeviceId]
		if d == c {
			delete(cm.devices, d.DeviceId)
		}
	}

	return
}
