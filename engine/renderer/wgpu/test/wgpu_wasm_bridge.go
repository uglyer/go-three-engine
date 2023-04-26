package main

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_bridge"
)

func main() {
	bridge := wgpu_bridge.NewBridge()
	canvas := bridge.CreateCanvas(&wgpu_bridge.CanvasDescriptor{
		Width:    640,
		Height:   480,
		Title:    "Test",
		ParentId: "root",
	})
	adapter, err := canvas.RequestAdapter(&wgpu_bridge.AdapterDescriptor{
		PowerPreference: wgpu_bridge.PowerPreference_Default,
	})
	if err != nil {
		fmt.Errorf("获取适配器失败:%v", err)
	}
	device, err := adapter.RequestDevice(nil)
	if err != nil {
		fmt.Errorf("获取gpu设备失败:%v", err)
	}
	println(device)
}
