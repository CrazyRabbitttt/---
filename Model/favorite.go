package Model

import "gorm.io/gorm"

type Favoriete struct {
	gorm.Model
	UserId  uint `json:"user_id"`
	VideoId uint `json:"video_id"`
	State   uint //用于进行是否关注的判断
}
