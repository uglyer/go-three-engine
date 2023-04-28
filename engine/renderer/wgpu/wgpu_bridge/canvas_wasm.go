//go:build wasm

package wgpu_bridge

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"log"
	"syscall/js"
)

type Canvas struct {
	descriptor    *CanvasDescriptor
	canvas        js.Value
	canvasContext js.Value
}

func NewCanvas(descriptor *CanvasDescriptor) (ICanvas, error) {
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
	return c, nil
}

func (c *Canvas) Drop() {
	log.Println("todo Drop c.canvasContext")
	parent := c.canvas.Get("parent")
	if wasm.IsUndefined(parent) {
		return
	}
	parent.Call("removeChild", parent)
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
	adapter, err := wasm.Await(gpu.Call("requestAdapter", obj))
	if err != nil {
		return nil, fmt.Errorf("requestAdapter error:%v", err)
	}
	wasm.ConsoleLog("xxx")
	wasm.ConsoleLog(*adapter)
	c.canvasContext = c.canvas.Call("getContext", "webgpu")
	wasm.ConsoleLog(c.canvasContext)
	if wasm.IsUndefined(c.canvasContext) {
		return nil, fmt.Errorf("获取上下文失败(gpupresent)")
	}
	return newAdapter(*adapter)
}
