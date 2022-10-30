package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/encrypted_middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup               system.ApiGroup
	ExampleApiGroup              example.ApiGroup
	Encrypted_middlewareApiGroup encrypted_middleware.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
