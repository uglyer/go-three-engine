//go:build wasm

package wgpu_bridge

var WasmPowerPreference = map[PowerPreference]string{
	PowerPreference_Default:         "high-performance",
	PowerPreference_LowPower:        "low-power",
	PowerPreference_HighPerformance: "high-performance",
}
