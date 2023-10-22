package geometry

import (
	"github.com/uglyer/go-three-engine/engine/core"
	"github.com/uglyer/go-three-engine/engine/math32"
	"sync"
)

type Geometry struct {
	mtx                   sync.Mutex
	attributeMap          map[string]*core.BufferAttribute
	boundingBoxNeedUpdate bool
	boundingBox           *math32.Box3
	boundingSphere        *math32.Sphere
}

// BoundingBox 获取几何体的包围盒（始终为最新值, 内部按需触发自动计算）
func (g *Geometry) BoundingBox() *math32.Box3 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if g.boundingBox != nil && !g.boundingBoxNeedUpdate {
		return g.boundingBox
	}
	if g.boundingBox == nil {
		g.boundingBox = math32.NewBox3Infinity()
	} else {
		g.boundingBox.MakeEmpty()
	}
	positionAttr := g.attributeMap["position"]
	for i := 0; i < positionAttr.Count; i += positionAttr.ItemSize {
		g.boundingBox.ExpandByPointXYZ(positionAttr.Array[i], positionAttr.Array[i+1], positionAttr.Array[i+2])
	}
	g.boundingBoxNeedUpdate = false
	return g.boundingBox
}

// BoundingSphere 获取几何体的包围球体（始终为最新值, 内部按需触发自动计算）
func (g *Geometry) BoundingSphere() *math32.Sphere {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if g.boundingSphere == nil {
		g.boundingSphere = math32.NewSphere(math32.NewVec3(), 0)
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
func (g *Geometry) ApplyMatrix4(mat4 *math32.Matrix4) {
	position := g.GetAttribute("position")
	if position != nil {
		position.OperateOnVector3(func(vertex *math32.Vector3) bool {
			vertex.ApplyMatrix4(mat4)
			return false
		})
	}
}

// Scale 缩放几何体
func (g *Geometry) Scale(x, y, z float32) {
	m := math32.NewMatrix4().MakeScale(x, y, z)
	g.ApplyMatrix4(m)
}

// Translate 平移几何体
func (g *Geometry) Translate(x, y, z float32) {
	m := math32.NewMatrix4().MakeScale(x, y, z)
	g.ApplyMatrix4(m)
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
