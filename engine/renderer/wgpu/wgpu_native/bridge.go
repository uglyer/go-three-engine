//go:build !wasm

package wgpu_native

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
)

type Bridge struct {
	instance *Instance
	format   wgpu.TextureFormat
}

func NewBridge() (wgpu.IBridge, error) {
	instance := CreateInstance(nil)
	return &Bridge{
		instance: instance,
	}, nil
}

func (b *Bridge) CreateCanvas(descriptor *wgpu.CanvasDescriptor) (wgpu.ICanvas, error) {
	canvas, err := NewCanvas(b.instance, descriptor)
	if err != nil {
		return nil, err
	}
	b.format = canvas.config.Format
	return canvas, nil
}

func (b *Bridge) RequestAnimationFrame(fn func()) {

}

func (b *Bridge) GetGPU() wgpu.IGPU {
	return b
}

func (b *Bridge) GetPreferredCanvasFormat() wgpu.TextureFormat {
	return b.format
}

func (b *Bridge) RequestAdapter(descriptor *wgpu.AdapterDescriptor) (wgpu.IAdapter, error) {
	// TODO 支持其他参数转发
	return b.instance.RequestAdapter(&RequestAdapterOptions{
		PowerPreference: descriptor.PowerPreference,
	})
}
