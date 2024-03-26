## mysql 操作

### 新建表

- 用户表

```sql
CREATE TABLE `blog_users`(
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`username` varchar(100) NOT NULL COMMENT '用户名称',
	`email` varchar(100) NOT NULL COMMENT '用户邮箱',
	`password` varchar(100) NOT NULL COMMENT '用户密码',
	`avater` varchar(100) DEFAULT '' COMMENT '用户头像地址',
	`desc` varchar(100) DEFAULT '' COMMENT '用户介绍',
	`state`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否禁用 0为禁用、1显示',
	`created_on`  int(10)  unsigned DEFAULT '0' COMMENT '创建时间',
	`modified_on`  int(10)  unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int(10)  unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户管理';
```

- 文章表

```sql
CREATE TABLE `blog_articles`(
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`title` varchar(100) DEFAULT '' COMMENT '文章名称',
	`desc` varchar(100) DEFAULT '' COMMENT '文章描述',
	`content` text COMMENT '文章内容',
	`cover_image_url` varchar(100) DEFAULT '' COMMENT '封面图片',
	`state`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否禁用 0为禁用、1显示',
	`is_draft`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否是草稿 0不是草稿、1为草稿',
	`created_on`  int(10)  unsigned DEFAULT '0' COMMENT '创建时间',	
	`modified_on`  int(10)  unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int(10)  unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',	
    `user_id` int(10) unsigned NOT NULL COMMENT '关联用户表的外键',
    FOREIGN KEY (user_id) REFERENCES blog_users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

```

**如果新建的时候没关联用户表 可以用下面的语句进行关联**

```sql
alter table blog_articles add user_id int(10) unsigned NOT NULL;
alter table blog_articles add constraint fk_user_id foreign key (user_id) references blog_users(id);
```

- 标签表

```sql
CREATE TABLE `blog_tag`(
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100) DEFAULT '' COMMENT '标签名称',
	`created_on`  int(10)  unsigned DEFAULT '0' COMMENT '创建时间',	
	`modified_on`  int(10)  unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int(10)  unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',	
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';

```


- 说说表

```sql
CREATE TABLE `blog_diaries`(
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`title` varchar(100) DEFAULT '' COMMENT '说说标题',
	`content` text COMMENT '说说内容',
	`created_on`  int(10)  unsigned DEFAULT '0' COMMENT '创建时间',	
	`modified_on`  int(10)  unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int(10)  unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',	
    `user_id` int(10) unsigned NOT NULL COMMENT '关联用户表的外键',
    FOREIGN KEY (user_id) REFERENCES blog_users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='说说管理';
```

- 标签图文表

```sql
CREATE TABLE `blog_tag_articles`(
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`tag_id` int(10) unsigned NOT NULL COMMENT '关联标签表的外键',
	`articles_id` int(10) unsigned NOT NULL COMMENT '关联文章表的外键',
	`created_on`  int(10)  unsigned DEFAULT '0' COMMENT '创建时间',	
	`modified_on`  int(10)  unsigned DEFAULT '0' COMMENT '修改时间',
	`deleted_on`  int(10)  unsigned DEFAULT '0' COMMENT '删除时间',
	`is_del`  tinyint(10)  unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',	
    FOREIGN KEY (tag_id) REFERENCES blog_tag(id),
    FOREIGN KEY (articles_id) REFERENCES blog_articles(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='说说管理';
```