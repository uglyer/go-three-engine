//go:build wasm

package wgpu_wasm

import (
	"fmt"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"log"
	"reflect"
	"strings"
	"syscall/js"
)

type Adapter struct {
	ref js.Value
}

func newAdapter(adapterRef js.Value) (wgpu.IAdapter, error) {
	return &Adapter{
		ref: adapterRef,
	}, nil
}

func (a *Adapter) Drop() {
	// TODO impl adapter drop
}

func (a *Adapter) RequestDevice(descriptor *wgpu.DeviceDescriptor) (wgpu.IDevice, error) {
	if descriptor != nil {
		log.Println("todo RequestDevice with descriptor")
	}
	device, err := wasm.Await(a.ref.Call("requestDevice"))
	if err != nil {
		return nil, fmt.Errorf("call requestDevice error:%v", err)
	}
	wasm.ConsoleLog("a.ref", a.ref)
	wasm.ConsoleLog("device", *device)
	return newDevice(*device)
}
func (a *Adapter) EnumerateFeatures() []wgpu.FeatureName {
	// TODO impl adapter EnumerateFeatures
	return nil
}
func (a *Adapter) GetLimits() wgpu.SupportedLimits {
	// TODO impl adapter EnumerateFeatures
	// TODO 使用反射实现类型映射转换
	//jsLimits := a.ref.Get("limits")
	limit := wgpu.Limits{}
	valueOfLimit := reflect.ValueOf(&limit).Elem()
	typeOfLimit := reflect.TypeOf(limit)
	wasm.ConsoleLog("valueOfLimit.Kind().String()", valueOfLimit.Kind().String())
	size := typeOfLimit.NumField()
	for i := 0; i < size; i++ {
		name := typeOfLimit.Field(i).Name
		field := valueOfLimit.FieldByName(name)
		name = strings.ToLower(name[0:1]) + name[1:]
		wasm.ConsoleLog("name", name)
		field.SetUint(1)
		//field.SetUint(uint64(jsLimits.Get(name).Int()))
	}
	//limits.Get("maxBindGroups")
	wasm.ConsoleLog("device.GetLimits", limit)
	return wgpu.SupportedLimits{
		Limits: &wgpu.Limits{
			//MaxTextureDimension1D:                     uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxTextureDimension2D:                     uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxTextureDimension3D:                     uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxTextureArrayLayers:                     uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxBindGroups:                             uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxBindingsPerBindGroup:                   uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxDynamicUniformBuffersPerPipelineLayout: uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxDynamicStorageBuffersPerPipelineLayout: uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxSampledTexturesPerShaderStage:          uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxSamplersPerShaderStage:                 uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxStorageBuffersPerShaderStage:           uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxStorageTexturesPerShaderStage:          uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxUniformBuffersPerShaderStage:           uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxUniformBufferBindingSize:               uint64(limits.Get("maxTextureDimension1D").Int()),
			//MaxStorageBufferBindingSize:               uint64(limits.Get("maxTextureDimension1D").Int()),
			//MinUniformBufferOffsetAlignment:           uint32(limits.Get("maxTextureDimension1D").Int()),
			//MinStorageBufferOffsetAlignment:           uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxVertexBuffers:                          uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxBufferSize:                             uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxVertexAttributes:                       uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxVertexBufferArrayStride:                uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxInterStageShaderComponents:             uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxInterStageShaderVariables:              uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxColorAttachments:                       uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxComputeWorkgroupStorageSize:            uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxComputeInvocationsPerWorkgroup:         uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxComputeWorkgroupSizeX:                  uint32(limits.Get("maxTextureDimension1D").Int()),
			//MaxComputeWorkgroupSizeY:                  0,
			//MaxComputeWorkgroupSizeZ:                  0,
			//MaxComputeWorkgroupsPerDimension:          0,
			//MaxPushConstantSize:                       0,
		},
	}
}
func (a *Adapter) GetProperties() wgpu.AdapterProperties {
	// TODO impl adapter EnumerateFeatures
	return wgpu.AdapterProperties{}
}
func (a *Adapter) HasFeature(feature wgpu.FeatureName) bool {
	// TODO impl adapter EnumerateFeatures
	return false
}
