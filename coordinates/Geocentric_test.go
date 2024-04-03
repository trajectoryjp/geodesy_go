package coordinates

import (
	"fmt"
	"math"
	"testing"

	"github.com/go-gl/mathgl/mgl64"
)

// Original point data source: https://www.ngs.noaa.gov/NCAT/
// (Note that we are not affiliated with NOAA and we do not claim copyright for their information we use.
// Please see the disclaimer https://www.ngs.noaa.gov/disclaimer.html for more detials.)

// var errorCriterion an error criterion of 1 meter
const errorCriterion float64 = 1

// var movingDistance is the distance in meters each point will be moved for tests
const movingDistance float64 = 1

func TestGeodeticFromGeocentric90(t *testing.T) {

	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geocentric{0, 6378137.000, 0}
	testGeodeticFromGeocentric(t, point90)
}

func TestGeodeticFromGeocentric45(t *testing.T) {

	// point at Lon:0 Lat: 45 Alt: 0
	point45 := Geocentric{4517590.879, 0.0, 4487348.409}
	testGeodeticFromGeocentric(t, point45)

}

func TestGeodeticFromGeocentric0(t *testing.T) {

	// point at Lon:0 Lat: 0 Alt: 0
	point0 := Geocentric{6378137.000, 0.0, 0.0}
	testGeodeticFromGeocentric(t, point0)

}

func TestGeocentricFromGeodetic90(t *testing.T) {

	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geocentric{0, 6378137.000, 0}
	testGeocentricFromGeodetic(t, point90)

}

func TestGeocentricFromGeodetic45(t *testing.T) {

	// point at Lon:0 Lat: 90 Alt: 0
	point45 := Geocentric{4517590.879, 0.0, 4487348.409}
	testGeocentricFromGeodetic(t, point45)

}

func TestGeocentricFromGeodetic0(t *testing.T) {

	// point at Lon:0 Lat: 0 Alt: 0
	point0 := Geocentric{6378137.000, 0.0, 0.0}
	testGeocentricFromGeodetic(t, point0)

}

func testGeocentricFromGeodetic(t *testing.T, point Geocentric) {

	// general pattern: geocentric (1) -> geodetic -> geocentric (2) then measure
	// the distance between geocentric 1 and 2.

	pointOrigin1 := point

	// convert to geodetic
	pointGeodetic := GeodeticFromGeocentric(pointOrigin1)

	// convert back to geodetic
	pointOrigin2 := GeocentricFromGeodetic(pointGeodetic)

	// measure distance between pointOrigin1 and pointOrigin2
	absoluteDifference := mgl64.Vec3(pointOrigin1).Sub(mgl64.Vec3(pointOrigin2))

	// the error value is the length between pointOrigin1 and pointOrigin2.
	errorVal := absoluteDifference.Len()

	if errorVal > errorCriterion {
		t.Fatalf("The error value (%2.16f) for on point %v is too large\n", errorVal, point)
	}

	fmt.Printf(
		`Origin1: (cartesian) %2.16f
Origin1: (geodetic) %2.16f
Origin2: (cartesian) %2.16f
Error: %2.16f
All error values are less than or equal to 1 (PASS)
`,
		pointOrigin1,
		pointGeodetic,
		pointOrigin2,
		errorVal,
	)

}

func testGeodeticFromGeocentric(t *testing.T, point Geocentric) {

	// make 3 sets of two cartesian (geocentric) points, the second of each moved by 1m in each
	// of x, y, and z directions
	pointOrigin := point
	pointOriginX := Geocentric{*pointOrigin.X() + movingDistance, *pointOrigin.Y(), *pointOrigin.Z()}

	pointOriginY := Geocentric{*pointOrigin.X(), *pointOrigin.Y() - movingDistance, *pointOrigin.Z()}

	pointOriginZ := Geocentric{*pointOrigin.X(), *pointOrigin.Y(), *pointOrigin.Z() + movingDistance}

	movedPoints := []Geocentric{pointOriginX, pointOriginY, pointOriginZ}

	pointOriginGeodetic := GeodeticFromGeocentric(pointOrigin)
	pointOriginSpherical := Spherical(pointOriginGeodetic)

	pointsGeodetic := []Geodetic{}
	distances := []float64{}
	errorVals := []float64{}

	// convert each pair to lat/lon (geodetic)
	for _, point := range movedPoints {

		// convert to geodetic
		pointGeodetic := GeodeticFromGeocentric(point)

		pointsGeodetic = append(pointsGeodetic, pointGeodetic)

		// save in sphereical struct
		pointSpherical := Spherical(pointGeodetic)

		// use .GetLengthTo() to measure distance and append to distances
		distance := pointOriginSpherical.GetLengthTo(pointSpherical)

		distances = append(distances, distance)

		// check that the error value is less than errorCriterion meters. We expect the difference
		// in distance to be movingDistance meter, so subtract movingDistance from distance
		errorVal := math.Abs(distance - movingDistance)
		if errorVal > errorCriterion {
			t.Fatalf("The error value for on point %v is too large\n", point)
		}

		errorVals = append(errorVals, errorVal)

	}

	fmt.Printf(
		` Origin: (cartesian) %2.16f (geodetic) %2.16f, 
ShiftX+1: (cartesian) %2.16f (geodetic) %2.16f,
ShiftY-1: (cartesian) %2.16f (geodetic) %2.16f,
ShiftZ+1: (cartesian) %2.16f (geodetic) %2.16f,
Distance 1 (delta x): %2.16f
Distance 2 (delta y): %2.16f
Distance 3 (delta z): %2.16f
Error X: %2.16f
Error Y: %2.16f
Error Z: %2.16f
All error values are less than or equal to 1 (PASS)
`,
		pointOrigin, pointOriginGeodetic,
		pointOriginX, pointsGeodetic[0],
		pointOriginY, pointsGeodetic[1],
		pointOriginZ, pointsGeodetic[2],
		distances[0], distances[1], distances[2],
		errorVals[0], errorVals[1], errorVals[2],
	)
}
