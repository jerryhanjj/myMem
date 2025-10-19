package models

import (
	"time"
)

// TimeRecord 时间记录模型
type TimeRecord struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null" binding:"required"`
	Description string    `json:"description"`
	TargetDate  time.Time `json:"target_date" gorm:"not null" binding:"required"`
	RecordType  string    `json:"record_type" gorm:"not null" binding:"required"` // "countdown" 或 "elapsed"
	Category    string    `json:"category"`
	Color       string    `json:"color" gorm:"default:'#409EFF'"` // 默认 Element Plus 主色
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TimeRecordCreate 创建时间记录的请求结构
type TimeRecordCreate struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	TargetDate  time.Time `json:"target_date" binding:"required"`
	RecordType  string    `json:"record_type" binding:"required,oneof=countdown elapsed"`
	Category    string    `json:"category"`
	Color       string    `json:"color"`
}

// TimeRecordUpdate 更新时间记录的请求结构
type TimeRecordUpdate struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TargetDate  time.Time `json:"target_date"`
	RecordType  string    `json:"record_type" binding:"omitempty,oneof=countdown elapsed"`
	Category    string    `json:"category"`
	Color       string    `json:"color"`
}
