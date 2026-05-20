# geodesy

[![Go Version](https://img.shields.io/badge/go-1.20%2B-00ADD8)](https://go.dev)
[![Build](https://github.com/trajectoryjp/geodesy_go/actions/workflows/go.yml/badge.svg)](https://github.com/trajectoryjp/geodesy_go/actions/workflows/go.yml)
[![Coverage](https://codecov.io/gh/trajectoryjp/geodesy_go/branch/main/graph/badge.svg)](https://codecov.io/gh/trajectoryjp/geodesy_go)
[![License](https://img.shields.io/github/license/trajectoryjp/geodesy_go)](LICENSE)
[![Release](https://img.shields.io/github/v/release/trajectoryjp/geodesy_go)](https://github.com/trajectoryjp/geodesy_go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/trajectoryjp/geodesy_go)](https://goreportcard.com/report/github.com/trajectoryjp/geodesy_go)
[![Docs](https://pkg.go.dev/badge/github.com/trajectoryjp/geodesy_go.svg)](https://pkg.go.dev/github.com/trajectoryjp/geodesy_go)

Go module geodesy provides the coordinate system conversion and the distance calculation.
The default coordinate system is WGS84.
This wraps [github.com/wroge/wgs84](https://github.com/wroge/wgs84) with object orientation,
therefore you can change the coordinate system with its system.

## Prerequisites

- **[Go](https://go.dev)**: **1.20+**.

## Installation

Simply append the following import to your code, and then `go [build|run|test]`
will automatically fetch the necessary dependencies:

```go
import "github.com/trajectoryjp/geodesy_go"
```

## Contribution

You are very welcome to:

- Create pull requests of any kind
- Let me know if you are using this library and find it useful
- Open issues with request for support because they will help you and many others
