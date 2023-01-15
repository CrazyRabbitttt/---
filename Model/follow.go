package Model

import "gorm.io/gorm"

type Following struct {
	gorm.Model
	HostId  uint //本机
	GuestId uint //客机
}

type Followers struct {
	gorm.Model
	HostId  uint
	GuestId uint
}
