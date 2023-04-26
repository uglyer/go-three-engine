package wasm

import "syscall/js"

func NewObject() js.Value {
	return js.Global().Get("Object").New()
}

func ConsoleLog(v ...any) {
	js.Global().Get("console").Call("log", v...)
}

func IsUndefined(v js.Value) bool {
	return v.Type() == js.TypeUndefined
}
