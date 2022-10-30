package encrypted_middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/encrypted_middleware"
	encrypted_middlewareReq "github.com/flipped-aurora/gin-vue-admin/server/model/encrypted_middleware/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UnitApi struct {
}

var unitService = service.ServiceGroupApp.Encrypted_middlewareServiceGroup.UnitService

// CreateUnit 创建Unit
// @Tags Unit
// @Summary 创建Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body encrypted_middleware.Unit true "创建Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /unit/createUnit [post]
func (unitApi *UnitApi) CreateUnit(c *gin.Context) {
	var unit encrypted_middleware.Unit
	err := c.ShouldBindJSON(&unit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := unitService.CreateUnit(unit); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUnit 删除Unit
// @Tags Unit
// @Summary 删除Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body encrypted_middleware.Unit true "删除Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /unit/deleteUnit [delete]
func (unitApi *UnitApi) DeleteUnit(c *gin.Context) {
	var unit encrypted_middleware.Unit
	err := c.ShouldBindJSON(&unit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := unitService.DeleteUnit(unit); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUnitByIds 批量删除Unit
// @Tags Unit
// @Summary 批量删除Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /unit/deleteUnitByIds [delete]
func (unitApi *UnitApi) DeleteUnitByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := unitService.DeleteUnitByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUnit 更新Unit
// @Tags Unit
// @Summary 更新Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body encrypted_middleware.Unit true "更新Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /unit/updateUnit [put]
func (unitApi *UnitApi) UpdateUnit(c *gin.Context) {
	var unit encrypted_middleware.Unit
	err := c.ShouldBindJSON(&unit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := unitService.UpdateUnit(unit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUnit 用id查询Unit
// @Tags Unit
// @Summary 用id查询Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query encrypted_middleware.Unit true "用id查询Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /unit/findUnit [get]
func (unitApi *UnitApi) FindUnit(c *gin.Context) {
	var unit encrypted_middleware.Unit
	err := c.ShouldBindQuery(&unit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reunit, err := unitService.GetUnit(unit.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reunit": reunit}, c)
	}
}

// GetUnitList 分页获取Unit列表
// @Tags Unit
// @Summary 分页获取Unit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query encrypted_middlewareReq.UnitSearch true "分页获取Unit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /unit/getUnitList [get]
func (unitApi *UnitApi) GetUnitList(c *gin.Context) {
	var pageInfo encrypted_middlewareReq.UnitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := unitService.GetUnitInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
