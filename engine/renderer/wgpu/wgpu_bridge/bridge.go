package wgpu_bridge

type IDrop interface {
	Drop()
}

type IBridge interface {
	CreateCanvas(descriptor *CanvasDescriptor) (ICanvas, error)
	GetGPU() (IGpu, error)
}

// IGpu navigator.gpu
// The GPU interface of the WebGPU API is the starting point for using WebGPU. It can be used to return a GPUAdapter
// from which you can request devices, configure features and limits, and more.
// The GPU object for the current context is accessed via the Navigator.gpu or WorkerNavigator.gpu properties.
type IGpu interface {
	// getPreferredCanvasFormat The getPreferredCanvasFormat() method of the GPU interface returns the optimal canvas
	// texture format for displaying 8-bit depth, standard dynamic range content on the current system.
	// This is commonly used to provide a GPUCanvasContext.configure() call with the optimal format value for the
	// current system. This is recommended — if you don't use the preferred format when configuring the canvas context,
	// you may incur additional overhead, such as additional texture copies, depending on the platform.
	getPreferredCanvasFormat() TextureFormat
	// RequestAdapter The requestAdapter() method of the GPU interface returns a Promise that fulfills with a GPUAdapter
	// object instance. From this you can request a GPUDevice, adapter info, features, and limits.
	// Note that the user agent chooses whether to return an adapter. If so, it chooses according to the provided
	// options. If no options are provided, the device will provide access to the default adapter, which is usually
	// good enough for most purposes.
	RequestAdapter(descriptor *AdapterDescriptor) (IAdapter, error)
}

// ICanvas 暂定等同于 GPUCanvasContext
type ICanvas interface {
	IDrop
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
