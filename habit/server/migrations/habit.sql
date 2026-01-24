DROP TABLE IF EXISTS `admin_sys_config`;
CREATE TABLE `admin_sys_config` (
                                    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
                                    `config_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigName',
                                    `config_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigKey',
                                    `config_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigValue',
                                    `config_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigType',
                                    `is_frontend` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否前台',
                                    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark',
                                    `create_by` bigint DEFAULT NULL COMMENT '创建者',
                                    `update_by` bigint DEFAULT NULL COMMENT '更新者',
                                    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置管理';


DROP TABLE IF EXISTS `admin_sys_user`;
CREATE TABLE `admin_sys_user` (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编码',
                                  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户名',
                                  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '密码',
                                  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '昵称',
                                  `role` int DEFAULT 1 COMMENT '1:superadmin 2:user',
                                  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '加盐',
                                  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
                                  `status`  int DEFAULT 1 COMMENT '1:启用 2:禁用',
                                  `create_by` bigint DEFAULT NULL COMMENT '创建者',
                                  `update_by` bigint DEFAULT NULL COMMENT '更新者',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='系统用户管理';



DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户id',
                            `level_id` int NOT NULL DEFAULT '1' COMMENT '用户等级编号',
                            `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '账号名称/用户名',
                            `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
                            `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '头像路径',
                            `pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
                            `ref_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '推荐码',
                            `ref_id` int NOT NULL DEFAULT '0' COMMENT '推荐id',
                            `friend_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邀请码',
                            `friend_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邀请码',
                            `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态(1-正常 2-异常)',
                            `online_status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '在线状态(1-离线 2-在线)',
                            `register_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
                            `register_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '注册IP',
                            `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
                            `last_login_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '最后登录IP',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',

                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE KEY `uk_username` (`username`),
                            UNIQUE KEY `uk_ref_code` (`ref_code`),
                            KEY `idx_ref_id` (`ref_id`),
                            KEY `idx_status` (`status`),
                            KEY `idx_online_status` (`online_status`),
                            KEY `idx_register_at` (`register_at`),
                            KEY `idx_last_login_at` (`last_login_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户管理';

DROP TABLE IF EXISTS `app_user_wallet`;
CREATE TABLE `app_user_wallet` (
                            `user_id` bigint DEFAULT 0  COMMENT '用户id',
                            `pay_pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提现密码',
                            `pay_status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '提现状态(1-启用 2-禁用)',
                            `balance`decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '余额',
                            `frozen`decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '冻结金额',
                            `total_r`decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '总充值',
                            `total_w`decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '总提现',
                            `total_re`decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '打卡总收益',
                            `total_i`decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '邀请总收益',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                            UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户钱包';


DROP TABLE IF EXISTS `app_user_withdraw`;
CREATE TABLE `app_user_withdraw` (
                                      `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '提现订单ID',
                                      `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                                      `amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '提现金额',
                                      `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '提现地址',
                                      `apply_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '申请IP',
                                      `free` DECIMAL(30,2) NOT NULL DEFAULT 0.03 COMMENT '提现手续费',
                                      `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1待审核 2通过 3拒绝 4打款完成',
                                      `reject_reason` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '拒绝原因',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
                                      `reviewed_at` datetime DEFAULT NULL COMMENT '审核时间',
                                      `review_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '审核IP',
                                      `review_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '审核人ID',
                                      PRIMARY KEY (`id`),
                                      KEY `idx_user_id` (`user_id`),
                                      KEY `idx_status` (`status`),
                                      KEY `idx_created_at` (`created_at`),
                                      KEY `idx_review_id` (`review_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='提现申请表';





