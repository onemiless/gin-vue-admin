import service from '@/utils/request'

// @Tags EngineeringChange
// @Summary 创建工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EngineeringChange true "创建工程变更信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /engineeringChange/createEngineeringChange [post]
export const createEngineeringChange = (data) => {
  return service({
    url: '/engineeringChange/createEngineeringChange',
    method: 'post',
    data
  })
}

// @Tags EngineeringChange
// @Summary 删除工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EngineeringChange true "删除工程变更信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /engineeringChange/deleteEngineeringChange [delete]
export const deleteEngineeringChange = (params) => {
  return service({
    url: '/engineeringChange/deleteEngineeringChange',
    method: 'delete',
    params
  })
}

// @Tags EngineeringChange
// @Summary 批量删除工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工程变更信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /engineeringChange/deleteEngineeringChange [delete]
export const deleteEngineeringChangeByIds = (params) => {
  return service({
    url: '/engineeringChange/deleteEngineeringChangeByIds',
    method: 'delete',
    params
  })
}

// @Tags EngineeringChange
// @Summary 更新工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EngineeringChange true "更新工程变更信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /engineeringChange/updateEngineeringChange [put]
export const updateEngineeringChange = (data) => {
  return service({
    url: '/engineeringChange/updateEngineeringChange',
    method: 'put',
    data
  })
}

// @Tags EngineeringChange
// @Summary 用id查询工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EngineeringChange true "用id查询工程变更信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /engineeringChange/findEngineeringChange [get]
export const findEngineeringChange = (params) => {
  return service({
    url: '/engineeringChange/findEngineeringChange',
    method: 'get',
    params
  })
}

// @Tags EngineeringChange
// @Summary 分页获取工程变更信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工程变更信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /engineeringChange/getEngineeringChangeList [get]
export const getEngineeringChangeList = (params) => {
  return service({
    url: '/engineeringChange/getEngineeringChangeList',
    method: 'get',
    params
  })
}
