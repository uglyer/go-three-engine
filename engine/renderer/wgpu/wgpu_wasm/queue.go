//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/wasm"
	"syscall/js"
)

type Queue struct {
	ref js.Value
}

func newQueue(ref js.Value) wgpu.IQueue {
	return &Queue{ref: ref}
}

func (q *Queue) WriteTexture(destination *wgpu.ImageCopyTexture, data []byte, dataLayout *wgpu.TextureDataLayout, writeSize *wgpu.Extent3D) {
	dest := map[string]any{
		"texture": destination.Texture.(*Texture).ref,
	}
	if &destination.MipLevel != nil {
		dest["mipLevel"] = destination.MipLevel
	}
	if &destination.Aspect != nil {
		dest["aspect"] = destination.Aspect.String()
	}
	if destination.Origin != nil {
		dest["origin"] = ConvertOrigin3DToArray(destination.Origin)
	}
	dataJsValue := wasm.BytesToJsValue(data)
	dataLayoutObj := make(map[string]any)
	if &dataLayout.Offset != nil {
		dataLayoutObj["offset"] = dataLayout.Offset
	}
	if &dataLayout.BytesPerRow != nil {
		dataLayoutObj["bytesPerRow"] = dataLayout.BytesPerRow
	}
	if &dataLayout.RowsPerImage != nil {
		dataLayoutObj["rowsPerImage"] = dataLayout.RowsPerImage
	}
	size := ConvertExtends3DToArray(writeSize)
	q.ref.Call("writeTexture", dest, dataJsValue, dataLayoutObj, size)
}

func (q *Queue) WriteBuffer(buffer wgpu.IBuffer, bufferOffset uint64, data []byte) {
	bufferRef := buffer.(*Buffer).ref
	q.ref.Call("writeBuffer", bufferRef, bufferOffset, wasm.BytesToJsValue(data))
}

func (q *Queue) Submit(commands ...wgpu.ICommandBuffer) (submissionIndex wgpu.SubmissionIndex) {
	// TODO impl Submit
	return 0
}
