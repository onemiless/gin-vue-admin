// 自动生成模板TestFileAndImg
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// TestFileAndImg 结构体  TestFileAndImg
type TestFileAndImg struct {
	global.GVA_MODEL
	File      datatypes.JSON `json:"file" form:"file" gorm:"column:file;comment:;type:nvarchar(max);"`          //file
	Img       string         `json:"img" form:"img" gorm:"column:img;comment:;type:nvarchar(max);"`             //img
	Multimg   datatypes.JSON `json:"multimg" form:"multimg" gorm:"column:multimg;comment:;type:nvarchar(max);"` //multimg
	CreatedBy uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint           `gorm:"column:deleted_by;comment:删除者"`
}

// TableName TestFileAndImg TestFileAndImg自定义表名 test_file_and_img
func (TestFileAndImg) TableName() string {
	return "test_file_and_img"
}
