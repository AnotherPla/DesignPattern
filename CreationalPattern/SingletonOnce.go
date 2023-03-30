package creationalpattern

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton2 struct{}

func (s *singleton2) Something() {
	fmt.Println("once singleton something")
}

var instance2 *singleton2

func GetInstance2() *singleton2 {
	//once.Do中的函数只会执行一次
	once.Do(func() {
		instance2 = new(singleton2)
	})

	return instance2
}
