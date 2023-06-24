//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"log"
	"reflect"
	"strings"
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
	// TODO 反射值更新未成功
	jsLimits := a.ref.Get("limits")
	limit := wgpu.Limits{}
	valueOfLimit := reflect.ValueOf(&limit).Elem()
	typeOfLimit := reflect.TypeOf(limit)
	wasm.ConsoleLog("jsLimits", jsLimits)
	size := typeOfLimit.NumField()
	for i := 0; i < size; i++ {
		name := typeOfLimit.Field(i).Name
		field := valueOfLimit.Field(i)
		name = strings.ToLower(name[0:1]) + name[1:]
		jsV, ok := wasm.TryGet(jsLimits, name)
		if !ok || jsV.IsUndefined() {
			continue
		}
		field.SetUint(uint64(jsV.Int()))
	}
	wasm.ConsoleLog("device.GetLimits", uint64(jsLimits.Get("maxBindGroups").Int()))
	wasm.ConsoleLog("device.GetLimits", int(limit.MaxBindGroups))
	return wgpu.SupportedLimits{
		Limits: &limit,
	}
}

func (a *Adapter) GetProperties() wgpu.AdapterProperties {
	// TODO impl adapter EnumerateFeatures
	return wgpu.AdapterProperties{}
}
func (a *Adapter) HasFeature(feature wgpu.FeatureName) bool {
	// TODO impl adapter EnumerateFeatures
	return false
}
