//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_bridge"
	"log"
	"syscall/js"
)

type Device struct {
	deviceRef js.Value
}

func newDevice(deviceRef js.Value) (wgpu.IDevice, error) {
	return &Device{deviceRef: deviceRef}, nil
}

func (d *Device) Drop() {

}
func (d *Device) GetQueue() wgpu_bridge.IGpuQueue {
	log.Panicln("todo impl GetQueue")
	return nil
}
func (d *Device) CreateCommandEncoder() (wgpu.ICommandEncoder, error) {
	return nil, fmt.Errorf("todo impl CreateCommandEncoder")
}
func (d *Device) CreateBuffer() (wgpu.IBuffer, error) {
	return nil, fmt.Errorf("todo impl CreateBuffer")
}
func (d *Device) CreateTexture() (wgpu.ITexture, error) {
	return nil, fmt.Errorf("todo impl CreateTexture")
}
func (d *Device) CreateRenderPipeline() (wgpu.IGPURenderPipeLine, error) {
	return nil, fmt.Errorf("todo impl CreateRenderPipeline")
}
