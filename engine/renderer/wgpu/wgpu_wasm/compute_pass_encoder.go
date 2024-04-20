//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type ComputePassEncoder struct {
	// fields for ComputePassEncoder struct
	ref js.Value
}

func newComputePassEncoder(ref js.Value) *ComputePassEncoder {
	return &ComputePassEncoder{ref: ref}
}

func (ce *ComputePassEncoder) DispatchWorkgroups(workgroupCountX, workgroupCountY, workgroupCountZ uint32) {
	ce.ref.Call("dispatchWorkgroups", workgroupCountX, workgroupCountY, workgroupCountZ)
}

func (ce *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer wgpu.IBuffer, indirectOffset uint64) {
	ce.ref.Call("dispatchWorkgroupsIndirect", indirectBuffer.(*Buffer).ref, indirectOffset)
}

func (ce *ComputePassEncoder) End() {
	ce.ref.Call("end")
}

func (ce *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	ce.ref.Call("insertDebugMarker", markerLabel)
}

func (ce *ComputePassEncoder) PopDebugGroup() {
	ce.ref.Call("popDebugGroup")
}

func (ce *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	ce.ref.Call("pushDebugGroup", groupLabel)
}

func (ce *ComputePassEncoder) SetBindGroup(groupIndex uint32, group wgpu.IGPUBindGroup, dynamicOffsets []uint32) {
	ce.ref.Call("setBindGroup", groupIndex, group.(*BindGroup).ref, dynamicOffsets)
}

func (ce *ComputePassEncoder) SetPipeline(pipeline wgpu.IComputePipeline) {
	ce.ref.Call("pushDebugGroup", pipeline.(*ComputePipeline).ref)
}

type ComputePipeline struct {
	ref js.Value
}

func newComputePipeline(ref js.Value) wgpu.IComputePipeline {
	return &ComputePipeline{ref: ref}
}

func (c *ComputePipeline) Release() {
	//TODO implement me
	panic("implement me")
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) wgpu.IGPUBindGroupLayout {
	ref := c.ref.Call("getBindGroupLayout", groupIndex)
	return &BindGroupLayout{ref: ref}
}
