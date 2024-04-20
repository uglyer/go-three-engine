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
	ref                   js.Value
	preferredCanvasFormat wgpu.TextureFormat
}

func newGPU() (wgpu.IGPU, error) {
	gpu := js.Global().Get("navigator").Get("gpu")
	if wasm.IsUndefined(gpu) {
		return nil, fmt.Errorf("navigator gpu is null")
	}
	return &GPU{ref: gpu, preferredCanvasFormat: wgpu.StringToTextureFormat(gpu.Call("getPreferredCanvasFormat").String())}, nil
}

func (g *GPU) GetPreferredCanvasFormat() wgpu.TextureFormat {
	return g.preferredCanvasFormat
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

func (p *BindGroup) Release()       { panic("todo impl Release") }
func (p *BindGroupLayout) Release() { panic("todo impl Release") }
func (p *CommandBuffer) Release()   { panic("todo impl Release") }
func (p *PipelineLayout) Release()  { panic("todo impl Release") }
func (p *QuerySet) Release()        { panic("todo impl Release") }
func (p *RenderBundle) Release()    { panic("todo impl Release") }
func (p *Sampler) Release()         { panic("todo impl Release") }
func (p *ShaderModule) Release()    { panic("todo impl Release") }
func (p *TextureView) Release()     { panic("todo impl Release") }
