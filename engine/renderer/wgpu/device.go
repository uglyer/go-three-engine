package wgpu

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
	ResourceBuffer      IBuffer
	Offset              uint64
	Size                uint64
	ResourceSampler     *IGPUSampler
	ResourceTextureView *ITextureView
}

type GPUBindGroupDescriptor struct {
	Label  string
	Layout IGPUBindGroupLayout
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
	Entries []*BindGroupLayoutEntry
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

type CommandEncoderDescriptor struct {
	Label string
}

type ShaderModuleSPIRVDescriptor struct {
	Code []byte
}

type ShaderModuleWGSLDescriptor struct {
	Code string
}

type ShaderModuleGLSLDescriptor struct {
	Code        string
	Defines     map[string]string
	ShaderStage ShaderStage
}

type ShaderModuleDescriptor struct {
	Label           string
	SPIRVDescriptor *ShaderModuleSPIRVDescriptor
	WGSLDescriptor  *ShaderModuleWGSLDescriptor
	GLSLDescriptor  *ShaderModuleGLSLDescriptor
}

type ProgrammableStageDescriptor struct {
	Module     IShaderModule
	EntryPoint string

	// unused in wgpu
	// Constants  []ConstantEntry
}

type ComputePipelineDescriptor struct {
	Label   string
	Layout  IPipelineLayout
	Compute ProgrammableStageDescriptor
}

type PushConstantRange struct {
	Stages ShaderStage
	Start  uint32
	End    uint32
}

type PipelineLayoutDescriptor struct {
	Label              string
	BindGroupLayouts   []IGPUBindGroupLayout
	PushConstantRanges []PushConstantRange
}

type RenderBundleEncoderDescriptor struct {
	Label              string
	ColorFormats       []TextureFormat
	DepthStencilFormat TextureFormat
	SampleCount        uint32
	DepthReadOnly      bool
	StencilReadOnly    bool
}

type WrappedSubmissionIndex struct {
	Queue           *IQueue
	SubmissionIndex SubmissionIndex
}

type BlendComponent struct {
	Operation BlendOperation
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

type BlendState struct {
	Color *BlendComponent
	Alpha *BlendComponent
}

type ColorTargetState struct {
	Format    TextureFormat
	Blend     *BlendState
	WriteMask ColorWriteMask
}

type FragmentState struct {
	Module     IShaderModule
	EntryPoint string
	Targets    []*ColorTargetState

	// unused in wgpu
	// Constants  []ConstantEntry
}

type VertexAttribute struct {
	Format         VertexFormat
	Offset         uint64
	ShaderLocation uint32
}

type VertexBufferLayout struct {
	ArrayStride uint64
	StepMode    VertexStepMode
	Attributes  []*VertexAttribute
}

type VertexState struct {
	Module     IShaderModule
	EntryPoint string
	Buffers    []VertexBufferLayout

	// unused in wgpu
	// Constants  []ConstantEntry
}

type PrimitiveState struct {
	Topology         PrimitiveTopology
	StripIndexFormat IndexFormat
	FrontFace        FrontFace
	CullMode         CullMode
}

type StencilFaceState struct {
	Compare     CompareFunction
	FailOp      StencilOperation
	DepthFailOp StencilOperation
	PassOp      StencilOperation
}

type DepthStencilState struct {
	Format              TextureFormat
	DepthWriteEnabled   bool
	DepthCompare        CompareFunction
	StencilFront        *StencilFaceState
	StencilBack         *StencilFaceState
	StencilReadMask     uint32
	StencilWriteMask    uint32
	DepthBias           int32
	DepthBiasSlopeScale float32
	DepthBiasClamp      float32
}

type MultisampleState struct {
	Count                  uint32
	Mask                   uint32
	AlphaToCoverageEnabled bool
}

type RenderPipelineDescriptor struct {
	Label        string
	Layout       IPipelineLayout
	Vertex       *VertexState
	Primitive    *PrimitiveState
	DepthStencil *DepthStencilState
	Multisample  *MultisampleState
	Fragment     *FragmentState
}

type TextureDescriptor struct {
	Label         string
	Usage         TextureUsage
	Dimension     TextureDimension
	Size          *Extent3D
	Format        TextureFormat
	MipLevelCount uint32
	SampleCount   uint32
}
