package coordinates

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// Spherical is a type for spherical coordinates.
// This is for calculating [Great-circle navigation].
//
// [Great-circle navigation]: https://en.wikipedia.org/wiki/Great-circle_navigation
type Spherical mgl64.Vec3

// Longitude returns the longitude.
// The unit is degree.
func (spherical *Spherical) Longitude() *float64 {
	return &spherical[0]
}

// Latitude returns the latitude.
// The unit is degree.
func (spherical *Spherical) Latitude() *float64 {
	return &spherical[1]
}

// Altitude returns the altitude.
// The unit is meter.
func (spherical *Spherical) Altitude() *float64 {
	return &spherical[2]
}

// GetLengthTo returns the length to the arrival.
// The unit is meter.
func (departure Spherical) GetLengthTo(arrival Spherical) float64 {
	greatCircleDistance := departure.GetGreatCircleDistanceTo(arrival)
	altitudeDifference := *arrival.Altitude() - *departure.Altitude()
	return math.Sqrt(greatCircleDistance*greatCircleDistance + altitudeDifference*altitudeDifference)
}

// GetGreatCircleDistanceTo returns the great circle distance to the arrival.
// The unit is meter.
func (departure Spherical) GetGreatCircleDistanceTo(arrival Spherical) float64 {
	φDeparture := *departure.Latitude() * math.Pi / 180.0
	φArrival := *arrival.Latitude() * math.Pi / 180.0
	λDeparture := *departure.Longitude() * math.Pi / 180.0
	λArrival := *arrival.Longitude() * math.Pi / 180.0
	λDefference := λArrival - λDeparture

	sinφDeparture := math.Sin(φDeparture)
	sinφArrival := math.Sin(φArrival)
	cosφDeparture := math.Cos(φDeparture)
	cosφArrival := math.Cos(φArrival)
	sinλDefference := math.Sin(λDefference)
	cosλDefference := math.Cos(λDefference)

	b := cosφDeparture*sinφArrival - sinφDeparture*cosφArrival*cosλDefference
	c := cosφArrival * sinλDefference
	d := sinφDeparture*sinφArrival + cosφDeparture*cosφArrival*cosλDefference
	tanσDifference := math.Sqrt(b*b + c*c) / d
	σDifference := math.Atan(tanσDifference)
	return GeodeticReferenceSystem.Datum.A() * σDifference
}

// GetDirectionTo returns the direction to the arrival.
// The unit is degree.
func (departure Spherical) GetDirectionTo(arrival Spherical) float64 {
	φDeparture := *departure.Latitude() * math.Pi / 180.0
	φArrival := *arrival.Latitude() * math.Pi / 180.0
	λDeparture := *departure.Longitude() * math.Pi / 180.0
	λArrival := *arrival.Longitude() * math.Pi / 180.0
	λDefference := λArrival - λDeparture

	sinφDeparture := math.Sin(φDeparture)
	sinφArrival := math.Sin(φArrival)
	cosφDeparture := math.Cos(φDeparture)
	cosφArrival := math.Cos(φArrival)
	sinλDefference := math.Sin(λDefference)
	cosλDefference := math.Cos(λDefference)

	tanα := cosφArrival*sinλDefference / (cosφDeparture*sinφArrival - sinφDeparture*cosφArrival*cosλDefference)
	return math.Atan(tanα) * 180.0 / math.Pi
}

// GetDirectionOnEquator returns the direction on the equator.
// The unit is degree.
func (spherical Spherical) GetDirectionOnEquator(direction float64) float64 {
	α := direction * math.Pi / 180.0
	φ := *spherical.Latitude() * math.Pi / 180.0

	sinα := math.Sin(α)
	cosα := math.Cos(α)
	sinφ := math.Sin(φ)
	cosφ := math.Cos(φ)

	tanα0 := sinα*cosφ / math.Sqrt(cosα*cosα+sinα*sinα*sinφ*sinφ)
	return math.Atan(tanα0) * 180.0 / math.Pi
}
