import service from '@/utils/request'

// @Tags ProcessFileInformation
// @Summary 创建工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProcessFileInformation true "创建工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /processFileInformation/createProcessFileInformation [post]
export const createProcessFileInformation = (data) => {
  return service({
    url: '/processFileInformation/createProcessFileInformation',
    method: 'post',
    data
  })
}

// @Tags ProcessFileInformation
// @Summary 删除工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProcessFileInformation true "删除工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /processFileInformation/deleteProcessFileInformation [delete]
export const deleteProcessFileInformation = (params) => {
  return service({
    url: '/processFileInformation/deleteProcessFileInformation',
    method: 'delete',
    params
  })
}

// @Tags ProcessFileInformation
// @Summary 批量删除工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /processFileInformation/deleteProcessFileInformation [delete]
export const deleteProcessFileInformationByIds = (params) => {
  return service({
    url: '/processFileInformation/deleteProcessFileInformationByIds',
    method: 'delete',
    params
  })
}

// @Tags ProcessFileInformation
// @Summary 更新工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProcessFileInformation true "更新工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /processFileInformation/updateProcessFileInformation [put]
export const updateProcessFileInformation = (data) => {
  return service({
    url: '/processFileInformation/updateProcessFileInformation',
    method: 'put',
    data
  })
}

// @Tags ProcessFileInformation
// @Summary 用id查询工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProcessFileInformation true "用id查询工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /processFileInformation/findProcessFileInformation [get]
export const findProcessFileInformation = (params) => {
  return service({
    url: '/processFileInformation/findProcessFileInformation',
    method: 'get',
    params
  })
}

// @Tags ProcessFileInformation
// @Summary 分页获取工艺文件信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工艺文件信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /processFileInformation/getProcessFileInformationList [get]
export const getProcessFileInformationList = (params) => {
  return service({
    url: '/processFileInformation/getProcessFileInformationList',
    method: 'get',
    params
  })
}
