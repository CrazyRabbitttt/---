package Model

import "gorm.io/gorm"

//原始时候创建的表，不是基于struct创建的表
type Video1 struct {
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
}

type Video struct {
	gorm.Model
	AuthorId      uint   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string `json:"title"`
}
