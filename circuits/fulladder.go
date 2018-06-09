package circuits

import (
	"github.com/philipf/adder-circuit/channels"
	"github.com/philipf/adder-circuit/gates"
)

// FullAdder builds a full adder circuit
// https://en.wikipedia.org/wiki/Adder_(electronics)
func FullAdder(a, b, c <-chan bool) (sum, carry <-chan bool) {

	a0, a1 := channels.Split(a)
	b0, b1 := channels.Split(b)
	c0, c1 := channels.Split(c)

	xor1 := gates.XorGate(a0, b0)
	xor1_0, xor1_1 := channels.Split(xor1)

	xor2 := gates.XorGate(xor1_0, c0)
	and1 := gates.AndGate(a1, b1)
	and2 := gates.AndGate(xor1_1, c1)

	or1 := gates.OrGate(and1, and2)

	return xor2, or1
}
