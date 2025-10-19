package handlers

import (
	"mymem-backend/database"
	"mymem-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatisticsResponse 统计数据响应
type StatisticsResponse struct {
	TotalRecords     int64               `json:"total_records"`
	CountdownRecords int64               `json:"countdown_records"`
	ElapsedRecords   int64               `json:"elapsed_records"`
	CategoryStats    []CategoryStat      `json:"category_stats"`
	RecentRecords    []models.TimeRecord `json:"recent_records"`
}

// CategoryStat 分类统计
type CategoryStat struct {
	Category   string  `json:"category"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}

// GetStatistics 获取统计数据
func GetStatistics(c *gin.Context) {
	db := database.GetDB()

	// 总记录数
	var totalRecords int64
	db.Model(&models.TimeRecord{}).Count(&totalRecords)

	// 倒计时记录数
	var countdownRecords int64
	db.Model(&models.TimeRecord{}).Where("record_type = ?", "countdown").Count(&countdownRecords)

	// 累计天数记录数
	var elapsedRecords int64
	db.Model(&models.TimeRecord{}).Where("record_type = ?", "elapsed").Count(&elapsedRecords)

	// 分类统计
	type CategoryResult struct {
		Category string
		Count    int64
	}
	var categoryResults []CategoryResult
	db.Model(&models.TimeRecord{}).
		Select("COALESCE(category, '未分类') as category, COUNT(*) as count").
		Group("category").
		Order("count DESC").
		Find(&categoryResults)

	// 计算百分比
	categoryStats := make([]CategoryStat, len(categoryResults))
	for i, result := range categoryResults {
		percentage := float64(0)
		if totalRecords > 0 {
			percentage = float64(result.Count) / float64(totalRecords) * 100
		}
		categoryStats[i] = CategoryStat{
			Category:   result.Category,
			Count:      result.Count,
			Percentage: percentage,
		}
	}

	// 最近创建的记录（前5条）
	var recentRecords []models.TimeRecord
	db.Order("created_at DESC").Limit(5).Find(&recentRecords)

	response := StatisticsResponse{
		TotalRecords:     totalRecords,
		CountdownRecords: countdownRecords,
		ElapsedRecords:   elapsedRecords,
		CategoryStats:    categoryStats,
		RecentRecords:    recentRecords,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    response,
		"message": "success",
	})
}

// BatchDeleteRecords 批量删除记录
func BatchDeleteRecords(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No IDs provided"})
		return
	}

	result := database.GetDB().Delete(&models.TimeRecord{}, ids)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          200,
		"message":       "Records deleted successfully",
		"deleted_count": result.RowsAffected,
	})
}

// ClearAllRecords 清空所有记录
func ClearAllRecords(c *gin.Context) {
	result := database.GetDB().Where("1 = 1").Delete(&models.TimeRecord{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          200,
		"message":       "All records cleared successfully",
		"deleted_count": result.RowsAffected,
	})
}
