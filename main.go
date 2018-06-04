package main

func main() {

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

	c <- true
	a <- true
	b <- false

	// println("xor1: ", <-xor1)
	//println("a1:", <-a1)
	// println("b1:", <-b1)

	println("carry: ", <-or1)
	println("sum: ", <-xor2)

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
			close(o)
		}
	}()

	return outChannels
}
