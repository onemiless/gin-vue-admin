import service from '@/utils/request'

// @Tags TecBaseCraftsmanship
// @Summary 创建入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseCraftsmanship true "创建入篮量和程序号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseCraftsmanship/createTecBaseCraftsmanship [post]
export const createTecBaseCraftsmanship = (data) => {
  return service({
    url: '/tecBaseCraftsmanship/createTecBaseCraftsmanship',
    method: 'post',
    data
  })
}

// @Tags TecBaseCraftsmanship
// @Summary 删除入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseCraftsmanship true "删除入篮量和程序号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseCraftsmanship/deleteTecBaseCraftsmanship [delete]
export const deleteTecBaseCraftsmanship = (params) => {
  return service({
    url: '/tecBaseCraftsmanship/deleteTecBaseCraftsmanship',
    method: 'delete',
    params
  })
}

// @Tags TecBaseCraftsmanship
// @Summary 批量删除入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除入篮量和程序号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseCraftsmanship/deleteTecBaseCraftsmanship [delete]
export const deleteTecBaseCraftsmanshipByIds = (params) => {
  return service({
    url: '/tecBaseCraftsmanship/deleteTecBaseCraftsmanshipByIds',
    method: 'delete',
    params
  })
}

// @Tags TecBaseCraftsmanship
// @Summary 更新入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TecBaseCraftsmanship true "更新入篮量和程序号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseCraftsmanship/updateTecBaseCraftsmanship [put]
export const updateTecBaseCraftsmanship = (data) => {
  return service({
    url: '/tecBaseCraftsmanship/updateTecBaseCraftsmanship',
    method: 'put',
    data
  })
}

// @Tags TecBaseCraftsmanship
// @Summary 用id查询入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TecBaseCraftsmanship true "用id查询入篮量和程序号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseCraftsmanship/findTecBaseCraftsmanship [get]
export const findTecBaseCraftsmanship = (params) => {
  return service({
    url: '/tecBaseCraftsmanship/findTecBaseCraftsmanship',
    method: 'get',
    params
  })
}

// @Tags TecBaseCraftsmanship
// @Summary 分页获取入篮量和程序号列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取入篮量和程序号列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseCraftsmanship/getTecBaseCraftsmanshipList [get]
export const getTecBaseCraftsmanshipList = (params) => {
  return service({
    url: '/tecBaseCraftsmanship/getTecBaseCraftsmanshipList',
    method: 'get',
    params
  })
}
// @Tags TecBaseCraftsmanship
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseCraftsmanship/findTecBaseCraftsmanshipDataSource [get]
export const getTecBaseCraftsmanshipDataSource = () => {
  return service({
    url: '/tecBaseCraftsmanship/getTecBaseCraftsmanshipDataSource',
    method: 'get',
  })
}
