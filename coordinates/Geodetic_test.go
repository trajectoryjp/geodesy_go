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

func TestGenerateLocalFromGeocentric90(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geodetic{0, 90, 0}
	testGenerateLocalFromGeocentric(t, point90)
}

func TestGenerateLocalFromGeocentric45(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geodetic{0, 45, 0}
	testGenerateLocalFromGeocentric(t, point90)
}

func TestGenerateLocalFromGeocentric0(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := Geodetic{0, 0, 0}
	testGenerateLocalFromGeocentric(t, point90)
}

func TestGenerateGeocentricFromLocal90(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := mgl64.Vec3{0, 6378137.000, 0}
	testGenerateGeocentricFromLocal(t, point90)
}

func TestGenerateGeocentricFromLocal45(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := mgl64.Vec3{4517590.879, 0.0, 4487348.409}
	testGenerateGeocentricFromLocal(t, point90)
}

func TestGenerateGeocentricFromLocal0(t *testing.T) {
	// point at Lon:0 Lat: 90 Alt: 0
	point90 := mgl64.Vec3{6378137.000, 0.0, 0.0}
	testGenerateGeocentricFromLocal(t, point90)
}

func testGenerateLocalFromGeocentric(t *testing.T, point Geodetic) {

	// general pattern: geocentric1 -> local1 ; geocentric2 -> local2 ; compare local1 and local2
	// ToLocalFunction takes the "world" conversion parameters from point (geodetic)
	// to local and then saves these conversion parameters in a function to apply
	// to another geodetic point.

	// generate a moved point in geodetic
	pointCartesian := GeocentricFromGeodetic(point)
	movedPointCartesian := Geocentric{*pointCartesian.X(), *pointCartesian.Y() - movingDistance, *pointCartesian.Z()}
	movedPoint := GeodeticFromGeocentric(movedPointCartesian)

	// generate a conversion function
	ToLocalFunction := point.GenerateLocalFromGeocentric()

	// convert point to local
	local := ToLocalFunction(pointCartesian)

	// convert movedPoint to local using same function
	localMovedPoint := ToLocalFunction(movedPointCartesian)

	// measure distance between point and movedPoint. Subtract movingDistance to get errorVal
	absoluteDifference := mgl64.Vec3(local).Sub(mgl64.Vec3(localMovedPoint))
	errorVal := math.Abs(absoluteDifference.Len() - movingDistance)

	if errorVal > errorCriterion {
		t.Fatalf("The error value (%2.16f) for on point %v is too large\n", errorVal, point)
	}

	fmt.Printf(
		`Original Point: (cartesian) %2.16f
Moved Point: (cartesian) %2.16f
Original Point to Local: %2.16f
Moved Point to Local: %2.16f
Error: %2.16f
All error values are less than or equal to 1 (PASS)
`,
		point,
		movedPoint,
		local,
		localMovedPoint,
		errorVal,
	)
}

func testGenerateGeocentricFromLocal(t *testing.T, point mgl64.Vec3) {

	// general pattern: local -> geocentric1 ; local2 -> geocentric2 ; compare geocentric1 and geocentric2

	// generate a moved local point
	movedPoint := mgl64.Vec3{point.X(), point.Y() - movingDistance, point.Z()}

	// convert both local points to geocentric
	pointGeodetic := GeodeticFromGeocentric(Geocentric{point.X(), point.Y(), point.Z()})
	//movedPointGeocentric := GeodeticFromGeocentric(movedPoint)

	// generate a conversion function (local -> geocentric)
	ToGeocentricFunction := pointGeodetic.GenerateGeocentricFromLocal()

	// convert local points to geocentric using same function
	geocentric := ToGeocentricFunction(point)
	movedGeocentric := ToGeocentricFunction(movedPoint)

	// measure distance between point and movedPoint. Subtract movingDistance to get errorVal
	absoluteDifference := mgl64.Vec3(geocentric).Sub(mgl64.Vec3(movedGeocentric))
	errorVal := math.Abs(absoluteDifference.Len() - movingDistance)

	if errorVal > errorCriterion {
		t.Fatalf("The error value (%2.16f) for on point %v is too large\n", errorVal, point)
	}

	fmt.Printf(
		`Original Point: (cartesian) %2.16f
Moved Point: (cartesian) %2.16f
Original Point to Local: %2.16f
Moved Point to Local: %2.16f
Error: %2.16f
All error values are less than or equal to 1 (PASS)
`,
		point,
		movedPoint,
		geocentric,
		movedGeocentric,
		errorVal,
	)
}
