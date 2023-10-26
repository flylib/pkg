package mysql

import (
	"gorm.io/gorm"
)

type Tx struct {
	*gorm.DB
}

func (c *Tx) ExecProcedure(sql string, args map[string]interface{}) *gorm.DB {
	return execSql(pcdRegexp, c.DB, sql, args)
}

func (c *Tx) ExecRawSql(sql string, args map[string]interface{}) *gorm.DB {
	return execSql(rawRegexp, c.DB, sql, args)
}
