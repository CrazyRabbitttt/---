package service

import (
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"errors"
	"gorm.io/gorm"
)

func IsFollowing(HostId uint, GuestId uint) bool {
	var tmpReleation = &Model.Following{}
	db := ConnSql.ThemodelOfSql()

	err := db.Table("favorite").Where("host_id = ? AND guest_id = ?", HostId, GuestId).First(&tmpReleation)
	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}
