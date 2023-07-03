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
	colorAttachments := make([]any, len(descriptor.ColorAttachments))
	for i, it := range descriptor.ColorAttachments {
		item := map[string]any{
			"clearValue": ConvertColorToArray(it.ClearValue),
			"loadOp":     it.LoadOp.String(),
			"storeOp":    it.StoreOp.String(),
			"view":       it.View.(*TextureView).ref,
		}
		if it.ResolveTarget != nil {
			item["resolveTarget"] = it.ResolveTarget.(*TextureView).ref
		}
		colorAttachments[i] = item
	}
	desc := map[string]any{
		"label":            descriptor.Label,
		"colorAttachments": colorAttachments,
	}
	if descriptor.DepthStencilAttachment != nil {
		src := descriptor.DepthStencilAttachment
		data := map[string]any{
			"view": src.View.(*TextureView).ref,
		}
		if &src.DepthClearValue != nil {
			data["depthClearValue"] = src.DepthClearValue
		}
		if &src.DepthLoadOp != nil {
			data["depthLoadOp"] = src.DepthLoadOp.String()
		}
		if &src.DepthReadOnly != nil {
			data["depthReadOnly"] = src.DepthReadOnly
		}
		if &src.DepthStoreOp != nil {
			data["depthStoreOp"] = src.DepthStoreOp.String()
		}
		if &src.StencilClearValue != nil {
			data["stencilClearValue"] = src.StencilClearValue
		}
		if &src.StencilLoadOp != nil {
			data["stencilLoadOp"] = src.StencilLoadOp.String()
		}
		if &src.StencilReadOnly != nil {
			data["stencilReadOnly"] = src.StencilReadOnly
		}
		if &src.StencilStoreOp != nil {
			data["stencilStoreOp"] = src.StencilStoreOp.String()
		}
		desc["depthStencilAttachment"] = data
	}
	ref := ce.ref.Call("beginRenderPass", desc)
	return newRenderPassEncoder(ref)
}

func (ce *CommandEncoder) ClearBuffer(buffer wgpu.IBuffer, offset uint64, size uint64) {
	if &offset != nil && &size != nil {
		ce.ref.Call("clearBuffer", buffer.(*Buffer).ref, offset, size)
	} else if &offset != nil {
		ce.ref.Call("clearBuffer", buffer.(*Buffer).ref, offset)
	} else {
		ce.ref.Call("clearBuffer", buffer.(*Buffer).ref)
	}
}

func (ce *CommandEncoder) CopyBufferToBuffer(source wgpu.IBuffer, sourceOffset uint64, destination wgpu.IBuffer, destinatonOffset uint64, size uint64) {
	ce.ref.Call("copyBufferToBuffer", source.(*Buffer).ref, sourceOffset, destination.(*Buffer).ref, destinatonOffset, size)
}

func (ce *CommandEncoder) CopyBufferToTexture(source *wgpu.ImageCopyBuffer, destination *wgpu.ImageCopyTexture, copySize *wgpu.Extent3D) {
	ce.ref.Call("copyBufferToTexture", ConvertImageCopyBufferToMap(source), ConvertImageImageCopyTexture(destination), ConvertExtends3DToArray(copySize))
}

func (ce *CommandEncoder) CopyTextureToBuffer(source *wgpu.ImageCopyTexture, destination *wgpu.ImageCopyBuffer, copySize *wgpu.Extent3D) {
	ce.ref.Call("copyTextureToBuffer", ConvertImageImageCopyTexture(source), ConvertImageCopyBufferToMap(destination), ConvertExtends3DToArray(copySize))
}

func (ce *CommandEncoder) CopyTextureToTexture(source *wgpu.ImageCopyTexture, destination *wgpu.ImageCopyTexture, copySize *wgpu.Extent3D) {
	ce.ref.Call("copyTextureToTexture", ConvertImageImageCopyTexture(source), ConvertImageImageCopyTexture(destination), ConvertExtends3DToArray(copySize))
}

func (ce *CommandEncoder) Finish(descriptor *wgpu.CommandBufferDescriptor) wgpu.ICommandBuffer {
	if descriptor != nil {
		desc := map[string]any{
			"label": descriptor.Label,
		}
		ref := ce.ref.Call("finish", desc)
		return &CommandBuffer{ref: ref}
	}
	ref := ce.ref.Call("finish")
	return &CommandBuffer{ref: ref}
}

func (ce *CommandEncoder) InsertDebugMarker(markerLabel string) {
	ce.ref.Call("insertDebugMarker", markerLabel)
}

func (ce *CommandEncoder) PopDebugGroup() {
	// TODO impl PopDebugGroup
}

func (ce *CommandEncoder) PushDebugGroup(groupLabel string) {
	// TODO impl PushDebugGroup
}
