//go:build windows

package wgpuext_glfw // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_native"
	"unsafe"
)

/*

#include <windows.h>

*/
import "C"

func GetSurfaceDescriptor(w *glfw.Window) *wgpu_native.SurfaceDescriptor {
	return &wgpu_native.SurfaceDescriptor{
		WindowsHWND: &wgpu_native.SurfaceDescriptorFromWindowsHWND{
			Hwnd:      unsafe.Pointer(w.GetWin32Window()),
			Hinstance: unsafe.Pointer(C.GetModuleHandle(nil)),
		},
	}
}
