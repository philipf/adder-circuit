// not work
package main

func main() {

	a := make(chan bool)
	b := make(chan bool)
	xor1 := make(chan bool)

	c := make(chan bool)
	sum := make(chan bool)

	go xorGate(a, b, xor1)
	go xorGate(c, xor1, sum)

	a <- false
	b <- false

	c <- false

	println("result 2 : ", <-sum)

	close(sum)
}

func orGate(a, b <-chan bool, r chan<- bool) {
	for {
		// Important to read BOTH the values of the channel before proceeding
		// otherwise the boolean evaluation takes the first one that becomes available.
		_a := <-a
		_b := <-b

		r <- (_a || _b)
	}
}

func xorGate(a, b <-chan bool, r chan<- bool) {
	for {
		// Important to read BOTH the values of the channel before proceeding
		// otherwise the boolean evaluation takes the first one that becomes available.
		_a := <-a
		_b := <-b

		r <- (_a != _b)
	}
}

func andGate(a, b <-chan bool, r chan<- bool) {
	for {
		// Important to read BOTH the values of the channel before proceeding
		// otherwise the boolean evaluation takes the first one that becomes available.
		_a := <-a
		_b := <-b

		r <- (_a && _b)
	}
}
