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
			 publishtime DATE NULL,
			 views	INT	NULL,
			 comments	INT	NULL,
			 url	TEXT	NULL,
			 cover	TEXT	NULL
	 );
	 `

	db.Exec(sql_table)

	return db
}

func newsadd(db *sql.DB, field string, title string, author string, publishtime string, views int, comments int, url string, cover string) error {

	//插入数据
	stmt, err := db.Prepare(
		"INSERT INTO newsinfo(field, title, author, publishtime, views, comments, url, cover) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Line43-Error: ", err)
		return err
	}
	_, err = stmt.Exec(
		field,
		title,
		author,
		publishtime,
		views,
		comments,
		url,
		cover)
	if err != nil {
		log.Fatal("Line56-Error: ", err)
		return err
	}

	return nil

}

// Newsinfo 结构体
type Newsinfo struct {
	field       string
	title       string
	author      string
	publishtime string
	views       int
	comments    int
	url         string
	cover       string
}

// newsquery 根据选择的领域进行数据库查询筛选
func newsquery(field string) []*Newsinfo {
	db := dbinit()
	rows, err := db.Query("SELECT * FROM newsinfo WHERE field=?", field)
	if err != nil {
		log.Fatal("Line81-Error: ", err)
		return nil
	}
	defer rows.Close()

	allnews := []*Newsinfo{}
	for rows.Next() {
		n := new(Newsinfo)
		err := rows.Scan(&n.field, &n.title, &n.author, &n.publishtime, &n.views, &n.comments, &n.url, &n.cover)
		if err != nil {
			log.Fatal("Line91-Error: ", err)
			return nil
		}
		allnews = append(allnews, n)
	}
	return allnews
}

// domains 查询数据库中的所有领域
func domains() []string {
	db := dbinit()
	rows, err := db.Query("SELECT field FROM newsinfo")
	if err != nil {
		log.Fatal("Line104-Error: ", err)
		return nil
	}
	results := []string{}
	for rows.Next() {
		var field string
		err = rows.Scan(&field)
		if err != nil {
			log.Fatal("Line112-Error: ", err)
			return nil
		}
		results = append(results, field)
	}
	return results
}

//判断是否已存在该新闻
func exist(db *sql.DB, url string) bool {
	rows, err := db.Query("SELECT * FROM newsinfo WHERE url=?", url)
	if err != nil {
		log.Fatal("Line103-Error: ", err)
		return false
	}
	if rows != nil {
		return true
	} else {
		return false
	}
}

// newsdelete 清空数据库表 TODO: 暂支持手动清空
func newsdelete() error {
	db := dbinit()
	stmt, err := db.Prepare("DELETE FROM newsinfo")
	if err != nil {
		log.Fatal("Line104-Error: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("Line110-Error: ", err)
		return err
	}

	return nil
}
