//go:build wasm

package wgpu_bridge

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/wasm"
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
	gpu := js.Global().Get("navigator").Get("gpu")
	if wasm.IsUndefined(gpu) {
		return nil, fmt.Errorf("navigator gpu is null")
	}
	obj := wasm.NewObject()
	if descriptor != nil {
		obj.Set("powerPreference", WasmPowerPreference[descriptor.PowerPreference])
	}
	wasm.ConsoleLog("xxx", obj)
	promise := gpu.Call("requestAdapter", obj)
	wasm.ConsoleLog("xxx")
	wasm.ConsoleLog(promise)
	return nil, fmt.Errorf("todo impl RequestAdapter")
}
