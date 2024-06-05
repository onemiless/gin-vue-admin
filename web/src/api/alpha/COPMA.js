import service from '@/utils/request'

// @Tags COPMA
// @Summary 创建ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.COPMA true "创建ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /copma/createCOPMA [post]
export const createCOPMA = (data) => {
  return service({
    url: '/copma/createCOPMA',
    method: 'post',
    data
  })
}

// @Tags COPMA
// @Summary 删除ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.COPMA true "删除ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /copma/deleteCOPMA [delete]
export const deleteCOPMA = (params) => {
  return service({
    url: '/copma/deleteCOPMA',
    method: 'delete',
    params
  })
}

// @Tags COPMA
// @Summary 批量删除ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /copma/deleteCOPMA [delete]
export const deleteCOPMAByIds = (params) => {
  return service({
    url: '/copma/deleteCOPMAByIds',
    method: 'delete',
    params
  })
}

// @Tags COPMA
// @Summary 更新ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.COPMA true "更新ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /copma/updateCOPMA [put]
export const updateCOPMA = (data) => {
  return service({
    url: '/copma/updateCOPMA',
    method: 'put',
    data
  })
}

// @Tags COPMA
// @Summary 用id查询ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.COPMA true "用id查询ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /copma/findCOPMA [get]
export const findCOPMA = (params) => {
  return service({
    url: '/copma/findCOPMA',
    method: 'get',
    params
  })
}

// @Tags COPMA
// @Summary 分页获取ERP客户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ERP客户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /copma/getCOPMAList [get]
export const getCOPMAList = (params) => {
  return service({
    url: '/copma/getCOPMAList',
    method: 'get',
    params
  })
}
