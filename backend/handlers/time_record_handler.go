package handlers

import (
	"mymem-backend/database"
	"mymem-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllRecords 获取所有时间记录
func GetAllRecords(c *gin.Context) {
	var records []models.TimeRecord
	result := database.GetDB().Order("created_at DESC").Find(&records)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    records,
		"message": "success",
	})
}

// GetRecord 获取单个时间记录
func GetRecord(c *gin.Context) {
	id := c.Param("id")
	var record models.TimeRecord

	result := database.GetDB().First(&record, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    record,
		"message": "success",
	})
}

// CreateRecord 创建时间记录
func CreateRecord(c *gin.Context) {
	var input models.TimeRecordCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.TimeRecord{
		Title:       input.Title,
		Description: input.Description,
		TargetDate:  input.TargetDate,
		RecordType:  input.RecordType,
		Category:    input.Category,
		Color:       input.Color,
	}

	// 设置默认颜色
	if record.Color == "" {
		record.Color = "#409EFF"
	}

	result := database.GetDB().Create(&record)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"data":    record,
		"message": "Record created successfully",
	})
}

// UpdateRecord 更新时间记录
func UpdateRecord(c *gin.Context) {
	id := c.Param("id")
	var record models.TimeRecord

	if err := database.GetDB().First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var input models.TimeRecordUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新字段
	updates := make(map[string]interface{})
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if !input.TargetDate.IsZero() {
		updates["target_date"] = input.TargetDate
	}
	if input.RecordType != "" {
		updates["record_type"] = input.RecordType
	}
	if input.Category != "" {
		updates["category"] = input.Category
	}
	if input.Color != "" {
		updates["color"] = input.Color
	}

	if err := database.GetDB().Model(&record).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 重新查询更新后的记录
	database.GetDB().First(&record, id)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    record,
		"message": "Record updated successfully",
	})
}

// DeleteRecord 删除时间记录
func DeleteRecord(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result := database.GetDB().Delete(&models.TimeRecord{}, idInt)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Record deleted successfully",
	})
}
