package gates

import "testing"

func TestXorInputFalseFalse(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := XorGate(a, b)

	a <- false
	Assert(t, r, false)

	b <- false
	Assert(t, r, false)
}

func TestXorInputTrueTrue(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := XorGate(a, b)

	a <- true
	Assert(t, r, true)

	b <- true
	Assert(t, r, false)
}

func TestXorInputFalseTrue(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := XorGate(a, b)

	a <- false
	Assert(t, r, false)

	b <- true
	Assert(t, r, true)
}

func TestXorInputTrueFalse(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := XorGate(a, b)

	a <- true
	Assert(t, r, true)

	b <- false
	Assert(t, r, true)
}
