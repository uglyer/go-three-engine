//go:build wasm

package wgpu_bridge

import "fmt"

type Adapter struct {
}

func (a *Adapter) Drop() {

}

func (a *Adapter) RequestDevice(descriptor *DeviceDescriptor) (IDevice, error) {
	return nil, fmt.Errorf("todo implRequestDevice ")
}
