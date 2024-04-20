//go:build !wasm

package wgpu_native

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	wgpuext_glfw "github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_native/wgpuext/glfw"
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
	window, err := glfw.CreateWindow(descriptor.Width, descriptor.Width, descriptor.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	surface := b.instance.CreateSurface(wgpuext_glfw.GetSurfaceDescriptor(window))
	panic("TODO impl createCanvas")
}

func (b *Bridge) RequestAnimationFrame(fn func()) {

}

func (b *Bridge) GetGPU() wgpu.IGPU {
	return b.instance
}
