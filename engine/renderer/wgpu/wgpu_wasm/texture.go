package wgpu_wasm

import "C"
import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type Texture struct {
	ref js.Value
}

func newTexture(ref js.Value) wgpu.ITexture {
	return &Texture{ref: ref}
}

func (p *Texture) CreateView(descriptor *wgpu.TextureViewDescriptor) wgpu.ITextureView {
	if descriptor == nil {
		ref := p.ref.Call("createView")
		return &TextureView{ref: ref}
	}
	desc := map[string]any{
		"label": descriptor.Label,
	}
	if &descriptor.ArrayLayerCount != nil {
		desc["arrayLayerCount"] = descriptor.ArrayLayerCount
	}
	if &descriptor.Aspect != nil {
		desc["aspect"] = descriptor.Aspect.String()
	}
	if &descriptor.BaseArrayLayer != nil {
		desc["baseArrayLayer"] = descriptor.BaseArrayLayer
	}
	if &descriptor.BaseMipLevel != nil {
		desc["baseMipLevel"] = descriptor.BaseMipLevel
	}
	if &descriptor.Dimension != nil {
		desc["dimension"] = descriptor.Dimension.String()
	}
	if &descriptor.Format != nil {
		desc["format"] = descriptor.Format.String()
	}
	ref := p.ref.Call("createView", desc)
	return &TextureView{ref: ref}
}

func (p *Texture) Drop() {
	panic("todo impl Drop")
}
