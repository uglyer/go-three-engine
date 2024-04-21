//go:build !wasm

package wgpu_native

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
)

type Bridge struct {
	instance *Instance
}

func NewBridge() (wgpu.IBridge, error) {
	instance := CreateInstance(nil)
	return &Bridge{
		instance: instance,
	}, nil
}

func (b *Bridge) CreateCanvas(descriptor *wgpu.CanvasDescriptor) (wgpu.ICanvas, error) {
	return NewCanvas(b.instance, descriptor)
}

func (b *Bridge) RequestAnimationFrame(fn func()) {

}

func (b *Bridge) GetGPU() wgpu.IGPU {
	return b.instance
}
