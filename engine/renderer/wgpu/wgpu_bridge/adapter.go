package wgpu_bridge

type DeviceDescriptor struct {
	// Label 标记
	Label string
	// Extensions 需要支持的扩展
	Extensions []DeviceExtensionName
	// Limits 资源限制
	Limits Limits
}
