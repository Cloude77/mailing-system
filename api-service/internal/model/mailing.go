package model

import (
	"time"
)

type Mailing struct {
	ID          int64     `json:"id"`
	Message     string    `json:"message"`
	GroupName   string    `json:"groupName"`
	SendTime    time.Time `json:"send_time"`
	Periodicity string    `json:"periodicity,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type MailingRequest struct {
	Message     string    `json:"message" binding:"required"`
	GroupName   string    `json:"group_name" binding:"required"`
	SendTime    time.Time `json:"send_time" binding:"required"`
	Periodicity string    `json:"periodicity,omitempty"`
}
