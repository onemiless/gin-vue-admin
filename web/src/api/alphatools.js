import service from '@/utils/request'


// @Tags getEntryNumber
// @Summary 用于获取该表的流水号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data tablename true "表名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /costCollection/findCostCollection [get]
export const getEntryNumber = (params) => {
  return service({
    url: '/entryNumber/getEntryNumber',
    method: 'get',
    params
  })
}



