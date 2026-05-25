// Package coordinates contains types and functions of coordinates for geodesy.
// The all types are based on [github.com/go-gl/mathgl/mgl64.Vec3], therefore you can convert each other.
package coordinates

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/wroge/wgs84"
)

// Geodetic is a type for geodetic coordinates.
// The order for geodetic points should be {Longitude, Latitude, Altitude}
type Geodetic mgl64.Vec3

// GeodeticReferenceSystem is the reference system for Geodetic.
// The default is WGS84.
var GeodeticReferenceSystem = wgs84.LonLat()

// GeodeticFromGeocentric converts geocentric coordinates to geodetic coordinates.
// A Geodetic object in the order of {X, Y, Z} should be entered and {Longitude, Latitude, Altitude}
// will be returned
func GeodeticFromGeocentric(geocentric Geocentric) (geodetic Geodetic) {
	geodetic[0], geodetic[1], geodetic[2] = GeocentricReferenceSystem.To(GeodeticReferenceSystem)(geocentric[0], geocentric[1], geocentric[2])
	return
}

// Longitude returns the longitude.
// The unit is degree.
func (geodetic *Geodetic) Longitude() *float64 {
	return &geodetic[0]
}

// Latitude returns the latitude.
// The unit is degree.
func (geodetic *Geodetic) Latitude() *float64 {
	return &geodetic[1]
}

// Altitude returns the altitude.
// The unit is meter.
func (geodetic *Geodetic) Altitude() *float64 {
	return &geodetic[2]
}

// GenerateLocalFromGeocentric generates a function to convert geocentric coordinates to local coordinates.
func (origin Geodetic) GenerateLocalFromGeocentric() func(geocentric Geocentric) mgl64.Vec3 {
	vector := mgl64.Vec3(GeocentricFromGeodetic(origin))
	γᵣ := vector.Len()
	ψ := math.Asin(vector.Z() / γᵣ)
	λ := origin[0] * math.Pi / 180.0
	rotator := mgl64.Rotate3DZ(-math.Pi / 2.0).Mul3(mgl64.Rotate3DY(ψ - math.Pi/2.0)).Mul3(mgl64.Rotate3DZ(-λ))

	return func(geocentric Geocentric) mgl64.Vec3 {
		local := rotator.Mul3x1(mgl64.Vec3(geocentric))
		local[2] -= γᵣ
		return local
	}
}

// GenerateGeocentricFromLocal generates a function to convert local coordinates to geocentric coordinates.
func (origin Geodetic) GenerateGeocentricFromLocal() func(local mgl64.Vec3) Geocentric {
	vector := mgl64.Vec3(GeocentricFromGeodetic(origin))
	γᵣ := vector.Len()
	ψ := math.Asin(vector.Z() / γᵣ)
	λ := origin[0] * math.Pi / 180.0
	rotator := mgl64.Rotate3DZ(λ).Mul3(mgl64.Rotate3DY(math.Pi/2.0 - ψ)).Mul3(mgl64.Rotate3DZ(math.Pi / 2.0))

	return func(local mgl64.Vec3) Geocentric {
		local[2] += γᵣ
		geocentric := Geocentric(rotator.Mul3x1(mgl64.Vec3(local)))
		return geocentric
	}
}
