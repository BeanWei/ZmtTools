package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func dbinit() *sql.DB {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./src.db")
	if err != nil {
		log.Fatal("Line14-Error: ", err)
	}

	//创建表
	sql_table := `
	 CREATE TABLE IF NOT EXISTS newsinfo(
			 uid INTEGER PRIMARY KEY AUTOINCREMENT,
			 field	TEXT	NULL,
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

	return db
}

func newsadd(db *sql.DB, field string, title string, author string, authorid string, publishtime string, views int, comments int, url string) error {

	//插入数据
	stmt, err := db.Prepare(
		"INSERT INTO newsinfo(field, title, author, authorid, publishtime, views, comments, url) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Line44-Error: ", err)
		return err
	}
	_, err = stmt.Exec(
		field,
		title,
		author,
		authorid,
		publishtime,
		views,
		comments,
		url)
	if err != nil {
		log.Fatal("Line56-Error: ", err)
		return err
	}

	return nil

}
