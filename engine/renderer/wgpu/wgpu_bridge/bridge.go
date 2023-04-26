package wgpu_bridge

type IDrop interface {
	Drop()
}

type IBridge interface {
	IDrop
	CreateCanvas(width, height int, title string) ICanvas
}

type ICanvas interface {
	IDrop
	RequestAdapter() IAdapter
}
type IAdapter interface {
	IDrop
	RequestDevice(descriptor *DeviceDescriptor) IDevice
}

type IDevice interface {
	IDrop
	GetQueue() IQueue
	CreateCommandEncoder() (IGpuCommandEncoder, error)
	CreateSwapChain() (IGpuSwapChain, error)
	CreateBuffer() (IGpuBuffer, error)
	CreateTexture() (IGpuTexture, error)
	CreateRenderPipeline() (IGpuPipeLine, error)
}

type IQueue interface {
	WriteTexture()
	WriteBuffer()
	Submit(commandBuffer ...IGpuCommandBuffer)
}

type IGpuCommandEncoder interface {
	BeginRenderPass()
	Finish() IGpuCommandBuffer
}

type IRenderPass interface {
	SetPipeline() IRenderPass
	SetBindGroup() IRenderPass
	SetIndexBuffer() IRenderPass
	SetVertexBuffer() IRenderPass
	DrawIndexed() IRenderPass
	EndPass() IRenderPass
}

type IGpuCommandBuffer interface {
}

type IGpuSwapChain interface {
	IDrop
	GetCurrentTextureView() (IGpuTextureView, error)
	Present()
}

type IGpuPipeLine interface {
	IDrop
}

type IGpuTextureView interface {
	IDrop
}

type IGpuBuffer interface {
	IDrop
}
type IGpuTexture interface {
	IDrop
	CreateView() IGpuTextureView
}
