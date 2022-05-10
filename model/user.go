package model

type TUser struct {
	Id            int64
	Username      string
	Password      string
	Token         string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	StatusCode int32
	StatusMsg  string
	UserId     int64
	Token      string
}
