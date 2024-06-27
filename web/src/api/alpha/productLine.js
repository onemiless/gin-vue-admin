import service from '@/utils/request'

// @Tags ProductLine
// @Summary 创建产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLine true "创建产线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /productLine/createProductLine [post]
export const createProductLine = (data) => {
  return service({
    url: '/productLine/createProductLine',
    method: 'post',
    data
  })
}

// @Tags ProductLine
// @Summary 删除产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLine true "删除产线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productLine/deleteProductLine [delete]
export const deleteProductLine = (params) => {
  return service({
    url: '/productLine/deleteProductLine',
    method: 'delete',
    params
  })
}

// @Tags ProductLine
// @Summary 批量删除产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除产线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productLine/deleteProductLine [delete]
export const deleteProductLineByIds = (params) => {
  return service({
    url: '/productLine/deleteProductLineByIds',
    method: 'delete',
    params
  })
}

// @Tags ProductLine
// @Summary 更新产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLine true "更新产线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /productLine/updateProductLine [put]
export const updateProductLine = (data) => {
  return service({
    url: '/productLine/updateProductLine',
    method: 'put',
    data
  })
}

// @Tags ProductLine
// @Summary 用id查询产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProductLine true "用id查询产线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /productLine/findProductLine [get]
export const findProductLine = (params) => {
  return service({
    url: '/productLine/findProductLine',
    method: 'get',
    params
  })
}

// @Tags ProductLine
// @Summary 分页获取产线信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取产线信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productLine/getProductLineList [get]
export const getProductLineList = (params) => {
  return service({
    url: '/productLine/getProductLineList',
    method: 'get',
    params
  })
}
