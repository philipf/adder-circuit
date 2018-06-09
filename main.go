package main

import (
	"time"

	"github.com/philipf/adder-circuit/channels"
	"github.com/philipf/adder-circuit/gates"
)

func main() {
	testAll()
}

func testAll() {
	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)

	a0, a1 := channels.Split(a)
	b0, b1 := channels.Split(b)
	c0, c1 := channels.Split(c)

	xor1 := gates.XorGate(a0, b0)
	xor1_0, xor1_1 := channels.Split(xor1)

	xor2 := gates.XorGate(xor1_0, c0)
	and1 := gates.AndGate(a1, b1)
	and2 := gates.AndGate(xor1_1, c1)
	or1 := gates.OrGate(and1, and2)

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
			println(" >> sum: ", sum)
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
