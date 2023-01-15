#### User表结构

> 存储基本的 **User**的信息，包括`id`,`name`等基本的字段，用户名密码组成的token在`UserLoginInfo`中存储
> `id`, `name`,`follow_count`,`follower_count`,`is_follow`

#### UserLoginInfo

> `id`, `token`, `name`



#### Register

> 用户的自增的`id`完全通过数据库进行创建，传入数据库的时候目前仅仅传入`name`，数据库表中自动填充`id`.
>
> 同时我们需要将 用户名➕密码组成的 `token`➕查到用户的`id`,传入`UserLoginInfo`表中 **用于登陆**

#### Login

> 因为我们在注册的时候已经将登陆信息同步到`UserLoginInfo`表中，直接查询表字段即可

#### Publish
> 进行文件的上传，目前仅仅是上传到./public目录下面，通过传递token进行User的查找
> 首先从`UserLoginInfo`中根据`token`找到`Id`,再次回表进行查询 `user`
> PostMan测试成功，在本地./public中上传