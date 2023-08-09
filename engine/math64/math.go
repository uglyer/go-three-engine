// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package math64 implements basic math functions which operate
// directly on float64 numbers without casting and contains
// types of common entities used in 3D Graphics such as vectors,
// matrices, quaternions and others.
package math64

import (
	"math"
)

const Pi = math.Pi
const degreeToRadiansFactor = math.Pi / 180
const radianToDegreesFactor = 180.0 / math.Pi

var Infinity = math.Inf(1)

// DegToRad converts a number from degrees to radians
func DegToRad(degrees float64) float64 {

	return degrees * degreeToRadiansFactor
}

// RadToDeg converts a number from radians to degrees
func RadToDeg(radians float64) float64 {

	return radians * radianToDegreesFactor
}

// Clamp clamps x to the provided closed interval [a, b]
func Clamp(x, a, b float64) float64 {

	if x < a {
		return a
	}
	if x > b {
		return b
	}
	return x
}

// ClampInt clamps x to the provided closed interval [a, b]
func ClampInt(x, a, b int) int {

	if x < a {
		return a
	}
	if x > b {
		return b
	}
	return x
}

func Abs(v float64) float64 {
	return math.Abs(v)
}

func Acos(v float64) float64 {
	return math.Acos(v)
}

func Asin(v float64) float64 {
	return math.Asin(v)
}

func Atan(v float64) float64 {
	return math.Atan(v)
}

func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

func Ceil(v float64) float64 {
	return math.Ceil(v)
}

func Cos(v float64) float64 {
	return math.Cos(v)
}

func Floor(v float64) float64 {
	return math.Floor(v)
}

func Inf(sign int) float64 {
	return math.Inf(sign)
}

func IsNaN(v float64) bool {
	return math.IsNaN(v)
}

func Sin(v float64) float64 {
	return math.Sin(v)
}

func Sqrt(v float64) float64 {
	return math.Sqrt(v)
}

func Max(a, b float64) float64 {
	return math.Max(a, b)
}

func Min(a, b float64) float64 {
	return math.Min(a, b)
}

func Mod(a, b float64) float64 {
	return math.Mod(a, b)
}

func NaN() float64 {
	return math.NaN()
}

func Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func Tan(v float64) float64 {
	return math.Tan(v)
}
