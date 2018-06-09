package gates

// XorGate represents a logic XOR gate
// https://en.wikipedia.org/wiki/XOR_gate
func XorGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	_a := false
	_b := false

	go func() {
		for {
			_a = <-a
			r <- (_a != _b)
		}
	}()

	go func() {
		for {
			_b = <-b
			r <- (_a != _b)
		}
	}()

	return r
}
