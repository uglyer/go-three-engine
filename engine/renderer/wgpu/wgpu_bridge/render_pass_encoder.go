package wgpu_bridge

type RenderPassColorAttachment struct {
	View          *IGPUTextureView
	ResolveTarget *IGPUTextureView
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

type RenderPassDepthStencilAttachment struct {
	View              *IGPUTextureView
	DepthLoadOp       LoadOp
	DepthStoreOp      StoreOp
	DepthClearValue   float32
	DepthReadOnly     bool
	StencilLoadOp     LoadOp
	StencilStoreOp    StoreOp
	StencilClearValue uint32
	StencilReadOnly   bool
}

type RenderPassDescriptor struct {
	Label                  string
	ColorAttachments       []RenderPassColorAttachment
	DepthStencilAttachment *RenderPassDepthStencilAttachment
}

type RenderPassEncoder interface {
	IDrop
	Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32)
	DrawIndexed(indexCount, instanceCount, firstIndex uint32, baseVertex int32, firstInstance uint32)
	DrawIndexedIndirect(indirectBuffer IGPUBuffer, indirectOffset uint64)
	DrawIndirect(indirectBuffer IGPUBuffer, indirectOffset uint64)
	End()
	ExecuteBundles(bundles ...IGPURenderBundle)
	InsertDebugMarker(markerLabel string)
	PopDebugGroup()
	PushDebugGroup(groupLabel string)
	SetBindGroup(groupIndex uint32, group IGPUBindGroup, dynamicOffsets []uint32)
	SetBlendConstant(color *Color)
	SetIndexBuffer(buffer IGPUBuffer, format IndexFormat, offset, size uint64)
	SetPipeline(pipeline IGPURenderPipeLine)
	SetScissorRect(x, y, width, height uint32)
	SetStencilReference(reference uint32)
	SetVertexBuffer(slot uint32, buffer IGPUBuffer, offset, size uint64)
	SetViewport(x, y, width, height, minDepth, maxDepth float32)
	SetPushConstants(stages ShaderStage, offset uint32, data []byte)
	MultiDrawIndirect(encoder *RenderPassEncoder, buffer IGPUBuffer, offset uint64, count uint32)
	MultiDrawIndexedIndirect(encoder *RenderPassEncoder, buffer IGPUBuffer, offset uint64, count uint32)
	MultiDrawIndirectCount(encoder *RenderPassEncoder, buffer IGPUBuffer, offset uint64, countBuffer IGPUBuffer, countBufferOffset uint64, maxCount uint32)
	MultiDrawIndexedIndirectCount(encoder *RenderPassEncoder, buffer IGPUBuffer, offset uint64, countBuffer IGPUBuffer, countBufferOffset uint64, maxCount uint32)
}

type CommandBufferDescriptor struct {
	Label string
}

type CommandBuffer interface {
	Drop()
}
