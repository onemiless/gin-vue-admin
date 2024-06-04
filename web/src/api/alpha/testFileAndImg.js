import service from '@/utils/request'

// @Tags TestFileAndImg
// @Summary 创建TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestFileAndImg true "创建TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /testFileAndImg/createTestFileAndImg [post]
export const createTestFileAndImg = (data) => {
  return service({
    url: '/testFileAndImg/createTestFileAndImg',
    method: 'post',
    data
  })
}

// @Tags TestFileAndImg
// @Summary 删除TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestFileAndImg true "删除TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testFileAndImg/deleteTestFileAndImg [delete]
export const deleteTestFileAndImg = (params) => {
  return service({
    url: '/testFileAndImg/deleteTestFileAndImg',
    method: 'delete',
    params
  })
}

// @Tags TestFileAndImg
// @Summary 批量删除TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testFileAndImg/deleteTestFileAndImg [delete]
export const deleteTestFileAndImgByIds = (params) => {
  return service({
    url: '/testFileAndImg/deleteTestFileAndImgByIds',
    method: 'delete',
    params
  })
}

// @Tags TestFileAndImg
// @Summary 更新TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestFileAndImg true "更新TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testFileAndImg/updateTestFileAndImg [put]
export const updateTestFileAndImg = (data) => {
  return service({
    url: '/testFileAndImg/updateTestFileAndImg',
    method: 'put',
    data
  })
}

// @Tags TestFileAndImg
// @Summary 用id查询TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestFileAndImg true "用id查询TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testFileAndImg/findTestFileAndImg [get]
export const findTestFileAndImg = (params) => {
  return service({
    url: '/testFileAndImg/findTestFileAndImg',
    method: 'get',
    params
  })
}

// @Tags TestFileAndImg
// @Summary 分页获取TestFileAndImg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestFileAndImg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testFileAndImg/getTestFileAndImgList [get]
export const getTestFileAndImgList = (params) => {
  return service({
    url: '/testFileAndImg/getTestFileAndImgList',
    method: 'get',
    params
  })
}
