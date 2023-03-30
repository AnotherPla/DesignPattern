package creationalpattern

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lock sync.Mutex

// 标志位
var initialized uint32

type singleton struct{}

func (s *singleton) Something() {
	fmt.Println("singleton something")
}

var instance *singleton

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		atomic.StoreUint32(&initialized, 1)
		instance = new(singleton)
	}

	return instance
}
