import service from '@/utils/request'

// @Tags ProofingInformation
// @Summary 创建打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProofingInformation true "创建打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /proofingInformation/createProofingInformation [post]
export const createProofingInformation = (data) => {
  return service({
    url: '/proofingInformation/createProofingInformation',
    method: 'post',
    data
  })
}

// @Tags ProofingInformation
// @Summary 删除打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProofingInformation true "删除打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /proofingInformation/deleteProofingInformation [delete]
export const deleteProofingInformation = (params) => {
  return service({
    url: '/proofingInformation/deleteProofingInformation',
    method: 'delete',
    params
  })
}

// @Tags ProofingInformation
// @Summary 批量删除打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /proofingInformation/deleteProofingInformation [delete]
export const deleteProofingInformationByIds = (params) => {
  return service({
    url: '/proofingInformation/deleteProofingInformationByIds',
    method: 'delete',
    params
  })
}

// @Tags ProofingInformation
// @Summary 更新打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProofingInformation true "更新打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /proofingInformation/updateProofingInformation [put]
export const updateProofingInformation = (data) => {
  return service({
    url: '/proofingInformation/updateProofingInformation',
    method: 'put',
    data
  })
}

// @Tags ProofingInformation
// @Summary 用id查询打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProofingInformation true "用id查询打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /proofingInformation/findProofingInformation [get]
export const findProofingInformation = (params) => {
  return service({
    url: '/proofingInformation/findProofingInformation',
    method: 'get',
    params
  })
}

// @Tags ProofingInformation
// @Summary 分页获取打样信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取打样信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /proofingInformation/getProofingInformationList [get]
export const getProofingInformationList = (params) => {
  return service({
    url: '/proofingInformation/getProofingInformationList',
    method: 'get',
    params
  })
}
