//go:build darwin

package wgpuext_glfw // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_native"
	"unsafe"
)

/*

#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework QuartzCore

#import <Cocoa/Cocoa.h>
#import <QuartzCore/CAMetalLayer.h>

CFTypeRef metalLayerFromNSWindow(CFTypeRef nsWindowRef) {
	NSWindow *ns_window = (__bridge NSWindow *)nsWindowRef;
	id metal_layer = NULL;
	[ns_window.contentView setWantsLayer:YES];
	metal_layer = [CAMetalLayer layer];
	[ns_window.contentView setLayer:metal_layer];
	return metal_layer;
}

*/
import "C"

func GetSurfaceDescriptor(w *glfw.Window) *wgpu_native.SurfaceDescriptor {
	return &wgpu_native.SurfaceDescriptor{
		MetalLayer: &wgpu_native.SurfaceDescriptorFromMetalLayer{
			Layer: unsafe.Pointer(C.metalLayerFromNSWindow((C.CFTypeRef)(unsafe.Pointer(w.GetCocoaWindow())))),
		},
	}
}
