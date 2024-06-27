package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	"gorm.io/gorm"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(alpha.ProductLine{}, alpha.TecBaseCraftsmanship{})
}
