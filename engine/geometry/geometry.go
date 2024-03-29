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
	if g.boundingSphere != nil && !g.boundingBoxNeedUpdate {
		return g.boundingSphere
	}
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

// ComputeNormal 重新计算法向量
func (g *Geometry) ComputeNormal() {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	position := g.GetAttribute("position")
	normal := g.GetAttribute("normal")
	if normal != nil {
		// TODO 释放
	}
	normal = &core.BufferAttribute{
		Array:    make([]float32, len(position.Array)),
		ItemSize: 3,
		Count:    position.Count,
	}
	a, b, c := math32.NewVec3(), math32.NewVec3(), math32.NewVec3()
	for index := 0; index < position.Count; index += 9 {
		a.Set(position.Array[index*position.ItemSize], position.Array[index*position.ItemSize+1], position.Array[index*position.ItemSize+2])
		b.Set(position.Array[index*position.ItemSize+3], position.Array[index*position.ItemSize+4], position.Array[index*position.ItemSize+5])
		c.Set(position.Array[index*position.ItemSize+6], position.Array[index*position.ItemSize+7], position.Array[index*position.ItemSize+8])
		v := b.Clone().Sub(a).Cross(c.Clone().Sub(a)).Normalize()
		normal.Array[index*position.ItemSize], normal.Array[index*position.ItemSize+1], normal.Array[index*position.ItemSize+2] = v.X, v.Y, v.Z
		normal.Array[index*position.ItemSize+3], normal.Array[index*position.ItemSize+4], normal.Array[index*position.ItemSize+5] = v.X, v.Y, v.Z
		normal.Array[index*position.ItemSize+6], normal.Array[index*position.ItemSize+6], normal.Array[index*position.ItemSize+6] = v.X, v.Y, v.Z
	}
	g.SetAttribute("normal", normal)
	return
}

// setFromPoints 通过顶点列表更新 position
func (g *Geometry) setFromPoints(points []*math32.Vector3) {
	// TODO 释放旧数据
	position := &core.BufferAttribute{
		Array:    make([]float32, len(points)*3),
		ItemSize: 3,
		Count:    len(points),
	}
	for index, v := range points {
		position.Array[index*position.ItemSize], position.Array[index*position.ItemSize+1], position.Array[index*position.ItemSize+2] = v.X, v.Y, v.Z
	}
	g.SetAttribute("position", position)
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
	g.boundingBoxNeedUpdate = true
}

// Translate 平移几何体
func (g *Geometry) Translate(x, y, z float32) {
	m := math32.NewMatrix4().MakeScale(x, y, z)
	g.ApplyMatrix4(m)
	g.boundingBoxNeedUpdate = true
}

// MakeCenter 使得几何体居中
func (g *Geometry) MakeCenter() {
	box := g.BoundingBox()
	center := box.Center(math32.NewVec3()).Negate()
	g.Translate(center.X, center.Y, center.Z)
}

// Clone 克隆一个几何体对象
func (g *Geometry) Clone() core.IGeometry {
	panic("Clone")
}

// Copy 拷贝目标的几何体对象到当前
func (g *Geometry) Copy(target *Geometry) {
	g.boundingSphere.Copy(target.boundingSphere)
	g.boundingBox.Copy(target.boundingBox)
	for k, v := range target.attributeMap {
		g.attributeMap[k] = v
	}
}
