package xdb

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"

	"fx/pkg/logging"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Xorm struct {
	Db       *gorm.DB
	Category string
	Database string
	Limit    int
}

func (xm *Xorm) parseTable() (string, bool) {
	switch xm.Category {
	case "mysql":
		return fmt.Sprintf("SELECT table_name as Name FROM information_schema.tables WHERE table_schema='%s'", xm.Database), true
	case "postgres":
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	case "kingbase":
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	case "sqlite":
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	case "oracle":
		return "select table_name from sys.dba_tables where owner='schema名'", true
	case "sqlserver":
		return "SELECT table_name FROM information_schema.tables WHERE table_schema = 'mydatabasename' AND table_type = 'base table' ", true
	default:
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	}
}

func (xm *Xorm) parseSchemas(table string) (string, bool) {

	switch xm.Category {
	case "mysql":
		return fmt.Sprintf("SELECT COLUMN_NAME as Name FROM INFORMATION_SCHEMA.COLUMNS where table_schema='%s' AND table_name ='%s'", xm.Database, table), true
	case "postgres":
		return fmt.Sprintf("SELECT column_name as Name FROM INFORMATION_SCHEMA.COLUMNS where table_name ='%s'", table), true
	case "sqlite":
		return "", true
	case "oracle":
		return "", true
	case "sqlserver":
		return "", true
	default:
		return "", false
	}
}

func (xm *Xorm) parseData(table string) (string, bool) {
	switch xm.Category {
	case "mysql":
		if xm.Limit == 0 {
			return fmt.Sprintf("SELECT * FROM %s", table), true
		} else {
			return fmt.Sprintf("SELECT * FROM %s limit %s", table, xm.Limit), true
		}
	case "postgres":
		if xm.Limit == 0 {
			return fmt.Sprintf("SELECT * FROM %s", table), true
		} else {
			return fmt.Sprintf("SELECT * FROM %s limit %s", table, xm.Limit), true
		}
	case "sqlite":
		if xm.Limit == 0 {
			return fmt.Sprintf("SELECT * FROM %s", table), true
		} else {
			return fmt.Sprintf("SELECT * FROM %s limit %s", table, xm.Limit), true
		}
	case "oracle":
		if xm.Limit == 0 {
			return fmt.Sprintf("SELECT * FROM %s", table), true
		} else {
			return fmt.Sprintf("SELECT * FROM %s limit %s", table, xm.Limit), true
		}
	case "sqlserver":
		if xm.Limit == 0 {
			return fmt.Sprintf("SELECT * FROM %s", table), true
		} else {
			return fmt.Sprintf("SELECT * FROM %s limit %s", table, xm.Limit), true
		}
	default:
		if xm.Limit == 0 {
			return fmt.Sprintf("SELECT * FROM %s", table), true
		} else {
			return fmt.Sprintf("SELECT * FROM %s limit %s", table, xm.Limit), true
		}
	}
}

func (xm *Xorm) Tables() []string {
	var data []Tables
	results := []string{}
	if xsql, ok := xm.parseTable(); ok {
		xm.Db.Raw(xsql).First(&data)
		for _, table := range data {
			results = append(results, table.Name)
		}
	}
	return results
}

func (xm *Xorm) Xschemas(table string) []string {
	var data []Schemas
	results := []string{}
	if xsql, ok := xm.parseSchemas(table); ok {
		xm.Db.Raw(xsql).First(&data)
		for _, table := range data {
			results = append(results, table.Name)
		}
	}
	return results
}

func (xm *Xorm) parseCreateTable(table string) (string, bool) {
	switch xm.Category {
	case "mysql":
		return fmt.Sprintf("show create table %s", table), true
	case "postgres":
		return fmt.Sprintf(` SELECT 'CREATE TABLE ' || '%s' || ' (' || ' ' || '' || 
							    string_agg(column_list.column_expr, ', ' || ' ' || '') || 
							    '' || '' || ');'
							FROM (
							  SELECT '    ' || column_name || ' ' || data_type || 
							       coalesce('(' || character_maximum_length || ')', '') || 
							       case when is_nullable = 'YES' then '' else ' NOT NULL' end as column_expr
							  FROM information_schema.columns
							  WHERE table_schema = 'public' AND table_name = '%s'
							  ORDER BY ordinal_position ) column_list;`, table, table), true
	case "sqlite":
		return fmt.Sprintf("show create table %s", table), true
	case "oracle":
		return fmt.Sprintf("show create table %s", table), true
	case "sqlserver":
		return fmt.Sprintf("show create table %s", table), true
	default:
		return fmt.Sprintf("show create table %s", table), true
	}
}

func (xm *Xorm) CreateTableSql(table string) string {
	m := map[string]interface{}{}
	if xsql, ok := xm.parseCreateTable(table); ok {
		xm.Db.Raw(xsql).Scan(&m)
	}
	for _, v := range m {
		v2 := cast.ToString(v)
		if strings.Index(strings.ToLower(v2), "create table") >= 0 {
			return v2
		}
	}
	return ""
}

func (xm *Xorm) Data(table string) []map[string]interface{} {
	var results []map[string]interface{}
	if xsql, ok := xm.parseData(table); ok {
		xm.Db.Raw(xsql).Find(&results)
	}
	return results
}

func Run(config map[string]string) *Xorm {
	var DB *gorm.DB
	var err error
	var dsn string
	if config["category"] == "mysql" {
		dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s", config["username"], config["password"], config["network"], config["server"], config["port"], config["database"], config["charset"])
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if config["category"] == "postgres" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config["server"], config["username"], config["password"], config["database"], config["port"])
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		logging.Error(err.Error())
	}
	return &Xorm{Db: DB, Category: config["category"], Database: config["database"]}
}
