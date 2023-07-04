//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type RenderPassEncoder struct {
	// fields for IRenderPassEncoder struct
	ref js.Value
}

func newRenderPassEncoder(ref js.Value) wgpu.IRenderPassEncoder {
	return &RenderPassEncoder{ref: ref}
}

func (rpe *RenderPassEncoder) Drop() {
	//TODO implement me
	panic("implement me")
}

func (rpe *RenderPassEncoder) Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	rpe.ref.Call("draw", vertexCount, instanceCount, firstVertex, firstInstance)
}

func (rpe *RenderPassEncoder) DrawIndexed(indexCount, instanceCount, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	rpe.ref.Call("drawIndexed", indexCount, instanceCount, firstIndex, baseVertex, firstInstance)
}

func (rpe *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer wgpu.IBuffer, indirectOffset uint64) {
	rpe.ref.Call("drawIndexedIndirect", indirectBuffer.(*Buffer).ref, indirectOffset)
}

func (rpe *RenderPassEncoder) DrawIndirect(indirectBuffer wgpu.IBuffer, indirectOffset uint64) {
	rpe.ref.Call("drawIndirect", indirectBuffer.(*Buffer).ref, indirectOffset)
}

func (rpe *RenderPassEncoder) End() {
	rpe.ref.Call("end")
}

func (rpe *RenderPassEncoder) ExecuteBundles(bundles ...wgpu.IRenderBundle) {
	list := make([]js.Value, len(bundles))
	for i, it := range bundles {
		list[i] = it.(*RenderBundle).ref
	}
}

func (rpe *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	rpe.ref.Call("insertDebugMarker", markerLabel)
}

func (rpe *RenderPassEncoder) PopDebugGroup() {
	rpe.ref.Call("popDebugGroup")
}

func (rpe *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	rpe.ref.Call("pushDebugGroup", groupLabel)
}

func (rpe *RenderPassEncoder) SetBindGroup(groupIndex uint32, group wgpu.IGPUBindGroup, dynamicOffsets []uint32) {
	rpe.ref.Call("setBindGroup", groupIndex, group.(*BindGroup).ref, dynamicOffsets)
}

func (rpe *RenderPassEncoder) SetBlendConstant(color *wgpu.Color) {
	rpe.ref.Call("setBlendConstant", ConvertColorToArray(color))
}

func (rpe *RenderPassEncoder) SetIndexBuffer(buffer wgpu.IBuffer, format wgpu.IndexFormat, offset, size uint64) {
	rpe.ref.Call("setBlendConstant", buffer.(*Buffer).ref, format.String(), offset, size)
}

func (rpe *RenderPassEncoder) SetPipeline(pipeline wgpu.IRenderPipeLine) {
	rpe.ref.Call("setPipeline", pipeline.(*RenderPipeline).ref)
}

func (rpe *RenderPassEncoder) SetScissorRect(x, y, width, height uint32) {
	rpe.ref.Call("setScissorRect", x, y, width, height)
}

func (rpe *RenderPassEncoder) SetStencilReference(reference uint32) {
	rpe.ref.Call("setStencilReference", reference)
}

func (rpe *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer wgpu.IBuffer, offset, size uint64) {
	rpe.ref.Call("setVertexBuffer", slot, buffer.(*Buffer).ref, offset, size)
}

func (rpe *RenderPassEncoder) SetViewport(x, y, width, height, minDepth, maxDepth float32) {
	rpe.ref.Call("setViewport", x, y, width, height, minDepth, maxDepth)
}

func (rpe *RenderPassEncoder) SetPushConstants(stages wgpu.ShaderStage, offset uint32, data []byte) {
	panic(wgpu.WASM_NOT_SUPPORT)
}

func (rpe *RenderPassEncoder) MultiDrawIndirect(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, count uint32) {
	panic(wgpu.WASM_NOT_SUPPORT)
}

func (rpe *RenderPassEncoder) MultiDrawIndexedIndirect(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, count uint32) {
	panic(wgpu.WASM_NOT_SUPPORT)
}

func (rpe *RenderPassEncoder) MultiDrawIndirectCount(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, countBuffer wgpu.IBuffer, countBufferOffset uint64, maxCount uint32) {
	panic(wgpu.WASM_NOT_SUPPORT)
}

func (rpe *RenderPassEncoder) MultiDrawIndexedIndirectCount(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, countBuffer wgpu.IBuffer, countBufferOffset uint64, maxCount uint32) {
	panic(wgpu.WASM_NOT_SUPPORT)
}

type RenderPipeline struct {
	ref js.Value
}

func newRenderPipeline(ref js.Value) wgpu.IRenderPipeLine {
	return &RenderPipeline{ref: ref}
}

func (r RenderPipeline) Drop() {
	//TODO implement me
	panic("implement me")
}

func (r RenderPipeline) GetBindGroupLayout(groupIndex uint32) wgpu.IGPUBindGroupLayout {
	ref := r.ref.Call("getBindGroupLayout", groupIndex)
	return &BindGroupLayout{ref: ref}
}
