package camera

import (
	"github.com/uglyer/go-three-engine/engine/math32"
	"github.com/uglyer/go-three-engine/engine/object3d"
)

// Type 相机投影方式类型
type Type int

// The possible camera projections.
const (
	Perspective = Type(iota)
	Orthographic
)

// Camera 相机对象
type Camera struct {
	object3d.Node
	// cameraType 相机类型（正交相机/透视相机）
	cameraType Type
	// aspect 宽高比
	aspect float32
	// near 此范围内的内容不渲染
	near float32
	// far 超出此范围的内容不渲染
	far float32
	// fov 水平朝向视角
	fov float32
	// zoom 缩放值
	zoom float32
	// projChanged 投影变更事件
	projChanged bool
	// projMatrix 相机投影矩阵
	projectionMatrix *math32.Matrix4
	// projMatrix 相机投影逆矩阵
	projectionMatrixInverse *math32.Matrix4
}

// updateProjectionMatrix 更新投影矩阵
func (c *Camera) updateProjectionMatrix() {
	c.projectionMatrix.MakePerspective(c.fov, c.aspect, c.near, c.far)
	c.projectionMatrixInverse.GetInverse(c.projectionMatrix)
}
