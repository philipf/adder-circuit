package main

func main() {

	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)

	xor1 := xorGate(a, b)
	sum := xorGate(xor1, c)
	fo := fanOut(sum, 2)

	a <- true
	b <- true
	c <- true

	println("sum1: ", <-fo[0])
	println("sum2: ", <-fo[1])

}

func orGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	go func() {
		// Important to read BOTH the values of the channel before proceeding
		// otherwise the boolean evaluation takes the first one that becomes available.
		_a := <-a
		_b := <-b

		r <- (_a || _b)
		close(r)
	}()

	return r
}

func xorGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	go func() {
		_a := <-a
		_b := <-b

		r <- (_a != _b)
		close(r)
	}()

	return r
}

func andGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	go func() {
		_a := <-a
		_b := <-b

		r <- (_a && _b)
		close(r)
	}()

	return r
}

func fanOut(c <-chan bool, num int) []chan bool {
	outChannels := make([]chan bool, num)

	for i := range outChannels {
		outChannels[i] = make(chan bool)
	}

	go func() {
		b := <-c

		for _, o := range outChannels {
			o <- b
		}
	}()

	return outChannels
}
