//go:build wasm

package wgpu_bridge

import (
	"fmt"
	"log"
	"syscall/js"
)

type Device struct {
	deviceRef js.Value
}

func newDevice(deviceRef js.Value) (IDevice, error) {
	return &Device{deviceRef: deviceRef}, nil
}

func (d *Device) Drop() {

}
func (d *Device) GetQueue() IQueue {
	log.Panicln("todo impl GetQueue")
	return nil
}
func (d *Device) CreateCommandEncoder() (IGpuCommandEncoder, error) {
	return nil, fmt.Errorf("todo impl CreateCommandEncoder")
}
func (d *Device) CreateBuffer() (IGpuBuffer, error) {
	return nil, fmt.Errorf("todo impl CreateBuffer")
}
func (d *Device) CreateTexture() (IGpuTexture, error) {
	return nil, fmt.Errorf("todo impl CreateTexture")
}
func (d *Device) CreateRenderPipeline() (IGpuPipeLine, error) {
	return nil, fmt.Errorf("todo impl CreateRenderPipeline")
}
