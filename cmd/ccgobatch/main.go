package main

import h3 "github.com/akhenakh/goh3"

func main() {
	c := h3.NewBatch()
	for i := 0; i < 1_000_000; i++ {
		h := c.FromGeo(h3.GeoCoord{Latitude: 48.8, Longitude: 2.2}, 15)
		c.ToGeo(h)
	}
}
