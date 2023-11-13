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

func TestLinearRegression(t *testing.T) {
	X := linear.NewArrayMatrix(2, 2)
	X.Set(0, 0, 1)
	X.Set(1, 0, 0)
	X.Set(0, 1, 1)
	X.Set(1, 1, 2)

	y := linear.NewArrayVector(2)
	y.Set(0, 6)
	y.Set(1, 0)

	theta_hat := LinearRegre