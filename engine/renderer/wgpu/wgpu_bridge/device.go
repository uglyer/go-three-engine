package wgpu_bridge

type IGPUSampler interface {
	IDrop
}

type IGPUBindGroup interface {
	IDrop
}

type IGPUBindGroupLayout interface {
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

type BufferBindingLayout struct {
	Type             BufferBindingType
	HasDynamicOffset bool
	MinBindingSize   uint64
}

type SamplerBindingLayout struct {
	Type SamplerBindingType
}

type TextureBindingLayout struct {
	SampleType    TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled  bool
}

type StorageTextureBindingLayout struct {
	Access        StorageTextureAccess
	Format        TextureFormat
	ViewDimension TextureViewDimension
}

type BindGroupLayoutEntry struct {
	Binding        uint32
	Visibility     ShaderStage
	Buffer         BufferBindingLayout
	Sampler        SamplerBindingLayout
	Texture        TextureBindingLayout
	StorageTexture StorageTextureBindingLayout
}

type BindGroupLayoutDescriptor struct {
	Label   string
	Entries []BindGroupLayoutEntry
}

type DeviceDescriptor struct {
	// Label 标记
	Label string
	// RequiredFeatures 需要支持的扩展
	RequiredFeatures []FeatureName
	// RequiredLimits 资源限制
	RequiredLimits Limits
}

type GPUBufferDescriptor struct {
	Label            string
	Usage            BufferUsage
	Size             uint64
	MappedAtCreation bool
}

type GPUCommandEncoderDescriptor struct {
	Label string
}
