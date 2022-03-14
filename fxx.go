package fxx

import (
	"reflect"

	"go.uber.org/fx"
)

func MultiProvide(constructors ...interface{}) fx.Option {
	return fx.Provide(
		joinConstructor(constructors)...,
	)
}
func joinConstructor(modules ...interface{}) []interface{} {
	_modules := make([]interface{}, 0)
	for _, v := range modules {
		rt := reflect.TypeOf(v)
		// fmt.Println(v, rt.Kind().String())
		switch rt.Kind() {
		case reflect.Slice:

			_modules = append(_modules, joinConstructor(v.([]interface{})...)...)
		case reflect.Func:
			_modules = append(_modules, v)
		default:
			_modules = append(_modules, v)
		}
	}

	return _modules
}

func Modules(modules ...interface{}) []interface{} {
	return modules
}
