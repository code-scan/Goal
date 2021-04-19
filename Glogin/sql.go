package Glogin

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"strings"
)

func login(host, port, username, password, type_ string) (*gorm.DB, bool) {
	var connStr string
	switch type_ {
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=%ds", username, password, host, port, 10)
	case "postgres":
		connStr = fmt.Sprintf("host=%s port=%s user=%s dbname=test sslmode=disable password=%s timeout=%ds", host, port, username, password, 10)
	case "mssql":
		connStr = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=test;timeout=%ds", host, username, password, port, 10)
	}
	db, err := gorm.Open(type_, connStr)
	if err != nil {
		return nil, false
	}
	return db, true
}
func SqlLogin(host, port, username, password, type_ string) bool {
	db, isLogin := login(host, port, username, password, type_)
	if isLogin {
		defer db.Close()
		return true
	}
	return false
}
func SqlQuery(host, port, username, password, type_, sql string) []interface{} {
	var result []interface{}

	db, isLogin := login(host, port, username, password, type_)
	if isLogin == false {
		return result
	}
	isRaw := false
	raw := []string{
		"update ",
		"insert ",
		"delete ",
		"drop ",
		"create",
		"set ",
		"exec ",
		"execute ",
	}
	for _, r := range raw {
		if strings.Contains(sql, r) {
			isRaw = true
			break
		}
	}
	if isRaw {
		db.Exec(sql)
		return result
	}

	rows, err := db.Raw(sql).Rows()
	if err != nil {
		result = append(result, err.Error())
		return result
	}
	var colums []string
	//result := make(map[]interface{})
	for rows.Next() {

		if colums == nil {
			colums, _ = rows.Columns()
		}
		result = append(result, colums)
		columns := make([]interface{}, len(colums))
		columnPointers := make([]interface{}, len(colums))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		//m := make(map[string]interface{})
		m := make(map[string]interface{})
		for i, colName := range colums {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
			//等价
			//val = columns[i]
			//m[colName] = val
		}
		}
		log.Println([1])

		//log.Println(m)
		result = append(result, m)
		//result = append(result, row)
	}
	return result
	//log.Println(result)

}
