package wgpu_bridge

type IDrop interface {
	Drop()
}

type IBridge interface {
	CreateCanvas(descriptor *CanvasDescriptor) (ICanvas, error)
}

// ICanvas 暂定等同于 GPUCanvasContext
type ICanvas interface {
	IDrop
	RequestAdapter(descriptor *AdapterDescriptor) (IAdapter, error)
	// Configure The configure() method of the GPUCanvasContext interface configures the context to use for
	// rendering with a given GPUDevice. When called the canvas will initially be cleared to transparent black.
	Configure(descriptor *ConfigureDescriptor) error
	// UnConfigure The unconfigure() method of the GPUCanvasContext interface removes any previously-set context
	// configuration, and destroys any textures returned via getCurrentTexture() while the canvas context was configured.
	UnConfigure()
	// getCurrentTexture The getCurrentTexture() method of the GPUCanvasContext interface returns the next GPUTexture
	// to be composited to the document by the canvas context.
	getCurrentTexture() IGpuTexture
}

type IAdapter interface {
	IDrop
	RequestDevice(descriptor *DeviceDescriptor) (IDevice, error)
}

type IDevice interface {
	IDrop
	GetQueue() IQueue
	CreateCommandEncoder() (IGpuCommandEncoder, error)
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
