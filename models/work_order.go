package models

import "time"

type WorkOrderStatus string

const (
	StatusOpen WorkOrderStatus = "OPEN"
	StatusInProgress WorkOrderStatus = "IN_PROGRESS"
	StatusCompleted WorkOrderStatus = "COMPLETED"
	StatusCancelled WorkOrderStatus = "CANCELLED"
)

type WorkOrder struct {
	ID uint 'gorm:"primaryKey" json:"id"'
	Title string 'gorm:"type:varchar(55);not null" json:"title"'
	Description *string 'gorm:"type:text" json:"description"'
	Status WorkOrderStatus 'gorm"type:enum('OPEN','IN_PROGRESS','COMPLETED','CANCELLED');default:'OPEN'" json:"status"'
	CreatedAt time.Time 'json:"created_at"'
	UpdatedAt time.Time 'json:"updated_at"'
}