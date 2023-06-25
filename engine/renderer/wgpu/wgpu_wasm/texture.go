package wgpu_wasm

import "C"
import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type Texture struct {
	ref js.Value
}

func newTexture(ref js.Value) *Texture {
	return &Texture{ref: ref}
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
