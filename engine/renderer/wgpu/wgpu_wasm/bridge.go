//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type Bridge struct {
	gpu wgpu.IGPU
}

func NewBridge() (wgpu.IBridge, error) {
	gpu, err := newGPU()
	if err != nil {
		return nil, err
	}
	return &Bridge{
		gpu: gpu,
	}, nil
}

func (b *Bridge) CreateCanvas(descriptor *wgpu.CanvasDescriptor) (wgpu.ICanvas, error) {
	return newCanvas(descriptor)
}

func (b *Bridge) RequestAnimationFrame(fn func()) {
	js.Global().Call("_requestAnimationFrame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fn()
		return nil
	}))
}

func (b *Bridge) GetGPU() wgpu.IGPU {
	return b.gpu
}
