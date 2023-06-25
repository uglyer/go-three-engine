//go:build wasm

package wgpu_wasm

import "C"
import (
	"errors"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

type Device struct {
	ref   js.Value
	queue wgpu.IQueue
}

func newDevice(deviceRef js.Value) (wgpu.IDevice, error) {
	queue := newQueue(deviceRef.Get("queue"))
	return &Device{ref: deviceRef, queue: queue}, nil
}

func (d *Device) Drop() {
	// TODO impl GetQueue
	return
}

func (d *Device) GetQueue() wgpu.IQueue {
	return d.queue
}

func (d *Device) CreateBindGroup(descriptor *wgpu.GPUBindGroupDescriptor) (wgpu.IGPUBindGroup, error) {
	// TODO impl CreateBindGroup
	return nil, errors.New("todo impl CreateBindGroup")
}

func (d *Device) CreateBindGroupLayout(descriptor *wgpu.GPUBindGroupDescriptor) (wgpu.IGPUBindGroupLayout, error) {
	// TODO impl CreateBindGroupLayout
	return nil, errors.New("todo impl CreateBindGroupLayout")
}

func (d *Device) CreateBuffer(descriptor *wgpu.BufferDescriptor) (wgpu.IBuffer, error) {
	// TODO impl CreateBuffer
	return nil, errors.New("todo impl CreateBuffer")
}

func (d *Device) CreateCommandEncoder(descriptor *wgpu.CommandEncoderDescriptor) (wgpu.ICommandEncoder, error) {
	// TODO impl CreateCommandEncoder
	return nil, errors.New("todo impl CreateCommandEncoder")
}
func (d *Device) CreateShaderModule(descriptor *wgpu.ShaderModuleDescriptor) (wgpu.IShaderModule, error) {
	// TODO impl CreateCommandEncoder
	//return nil, errors.New("todo impl CreateShaderModule")
	desc := make(map[string]any)
	if descriptor != nil {
		if descriptor.Label != "" {
			desc["label"] = descriptor.Label
		}
		switch {
		case descriptor.WGSLDescriptor != nil:
			desc["code"] = descriptor.WGSLDescriptor.Code
			break
		case descriptor.SPIRVDescriptor != nil:
		case descriptor.GLSLDescriptor != nil:
			panic("暂不支持的类型")
		}
	}
	result := d.ref.Call("createShaderModule", desc)
	return &ShaderModule{ref: result}, nil
}

func (d *Device) CreateTexture(descriptor *wgpu.TextureDescriptor) (wgpu.ITexture, error) {
	// TODO impl CreateTexture
	return nil, errors.New("todo impl CreateTexture")
}

func (d *Device) CreateRenderPipeline(descriptor *wgpu.RenderPipelineDescriptor) (wgpu.IRenderPipeLine, error) {
	desc := make(map[string]any)
	if descriptor != nil {
		if descriptor.Label != "" {
			desc["label"] = descriptor.Label
		}

		if descriptor.Layout != nil {
			desc["layout"] = descriptor.Layout.(*PipelineLayout).ref
		} else {
			desc["layout"] = "auto"
		}

		// vertex
		{
			vertex := descriptor.Vertex

			var vert = make(map[string]any)

			if vertex.Module != nil {
				vert["module"] = vertex.Module.(*ShaderModule).ref
			}

			if vertex.EntryPoint != "" {
				vert["entryPoint"] = vertex.EntryPoint
			}

			bufferCount := len(vertex.Buffers)
			if bufferCount > 0 {
				// TODO 实现 buffer 处理
			}

			desc["vertex"] = vert
		}
		primitive := map[string]any{
			"topology":  descriptor.Primitive.Topology.String(),
			"frontFace": descriptor.Primitive.FrontFace.String(),
			"cullMode":  descriptor.Primitive.CullMode.String(),
		}
		if descriptor.Primitive.StripIndexFormat != wgpu.IndexFormat_Undefined {
			primitive["stripIndexFormat"] = descriptor.Primitive.StripIndexFormat.String()
		}

		desc["primitive"] = primitive

		if descriptor.DepthStencil != nil {
			depthStencil := descriptor.DepthStencil

			ds := make(map[string]any)

			//ds["nextInChain"] = nil
			ds["format"] = depthStencil.Format.String()
			ds["depthWriteEnabled"] = depthStencil.DepthWriteEnabled
			ds["depthCompare"] = depthStencil.DepthCompare.String()
			ds["stencilFront"] = map[string]any{
				"compare":     depthStencil.StencilFront.Compare.String(),
				"failOp":      depthStencil.StencilFront.FailOp.String(),
				"depthFailOp": depthStencil.StencilFront.DepthFailOp.String(),
				"passOp":      depthStencil.StencilFront.PassOp.String(),
			}
			ds["stencilBack"] = map[string]any{
				"compare":     depthStencil.StencilBack.Compare.String(),
				"failOp":      depthStencil.StencilBack.FailOp.String(),
				"depthFailOp": depthStencil.StencilBack.DepthFailOp.String(),
				"passOp":      depthStencil.StencilBack.PassOp.String(),
			}
			ds["stencilReadMask"] = depthStencil.StencilReadMask
			ds["stencilWriteMask"] = depthStencil.StencilWriteMask
			ds["depthBias"] = depthStencil.DepthBias
			ds["depthBiasSlopeScale"] = depthStencil.DepthBiasSlopeScale
			ds["depthBiasClamp"] = depthStencil.DepthBiasClamp

			desc["depthStencil"] = ds
		}

		desc["multisample"] = map[string]any{
			"count":                  descriptor.Multisample.Count,
			"mask":                   descriptor.Multisample.Mask,
			"alphaToCoverageEnabled": descriptor.Multisample.AlphaToCoverageEnabled,
		}

		if descriptor.Fragment != nil {
			fragment := descriptor.Fragment

			frag := make(map[string]any)

			frag["nextInChain"] = nil
			if fragment.EntryPoint != "" {
				frag["entryPoint"] = fragment.EntryPoint
			}

			if fragment.Module != nil {
				frag["module"] = fragment.Module.(*ShaderModule).ref
			}

			targetCount := len(fragment.Targets)
			if targetCount > 0 {
				targets := make([]any, targetCount)

				for i, v := range fragment.Targets {
					target := map[string]any{
						"format":    v.Format.String(),
						"writeMask": uint32(v.WriteMask),
					}
					if v.Blend != nil {
						target["blend"] = map[string]any{
							"color": map[string]any{
								"operation": v.Blend.Color.Operation.String(),
								"srcFactor": v.Blend.Color.SrcFactor.String(),
								"dstFactor": v.Blend.Color.DstFactor.String(),
							},
							"alpha": map[string]any{
								"operation": v.Blend.Alpha.Operation.String(),
								"srcFactor": v.Blend.Alpha.SrcFactor.String(),
								"dstFactor": v.Blend.Alpha.DstFactor.String(),
							},
						}
					}
					targets[i] = target
				}

				frag["targetCount"] = targetCount
				frag["targets"] = targets
			} else {
				frag["targetCount"] = 0
				frag["targets"] = nil
			}

			desc["fragment"] = frag
		}
	}
	obj := wasm.NewObj(desc)
	wasm.ConsoleLog("CreateRenderPipeline obj", obj)
	result := d.ref.Call("createRenderPipeline", obj)
	wasm.ConsoleLog("CreateRenderPipeline result", result)
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
