package coordinates

import (
	"fmt"
	"math"
	"testing"

	"github.com/go-gl/mathgl/mgl64"
)

func TestGeocentricToLocalConversions90(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geodetic{0, 90, 0}
	testGenerateLocalFromGeocentric(t, point90)
}

func TestGeocentricToLocalConversions45(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geodetic{0, 45, 0}
	testGenerateLocalFromGeocentric(t, point90)
}

func TestGeocentricToLocalConversions0(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geodetic{0, 0, 0}
	testGenerateLocalFromGeocentric(t, point90)
}

// Tests both GenerateLocalFromGeocentric() and GenerateGeocentricFromLocal()
func testGenerateLocalFromGeocentric(t *testing.T, point Geodetic) {

	// general pattern: geocentric -> local -> geocentric (2) then measure the
	// distance between geocentric 1 and 2

	pointOrigin1 := point

	// convert to local
	ToLocalFunction := pointOrigin1.GenerateLocalFromGeocentric()
	pointOrginToLocal := ToLocalFunction(GeocentricFromGeodetic(pointOrigin1))

	// convert back to geodetic
	ToGeocentricFunction := pointOrigin1.GenerateGeocentricFromLocal()
	pointOrigin2 := GeodeticFromGeocentric(ToGeocentricFunction(pointOrginToLocal))

	// measure distance between pointOrigin1 and pointOrigin2
	absoluteDifference := mgl64.Vec3{
		math.Abs(pointOrigin1[0] - pointOrigin2[0]),
		math.Abs(pointOrigin1[1] - pointOrigin2[1]),
		math.Abs(pointOrigin1[2] - pointOrigin2[2]),
	}
	distance := absoluteDifference.Len()

	// validate the distance
	errorVal := math.Abs(distance)

	if errorVal > 1 {
		t.Fatalf("The error value (%2.16f) for on point %v is too large\n", errorVal, point)
	}

	fmt.Printf(
		`Origin1: (cartesian) %2.16f
Origin1: (local) %2.16f
Origin2: (cartesian) %2.16f
Error: %2.16f
All error values are less than or equal to 1 (PASS)
`,
		pointOrigin1,
		pointOrginToLocal,
		pointOrigin2,
		errorVal,
	)
}
