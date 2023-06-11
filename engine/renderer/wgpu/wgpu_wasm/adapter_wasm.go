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
	adapterRef js.Value
}

func newAdapter(adapterRef js.Value) (wgpu.IAdapter, error) {
	return &Adapter{
		adapterRef: adapterRef,
	}, nil
}

func (a *Adapter) Drop() {

}

func (a *Adapter) RequestDevice(descriptor *wgpu.DeviceDescriptor) (wgpu.IDevice, error) {
	if descriptor != nil {
		log.Println("todo RequestDevice with descriptor")
	}
	device, err := wasm.Await(a.adapterRef.Call("requestDevice"))
	if err != nil {
		return nil, fmt.Errorf("call requestDevice error:%v", err)
	}
	wasm.ConsoleLog("device", *device)
	return newDevice(*device)
}
