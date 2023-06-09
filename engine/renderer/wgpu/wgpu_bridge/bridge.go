package wgpu_bridge

type IDrop interface {
	Drop()
}

type IBridge interface {
	CreateCanvas(descriptor *CanvasDescriptor) (ICanvas, error)
	GetGPU() (IGPU, error)
}

// IGPU navigator.gpu
// The GPU interface of the WebGPU API is the starting point for using WebGPU. It can be used to return a GPUAdapter
// from which you can request devices, configure features and limits, and more.
// The GPU object for the current context is accessed via the Navigator.gpu or WorkerNavigator.gpu properties.
type IGPU interface {
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
	getCurrentTexture() IGPUTexture
}

type IAdapter interface {
	IDrop
	// RequestDevice The requestDevice() method of the GPUAdapter interface returns a Promise that fulfills
	// with a GPUDevice object, which is the primary interface for communicating with the GPU.
	RequestDevice(descriptor *DeviceDescriptor) (IDevice, error)
}

type IDevice interface {
	IDrop
	// GetQueue A GPUQueue object instance.
	GetQueue() IGpuQueue
	// CreateBindGroup The createBindGroup() method of the GPUDevice interface creates a GPUBindGroup based
	// on a GPUBindGroupLayout that defines a set of resources to be bound together in a group and how those
	// resources are used in shader stages.
	CreateBindGroup(descriptor *GPUBindGroupDescriptor) (IBindGroup, error)
	// CreateBindGroupLayout The createBindGroupLayout() method of the GPUDevice interface creates a GPUBindGroupLayout
	// that defines the structure and purpose of related GPU resources such as buffers that will be used in a pipeline,
	// and is used as a template when creating GPUBindGroups.
	CreateBindGroupLayout(descriptor *GPUBindGroupDescriptor) (any, error)
	// CreateBuffer The createBuffer() method of the GPUDevice interface creates a GPUBuffer in which to store
	// raw data to use in GPU operations.
	CreateBuffer(descriptor *GPUBufferDescriptor) (IGPUBuffer, error)
	CreateCommandEncoder() (IGpuCommandEncoder, error)
	CreateTexture() (IGPUTexture, error)
	CreateRenderPipeline() (IGpuPipeLine, error)
}

type IGpuQueue interface {
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
	GetCurrentTextureView() (IGPUTextureView, error)
	Present()
}

type IGpuPipeLine interface {
	IDrop
}

type IGPUTextureView interface {
	IDrop
}

type IGPUTexture interface {
	IDrop
	CreateView() IGPUTextureView
}
