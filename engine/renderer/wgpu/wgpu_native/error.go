//go:build !wasm

package wgpu_native

import "fmt"

type Error struct {
	Type    ErrorType
	Message string
}

func (v *Error) Error() string {
	return fmt.Sprintf("%s : %s", v.Type.String(), v.Message)
}
