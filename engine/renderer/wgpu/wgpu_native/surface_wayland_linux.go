//go:build linux && !android && wayland

package wgpu_native // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

func GetSurfaceDescriptor(w *glfw.Window) *SurfaceDescriptor {
	return &SurfaceDescriptor{
		WaylandSurface: &SurfaceDescriptorFromWaylandSurface{
			Display: unsafe.Pointer(glfw.GetWaylandDisplay()),
			Surface: unsafe.Pointer(w.GetWaylandWindow()),
		},
	}
}
