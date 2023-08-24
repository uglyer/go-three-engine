package object3d

import (
	"github.com/uglyer/go-three-engine/engine/core"
	"github.com/uglyer/go-three-engine/engine/math64"
)

type Node struct {
	core.IEventDispatcher
}

func NewNode() *Node {
	return &Node{
		IEventDispatcher: core.NewEventDispatcher(),
	}
}

func (n *Node) IsVisible() bool {
	panic("IsVisible")
}

func (n *Node) SetIsVisible(state bool) {
	panic("SetIsVisible")
}

func (n *Node) Name() string {
	panic("Name")
}

func (n *Node) SetName(string) {
	panic("SetName")
}

func (n *Node) Parent() core.INode {
	panic("Parent")
}

func (n *Node) Children() []core.INode {
	panic("Children")
}

func (n *Node) IsAncestorOf(core.INode) bool {
	panic("IsAncestorOf")
}

func (n *Node) UpdateMatrixWorld() {
	panic("UpdateMatrixWorld")
}

func (n *Node) BoundingBox() *math64.Box3 {
	panic("BoundingBox")
}

func (n *Node) Dispose() {
	panic("Dispose")
}

func (n *Node) Position() *math64.Vector3 {
	panic("Position")
}

func (n *Node) Rotation() *math64.Vector3 {
	panic("Rotation")
}

func (n *Node) Quaternion() *math64.Quaternion {
	panic("Quaternion")
}

func (n *Node) Scale() *math64.Vector3 {
	panic("Scale")
}
