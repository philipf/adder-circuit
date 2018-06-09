package circuits

import (
	"testing"
	"time"
)

func runTest(t *testing.T,
	inputA, inputB, inputC,
	expectedCarry, expectedSum bool) {

	a := make(chan bool, 1)
	defer close(a)

	b := make(chan bool, 1)
	defer close(b)

	cIn := make(chan bool, 1)
	defer close(cIn)

	done := make(chan bool)

	s, cOut := FullAdder(a, b, cIn)

	sum := false
	carry := false

	go func() {
		for {
			sum = <-s
			carry = <-cOut
		}
	}()

	go func() {
		a <- inputA
		b <- inputB
		cIn <- inputC
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		done <- true
	}()

	<-done

	if sum != expectedSum {
		t.Errorf("Expected sum to be %t", expectedSum)
	}

	if carry != expectedCarry {
		t.Errorf("Expected carry to be %t", expectedCarry)
	}

}

// Test without carry bit
func TestFullAdder_FalseFalseFalse(t *testing.T) {
	runTest(t, false, false, false, false, false)
}

func TestFullAdder_FalseTrueFalse(t *testing.T) {
	runTest(t, false, true, false, false, true)
}

func TestFullAdder_TrueFalseFalse(t *testing.T) {
	runTest(t, true, false, false, false, true)
}

func TestFullAdder_TrueTrueFalse(t *testing.T) {
	runTest(t, true, true, false, true, false)
}

// Test with carry bit
func TestFullAdder_FalseFalseTrue(t *testing.T) {
	runTest(t, false, false, true, false, true)
}

func TestFullAdder_FalseTrueTrue(t *testing.T) {
	runTest(t, false, true, true, true, false)
}

func TestFullAdder_TrueFalseTrue(t *testing.T) {
	runTest(t, true, false, true, true, false)
}

func TestFullAdder_TrueTrueTrue(t *testing.T) {
	runTest(t, true, true, true, true, true)
}
