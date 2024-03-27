package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

var AutoCodeMssql = new(autoCodeMssql)

type autoCodeMssql struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := "select name AS 'database' from sysdatabases;"
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table

	sql := fmt.Sprintf(`select name as 'table_name' from [%s].dbo.sysobjects where xtype='U'`, dbName)
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	//sql := fmt.Sprintf(`
	//	SELECT
	//		sc.name AS column_name,
	//		st.name AS data_type,
	//		sc.length AS data_type_long,
	//		CASE
	//			WHEN pk.index_id IS NOT NULL THEN 1
	//			ELSE 0
	//		END AS primary_key
	//	FROM
	//		[%s].dbo.syscolumns sc
	//	JOIN
	//		systypes st ON sc.xtype=st.xtype
	//	LEFT JOIN
	//		[%s].dbo.sysobjects so ON so.name='%s' AND so.xtype='U'
	//	LEFT JOIN
	//		[%s].dbo.sysindexes si ON si.id = so.id AND si.indid < 2
	//	LEFT JOIN
	//		[%s].dbo.sysindexkeys sik ON sik.id = si.id AND sik.indid = si.indid AND sik.colid = sc.colid
	//	LEFT JOIN
	//		[%s].dbo.syskeyconstraints pk ON pk.constid = sik.constid
	//	WHERE
	//		st.usertype=0 AND sc.id = so.id
	//`, dbName, dbName, tableName, dbName, dbName, dbName)

	sql := fmt.Sprintf(`
		SELECT
	COLUMN_NAME AS 'column_name',
		DATA_TYPE AS 'data_type',
		CHARACTER_MAXIMUM_LENGTH AS 'data_type_long',
		CASE
	WHEN COLUMN_NAME IN (
		SELECT COLUMN_NAME
	FROM  [%s].INFORMATION_SCHEMA.KEY_COLUMN_USAGE
	WHERE OBJECTPROPERTY(OBJECT_ID(CONSTRAINT_SCHEMA + '.' + CONSTRAINT_NAME), 'IsPrimaryKey') = 1
	AND TABLE_NAME = '%s'
	)
	THEN 1
	ELSE 0
	END AS 'primary_key'
	FROM [%s].INFORMATION_SCHEMA.COLUMNS
	WHERE TABLE_NAME = '%s';
	`, dbName, tableName, dbName, tableName)

	//sql := fmt.Sprintf(`
	//SELECT
	//  --表名=CASE WHEN C.column_id=1 THEN O.name ELSE N'' END,
	//  --表说明=ISNULL(CASE WHEN C.column_id=1 THEN PTB.[value] END,N''),
	//  --字段序号=C.column_id,
	//  column_name=C.name,
	//
	//  data_type=T.name,
	//      IDX.is_primary_key,
	//  --data_type_long=C.max_length,
	//          CASE
	//      WHEN IDX.is_primary_key =1 THEN 1
	//      ELSE 0
	//  END AS primary_key
	//
	//FROM sys.columns C
	//  INNER JOIN sys.objects O
	//      ON C.[object_id]=O.[object_id]
	//          AND O.type='U'
	//          AND O.is_ms_shipped=0
	//  INNER JOIN sys.types T
	//      ON C.user_type_id=T.user_type_id
	//  LEFT JOIN sys.default_constraints D
	//      ON C.[object_id]=D.parent_object_id
	//          AND C.column_id=D.parent_column_id
	//          AND C.default_object_id=D.[object_id]
	//LEFT JOIN sys.extended_properties PFD
	//      ON PFD.class=1
	//          AND C.[object_id]=PFD.major_id
	//          AND C.column_id=PFD.minor_id
	//--             AND PFD.name='Caption'  -- 字段说明对应的描述名称(一个字段可以添加多个不同name的描述)
	//  LEFT JOIN sys.extended_properties PTB
	//      ON PTB.class=1
	//          AND PTB.minor_id=0
	//          AND C.[object_id]=PTB.major_id
	//--             AND PFD.name='Caption'  -- 表说明对应的描述名称(一个表可以添加多个不同name的描述)
	//  LEFT JOIN                       -- 索引及主键信息
	//  (
	//      SELECT
	//          IDXC.[object_id],
	//          IDXC.column_id,
	//          Sort=CASE INDEXKEY_PROPERTY(IDXC.[object_id],IDXC.index_id,IDXC.index_column_id,'IsDescending')
	//              WHEN 1 THEN 'DESC' WHEN 0 THEN 'ASC' ELSE '' END,
	//          --PrimaryKey=CASE WHEN IDX.is_primary_key=1 THEN N'√'ELSE N'' END,
	//                      IDX.is_primary_key,
	//          IndexName=IDX.Name
	//      FROM sys.indexes IDX
	//      INNER JOIN sys.index_columns IDXC
	//ON IDX.[object_id]=IDXC.[object_id]
	//              AND IDX.index_id=IDXC.index_id
	//      LEFT JOIN sys.key_constraints KC
	//          ON IDX.[object_id]=KC.[parent_object_id]
	//              AND IDX.index_id=KC.unique_index_id
	//      INNER JOIN  -- 对于一个列包含多个索引的情况,只显示第1个索引信息
	//      (
	//          SELECT [object_id], Column_id, index_id=MIN(index_id)
	//          FROM sys.index_columns
	//          GROUP BY [object_id], Column_id
	//      ) IDXCUQ
	//          ON IDXC.[object_id]=IDXCUQ.[object_id]
	//              AND IDXC.Column_id=IDXCUQ.Column_id
	//              AND IDXC.index_id=IDXCUQ.index_id
	//  ) IDX
	//      ON C.[object_id]=IDX.[object_id]
	//          AND C.column_id=IDX.column_id
	//WHERE O.name = '%s'
	//  -- 如果只查询指定表,加上此条件
	//ORDER BY O.name,C.column_id
	//
	//`, tableName)

	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}
