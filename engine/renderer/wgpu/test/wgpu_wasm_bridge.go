package main

import (
	_ "embed"
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_wasm"
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
	state := &State{
		Bridge: bridge,
	}
	err = state.Init()
	if err != nil {
		log.Fatalf("初始化失败:%v", err)
	}
	var renderFunc func()
	renderFunc = func() {
		// RequestAnimationFrame
		bridge.RequestAnimationFrame(renderFunc)
		//wasm.ConsoleLog("xxx")
		err = state.Render()
		if err != nil {
			log.Fatalf("渲染异常:%v", err)
		}
	}
	bridge.RequestAnimationFrame(renderFunc)
	//js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	//	wasm.ConsoleLog("xxxxx")
	//	return nil
	//}))
	//js.Global().Call("_requestAnimationFrame", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	//	wasm.ConsoleLog("111")
	//	return nil
	//}))
	// 避免主线程退出内容被销毁
	select {}
}

type State struct {
	Bridge   wgpu.IBridge
	Canvas   wgpu.ICanvas
	Device   wgpu.IDevice
	Adapter  wgpu.IAdapter
	Pipeline wgpu.IRenderPipeLine
	Config   *wgpu.TextureDescriptor
	View     wgpu.ITexture
}

func (s *State) Init() error {
	width := 640
	height := 480
	canvas, err := s.Bridge.CreateCanvas(&wgpu.CanvasDescriptor{
		Width:    width,
		Height:   height,
		Title:    "Test",
		ParentId: "root",
	})
	if err != nil {
		return fmt.Errorf("创建画布失败:%v", err)
	}
	s.Canvas = canvas
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
	format := s.Bridge.GetGPU().GetPreferredCanvasFormat()
	s.Config = &wgpu.TextureDescriptor{
		Size: &wgpu.Extent3D{
			Width:              uint32(width),
			Height:             uint32(height),
			DepthOrArrayLayers: 1,
		},
		SampleCount: 4,
		Dimension:   wgpu.TextureDimension_2D,
		Format:      format,
		Usage:       wgpu.TextureUsage_RenderAttachment,
	}
	err = canvas.Configure(&wgpu.ConfigureDescriptor{
		Device:    s.Device,
		Format:    format,
		AlphaMode: wgpu.CompositeAlphaMode_PreMultiplied,
	})
	if err != nil {
		return fmt.Errorf("配置画布失败:%v", err)
	}
	//{
	//size: [canvas.width, canvas.height],
	//sampleCount: 4,
	//format: presentationFormat,
	//usage: GPUTextureUsage.RENDER_ATTACHMENT,
	//}
	s.View, err = s.Device.CreateTexture(s.Config)
	if err != nil {
		return fmt.Errorf("创建纹理失败:%v", err)
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
					Format: format,
					//Blend:  nil,
					WriteMask: wgpu.ColorWriteMask_All,
				},
			},
		},
		Primitive: &wgpu.PrimitiveState{
			Topology:  wgpu.PrimitiveTopology_TriangleList,
			FrontFace: wgpu.FrontFace_CCW,
		},
		Multisample: &wgpu.MultisampleState{
			Count:                  4,
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
	currentTexture := s.Canvas.GetCurrentTexture().CreateView(nil)
	renderAttachment := s.View.CreateView(nil)

	encoder, err := s.Device.CreateCommandEncoder(&wgpu.CommandEncoderDescriptor{
		Label: "Command Encoder",
	})
	if err != nil {
		return err
	}

	renderPass := encoder.BeginRenderPass(&wgpu.RenderPassDescriptor{
		ColorAttachments: []*wgpu.RenderPassColorAttachment{
			{
				View:          renderAttachment,
				ResolveTarget: currentTexture,
				LoadOp:        wgpu.LoadOp_Clear,
				StoreOp:       wgpu.StoreOp_Store,
				ClearValue:    &wgpu.Color{R: 0.15, G: 0.15, B: 0.15, A: 1},
			},
		},
	})
	renderPass.SetPipeline(s.Pipeline)
	renderPass.Draw(3, 1, 0, 0)
	renderPass.End()

	s.Device.GetQueue().Submit(encoder.Finish(nil))
	return nil
}
