package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numPhilosophers = 5

// Philosopher struct to represent each philosopher
type Philosopher struct {
	id            int
	leftFork      *sync.Mutex
	rightFork     *sync.Mutex
	doneEating    chan bool
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking...\n", p.id)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
}

func (p *Philosopher) eat() {
	// Try to acquire both forks (left and right) to eat
	p.leftFork.Lock()
	p.rightFork.Lock()

	// Eating
	fmt.Printf("Philosopher %d is eating...\n", p.id)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

	// Done eating, release forks
	p.leftFork.Unlock()
	p.rightFork.Unlock()

	fmt.Printf("Philosopher %d has finished eating.\n", p.id)
	p.doneEating <- true // Notify that the philosopher is done eating
}

func philosopherRoutine(p *Philosopher, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Philosopher thinks first, then tries to eat
		p.think()
		p.eat()
	}
}

func main() {
	// Mutexes for the forks (resources)
	var forks [numPhilosophers]*sync.Mutex
	for i := range forks {
		forks[i] = &sync.Mutex{}
	}

	// Philosophers
	var philosophers [numPhilosophers]Philosopher
	var wg sync.WaitGroup

	// Create philosophers and associate them with forks
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = Philosopher{
			id:            i,
			leftFork:      forks[i],
			rightFork:     forks[(i+1)%numPhilosophers], // the next fork is the right fork
			doneEating:    make(chan bool),
		}
		wg.Add(1)
		go philosopherRoutine(&philosophers[i], &wg)
	}

	// Run the simulation for a limited amount of time (to avoid infinite loops)
	time.Sleep(10 * time.Second) // Change this to however long you'd like the simulation to run

	// Shutdown all philosophers after the simulation ends
	for i := 0; i < numPhilosophers; i++ {
		<-philosophers[i].doneEating // Wait for each philosopher to finish eating
	}

	// Wait for all philosophers to finish their routines
	wg.Wait()
	fmt.Println("Simulation finished.")
}
