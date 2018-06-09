package gates

import "testing"

// Assert for unit tests
func Assert(t *testing.T, r <-chan bool, expectedResult bool) {
	result := <-r

	if result != expectedResult {
		t.Errorf("Expected %t but actual was %t", expectedResult, result)
	}
}
