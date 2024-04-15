import service from '@/utils/request'

// @Tags MdFirstLevel
// @Summary 创建层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdFirstLevel true "创建层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdFirstLevel/createMdFirstLevel [post]
export const createMdFirstLevel = (data) => {
  return service({
    url: '/mdFirstLevel/createMdFirstLevel',
    method: 'post',
    data
  })
}

// @Tags MdFirstLevel
// @Summary 删除层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdFirstLevel true "删除层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdFirstLevel/deleteMdFirstLevel [delete]
export const deleteMdFirstLevel = (params) => {
  return service({
    url: '/mdFirstLevel/deleteMdFirstLevel',
    method: 'delete',
    params
  })
}

// @Tags MdFirstLevel
// @Summary 批量删除层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdFirstLevel/deleteMdFirstLevel [delete]
export const deleteMdFirstLevelByIds = (params) => {
  return service({
    url: '/mdFirstLevel/deleteMdFirstLevelByIds',
    method: 'delete',
    params
  })
}

// @Tags MdFirstLevel
// @Summary 更新层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdFirstLevel true "更新层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdFirstLevel/updateMdFirstLevel [put]
export const updateMdFirstLevel = (data) => {
  return service({
    url: '/mdFirstLevel/updateMdFirstLevel',
    method: 'put',
    data
  })
}

// @Tags MdFirstLevel
// @Summary 用id查询层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MdFirstLevel true "用id查询层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdFirstLevel/findMdFirstLevel [get]
export const findMdFirstLevel = (params) => {
  return service({
    url: '/mdFirstLevel/findMdFirstLevel',
    method: 'get',
    params
  })
}

// @Tags MdFirstLevel
// @Summary 分页获取层级一列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取层级一列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdFirstLevel/getMdFirstLevelList [get]
export const getMdFirstLevelList = (params) => {
  return service({
    url: '/mdFirstLevel/getMdFirstLevelList',
    method: 'get',
    params
  })
}
