package gates

import "testing"

func TestAndInputFalseFalse(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := AndGate(a, b)

	a <- false
	Assert(t, r, false)

	b <- false
	Assert(t, r, false)
}

func TestAndInputTrueTrue(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := AndGate(a, b)

	a <- true
	Assert(t, r, false)

	b <- true
	Assert(t, r, true)
}

func TestAndInputFalseTrue(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := AndGate(a, b)

	a <- false
	Assert(t, r, false)

	b <- true
	Assert(t, r, false)
}

func TestAndInputTrueFalse(t *testing.T) {
	a := make(chan bool)
	b := make(chan bool)

	r := AndGate(a, b)

	a <- true
	Assert(t, r, false)

	b <- false
	Assert(t, r, false)
}
