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
	ref               js.Value
	enumerateFeatures []wgpu.FeatureName
}

func newAdapter(adapterRef js.Value) (wgpu.IAdapter, error) {
	wasm.ConsoleLog("adapterRef", adapterRef)
	return &Adapter{
		ref: adapterRef,
	}, nil
}

func (a *Adapter) Release() {
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
	if a.enumerateFeatures == nil {
		featuresRef := a.ref.Get("features")
		list := make([]wgpu.FeatureName, 0)
		featuresRef.Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			list = append(list, wgpu.FeatureName(args[0].String()))
			return nil
		}))
		a.enumerateFeatures = list
	}
	return a.enumerateFeatures
}
func (a *Adapter) GetLimits() wgpu.SupportedLimits {
	jsLimits := a.ref.Get("limits")
	limit := wgpu.Limits{}
	valueOfLimit := reflect.ValueOf(&limit).Elem()
	typeOfLimit := reflect.TypeOf(limit)
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
	return wgpu.SupportedLimits{
		Limits: &limit,
	}
}

func (a *Adapter) GetProperties() wgpu.AdapterProperties {
	// TODO impl adapter EnumerateFeatures
	return wgpu.AdapterProperties{}
}
func (a *Adapter) HasFeature(feature wgpu.FeatureName) bool {
	features := a.EnumerateFeatures()
	for _, it := range features {
		if it == feature {
			return true
		}
	}
	return false
}
