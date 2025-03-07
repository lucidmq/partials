package partials

import (
	"fmt"
	"partials/components/blogSections"
	calltoactionsections "partials/components/callToActionSections"
	"partials/components/eccomerceSections"
	faqsections "partials/components/faqSections"
	"partials/components/featureSections"
	"partials/components/gridCardsSections"
	headerTitleSections "partials/components/headersTitlesSections"
	"partials/components/heroComponents"
	"partials/components/navigationComponents"
	"partials/components/otherComponents"
	"partials/components/profileComponents"
	testamonials "partials/components/testamonialSections"
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

	blogFunc := blogSections.GetFuncMap()
	addValueFuncs(m, blogFunc)
	addFuncs(funcMap, blogFunc)

	featureSections := featureSections.GetFuncMap()
	addValueFuncs(m, featureSections)
	addFuncs(funcMap, featureSections)

	herosComp := heroComponents.GetFuncMap()
	addValueFuncs(m, herosComp)
	addFuncs(funcMap, herosComp)

	profComp := profileComponents.GetFuncMap()
	addValueFuncs(m, profComp)
	addFuncs(funcMap, profComp)

	otherComps := otherComponents.GetFuncMap()
	addValueFuncs(m, otherComps)
	addFuncs(funcMap, otherComps)

	faqCops := faqsections.GetFuncMap()
	addValueFuncs(m, faqCops)
	addFuncs(funcMap, faqCops)

	navComps := navigationComponents.GetFuncMap()
	addValueFuncs(m, navComps)
	addFuncs(funcMap, navComps)

	headersTitlesSectionsComps := headerTitleSections.GetFuncMap()
	addValueFuncs(m, headersTitlesSectionsComps)
	addFuncs(funcMap, headersTitlesSectionsComps)

	eccomerceSectionsCompos := eccomerceSections.GetFuncMap()
	addValueFuncs(m, eccomerceSectionsCompos)
	addFuncs(funcMap, eccomerceSectionsCompos)

	gridCardsSectionsComps := gridCardsSections.GetFuncMap()
	addValueFuncs(m, gridCardsSectionsComps)
	addFuncs(funcMap, gridCardsSectionsComps)

	testamonialSectionsComps := testamonials.GetFuncMap()
	addValueFuncs(m, testamonialSectionsComps)
	addFuncs(funcMap, testamonialSectionsComps)

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
