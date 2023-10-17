package core

import (
	"github.com/uglyer/go-three-engine/engine/math32"
)

// BufferAttribute 几何体属性数据
type BufferAttribute struct {
	// Array
	Array []float32
	// ItemSize 单组长度
	ItemSize int
	// Count 成员个数 (Array/ItemSize)
	Count int
}

// CopyToVector3 拷贝顶点到三维向量中
func (b *BufferAttribute) CopyToVector3(v *math32.Vector3, index int) *math32.Vector3 {
	return v.Set(b.Array[index*b.ItemSize], b.Array[index*b.ItemSize+1], b.Array[index*b.ItemSize+2])
}

// OperateOnVector3 操作三维向量
func (b *BufferAttribute) OperateOnVector3(cb func(vertex *math32.Vector3) bool) {
	v := math32.NewVec3()
	for index := 0; index < b.Count; index += b.ItemSize {
		v.Set(b.Array[index*b.ItemSize], b.Array[index*b.ItemSize+1], b.Array[index*b.ItemSize+2])
		if cb(v) {
			break
		}
	}
}
