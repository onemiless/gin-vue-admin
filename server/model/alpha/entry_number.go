package alpha

// 该表当天应该对应的流水号
type EntryNumber struct {
	TableName string `json:"tableName" form:"tableName" ` //表名       string `json:"typeName" form:"typeName" `  //显示内容
	Param     string `json:"param" form:"param" `
	FiledName string `json:"filedName" form:"filedName" `
}
