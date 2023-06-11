//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

type GPU struct {
	jsRef js.Value
}

func newGPU() (wgpu.IGPU, error) {
	gpu := js.Global().Get("navigator").Get("gpu")
	if wasm.IsUndefined(gpu) {
		return nil, fmt.Errorf("navigator gpu is null")
	}
	return &GPU{jsRef: gpu}, nil
}

func (g *GPU) GetPreferredCanvasFormat() wgpu.TextureFormat {
	return wgpu.TextureFormat_BGRA8Unorm
}

func (g *GPU) RequestAdapter(descriptor *wgpu.AdapterDescriptor) (wgpu.IAdapter, error) {
	obj := wasm.NewObject()
	if descriptor != nil {
		obj.Set("powerPreference", descriptor.PowerPreference.String())
	}
	adapter, err := wasm.Await(g.jsRef.Call("requestAdapter", obj))
	if err != nil {
		return nil, fmt.Errorf("requestAdapter error:%v", err)
	}
	return newAdapter(*adapter)
}
