package gates

// OrGate represents a logic OR gate
// https://en.wikipedia.org/wiki/OR_gate
func OrGate(a, b <-chan bool) <-chan bool {
	r := make(chan bool)

	_a := false
	_b := false

	go func() {
		for {
			_a = <-a
			r <- (_a || _b)
		}
	}()

	go func() {
		for {
			_b = <-b
			r <- (_a || _b)
		}
	}()

	return r
}
