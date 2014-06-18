package distance

import (
	"fmt"
)

type Distance float64

const (
	Meter      Distance = 1
	Centimeter          = Meter / 100
	Millimeter          = Centimeter / 10
	Kilometer           = 1000 * Meter
	Mile                = 1.609344 * Kilometer
	Yard                = Mile / 1760
	Foot                = Yard / 3
	Inch                = Foot / 12
)

// String outputs a readable version of the distance in meters
func (d Distance) String() string {
	return fmt.Sprintf("%v", d.Meters()) + "m"
}

func (d Distance) Meters() float64 {
	return float64(d)
}

func (d Distance) Centimeters() float64 {
	return float64(d / Centimeter)
}

func (d Distance) Millimeters() float64 {
	return float64(d / Millimeter)
}

func (d Distance) Kilometers() float64 {
	return float64(d / Kilometer)
}

func (d Distance) Miles() float64 {
	return float64(d / Mile)
}

func (d Distance) Inches() float64 {
	return float64(d / Inch)
}

func (d Distance) Yards() float64 {
	return float64(d / Yard)
}

func (d Distance) Feet() float64 {
	return float64(d / Foot)
}
