package wgpu

import "C"

type BufferMapCallback func(BufferMapAsyncStatus)

type IBuffer interface {
	IGPUBasic
	MapAsync(mode MapMode, offset uint64, size uint64, callback BufferMapCallback) error
	GetMappedRange(offset, size uint) []byte
	Unmap() error
	GetSize() uint64
	GetUsage() BufferUsage
	Release()
}
