package encrypted_middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UnitRouter struct {
}

// InitUnitRouter 初始化 Unit 路由信息
func (s *UnitRouter) InitUnitRouter(Router *gin.RouterGroup) {
	unitRouter := Router.Group("unit").Use(middleware.OperationRecord())
	unitRouterWithoutRecord := Router.Group("unit")
	var unitApi = v1.ApiGroupApp.Encrypted_middlewareApiGroup.UnitApi
	{
		unitRouter.POST("createUnit", unitApi.CreateUnit)             // 新建Unit
		unitRouter.DELETE("deleteUnit", unitApi.DeleteUnit)           // 删除Unit
		unitRouter.DELETE("deleteUnitByIds", unitApi.DeleteUnitByIds) // 批量删除Unit
		unitRouter.PUT("updateUnit", unitApi.UpdateUnit)              // 更新Unit
	}
	{
		unitRouterWithoutRecord.GET("findUnit", unitApi.FindUnit)       // 根据ID获取Unit
		unitRouterWithoutRecord.GET("getUnitList", unitApi.GetUnitList) // 获取Unit列表
	}
}
