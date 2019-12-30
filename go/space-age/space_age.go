package space

import (
	"math"
)

// Planet custom type
type Planet string

const earthSecondsPerYear = 31557600
const mercuryToEarthRate = 0.2408467 * earthSecondsPerYear
const venusToEarthRate = 0.61519726 * earthSecondsPerYear
const marsToEarthRate = 1.8808158 * earthSecondsPerYear
const jupiterToEarthRate = 11.862615 * earthSecondsPerYear
const saturnToEarthRate = 29.447498 * earthSecondsPerYear
const uranusToEarthRate = 84.016846 * earthSecondsPerYear
const neptuneToEarthRate = 164.79132 * earthSecondsPerYear

var rateLookup = map[Planet]float64{
	"Earth":   earthSecondsPerYear,
	"Mercury": mercuryToEarthRate,
	"Venus":   venusToEarthRate,
	"Mars":    marsToEarthRate,
	"Jupiter": jupiterToEarthRate,
	"Saturn":  saturnToEarthRate,
	"Uranus":  uranusToEarthRate,
	"Neptune": neptuneToEarthRate,
}

// Age function to calculate ages on planets
func Age(seconds float64, planet Planet) float64 {
	return math.Round(seconds/rateLookup[planet]*100) / 100
}
