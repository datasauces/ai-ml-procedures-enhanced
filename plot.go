
package ml

import (
	"image"
	"image/color"
	"math"

	"github.com/ornerylawn/linear"
)

type Color struct {
	R, G, B, A float64 // not "pre-multiplied"
}

// RGBA returns the "pre-multiplied" 8-bit per channel color.
func (c Color) RGBA() (r, g, b, a uint32) {
	return uint32(c.R * c.A * 255.0),
		uint32(c.G * c.A * 255.0),
		uint32(c.B * c.A * 255.0),
		uint32(c.A * 255.0)
}

func ColorFromRaster(c color.Color) Color {
	r, g, b, a := c.RGBA() // "pre-multiplied"
	alpha := float64(a) / 255.0
	return Color{
		(float64(r) / alpha) / 255.0,
		(float64(g) / alpha) / 255.0,
		(float64(b) / alpha) / 255.0,
		alpha,
	}
}

type Point struct {
	X, Y float64
}

type Canvas interface {
	SetStrokeColor(color Color)
	SetFillColor(color Color)