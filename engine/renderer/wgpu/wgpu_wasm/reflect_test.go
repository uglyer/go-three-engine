package wgpu_wasm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestReflect(t *testing.T) {
	u := User{Name: "John", Age: 30}
	v := reflect.ValueOf(&u).Elem()
	typeOfLimit := reflect.TypeOf(u)
	size := typeOfLimit.NumField()
	for i := 0; i < size; i++ {
		name := typeOfLimit.Field(i).Name
		field := v.FieldByName(name)
		name = strings.ToLower(name[0:1]) + name[1:]

		field.SetUint(1)
		//field.SetUint(uint64(jsLimits.Get(name).Int()))
	}
	v.FieldByName("Name").SetString("Jane")
	v.FieldByName("Age").SetInt(35)
	fmt.Println(u)
}
