package wgpu

func DefaultLimits() Limits {
	return Limits{
		MaxTextureDimension1D:                     LimitU32Undefined,
		MaxTextureDimension2D:                     LimitU32Undefined,
		MaxTextureDimension3D:                     LimitU32Undefined,
		MaxTextureArrayLayers:                     LimitU32Undefined,
		MaxBindGroups:                             LimitU32Undefined,
		MaxBindingsPerBindGroup:                   LimitU32Undefined,
		MaxDynamicUniformBuffersPerPipelineLayout: LimitU32Undefined,
		MaxDynamicStorageBuffersPerPipelineLayout: LimitU32Undefined,
		MaxSampledTexturesPerShaderStage:          LimitU32Undefined,
		MaxSamplersPerShaderStage:                 LimitU32Undefined,
		MaxStorageBuffersPerShaderStage:           LimitU32Undefined,
		MaxStorageTexturesPerShaderStage:          LimitU32Undefined,
		MaxUniformBuffersPerShaderStage:           LimitU32Undefined,
		MaxUniformBufferBindingSize:               LimitU64Undefined,
		MaxStorageBufferBindingSize:               LimitU64Undefined,
		MinUniformBufferOffsetAlignment:           LimitU32Undefined,
		MinStorageBufferOffsetAlignment:           LimitU32Undefined,
		MaxVertexBuffers:                          LimitU32Undefined,
		MaxBufferSize:                             LimitU64Undefined,
		MaxVertexAttributes:                       LimitU32Undefined,
		MaxVertexBufferArrayStride:                LimitU32Undefined,
		MaxInterStageShaderComponents:             LimitU32Undefined,
		MaxInterStageShaderVariables:              LimitU32Undefined,
		MaxColorAttachments:                       LimitU32Undefined,
		MaxColorAttachmentBytesPerSample:          LimitU32Undefined,
		MaxComputeWorkgroupStorageSize:            LimitU32Undefined,
		MaxComputeInvocationsPerWorkgroup:         LimitU32Undefined,
		MaxComputeWorkgroupSizeX:                  LimitU32Undefined,
		MaxComputeWorkgroupSizeY:                  LimitU32Undefined,
		MaxComputeWorkgroupSizeZ:                  LimitU32Undefined,
		MaxComputeWorkgroupsPerDimension:          LimitU32Undefined,
		MaxPushConstantSize:                       LimitU32Undefined,
	}
}
