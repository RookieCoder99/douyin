package model

type TVideo struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

type Video struct {
	Id            int64
	Author        TUser
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

type FeedRequest struct {
	LatestTime int64 // 可选参数，限制返回视频的最新投稿时间戳，精确到s， 不填表示当前时间
}

type FeedResponse struct {
	StatusCode int32
	StatusMsg  string
	VideoList  Video
	NextTime   int64
}
