//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"log"
	"syscall/js"
)

type Canvas struct {
	descriptor    *wgpu.CanvasDescriptor
	canvas        js.Value
	canvasContext js.Value
}

func newCanvas(descriptor *wgpu.CanvasDescriptor) (wgpu.ICanvas, error) {
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
	c.canvasContext = c.canvas.Call("getContext", "webgpu")
	wasm.ConsoleLog(c.canvasContext)
	if wasm.IsUndefined(c.canvasContext) {
		return nil, fmt.Errorf("获取上下文失败(gpupresent)")
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

func (c *Canvas) Configure(descriptor *wgpu.ConfigureDescriptor) error {
	device, ok := descriptor.Device.(*Device)
	if !ok {
		return fmt.Errorf("device type error")
	}
	config := wasm.NewObject()
	config.Set("format", "bgra8unorm")
	config.Set("alphaMode", "premultiplied")
	config.Set("device", device.ref)
	wasm.ConsoleLog("Configure", config)
	c.canvasContext.Call("configure", config)
	return nil
}

func (c *Canvas) UnConfigure() {
	c.canvasContext.Call("unConfigure")
}

func (c *Canvas) GetCurrentTexture() wgpu.ITexture {
	ref := c.canvasContext.Call("getCurrentTexture")
	return newTexture(ref)
}
