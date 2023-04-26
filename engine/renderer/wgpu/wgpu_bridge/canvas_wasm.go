//go:build wasm

package wgpu_bridge

import "fmt"

type Canvas struct {
	descriptor *CanvasDescriptor
}

func NewCanvas(descriptor *CanvasDescriptor) ICanvas {
	canvas := &Canvas{descriptor}
	return canvas
}

func (c *Canvas) Drop() {
}

func (c *Canvas) RequestAdapter(descriptor *AdapterDescriptor) (IAdapter, error) {
	return nil, fmt.Errorf("todo impl RequestAdapter")
}
