package wgpu_bridge

type BufferMapCallback func(BufferMapAsyncStatus)

type IGPUBuffer interface {
	IDrop
	Destroy()
	MapAsync(mode MapMode, offset uint64, size uint64, callback BufferMapCallback)
	GetMappedRange(offset, size uint) []byte
	Unmap()
}
