package core

import "github.com/uglyer/go-three-engine/engine/math64"

type IEventDispatcher interface {
	// UUID 获取当前对象的 UUID
	UUID() string
	AddEventListener(eventType string, listener func(Event)) int
	RemoveEventListener(eventType string, eventId int)
	DispatchEvent(eventType string, data interface{})
}

type INode interface {
	IEventDispatcher

	IsVisible() bool

	SetIsVisible(state bool)

	Name() string

	SetName(string)

	Parent() INode

	Children() []INode

	AddChild(node INode)

	RemoveInParent()

	IsAncestorOf(INode) bool

	UpdateMatrixWorld()

	BoundingBox() *math64.Box3

	Dispose()

	Position() *math64.Vector3

	Rotation() *math64.Vector3

	Quaternion() *math64.Quaternion

	Scale() *math64.Vector3
}
