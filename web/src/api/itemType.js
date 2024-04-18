import service from '@/utils/request'

// @Tags ItemType
// @Summary 创建零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ItemType true "创建零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /itemtype/createItemType [post]
export const createItemType = (data) => {
  return service({
    url: '/itemtype/createItemType',
    method: 'post',
    data
  })
}

// @Tags ItemType
// @Summary 删除零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ItemType true "删除零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itemtype/deleteItemType [delete]
export const deleteItemType = (params) => {
  return service({
    url: '/itemtype/deleteItemType',
    method: 'delete',
    params
  })
}

// @Tags ItemType
// @Summary 批量删除零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itemtype/deleteItemType [delete]
export const deleteItemTypeByIds = (params) => {
  return service({
    url: '/itemtype/deleteItemTypeByIds',
    method: 'delete',
    params
  })
}

// @Tags ItemType
// @Summary 更新零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ItemType true "更新零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /itemtype/updateItemType [put]
export const updateItemType = (data) => {
  return service({
    url: '/itemtype/updateItemType',
    method: 'put',
    data
  })
}

// @Tags ItemType
// @Summary 用id查询零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ItemType true "用id查询零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /itemtype/findItemType [get]
export const findItemType = (params) => {
  return service({
    url: '/itemtype/findItemType',
    method: 'get',
    params
  })
}

// @Tags ItemType
// @Summary 分页获取零件类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取零件类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itemtype/getItemTypeList [get]
export const getItemTypeList = (params) => {
  return service({
    url: '/itemtype/getItemTypeList',
    method: 'get',
    params
  })
}

//获取父菜单
export const getOptionsFromBackend=(params)=>{
  return service({
    url: '/itemtype/getItemTypeListParent',
    method: 'get',
    params
  })
}
//根据父菜单，获取子菜单
export const getOptionsFromBackendByParentId=(params)=>{
  return service({
    url: '/itemtype/getItemTypeListParent',
    method: 'get',
    params
  })
}