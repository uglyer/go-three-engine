package wgpu_bridge

type Limits struct {
	MaxTextureDimension1D                     uint32
	MaxTextureDimension2D                     uint32
	MaxTextureDimension3D                     uint32
	MaxTextureArrayLayers                     uint32
	MaxBindGroups                             uint32
	MaxBindingsPerBindGroup                   uint32
	MaxDynamicUniformBuffersPerPipelineLayout uint32
	MaxDynamicStorageBuffersPerPipelineLayout uint32
	MaxSampledTexturesPerShaderStage          uint32
	MaxSamplersPerShaderStage                 uint32
	MaxStorageBuffersPerShaderStage           uint32
	MaxStorageTexturesPerShaderStage          uint32
	MaxUniformBuffersPerShaderStage           uint32
	MaxUniformBufferBindingSize               uint64
	MaxStorageBufferBindingSize               uint64
	MinUniformBufferOffsetAlignment           uint32
	MinStorageBufferOffsetAlignment           uint32
	MaxVertexBuffers                          uint32
	MaxBufferSize                             uint64
	MaxVertexAttributes                       uint32
	MaxVertexBufferArrayStride                uint32
	MaxInterStageShaderComponents             uint32
	MaxInterStageShaderVariables              uint32
	MaxColorAttachments                       uint32
	MaxComputeWorkgroupStorageSize            uint32
	MaxComputeInvocationsPerWorkgroup         uint32
	MaxComputeWorkgroupSizeX                  uint32
	MaxComputeWorkgroupSizeY                  uint32
	MaxComputeWorkgroupSizeZ                  uint32
	MaxComputeWorkgroupsPerDimension          uint32

	MaxPushConstantSize uint32
}

type FeaturesName uint32

const (
	DeviceExtensionName_Undefined                          FeaturesName = 0x00000000
	DeviceExtensionName_DepthClipControl                   FeaturesName = 0x00000001
	DeviceExtensionName_Depth32FloatStencil8               FeaturesName = 0x00000002
	DeviceExtensionName_TimestampQuery                     FeaturesName = 0x00000003
	DeviceExtensionName_PipelineStatisticsQuery            FeaturesName = 0x00000004
	DeviceExtensionName_TextureCompressionBC               FeaturesName = 0x00000005
	DeviceExtensionName_TextureCompressionETC2             FeaturesName = 0x00000006
	DeviceExtensionName_TextureCompressionASTC             FeaturesName = 0x00000007
	DeviceExtensionName_IndirectFirstInstance              FeaturesName = 0x00000008
	DeviceExtensionName_ShaderF16                          FeaturesName = 0x00000009
	DeviceExtensionName_RG11B10UfloatRenderable            FeaturesName = 0x0000000A
	NativeFeature_PUSH_CONSTANTS                           FeaturesName = 0x60000001
	NativeFeature_TEXTURE_ADAPTER_SPECIFIC_FORMAT_FEATURES FeaturesName = 0x60000002
	NativeFeature_MULTI_DRAW_INDIRECT                      FeaturesName = 0x60000003
	NativeFeature_MULTI_DRAW_INDIRECT_COUNT                FeaturesName = 0x60000004
	NativeFeature_VERTEX_WRITABLE_STORAGE                  FeaturesName = 0x60000005
)

type PowerPreference uint32

const (
	PowerPreference_Default  PowerPreference = 0x00000000
	PowerPreference_LowPower PowerPreference = 0x00000001
)
