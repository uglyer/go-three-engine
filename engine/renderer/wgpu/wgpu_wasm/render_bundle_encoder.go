package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type RenderBundleEncoder struct {
	ref js.Value
}

func newRenderBundleEncoder(ref js.Value) wgpu.IRenderBundleEncoder {
	return &RenderBundleEncoder{ref: ref}
}

func (rbe *RenderBundleEncoder) Release() {
	//TODO implement me
	panic("implement me")
}

func (rbe *RenderBundleEncoder) Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	// TODO impl Draw
}

func (rbe *RenderBundleEncoder) DrawIndexed(indexCount, instanceCount, firstIndex, baseVertex, firstInstance uint32) {
	// TODO impl DrawIndexed
}

func (rbe *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer wgpu.IBuffer, indirectOffset uint64) {
	// TODO impl DrawIndexedIndirect
}

func (rbe *RenderBundleEncoder) DrawIndirect(indirectBuffer wgpu.IBuffer, indirectOffset uint64) {
	// TODO impl DrawIndirect
}

func (rbe *RenderBundleEncoder) Finish(descriptor *wgpu.RenderBundleDescriptor) wgpu.IRenderBundle {
	// TODO impl Finish
	return nil
}

func (rbe *RenderBundleEncoder) InsertDebugMarker(markerLabel string) {
	// TODO impl InsertDebugMarker
}

func (rbe *RenderBundleEncoder) PopDebugGroup() {
	// TODO impl PopDebugGroup
}

func (rbe *RenderBundleEncoder) PushDebugGroup(groupLabel string) {
	// TODO impl PushDebugGroup
}

func (rbe *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group wgpu.IGPUBindGroup, dynamicOffsets []uint32) {
	// TODO impl SetBindGroup
}

func (rbe *RenderBundleEncoder) SetIndexBuffer(buffer wgpu.IBuffer, format wgpu.IndexFormat, offset, size uint64) {
	// TODO impl SetIndexBuffer
}

func (rbe *RenderBundleEncoder) SetPipeline(pipeline wgpu.IRenderPipeLine) {
	// TODO impl SetPipeline
}

func (rbe *RenderBundleEncoder) SetVertexBuffer(slot uint32, buffer wgpu.IBuffer, offset, size uint64) {
	// TODO impl SetVertexBuffer
}
