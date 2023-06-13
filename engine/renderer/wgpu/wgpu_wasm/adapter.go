//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"log"
	"syscall/js"
)

type Adapter struct {
	ref js.Value
}

func newAdapter(adapterRef js.Value) (wgpu.IAdapter, error) {
	return &Adapter{
		ref: adapterRef,
	}, nil
}

func (a *Adapter) Drop() {
	// TODO impl adapter drop
}

func (a *Adapter) RequestDevice(descriptor *wgpu.DeviceDescriptor) (wgpu.IDevice, error) {
	if descriptor != nil {
		log.Println("todo RequestDevice with descriptor")
	}
	device, err := wasm.Await(a.ref.Call("requestDevice"))
	if err != nil {
		return nil, fmt.Errorf("call requestDevice error:%v", err)
	}
	wasm.ConsoleLog("a.ref", a.ref)
	wasm.ConsoleLog("device", *device)
	return newDevice(*device)
}
func (a *Adapter) EnumerateFeatures() []wgpu.FeatureName {
	// TODO impl adapter EnumerateFeatures
	return nil
}
func (a *Adapter) GetLimits() wgpu.SupportedLimits {
	// TODO impl adapter EnumerateFeatures
	return wgpu.SupportedLimits{}
}
func (a *Adapter) GetProperties() wgpu.AdapterProperties {
	// TODO impl adapter EnumerateFeatures
	return wgpu.AdapterProperties{}
}
func (a *Adapter) HasFeature(feature wgpu.FeatureName) bool {
	// TODO impl adapter EnumerateFeatures
	return false
}
