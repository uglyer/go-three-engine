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
	// ParentNode 父节点
	ParentNode core.INode
	// ChildrenNodes 子节点
	ChildrenNodes []core.INode
	position      *math64.Vector3
	quaternion    *math64.Quaternion
	// rotation 欧拉角为缓存对象, 以四元数为准
	rotation *math64.Vector3
	scale    *math64.Vector3
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
	return n.ParentNode
}

func (n *Node) Children() []core.INode {
	return n.ChildrenNodes
}

func (n *Node) IsAncestorOf(core.INode) bool {
	panic("IsAncestorOf")
}

func (n *Node) AddChild(node core.INode) {
	n.ChildrenNodes = append(n.ChildrenNodes, node)
	// TODO 更新父节点
}

func (n *Node) Remove(node core.INode) {
	panic("Remove")
}

func (n *Node) RemoveInParent() {
	panic("AddChild")
}

func (n *Node) UpdateMatrixWorld() {
	panic("UpdateMatrixWorld")
}

func (n *Node) BoundingBox() *math64.Box3 {
	panic("BoundingBox")
}

func (n *Node) Dispose() {
	n.DispatchEvent("dispose", nil)
}

func (n *Node) Position() *math64.Vector3 {
	return n.position
}

func (n *Node) RotationReadonly() *math64.Vector3 {
	n.rotation.SetFromQuaternion(n.quaternion)
	return n.rotation
}

func (n *Node) Quaternion() *math64.Quaternion {
	return n.quaternion
}

func (n *Node) Scale() *math64.Vector3 {
	return n.scale
}
