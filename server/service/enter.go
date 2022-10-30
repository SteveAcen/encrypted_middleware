package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/encrypted_middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup               system.ServiceGroup
	ExampleServiceGroup              example.ServiceGroup
	Encrypted_middlewareServiceGroup encrypted_middleware.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
