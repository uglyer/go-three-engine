//go:build windows

package wgpu_native // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

/*

#include <windows.h>

*/
import "C"

func GetSurfaceDescriptor(w *glfw.Window) *SurfaceDescriptor {
	return &SurfaceDescriptor{
		WindowsHWND: &SurfaceDescriptorFromWindowsHWND{
			Hwnd:      unsafe.Pointer(w.GetWin32Window()),
			Hinstance: unsafe.Pointer(C.GetModuleHandle(nil)),
		},
	}
}
