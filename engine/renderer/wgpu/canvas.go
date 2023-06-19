package wgpu

type CanvasDescriptor struct {
	Width    int
	Height   int
	Title    string
	ParentId string
	CanvasId string
}

type ConfigureDescriptor struct {

	// Device The GPUDevice that the rendering information for the context will come from.
	Device IDevice

	// Format The format that textures returned by getCurrentTexture() will have.
	// This can be bgra8unorm, rgba8unorm, or rgba16float.
	// The optimal canvas texture format for the current system can be returned by GPU.getPreferredCanvasFormat().
	// Using this is recommended — if you don't use the preferred format when configuring the canvas context,
	// you may incur additional overhead, such as additional texture copies, depending on the platform.
	Format TextureFormat

	// AlphaMode Optional TODO AlphaMode
	//AlphaMode CompositeAlphaMode

	// ViewFormats Optional
	ViewFormats []TextureFormat
}
