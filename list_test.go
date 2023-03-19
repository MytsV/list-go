package list

import (
	"math"
	"testing"
)

func TestLength(t *testing.T) {
	list := List{}
	if list.Length() != int(math.NaN()) {
		t.Fatalf("...")
	}
}
