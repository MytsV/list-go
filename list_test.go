package list

import (
	"testing"
	"math"
)

func TestLength(t *testing.T) {
	list := List{}
	if (list.Length() != int(math.NaN())) {
		t.Fatalf("...")
	}
}