package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	LeftCs, RightCS *ChopS
	ID              int
	EatingTimes     int
}

//allow only two eatings at any time
func host(phChan chan *Philo) {
	for {
		if len(phChan) == 2 {
			<-phChan
			<-phChan
		}
	}
}

func (p Philo) eat(phChan chan *Philo, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		phChan <- &p

		p.LeftCs.Lock()
		p.RightCS.Lock()

		fmt.Println("Philospher ", p.ID, " is eating", " ", i+1, " of 3 times")
		p.EatingTimes++
		p.RightCS.Unlock()
		p.LeftCs.Unlock()

		fmt.Println("Philospher ", p.ID, " finished")

		if p.EatingTimes == 3 {
			wg.Done() //done only if philo eats three times
		}
	}
}

func main() {

	var wg sync.WaitGroup

	var phChan chan *Philo
	phChan = make(chan *Philo, 2)

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1, 0}
	}

	go host(phChan)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(phChan, &wg)
	}

	wg.Wait()
}
