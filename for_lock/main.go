package main

import (
	"sync"
	"time"
)

var (
	lock  = sync.Mutex{}
	money = 100
)

func main() {
	go add()
	go remove()

	time.Sleep(3000 * time.Millisecond)
	println(money)
}

func add() {
	for i := 1; i <= 100000; i++ {
		lock.Lock()
		money += 1
		lock.Unlock()
	}
	println("add done")
}

func remove() {
	for i := 1; i <= 100000; i++ {
		lock.Lock()
		money -= 1
		lock.Unlock()
	}
	println("remove done")
}
