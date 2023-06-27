package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

type Buffer struct {
	ref js.Value
}

func (b *Buffer) Drop() {
	panic("todo impl Drop")
}

func (b *Buffer) Destroy() {
	b.ref.Call("destroy")
}

func (b *Buffer) MapAsync(mode wgpu.MapMode, offset uint64, size uint64, callback wgpu.BufferMapCallback) {
	promise := b.ref.Call("mapAsync", int(mode), offset, size)
	go func() {
		_, err := wasm.Await(promise)
		if err != nil {
			wasm.ConsoleError("MapAsync", err.Error())
			callback(wgpu.BufferMapAsyncStatus_Error)
			return
		}
		callback(wgpu.BufferMapAsyncStatus_Success)
	}()
}

func (b *Buffer) GetMappedRange(offset, size uint) []byte {
	arrayBuffer := b.ref.Call("getMappedRange", offset, size)
	// 声明一个和文件大小一样的切片
	buffer := make([]byte, arrayBuffer.Length())
	js.CopyBytesToGo(buffer, arrayBuffer)
	return buffer
}

func (b *Buffer) Unmap() {
	b.ref.Call("unmap")
}
