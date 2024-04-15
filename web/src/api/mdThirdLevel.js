import service from '@/utils/request'

// @Tags MdThirdLevel
// @Summary 创建层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdThirdLevel true "创建层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdThirdLevel/createMdThirdLevel [post]
export const createMdThirdLevel = (data) => {
  return service({
    url: '/mdThirdLevel/createMdThirdLevel',
    method: 'post',
    data
  })
}

// @Tags MdThirdLevel
// @Summary 删除层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdThirdLevel true "删除层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdThirdLevel/deleteMdThirdLevel [delete]
export const deleteMdThirdLevel = (params) => {
  return service({
    url: '/mdThirdLevel/deleteMdThirdLevel',
    method: 'delete',
    params
  })
}

// @Tags MdThirdLevel
// @Summary 批量删除层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdThirdLevel/deleteMdThirdLevel [delete]
export const deleteMdThirdLevelByIds = (params) => {
  return service({
    url: '/mdThirdLevel/deleteMdThirdLevelByIds',
    method: 'delete',
    params
  })
}

// @Tags MdThirdLevel
// @Summary 更新层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MdThirdLevel true "更新层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdThirdLevel/updateMdThirdLevel [put]
export const updateMdThirdLevel = (data) => {
  return service({
    url: '/mdThirdLevel/updateMdThirdLevel',
    method: 'put',
    data
  })
}

// @Tags MdThirdLevel
// @Summary 用id查询层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MdThirdLevel true "用id查询层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdThirdLevel/findMdThirdLevel [get]
export const findMdThirdLevel = (params) => {
  return service({
    url: '/mdThirdLevel/findMdThirdLevel',
    method: 'get',
    params
  })
}

// @Tags MdThirdLevel
// @Summary 分页获取层级三列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取层级三列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdThirdLevel/getMdThirdLevelList [get]
export const getMdThirdLevelList = (params) => {
  return service({
    url: '/mdThirdLevel/getMdThirdLevelList',
    method: 'get',
    params
  })
}
