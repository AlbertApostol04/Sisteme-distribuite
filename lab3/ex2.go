package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct{ counter int }

func (l *LamportClock) Tick() { l.counter++ }

func (l *LamportClock) Send() { l.Tick() }

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

const rounds = 3

func process(id int, in <-chan Message, out chan<- Message, events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{}

	for r := 0; r < rounds; r++ {
		clock.counter += rand.Intn(4)
		clock.Tick()
		events <- Event{id, clock.counter, "internal"}

		clock.counter += rand.Intn(4)
		clock.Send()
		out <- Message{sender: id, ts: clock.counter}
		events <- Event{id, clock.counter, "send"}

		clock.counter += rand.Intn(4)
		msg := <-in
		clock.Receive(msg.ts)
		events <- Event{id, clock.counter, "recv"}
	}
}

func main() {
	c12 := make(chan Message, rounds)
	c23 := make(chan Message, rounds)
	c31 := make(chan Message, rounds)

	events := make(chan Event, 3*rounds*3)

	var wg sync.WaitGroup
	wg.Add(3)
	go process(1, c31, c12, events, &wg)
	go process(2, c12, c23, events, &wg)
	go process(3, c23, c31, events, &wg)

	wg.Wait()
	close(events)

	all := []Event{}
	for e := range events {
		all = append(all, e)
	}

	sort.Slice(all, func(i, j int) bool { return less(all[i], all[j]) })

	for _, e := range all {
		fmt.Printf("C = %3d   P%d   %s\n", e.internal_counter, e.process, e.kind)
	}
}
