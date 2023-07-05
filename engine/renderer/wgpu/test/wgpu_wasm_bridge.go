package main

import (
	_ "embed"
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_wasm"
	"log"
	"time"
)

//go:embed shader.wgsl
var shader string

func main() {

	//wasm.ConsoleLog(js.Global().Get("Object").New(map[string]any{"x": "xx"}))
	bridge, err := wgpu_wasm.NewBridge()
	if err != nil {
		log.Fatalf("获取桥接器失败:%v", err)
	}
	state := &State{
		Bridge: bridge,
	}
	err = state.Init()
	if err != nil {
		log.Fatalf("初始化失败:%v", err)
	}
	for {
		err = state.Render()
		if err != nil {
			log.Fatalf("渲染异常:%v", err)
		}
		time.Sleep(16 * time.Millisecond)
	}
}

type State struct {
	Bridge    wgpu.IBridge
	Canvas    wgpu.ICanvas
	Device    wgpu.IDevice
	Adapter   wgpu.IAdapter
	Pipeline  wgpu.IRenderPipeLine
	Config    *wgpu.TextureDescriptor
	SwapChain wgpu.ITexture
}

func (s *State) Init() error {
	canvas, err := s.Bridge.CreateCanvas(&wgpu.CanvasDescriptor{
		Width:    640,
		Height:   480,
		Title:    "Test",
		ParentId: "root",
	})
	if err != nil {
		return fmt.Errorf("创建画布失败:%v", err)
	}
	s.Adapter, err = s.Bridge.GetGPU().RequestAdapter(&wgpu.AdapterDescriptor{
		PowerPreference: wgpu.PowerPreference_HighPerformance,
	})
	if err != nil {
		return fmt.Errorf("获取适配器失败:%v", err)
	}
	s.Device, err = s.Adapter.RequestDevice(nil)
	if err != nil {
		return fmt.Errorf("获取gpu设备失败:%v", err)
	}
	s.Config = &wgpu.TextureDescriptor{
		Size: &wgpu.Extent3D{
			Width:              640,
			Height:             480,
			DepthOrArrayLayers: 1,
		},
		SampleCount: 1,
		Format:      wgpu.TextureFormat_RGBA8Unorm,
		Usage:       wgpu.TextureUsage_RenderAttachment,
	}
	s.SwapChain, err = s.Device.CreateTexture(s.Config)
	if err != nil {
		return err
	}
	err = canvas.Configure(&wgpu.ConfigureDescriptor{
		Device: s.Device,
	})
	if err != nil {
		return fmt.Errorf("配置画布失败:%v", err)
	}
	shaderModule, err := s.Device.CreateShaderModule(&wgpu.ShaderModuleDescriptor{
		Label: "draw.wgsl",
		WGSLDescriptor: &wgpu.ShaderModuleWGSLDescriptor{
			Code: shader,
		},
	})
	if err != nil {
		log.Fatalf("CreateShaderModule:%v", err)
	}
	format := s.Bridge.GetGPU().GetPreferredCanvasFormat()
	s.Pipeline, err = s.Device.CreateRenderPipeline(&wgpu.RenderPipelineDescriptor{
		//Layout: 'auto',
		Vertex: &wgpu.VertexState{
			Module:     shaderModule,
			EntryPoint: "vs_main",
		},
		Fragment: &wgpu.FragmentState{
			Module:     shaderModule,
			EntryPoint: "fs_main",
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
		return fmt.Errorf("CreateShaderModule:%v", err)
	}
	return nil
}

func (s *State) Resize(width, height int) error {
	return nil
}

func (s *State) Render() error {
	nextTexture := s.SwapChain.CreateView(nil)

	encoder, err := s.Device.CreateCommandEncoder(&wgpu.CommandEncoderDescriptor{
		Label: "Command Encoder",
	})
	if err != nil {
		return err
	}

	renderPass := encoder.BeginRenderPass(&wgpu.RenderPassDescriptor{
		ColorAttachments: []wgpu.RenderPassColorAttachment{
			{
				View:       nextTexture,
				LoadOp:     wgpu.LoadOp_Clear,
				StoreOp:    wgpu.StoreOp_Store,
				ClearValue: &wgpu.Color{G: 1, A: 1},
			},
		},
	})

	renderPass.SetPipeline(s.Pipeline)
	renderPass.Draw(3, 1, 0, 0)
	renderPass.End()

	s.Device.GetQueue().Submit(encoder.Finish(nil))
	return nil
}
