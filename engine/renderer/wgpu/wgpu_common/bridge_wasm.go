//go:build wasm

package wgpu_common

import (
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu"
	"github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_wasm"
)

func NewBridge() (wgpu.IBridge, error) {
	return wgpu_wasm.NewBridge()
}
