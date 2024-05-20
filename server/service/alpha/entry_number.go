package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	"strconv"
	"time"
)

type IEntryNumberService struct {
}

// 获取流水号
func (entryNumberService *IEntryNumberService) GetEntryNumber(entryNumber alpha.EntryNumber) (number string, err error) {
	//err = global.GVA_DB.Create(itemtype).Error
	date := time.Now().Format("2006-01-02")
	//根据表名动态查询
	if entryNumber.Param == "" {
		err = global.GVA_DB.Raw("select count(created_at ) from " + entryNumber.TableName + " where created_at BETWEEN '" + date + " 00:00:00' AND '" + date + " 23:59:59'").Scan(&number).Error
	} else {
		//if entryNumber.TableName == "cost_collection" {
		err = global.GVA_DB.Raw("select count(created_at ) from "+entryNumber.TableName+" where "+entryNumber.FiledName+" = ?", entryNumber.Param).Scan(&number).Error
		//}
	}
	if err == nil {

		numberTemp, err := strconv.Atoi(number)
		if err != nil {
			return "001", err
		}
		numberTemp = numberTemp + 1
		//数字转字符串,并填充流水号格式为001
		numberTemp = numberTemp + 1000
		numberTemp = numberTemp % 10000
		//截取后三位
		number = strconv.Itoa(numberTemp)[1:]

	} else {
		number = "001"
		err = nil
	}
	return number, err
}
