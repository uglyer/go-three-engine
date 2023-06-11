//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
)

type ComputePassEncoder struct {
	// fields for ComputePassEncoder struct
}

func (ce *ComputePassEncoder) DispatchWorkgroups(workgroupCountX, workgroupCountY, workgroupCountZ uint32) {
	// TODO impl DispatchWorkgroups
}

func (ce *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *wgpu.IBuffer, indirectOffset uint64) {
	// TODO impl DispatchWorkgroupsIndirect
}

func (ce *ComputePassEncoder) End() {
	// TODO impl End
}

func (ce *ComputePassEncoder) InsertDebugMarker(markerLabel string) {
	// TODO impl InsertDebugMarker
}

func (ce *ComputePassEncoder) PopDebugGroup() {
	// TODO impl PopDebugGroup
}

func (ce *ComputePassEncoder) PushDebugGroup(groupLabel string) {
	// TODO impl PushDebugGroup
}

func (ce *ComputePassEncoder) SetBindGroup(groupIndex uint32, group wgpu.IGPUBindGroup, dynamicOffsets []uint32) {
	// TODO impl SetBindGroup
}

func (ce *ComputePassEncoder) SetPipeline(pipeline *wgpu.IComputePipeline) {
	// TODO impl SetPipeline
}