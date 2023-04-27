package main

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_bridge"
	"log"
)

func main() {
	bridge := wgpu_bridge.NewBridge()
	canvas, err := bridge.CreateCanvas(&wgpu_bridge.CanvasDescriptor{
		Width:    640,
		Height:   480,
		Title:    "Test",
		ParentId: "root",
	})
	if err != nil {
		log.Fatalf("创建画布失败:%v", err)
	}
	adapter, err := canvas.RequestAdapter(&wgpu_bridge.AdapterDescriptor{
		PowerPreference: wgpu_bridge.PowerPreference_HighPerformance,
	})
	if err != nil {
		log.Fatalf("获取适配器失败:%v", err)
	}
	device, err := adapter.RequestDevice(nil)
	if err != nil {
		log.Fatalf("获取gpu设备失败:%v", err)
	}
	println(device)
}
