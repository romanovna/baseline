package db

import (
	"crud-api/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectSQL(host, port, uname, pass, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		uname, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Post{},
	)
}

func TestDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "password123", "localhost", "3306", "posts_test")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("DB error", err)
	}
	return db
}

func TruncateTable() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "password123", "localhost", "3306", "posts_test")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = db.Exec("TRUNCATE posts")
}
