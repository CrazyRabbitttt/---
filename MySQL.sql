
#### 创建表都是以tik_作为前缀进行表的设置
#用户表
CREATE TABLE `tik_user` (
                            `id` BIGINT not null auto_increment,		#自增创建ID
                            `name` 		VARCHAR(40) DEFAULT '',
                            `follow_count` int DEFAULT '0',
                            `follower_count` int DEFAULT '0',
                            `is_follow` bool DEFAULT '0',
                            `password` VARCHAR(40) DEFAULT '',
                            PRIMARY KEY(`id`)
)

    AuthorId      uint   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string `json:"title"`

create table `tik_video` (
        `video_id` int  not null auto_increment,
        `author`   int  not null default '0',
        `play_url`  varchar(60) default '',
        `cover_url` varchar(60) default '',
        `favorite_count` int not null default '0',
        `comment_count`  int not null default '0',
        `title`     varchar(50) default '',
        primary key (video_id)
)

