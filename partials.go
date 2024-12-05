package partials

import (
	"fmt"
	"partials/components/blogsections"
	calltoactionsections "partials/components/callToActionSections"
	"reflect"
	"text/template"
	"unicode"
)

var (
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

func GetFuncMap() template.FuncMap {
	funcMap := template.FuncMap{}

	ctaFuncs := calltoactionsections.GetFuncMap()
	m := make(map[string]reflect.Value)
	addValueFuncs(m, ctaFuncs)
	addFuncs(funcMap, ctaFuncs)

	blogFunc := blogsections.GetFuncMap()
	addValueFuncs(m, blogFunc)
	addFuncs(funcMap, blogFunc)

	return funcMap
}

// addValueFuncs adds to values the functions in funcs, converting them to reflect.Values.
func addValueFuncs(out map[string]reflect.Value, in template.FuncMap) {
	for name, fn := range in {
		if !goodName(name) {
			panic(fmt.Errorf("function name %q is not a valid identifier", name))
		}
		v := reflect.ValueOf(fn)
		if v.Kind() != reflect.Func {
			panic("value for " + name + " not a function")
		}
		if !goodFunc(v.Type()) {
			panic(fmt.Errorf("can't install method/function %q with %d results", name, v.Type().NumOut()))
		}
		out[name] = v
	}
}

// addFuncs adds to values the functions in funcs. It does no checking of the input -
// call addValueFuncs first.
func addFuncs(out, in template.FuncMap) {
	for name, fn := range in {
		out[name] = fn
	}
}

// goodFunc reports whether the function or method has the right result signature.
func goodFunc(typ reflect.Type) bool {
	// We allow functions with 1 result or 2 results where the second is an error.
	switch {
	case typ.NumOut() == 1:
		return true
	case typ.NumOut() == 2 && typ.Out(1) == errorType:
		return true
	}
	return false
}

// goodName reports whether the function name is a valid identifier.
func goodName(name string) bool {
	if name == "" {
		return false
	}
	for i, r := range name {
		switch {
		case r == '_':
		case i == 0 && !unicode.IsLetter(r):
			return false
		case !unicode.IsLetter(r) && !unicode.IsDigit(r):
			return false
		}
	}
	return true
}

//https://stackoverflow.com/questions/43153133/how-to-assign-a-package-to-a-variable-in-go
