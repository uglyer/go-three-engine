//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type CommandEncoder struct {
	// fields for CommandEncoder struct
	ref js.Value
}

func (ce *CommandEncoder) Drop() {
	// TODO impl BeginComputePass
	return
}

func (ce *CommandEncoder) BeginComputePass(descriptor *wgpu.ComputePassDescriptor) wgpu.IComputePassEncoder {
	if descriptor == nil {
		ref := ce.ref.Call("beginComputePass")
		return newComputePassEncoder(ref)
	}
	desc := map[string]any{
		"label": descriptor.Label,
	}
	ref := ce.ref.Call("beginComputePass", desc)
	return newComputePassEncoder(ref)
}

func (ce *CommandEncoder) BeginRenderPass(descriptor *wgpu.RenderPassDescriptor) wgpu.IRenderPassEncoder {
	// TODO impl BeginRenderPass
	return nil
}

func (ce *CommandEncoder) ClearBuffer(buffer wgpu.IBuffer, offset uint64, size uint64) {
	// TODO impl ClearBuffer
}

func (ce *CommandEncoder) CopyBufferToBuffer(source wgpu.IBuffer, sourceOffset uint64, destination wgpu.IBuffer, destinatonOffset uint64, size uint64) {
	// TODO impl CopyBufferToBuffer
}

func (ce *CommandEncoder) CopyBufferToTexture(source *wgpu.ImageCopyBuffer, destination *wgpu.ImageCopyTexture, copySize *wgpu.Extent3D) {
	// TODO impl CopyBufferToTexture
}

func (ce *CommandEncoder) CopyTextureToBuffer(source *wgpu.ImageCopyTexture, destination *wgpu.ImageCopyBuffer, copySize *wgpu.Extent3D) {
	// TODO impl CopyTextureToBuffer
}

func (ce *CommandEncoder) CopyTextureToTexture(source *wgpu.ImageCopyTexture, destination *wgpu.ImageCopyTexture, copySize *wgpu.Extent3D) {
	// TODO impl CopyTextureToTexture
}

func (ce *CommandEncoder) Finish(descriptor *wgpu.CommandBufferDescriptor) wgpu.CommandBuffer {
	// TODO impl Finish
	return nil
}

func (ce *CommandEncoder) InsertDebugMarker(markerLabel string) {
	// TODO impl InsertDebugMarker
}

func (ce *CommandEncoder) PopDebugGroup() {
	// TODO impl PopDebugGroup
}

func (ce *CommandEncoder) PushDebugGroup(groupLabel string) {
	// TODO impl PushDebugGroup
}
