import service from '@/utils/request'

// @Tags CMSXC
// @Summary 创建车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CMSXC true "创建车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsxc/createCMSXC [post]
export const createCMSXC = (data) => {
  return service({
    url: '/cmsxc/createCMSXC',
    method: 'post',
    data
  })
}

// @Tags CMSXC
// @Summary 删除车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CMSXC true "删除车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsxc/deleteCMSXC [delete]
export const deleteCMSXC = (params) => {
  return service({
    url: '/cmsxc/deleteCMSXC',
    method: 'delete',
    params
  })
}

// @Tags CMSXC
// @Summary 批量删除车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsxc/deleteCMSXC [delete]
export const deleteCMSXCByIds = (params) => {
  return service({
    url: '/cmsxc/deleteCMSXCByIds',
    method: 'delete',
    params
  })
}

// @Tags CMSXC
// @Summary 更新车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CMSXC true "更新车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsxc/updateCMSXC [put]
export const updateCMSXC = (data) => {
  return service({
    url: '/cmsxc/updateCMSXC',
    method: 'put',
    data
  })
}

// @Tags CMSXC
// @Summary 用id查询车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CMSXC true "用id查询车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsxc/findCMSXC [get]
export const findCMSXC = (params) => {
  return service({
    url: '/cmsxc/findCMSXC',
    method: 'get',
    params
  })
}

// @Tags CMSXC
// @Summary 分页获取车厂规范信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取车厂规范信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsxc/getCMSXCList [get]
export const getCMSXCList = (params) => {
  return service({
    url: '/cmsxc/getCMSXCList',
    method: 'get',
    params
  })
}
