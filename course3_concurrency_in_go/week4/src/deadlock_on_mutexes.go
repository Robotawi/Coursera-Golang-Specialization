package main

import "sync"

var mx1 sync.Mutex
var mx2 sync.Mutex

func worker1() {
	// wait to get something from ch1
	mx1.Lock()
	// then send on ch2
	mx2.Unlock()
}

func worker2() {
	// wait to get something from ch1
	mx2.Lock()
	// then send on ch2
	mx1.Unlock()
}

func main() {

	worker1()

	worker2()

}
