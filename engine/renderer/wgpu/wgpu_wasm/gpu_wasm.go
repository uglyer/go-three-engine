//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

type GPU struct {
}

func (g *GPU) GetPreferredCanvasFormat() wgpu.TextureFormat {

}

func (g *GPU) RequestAdapter(descriptor *wgpu.AdapterDescriptor) (wgpu.IAdapter, error) {
	gpu := js.Global().Get("navigator").Get("gpu")
	if wasm.IsUndefined(gpu) {
		return nil, fmt.Errorf("navigator gpu is null")
	}
	obj := wasm.NewObject()
	if descriptor != nil {
		obj.Set("powerPreference", descriptor.PowerPreference.String())
	}
	adapter, err := wasm.Await(gpu.Call("requestAdapter", obj))
	if err != nil {
		return nil, fmt.Errorf("requestAdapter error:%v", err)
	}
	return newAdapter(*adapter)
}
