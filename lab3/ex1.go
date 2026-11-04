package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()

	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

func main() {
	offsetInitial := 2 * time.Second
	server := &Clock{offset: 0}

	client := &Clock{offset: offsetInitial}
	theta, delta := measure(client, server)

	fmt.Println("a) O singura masurare")
	fmt.Println("   theta =", theta)
	fmt.Println("   delta =", delta)
	fmt.Println("   offset initial:       ", client.offset)

	client.offset += theta
	offsetA := client.offset
	fmt.Println("   offset dupa corectie: ", offsetA)

	if abs(offsetA) < abs(offsetInitial) {
		fmt.Println("   Offset-ul a scazut")
	} else {
		fmt.Println("   Offset-ul NU a scazut")
	}

	fmt.Println()
	fmt.Println("b) 10 masurari")

	var deltaMin, thetaBest time.Duration
	for i := 0; i < 10; i++ {
		c := &Clock{offset: offsetInitial}
		th, de := measure(c, server)
		fmt.Printf("   %2d: theta = %v, delta = %v\n", i+1, th, de)

		if i == 0 || de < deltaMin {
			deltaMin = de
			thetaBest = th
		}
	}

	offsetB := offsetInitial + thetaBest
	fmt.Println()
	fmt.Println("   delta minim:                  ", deltaMin)
	fmt.Println("   theta (la delta minim):       ", thetaBest)
	fmt.Println("   offset dupa corectie, punct b:", offsetB)
	fmt.Println("   offset dupa corectie, punct a:", offsetA)
}
