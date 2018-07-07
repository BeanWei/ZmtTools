package main

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	Filename	string
	dbh	*sql.DB
	mu  sync.RWMutex
}

func (s *Storage) Init() error {
	if s.dbh == nil {
		db, err := sql.Open("sqlite3", s.Filename)
		if err != nil {
			return err
		}
		err = db.Ping()
		if err != nil {
			return err
		}
		s.dbh = db
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
	statement, _ := s.dbh.Prepare(sql_table)
	_, err := statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Clear() error {
	statement, err := s.dbh.Prepare("DROP TABLE newsinfo")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Close() error {
	err := s.dbh.Close()
	return err
}

func (s *Storage) ClearOldNews() error {
	rows, err := s.dbh.Query("SELECT publishtime FROM newsinfo")
	if err != nil {
		return err
	}
	for rows.Next() {
		var publishtime string
		err = rows.Scan(&publishtime)
		if err != nil {
			return err
		}
		today := time.Now().Format("2006-01-02")
		if today != publishtime[0:10] {
			stmt, err := s.dbh.Prepare("DELETE FROM newsinfo WHERE publishtime = ?")
			if err != nil {
				return err
			}
			_, err = stmt.Exec(publishtime)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Storage) AddNews(field string, title string, author string, publishtime string, views int, comments int, url string, cover string) error {

	// 加锁
	s.mu.Lock()
	defer s.mu.Unlock()
	
	//插入数据
	statement, err := s.dbh.Prepare(
		"INSERT INTO newsinfo(field, title, author, publishtime, views, comments, url, cover) values(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		field,
		title,
		author,
		publishtime,
		views,
		comments,
		url,
		cover)
	if err != nil {
		return err
	}
	return nil
}

// Newsinfo 结构体
type Newsinfo struct {
	uid					int
	field       string
	title       string
	author      string
	publishtime string
	views       int
	comments    int
	url         string
	cover       string
}

// Newsquery 根据选择的领域进行数据库查询筛选
func (s *Storage) Newsquery(field string) ([]map[string]interface{}, error) {
	rows, err := s.dbh.Query("SELECT * FROM newsinfo WHERE field=?", field)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	allnews := []map[string]interface{}{}
	for rows.Next() {
		// n := new(Newsinfo)
		// err := rows.Scan(&n.uid, &n.field, &n.title, &n.author, &n.publishtime, &n.views, &n.comments, &n.url, &n.cover)
		var		uid					int
		var		field       string
		var		title       string
		var		author      string
		var		publishtime string
		var		views       int
		var		comments    int
		var		url         string
		var		cover       string
		err := rows.Scan(&uid, &field, &title, &author, &publishtime, &views, &comments, &url, &cover)
		if err != nil {
			return nil, err
		}
		newsmap := map[string]interface{}{
			"uid": uid,
			"field": field,
			"title": title,
			"author": author,
			"read":	views,
			"comment": comments,
			"publishtime": publishtime,
			"cover": cover,
			"url": url,
		}
		allnews = append(allnews, newsmap)
	}
	log.Println(allnews)
	return allnews, nil
} 

// Domains 查询数据库中的所有领域
func (s *Storage) Domains() ([]string, error) {
	rows, err := s.dbh.Query("SELECT field FROM newsinfo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := []string{}
	for rows.Next() {
		var field string
		err = rows.Scan(&field)
		if err != nil {
			return nil,err
		}
		results = append(results, field)
	}
	//数组去重
	return RemRepArry(results), nil
}

// ExistURL 判断是否已存在该新闻
func (s *Storage) ExistURL(url string) bool {
	var count int
	statement, err := s.dbh.Prepare("SELECT COUNT(*) FROM newsinfo where url = ?")
	if err != nil {
		 log.Printf("Query ExistURL failure: %s", err.Error())
		 return true
	}
	row := statement.QueryRow(url)
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
		return true
 	}
	if count >= 1 {
		return true
	}
	return false
}

