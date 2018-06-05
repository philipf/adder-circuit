package main

import (
	"time"
)

func main() {
	testAll()
}

func testAll() {
	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)

	a0, a1 := split(a)
	b0, b1 := split(b)
	c0, c1 := split(c)

	xor1 := xorGate(a0, b0)
	xor1_0, xor1_1 := split(xor1)

	xor2 := xorGate(xor1_0, c0)
	and1 := andGate(a1, b1)
	and2 := andGate(xor1_1, c1)
	or1 := orGate(and1, and2)

	sum := false
	carry := false

	go func() {
		println("ready player 1")
		for {
			carry = <-or1
			println(" > carry: ", carry)
		}
	}()

	go func() {
		println("ready player 2")
		for {
			sum = <-xor2
			println(" > sum: ", sum)
		}
	}()

	c <- true
	a <- true
	b <- true

	time.Sleep(200 * time.Millisecond)
	println("\n-- Result --")
	println("sum  :", sum)
	println("carry:", carry)
}

func orGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	_a := false
	_b := false

	go func() {
		for {
			_a = <-a
			println("orGate (a)")
			r <- (_a || _b)
		}
	}()

	go func() {
		for {
			_b = <-b
			println("orGate (b)")
			r <- (_a || _b)
		}
	}()

	return r
}

func xorGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	_a := false
	_b := false

	go func() {
		for {
			_a = <-a
			println("xorGate (a)")
			r <- (_a != _b)
		}
	}()

	go func() {
		for {
			_b = <-b
			println("xorGate (b)")
			r <- (_a != _b)
		}
	}()

	return r
}

func andGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	_a := false
	_b := false

	go func() {
		for {
			_a = <-a
			println("andGate (a)")
			r <- (_a && _b)
		}
	}()

	go func() {
		for {
			_b = <-b
			println("andGate (b)")
			r <- (_a && _b)
		}
	}()

	return r
}

// https://stackoverflow.com/questions/16930251/go-one-producer-many-consumers
func fanOut(ch <-chan bool, num int) []chan bool {
	cs := make([]chan bool, num)

	for i := range cs {
		cs[i] = make(chan bool)
	}

	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}
		for _, c := range cs {
			// close all our fanOut channels when the input channel is exhausted.
			close(c)
		}
	}()

	return cs
}

func split(ch <-chan bool) (chan bool, chan bool) {
	r := fanOut(ch, 2)
	return r[0], r[1]
}
