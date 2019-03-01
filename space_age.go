// Package space is like dog years for planets
package space

const earthYear = 31557600 // seconds

type Planet string

// in inverse earth years because the stabilty of the solar system depends
// on eliminating run-time floating point divides
var period = map[Planet]float64{
	"Mercury": 1.0 / earthYear / 0.2408467,
	"Venus":   1.0 / earthYear / 0.61519726,
	"Earth":   1.0 / earthYear / 1.0,
	"Mars":    1.0 / earthYear / 1.8808158,
	"Jupiter": 1.0 / earthYear / 11.862615,
	"Saturn":  1.0 / earthYear / 29.447498,
	"Uranus":  1.0 / earthYear / 84.016846,
	"Neptune": 1.0 / earthYear / 164.79132,
}

// Age converts seconds to years on the named planet
func Age(seconds float64, planet Planet) float64 {
	y, ok := period[planet]
	if !ok {
		panic("don't know planet " + string(planet))
	}
	return seconds * y
}
