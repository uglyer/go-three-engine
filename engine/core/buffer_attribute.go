package core

// BufferAttribute 几何体属性数据
type BufferAttribute struct {
	// Array
	Array []float64
	// ItemSize 单组长度
	ItemSize int
	// Count 成员个数 (Array/ItemSize)
	Count int
}
