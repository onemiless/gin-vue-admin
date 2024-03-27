import service from '@/utils/request'

// @Tags MdUnitMeasure
// @Summary 创建品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdUnitMeasure true "创建品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdUnitMeasure/createMdUnitMeasure [post]
export const createMdUnitMeasure = (data) => {
  return service({
    url: '/mdUnitMeasure/createMdUnitMeasure',
    method: 'post',
    data
  })
}

// @Tags MdUnitMeasure
// @Summary 删除品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdUnitMeasure true "删除品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdUnitMeasure/deleteMdUnitMeasure [delete]
export const deleteMdUnitMeasure = (params) => {
  return service({
    url: '/mdUnitMeasure/deleteMdUnitMeasure',
    method: 'delete',
    params
  })
}

// @Tags MdUnitMeasure
// @Summary 批量删除品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdUnitMeasure/deleteMdUnitMeasure [delete]
export const deleteMdUnitMeasureByIds = (params) => {
  return service({
    url: '/mdUnitMeasure/deleteMdUnitMeasureByIds',
    method: 'delete',
    params
  })
}

// @Tags MdUnitMeasure
// @Summary 更新品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdUnitMeasure true "更新品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdUnitMeasure/updateMdUnitMeasure [put]
export const updateMdUnitMeasure = (data) => {
  return service({
    url: '/mdUnitMeasure/updateMdUnitMeasure',
    method: 'put',
    data
  })
}

// @Tags MdUnitMeasure
// @Summary 用id查询品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MdUnitMeasure true "用id查询品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdUnitMeasure/findMdUnitMeasure [get]
export const findMdUnitMeasure = (params) => {
  return service({
    url: '/mdUnitMeasure/findMdUnitMeasure',
    method: 'get',
    params
  })
}

// @Tags MdUnitMeasure
// @Summary 分页获取品号单位设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取品号单位设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdUnitMeasure/getMdUnitMeasureList [get]
export const getMdUnitMeasureList = (params) => {
  return service({
    url: '/mdUnitMeasure/getMdUnitMeasureList',
    method: 'get',
    params
  })
}
