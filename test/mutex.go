package test

import (
	"fmt"
	"sync"
	"time"
)

var number int
var mu *sync.Mutex
var mua *sync.RWMutex

func Add(i int, mutex *sync.Mutex) {
	mutex.Lock()
	// fmt.Println(i, " lock ", number, mutex)
	number++
	mutex.Unlock()
	// fmt.Println(i, " unlo ", number, mutex)
}

func Mutex() {
	mu = &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go Add(i, mu)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(number)
}
