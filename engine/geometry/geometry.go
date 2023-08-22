package geometry

import (
	"github.com/uglyer/go-three-engine/engine/core"
	"github.com/uglyer/go-three-engine/engine/math64"
)

type Geometry struct {
}

// BoundingBox 获取几何体的包围盒（始终为最新值, 内部按需触发自动计算）
func (g *Geometry) BoundingBox() *math64.Box3 {
	panic("BoundingBox")
}

// BoundingSphere 获取几何体的包围球体（始终为最新值, 内部按需触发自动计算）
func (g *Geometry) BoundingSphere() *math64.Sphere {
	panic("BoundingSphere")
}

// GetAttribute 通过名称获取指定的缓冲区属性
func (g *Geometry) GetAttribute(name string) *core.BufferAttribute {
	panic("GetAttribute")
}

// SetAttribute 通过名称设置指定的缓冲区属性
func (g *Geometry) SetAttribute(name string, attr *core.BufferAttribute) {
	panic("GetAttribute")
}

// DeleteAttribute 通过名称删除指定的缓冲区属性
func (g *Geometry) DeleteAttribute(name string) {
	panic("DeleteAttribute")
}

// HasAttribute 通过名称校验是否有指定的缓冲区属性
func (g *Geometry) HasAttribute(name string) bool {
	panic("HasAttribute")
}

// ApplyMatrix4 应用指定的矩阵
func (g *Geometry) ApplyMatrix4(mat4 *math64.Matrix4) {
	panic("ApplyMatrix4")
}

// MakeCenter 使得几何体居中
func (g *Geometry) MakeCenter() {
	panic("MakeCenter")
}

// Clone 克隆一个几何体对象
func (g *Geometry) Clone() *core.IGeometry {
	panic("Clone")
}

// Copy 拷贝目标的几何体对象到当前
func (g *Geometry) Copy(target *core.IGeometry) {
	panic("Copy")
}
