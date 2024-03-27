package system

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/sqlserver"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/gookit/color"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type MssqlInitHandler struct{}

func NewMssqlInitHandler() *MssqlInitHandler {
	return &MssqlInitHandler{}
}

// WriteConfig pgsql 回写配置
func (h MssqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mssql)
	if !ok {
		return errors.New("mssql config invalid")
	}
	global.GVA_CONFIG.System.DbType = "mssql"
	global.GVA_CONFIG.Mssql = c
	global.GVA_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	return global.GVA_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 pg
func (h MssqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mssql" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToMssqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.MssqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE %s;", c.Dbname)
	if err = createDatabase(dsn, "sqlserver", createSql); err != nil {
		return nil, err
	} // 创建数据库

	var db *gorm.DB
	//sqlserver.Open(dsn), &gorm.Config{}
	if db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h MssqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h MssqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for i := 0; i < len(inits); i++ {
		if inits[i].DataInserted(next) {
			color.Info.Printf(InitDataExist, Mssql, inits[i].InitializerName())
			continue
		}
		if n, err := inits[i].InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Mssql, inits[i].InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Mssql, inits[i].InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Mssql)
	return nil
}
