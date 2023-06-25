//go:build wasm

package wgpu_wasm

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"syscall/js"
)

type Queue struct {
	ref js.Value
}

func newQueue(ref js.Value) wgpu.IQueue {
	return &Queue{ref: ref}
}

func (q *Queue) WriteTexture(destination *wgpu.ImageCopyTexture, data []byte, dataLayout *wgpu.TextureDataLayout, writeSize *wgpu.Extent3D) {
	// TODO impl WriteTexture
}

func (q *Queue) WriteBuffer(buffer wgpu.IBuffer, bufferOffset uint64, data []byte) {
	// TODO impl WriteBuffer
}

func (q *Queue) Submit(commands ...wgpu.ICommandBuffer) (submissionIndex wgpu.SubmissionIndex) {
	// TODO impl Submit
	return 0
}
