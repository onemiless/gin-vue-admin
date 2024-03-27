import service from '@/utils/request'

// @Tags MdClient
// @Summary 创建客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdClient true "创建客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdClient/createMdClient [post]
export const createMdClient = (data) => {
  return service({
    url: '/mdClient/createMdClient',
    method: 'post',
    data
  })
}

// @Tags MdClient
// @Summary 删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdClient true "删除客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdClient/deleteMdClient [delete]
export const deleteMdClient = (params) => {
  return service({
    url: '/mdClient/deleteMdClient',
    method: 'delete',
    params
  })
}

// @Tags MdClient
// @Summary 批量删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdClient/deleteMdClient [delete]
export const deleteMdClientByIds = (params) => {
  return service({
    url: '/mdClient/deleteMdClientByIds',
    method: 'delete',
    params
  })
}

// @Tags MdClient
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdClient true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdClient/updateMdClient [put]
export const updateMdClient = (data) => {
  return service({
    url: '/mdClient/updateMdClient',
    method: 'put',
    data
  })
}

// @Tags MdClient
// @Summary 用id查询客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MdClient true "用id查询客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdClient/findMdClient [get]
export const findMdClient = (params) => {
  return service({
    url: '/mdClient/findMdClient',
    method: 'get',
    params
  })
}

// @Tags MdClient
// @Summary 分页获取客户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取客户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdClient/getMdClientList [get]
export const getMdClientList = (params) => {
  return service({
    url: '/mdClient/getMdClientList',
    method: 'get',
    params
  })
}
