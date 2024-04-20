//go:build wasm

package wgpu_wasm

import "C"
import (
	"errors"
	"fmt"
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

func (d *Device) Release() {
	// TODO impl GetQueue
	return
}

func (d *Device) GetQueue() wgpu.IQueue {
	return d.queue
}

func (d *Device) CreateBindGroup(descriptor *wgpu.GPUBindGroupDescriptor) (wgpu.IGPUBindGroup, error) {
	layout, ok := descriptor.Layout.(*BindGroupLayout)
	if !ok {
		return nil, fmt.Errorf("layout type error")
	}
	des := map[string]any{
		"layout": layout.ref,
	}
	if descriptor.Label != "" {
		des["label"] = descriptor.Label
	}
	targetCount := len(descriptor.Entries)
	if targetCount > 0 {
		entries := make([]any, targetCount)
		for i, it := range descriptor.Entries {
			data := map[string]any{
				"binding": it.Binding,
			}
			if it.Size != 0 {
				data["size"] = it.Size
			}
			if it.Offset != 0 {
				data["offset"] = it.Offset
			}
			if it.ResourceBuffer != nil {
				data["resource"] = it.ResourceBuffer.(*Buffer).ref
			} else if it.ResourceSampler != nil {
				data["resource"] = it.ResourceSampler.(*Sampler).ref
			} else if it.ResourceTextureView != nil {
				data["resource"] = it.ResourceTextureView.(*Sampler).ref
			} else {
				return nil, fmt.Errorf("resource is nil")
			}
			entries[i] = data
		}
	}
	ref := d.ref.Call("createBindGroup", des)
	return &BindGroupLayout{ref: ref}, nil
}

func (d *Device) CreateBindGroupLayout(descriptor *wgpu.BindGroupLayoutDescriptor) (wgpu.IGPUBindGroupLayout, error) {
	des := map[string]any{}
	if descriptor.Label != "" {
		des["label"] = descriptor.Label
	}
	targetCount := len(descriptor.Entries)
	if targetCount > 0 {
		entries := make([]any, targetCount)
		for i, it := range descriptor.Entries {
			data := map[string]any{
				"binding":    it.Binding,
				"visibility": int(it.Visibility),
			}
			if it.Buffer != nil {
				data["buffer"] = map[string]any{
					"type":             it.Buffer.Type.String(),
					"hasDynamicOffset": it.Buffer.HasDynamicOffset,
					"minBindingSize":   it.Buffer.MinBindingSize,
				}
			}
			if it.Texture != nil {
				data["texture"] = map[string]any{
					"multisampled":  it.Texture.Multisampled,
					"sampleType":    it.Texture.SampleType.String(),
					"viewDimension": it.Texture.ViewDimension.String(),
				}
			}
			if it.StorageTexture != nil {
				data["storageTexture"] = map[string]any{
					"access":        it.StorageTexture.Access.String(),
					"format":        it.StorageTexture.Format.String(),
					"viewDimension": it.StorageTexture.ViewDimension.String(),
				}
			}
			if it.Sampler != nil {
				data["sampler"] = map[string]any{
					"type": it.Sampler.Type.String(),
				}
			}
			entries[i] = data
		}
	}
	ref := d.ref.Call("createBindGroupLayout", des)
	return &BindGroupLayout{ref: ref}, nil
}

func (d *Device) CreateBuffer(descriptor *wgpu.BufferDescriptor) (wgpu.IBuffer, error) {
	desc := map[string]any{
		"label":            descriptor.Label,
		"mappedAtCreation": descriptor.MappedAtCreation,
		"size":             descriptor.Size,
		"usage":            int(descriptor.Usage),
	}
	ref := d.ref.Call("createBuffer", desc)
	return &Buffer{ref: ref}, nil
}

func (d *Device) CreateCommandEncoder(descriptor *wgpu.CommandEncoderDescriptor) (wgpu.ICommandEncoder, error) {
	desc := map[string]any{
		"label": descriptor.Label,
	}
	ref := d.ref.Call("createCommandEncoder", desc)
	return &CommandEncoder{ref: ref}, nil
}

func (d *Device) CreateShaderModule(descriptor *wgpu.ShaderModuleDescriptor) (wgpu.IShaderModule, error) {
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
	desc := map[string]any{
		"label":  descriptor.Label,
		"format": descriptor.Format.String(),
		"size": map[string]any{
			"width":              descriptor.Size.Width,
			"height":             descriptor.Size.Height,
			"depthOrArrayLayers": descriptor.Size.DepthOrArrayLayers,
		},
		"usage": int(descriptor.Usage),
	}
	if &descriptor.Dimension != nil {
		desc["dimension"] = descriptor.Dimension.String()
	}
	if descriptor.MipLevelCount != 0 {
		desc["mipLevelCount"] = descriptor.MipLevelCount
	}
	if descriptor.SampleCount != 0 {
		desc["sampleCount"] = descriptor.SampleCount
	}
	ref := d.ref.Call("createTexture", desc)
	return &Texture{ref: ref}, nil
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
				buffers := make([]any, bufferCount)
				for i, it := range vertex.Buffers {
					data := map[string]any{
						"arrayStride": it.ArrayStride,
					}
					attributes := make([]any, len(it.Attributes))
					for i2, attr := range it.Attributes {
						attributes[i2] = map[string]any{
							"format":         attr.Format.String(),
							"offset":         attr.Offset,
							"shaderLocation": attr.ShaderLocation,
						}
					}
					data["attributes"] = attributes
					if &it.StepMode != nil {
						data["stepMode"] = it.StepMode.String()
					}
					buffers[i] = data
				}
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

		if descriptor.Multisample != nil {
			desc["multisample"] = map[string]any{
				"count":                  descriptor.Multisample.Count,
				"mask":                   descriptor.Multisample.Mask,
				"alphaToCoverageEnabled": descriptor.Multisample.AlphaToCoverageEnabled,
			}
		}

		if descriptor.Fragment != nil {
			fragment := descriptor.Fragment

			frag := make(map[string]any)

			//frag["nextInChain"] = nil
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
						"format": v.Format.String(),
					}
					if v.WriteMask != 0 {
						target["writeMask"] = uint32(v.WriteMask)
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
	ref := d.ref.Call("createRenderPipeline", obj)
	js.Global().Set("pipeline", ref)
	return newRenderPipeline(ref), nil
}

func (d *Device) GetErr() (err error) {
	// TODO impl getErr
	return nil
}

func (d *Device) StoreErr(typ wgpu.ErrorType, message string) {
	// TODO impl storeErr
}

func (d *Device) CreateComputePipeline(descriptor *wgpu.ComputePipelineDescriptor) (wgpu.IComputePipeline, error) {
	desc := map[string]any{
		"label": descriptor.Label,
	}
	layout, ok := descriptor.Layout.(*PipelineLayout)
	if !ok {
		return nil, fmt.Errorf("layout type error")
	}
	desc["layout"] = layout
	if descriptor.Compute != nil {
		module, ok := descriptor.Compute.Module.(*ShaderModule)
		if !ok {
			return nil, fmt.Errorf("compute.module type error")
		}
		desc["compute"] = map[string]any{
			"module":     module.ref,
			"entryPoint": descriptor.Compute.EntryPoint,
		}
	}
	ref := d.ref.Call("createComputePipeline", desc)
	return newComputePipeline(ref), nil
}

func (d *Device) CreatePipelineLayout(descriptor *wgpu.PipelineLayoutDescriptor) (wgpu.IPipelineLayout, error) {
	desc := map[string]any{
		"label": descriptor.Label,
	}
	if len(descriptor.BindGroupLayouts) > 0 {
		bindGroupLayouts := make([]any, len(descriptor.BindGroupLayouts))
		for i, it := range descriptor.BindGroupLayouts {
			bindGroupLayouts[i] = it.(*BindGroupLayout).ref
		}
		desc["bindGroupLayouts"] = bindGroupLayouts
	}
	if descriptor.PushConstantRanges != nil {
		return nil, errors.New("PushConstantRanges is not support yet")
	}
	ref := d.ref.Call("createPipelineLayout", desc)
	return &PipelineLayout{ref: ref}, nil
}

func (d *Device) CreateRenderBundleEncoder(descriptor *wgpu.RenderBundleEncoderDescriptor) (wgpu.IRenderBundleEncoder, error) {
	desc := map[string]any{
		"label":           descriptor.Label,
		"depthReadOnly":   descriptor.DepthReadOnly,
		"stencilReadOnly": descriptor.StencilReadOnly,
	}
	if &descriptor.SampleCount != nil {
		desc["sampleCount"] = descriptor.SampleCount
	}
	if &descriptor.DepthStencilFormat != nil {
		desc["depthStencilFormat"] = descriptor.DepthStencilFormat.String()
	}
	colorFormatsCount := len(descriptor.ColorFormats)
	if colorFormatsCount > 0 {
		colorFormats := make([]any, colorFormatsCount)
		for i, it := range descriptor.ColorFormats {
			colorFormats[i] = it.String()
		}
		desc["colorFormats"] = colorFormats
	}
	ref := d.ref.Call("createPipelineLayout", desc)
	return newRenderBundleEncoder(ref), nil
}

func (d *Device) Poll(wait bool, wrappedSubmissionIndex *wgpu.WrappedSubmissionIndex) (queueEmpty bool) {
	// TODO impl Poll
	return false
}
