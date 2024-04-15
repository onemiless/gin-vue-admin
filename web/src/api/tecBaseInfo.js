import service from '@/utils/request'

// @Tags TecBaseInfo
// @Summary 创建技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseInfo true "创建技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseInfo/createTecBaseInfo [post]
export const createTecBaseInfo = (data) => {
  return service({
    url: '/tecBaseInfo/createTecBaseInfo',
    method: 'post',
    data
  })
}

// @Tags TecBaseInfo
// @Summary 删除技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseInfo true "删除技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseInfo/deleteTecBaseInfo [delete]
export const deleteTecBaseInfo = (params) => {
  return service({
    url: '/tecBaseInfo/deleteTecBaseInfo',
    method: 'delete',
    params
  })
}

// @Tags TecBaseInfo
// @Summary 批量删除技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseInfo/deleteTecBaseInfo [delete]
export const deleteTecBaseInfoByIds = (params) => {
  return service({
    url: '/tecBaseInfo/deleteTecBaseInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags TecBaseInfo
// @Summary 更新技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseInfo true "更新技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseInfo/updateTecBaseInfo [put]
export const updateTecBaseInfo = (data) => {
  return service({
    url: '/tecBaseInfo/updateTecBaseInfo',
    method: 'put',
    data
  })
}

// @Tags TecBaseInfo
// @Summary 用id查询技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TecBaseInfo true "用id查询技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseInfo/findTecBaseInfo [get]
export const findTecBaseInfo = (params) => {
  return service({
    url: '/tecBaseInfo/findTecBaseInfo',
    method: 'get',
    params
  })
}

// @Tags TecBaseInfo
// @Summary 分页获取技术部基础信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取技术部基础信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseInfo/getTecBaseInfoList [get]
export const getTecBaseInfoList = (params) => {
  return service({
    url: '/tecBaseInfo/getTecBaseInfoList',
    method: 'get',
    params
  })
}
