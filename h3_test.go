package h3bench

import (
	"testing"

	h3 "github.com/akhenakh/goh3"
	ch3 "github.com/uber/h3-go/v3"
)

// validH3Index resolution 5
const (
	validCGOH3Index  = ch3.H3Index(0x850dab63fffffff)
	validCCGOH3Index = h3.H3Index(0x850dab63fffffff)
)

func BenchmarkToGeoCCGO(b *testing.B) {
	c := h3.NewBatch()
	for i := 0; i < b.N; i++ {
		c.ToGeo(validCCGOH3Index)
	}
}

func BenchmarkToGeoCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch3.ToGeo(validCGOH3Index)
	}
}

func BenchmarkFromToCCGO(b *testing.B) {
	c := h3.NewBatch()
	for i := 0; i < b.N; i++ {
		h := c.FromGeo(h3.GeoCoord{Latitude: 48.8, Longitude: 2.2}, 15)
		c.ToGeo(h)
	}
}

func BenchmarkFromToCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := ch3.FromGeo(ch3.GeoCoord{Latitude: 48.8, Longitude: 2.2}, 15)
		ch3.ToGeo(h)
	}
}
