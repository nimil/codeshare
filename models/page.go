package models

import (
	"database/sql"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db *sql.DB
var dbPath string = "./data/travel_guide.db"

// SetDBPath 设置数据库文件路径
func SetDBPath(path string) {
	dbPath = path
}

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// 创建页面表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS pages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			path TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func GetPageByPath(path string) (*Page, error) {
	page := &Page{}
	err := db.QueryRow("SELECT id, title, content, path, created_at, updated_at FROM pages WHERE path = ?", path).
		Scan(&page.ID, &page.Title, &page.Content, &page.Path, &page.CreatedAt, &page.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return page, err
}

// PageExists 检查指定路径的页面是否存在
func PageExists(path string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pages WHERE path = ?)", path).Scan(&exists)
	return exists, err
}

func (p *Page) Save() error {
	// 检查页面是否已存在
	exists, err := PageExists(p.Path)
	if err != nil {
		return err
	}

	if !exists {
		// 插入新页面
		result, err := db.Exec(
			"INSERT INTO pages (title, content, path, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
			p.Title, p.Content, p.Path, time.Now(), time.Now(),
		)
		if err != nil {
			return err
		}
		p.ID, _ = result.LastInsertId()
	} else {
		// 更新现有页面
		_, err := db.Exec(
			"UPDATE pages SET title = ?, content = ?, updated_at = ? WHERE path = ?",
			p.Title, p.Content, time.Now(), p.Path,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllPages() ([]Page, error) {
	rows, err := db.Query("SELECT id, title, content, path, created_at, updated_at FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
		var page Page
		err := rows.Scan(&page.ID, &page.Title, &page.Content, &page.Path, &page.CreatedAt, &page.UpdatedAt)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
} 