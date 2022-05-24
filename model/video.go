package model

import "time"

type TVideo struct {
	ID        int64 `gorm:"primaryKey autoIncrement"`
	AuthorID  int64
	PlayUrl   string
	CoverUrl  string
	CreatedAt time.Time
	Title     string
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	User          User   `json:"author" gorm:"embedded"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title, omitempty"`
}
