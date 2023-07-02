package wgpu_wasm

import "github.com/uglyer/go-three-engine/engine/renderer/wgpu"

func ConvertOrigin3DToArray(point *wgpu.Origin3D) []uint32 {
	return []uint32{point.X, point.Y, point.Z}
}

func ConvertExtends3DToArray(point *wgpu.Extent3D) []uint32 {
	return []uint32{point.Width, point.Height, point.DepthOrArrayLayers}
}
