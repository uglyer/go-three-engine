package wgpu

type IRelease interface {
	Release()
}

type IBridge interface {
	CreateCanvas(descriptor *CanvasDescriptor) (ICanvas, error)
	RequestAnimationFrame(fn func())
	GetGPU() IGPU
}

// IGPU navigator.gpu
// The GPU interface of the WebGPU API is the starting point for using WebGPU. It can be used to return a GPUAdapter
// from which you can request devices, configure features and limits, and more.
// The GPU object for the current context is accessed via the Navigator.gpu or WorkerNavigator.gpu properties.
type IGPU interface {
	// GetPreferredCanvasFormat The getPreferredCanvasFormat() method of the GPU interface returns the optimal canvas
	// texture format for displaying 8-bit depth, standard dynamic range content on the current system.
	// This is commonly used to provide a GPUCanvasContext.configure() call with the optimal format value for the
	// current system. This is recommended — if you don't use the preferred format when configuring the canvas context,
	// you may incur additional overhead, such as additional texture copies, depending on the platform.
	GetPreferredCanvasFormat() TextureFormat
	// RequestAdapter The requestAdapter() method of the GPU interface returns a Promise that fulfills with a GPUAdapter
	// object instance. From this you can request a GPUDevice, adapter info, features, and limits.
	// Note that the user agent chooses whether to return an adapter. If so, it chooses according to the provided
	// options. If no options are provided, the device will provide access to the default adapter, which is usually
	// good enough for most purposes.
	RequestAdapter(descriptor *AdapterDescriptor) (IAdapter, error)
}

// ICanvas 暂定等同于 GPUCanvasContext
type ICanvas interface {
	IRelease
	// Configure The configure() method of the GPUCanvasContext interface configures the context to use for
	// rendering with a given GPUDevice. When called the canvas will initially be cleared to transparent black.
	Configure(descriptor *ConfigureDescriptor) error
	// UnConfigure The unconfigure() method of the GPUCanvasContext interface removes any previously-set context
	// configuration, and destroys any textures returned via getCurrentTexture() while the canvas context was configured.
	UnConfigure()
	// GetCurrentTexture The getCurrentTexture() method of the GPUCanvasContext interface returns the next GPUTexture
	// to be composited to the document by the canvas context.
	GetCurrentTexture() ITexture
}

type IAdapter interface {
	IRelease
	// RequestDevice The requestDevice() method of the GPUAdapter interface returns a Promise that fulfills
	// with a GPUDevice object, which is the primary interface for communicating with the GPU.
	RequestDevice(descriptor *DeviceDescriptor) (IDevice, error)
	EnumerateFeatures() []FeatureName
	GetLimits() SupportedLimits
	GetProperties() AdapterProperties
	HasFeature(feature FeatureName) bool
}

type IDevice interface {
	IRelease
	// GetQueue A GPUQueue object instance.
	GetQueue() IQueue
	// CreateBindGroup The createBindGroup() method of the GPUDevice interface creates a GPUBindGroup based
	// on a GPUBindGroupLayout that defines a set of resources to be bound together in a group and how those
	// resources are used in shader stages.
	CreateBindGroup(descriptor *GPUBindGroupDescriptor) (IGPUBindGroup, error)
	// CreateBindGroupLayout The createBindGroupLayout() method of the GPUDevice interface creates a GPUBindGroupLayout
	// that defines the structure and purpose of related GPU resources such as buffers that will be used in a pipeline,
	// and is used as a template when creating GPUBindGroups.
	CreateBindGroupLayout(descriptor *BindGroupLayoutDescriptor) (IGPUBindGroupLayout, error)
	// CreateBuffer The createBuffer() method of the GPUDevice interface creates a GPUBuffer in which to store
	// raw data to use in GPU operations.
	CreateBuffer(descriptor *BufferDescriptor) (IBuffer, error)
	// CreateCommandEncoder The createCommandEncoder() method of the GPUDevice interface creates a GPUCommandEncoder,
	// used to encode commands to be issued to the GPU.
	CreateCommandEncoder(descriptor *CommandEncoderDescriptor) (ICommandEncoder, error)
	CreateShaderModule(descriptor *ShaderModuleDescriptor) (IShaderModule, error)
	CreateTexture(descriptor *TextureDescriptor) (ITexture, error)
	CreateRenderPipeline(descriptor *RenderPipelineDescriptor) (IRenderPipeLine, error)
	GetErr() (err error)
	StoreErr(typ ErrorType, message string)
	CreateComputePipeline(descriptor *ComputePipelineDescriptor) (IComputePipeline, error)
	CreatePipelineLayout(descriptor *PipelineLayoutDescriptor) (IPipelineLayout, error)
	CreateRenderBundleEncoder(descriptor *RenderBundleEncoderDescriptor) (IRenderBundleEncoder, error)
	Poll(wait bool, wrappedSubmissionIndex *WrappedSubmissionIndex) (queueEmpty bool)
}

type IQueue interface {
	WriteTexture(destination *ImageCopyTexture, data []byte, dataLayout *TextureDataLayout, writeSize *Extent3D)
	WriteBuffer(buffer IBuffer, bufferOffset uint64, data []byte)
	Submit(commands ...ICommandBuffer) (submissionIndex SubmissionIndex)
}

type ICommandEncoder interface {
	IRelease
	BeginComputePass(descriptor *ComputePassDescriptor) IComputePassEncoder
	BeginRenderPass(descriptor *RenderPassDescriptor) IRenderPassEncoder
	ClearBuffer(buffer IBuffer, offset uint64, size uint64)
	CopyBufferToBuffer(source IBuffer, sourceOffset uint64, destination IBuffer, destinatonOffset uint64, size uint64)
	CopyBufferToTexture(source *ImageCopyBuffer, destination *ImageCopyTexture, copySize *Extent3D)
	CopyTextureToBuffer(source *ImageCopyTexture, destination *ImageCopyBuffer, copySize *Extent3D)
	CopyTextureToTexture(source *ImageCopyTexture, destination *ImageCopyTexture, copySize *Extent3D)
	Finish(descriptor *CommandBufferDescriptor) ICommandBuffer
	InsertDebugMarker(markerLabel string)
	PopDebugGroup()
	PushDebugGroup(groupLabel string)
}

type IComputePassEncoder interface {
	DispatchWorkgroups(workgroupCountX, workgroupCountY, workgroupCountZ uint32)
	DispatchWorkgroupsIndirect(indirectBuffer IBuffer, indirectOffset uint64)
	End()
	InsertDebugMarker(markerLabel string)
	PopDebugGroup()
	PushDebugGroup(groupLabel string)
	SetBindGroup(groupIndex uint32, group IGPUBindGroup, dynamicOffsets []uint32)
	SetPipeline(pipeline IComputePipeline)
}

type IRenderPassEncoder interface {
	IRelease
	Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32)
	DrawIndexed(indexCount, instanceCount, firstIndex uint32, baseVertex int32, firstInstance uint32)
	DrawIndexedIndirect(indirectBuffer IBuffer, indirectOffset uint64)
	DrawIndirect(indirectBuffer IBuffer, indirectOffset uint64)
	End()
	ExecuteBundles(bundles ...IRenderBundle)
	InsertDebugMarker(markerLabel string)
	PopDebugGroup()
	PushDebugGroup(groupLabel string)
	SetBindGroup(groupIndex uint32, group IGPUBindGroup, dynamicOffsets []uint32)
	SetBlendConstant(color *Color)
	SetIndexBuffer(buffer IBuffer, format IndexFormat, offset, size uint64)
	SetPipeline(pipeline IRenderPipeLine)
	SetScissorRect(x, y, width, height uint32)
	SetStencilReference(reference uint32)
	SetVertexBuffer(slot uint32, buffer IBuffer, offset, size uint64)
	SetViewport(x, y, width, height, minDepth, maxDepth float32)
	SetPushConstants(stages ShaderStage, offset uint32, data []byte)
	MultiDrawIndirect(encoder IRenderPassEncoder, buffer IBuffer, offset uint64, count uint32)
	MultiDrawIndexedIndirect(encoder IRenderPassEncoder, buffer IBuffer, offset uint64, count uint32)
	MultiDrawIndirectCount(encoder IRenderPassEncoder, buffer IBuffer, offset uint64, countBuffer IBuffer, countBufferOffset uint64, maxCount uint32)
	MultiDrawIndexedIndirectCount(encoder IRenderPassEncoder, buffer IBuffer, offset uint64, countBuffer IBuffer, countBufferOffset uint64, maxCount uint32)
}

type IRenderBundleEncoder interface {
	IRelease
	Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32)
	DrawIndexed(indexCount, instanceCount, firstIndex, baseVertex, firstInstance uint32)
	DrawIndexedIndirect(indirectBuffer IBuffer, indirectOffset uint64)
	DrawIndirect(indirectBuffer IBuffer, indirectOffset uint64)
	Finish(descriptor *RenderBundleDescriptor) IRenderBundle
	InsertDebugMarker(markerLabel string)
	PopDebugGroup()
	PushDebugGroup(groupLabel string)
	SetBindGroup(groupIndex uint32, group IGPUBindGroup, dynamicOffsets []uint32)
	SetIndexBuffer(buffer IBuffer, format IndexFormat, offset, size uint64)
	SetPipeline(pipeline IRenderPipeLine)
	SetVertexBuffer(slot uint32, buffer IBuffer, offset, size uint64)
}

type IComputePipeline interface {
	IRelease
	GetBindGroupLayout(groupIndex uint32) IGPUBindGroupLayout
}

type IRenderPipeLine interface {
	IRelease
	GetBindGroupLayout(groupIndex uint32) IGPUBindGroupLayout
}

type ITextureView interface {
	IRelease
}

type ITexture interface {
	IRelease
	CreateView(descriptor *TextureViewDescriptor) ITextureView
}

type IShaderModule interface {
	IRelease
}

type IPipelineLayout interface {
	IRelease
}

type IRenderBundle interface {
	IRelease
}
type ICommandBuffer interface {
	IRelease
}
