-- +migrate Up
create table `unko` (
    `id` int(11) unsigned not null auto_increment comment '唯一に定まるやーつ',
    `type` varchar(256) not null unique comment 'うんこの種類',
    `num` int(11) unsigned not null comment 'うんこの数',
    `created_at` datetime not null default current_timestamp comment '作成日時',
    `updated_at` datetime not null default current_timestamp on update current_timestamp comment '更新日時',
    `deleted_at` datetime comment '削除日時',
    primary key (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `unko`;
