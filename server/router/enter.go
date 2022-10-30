package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/encrypted_middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System               system.RouterGroup
	Example              example.RouterGroup
	Encrypted_middleware encrypted_middleware.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
