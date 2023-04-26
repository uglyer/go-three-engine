//go:build wasm

package wgpu_bridge

import (
	"fmt"
	"syscall/js"
)

type Canvas struct {
	descriptor *CanvasDescriptor
	canvas     js.Value
}

func NewCanvas(descriptor *CanvasDescriptor) ICanvas {
	// Create or get WebGlCanvas
	c := &Canvas{
		descriptor: descriptor,
	}
	doc := js.Global().Get("document")
	if descriptor.CanvasId == "" {
		c.canvas = doc.Call("createElement", "Canvas")
	} else {
		c.canvas = doc.Call("getElementById", descriptor.CanvasId)
	}
	c.canvas.Set("width", descriptor.Width)
	c.canvas.Set("height", descriptor.Height)
	if descriptor.ParentId != "" {
		doc.Call("getElementById", descriptor.ParentId).Call("appendChild", c.canvas)
	}
	return c
}

func (c *Canvas) Drop() {
}

func (c *Canvas) RequestAdapter(descriptor *AdapterDescriptor) (IAdapter, error) {
	return nil, fmt.Errorf("todo impl RequestAdapter")
}
