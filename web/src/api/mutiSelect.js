import service from '@/utils/request'

// @Tags MutiSelect
// @Summary 创建多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MutiSelect true "创建多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mutiSelect/createMutiSelect [post]
export const createMutiSelect = (data) => {
  return service({
    url: '/mutiSelect/createMutiSelect',
    method: 'post',
    data
  })
}

// @Tags MutiSelect
// @Summary 删除多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MutiSelect true "删除多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mutiSelect/deleteMutiSelect [delete]
export const deleteMutiSelect = (params) => {
  return service({
    url: '/mutiSelect/deleteMutiSelect',
    method: 'delete',
    params
  })
}

// @Tags MutiSelect
// @Summary 批量删除多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mutiSelect/deleteMutiSelect [delete]
export const deleteMutiSelectByIds = (params) => {
  return service({
    url: '/mutiSelect/deleteMutiSelectByIds',
    method: 'delete',
    params
  })
}

// @Tags MutiSelect
// @Summary 更新多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MutiSelect true "更新多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mutiSelect/updateMutiSelect [put]
export const updateMutiSelect = (data) => {
  return service({
    url: '/mutiSelect/updateMutiSelect',
    method: 'put',
    data
  })
}

// @Tags MutiSelect
// @Summary 用id查询多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MutiSelect true "用id查询多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mutiSelect/findMutiSelect [get]
export const findMutiSelect = (params) => {
  return service({
    url: '/mutiSelect/findMutiSelect',
    method: 'get',
    params
  })
}

// @Tags MutiSelect
// @Summary 分页获取多级选择信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取多级选择信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mutiSelect/getMutiSelectList [get]
export const getMutiSelectList = (params) => {
  return service({
    url: '/mutiSelect/getMutiSelectList',
    method: 'get',
    params
  })
}
