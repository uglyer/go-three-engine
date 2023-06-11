package wgpu_wasm

import "C"
import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
)

type Texture struct {
}

func (p *Texture) CreateView(descriptor *wgpu.TextureViewDescriptor) wgpu.ITextureView {
	panic("todo impl CreateView")
}

func (p *Texture) Destroy() {
	panic("todo impl Destroy")
}

func (p *Texture) Drop() {
	panic("todo impl Drop")
}
