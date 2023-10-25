package core

import (
	"github.com/uglyer/go-three-engine/engine/math64"
)

// IGeometry 几何体的接口定义
type IGeometry interface {
	IEventDispatcher
	// BoundingBox 获取几何体的包围盒（始终为最新值, 内部按需触发自动计算）
	BoundingBox() *math64.Box3
	// BoundingSphere 获取几何体的包围球体（始终为最新值, 内部按需触发自动计算）
	BoundingSphere() *math64.Sphere
	// GetAttribute 通过名称获取指定的缓冲区属性
	GetAttribute(name string) *BufferAttribute
	// SetAttribute 通过名称设置指定的缓冲区属性
	SetAttribute(name string, attr *BufferAttribute)
	// DeleteAttribute 通过名称删除指定的缓冲区属性
	DeleteAttribute(name string)
	// HasAttribute 通过名称校验是否有指定的缓冲区属性
	HasAttribute(name string) bool
	// ApplyMatrix4 应用指定的矩阵
	ApplyMatrix4(mat4 *math64.Matrix4)
	// Scale 缩放几何体
	Scale(x, y, z float32)
	// Translate 平移几何体
	Translate(x, y, z float32)
	// MakeCenter 使得几何体居中
	MakeCenter()
	// Clone 克隆一个几何体对象
	Clone() IGeometry
	// Copy 拷贝目标的几何体对象到当前
	Copy(target *IGeometry)
}
