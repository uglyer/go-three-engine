package wasm

import (
	"syscall/js"
)

func NewObject() js.Value {
	return js.Global().Get("Object").New()
}

func ConsoleLog(v ...any) {
	js.Global().Get("console").Call("log", v...)
}

func IsUndefined(v js.Value) bool {
	return v.Type() == js.TypeUndefined
}

func Await(target js.Value) js.Value {
	wait := make(chan js.Value)
	defer close(wait)
	target.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		wait <- args[0]
		return nil
	}))
	return <-wait
}
