package main

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_wasm"
	"log"
)

func main() {

	//wasm.ConsoleLog(js.Global().Get("Object").New(map[string]any{"x": "xx"}))
	bridge, err := wgpu_wasm.NewBridge()
	if err != nil {
		log.Fatalf("获取桥接器失败:%v", err)
	}
	canvas, err := bridge.CreateCanvas(&wgpu.CanvasDescriptor{
		Width:    640,
		Height:   480,
		Title:    "Test",
		ParentId: "root",
	})
	if err != nil {
		log.Fatalf("创建画布失败:%v", err)
	}
	adapter, err := bridge.GetGPU().RequestAdapter(&wgpu.AdapterDescriptor{
		PowerPreference: wgpu.PowerPreference_HighPerformance,
	})
	if err != nil {
		log.Fatalf("获取适配器失败:%v", err)
	}
	device, err := adapter.RequestDevice(nil)
	if err != nil {
		log.Fatalf("获取gpu设备失败:%v", err)
	}
	err = canvas.Configure(&wgpu.ConfigureDescriptor{
		Device: device,
	})
	if err != nil {
		log.Fatalf("配置画布失败:%v", err)
	}
}
