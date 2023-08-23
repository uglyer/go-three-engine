package core

import "github.com/uglyer/go-three-engine/engine/math64"

// BufferAttribute 几何体属性数据
type BufferAttribute struct {
	// Array
	Array []float64
	// ItemSize 单组长度
	ItemSize int
	// Count 成员个数 (Array/ItemSize)
	Count int
}

// CopyToVector3 拷贝顶点到三维向量中
func (b *BufferAttribute) CopyToVector3(v *math64.Vector3, index int) *math64.Vector3 {
	return v.Set(b.Array[index*b.ItemSize], b.Array[index*b.ItemSize+1], b.Array[index*b.ItemSize+2])
}
