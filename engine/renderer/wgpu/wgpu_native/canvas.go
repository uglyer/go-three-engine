package wgpu_native

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
)

type Canvas struct {
	window    *glfw.Window
	surface   *Surface
	device    *Device
	config    *SwapChainDescriptor
	swapChain *SwapChain
}

func NewCanvas(instance *Instance, descriptor *wgpu.CanvasDescriptor) (*Canvas, error) {
	window, err := glfw.CreateWindow(descriptor.Width, descriptor.Width, descriptor.Title, nil, nil)
	if err != nil {
		return nil, err
	}
	surface := instance.CreateSurface(GetSurfaceDescriptor(window))
	adapter, err := instance.RequestAdapter(&RequestAdapterOptions{
		ForceFallbackAdapter: true,
		CompatibleSurface:    surface,
	})
	caps := surface.GetCapabilities(adapter)
	config := &SwapChainDescriptor{
		Usage:       wgpu.TextureUsage_RenderAttachment,
		Format:      caps.Formats[0],
		Width:       uint32(descriptor.Width),
		Height:      uint32(descriptor.Height),
		PresentMode: wgpu.PresentMode_Fifo,
		AlphaMode:   caps.AlphaModes[0],
	}
	device, err := adapter.RequestDevice(nil)
	if err != nil {
		return nil, err
	}
	swapChain, err := device.(*Device).CreateSwapChain(surface, config)
	if err != nil {
		return nil, err
	}
	return &Canvas{
		window:    window,
		surface:   surface,
		device:    device,
		config:    config,
		swapChain: swapChain,
	}, nil
}

func (c *Canvas) Release() {
	panic("TODO impl Canvas release")
}

func (c *Canvas) Configure(descriptor *wgpu.ConfigureDescriptor) error {
	panic("TODO impl Canvas Configure")
}

func (c *Canvas) UnConfigure() {
	panic("TODO impl Canvas UnConfigure")
}

func (c *Canvas) GetCurrentTexture() wgpu.ITexture {
	panic("TODO impl Canvas GetCurrentTexture")
	//return c.swapChain.GetCurrentTextureView()
}
