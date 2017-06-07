package even

import "testing"

func TestEven(t *testing.T) {
	i := 5
	if !Even(i) {
		t.Log("%d is not even!", i)
		t.Fail()
	}
}
