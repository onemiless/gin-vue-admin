import service from '@/utils/request'

// @Tags CostCollection
// @Summary 创建成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CostCollection true "创建成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /costCollection/createCostCollection [post]
export const createCostCollection = (data) => {
  return service({
    url: '/costCollection/createCostCollection',
    method: 'post',
    data
  })
}

// @Tags CostCollection
// @Summary 删除成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CostCollection true "删除成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /costCollection/deleteCostCollection [delete]
export const deleteCostCollection = (params) => {
  return service({
    url: '/costCollection/deleteCostCollection',
    method: 'delete',
    params
  })
}

// @Tags CostCollection
// @Summary 批量删除成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /costCollection/deleteCostCollection [delete]
export const deleteCostCollectionByIds = (params) => {
  return service({
    url: '/costCollection/deleteCostCollectionByIds',
    method: 'delete',
    params
  })
}

// @Tags CostCollection
// @Summary 更新成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CostCollection true "更新成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /costCollection/updateCostCollection [put]
export const updateCostCollection = (data) => {
  return service({
    url: '/costCollection/updateCostCollection',
    method: 'put',
    data
  })
}

// @Tags CostCollection
// @Summary 用id查询成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CostCollection true "用id查询成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /costCollection/findCostCollection [get]
export const findCostCollection = (params) => {
  return service({
    url: '/costCollection/findCostCollection',
    method: 'get',
    params
  })
}
export const findCostCollectionNumber = ()=>{
  return service({
    url:'/costCollection/getCostCollectionNumer',
    method: 'get'
  })
}

// @Tags CostCollection
// @Summary 分页获取成本信息收集信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取成本信息收集信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /costCollection/getCostCollectionList [get]
export const getCostCollectionList = (params) => {
  return service({
    url: '/costCollection/getCostCollectionList',
    method: 'get',
    params
  })
}
