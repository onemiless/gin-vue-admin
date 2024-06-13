import service from '@/utils/request'

// @Tags TecBaseInfoExt
// @Summary 创建零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseInfoExt true "创建零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseInfoExt/createTecBaseInfoExt [post]
export const createTecBaseInfoExt = (data) => {
  return service({
    url: '/tecBaseInfoExt/createTecBaseInfoExt',
    method: 'post',
    data
  })
}

// @Tags TecBaseInfoExt
// @Summary 删除零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseInfoExt true "删除零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseInfoExt/deleteTecBaseInfoExt [delete]
export const deleteTecBaseInfoExt = (params) => {
  return service({
    url: '/tecBaseInfoExt/deleteTecBaseInfoExt',
    method: 'delete',
    params
  })
}

// @Tags TecBaseInfoExt
// @Summary 批量删除零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseInfoExt/deleteTecBaseInfoExt [delete]
export const deleteTecBaseInfoExtByIds = (params) => {
  return service({
    url: '/tecBaseInfoExt/deleteTecBaseInfoExtByIds',
    method: 'delete',
    params
  })
}

// @Tags TecBaseInfoExt
// @Summary 更新零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseInfoExt true "更新零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseInfoExt/updateTecBaseInfoExt [put]
export const updateTecBaseInfoExt = (data) => {
  return service({
    url: '/tecBaseInfoExt/updateTecBaseInfoExt',
    method: 'put',
    data
  })
}

// @Tags TecBaseInfoExt
// @Summary 用id查询零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TecBaseInfoExt true "用id查询零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseInfoExt/findTecBaseInfoExt [get]
export const findTecBaseInfoExt = (params) => {
  return service({
    url: '/tecBaseInfoExt/findTecBaseInfoExt',
    method: 'get',
    params
  })
}

// @Tags TecBaseInfoExt
// @Summary 分页获取零件基础信息扩展列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取零件基础信息扩展列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseInfoExt/getTecBaseInfoExtList [get]
export const getTecBaseInfoExtList = (params) => {
  return service({
    url: '/tecBaseInfoExt/getTecBaseInfoExtList',
    method: 'get',
    params
  })
}

export const checkIsDuplicate = (params) =>{
  return service({
    url: '/tecBaseInfoExt/checkIsDuplicate',
    method: 'get',
    params
  })
}
