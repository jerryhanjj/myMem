package database

import (
	"log"
	"mymem-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mymem.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库模式
	err = DB.AutoMigrate(&models.TimeRecord{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database initialized successfully")
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
