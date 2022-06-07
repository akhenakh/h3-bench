package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	h3 "github.com/akhenakh/goh3"
)

func main() {
	var resolution = flag.Int("resolution", 15, "resolution 0-15 inclusive")
	flag.Parse()

	fileScanner := bufio.NewScanner(os.Stdin)

	fileScanner.Split(bufio.ScanLines)
	g := h3.GeoCoord{}
	for fileScanner.Scan() {
		coords := strings.Split(fileScanner.Text(), " ")
		lat, err := strconv.ParseFloat(coords[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		g.Latitude = lat
		lng, err := strconv.ParseFloat(coords[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		g.Longitude = lng

		h := h3.FromGeo(g, *resolution)
		fmt.Printf("%#x\n", h)
	}
}
