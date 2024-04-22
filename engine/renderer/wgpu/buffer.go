package wgpu

type BufferMapCallback func(BufferMapAsyncStatus)

type IBuffer interface {
	IGPUBasic
	MapAsync(mode MapMode, offset uint64, size uint64, callback BufferMapCallback)
	GetMappedRange(offset, size uint) []byte
	Unmap()
}
