package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

func ConvertOrigin3DToArray(point *wgpu.Origin3D) js.Value {
	return wasm.NewArray([]uint32{point.X, point.Y, point.Z})
}

func ConvertExtends3DToArray(point *wgpu.Extent3D) js.Value {
	return wasm.NewArray([]uint32{point.Width, point.Height, point.DepthOrArrayLayers})
}

func ConvertColorToArray(point *wgpu.Color) js.Value {
	return wasm.NewArray([]float64{point.R, point.G, point.B, point.A})
}

func ConvertImageCopyBufferToMap(source *wgpu.ImageCopyBuffer) map[string]any {
	sourceMap := map[string]any{
		"buffer": source.Buffer.(*Buffer).ref,
	}
	if source.Layout != nil {
		if &source.Layout.Offset != nil {
			sourceMap["offset"] = source.Layout.Offset
		}
		if &source.Layout.RowsPerImage != nil {
			sourceMap["rowsPerImage"] = source.Layout.RowsPerImage
		}
		if &source.Layout.BytesPerRow != nil {
			sourceMap["bytesPerRow"] = source.Layout.BytesPerRow
		}
	}
	return sourceMap
}

func ConvertImageImageCopyTexture(destination *wgpu.ImageCopyTexture) map[string]any {
	destMap := map[string]any{
		"texture": destination.Texture.(*Texture).ref,
	}
	if &destination.Aspect != nil {
		destMap["aspect"] = destination.Aspect.String()
	}
	if &destination.MipLevel != nil {
		destMap["mipLevel"] = destination.MipLevel
	}
	if destination.Origin != nil {
		destMap["origin"] = ConvertOrigin3DToArray(destination.Origin)
	}
	return destMap
}
