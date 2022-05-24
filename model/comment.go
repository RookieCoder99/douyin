package model

import "time"

type TComment struct {
	ID          int64 `gorm:"primaryKey autoIncrement"`
	VideoID     int64
	CommentText string
	UserID      int64
	CreatedAt   time.Time
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user" gorm:"embedded"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
