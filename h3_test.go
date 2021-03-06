package h3bench

import (
	"fmt"
	"testing"

	h3 "github.com/akhenakh/goh3"
	"github.com/stretchr/testify/assert"
	ch3 "github.com/uber/h3-go/v3"
)

// validH3Index resolution 5
const (
	validCGOH3Index  = ch3.H3Index(0x850dab63fffffff)
	validCCGOH3Index = h3.H3Index(0x850dab63fffffff)
)

func TestCCGOFromGeo(t *testing.T) {
	geo := h3.GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	resolution := 9
	assert.Equal(t, fmt.Sprintf("%#x", h3.FromGeo(geo, resolution)), "0x8928308280fffff")
	// Output:
	// 0x8928308280fffff
}

func TestCCOFromGeo(t *testing.T) {
	geo := ch3.GeoCoord{
		Latitude:  37.775938728915946,
		Longitude: -122.41795063018799,
	}
	resolution := 9
	assert.Equal(t, fmt.Sprintf("%#x", ch3.FromGeo(geo, resolution)), "0x8928308280fffff")
	// Output:
	// 0x8928308280fffff
}

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
