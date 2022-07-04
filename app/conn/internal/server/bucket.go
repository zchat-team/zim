package server

import (
	"sync"
)

type Bucket struct {
	sync.RWMutex
	conns map[string]*Connection
}
