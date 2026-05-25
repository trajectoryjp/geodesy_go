package coordinates

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/wroge/wgs84"
)

// Geocentric is a type for geocentric coordinates. The {X, Y, Z} order should be followed.
type Geocentric mgl64.Vec3

// GeocentricReferenceSystem is the reference system for Geocentric.
var GeocentricReferenceSystem = wgs84.GeocentricReferenceSystem{}

// GeocentricFromGeodetic converts geodetic coordinates to geocentric coordinates.
// The order for the point to convert should be {Longitude, Latitude, Altitude} and
// {X, Y, Z} will be returned.
func GeocentricFromGeodetic(geodetic Geodetic) (geocentric Geocentric) {
	geocentric[0], geocentric[1], geocentric[2] = GeodeticReferenceSystem.To(GeocentricReferenceSystem)(geodetic[0], geodetic[1], geodetic[2])
	return
}

// X returns the x coordinate.
// The unit is meter.
func (geocentric *Geocentric) X() *float64 {
	return &geocentric[0]
}

// Y returns the y coordinate.
// The unit is meter.
func (geocentric *Geocentric) Y() *float64 {
	return &geocentric[1]
}

// Z returns the z coordinate.
// The unit is meter.
func (geocentric *Geocentric) Z() *float64 {
	return &geocentric[2]
}
