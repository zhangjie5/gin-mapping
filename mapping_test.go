package mapping

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

type A struct {}

func (a *A) Group() string {

	return "a"
}

func (a *A) GetHello(c *gin.Context) {

	c.String(200, "Hello, World")
}

func TestMapping(t *testing.T) {
	api := &A{}
	g := gin.Default()
	Register(g, api)
	for _, v := range g.Routes() {
		assert.Equal(t, v.Path, "/a/Hello")
	}
}
