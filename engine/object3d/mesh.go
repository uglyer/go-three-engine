package object3d

import "github.com/uglyer/go-three-engine/engine/geometry"

type Mesh struct {
	Node
	Geometry *geometry.Geometry
}
