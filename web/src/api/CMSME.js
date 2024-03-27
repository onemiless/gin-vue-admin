import service from '@/utils/request'

// @Tags CMSME
// @Summary 创建ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CMSME true "创建ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsme/createCMSME [post]
export const createCMSME = (data) => {
  return service({
    url: '/cmsme/createCMSME',
    method: 'post',
    data
  })
}

// @Tags CMSME
// @Summary 删除ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CMSME true "删除ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsme/deleteCMSME [delete]
export const deleteCMSME = (params) => {
  return service({
    url: '/cmsme/deleteCMSME',
    method: 'delete',
    params
  })
}

// @Tags CMSME
// @Summary 批量删除ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsme/deleteCMSME [delete]
export const deleteCMSMEByIds = (params) => {
  return service({
    url: '/cmsme/deleteCMSMEByIds',
    method: 'delete',
    params
  })
}

// @Tags CMSME
// @Summary 更新ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CMSME true "更新ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsme/updateCMSME [put]
export const updateCMSME = (data) => {
  return service({
    url: '/cmsme/updateCMSME',
    method: 'put',
    data
  })
}

// @Tags CMSME
// @Summary 用id查询ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CMSME true "用id查询ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsme/findCMSME [get]
export const findCMSME = (params) => {
  return service({
    url: '/cmsme/findCMSME',
    method: 'get',
    params
  })
}

// @Tags CMSME
// @Summary 分页获取ERP部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ERP部门列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsme/getCMSMEList [get]
export const getCMSMEList = (params) => {
  return service({
    url: '/cmsme/getCMSMEList',
    method: 'get',
    params
  })
}
