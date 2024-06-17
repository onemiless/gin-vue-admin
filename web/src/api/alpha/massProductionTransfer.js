import service from '@/utils/request'

// @Tags MassProductionTransfer
// @Summary 创建量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MassProductionTransfer true "创建量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /massProductionTransfer/createMassProductionTransfer [post]
export const createMassProductionTransfer = (data) => {
  return service({
    url: '/massProductionTransfer/createMassProductionTransfer',
    method: 'post',
    data
  })
}

// @Tags MassProductionTransfer
// @Summary 删除量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MassProductionTransfer true "删除量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /massProductionTransfer/deleteMassProductionTransfer [delete]
export const deleteMassProductionTransfer = (params) => {
  return service({
    url: '/massProductionTransfer/deleteMassProductionTransfer',
    method: 'delete',
    params
  })
}

// @Tags MassProductionTransfer
// @Summary 批量删除量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /massProductionTransfer/deleteMassProductionTransfer [delete]
export const deleteMassProductionTransferByIds = (params) => {
  return service({
    url: '/massProductionTransfer/deleteMassProductionTransferByIds',
    method: 'delete',
    params
  })
}

// @Tags MassProductionTransfer
// @Summary 更新量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MassProductionTransfer true "更新量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /massProductionTransfer/updateMassProductionTransfer [put]
export const updateMassProductionTransfer = (data) => {
  return service({
    url: '/massProductionTransfer/updateMassProductionTransfer',
    method: 'put',
    data
  })
}

// @Tags MassProductionTransfer
// @Summary 用id查询量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MassProductionTransfer true "用id查询量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /massProductionTransfer/findMassProductionTransfer [get]
export const findMassProductionTransfer = (params) => {
  return service({
    url: '/massProductionTransfer/findMassProductionTransfer',
    method: 'get',
    params
  })
}

// @Tags MassProductionTransfer
// @Summary 分页获取量产转移列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取量产转移列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /massProductionTransfer/getMassProductionTransferList [get]
export const getMassProductionTransferList = (params) => {
  return service({
    url: '/massProductionTransfer/getMassProductionTransferList',
    method: 'get',
    params
  })
}
