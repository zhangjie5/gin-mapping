package mapping

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
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
			if mName == "Group" {
				continue
			}
			mVal := val.Method(i)
			handle := gin.HandlerFunc(mVal.Interface().(func(ctx *gin.Context)))

			for _, e := range httpMethod {
				if strings.HasPrefix(mName, e) {
					mapping := strings.TrimPrefix(mName, e)
					g.Handle(strings.ToUpper(e), mapping, handle)
				}
			}
		}
	} else {
		panic(val.String() + " is not implements GinApi interface")
	}
}