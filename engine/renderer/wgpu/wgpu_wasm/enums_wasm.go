//go:build wasm

package wgpu_wasm

import "github.com/uglyer/go-three-engine/engine/renderer/wgpu/wgpu_bridge"

var WasmPowerPreference = map[wgpu_bridge.PowerPreference]string{
	wgpu_bridge.PowerPreference_Default:         "high-performance",
	wgpu_bridge.PowerPreference_LowPower:        "low-power",
	wgpu_bridge.PowerPreference_HighPerformance: "high-performance",
}
