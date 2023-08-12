// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math64

// Box2 represents a 2D bounding box defined by two points:
// the point with minimum coordinates and the point with maximum coordinates.
type Box2 struct {
	Min *Vector2
	Max *Vector2
}

// NewBox2 creates and returns a pointer to a new Box2 defined
// by its minimum and maximum coordinates.
func NewBox2(min, max *Vector2) *Box2 {

	b := &Box2{Min: min.Clone(), Max: max.Clone()}
	return b
}

// Set sets this bounding box minimum and maximum coordinates.
// Returns pointer to this updated bounding box.
func (b *Box2) Set(min, max *Vector2) *Box2 {

	if min != nil {
		b.Min.Copy(min)
	} else {
		b.Min.Set(Infinity, Infinity)
	}
	if max != nil {
		b.Max.Copy(max)
	} else {
		b.Max.Set(-Infinity, -Infinity)
	}
	return b
}

// SetFromPoints set this bounding box from the specified array of points.
// Returns pointer to this updated bounding box.
func (b *Box2) SetFromPoints(points []*Vector2) *Box2 {

	b.MakeEmpty()
	for i := 0; i < len(points); i++ {
		b.ExpandByPoint(points[i])
	}
	return b
}

// SetFromCenterAndSize set this bounding box from a center point and size.
// Size is a vector from the minimum point to the maximum point.
// Returns pointer to this updated bounding box.
func (b *Box2) SetFromCenterAndSize(center, size *Vector2) *Box2 {

	var v1 Vector2
	halfSize := v1.Copy(size).MultiplyScalar(0.5)
	b.Min.Copy(center).Sub(halfSize)
	b.Max.Copy(center).Add(halfSize)
	return b
}

// Copy copy other to this bounding box.
// Returns pointer to this updated bounding box.
func (b *Box2) Copy(box *Box2) *Box2 {

	b.Min.Copy(box.Min)
	b.Max.Copy(box.Max)
	return b
}

func (b *Box2) Clone() *Box2 {
	return NewBox2(b.Min, b.Max)
}

// MakeEmpty set this bounding box to empty.
// Returns pointer to this updated bounding box.
func (b *Box2) MakeEmpty() *Box2 {

	b.Min.X = Infinity
	b.Min.Y = Infinity
	b.Max.X = -Infinity
	b.Max.Y = -Infinity
	return b
}

// Empty returns if this bounding box is empty.
func (b *Box2) Empty() bool {

	return (b.Max.X < b.Min.X) || (b.Max.Y < b.Min.Y)
}

// Center calculates the center point of this bounding box and
// stores its pointer to optionalTarget, if not nil, and also returns it.
func (b *Box2) Center(optionalTarget *Vector2) *Vector2 {

	var result *Vector2
	if optionalTarget == nil {
		result = NewVector2(0, 0)
	} else {
		result = optionalTarget
	}
	return result.AddVectors(b.Min, b.Max).MultiplyScalar(0.5)
}

// Size calculates the size of this bounding box: the vector  from
// its minimum point to its maximum point.
// Store pointer to the calculated size into optionalTarget, if not nil,
// and also returns it.
func (b *Box2) Size(optionalTarget *Vector2) *Vector2 {

	var result *Vector2
	if optionalTarget == nil {
		result = NewVector2(0, 0)
	} else {
		result = optionalTarget
	}
	return result.SubVectors(b.Max, b.Min)
}

// ExpandByPoint may expand this bounding box to include the specified point.
// Returns pointer to this updated bounding box.
func (b *Box2) ExpandByPoint(point *Vector2) *Box2 {

	b.Min.Min(point)
	b.Max.Max(point)
	return b
}

// ExpandByVector expands this bounding box by the specified vector.
// Returns pointer to this updated bounding box.
func (b *Box2) ExpandByVector(vector *Vector2) *Box2 {

	b.Min.Sub(vector)
	b.Max.Add(vector)
	return b
}

// ExpandByScalar expands this bounding box by the specified scalar.
// Returns pointer to this updated bounding box.
func (b *Box2) ExpandByScalar(scalar float64) *Box2 {

	b.Min.AddScalar(-scalar)
	b.Max.AddScalar(scalar)
	return b
}

// ContainsPoint returns if this bounding box contains the specified point.
func (b *Box2) ContainsPoint(point *Vector2) bool {

	if point.X < b.Min.X || point.X > b.Max.X ||
		point.Y < b.Min.Y || point.Y > b.Max.Y {
		return false
	}
	return true
}

// ContainsBox returns if this bounding box contains other box.
func (b *Box2) ContainsBox(other *Box2) bool {

	if (b.Min.X <= other.Min.X) && (other.Max.X <= b.Max.X) &&
		(b.Min.Y <= other.Min.Y) && (other.Max.Y <= b.Max.Y) {
		return true

	}
	return false
}

// IsIntersectionBox returns if other box intersects this one.
func (b *Box2) IsIntersectionBox(other *Box2) bool {

	// using 6 splitting planes to rule out intersections.
	if other.Max.X < b.Min.X || other.Min.X > b.Max.X ||
		other.Max.Y < b.Min.Y || other.Min.Y > b.Max.Y {
		return false
	}
	return true
}

// ClampPoint calculates a new point which is the specified point clamped inside this box.
// Stores the pointer to this new point into optionaTarget, if not nil, and also returns it.
func (b *Box2) ClampPoint(point *Vector2, optionalTarget *Vector2) *Vector2 {

	var result *Vector2
	if optionalTarget == nil {
		result = NewVector2(0, 0)
	} else {
		result = optionalTarget
	}
	return result.Copy(point).Clamp(b.Min, b.Max)
}

// DistanceToPoint returns the distance from this box to the specified point.
func (b *Box2) DistanceToPoint(point *Vector2) float64 {

	v1 := NewVector2(0, 0)
	clampedPoint := v1.Copy(point).Clamp(b.Min, b.Max)
	return clampedPoint.Sub(point).Length()
}

// Intersect sets this box to the intersection with other box.
// Returns pointer to this updated bounding box.
func (b *Box2) Intersect(other *Box2) *Box2 {

	b.Min.Max(other.Min)
	b.Max.Min(other.Max)
	return b
}

// Union set this box to the union with other box.
// Returns pointer to this updated bounding box.
func (b *Box2) Union(other *Box2) *Box2 {

	b.Min.Min(other.Min)
	b.Max.Max(other.Max)
	return b
}

// Translate translates the position of this box by offset.
// Returns pointer to this updated box.
func (b *Box2) Translate(offset *Vector2) *Box2 {

	b.Min.Add(offset)
	b.Max.Add(offset)
	return b
}

// Equals returns if this box is equal to other
func (b *Box2) Equals(other *Box2) bool {

	return other.Min.Equals(b.Min) && other.Max.Equals(b.Max)
}
