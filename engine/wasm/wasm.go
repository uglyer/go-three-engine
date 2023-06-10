//go:build wasm

package wasm

import (
	"fmt"
	"syscall/js"
)

func NewObject() js.Value {
	return js.Global().Get("Object").New()
}

func NewObj(m map[string]any) js.Value {
	return js.Global().Get("Object").New(m)
}

func ConsoleLog(v ...any) {
	js.Global().Get("console").Call("log", v...)
}
func ConsoleError(v ...any) {
	js.Global().Get("console").Call("error", v...)
}

func IsUndefined(v js.Value) bool {
	return v.IsNull() || v.IsUndefined()
}

func Await(target js.Value) (*js.Value, error) {
	wait := make(chan js.Value)
	catch := make(chan js.Value)
	defer close(wait)
	defer close(catch)
	target.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		wait <- args[0]
		return nil
	}))
	target.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		catch <- args[0]
		return nil
	}))
	select {
	case result := <-wait:
		return &result, nil
	case err := <-catch:
		return nil, fmt.Errorf("promise error:%s", err.Call("toString"))
	}
}
