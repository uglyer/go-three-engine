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
	// TODO impl SetBindGroup
}

func (rpe *RenderPassEncoder) SetBlendConstant(color *wgpu.Color) {
	// TODO impl SetBlendConstant
}

func (rpe *RenderPassEncoder) SetIndexBuffer(buffer wgpu.IBuffer, format wgpu.IndexFormat, offset, size uint64) {
	// TODO impl SetIndexBuffer
}

func (rpe *RenderPassEncoder) SetPipeline(pipeline wgpu.IRenderPipeLine) {
	// TODO impl SetPipeline
}

func (rpe *RenderPassEncoder) SetScissorRect(x, y, width, height uint32) {
	// TODO impl SetScissorRect
}

func (rpe *RenderPassEncoder) SetStencilReference(reference uint32) {
	// TODO impl SetStencilReference
}

func (rpe *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer wgpu.IBuffer, offset, size uint64) {
	// TODO impl SetVertexBuffer
}

func (rpe *RenderPassEncoder) SetViewport(x, y, width, height, minDepth, maxDepth float32) {
	// TODO impl SetViewport
}

func (rpe *RenderPassEncoder) SetPushConstants(stages wgpu.ShaderStage, offset uint32, data []byte) {
	// TODO impl SetPushConstants
}

func (rpe *RenderPassEncoder) MultiDrawIndirect(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, count uint32) {
	// TODO impl MultiDrawIndirect
}

func (rpe *RenderPassEncoder) MultiDrawIndexedIndirect(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, count uint32) {
	// TODO impl MultiDrawIndexedIndirect
}

func (rpe *RenderPassEncoder) MultiDrawIndirectCount(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, countBuffer wgpu.IBuffer, countBufferOffset uint64, maxCount uint32) {
	// TODO impl MultiDrawIndirectCount
}

func (rpe *RenderPassEncoder) MultiDrawIndexedIndirectCount(encoder wgpu.IRenderPassEncoder, buffer wgpu.IBuffer, offset uint64, countBuffer wgpu.IBuffer, countBufferOffset uint64, maxCount uint32) {
	// TODO impl MultiDrawIndexedIndirectCount
}

type RenderPipeline struct {
}

func newRenderPipeline() wgpu.IRenderPipeLine {
	return &RenderPipeline{}
}

func (r RenderPipeline) Drop() {
	//TODO implement me
	panic("implement me")
}

func (r RenderPipeline) GetBindGroupLayout(groupIndex uint32) wgpu.IGPUBindGroupLayout {
	//TODO implement me
	panic("implement me")
}
