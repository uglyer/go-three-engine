package materials

import "github.com/uglyer/go-three-engine/engine/core"

// Side represents the material's visible side(s).
type Side int

// The face side(s) to be rendered. The non-rendered side will be culled to improve performance.
const (
	SideFront = Side(iota)
	SideBack
	SideDouble
)

// Blending specifies the blending mode.
type Blending int

// The various blending types.
const (
	BlendNone = Blending(iota)
	BlendNormal
	BlendAdditive
	BlendSubtractive
	BlendMultiply
	BlendCustom
)

// UseLights is a bitmask that specifies which types of lights affect the material.
type UseLights int

// The possible UseLights values.
const (
	UseLightNone        UseLights = 0x00
	UseLightAmbient     UseLights = 0x01
	UseLightDirectional UseLights = 0x02
	UseLightPoint       UseLights = 0x04
	UseLightSpot        UseLights = 0x08
	UseLightAll         UseLights = 0xFF
)

type Material struct {
	core.EventDispatcher
}
