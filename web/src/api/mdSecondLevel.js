import service from '@/utils/request'

// @Tags MdSecondLevel
// @Summary 创建层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdSecondLevel true "创建层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdSecondLevel/createMdSecondLevel [post]
export const createMdSecondLevel = (data) => {
  return service({
    url: '/mdSecondLevel/createMdSecondLevel',
    method: 'post',
    data
  })
}

// @Tags MdSecondLevel
// @Summary 删除层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdSecondLevel true "删除层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdSecondLevel/deleteMdSecondLevel [delete]
export const deleteMdSecondLevel = (params) => {
  return service({
    url: '/mdSecondLevel/deleteMdSecondLevel',
    method: 'delete',
    params
  })
}

// @Tags MdSecondLevel
// @Summary 批量删除层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdSecondLevel/deleteMdSecondLevel [delete]
export const deleteMdSecondLevelByIds = (params) => {
  return service({
    url: '/mdSecondLevel/deleteMdSecondLevelByIds',
    method: 'delete',
    params
  })
}

// @Tags MdSecondLevel
// @Summary 更新层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdSecondLevel true "更新层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdSecondLevel/updateMdSecondLevel [put]
export const updateMdSecondLevel = (data) => {
  return service({
    url: '/mdSecondLevel/updateMdSecondLevel',
    method: 'put',
    data
  })
}

// @Tags MdSecondLevel
// @Summary 用id查询层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MdSecondLevel true "用id查询层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdSecondLevel/findMdSecondLevel [get]
export const findMdSecondLevel = (params) => {
  return service({
    url: '/mdSecondLevel/findMdSecondLevel',
    method: 'get',
    params
  })
}

// @Tags MdSecondLevel
// @Summary 分页获取层级二列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取层级二列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdSecondLevel/getMdSecondLevelList [get]
export const getMdSecondLevelList = (params) => {
  return service({
    url: '/mdSecondLevel/getMdSecondLevelList',
    method: 'get',
    params
  })
}
