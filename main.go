package main

import (
	"time"
)

func main() {
	// testOrGate()
	testAll()
}

func testAll() {

	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)

	fa := fanOut(a, 2)
	a0 := fa[0]
	a1 := fa[1]

	fb := fanOut(b, 2)
	b0 := fb[0]
	b1 := fb[1]

	fc := fanOut(c, 2)
	c0 := fc[0]
	c1 := fc[1]

	xor1 := xorGate(a0, b0)

	fxor1 := fanOut(xor1, 2)
	xor1_0 := fxor1[0]
	xor1_1 := fxor1[1]

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

	a <- true
	// b <- true
	// c <- true

	time.Sleep(200 * time.Millisecond)
	println("\n-- Result --")
	println("sum  : ", sum)
	println("carry:", carry)
}

func testOrGate() {
	a := make(chan bool)
	b := make(chan bool)
	// andGate := make(chan bool)

	fb := fanOut(b, 2)
	b0 := fb[0]
	b1 := fb[1]

	orGate1 := orGate(a, b0)
	andGate1 := andGate(orGate1, b1)
	//	andGate2 := andGate(andGate1, b)

	// go changeInput(a, b)
	a <- true
	b <- true

	go func() {
		for {
			result := <-andGate1
			println(result)
		}
	}()

	time.Sleep(5000 * time.Millisecond)

	println("done!!!")

}

func changeInput(a chan bool, b chan bool) {
	for {
		a <- false
		//b <- false
		time.Sleep(2 * time.Millisecond)
		println("--1---")

		a <- true
		//b <- false
		time.Sleep(2 * time.Millisecond)
		println("--2--")
	}

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
