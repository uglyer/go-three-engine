package wgpu_bridge

type AdapterDescriptor struct {
	PowerPreference PowerPreference
}

type SupportedLimits struct {
	Limits Limits
}

type AdapterProperties struct {
	VendorID          uint32
	DeviceID          uint32
	Name              string
	DriverDescription string
	AdapterType       AdapterType
	BackendType       BackendType
}
