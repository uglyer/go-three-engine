package object3d

import (
	"github.com/uglyer/go-three-engine/engine/core"
	"github.com/uglyer/go-three-engine/engine/math64"
)

type Node struct {
	core.IEventDispatcher
	// NodeName 当前节点的名称（业务名称）, 可不唯一
	NodeName string
	// Visible 当前节点是否可见
	Visible bool
}

func NewNode() *Node {
	return &Node{
		IEventDispatcher: core.NewEventDispatcher(),
		Visible:          true,
	}
}

func (n *Node) IsVisible() bool {
	return n.Visible
}

func (n *Node) SetIsVisible(state bool) {
	n.Visible = state
}

func (n *Node) Name() string {
	return n.NodeName
}

func (n *Node) SetName(name string) {
	n.NodeName = name
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
