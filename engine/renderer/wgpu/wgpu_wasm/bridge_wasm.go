//go:build wasm

package wgpu_wasm

import "github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_bridge"

type Bridge struct {
}

func NewBridge() wgpu_bridge.IBridge {
	return &Bridge{}
}

func (b *Bridge) CreateCanvas(descriptor *wgpu_bridge.CanvasDescriptor) (wgpu_bridge.ICanvas, error) {
	return NewCanvas(descriptor)
}
