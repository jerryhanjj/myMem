package main

import (
	"log"
	"mymem-backend/database"
	"mymem-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	database.InitDatabase()

	// 创建 Gin 路由
	r := gin.Default()

	// 配置 CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config))

	// API 路由组
	api := r.Group("/api")
	{
		records := api.Group("/records")
		{
			records.GET("", handlers.GetAllRecords)       // 获取所有记录
			records.GET("/:id", handlers.GetRecord)       // 获取单个记录
			records.POST("", handlers.CreateRecord)       // 创建记录
			records.PUT("/:id", handlers.UpdateRecord)    // 更新记录
			records.DELETE("/:id", handlers.DeleteRecord) // 删除记录
		}

		// 统计相关路由
		api.GET("/statistics", handlers.GetStatistics)                 // 获取统计数据
		api.POST("/records/batch-delete", handlers.BatchDeleteRecords) // 批量删除
		api.DELETE("/records/clear", handlers.ClearAllRecords)         // 清空所有记录
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "MyMem API is running",
		})
	})

	// 启动服务器
	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
