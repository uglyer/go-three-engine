//go:build !wasm

package wgpu_native

/*

#include <stdlib.h>
#include "./lib/wgpu.h"

*/
import "C"
import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"unsafe"
)

type Surface struct {
	ref C.WGPUSurface
}

func (p *Surface) GetPreferredFormat(adapter *Adapter) wgpu.TextureFormat {
	format := C.wgpuSurfaceGetPreferredFormat(p.ref, adapter.ref)
	return wgpu.TextureFormat(format)
}

type SurfaceCapabilities struct {
	Formats      []wgpu.TextureFormat
	PresentModes []wgpu.PresentMode
	AlphaModes   []wgpu.CompositeAlphaMode
}

func (p *Surface) GetCapabilities(adapter *Adapter) (ret SurfaceCapabilities) {
	var caps C.WGPUSurfaceCapabilities
	C.wgpuSurfaceGetCapabilities(p.ref, adapter.ref, &caps)

	if caps.alphaModeCount == 0 && caps.formatCount == 0 && caps.presentModeCount == 0 {
		return
	}
	if caps.formatCount > 0 {
		caps.formats = (*C.WGPUTextureFormat)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))) * caps.formatCount))
		defer C.free(unsafe.Pointer(caps.formats))
	}
	if caps.presentModeCount > 0 {
		caps.presentModes = (*C.WGPUPresentMode)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUPresentMode(0))) * caps.presentModeCount))
		defer C.free(unsafe.Pointer(caps.presentModes))
	}
	if caps.alphaModeCount > 0 {
		caps.alphaModes = (*C.WGPUCompositeAlphaMode)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUCompositeAlphaMode(0))) * caps.alphaModeCount))
		defer C.free(unsafe.Pointer(caps.alphaModes))
	}

	C.wgpuSurfaceGetCapabilities(p.ref, adapter.ref, &caps)

	if caps.formatCount > 0 {
		formatsTmp := unsafe.Slice((*wgpu.TextureFormat)(caps.formats), caps.formatCount)
		ret.Formats = make([]wgpu.TextureFormat, caps.formatCount)
		copy(ret.Formats, formatsTmp)
	}
	if caps.presentModeCount > 0 {
		presentModesTmp := unsafe.Slice((*wgpu.PresentMode)(caps.presentModes), caps.presentModeCount)
		ret.PresentModes = make([]wgpu.PresentMode, caps.presentModeCount)
		copy(ret.PresentModes, presentModesTmp)
	}
	if caps.alphaModeCount > 0 {
		alphaModesTmp := unsafe.Slice((*wgpu.CompositeAlphaMode)(caps.alphaModes), caps.alphaModeCount)
		ret.AlphaModes = make([]wgpu.CompositeAlphaMode, caps.alphaModeCount)
		copy(ret.AlphaModes, alphaModesTmp)
	}

	return
}

func (p *Surface) Release() {
	C.wgpuSurfaceRelease(p.ref)
}
