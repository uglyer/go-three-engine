//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
)

type Bridge struct {
}

func NewBridge() wgpu.IBridge {
	return &Bridge{}
}

func (b *Bridge) CreateCanvas(descriptor *wgpu.CanvasDescriptor) (wgpu.ICanvas, error) {
	return NewCanvas(descriptor)
}
func (b *Bridge) GetGPU() (wgpu.IGPU, error) {
	return
}
