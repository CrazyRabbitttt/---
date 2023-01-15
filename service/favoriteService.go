package service

import (
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"errors"
	"gorm.io/gorm"
)

func CheckFavorite(userId uint, videoId uint) bool {
	var tmpCheckFavorite Model.Favoriete
	db := ConnSql.ThemodelOfSql()

	err := db.Table("favorites").Where("user_id = ? AND video_id = ?", userId, videoId).First(&tmpCheckFavorite)
	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return false
		}
	}

	if tmpCheckFavorite.State == 0 {
		return false
	}
	return true
}
