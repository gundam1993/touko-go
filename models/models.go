package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;default:null"`
	Password string `gorm:"default:null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string
	Pv      int `gorm:"default:0"`
	UserID  int
	Display bool `gorm:"default:True"`
}

type Introduction struct {
	gorm.Model
	Content string
	UserID  int
}

var DB *gorm.DB

func InitDB(DBPath string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", DBPath)
	if err == nil {
		DB = db.AutoMigrate(&User{}, &Post{}, &Introduction{})
		return db, err
	}
	return nil, err
}
