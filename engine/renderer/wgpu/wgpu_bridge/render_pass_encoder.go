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
	Drop()
}

type CommandBufferDescriptor struct {
	Label string
}

type CommandBuffer interface {
	Drop()
}
