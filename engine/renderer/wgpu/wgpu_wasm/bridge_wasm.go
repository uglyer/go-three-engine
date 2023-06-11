//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
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
func (b *Bridge) GetGPU() wgpu.IGPU {
	return b.gpu
}
