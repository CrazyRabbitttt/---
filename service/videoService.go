package service

import (
	"Web-Go/ConnSql"
	"Web-Go/Model"
)

func GetVideoListById(userId uint) []Model.Video {
	db := ConnSql.ThemodelOfSql()
	var videoList []Model.Video
	db.Table("videos").Where("author_id = ?", userId).Find(&videoList)
	return videoList
}
