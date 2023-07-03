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
	sourceMap := map[string]any{
		"buffer": source.Buffer.(*Buffer).ref,
	}
	if source.Layout != nil {
		if &source.Layout.Offset != nil {
			sourceMap["offset"] = source.Layout.Offset
		}
		if &source.Layout.RowsPerImage != nil {
			sourceMap["rowsPerImage"] = source.Layout.RowsPerImage
		}
		if &source.Layout.BytesPerRow != nil {
			sourceMap["bytesPerRow"] = source.Layout.BytesPerRow
		}
	}
	destMap := map[string]any{
		"texture": destination.Texture.(*Texture).ref,
	}
	if &destination.Aspect != nil {
		destMap["aspect"] = destination.Aspect.String()
	}
	if &destination.MipLevel != nil {
		destMap["mipLevel"] = destination.MipLevel
	}
	if destination.Origin != nil {
		destMap["origin"] = ConvertOrigin3DToArray(destination.Origin)
	}
	ce.ref.Call("copyBufferToTexture", sourceMap, destMap, ConvertExtends3DToArray(copySize))
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
