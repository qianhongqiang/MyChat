package libnet

import "sync"

type channel struct {
	mutex sync.Mutex
	sessions map[int]*Session
	state interface{}
}