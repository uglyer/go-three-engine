package wgpu_bridge

type DeviceDescriptor struct {
	// Label 标记
	Label string
	// RequiredFeatures 需要支持的扩展
	RequiredFeatures []FeaturesName
	// RequiredLimits 资源限制
	RequiredLimits Limits
}
