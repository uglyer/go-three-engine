//go:build !wasm

package wgpu_native

func (p *Texture) AsImageCopy() *ImageCopyTexture {
	return &ImageCopyTexture{
		Texture:  p,
		MipLevel: 0,
		Origin:   Origin3D{},
		Aspect:   TextureAspect_All,
	}
}
