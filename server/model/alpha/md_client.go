// 自动生成模板MdClient
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 客户信息 结构体  MdClient
type MdClient struct {
	global.GVA_MODEL
	ClientCode string `json:"clientCode" form:"clientCode" gorm:"column:client_code;comment:客户编码;size:64;"binding:"required"`   //客户编码
	ClientName string `json:"clientName" form:"clientName" gorm:"column:client_name;comment:客户名称;size:255;"binding:"required"`  //客户名称
	ClientEn   string `json:"clientEn" form:"clientEn" gorm:"column:client_en;comment:英文名称;size:255;"`                          //英文名称
	ClientDes  string `json:"clientDes" form:"clientDes" gorm:"column:client_des;comment:描述;size:500;"`                         //描述
	ClientType string `json:"clientType" form:"clientType" gorm:"column:client_first;comment:拼音首字母;size:64;"binding:"required"` //拼音首字母
	Address    string `json:"address" form:"address" gorm:"column:address;comment:地址;size:500;"`                                //地址
	Email      string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:255;"`                                      //邮箱
	Tel        string `json:"tel" form:"tel" gorm:"column:tel;comment:客户电话;size:64;"`                                           //客户电话
	EnableFlag *int   `json:"enableFlag" form:"enableFlag" gorm:"column:enable_flag;comment:启用;"binding:"required"`             //启用
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;type:text;"`                         //备注
	Attr1      string `json:"attr1" form:"attr1" gorm:"column:attr1;comment:;size:64;"`                                         //attr1字段
	Attr2      string `json:"attr2" form:"attr2" gorm:"column:attr2;comment:;size:255;"`                                        //attr2字段
	Attr3      *int   `json:"attr3" form:"attr3" gorm:"column:attr3;comment:;"`                                                 //attr3字段
	Attr4      *int   `json:"attr4" form:"attr4" gorm:"column:attr4;comment:;"`                                                 //attr4字段
}

// TableName 客户信息 MdClient自定义表名 md_client
func (MdClient) TableName() string {
	return "md_client"
}
