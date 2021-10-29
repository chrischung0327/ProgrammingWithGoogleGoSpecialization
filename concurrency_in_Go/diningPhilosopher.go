package main

import (
	"fmt"
	"sync"
	"time"
)

var host = make(chan bool, 2)
var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	defer wg.Done()
	for j := 0; j < 3; j++ {
		host <- true
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat Philo: %d, times: %d\n", p.id, j+1)
		time.Sleep(time.Second)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		fmt.Printf("finishing to eat Philo: %d, times: %d\n", p.id, j+1)
		<-host
	}
}

func main() {
	// Initial
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
