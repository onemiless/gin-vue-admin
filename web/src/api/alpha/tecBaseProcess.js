import service from '@/utils/request'

// @Tags TecBaseProcess
// @Summary 创建工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseProcess true "创建工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseProcess/createTecBaseProcess [post]
export const createTecBaseProcess = (data) => {
  return service({
    url: '/tecBaseProcess/createTecBaseProcess',
    method: 'post',
    data
  })
}

// @Tags TecBaseProcess
// @Summary 删除工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseProcess true "删除工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseProcess/deleteTecBaseProcess [delete]
export const deleteTecBaseProcess = (params) => {
  return service({
    url: '/tecBaseProcess/deleteTecBaseProcess',
    method: 'delete',
    params
  })
}

// @Tags TecBaseProcess
// @Summary 批量删除工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseProcess/deleteTecBaseProcess [delete]
export const deleteTecBaseProcessByIds = (params) => {
  return service({
    url: '/tecBaseProcess/deleteTecBaseProcessByIds',
    method: 'delete',
    params
  })
}

// @Tags TecBaseProcess
// @Summary 更新工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseProcess true "更新工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseProcess/updateTecBaseProcess [put]
export const updateTecBaseProcess = (data) => {
  return service({
    url: '/tecBaseProcess/updateTecBaseProcess',
    method: 'put',
    data
  })
}

// @Tags TecBaseProcess
// @Summary 用id查询工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TecBaseProcess true "用id查询工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseProcess/findTecBaseProcess [get]
export const findTecBaseProcess = (params) => {
  return service({
    url: '/tecBaseProcess/findTecBaseProcess',
    method: 'get',
    params
  })
}

// @Tags TecBaseProcess
// @Summary 分页获取工艺、设备及治具基本信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工艺、设备及治具基本信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseProcess/getTecBaseProcessList [get]
export const getTecBaseProcessList = (params) => {
  return service({
    url: '/tecBaseProcess/getTecBaseProcessList',
    method: 'get',
    params
  })
}
