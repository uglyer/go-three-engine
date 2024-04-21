//go:build !wasm

package wgpu_native

/*

#include <stdlib.h>
#include "./lib/wgpu.h"

extern void gowebgpu_error_callback_c(WGPUErrorType type, char const * message, void * userdata);

static inline WGPUTextureView gowebgpu_texture_create_view(WGPUTexture texture, WGPUTextureViewDescriptor const * descriptor, WGPUDevice device, void * error_userdata) {
	WGPUTextureView ref = NULL;
	wgpuDevicePushErrorScope(device, WGPUErrorFilter_Validation);
	ref = wgpuTextureCreateView(texture, descriptor);
	wgpuDevicePopErrorScope(device, gowebgpu_error_callback_c, error_userdata);
	return ref;
}

static inline void gowebgpu_texture_release(WGPUTexture texture, WGPUDevice device) {
	wgpuDeviceRelease(device);
	wgpuTextureRelease(texture);
}

*/
import "C"
import (
	"errors"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"runtime/cgo"
	"unsafe"
)

type Texture struct {
	deviceRef C.WGPUDevice
	ref       C.WGPUTexture
}

type TextureViewDescriptor struct {
	Label           string
	Format          wgpu.TextureFormat
	Dimension       wgpu.TextureViewDimension
	BaseMipLevel    uint32
	MipLevelCount   uint32
	BaseArrayLayer  uint32
	ArrayLayerCount uint32
	Aspect          wgpu.TextureAspect
}

func (p *Texture) CreateView(descriptor *wgpu.TextureViewDescriptor) (*wgpu.ITextureView, error) {
	var desc *C.WGPUTextureViewDescriptor

	if descriptor != nil {
		desc = &C.WGPUTextureViewDescriptor{
			format:          C.WGPUTextureFormat(descriptor.Format),
			dimension:       C.WGPUTextureViewDimension(descriptor.Dimension),
			baseMipLevel:    C.uint32_t(descriptor.BaseMipLevel),
			mipLevelCount:   C.uint32_t(descriptor.MipLevelCount),
			baseArrayLayer:  C.uint32_t(descriptor.BaseArrayLayer),
			arrayLayerCount: C.uint32_t(descriptor.ArrayLayerCount),
			aspect:          C.WGPUTextureAspect(descriptor.Aspect),
		}

		if descriptor.Label != "" {
			label := C.CString(descriptor.Label)
			defer C.free(unsafe.Pointer(label))

			desc.label = label
		}
	}

	var err error = nil
	var cb errorCallback = func(_ ErrorType, message string) {
		err = errors.New("wgpu.(*Texture).CreateView(): " + message)
	}
	errorCallbackHandle := cgo.NewHandle(cb)
	defer errorCallbackHandle.Delete()

	ref := C.gowebgpu_texture_create_view(
		p.ref,
		desc,
		p.deviceRef,
		unsafe.Pointer(&errorCallbackHandle),
	)
	if err != nil {
		C.wgpuTextureViewRelease(ref)
		return nil, err
	}

	return &TextureView{ref}, nil
}

func (p *Texture) Destroy() {
	C.wgpuTextureDestroy(p.ref)
}

func (p *Texture) GetDepthOrArrayLayers() uint32 {
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(p.ref))
}

func (p *Texture) GetDimension() wgpu.TextureDimension {
	return wgpu.TextureDimension(C.wgpuTextureGetDimension(p.ref))
}

func (p *Texture) GetFormat() wgpu.TextureFormat {
	return wgpu.TextureFormat(C.wgpuTextureGetFormat(p.ref))
}

func (p *Texture) GetHeight() uint32 {
	return uint32(C.wgpuTextureGetHeight(p.ref))
}

func (p *Texture) GetMipLevelCount() uint32 {
	return uint32(C.wgpuTextureGetMipLevelCount(p.ref))
}

func (p *Texture) GetSampleCount() uint32 {
	return uint32(C.wgpuTextureGetSampleCount(p.ref))
}

func (p *Texture) GetUsage() TextureUsage {
	return TextureUsage(C.wgpuTextureGetUsage(p.ref))
}

func (p *Texture) GetWidth() uint32 {
	return uint32(C.wgpuTextureGetWidth(p.ref))
}

func (p *Texture) Release() {
	C.gowebgpu_texture_release(p.ref, p.deviceRef)
}
