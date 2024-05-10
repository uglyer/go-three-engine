//go:build !wasm

package wgpu_native

import "github.com/uglyer/go-three-engine/engine/renderer/wgpu"

type BufferInitDescriptor struct {
	Label    string
	Contents []byte
	Usage    wgpu.BufferUsage
}

func (p *Device) CreateBufferInit(descriptor *BufferInitDescriptor) (*Buffer, error) {
	if descriptor == nil {
		panic("got nil descriptor")
	}

	if len(descriptor.Contents) == 0 {
		return p.CreateBuffer(&BufferDescriptor{
			Label:            descriptor.Label,
			Size:             0,
			Usage:            descriptor.Usage,
			MappedAtCreation: false,
		})
	}

	unpaddedSize := len(descriptor.Contents)
	const alignMask = CopyBufferAlignment - 1
	paddedSize := max(((unpaddedSize + alignMask) & ^alignMask), CopyBufferAlignment)

	buffer, err := p.CreateBuffer(&BufferDescriptor{
		Label:            descriptor.Label,
		Size:             uint64(paddedSize),
		Usage:            descriptor.Usage,
		MappedAtCreation: true,
	})
	if err != nil {
		return nil, err
	}
	buf := buffer.GetMappedRange(0, uint(paddedSize))
	copy(buf, descriptor.Contents)
	buffer.Unmap()

	return buffer, nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
