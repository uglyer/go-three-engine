package core

import "github.com/uglyer/go-three-engine/engine/math64"

type IGeometry interface {
	BoundingBox() *math64.Box3

	BoundingSphere() *math64.Sphere
}
