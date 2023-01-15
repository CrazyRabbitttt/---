package controller

//用户登陆的时候的结构体，需要添加上token，同时添加好Name & Id

type UserLoginInfo struct {
	UserId int64  `json:"userId"` //进行索引是哪一个User
	Name   string `json:"name,omitempty"`
	Token  string `json:"token"`
}

//基本的Response类型

type Response struct {
	Statuscode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`   //关注数目
	FollowerCount int64  `json:"follower_count,omitempty"` //粉丝数目
	IsFollow      bool   `json:"is_follow,omitempty"`      //是否关注了
	Password      string `json:"password"`
}
