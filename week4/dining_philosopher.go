package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

const number_philosophers = 5
const max_eat_times = 3
const host_allows = 2

var host = make(chan bool, host_allows)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	name        string
	left, right *ChopStick
}

func (p Philosopher) eat() {
	defer wg.Done()
	for i := 0; i < max_eat_times; i++ {
		host <- true
		p.left.Lock()
		p.right.Lock()
		fmt.Printf("Starting to eat \t%s\t\t%d times\n", p.name, i+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Finishing eating \t%s\n\n", p.name)
		p.left.Unlock()
		p.right.Unlock()
		<-host
	}
}

func genName() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 6 + rand.Intn(14-6)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func createChopSticks(n int) []*ChopStick {
	ChopSticks := make([]*ChopStick, n)
	for i := 0; i < n; i++ {
		ChopSticks[i] = new(ChopStick)
	}
	return ChopSticks

}

func creatPhilosophers(ChopSticks []*ChopStick, n int) []*Philosopher {
	Philosophers := make([]*Philosopher, n)
	for i := 0; i < n; i++ {
		name := genName()
		fmt.Printf("Philosopher with name: %s is created !!!\n", name)
		Philosophers[i] = &Philosopher{name, ChopSticks[i], ChopSticks[(i+1)%n]}
	}
	return Philosophers
}

func main() {
	fmt.Println("Assignment week4: dining philosopher")
	ChopSticks := createChopSticks(number_philosophers)
	Philosophers := creatPhilosophers(ChopSticks, number_philosophers)

	for i := 0; i < number_philosophers; i++ {
		wg.Add(1)
		go Philosophers[i].eat()
	}

	wg.Wait()

}
