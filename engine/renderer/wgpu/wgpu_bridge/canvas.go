package wgpu_bridge

type CanvasDescriptor struct {
	Width    int
	Height   int
	Title    string
	ParentId string
	CanvasId string
}

type ConfigureDescriptor struct {
	Device    IDevice
	Format    any
	AlphaMode any
}
