package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type Buffer struct {
	ref js.Value
}

func (b *Buffer) Drop() {
	panic("todo impl Drop")
}

func (b *Buffer) Destroy() {
	panic("todo impl Destroy")
}

func (b *Buffer) MapAsync(mode wgpu.MapMode, offset uint64, size uint64, callback wgpu.BufferMapCallback){
	panic("todo impl MapAsync")
}

func (b *Buffer) GetMappedRange(offset, size uint) []byte {
	panic("todo impl GetMappedRange")
}

func (b *Buffer) Unmap() {
	panic("todo impl Unmap")
}
