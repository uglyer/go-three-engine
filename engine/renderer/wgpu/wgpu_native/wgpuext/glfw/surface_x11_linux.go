//go:build linux && !android && !wayland

package wgpuext_glfw // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_native"
	"unsafe"
)

func GetSurfaceDescriptor(w *glfw.Window) *wgpu_native.SurfaceDescriptor {
	return &wgpu.SurfaceDescriptor{
		XlibWindow: &wgpu_native.SurfaceDescriptorFromXlibWindow{
			Display: unsafe.Pointer(glfw.GetX11Display()),
			Window:  uint32(w.GetX11Window()),
		},
	}
}
