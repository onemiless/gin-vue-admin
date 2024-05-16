package alpha

// 该表当天应该对应的流水号
type EntryNumber struct {
	TableName string `json:"tableName" form:"tableName" ` //表名       string `json:"typeName" form:"typeName" `  //显示内容

}

// TableName 多级选择信息 MutiSelect自定义表名 muti_select
//func (MutiSelect) TableName() string {
//	return "muti_select"
//}
