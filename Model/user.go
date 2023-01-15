package Model

//具体的进行表字段的存储

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name,omitempty"`
	FollowCount   uint   `json:"follow_count,omitempty"`   //关注数目
	FollowerCount uint   `json:"follower_count,omitempty"` //粉丝数目
	IsFollow      bool   `json:"is_follow,omitempty"`      //是否关注了
	Password      string `json:"password"`
}
