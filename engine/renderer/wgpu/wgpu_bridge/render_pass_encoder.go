package wgpu_bridge

type RenderPassColorAttachment struct {
	View          ITextureView
	ResolveTarget ITextureView
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

type RenderPassDepthStencilAttachment struct {
	View        ITextureView
	DepthLoadOp LoadOp
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

type CommandBufferDescriptor struct {
	Label string
}

type CommandBuffer interface {
	Drop()
}
