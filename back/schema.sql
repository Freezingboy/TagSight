create database tag_sight;

use tag_sight;


-- 用户表
CREATE TABLE user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    username VARCHAR(50) NOT NULL COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码（加密存储）',
    PRIMARY KEY (id),
    UNIQUE KEY idx_username (username) COMMENT '用户名唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 标签表（关联用户，每个用户可创建自己的标签）
CREATE TABLE tag (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '所属用户ID',
    name VARCHAR(50) NOT NULL COMMENT '标签名（如 重要、工作）',
    color VARCHAR(20) COMMENT '标签颜色（如 #FF0000）',
    PRIMARY KEY (id),
    UNIQUE KEY idx_user_tag (user_id, name) COMMENT '同一用户的标签名唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 文件表（关联用户，记录文件归属）
CREATE TABLE file (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '所属用户ID',
    name VARCHAR(255) NOT NULL COMMENT '文件名',
    size INT NOT NULL COMMENT '文件大小（字节）',
    type VARCHAR(50) NOT NULL COMMENT '文件类型（如 pdf、jpg）',
    path VARCHAR(255) NOT NULL COMMENT '文件存储路径',
    upload_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    PRIMARY KEY (id),
    UNIQUE KEY idx_path (path) COMMENT '文件路径唯一',
    KEY idx_user (user_id) COMMENT '按用户查询文件'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 文件-标签关联表（关联用户，确保标签属于当前用户）
CREATE TABLE file_tag_relation (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关联ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '所属用户ID',
    file_id BIGINT UNSIGNED NOT NULL COMMENT '文件ID',
    tag_id BIGINT UNSIGNED NOT NULL COMMENT '标签ID',
    PRIMARY KEY (id),
    UNIQUE KEY idx_user_file_tag (user_id, file_id, tag_id) COMMENT '同一用户的文件-标签关联唯一',
    KEY idx_file (file_id),
    KEY idx_tag (tag_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
    