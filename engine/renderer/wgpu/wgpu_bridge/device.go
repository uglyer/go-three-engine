package wgpu_bridge

type IGPUBindGroupLayout interface {
	IDrop
}

type IGPUSampler interface {
	IDrop
}

type IBindGroup interface {
	IDrop
}

type GPUBindGroupEntry struct {
	// Binding A number representing a unique identifier for this resource binding, which matches the binding value
	// of a corresponding GPUBindGroupLayout entry. In addition, it matches the n index value of the corresponding
	// @binding(n) attribute in the shader (GPUShaderModule) used in the related pipeline.
	Binding uint32
	// Resource 开头三选一
	ResourceBuffer      IGPUBuffer
	Offset              uint64
	Size                uint64
	ResourceSampler     *IGPUSampler
	ResourceTextureView *IGPUTextureView
}

type GPUBindGroupDescriptor struct {
	Label  string
	Layout *IGPUBindGroupLayout
	// Entries An array of entry objects describing the resources to expose to the shader.
	// There will be one for each corresponding entry described by the GPUBindGroupLayout referenced in layout.
	// Each entry object has the following properties:
	Entries []*GPUBindGroupEntry
}

type DeviceDescriptor struct {
	// Label 标记
	Label string
	// RequiredFeatures 需要支持的扩展
	RequiredFeatures []FeatureName
	// RequiredLimits 资源限制
	RequiredLimits Limits
}

type BufferDescriptor struct {
	Label            string
	Usage            BufferUsage
	Size             uint64
	MappedAtCreation bool
}
