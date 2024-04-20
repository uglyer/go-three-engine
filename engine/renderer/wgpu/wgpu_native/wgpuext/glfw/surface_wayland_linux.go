//go:build linux && !android && wayland

package wgpuext_glfw // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_native"
	"unsafe"
)

func GetSurfaceDescriptor(w *glfw.Window) *wgpu_native.SurfaceDescriptor {
	return &wgpu_native.SurfaceDescriptor{
		WaylandSurface: &wgpu_native.SurfaceDescriptorFromWaylandSurface{
			Display: unsafe.Pointer(glfw.GetWaylandDisplay()),
			Surface: unsafe.Pointer(w.GetWaylandWindow()),
		},
	}
}
