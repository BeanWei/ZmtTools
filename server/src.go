package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Storage() {

	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./src.db")
	if err != nil {
		log.Fatal("Line15-Error: ", err)
	}

	//创建表
	sql_table := `
	 CREATE TABLE IF NOT EXISTS newsent(
			 uid INTEGER PRIMARY KEY AUTOINCREMENT,
			 title TEXT NULL,
			 author	TEXT NULL,
			 authorid	TEXT NULL,
			 publishtime DATE NULL,
			 views	INT	NULL,
			 comments	INT	NULL,
			 url	TEXT	NULL
	 );
	 `

	db.Exec(sql_table)
}
