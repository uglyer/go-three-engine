package geometry

import (
	"github.com/uglyer/go-three-engine/engine/core"
	"github.com/uglyer/go-three-engine/engine/math64"
	"sync"
)

type Geometry struct {
	mtx            sync.Mutex
	attributeMap   map[string]*core.BufferAttribute
	boundingBox    *math64.Box3
	boundingSphere *math64.Sphere
}

// BoundingBox 获取几何体的包围盒（始终为最新值, 内部按需触发自动计算）
func (g *Geometry) BoundingBox() *math64.Box3 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if g.boundingBox == nil {
		g.boundingBox = math64.NewBox3Infinity()
	} else {
		g.boundingBox.MakeEmpty()
	}
	positionAttr := g.attributeMap["position"]
	for i := 0; i < positionAttr.Count; i += positionAttr.ItemSize {
		g.boundingBox.ExpandByPointXYZ(positionAttr.Array[i], positionAttr.Array[i+1], positionAttr.Array[i+2])
	}
	return g.boundingBox
}

// BoundingSphere 获取几何体的包围球体（始终为最新值, 内部按需触发自动计算）
func (g *Geometry) BoundingSphere() *math64.Sphere {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if g.boundingSphere == nil {
		g.boundingSphere = math64.NewSphere(math64.NewVec3(), 0)
	} else {
		g.boundingSphere.Center.Set(0, 0, 0)
		g.boundingSphere.Radius = 0
	}
	return g.boundingSphere.SetFromBox3(g.boundingBox)
}

// GetAttribute 通过名称获取指定的缓冲区属性
func (g *Geometry) GetAttribute(name string) *core.BufferAttribute {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.attributeMap[name]
}

// SetAttribute 通过名称设置指定的缓冲区属性
func (g *Geometry) SetAttribute(name string, attr *core.BufferAttribute) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.attributeMap[name] = attr
}

// DeleteAttribute 通过名称删除指定的缓冲区属性
func (g *Geometry) DeleteAttribute(name string) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	delete(g.attributeMap, name)
}

// HasAttribute 通过名称校验是否有指定的缓冲区属性
func (g *Geometry) HasAttribute(name string) (ok bool) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	_, ok = g.attributeMap[name]
	return
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
