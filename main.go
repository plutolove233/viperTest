package main

import (
	"ViperTest/reader"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var config reader.Config

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", config.UserName, config.Password, config.IpAddrees, config.Port, config.DBName, config.Charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func addRecord(Db *sqlx.DB) {//添加
	result, err := Db.Exec("insert into user(userID,account,password,userName,isAdmin) values(?,?,?,?,?)", 67890, "shyhao", "666666", "Jackiex",1)
	if err != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("insert success, last id:[%d]\n", id)

}

func updateRecord(Db *sqlx.DB) {
	//更新uid=1的username
	result, err := Db.Exec("update user set username = 'anson' where userID = 1")
	if err != nil {
		fmt.Printf("update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

func deleteRecord(Db *sqlx.DB) {
	//删除uid=2的数据
	result, err := Db.Exec("delete from user where userID = 2")
	if err != nil {
		fmt.Printf("delete faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("delete success, affected rows:[%d]\n", num)
}

func main() {
	//import ini
	config.Read()

	var Db *sqlx.DB = connectMysql()
	defer Db.Close()

	addRecord(Db)
	//updateRecord(Db)
	//deleteRecord(Db)
}