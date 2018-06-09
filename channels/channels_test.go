package channels

import (
	"testing"
)

func TestFanoutWithOneOutput(t *testing.T) {
	c := make(chan bool)

	output := FanOut(c, 1)

	if len(output) != 1 {
		t.Errorf("Expected channels %d, actual %d", 1, len(output))
	}

	go func() {
		c <- true
		c <- false
		c <- true
		close(c)
	}()

	if !<-output[0] {
		t.Error("Expected true")
	}

	if <-output[0] {
		t.Error("Expected false")
	}

	if !<-output[0] {
		t.Error("Expected true")
	}
}

func TestFanoutWithTwoOutputs(t *testing.T) {
	c := make(chan bool)
	defer close(c)

	done := make(chan bool)
	defer close(done)

	output := FanOut(c, 2)

	if len(output) != 2 {
		t.Errorf("Expected channels %d, actual %d", 2, len(output))
	}

	go func() {
		c <- true
		c <- false
		done <- true
	}()

	if !<-output[0] {
		t.Error("Expected true")
	}

	if !<-output[1] {
		t.Error("Expected true")
	}

	if <-output[0] {
		t.Error("Expected false")
	}

	if <-output[1] {
		t.Error("Expected false")
	}

	<-done

	// if <-output[1] {
	// 	t.Error("Expected false")
	// }

}
