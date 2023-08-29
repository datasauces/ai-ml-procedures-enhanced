package ml

import (
	"math"
	"testing"

	"github.com/ornerylawn/linear"
)

func ExpectFloat(expect, got float64, t *testing.T) {
	if math.Abs(got-expect) > 1e-9