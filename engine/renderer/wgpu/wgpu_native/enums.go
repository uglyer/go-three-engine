// Code generated by github.com/rajveermalviya/go-webgpu/cmd/gen_enums. DO NOT EDIT.

//go:build !wasm

package wgpu_native

type AddressMode uint32

const (
	AddressMode_Repeat       AddressMode = 0x00000000
	AddressMode_MirrorRepeat AddressMode = 0x00000001
	AddressMode_ClampToEdge  AddressMode = 0x00000002
)

func (v AddressMode) String() string {
	switch v {
	case AddressMode_Repeat:
		return "Repeat"
	case AddressMode_MirrorRepeat:
		return "MirrorRepeat"
	case AddressMode_ClampToEdge:
		return "ClampToEdge"
	default:
		return ""
	}
}

type BlendFactor uint32

const (
	BlendFactor_Zero              BlendFactor = 0x00000000
	BlendFactor_One               BlendFactor = 0x00000001
	BlendFactor_Src               BlendFactor = 0x00000002
	BlendFactor_OneMinusSrc       BlendFactor = 0x00000003
	BlendFactor_SrcAlpha          BlendFactor = 0x00000004
	BlendFactor_OneMinusSrcAlpha  BlendFactor = 0x00000005
	BlendFactor_Dst               BlendFactor = 0x00000006
	BlendFactor_OneMinusDst       BlendFactor = 0x00000007
	BlendFactor_DstAlpha          BlendFactor = 0x00000008
	BlendFactor_OneMinusDstAlpha  BlendFactor = 0x00000009
	BlendFactor_SrcAlphaSaturated BlendFactor = 0x0000000A
	BlendFactor_Constant          BlendFactor = 0x0000000B
	BlendFactor_OneMinusConstant  BlendFactor = 0x0000000C
)

func (v BlendFactor) String() string {
	switch v {
	case BlendFactor_Zero:
		return "Zero"
	case BlendFactor_One:
		return "One"
	case BlendFactor_Src:
		return "Src"
	case BlendFactor_OneMinusSrc:
		return "OneMinusSrc"
	case BlendFactor_SrcAlpha:
		return "SrcAlpha"
	case BlendFactor_OneMinusSrcAlpha:
		return "OneMinusSrcAlpha"
	case BlendFactor_Dst:
		return "Dst"
	case BlendFactor_OneMinusDst:
		return "OneMinusDst"
	case BlendFactor_DstAlpha:
		return "DstAlpha"
	case BlendFactor_OneMinusDstAlpha:
		return "OneMinusDstAlpha"
	case BlendFactor_SrcAlphaSaturated:
		return "SrcAlphaSaturated"
	case BlendFactor_Constant:
		return "Constant"
	case BlendFactor_OneMinusConstant:
		return "OneMinusConstant"
	default:
		return ""
	}
}

type BlendOperation uint32

const (
	BlendOperation_Add             BlendOperation = 0x00000000
	BlendOperation_Subtract        BlendOperation = 0x00000001
	BlendOperation_ReverseSubtract BlendOperation = 0x00000002
	BlendOperation_Min             BlendOperation = 0x00000003
	BlendOperation_Max             BlendOperation = 0x00000004
)

func (v BlendOperation) String() string {
	switch v {
	case BlendOperation_Add:
		return "Add"
	case BlendOperation_Subtract:
		return "Subtract"
	case BlendOperation_ReverseSubtract:
		return "ReverseSubtract"
	case BlendOperation_Min:
		return "Min"
	case BlendOperation_Max:
		return "Max"
	default:
		return ""
	}
}

type BufferMapAsyncStatus uint32

const (
	BufferMapAsyncStatus_Success                 BufferMapAsyncStatus = 0x00000000
	BufferMapAsyncStatus_ValidationError         BufferMapAsyncStatus = 0x00000001
	BufferMapAsyncStatus_Unknown                 BufferMapAsyncStatus = 0x00000002
	BufferMapAsyncStatus_DeviceLost              BufferMapAsyncStatus = 0x00000003
	BufferMapAsyncStatus_DestroyedBeforeCallback BufferMapAsyncStatus = 0x00000004
	BufferMapAsyncStatus_UnmappedBeforeCallback  BufferMapAsyncStatus = 0x00000005
	BufferMapAsyncStatus_MappingAlreadyPending   BufferMapAsyncStatus = 0x00000006
	BufferMapAsyncStatus_OffsetOutOfRange        BufferMapAsyncStatus = 0x00000007
	BufferMapAsyncStatus_SizeOutOfRange          BufferMapAsyncStatus = 0x00000008
)

func (v BufferMapAsyncStatus) String() string {
	switch v {
	case BufferMapAsyncStatus_Success:
		return "Success"
	case BufferMapAsyncStatus_ValidationError:
		return "ValidationError"
	case BufferMapAsyncStatus_Unknown:
		return "Unknown"
	case BufferMapAsyncStatus_DeviceLost:
		return "DeviceLost"
	case BufferMapAsyncStatus_DestroyedBeforeCallback:
		return "DestroyedBeforeCallback"
	case BufferMapAsyncStatus_UnmappedBeforeCallback:
		return "UnmappedBeforeCallback"
	case BufferMapAsyncStatus_MappingAlreadyPending:
		return "MappingAlreadyPending"
	case BufferMapAsyncStatus_OffsetOutOfRange:
		return "OffsetOutOfRange"
	case BufferMapAsyncStatus_SizeOutOfRange:
		return "SizeOutOfRange"
	default:
		return ""
	}
}

type BufferMapState uint32

const (
	BufferMapState_Unmapped BufferMapState = 0x00000000
	BufferMapState_Pending  BufferMapState = 0x00000001
	BufferMapState_Mapped   BufferMapState = 0x00000002
)

func (v BufferMapState) String() string {
	switch v {
	case BufferMapState_Unmapped:
		return "Unmapped"
	case BufferMapState_Pending:
		return "Pending"
	case BufferMapState_Mapped:
		return "Mapped"
	default:
		return ""
	}
}


type ColorWriteMask uint32

const (
	ColorWriteMask_None  ColorWriteMask = 0x00000000
	ColorWriteMask_Red   ColorWriteMask = 0x00000001
	ColorWriteMask_Green ColorWriteMask = 0x00000002
	ColorWriteMask_Blue  ColorWriteMask = 0x00000004
	ColorWriteMask_Alpha ColorWriteMask = 0x00000008
	ColorWriteMask_All   ColorWriteMask = 0x0000000F
)

func (v ColorWriteMask) String() string {
	switch v {
	case ColorWriteMask_None:
		return "None"
	case ColorWriteMask_Red:
		return "Red"
	case ColorWriteMask_Green:
		return "Green"
	case ColorWriteMask_Blue:
		return "Blue"
	case ColorWriteMask_Alpha:
		return "Alpha"
	case ColorWriteMask_All:
		return "All"
	default:
		return ""
	}
}

type CompareFunction uint32

const (
	CompareFunction_Undefined    CompareFunction = 0x00000000
	CompareFunction_Never        CompareFunction = 0x00000001
	CompareFunction_Less         CompareFunction = 0x00000002
	CompareFunction_LessEqual    CompareFunction = 0x00000003
	CompareFunction_Greater      CompareFunction = 0x00000004
	CompareFunction_GreaterEqual CompareFunction = 0x00000005
	CompareFunction_Equal        CompareFunction = 0x00000006
	CompareFunction_NotEqual     CompareFunction = 0x00000007
	CompareFunction_Always       CompareFunction = 0x00000008
)

func (v CompareFunction) String() string {
	switch v {
	case CompareFunction_Undefined:
		return "Undefined"
	case CompareFunction_Never:
		return "Never"
	case CompareFunction_Less:
		return "Less"
	case CompareFunction_LessEqual:
		return "LessEqual"
	case CompareFunction_Greater:
		return "Greater"
	case CompareFunction_GreaterEqual:
		return "GreaterEqual"
	case CompareFunction_Equal:
		return "Equal"
	case CompareFunction_NotEqual:
		return "NotEqual"
	case CompareFunction_Always:
		return "Always"
	default:
		return ""
	}
}

type CompilationInfoRequestStatus uint32

const (
	CompilationInfoRequestStatus_Success    CompilationInfoRequestStatus = 0x00000000
	CompilationInfoRequestStatus_Error      CompilationInfoRequestStatus = 0x00000001
	CompilationInfoRequestStatus_DeviceLost CompilationInfoRequestStatus = 0x00000002
	CompilationInfoRequestStatus_Unknown    CompilationInfoRequestStatus = 0x00000003
)

func (v CompilationInfoRequestStatus) String() string {
	switch v {
	case CompilationInfoRequestStatus_Success:
		return "Success"
	case CompilationInfoRequestStatus_Error:
		return "Error"
	case CompilationInfoRequestStatus_DeviceLost:
		return "DeviceLost"
	case CompilationInfoRequestStatus_Unknown:
		return "Unknown"
	default:
		return ""
	}
}

type CompilationMessageType uint32

const (
	CompilationMessageType_Error   CompilationMessageType = 0x00000000
	CompilationMessageType_Warning CompilationMessageType = 0x00000001
	CompilationMessageType_Info    CompilationMessageType = 0x00000002
)

func (v CompilationMessageType) String() string {
	switch v {
	case CompilationMessageType_Error:
		return "Error"
	case CompilationMessageType_Warning:
		return "Warning"
	case CompilationMessageType_Info:
		return "Info"
	default:
		return ""
	}
}

type CompositeAlphaMode uint32

const (
	CompositeAlphaMode_Auto           CompositeAlphaMode = 0x00000000
	CompositeAlphaMode_Opaque         CompositeAlphaMode = 0x00000001
	CompositeAlphaMode_PreMultiplied  CompositeAlphaMode = 0x00000002
	CompositeAlphaMode_PostMultiplied CompositeAlphaMode = 0x00000003
	CompositeAlphaMode_Inherit        CompositeAlphaMode = 0x00000004
)

func (v CompositeAlphaMode) String() string {
	switch v {
	case CompositeAlphaMode_Auto:
		return "Auto"
	case CompositeAlphaMode_Opaque:
		return "Opaque"
	case CompositeAlphaMode_PreMultiplied:
		return "PreMultiplied"
	case CompositeAlphaMode_PostMultiplied:
		return "PostMultiplied"
	case CompositeAlphaMode_Inherit:
		return "Inherit"
	default:
		return ""
	}
}

type ComputePassTimestampLocation uint32

const (
	ComputePassTimestampLocation_Beginning ComputePassTimestampLocation = 0x00000000
	ComputePassTimestampLocation_End       ComputePassTimestampLocation = 0x00000001
)

func (v ComputePassTimestampLocation) String() string {
	switch v {
	case ComputePassTimestampLocation_Beginning:
		return "Beginning"
	case ComputePassTimestampLocation_End:
		return "End"
	default:
		return ""
	}
}

type CreatePipelineAsyncStatus uint32

const (
	CreatePipelineAsyncStatus_Success         CreatePipelineAsyncStatus = 0x00000000
	CreatePipelineAsyncStatus_ValidationError CreatePipelineAsyncStatus = 0x00000001
	CreatePipelineAsyncStatus_InternalError   CreatePipelineAsyncStatus = 0x00000002
	CreatePipelineAsyncStatus_DeviceLost      CreatePipelineAsyncStatus = 0x00000003
	CreatePipelineAsyncStatus_DeviceDestroyed CreatePipelineAsyncStatus = 0x00000004
	CreatePipelineAsyncStatus_Unknown         CreatePipelineAsyncStatus = 0x00000005
)

func (v CreatePipelineAsyncStatus) String() string {
	switch v {
	case CreatePipelineAsyncStatus_Success:
		return "Success"
	case CreatePipelineAsyncStatus_ValidationError:
		return "ValidationError"
	case CreatePipelineAsyncStatus_InternalError:
		return "InternalError"
	case CreatePipelineAsyncStatus_DeviceLost:
		return "DeviceLost"
	case CreatePipelineAsyncStatus_DeviceDestroyed:
		return "DeviceDestroyed"
	case CreatePipelineAsyncStatus_Unknown:
		return "Unknown"
	default:
		return ""
	}
}

type CullMode uint32

const (
	CullMode_None  CullMode = 0x00000000
	CullMode_Front CullMode = 0x00000001
	CullMode_Back  CullMode = 0x00000002
)

func (v CullMode) String() string {
	switch v {
	case CullMode_None:
		return "None"
	case CullMode_Front:
		return "Front"
	case CullMode_Back:
		return "Back"
	default:
		return ""
	}
}

type DeviceLostReason uint32

const (
	DeviceLostReason_Undefined DeviceLostReason = 0x00000000
	DeviceLostReason_Destroyed DeviceLostReason = 0x00000001
)

func (v DeviceLostReason) String() string {
	switch v {
	case DeviceLostReason_Undefined:
		return "Undefined"
	case DeviceLostReason_Destroyed:
		return "Destroyed"
	default:
		return ""
	}
}

type Dx12Compiler uint32

const (
	Dx12Compiler_Undefined Dx12Compiler = 0x00000000
	Dx12Compiler_Fxc       Dx12Compiler = 0x00000001
	Dx12Compiler_Dxc       Dx12Compiler = 0x00000002
)

func (v Dx12Compiler) String() string {
	switch v {
	case Dx12Compiler_Undefined:
		return "Undefined"
	case Dx12Compiler_Fxc:
		return "Fxc"
	case Dx12Compiler_Dxc:
		return "Dxc"
	default:
		return ""
	}
}

type ErrorFilter uint32

const (
	ErrorFilter_Validation  ErrorFilter = 0x00000000
	ErrorFilter_OutOfMemory ErrorFilter = 0x00000001
	ErrorFilter_Internal    ErrorFilter = 0x00000002
)

func (v ErrorFilter) String() string {
	switch v {
	case ErrorFilter_Validation:
		return "Validation"
	case ErrorFilter_OutOfMemory:
		return "OutOfMemory"
	case ErrorFilter_Internal:
		return "Internal"
	default:
		return ""
	}
}

type ErrorType uint32

const (
	ErrorType_NoError     ErrorType = 0x00000000
	ErrorType_Validation  ErrorType = 0x00000001
	ErrorType_OutOfMemory ErrorType = 0x00000002
	ErrorType_Internal    ErrorType = 0x00000003
	ErrorType_Unknown     ErrorType = 0x00000004
	ErrorType_DeviceLost  ErrorType = 0x00000005
)

func (v ErrorType) String() string {
	switch v {
	case ErrorType_NoError:
		return "NoError"
	case ErrorType_Validation:
		return "Validation"
	case ErrorType_OutOfMemory:
		return "OutOfMemory"
	case ErrorType_Internal:
		return "Internal"
	case ErrorType_Unknown:
		return "Unknown"
	case ErrorType_DeviceLost:
		return "DeviceLost"
	default:
		return "Unknown"
	}
}

type FeatureName uint32

const (
	FeatureName_Undefined                              FeatureName = 0x00000000
	FeatureName_DepthClipControl                       FeatureName = 0x00000001
	FeatureName_Depth32FloatStencil8                   FeatureName = 0x00000002
	FeatureName_TimestampQuery                         FeatureName = 0x00000003
	FeatureName_PipelineStatisticsQuery                FeatureName = 0x00000004
	FeatureName_TextureCompressionBC                   FeatureName = 0x00000005
	FeatureName_TextureCompressionETC2                 FeatureName = 0x00000006
	FeatureName_TextureCompressionASTC                 FeatureName = 0x00000007
	FeatureName_IndirectFirstInstance                  FeatureName = 0x00000008
	FeatureName_ShaderF16                              FeatureName = 0x00000009
	FeatureName_RG11B10UfloatRenderable                FeatureName = 0x0000000A
	FeatureName_BGRA8UnormStorage                      FeatureName = 0x0000000B
	FeatureName_Float32Filterable                      FeatureName = 0x0000000C
	NativeFeature_PushConstants                        FeatureName = 0x60000001
	NativeFeature_TextureAdapterSpecificFormatFeatures FeatureName = 0x60000002
	NativeFeature_MultiDrawIndirect                    FeatureName = 0x60000003
	NativeFeature_MultiDrawIndirectCount               FeatureName = 0x60000004
	NativeFeature_VertexWritableStorage                FeatureName = 0x60000005
)

func (v FeatureName) String() string {
	switch v {
	case FeatureName_Undefined:
		return "Undefined"
	case FeatureName_DepthClipControl:
		return "DepthClipControl"
	case FeatureName_Depth32FloatStencil8:
		return "Depth32FloatStencil8"
	case FeatureName_TimestampQuery:
		return "TimestampQuery"
	case FeatureName_PipelineStatisticsQuery:
		return "PipelineStatisticsQuery"
	case FeatureName_TextureCompressionBC:
		return "TextureCompressionBC"
	case FeatureName_TextureCompressionETC2:
		return "TextureCompressionETC2"
	case FeatureName_TextureCompressionASTC:
		return "TextureCompressionASTC"
	case FeatureName_IndirectFirstInstance:
		return "IndirectFirstInstance"
	case FeatureName_ShaderF16:
		return "ShaderF16"
	case FeatureName_RG11B10UfloatRenderable:
		return "RG11B10UfloatRenderable"
	case FeatureName_BGRA8UnormStorage:
		return "BGRA8UnormStorage"
	case FeatureName_Float32Filterable:
		return "Float32Filterable"
	case NativeFeature_PushConstants:
		return "NativeFeature_PushConstants"
	case NativeFeature_TextureAdapterSpecificFormatFeatures:
		return "NativeFeature_TextureAdapterSpecificFormatFeatures"
	case NativeFeature_MultiDrawIndirect:
		return "NativeFeature_MultiDrawIndirect"
	case NativeFeature_MultiDrawIndirectCount:
		return "NativeFeature_MultiDrawIndirectCount"
	case NativeFeature_VertexWritableStorage:
		return "NativeFeature_VertexWritableStorage"
	default:
		return ""
	}
}

type FilterMode uint32

const (
	FilterMode_Nearest FilterMode = 0x00000000
	FilterMode_Linear  FilterMode = 0x00000001
)

func (v FilterMode) String() string {
	switch v {
	case FilterMode_Nearest:
		return "Nearest"
	case FilterMode_Linear:
		return "Linear"
	default:
		return ""
	}
}

type FrontFace uint32

const (
	FrontFace_CCW FrontFace = 0x00000000
	FrontFace_CW  FrontFace = 0x00000001
)

func (v FrontFace) String() string {
	switch v {
	case FrontFace_CCW:
		return "CCW"
	case FrontFace_CW:
		return "CW"
	default:
		return ""
	}
}

type IndexFormat uint32

const (
	IndexFormat_Undefined IndexFormat = 0x00000000
	IndexFormat_Uint16    IndexFormat = 0x00000001
	IndexFormat_Uint32    IndexFormat = 0x00000002
)

func (v IndexFormat) String() string {
	switch v {
	case IndexFormat_Undefined:
		return "Undefined"
	case IndexFormat_Uint16:
		return "Uint16"
	case IndexFormat_Uint32:
		return "Uint32"
	default:
		return ""
	}
}

type InstanceBackend uint32

const (
	InstanceBackend_None          InstanceBackend = 0x00000000
	InstanceBackend_Vulkan        InstanceBackend = 0x00000002
	InstanceBackend_Metal         InstanceBackend = 0x00000004
	InstanceBackend_DX12          InstanceBackend = 0x00000008
	InstanceBackend_DX11          InstanceBackend = 0x00000010
	InstanceBackend_GL            InstanceBackend = 0x00000020
	InstanceBackend_Secondary     InstanceBackend = 0x00000030
	InstanceBackend_BrowserWebGPU InstanceBackend = 0x00000040
	InstanceBackend_Primary       InstanceBackend = 0x0000004E
)

func (v InstanceBackend) String() string {
	switch v {
	case InstanceBackend_None:
		return "None"
	case InstanceBackend_Vulkan:
		return "Vulkan"
	case InstanceBackend_Metal:
		return "Metal"
	case InstanceBackend_DX12:
		return "DX12"
	case InstanceBackend_DX11:
		return "DX11"
	case InstanceBackend_GL:
		return "GL"
	case InstanceBackend_Secondary:
		return "Secondary"
	case InstanceBackend_BrowserWebGPU:
		return "BrowserWebGPU"
	case InstanceBackend_Primary:
		return "Primary"
	default:
		return ""
	}
}

type LoadOp uint32

const (
	LoadOp_Undefined LoadOp = 0x00000000
	LoadOp_Clear     LoadOp = 0x00000001
	LoadOp_Load      LoadOp = 0x00000002
)

func (v LoadOp) String() string {
	switch v {
	case LoadOp_Undefined:
		return "Undefined"
	case LoadOp_Clear:
		return "Clear"
	case LoadOp_Load:
		return "Load"
	default:
		return ""
	}
}

type LogLevel uint32

const (
	LogLevel_Off   LogLevel = 0x00000000
	LogLevel_Error LogLevel = 0x00000001
	LogLevel_Warn  LogLevel = 0x00000002
	LogLevel_Info  LogLevel = 0x00000003
	LogLevel_Debug LogLevel = 0x00000004
	LogLevel_Trace LogLevel = 0x00000005
)

func (v LogLevel) String() string {
	switch v {
	case LogLevel_Off:
		return "Off"
	case LogLevel_Error:
		return "Error"
	case LogLevel_Warn:
		return "Warn"
	case LogLevel_Info:
		return "Info"
	case LogLevel_Debug:
		return "Debug"
	case LogLevel_Trace:
		return "Trace"
	default:
		return ""
	}
}

type MipmapFilterMode uint32

const (
	MipmapFilterMode_Nearest MipmapFilterMode = 0x00000000
	MipmapFilterMode_Linear  MipmapFilterMode = 0x00000001
)

func (v MipmapFilterMode) String() string {
	switch v {
	case MipmapFilterMode_Nearest:
		return "Nearest"
	case MipmapFilterMode_Linear:
		return "Linear"
	default:
		return ""
	}
}

type PresentMode uint32

const (
	PresentMode_Immediate PresentMode = 0x00000000
	PresentMode_Mailbox   PresentMode = 0x00000001
	PresentMode_Fifo      PresentMode = 0x00000002
)

func (v PresentMode) String() string {
	switch v {
	case PresentMode_Immediate:
		return "Immediate"
	case PresentMode_Mailbox:
		return "Mailbox"
	case PresentMode_Fifo:
		return "Fifo"
	default:
		return ""
	}
}

type PrimitiveTopology uint32

const (
	PrimitiveTopology_PointList     PrimitiveTopology = 0x00000000
	PrimitiveTopology_LineList      PrimitiveTopology = 0x00000001
	PrimitiveTopology_LineStrip     PrimitiveTopology = 0x00000002
	PrimitiveTopology_TriangleList  PrimitiveTopology = 0x00000003
	PrimitiveTopology_TriangleStrip PrimitiveTopology = 0x00000004
)

func (v PrimitiveTopology) String() string {
	switch v {
	case PrimitiveTopology_PointList:
		return "PointList"
	case PrimitiveTopology_LineList:
		return "LineList"
	case PrimitiveTopology_LineStrip:
		return "LineStrip"
	case PrimitiveTopology_TriangleList:
		return "TriangleList"
	case PrimitiveTopology_TriangleStrip:
		return "TriangleStrip"
	default:
		return ""
	}
}

type QueryType uint32

const (
	QueryType_Occlusion          QueryType = 0x00000000
	QueryType_PipelineStatistics QueryType = 0x00000001
	QueryType_Timestamp          QueryType = 0x00000002
)

func (v QueryType) String() string {
	switch v {
	case QueryType_Occlusion:
		return "Occlusion"
	case QueryType_PipelineStatistics:
		return "PipelineStatistics"
	case QueryType_Timestamp:
		return "Timestamp"
	default:
		return ""
	}
}

type QueueWorkDoneStatus uint32

const (
	QueueWorkDoneStatus_Success    QueueWorkDoneStatus = 0x00000000
	QueueWorkDoneStatus_Error      QueueWorkDoneStatus = 0x00000001
	QueueWorkDoneStatus_Unknown    QueueWorkDoneStatus = 0x00000002
	QueueWorkDoneStatus_DeviceLost QueueWorkDoneStatus = 0x00000003
)

func (v QueueWorkDoneStatus) String() string {
	switch v {
	case QueueWorkDoneStatus_Success:
		return "Success"
	case QueueWorkDoneStatus_Error:
		return "Error"
	case QueueWorkDoneStatus_Unknown:
		return "Unknown"
	case QueueWorkDoneStatus_DeviceLost:
		return "DeviceLost"
	default:
		return ""
	}
}

type RenderPassTimestampLocation uint32

const (
	RenderPassTimestampLocation_Beginning RenderPassTimestampLocation = 0x00000000
	RenderPassTimestampLocation_End       RenderPassTimestampLocation = 0x00000001
)

func (v RenderPassTimestampLocation) String() string {
	switch v {
	case RenderPassTimestampLocation_Beginning:
		return "Beginning"
	case RenderPassTimestampLocation_End:
		return "End"
	default:
		return ""
	}
}

type RequestAdapterStatus uint32

const (
	RequestAdapterStatus_Success     RequestAdapterStatus = 0x00000000
	RequestAdapterStatus_Unavailable RequestAdapterStatus = 0x00000001
	RequestAdapterStatus_Error       RequestAdapterStatus = 0x00000002
	RequestAdapterStatus_Unknown     RequestAdapterStatus = 0x00000003
)

func (v RequestAdapterStatus) String() string {
	switch v {
	case RequestAdapterStatus_Success:
		return "Success"
	case RequestAdapterStatus_Unavailable:
		return "Unavailable"
	case RequestAdapterStatus_Error:
		return "Error"
	case RequestAdapterStatus_Unknown:
		return "Unknown"
	default:
		return ""
	}
}

type RequestDeviceStatus uint32

const (
	RequestDeviceStatus_Success RequestDeviceStatus = 0x00000000
	RequestDeviceStatus_Error   RequestDeviceStatus = 0x00000001
	RequestDeviceStatus_Unknown RequestDeviceStatus = 0x00000002
)

func (v RequestDeviceStatus) String() string {
	switch v {
	case RequestDeviceStatus_Success:
		return "Success"
	case RequestDeviceStatus_Error:
		return "Error"
	case RequestDeviceStatus_Unknown:
		return "Unknown"
	default:
		return ""
	}
}

type StencilOperation uint32

const (
	StencilOperation_Keep           StencilOperation = 0x00000000
	StencilOperation_Zero           StencilOperation = 0x00000001
	StencilOperation_Replace        StencilOperation = 0x00000002
	StencilOperation_Invert         StencilOperation = 0x00000003
	StencilOperation_IncrementClamp StencilOperation = 0x00000004
	StencilOperation_DecrementClamp StencilOperation = 0x00000005
	StencilOperation_IncrementWrap  StencilOperation = 0x00000006
	StencilOperation_DecrementWrap  StencilOperation = 0x00000007
)

func (v StencilOperation) String() string {
	switch v {
	case StencilOperation_Keep:
		return "Keep"
	case StencilOperation_Zero:
		return "Zero"
	case StencilOperation_Replace:
		return "Replace"
	case StencilOperation_Invert:
		return "Invert"
	case StencilOperation_IncrementClamp:
		return "IncrementClamp"
	case StencilOperation_DecrementClamp:
		return "DecrementClamp"
	case StencilOperation_IncrementWrap:
		return "IncrementWrap"
	case StencilOperation_DecrementWrap:
		return "DecrementWrap"
	default:
		return ""
	}
}

type StoreOp uint32

const (
	StoreOp_Undefined StoreOp = 0x00000000
	StoreOp_Store     StoreOp = 0x00000001
	StoreOp_Discard   StoreOp = 0x00000002
)

func (v StoreOp) String() string {
	switch v {
	case StoreOp_Undefined:
		return "Undefined"
	case StoreOp_Store:
		return "Store"
	case StoreOp_Discard:
		return "Discard"
	default:
		return ""
	}
}

type TextureAspect uint32

const (
	TextureAspect_All         TextureAspect = 0x00000000
	TextureAspect_StencilOnly TextureAspect = 0x00000001
	TextureAspect_DepthOnly   TextureAspect = 0x00000002
)

func (v TextureAspect) String() string {
	switch v {
	case TextureAspect_All:
		return "All"
	case TextureAspect_StencilOnly:
		return "StencilOnly"
	case TextureAspect_DepthOnly:
		return "DepthOnly"
	default:
		return ""
	}
}

type TextureDimension uint32

const (
	TextureDimension_1D TextureDimension = 0x00000000
	TextureDimension_2D TextureDimension = 0x00000001
	TextureDimension_3D TextureDimension = 0x00000002
)

func (v TextureDimension) String() string {
	switch v {
	case TextureDimension_1D:
		return "1D"
	case TextureDimension_2D:
		return "2D"
	case TextureDimension_3D:
		return "3D"
	default:
		return ""
	}
}

type TextureUsage uint32

const (
	TextureUsage_None             TextureUsage = 0x00000000
	TextureUsage_CopySrc          TextureUsage = 0x00000001
	TextureUsage_CopyDst          TextureUsage = 0x00000002
	TextureUsage_TextureBinding   TextureUsage = 0x00000004
	TextureUsage_StorageBinding   TextureUsage = 0x00000008
	TextureUsage_RenderAttachment TextureUsage = 0x00000010
)

func (v TextureUsage) String() string {
	switch v {
	case TextureUsage_None:
		return "None"
	case TextureUsage_CopySrc:
		return "CopySrc"
	case TextureUsage_CopyDst:
		return "CopyDst"
	case TextureUsage_TextureBinding:
		return "TextureBinding"
	case TextureUsage_StorageBinding:
		return "StorageBinding"
	case TextureUsage_RenderAttachment:
		return "RenderAttachment"
	default:
		return ""
	}
}

type VertexFormat uint32

const (
	VertexFormat_Undefined VertexFormat = 0x00000000
	VertexFormat_Uint8x2   VertexFormat = 0x00000001
	VertexFormat_Uint8x4   VertexFormat = 0x00000002
	VertexFormat_Sint8x2   VertexFormat = 0x00000003
	VertexFormat_Sint8x4   VertexFormat = 0x00000004
	VertexFormat_Unorm8x2  VertexFormat = 0x00000005
	VertexFormat_Unorm8x4  VertexFormat = 0x00000006
	VertexFormat_Snorm8x2  VertexFormat = 0x00000007
	VertexFormat_Snorm8x4  VertexFormat = 0x00000008
	VertexFormat_Uint16x2  VertexFormat = 0x00000009
	VertexFormat_Uint16x4  VertexFormat = 0x0000000A
	VertexFormat_Sint16x2  VertexFormat = 0x0000000B
	VertexFormat_Sint16x4  VertexFormat = 0x0000000C
	VertexFormat_Unorm16x2 VertexFormat = 0x0000000D
	VertexFormat_Unorm16x4 VertexFormat = 0x0000000E
	VertexFormat_Snorm16x2 VertexFormat = 0x0000000F
	VertexFormat_Snorm16x4 VertexFormat = 0x00000010
	VertexFormat_Float16x2 VertexFormat = 0x00000011
	VertexFormat_Float16x4 VertexFormat = 0x00000012
	VertexFormat_Float32   VertexFormat = 0x00000013
	VertexFormat_Float32x2 VertexFormat = 0x00000014
	VertexFormat_Float32x3 VertexFormat = 0x00000015
	VertexFormat_Float32x4 VertexFormat = 0x00000016
	VertexFormat_Uint32    VertexFormat = 0x00000017
	VertexFormat_Uint32x2  VertexFormat = 0x00000018
	VertexFormat_Uint32x3  VertexFormat = 0x00000019
	VertexFormat_Uint32x4  VertexFormat = 0x0000001A
	VertexFormat_Sint32    VertexFormat = 0x0000001B
	VertexFormat_Sint32x2  VertexFormat = 0x0000001C
	VertexFormat_Sint32x3  VertexFormat = 0x0000001D
	VertexFormat_Sint32x4  VertexFormat = 0x0000001E
)

func (v VertexFormat) String() string {
	switch v {
	case VertexFormat_Undefined:
		return "Undefined"
	case VertexFormat_Uint8x2:
		return "Uint8x2"
	case VertexFormat_Uint8x4:
		return "Uint8x4"
	case VertexFormat_Sint8x2:
		return "Sint8x2"
	case VertexFormat_Sint8x4:
		return "Sint8x4"
	case VertexFormat_Unorm8x2:
		return "Unorm8x2"
	case VertexFormat_Unorm8x4:
		return "Unorm8x4"
	case VertexFormat_Snorm8x2:
		return "Snorm8x2"
	case VertexFormat_Snorm8x4:
		return "Snorm8x4"
	case VertexFormat_Uint16x2:
		return "Uint16x2"
	case VertexFormat_Uint16x4:
		return "Uint16x4"
	case VertexFormat_Sint16x2:
		return "Sint16x2"
	case VertexFormat_Sint16x4:
		return "Sint16x4"
	case VertexFormat_Unorm16x2:
		return "Unorm16x2"
	case VertexFormat_Unorm16x4:
		return "Unorm16x4"
	case VertexFormat_Snorm16x2:
		return "Snorm16x2"
	case VertexFormat_Snorm16x4:
		return "Snorm16x4"
	case VertexFormat_Float16x2:
		return "Float16x2"
	case VertexFormat_Float16x4:
		return "Float16x4"
	case VertexFormat_Float32:
		return "Float32"
	case VertexFormat_Float32x2:
		return "Float32x2"
	case VertexFormat_Float32x3:
		return "Float32x3"
	case VertexFormat_Float32x4:
		return "Float32x4"
	case VertexFormat_Uint32:
		return "Uint32"
	case VertexFormat_Uint32x2:
		return "Uint32x2"
	case VertexFormat_Uint32x3:
		return "Uint32x3"
	case VertexFormat_Uint32x4:
		return "Uint32x4"
	case VertexFormat_Sint32:
		return "Sint32"
	case VertexFormat_Sint32x2:
		return "Sint32x2"
	case VertexFormat_Sint32x3:
		return "Sint32x3"
	case VertexFormat_Sint32x4:
		return "Sint32x4"
	default:
		return ""
	}
}

type VertexStepMode uint32

const (
	VertexStepMode_Vertex              VertexStepMode = 0x00000000
	VertexStepMode_Instance            VertexStepMode = 0x00000001
	VertexStepMode_VertexBufferNotUsed VertexStepMode = 0x00000002
)

func (v VertexStepMode) String() string {
	switch v {
	case VertexStepMode_Vertex:
		return "Vertex"
	case VertexStepMode_Instance:
		return "Instance"
	case VertexStepMode_VertexBufferNotUsed:
		return "VertexBufferNotUsed"
	default:
		return ""
	}
}
