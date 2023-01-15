package controller

import (
	"Web-Go/Common"
)

type FeedUser struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   uint   `json:"follow_count,omitempty"`
	FollowerCount uint   `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type FeedVideo struct {
	Id            uint     `json:"id,omitempty"` //可省略,视频的唯一标识
	Author        FeedUser `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`      //封面地址
	FavoriteCount int64    `json:"favorite_count,omitempty"` //点赞总数
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         int64    `json:"time_chuo"` //时间戳
}

type FeedResponse struct {
	Common.Response
	VideoList []FeedVideo `json:"video_list"`
	NextTime  int64       `json:"next_time"`
}

type FeedNoVideoResponse struct {
	Common.Response
	NextTime int64 `json:"next_time"`
}

//
////传入的参数：latest_time,token
//func Feed(c *gin.Context) { //现在返回的是StatusCode :0 & Unixcurtime :23563263
//
//	//进行VedioList的查找
//	c.JSON(http.StatusOK, FeedResponse{ //将给的结构体序列化成为Json， 传到response Body中
//		Response:  Common.Response{StatusCode: 0}, //成功就是返回0
//		VideoList: DemoVideos,                     //Video List
//		NextTime:  time.Now().Unix(),
//	})
//}
