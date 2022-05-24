package model

import "time"

type TRelation struct {
	ID         int64 `gorm:"primaryKey autoIncrement"`
	FollowerId int64
	FollowId   int64
	CreatedAt  time.Time
}
