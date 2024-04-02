# Coordinates Module

The coordinates module provides structs and functions necessary for conversion between Geocentric (x, y z), Geodetic (Longitude, Latitude, Altitude), and Spherical coordinates and related measurements.

## External Data and License Notice

This guide will use geodetic coordinates from [OpenStreetMap](openstreetmap.org/copyright); please note that the data is available under the Open Database License. Please refer to the link for more information. The following code is not part of OpenStreeMap; it's provided by Trajectory under the MIT license. Please feel free to learn and contribute!

Our Coordinates module makes use of two amazing open source packages, please take a look at our source code for more information. Thanks!

## Definitions


## Converting between Geodetic and Geocentric

Support for converstion between Geodetic (Longitude, Latitude, Altitude) and Geocentric (X, Y, Z) is provided by `GeodeticFromGeocentric()` and `GeocentricFromGeodetic()`.

Let's take a point on the [Sumida River](https://en.wikipedia.org/wiki/Sumida_River), one of my favorite spots in Tokyo. We'll put our coordinates in the `Geodetic` struct in the following order: Longitude, Latitude, Altitude.

```go
// import packages
import (
	"fmt"
	"math"
	"github.com/go-gl/mathgl/mgl64"
    "github.com/wroge/wgs84"
    "github.com/trajectoryjp/geodesy_go/coordinates"

)

sumida := Geodetic{
	139.78645,	// longitude
	35.67254,	// latitude
	0, 			// altitude
}

```

We simply use `GeodeticFromGeocentric()` to convert to X, Y, Z format.

```go
sumidaGeocentric := GeocentricFromGeodetic(sumida)
// sumidaGeocentric: [-3961225.0964641897007823 3349100.8647583685815334 3698736.2743638330139220]
```

Converting back to Geodetic from Geocentric is just as simple.

```go
sumidaGeodetic := GeodeticFromGeocentric(sumidaGeocentric)
// sumidaGeodetic: [139.7864500000000305 35.6725400000000050 0.0000000000000000]
```

## Converting To Spherical and Measuring Distance

The `Spherical` type provides format for spherical systems. We have to convert to Spherical to measure the "as the bird flys" distance on the surface of the earth.

Let's add another point down the river, convert both to spherical, and measure the distance between the two.

Our module provides `GetLengthTo()`, which is a method on a spherical point that requires a second spherical point as an input.

```go
tokyoBay := Geodetic{
	139.76829,
	35.64328,
	0,
}

// convert both to Spherical
sumidaSpherical := Spherical(sumida)
tokyoBaySpherical := Spherical(tokyoBay)

// calculate the distance in meters
distance := sumidaSpherical.GetLengthTo(tokyoBaySpherical)

// sumidaSpherical: [139.7864500000000021 35.6725399999999979 0.0000000000000000],
// tokyoBaySpherical: [139.7682900000000075 35.6432799999999972 0.0000000000000000],
// distance: 3647.9246540185599770
```

## Converting to Local from Geocentric

Let's convert `sumida` to a local point. Here's how we'll do the conversion:

1. Start with a point in `Geodetic{}` form.
2. Convert the geodetic point to geocentric.
3. Create a "geocentric -> local" conversion function using the `GenerateLocalFromGeocentric()` method of the geodetic point.
4. Use the conversion function from 3 to convert the geocentric point to local.

We already completed steps 1 and 2 above, but let's regenerate here anyway in case you're joining part way through.

```go
// declare sumida as Geodetic{}
	sumida := Geodetic{
		139.78645, // longitude
		35.67254,  // latitude
		0,         // altitude
	}

// convert to geocentric
sumidaGeocentric := GeocentricFromGeodetic(sumida)

```

Next we convert use `sumida.GenerateLocalFromGeocentric()` method to create a function that will allow us to convert from Geocentric points to Local Points.

```go
ToLocalFunction := sumida.GenerateLocalFromGeocentric()
```

In step 4, we use this function to convert **geocentric** points to local. Notice, as with `tokyoBayLocal`, we can convert to geocentric directly inside the function call.

```go
sumidaLocal := ToLocalFunction(sumidaGeocentric)
// sumidaLocal: [-0.0000000006500660 -0.0000000013969839 0.0000000000000000]
tokyoBayLocal := ToLocalFunction(Geocentric(tokyoBay))
// tokyoBayLocal: [-117.4585155050558853 48.6053239105692398 -6370973.1957182800397277]
```

## Finding Distance between two local points

If we want to find the distance between two local points, we first create a difference vector using `mgl64.Vec3()` on the local point. This library has a `.Sub()` method that creates a resulting difference vector from another `mgl64.Vec3` point.

```go
differenceVector := mgl64.Vec3(sumidaLocal).Sub(mgl64.Vec3(tokyoBayLocal))
// differenceVector: [1644.7146617341784349 3246.3099505109712481 -9.2779252650216222
```

Then we find the length of this vector

```go
distance2 := differenceVector.Len()
// distance2: 3639.1895654399677369
```

If we compare it from the results from `.GetLengthTo()` (from above in `distance`), the two distances are quite similar.

```go
difference := math.Abs(distance2 - distance)
// difference: 8.7350885785922401
```
