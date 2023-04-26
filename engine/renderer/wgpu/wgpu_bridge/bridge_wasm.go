//go:build wasm

package wgpu_bridge

type Bridge struct {
}

func NewBridge() IBridge {
	return &Bridge{}
}

func (b *Bridge) CreateCanvas(descriptor *CanvasDescriptor) ICanvas {
	return NewCanvas(descriptor)
}
