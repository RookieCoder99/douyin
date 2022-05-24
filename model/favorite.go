package model

import "time"

type TFavorite struct {
	ID        int64 `gorm:"primaryKey autoIncrement"`
	VideoID   int64
	UserID    int64
	CreatedAt time.Time
}
