# admins who has highest authority
CREATE TABLE IF NOT EXISTS `administrator` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(200) NOT NULL COMMENT '姓名',
    `email` varchar(200) NOT NULL COMMENT '邮箱',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;