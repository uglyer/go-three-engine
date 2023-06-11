//go:build wasm

package wgpu_wasm

import (
	"errors"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type Device struct {
	deviceRef js.Value
}

func newDevice(deviceRef js.Value) (wgpu.IDevice, error) {
	return &Device{deviceRef: deviceRef}, nil
}

func (d *Device) Drop() {
	// TODO impl GetQueue
	return
}

func (d *Device) GetQueue() wgpu.IQueue {
	// TODO impl GetQueue
	return nil
}

func (d *Device) CreateBindGroup(descriptor *wgpu.GPUBindGroupDescriptor) (wgpu.IGPUBindGroup, error) {
	// TODO impl CreateBindGroup
	return nil, errors.New("todo impl CreateBindGroup")
}

func (d *Device) CreateBindGroupLayout(descriptor *wgpu.GPUBindGroupDescriptor) (wgpu.IGPUBindGroupLayout, error) {
	// TODO impl CreateBindGroupLayout
	return nil, errors.New("todo impl CreateBindGroupLayout")
}

func (d *Device) CreateBuffer(descriptor *wgpu.GPUBufferDescriptor) (wgpu.IBuffer, error) {
	// TODO impl CreateBuffer
	return nil, errors.New("todo impl CreateBuffer")
}

func (d *Device) CreateCommandEncoder(descriptor *wgpu.GPUCommandEncoderDescriptor) (wgpu.ICommandEncoder, error) {
	// TODO impl CreateCommandEncoder
	return nil, errors.New("todo impl CreateCommandEncoder")
}

func (d *Device) CreateTexture() (wgpu.ITexture, error) {
	// TODO impl CreateTexture
	return nil, errors.New("todo impl CreateTexture")
}

func (d *Device) CreateRenderPipeline() (wgpu.IRenderPipeLine, error) {
	// TODO impl CreateRenderPipeline
	return nil, errors.New("todo impl CreateRenderPipeline")
}

func (d *Device) GetErr() (err error) {
	// TODO impl getErr
	return nil
}

func (d *Device) StoreErr(typ wgpu.ErrorType, message string) {
	// TODO impl storeErr
}

func (d *Device) CreateComputePipeline(descriptor *wgpu.ComputePipelineDescriptor) (wgpu.IComputePipeline, error) {
	// TODO impl CreateComputePipeline
	return nil, errors.New("todo impl CreateComputePipeline")
}

func (d *Device) CreatePipelineLayout(descriptor *wgpu.PipelineLayoutDescriptor) (wgpu.IPipelineLayout, error) {
	// TODO impl CreatePipelineLayout
	return nil, errors.New("todo impl CreatePipelineLayout")
}

func (d *Device) CreateRenderBundleEncoder(descriptor *wgpu.RenderBundleEncoderDescriptor) (*wgpu.IRenderBundleEncoder, error) {
	// TODO impl CreateRenderBundleEncoder
	return nil, errors.New("todo impl CreateRenderBundleEncoder")
}

func (d *Device) Poll(wait bool, wrappedSubmissionIndex *wgpu.WrappedSubmissionIndex) (queueEmpty bool) {
	// TODO impl Poll
	return false
}
