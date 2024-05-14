import service from '@/utils/request'

// @Tags QualityBaseInfo
// @Summary 创建质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QualityBaseInfo true "创建质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /qualityBaseInfo/createQualityBaseInfo [post]
export const createQualityBaseInfo = (data) => {
  return service({
    url: '/qualityBaseInfo/createQualityBaseInfo',
    method: 'post',
    data
  })
}

// @Tags QualityBaseInfo
// @Summary 删除质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QualityBaseInfo true "删除质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qualityBaseInfo/deleteQualityBaseInfo [delete]
export const deleteQualityBaseInfo = (params) => {
  return service({
    url: '/qualityBaseInfo/deleteQualityBaseInfo',
    method: 'delete',
    params
  })
}

// @Tags QualityBaseInfo
// @Summary 批量删除质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qualityBaseInfo/deleteQualityBaseInfo [delete]
export const deleteQualityBaseInfoByIds = (params) => {
  return service({
    url: '/qualityBaseInfo/deleteQualityBaseInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags QualityBaseInfo
// @Summary 更新质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QualityBaseInfo true "更新质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /qualityBaseInfo/updateQualityBaseInfo [put]
export const updateQualityBaseInfo = (data) => {
  return service({
    url: '/qualityBaseInfo/updateQualityBaseInfo',
    method: 'put',
    data
  })
}

// @Tags QualityBaseInfo
// @Summary 用id查询质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QualityBaseInfo true "用id查询质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /qualityBaseInfo/findQualityBaseInfo [get]
export const findQualityBaseInfo = (params) => {
  return service({
    url: '/qualityBaseInfo/findQualityBaseInfo',
    method: 'get',
    params
  })
}

// @Tags QualityBaseInfo
// @Summary 分页获取质量要求信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取质量要求信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qualityBaseInfo/getQualityBaseInfoList [get]
export const getQualityBaseInfoList = (params) => {
  return service({
    url: '/qualityBaseInfo/getQualityBaseInfoList',
    method: 'get',
    params
  })
}
