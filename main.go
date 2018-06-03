// not work
package main

func main() {

	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)

	xor1 := xorGate(a, b)
	sum := orGate(xor1, c)

	a <- true
	b <- true
	c <- true

	println("sum: ", <-sum)

}

func orGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	go func() {
		for {
			// Important to read BOTH the values of the channel before proceeding
			// otherwise the boolean evaluation takes the first one that becomes available.
			_a := <-a
			_b := <-b

			r <- (_a || _b)
		}
	}()

	return r
}

func xorGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	go func() {
		for {
			_a := <-a
			_b := <-b

			r <- (_a != _b)
		}
		// close(r)
	}()

	return r
}

func andGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	go func() {
		for {
			_a := <-a
			_b := <-b

			r <- (_a && _b)
		}
	}()

	return r
}
