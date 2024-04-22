package wgpu

type AdapterDescriptor struct {
	PowerPreference PowerPreference
}

type SupportedLimits struct {
	Limits Limits
}

type AdapterProperties struct {
	VendorId          uint32
	VendorName        string
	Architecture      string
	DeviceId          uint32
	Name              string
	DriverDescription string
	AdapterType       AdapterType
	BackendType       BackendType
}
