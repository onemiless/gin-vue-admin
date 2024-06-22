import service from '@/utils/request'

// @Tags PackageRequirement
// @Summary 创建包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PackageRequirement true "创建包装要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /packageRequirement/createPackageRequirement [post]
export const createPackageRequirement = (data) => {
  return service({
    url: '/packageRequirement/createPackageRequirement',
    method: 'post',
    data
  })
}

// @Tags PackageRequirement
// @Summary 删除包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PackageRequirement true "删除包装要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /packageRequirement/deletePackageRequirement [delete]
export const deletePackageRequirement = (params) => {
  return service({
    url: '/packageRequirement/deletePackageRequirement',
    method: 'delete',
    params
  })
}

// @Tags PackageRequirement
// @Summary 批量删除包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除包装要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /packageRequirement/deletePackageRequirement [delete]
export const deletePackageRequirementByIds = (params) => {
  return service({
    url: '/packageRequirement/deletePackageRequirementByIds',
    method: 'delete',
    params
  })
}

// @Tags PackageRequirement
// @Summary 更新包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PackageRequirement true "更新包装要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /packageRequirement/updatePackageRequirement [put]
export const updatePackageRequirement = (data) => {
  return service({
    url: '/packageRequirement/updatePackageRequirement',
    method: 'put',
    data
  })
}

// @Tags PackageRequirement
// @Summary 用id查询包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PackageRequirement true "用id查询包装要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /packageRequirement/findPackageRequirement [get]
export const findPackageRequirement = (params) => {
  return service({
    url: '/packageRequirement/findPackageRequirement',
    method: 'get',
    params
  })
}

// @Tags PackageRequirement
// @Summary 分页获取包装要求信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取包装要求信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /packageRequirement/getPackageRequirementList [get]
export const getPackageRequirementList = (params) => {
  return service({
    url: '/packageRequirement/getPackageRequirementList',
    method: 'get',
    params
  })
}
