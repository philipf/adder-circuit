package gates

// AndGate represents a logic AND gate
// https://en.wikipedia.org/wiki/AND_gate
func AndGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	_a := false
	_b := false

	go func() {
		for {
			_a = <-a
			r <- (_a && _b)
		}
	}()

	go func() {
		for {
			_b = <-b
			r <- (_a && _b)
		}
	}()

	return r
}
