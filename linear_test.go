package ml

import (
	"math"
	"testing"

	"github.com/ornerylawn/linear"
)

func ExpectFloat(expect, got float64, t *testing.T) {
	if math.Abs(got-expect) > 1e-9 {
		t.Errorf("expected %f but got %f", expect, got)
	}
}

func ExpectInt(expect, got int, t *testing.T) {
	if got != expect {
		t.Errorf("expected %d but got %d", expect, got)
	}
}

func TestLinearRegression(t *tes