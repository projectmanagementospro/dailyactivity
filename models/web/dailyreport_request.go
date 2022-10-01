package web

import (
	"time"
)

type DailyReportRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	User_id     uint64    `json:"user_id"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
}

type DailyReportUpdateRequest struct {
	ID           uint      `json:"id" binding:"required"`
	ReportNumber uint64    `json:"report_number"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description"`
	User_id      uint64    `json:"user_id"`
	Location     string    `json:"location" binding:"required"`
	Date         time.Time `json:"date"`
	StartTime    time.Time `json:"start_time" binding:"required"`
	EndTime      time.Time `json:"end_time" binding:"required"`
	UpdatedBy    string    `json:"updated_by" binding:"required"`
	DeletedBy    string    `json:"deleted_by"`
}
