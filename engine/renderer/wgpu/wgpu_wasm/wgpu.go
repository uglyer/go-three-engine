//go:build wasm

package wgpu_wasm

import "C"
import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

type GPU struct {
	ref js.Value
}

func newGPU() (wgpu.IGPU, error) {
	gpu := js.Global().Get("navigator").Get("gpu")
	if wasm.IsUndefined(gpu) {
		return nil, fmt.Errorf("navigator gpu is null")
	}
	return &GPU{ref: gpu}, nil
}

func (g *GPU) GetPreferredCanvasFormat() wgpu.TextureFormat {
	println("GetPreferredCanvasFormat")
	v := g.ref.Call("getPreferredCanvasFormat")
	wasm.ConsoleLog("v", v)
	return wgpu.StringToTextureFormat(v.String())
}

func (g *GPU) RequestAdapter(descriptor *wgpu.AdapterDescriptor) (wgpu.IAdapter, error) {
	obj := wasm.NewObject()
	if descriptor != nil {
		obj.Set("powerPreference", descriptor.PowerPreference.String())
	}
	adapter, err := wasm.Await(g.ref.Call("requestAdapter", obj))
	if err != nil {
		return nil, fmt.Errorf("requestAdapter error:%v", err)
	}
	return newAdapter(*adapter)
}

type (
	BindGroup       struct{ ref js.Value }
	BindGroupLayout struct{ ref js.Value }
	CommandBuffer   struct{ ref js.Value }
	PipelineLayout  struct{ ref js.Value }
	QuerySet        struct{ ref js.Value }
	RenderBundle    struct{ ref js.Value }
	Sampler         struct{ ref js.Value }
	ShaderModule    struct{ ref js.Value }
	TextureView     struct{ ref js.Value }
)

func (p *BindGroup) Drop()       { panic("todo impl Drop") }
func (p *BindGroupLayout) Drop() { panic("todo impl Drop") }
func (p *CommandBuffer) Drop()   { panic("todo impl Drop") }
func (p *PipelineLayout) Drop()  { panic("todo impl Drop") }
func (p *QuerySet) Drop()        { panic("todo impl Drop") }
func (p *RenderBundle) Drop()    { panic("todo impl Drop") }
func (p *Sampler) Drop()         { panic("todo impl Drop") }
func (p *ShaderModule) Drop()    { panic("todo impl Drop") }
func (p *TextureView) Drop()     { panic("todo impl Drop") }
