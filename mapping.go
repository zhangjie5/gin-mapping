package mapping

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"unicode"
)

// GinHandle represents every handle need implements this interface
type GinHandle interface {

	// Mapping will use this method return value to group apis
	Group() string
}

// Use start with method name to define which http method should use
var httpMethod = []string{"Get", "Post", "Put", "Delete", "Options", "Head"}

// Mapping router with this struct methods except Group method.
func Register(e *gin.Engine, v interface{}) {
	val := reflect.ValueOf(v)
	typ := val.Type()
	if typ.Implements(reflect.TypeOf((*GinHandle)(nil)).Elem()) {
		groupMV := val.MethodByName("Group")
		group := groupMV.Call(nil)[0].Interface().(string)
		g := e.Group(group)
		for i := 0; i < val.NumMethod(); i++ {
			mTyp := typ.Method(i)
			mName := mTyp.Name
			mVal := val.Method(i)

			// will discard method if it first parameter is not *gin.Context
			if !isHandle(mVal) {
				continue
			}

			handle := gin.HandlerFunc(mVal.Interface().(func(ctx *gin.Context)))

			for _, e := range httpMethod {
				if strings.HasPrefix(mName, e) {

					// make first character lower
					mapping := formatMapping(strings.TrimPrefix(mName, e))
					g.Handle(strings.ToUpper(e), mapping, handle)
				}
			}
		}
	} else {
		panic(val.String() + " is not implements GinApi interface")
	}
}

func isHandle(v reflect.Value) bool {
	typ := v.Type()
	return typ.NumIn() > 0 && typ.In(0) == reflect.TypeOf(&gin.Context{})
}

func formatMapping(s string) string {
	if len(s) == 0 {
		return s
	} else if len(s) == 1 {
		return strings.ToLower(s)
	} else {
		arr := []rune(s)
		return string(unicode.ToLower(arr[0])) + string(arr[1:])
	}
}
