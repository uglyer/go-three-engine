package main

import (
	_ "embed"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_wasm"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"log"
)

//go:embed shader.wgsl
var shader string

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
	shaderModule, err := device.CreateShaderModule(&wgpu.ShaderModuleDescriptor{
		Label: "draw.wgsl",
		WGSLDescriptor: &wgpu.ShaderModuleWGSLDescriptor{
			Code: shader,
		},
	})
	if err != nil {
		log.Fatalf("CreateShaderModule:%v", err)
	}
	format := bridge.GetGPU().GetPreferredCanvasFormat()
	pipeline, err := device.CreateRenderPipeline(&wgpu.RenderPipelineDescriptor{
		//Layout: 'auto',
		Vertex: &wgpu.VertexState{
			Module:     shaderModule,
			EntryPoint: "main_vs",
		},
		Fragment: &wgpu.FragmentState{
			Module:     shaderModule,
			EntryPoint: "main_fs",
			Targets: []*wgpu.ColorTargetState{
				{
					Format:    format,
					Blend:     nil,
					WriteMask: wgpu.ColorWriteMask_All,
				},
			},
		},
		Primitive: &wgpu.PrimitiveState{
			Topology:  wgpu.PrimitiveTopology_TriangleList,
			FrontFace: wgpu.FrontFace_CCW,
		},
		Multisample: &wgpu.MultisampleState{
			Count:                  1,
			Mask:                   0xFFFFFFFF,
			AlphaToCoverageEnabled: false,
		},
	})
	if err != nil {
		log.Fatalf("CreateShaderModule:%v", err)
	}
	wasm.ConsoleLog(pipeline)
}
