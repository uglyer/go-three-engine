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
	refCount int // Current number of references

	// Shader specification
	shader       string // Shader name
	shaderUnique bool   // shader has only one instance (does not depend on lights or textures)
	//ShaderDefines gls.ShaderDefines // shader defines

	side        Side      // Face side(s) visibility
	blending    Blending  // Blending mode
	useLights   UseLights // Which light types to consider
	transparent bool      // Whether at all transparent
	wireframe   bool      // Whether to render only the wireframe
	lineWidth   float32   // Line width for lines and wireframe
	//textures    []*texture.Texture2D // List of textures

	polyOffsetFactor float32 // polygon offset factor
	polyOffsetUnits  float32 // polygon offset units

	depthMask bool   // Enable writing into the depth buffer
	depthTest bool   // Enable depth buffer test
	depthFunc uint32 // Active depth test function

	// TODO stencil properties

	// Equations used for custom blending (when blending=BlendCustom) // TODO implement methods
	blendRGB      uint32 // separate blending equation for RGB
	blendAlpha    uint32 // separate blending equation for Alpha
	blendSrcRGB   uint32 // separate blending func source RGB
	blendDstRGB   uint32 // separate blending func dest RGB
	blendSrcAlpha uint32 // separate blending func source Alpha
	blendDstAlpha uint32 // separate blending func dest Alpha
}
