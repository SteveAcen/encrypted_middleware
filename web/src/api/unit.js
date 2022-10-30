import service from '@/utils/request'

// @Tags Unit
// @Summary 创建Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Unit true "创建Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /unit/createUnit [post]
export const createUnit = (data) => {
  return service({
    url: '/unit/createUnit',
    method: 'post',
    data
  })
}

// @Tags Unit
// @Summary 删除Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Unit true "删除Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /unit/deleteUnit [delete]
export const deleteUnit = (data) => {
  return service({
    url: '/unit/deleteUnit',
    method: 'delete',
    data
  })
}

// @Tags Unit
// @Summary 删除Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /unit/deleteUnit [delete]
export const deleteUnitByIds = (data) => {
  return service({
    url: '/unit/deleteUnitByIds',
    method: 'delete',
    data
  })
}

// @Tags Unit
// @Summary 更新Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Unit true "更新Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /unit/updateUnit [put]
export const updateUnit = (data) => {
  return service({
    url: '/unit/updateUnit',
    method: 'put',
    data
  })
}

// @Tags Unit
// @Summary 用id查询Unit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Unit true "用id查询Unit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /unit/findUnit [get]
export const findUnit = (params) => {
  return service({
    url: '/unit/findUnit',
    method: 'get',
    params
  })
}

// @Tags Unit
// @Summary 分页获取Unit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Unit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /unit/getUnitList [get]
export const getUnitList = (params) => {
  return service({
    url: '/unit/getUnitList',
    method: 'get',
    params
  })
}
