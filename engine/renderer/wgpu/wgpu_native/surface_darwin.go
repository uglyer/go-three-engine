//go:build darwin

package wgpu_native // import "github.com/rajveermalviya/go-webgpu/wgpuext/glfw"

import (
	"github.com/go-gl/glfw/v3.3/glfw"
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

func GetSurfaceDescriptor(w *glfw.Window) *SurfaceDescriptor {
	return &SurfaceDescriptor{
		MetalLayer: &SurfaceDescriptorFromMetalLayer{
			Layer: unsafe.Pointer(C.metalLayerFromNSWindow((C.CFTypeRef)(unsafe.Pointer(w.GetCocoaWindow())))),
		},
	}
}
