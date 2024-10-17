package utils

import (
	"context"
	"fmt"
	"github.com/flyerxp/lib/v2/logger"
	"github.com/flyerxp/lib/v2/middleware/mysqlL"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SchemaMeta struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default interface{}
	Extra   string
}

func GetString(dbName string) {
	ctx := logger.GetContext(context.Background(), "cron")
	dbClient, err := mysqlL.GetEngine(ctx, dbName)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := dbClient.GetDb()
	tables := getTables(db)
	for _, table := range tables {
		metas := getTableInfo(table, db)
		result := changeMetas(table, metas)
		fmt.Println(result)
	}
}
func getTables(db *sqlx.DB) []string {
	var tables []string
	res, err := db.Query("SHOW TABLES")
	if err != nil {
		return []string{}
	}
	for res.Next() {
		var table string
		res.Scan(&table)
		tables = append(tables, table)
	}
	return tables
}
func getTableInfo(tableName string, db *sqlx.DB) (metas []SchemaMeta) {
	list, _ := db.Query(fmt.Sprintf("show columns from %s", tableName))
	for list.Next() {
		var data SchemaMeta
		err := list.Scan(&data.Field, &data.Type, &data.Null, &data.Key, &data.Default, &data.Extra)
		if err != nil {
			fmt.Println(err.Error())
		}
		metas = append(metas, data)
	}
	return metas
}
func changeMetas(tableName string, metas []SchemaMeta) string {
	var modelStr string
	for _, val := range metas {
		dataType := "interface{}"
		if val.Type[:3] == "int" {
			dataType = "int `json:\"" + val.Field + "\"` db:\"" + val.Field + "\"` "
		} else if val.Type[:4] == "enum" {
			dataType = "int `json:\"" + val.Field + "\"` db:\"" + val.Field + "\"` "
		} else if val.Type[:4] == "text" {
			dataType = "string `json:\"" + val.Field + "\"` db:\"" + val.Field + "\"` "
		} else if val.Type[:7] == "varchar" {
			dataType = "string `json:\"" + val.Field + "\"` db:\"" + val.Field + "\"` "
		} else if val.Type[:7] == "tinyint" {
			dataType = "bool `json:\"" + val.Field + "\"` db:\"" + val.Field + "\"` "
		} else if val.Type == "datetime" || val.Type == "timestamp" {
			dataType = "time.Time `json:\"" + val.Field + "\"` db:\"" + val.Field + "\"` "
		}
		field := val.Field
		field = Tf(field)
		modelStr += fmt.Sprintf("   %s %s\n", field, dataType)
	}
	tableName = strings.ToUpper(tableName[:1]) + tableName[1:]
	return fmt.Sprintf("type %s struct {\n %s }", tableName, modelStr)
}
func Tf(s string) string {
	as := strings.Split(s, "_")
	if len(as) == 1 {
		return strings.ToUpper(as[0][:1]) + as[0][1:]
	}
	for i, v := range as {
		as[i] = strings.ToUpper(v[:1]) + v[1:]
	}
	return strings.Join(as, "")
}
