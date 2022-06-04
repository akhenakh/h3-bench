package main

import h3 "github.com/akhenakh/goh3"

func main() {
	for i := 0; i < 1_000_000; i++ {
		h := h3.FromGeo(h3.GeoCoord{Latitude: 48.8, Longitude: 2.2}, 15)
		h3.ToGeo(h)
	}
}
