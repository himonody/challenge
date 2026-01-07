
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_user
-- ----------------------------
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user` (
                            `id` int NOT NULL AUTO_INCREMENT COMMENT '用户编码',
                            `level_id` int NOT NULL DEFAULT '1' COMMENT '用户等级编号',
                            `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '账号名称/用户名',
                            `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
                            `true_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '真实姓名',
                            `money` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '余额',
                            `freeze_money` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '冻结金额',
                            `email` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '电子邮箱',
                            `mobile_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '+86' COMMENT '用户手机号国家前缀',
                            `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号码',
                            `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '头像路径',
                            `pay_pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提现密码',
                            `pay_status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '提现状态(1-启用 2-禁用)',
                            `pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
                            `ref_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '推荐码',
                            `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级编号',
                            `parent_ids` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '所有父级编号',
                            `tree_sort` int NOT NULL DEFAULT '0' COMMENT '本级排序号（升序）',
                            `tree_sorts` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '所有级别排序号',
                            `tree_leaf` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '是否最末级',
                            `tree_level` int NOT NULL DEFAULT '0' COMMENT '层次级别',
                            `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态(1-正常 2-异常)',
                            `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',

                            `register_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
                            `register_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '注册IP',
                            `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
                            `last_login_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '最后登录IP',

                            `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
                            `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',

                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE KEY `uk_username` (`username`),
                            UNIQUE KEY `uk_email` (`email`),
                            UNIQUE KEY `uk_mobile` (`mobile_title`,`mobile`),
                            UNIQUE KEY `uk_ref_code` (`ref_code`),
                            KEY `idx_parent_id` (`parent_id`),
                            KEY `idx_status` (`status`),
                            KEY `idx_register_at` (`register_at`),
                            KEY `idx_last_login_at` (`last_login_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户管理';


-- ----------------------------
-- Records of app_user
-- ----------------------------
BEGIN;

INSERT INTO `app_user` (`id`, `level_id`, `username`, `nickname`, `true_name`,
                        `money`, `freeze_money`,
                        `email`, `mobile_title`, `mobile`, `avatar`,
                        `pay_pwd`, `pay_status`, `pwd`,
                        `ref_code`,
                        `parent_id`, `parent_ids`,
                        `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`,
                        `status`, `remark`,
                        `register_at`, `register_ip`,
                        `last_login_at`, `last_login_ip`,
                        `create_by`, `update_by`,
                        `created_at`, `updated_at`)
VALUES (1, 1, '- -', '- -', '- -',
        1.00, 0.00,
        'fb0cc809bbed1743bd7d2d8f444e2bae099e69819f4e072f7057bb1e4249bf3d',
        '86', '6d84b6afd68a5c7188779114f16c46e9',
        'http://www.bitxx.top/images/my_head-touch-icon-next.png',
        '', '1', '',
        'akIiWm',
        0, '0,',
        1, '1,', '2', 1,
        '1', '',
        '2023-04-03 21:09:13', NULL,
        NULL, NULL,
        0, 1,
        '2023-04-03 21:09:13', '2023-10-19 14:03:37'),
       (2, 2, '- -', '- -', '- -',
        0.00, 0.00,
        'dca887a13d1225ccd447dc52a712861c099e69819f4e072f7057bb1e4249bf3d',
        '86', '84ace68f39f53a315d8114c61413505d',
        'http://www.bitxx.top/images/my_head-touch-icon-next.png',
        '', '1', '',
        'GQFz6v',
        1, '0,1,',
        1, '1,1,', '1', 2,
        '1', '',
        '2023-04-03 21:29:34', NULL,
        NULL, NULL,
        0, 1,
        '2023-04-03 21:29:34', '2023-10-19 14:06:49'),
       (3, 1, '- -', '- -', '- -',
        0.00, 0.00,
        '4884f3537b62e668d33c6af76ddf6670099e69819f4e072f7057bb1e4249bf3d',
        '86', 'ff4273c3b1372055923122f9881b651b',
        'http://www.bitxx.top/images/my_head-touch-icon-next.png',
        '', '1', '',
        'tT1Fbk',
        1, '0,1,',
        2, '1,2,', '1', 2,
        '1', '',
        '2023-04-03 21:29:35', NULL,
        NULL, NULL,
        0, 1,
        '2023-04-03 21:29:35', '2023-10-19 14:06:37');

COMMIT;

-- ----------------------------
-- Table structure for app_user_login_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_login_log`;
CREATE TABLE `app_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '登录日志ID',
  `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
  `login_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  `login_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '登录IP',
  `device_fp` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '设备指纹',
  `user_agent` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'UA信息',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '登录状态 1成功 2失败 3风控拦截',
  `fail_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '失败原因/拦截原因',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间戳',
  PRIMARY KEY (`id`),
  KEY `idx_user_time` (`user_id`,`login_at`),
  KEY `idx_status` (`status`),
  KEY `idx_ip` (`login_ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户登录日志';



-- ----------------------------
-- Table structure for app_user_account_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_account_log`;
CREATE TABLE `app_user_account_log` (
                                        `id` int NOT NULL AUTO_INCREMENT COMMENT '账变编号',
                                        `user_id` int NOT NULL DEFAULT '0' COMMENT '用户编号',
                                        `change_money` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '账变金额',
                                        `before_money` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '账变前金额',
                                        `after_money` decimal(30,2) NOT NULL DEFAULT '0.00' COMMENT '账变后金额',
                                        `money_type` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '金额类型 1:余额（可扩展）',
                                        `change_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '账变类型编码',
                                        `operate_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '操作IP',
                                        `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态（1正常 2-异常）',
                                        `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
                                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                        `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
                                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                                        `remarks` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',

                                        PRIMARY KEY (`id`),
                                        KEY `idx_user_id_created` (`user_id`,`created_at`),
                                        KEY `idx_change_type` (`change_type`),
                                        KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='账变记录';

-- ----------------------------
-- Records of app_user_account_log
-- ----------------------------
BEGIN;

INSERT INTO `app_user_account_log` (
    `id`, `user_id`,
    `change_money`, `before_money`, `after_money`,
    `money_type`, `change_type`,
    `operate_ip`,
    `status`,
    `create_by`, `created_at`,
    `update_by`, `updated_at`,
    `remarks`
) VALUES
      (1,1,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (2,2,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (3,1,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (4,3,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (5,1,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (6,2,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (7,1,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (8,3,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL),
      (9,1,10.00,0.00,20.00,'1','1',NULL,'1',1,'2023-03-09 22:55:48',1,'2023-03-09 22:55:51',NULL);

COMMIT;


-- ----------------------------
-- Table structure for app_user_conf
-- ----------------------------
DROP TABLE IF EXISTS `app_user_conf`;
CREATE TABLE `app_user_conf` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
                                 `can_login` char(1) NOT NULL DEFAULT '0' COMMENT '1-允许登陆；2-不允许登陆',
                                 `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
                                 `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态（1-正常 2-异常）\n',
                                 `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
                                 `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE KEY `uk_user_id` (`user_id`),
                                 KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT COMMENT='用户配置';

-- ----------------------------
-- Records of app_user_conf
-- ----------------------------
BEGIN;
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, '1', '', '1', 198, 198, '2023-04-03 21:09:13', '2023-04-03 21:09:13');
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, '1', '', '1', 200, 200, '2023-04-03 21:29:34', '2023-04-03 21:29:34');
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 3, '1', '', '1', 201, 201, '2023-04-03 21:29:35', '2023-04-03 21:29:35');
COMMIT;

-- ----------------------------
-- Table structure for app_user_country_code
-- ----------------------------
DROP TABLE IF EXISTS `app_user_country_code`;
CREATE TABLE `app_user_country_code` (
                                         `id` int NOT NULL AUTO_INCREMENT,
                                         `country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家或地区',
                                         `code` varchar(12) NOT NULL DEFAULT '' COMMENT '区号',
                                         `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态(1-可用 2-停用)',
                                         `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
                                         `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
                                         `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
                                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`) USING BTREE,
                                         UNIQUE KEY `uk_code` (`code`),
                                         KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT COMMENT='国家区号';

-- ----------------------------
-- Records of app_user_country_code
-- ----------------------------
BEGIN;
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '新加坡', '65', '2', '', 1, 1, '2021-06-29 14:10:00', '2021-06-29 14:10:00');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '加拿大', '1', '2', '', 1, 1, '2021-06-29 14:10:21', '2021-06-29 14:10:21');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '韩国', '82', '2', '', 1, 1, '2021-06-29 14:10:36', '2021-06-29 14:10:36');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '日本', '81', '2', '', 1, 1, '2021-06-29 14:10:49', '2021-06-29 14:10:49');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '中国香港', '852', '2', '', 1, 1, '2021-06-29 14:11:02', '2021-06-29 14:11:02');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '中国澳门', '853', '2', '', 1, 1, '2021-06-29 14:11:15', '2021-06-29 14:11:15');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '中国台湾', '886', '2', '', 1, 1, '2021-06-29 14:11:25', '2021-06-29 14:11:25');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '泰国', '66', '2', '', 1, 1, '2021-06-29 14:11:36', '2021-06-29 14:11:36');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '缅甸', '95', '2', '', 1, 1, '2021-06-29 14:11:45', '2021-06-29 14:11:45');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '老挝', '856', '1', '', 1, 1, '2021-06-29 14:11:59', '2023-03-14 21:11:18');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '澳大利亚', '61', '2', '', 1, 1, '2021-06-29 14:12:14', '2021-06-29 14:12:14');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '俄罗斯', '7', '1', '', 1, 1, '2021-06-29 14:12:32', '2023-03-14 21:11:08');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '中国大陆', '86', '1', '', 1, 1, '2021-06-29 14:16:22', '2023-03-14 21:11:03');
COMMIT;

-- ----------------------------
-- Table structure for app_user_level
-- ----------------------------
DROP TABLE IF EXISTS `app_user_level`;
CREATE TABLE `app_user_level` (
                                  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                                  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '等级名称',
                                  `level_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '等级类型',
                                  `level` int NOT NULL DEFAULT '0' COMMENT '等级',
                                  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态(1-正常 2-异常)',
                                  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
                                  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
                                  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uk_level_type_level` (`level_type`,`level`),
                                  KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户等级';

-- ----------------------------
-- Records of app_user_level
-- ----------------------------
BEGIN;
INSERT INTO `app_user_level` (`id`, `name`, `level_type`, `level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 'test3', '2', 2, '1', '', 1, 1, '2023-03-09 17:05:24', '2023-03-09 17:05:24');
INSERT INTO `app_user_level` (`id`, `name`, `level_type`, `level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test34', '1', 1, '1', '', 1, 1, '2023-03-09 17:05:37', '2023-03-09 20:19:19');
COMMIT;

-- ----------------------------
-- Table structure for app_user_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_oper_log`;
CREATE TABLE `app_user_oper_log` (
                                     `id` int NOT NULL AUTO_INCREMENT COMMENT '日志编码',
                                     `user_id` int NOT NULL DEFAULT 0 COMMENT '用户编号',
                                     `action_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户行为类型编码',
                                     `operate_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '操作IP',
                                     `by_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '更新用户类型 1-app用户 2-后台用户',
                                     `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '状态(1-正常 2-异常)',
                                     `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
                                     `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
                                     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                                     `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息/原因',

                                     PRIMARY KEY (`id`) USING BTREE,
                                     KEY `idx_user_created` (`user_id`,`created_at`),
                                     KEY `idx_action_type` (`action_type`),
                                     KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户关键行为日志表';

-- ----------------------------
-- Records of app_user_oper_log
-- ----------------------------
BEGIN;

INSERT INTO `app_user_oper_log` (
    `id`, `user_id`,
    `action_type`, `operate_ip`,
    `by_type`, `status`,
    `create_by`, `update_by`,
    `created_at`, `updated_at`,
    `remark`
) VALUES
      (1,1,'',NULL,'2','1',1,1,'2023-03-11 15:39:31','2023-03-11 15:39:31',''),
      (2,2,'',NULL,'2','1',1,1,'2023-03-11 15:41:16','2023-03-11 15:41:16',''),
      (3,3,'',NULL,'1','1',1,1,'2023-03-11 15:45:44','2023-03-11 15:45:44',''),
      (4,1,'',NULL,'1','1',1,1,'2023-03-11 15:46:13','2023-03-11 15:46:13',''),
      (5,3,'2',NULL,'1','1',1,1,'2023-03-11 15:54:05','2023-03-11 15:54:05',''),
      (6,2,'1',NULL,'1','1',1,1,'2023-03-11 15:56:36','2023-03-11 15:56:36',''),
      (7,1,'2',NULL,'1','1',1,1,'2023-03-11 16:03:35','2023-03-11 16:03:35','');

COMMIT;


-- ----------------------------
-- Table structure for app_challenge_config
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_config`;
CREATE TABLE `app_challenge_config` (
                                    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '活动配置ID',
                                    `day_count` INT NOT NULL DEFAULT 0 COMMENT '挑战天数 1/7/21',
                                    `amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '单人挑战金额',
                                    `checkin_start` datetime  DEFAULT NULL COMMENT '每日打卡开始时间',
                                    `checkin_end` datetime  DEFAULT NULL COMMENT '每日打卡结束时间',
                                    `platform_bonus` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '平台补贴金额',
                                    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1启用 2停用',
                                    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `uk_day_amount` (`day_count`,`amount`),
                                    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin  COMMENT='打卡挑战活动配置';

-- ----------------------------
-- Table structure for app_challenge_checkin
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_checkin`;
CREATE TABLE `app_challenge_checkin` (
                                     `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '打卡ID',
                                     `challenge_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户挑战ID',
                                     `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',

                                     `checkin_date` DATE NOT NULL COMMENT '打卡日期 YYYYMMDD',
                                     `checkin_time` datetime  DEFAULT NULL COMMENT '打卡时间戳',

                                     `mood_code` TINYINT NOT NULL DEFAULT 0 COMMENT '心情枚举 1开心 2平静 3一般 4疲惫 5低落 6爆棚',
                                     `mood_text` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '用户心情文字描述（最多200字）',

                                     `content_type` TINYINT NOT NULL DEFAULT 1 COMMENT '打卡内容类型 1图片 2视频广告',


                                         `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1成功 2超时',
                                     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间戳',

                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `uk_challenge_date` (`challenge_id`,`checkin_date`),
                                     KEY `idx_user_date` (`user_id`,`checkin_date`),
                                     KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='每日打卡记录（心情 + 内容）';

-- ----------------------------
-- Table structure for app_challenge_checkin_image
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_checkin_image`;
CREATE TABLE `app_challenge_checkin_image` (
                                           `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '图片ID',
                                           `checkin_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '打卡ID',
                                           `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                                           `image_url` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '图片URL',
                                           `image_hash` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '图片Hash（防重复）',
                                           `sort_no` TINYINT NOT NULL DEFAULT 1 COMMENT '图片顺序',
                                           `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1正常 2屏蔽 3审核中',
                                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间戳',
                                           PRIMARY KEY (`id`),
                                           UNIQUE KEY `uk_checkin_hash` (`checkin_id`,`image_hash`),
                                           KEY `idx_checkin` (`checkin_id`),
                                           KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='打卡图片表';

-- ----------------------------
-- Table structure for app_challenge_checkin_video_ad
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_checkin_video_ad`;
CREATE TABLE `app_challenge_checkin_video_ad` (
                                              `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频广告打卡ID',

                                              `checkin_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联打卡ID',
                                              `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',

                                              `ad_platform` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '广告平台 如：csj、gdt、unity',
                                              `ad_unit_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '广告位ID',
                                              `ad_order_no` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '广告联盟返回的订单号（唯一）',

                                              `video_duration` INT NOT NULL DEFAULT 0 COMMENT '视频时长（秒）',
                                              `watch_duration` INT NOT NULL DEFAULT 0 COMMENT '实际观看时长（秒）',

                                              `reward_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '该广告产生的收益',
                                              `verify_status` TINYINT NOT NULL DEFAULT 0 COMMENT '校验状态 0待校验 1成功 2失败',

                                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '观看完成时间戳',
                                              `verified_at` datetime  DEFAULT NULL COMMENT '校验完成时间戳',

                                              PRIMARY KEY (`id`),
                                              UNIQUE KEY `uk_ad_order` (`ad_order_no`),
                                              UNIQUE KEY `uk_checkin` (`checkin_id`),
                                              KEY `idx_user` (`user_id`),
                                              KEY `idx_verify_status` (`verify_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='视频广告打卡记录';

-- ----------------------------
-- Table structure for app_challenge_user
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_user`;
CREATE TABLE `app_challenge_user` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户挑战ID',
                                  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                                  `config_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动配置ID',
                                  `pool_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '奖池ID',
                                  `challenge_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '用户挑战金额',
                                  `start_date` INT NOT NULL DEFAULT 0 COMMENT '活动开始日期 YYYYMMDD',
                                  `end_date` INT NOT NULL DEFAULT 0 COMMENT '活动结束日期 YYYYMMDD',
                                  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1进行中 2成功 3失败',
                                  `fail_reason` TINYINT NOT NULL DEFAULT 0 COMMENT '失败原因 0无 1未打卡 2作弊',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '报名时间戳',
                                  `finished_at` datetime  DEFAULT NULL COMMENT '完成时间戳',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uk_user_active` (`user_id`,`status`),
                                  KEY `idx_pool` (`pool_id`),
                                  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户参与挑战记录';

-- ----------------------------
-- Table structure for app_challenge_pool
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_pool`;
CREATE TABLE `app_challenge_pool` (
                                  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '奖池ID',
                                  `config_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '活动配置ID',
                                  `start_date` datetime  DEFAULT NULL COMMENT '活动开始日期',
                                  `end_date` datetime  DEFAULT NULL COMMENT '活动结束日期',
                                  `total_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '奖池当前总金额',
                                  `settled` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已结算 0否 1是',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `uk_config_date` (`config_id`,`start_date`),
                                  KEY `idx_settled` (`settled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='活动奖池表';

-- ----------------------------
-- Table structure for app_challenge_pool_flow
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_pool_flow`;
CREATE TABLE `app_challenge_pool_flow` (
                                       `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '奖池流水ID',
                                       `pool_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '奖池ID',
                                       `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                                       `amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '变动金额',
                                       `type` TINYINT NOT NULL DEFAULT 0 COMMENT '类型 1报名 2失败 3平台补贴 4结算',
                                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
                                       PRIMARY KEY (`id`),
                                       KEY `idx_pool` (`pool_id`),
                                       KEY `idx_user` (`user_id`),
                                       KEY `idx_type_time` (`type`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='奖池资金流水';

-- ----------------------------
-- Table structure for app_challenge_settlement
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_settlement`;
CREATE TABLE `app_challenge_settlement` (
                                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '结算ID',
                                        `challenge_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户挑战ID',
                                        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                                        `reward` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '最终获得金额',
                                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结算时间戳',
                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `uk_challenge_user` (`challenge_id`,`user_id`),
                                        KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='挑战结算结果';

-- ----------------------------
-- Table structure for app_challenge_daily_stat
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_daily_stat`;
CREATE TABLE `app_challenge_daily_stat` (
                                        `stat_date` DATE NOT NULL COMMENT '统计日期 YYYYMMDD',
                                        `join_user_cnt` INT NOT NULL DEFAULT 0 COMMENT '参与人数',
                                        `success_user_cnt` INT NOT NULL DEFAULT 0 COMMENT '成功人数',
                                        `fail_user_cnt` INT NOT NULL DEFAULT 0 COMMENT '失败人数',
                                        `join_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '参与总金额',
                                        `success_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '成功金额',
                                        `fail_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '失败金额',
                                        `platform_bonus` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '平台补贴',
                                        `pool_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '奖池金额',
                                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
                                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间戳',
                                        PRIMARY KEY (`stat_date`),
                                        KEY `idx_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='每日运营统计';

-- ----------------------------
-- Table structure for app_challenge_total_stat
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_total_stat`;
CREATE TABLE `app_challenge_total_stat` (
                                        `id` TINYINT NOT NULL DEFAULT 1 COMMENT '固定主键',
                                        `total_user_cnt` INT NOT NULL DEFAULT 0 COMMENT '累计用户数',
                                        `total_join_cnt` INT NOT NULL DEFAULT 0 COMMENT '累计参与人次',
                                        `total_success_cnt` INT NOT NULL DEFAULT 0 COMMENT '累计成功人次',
                                        `total_fail_cnt` INT NOT NULL DEFAULT 0 COMMENT '累计失败人次',
                                        `total_join_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '累计参与金额',
                                        `total_success_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '累计成功金额',
                                        `total_fail_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '累计失败金额',
                                        `total_platform_bonus` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '累计平台补贴',
                                        `total_pool_amount` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '累计奖池金额',
                                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间戳',
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='平台累计统计';

-- ----------------------------
-- Table structure for app_challenge_rank_daily
-- ----------------------------
DROP TABLE IF EXISTS `app_challenge_rank_daily`;
CREATE TABLE `app_challenge_rank_daily` (
                                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '排行ID',
                                        `rank_date` DATE NOT NULL COMMENT '排行日期',
                                        `rank_type` TINYINT NOT NULL DEFAULT 0 COMMENT '1邀请 2收益 3毅力',
                                        `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                                        `value` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '排行值',
                                        `rank_no` INT NOT NULL DEFAULT 0 COMMENT '排名',
                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `uk_rank` (`rank_date`,`rank_type`,`user_id`),
                                        KEY `idx_rank_type` (`rank_type`,`rank_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='排行榜日快照';

-- ----------------------------
-- Table structure for app_withdraw_order
-- ----------------------------
DROP TABLE IF EXISTS `app_withdraw_order`;
CREATE TABLE `app_withdraw_order` (
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

                                      PRIMARY KEY (`id`),
                                      KEY `idx_user_id` (`user_id`),
                                      KEY `idx_status` (`status`),
                                      KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='提现申请表';


-- ----------------------------
-- Table structure for app_risk_user
-- ----------------------------
DROP TABLE IF EXISTS `app_risk_user`;
CREATE TABLE `app_risk_user` (
                             `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                             `risk_level` TINYINT NOT NULL DEFAULT 0 COMMENT '0正常 1观察 2限制 3封禁',
                             `risk_score` INT NOT NULL DEFAULT 0 COMMENT '风险评分',
                             `reason` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '风险原因',
                             `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`user_id`),
                             KEY `idx_risk_level` (`risk_level`),
                             KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='风控用户表';

-- ----------------------------
-- Table structure for app_risk_device
-- ----------------------------
DROP TABLE IF EXISTS `app_risk_device`;
CREATE TABLE `app_risk_device` (
                               `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录ID',
                               `device_fp` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '设备指纹',
                               `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                               `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录时间',
                               PRIMARY KEY (`id`),
                               KEY `idx_fp` (`device_fp`),
                               KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备风控表';

-- ----------------------------
-- Table structure for app_risk_event
-- ----------------------------
DROP TABLE IF EXISTS `app_risk_event`;
CREATE TABLE `app_risk_event` (
                              `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '事件ID',
                              `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
                              `event_type` TINYINT NOT NULL DEFAULT 0 COMMENT '事件类型',
                              `detail` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '事件详情',
                              `score` INT NOT NULL DEFAULT 0 COMMENT '风险分',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发生时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_user` (`user_id`),
                              KEY `idx_event_type_time` (`event_type`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='风控事件记录';

DROP TABLE IF EXISTS `app_risk_rate_limit`;
CREATE TABLE `app_risk_rate_limit` (
                                       `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                       `scene` varchar(32) NOT NULL COMMENT '场景: register/login/sms/email',
                                       `identity_type` varchar(16) NOT NULL COMMENT '标识类型 ip/device/mobile/email',
                                       `identity_value` varchar(128) NOT NULL COMMENT '标识值',
                                       `count` int NOT NULL DEFAULT 1 COMMENT '次数',
                                       `window_start` datetime NOT NULL COMMENT '窗口开始时间',
                                       `window_end` datetime NOT NULL COMMENT '窗口结束时间',
                                       `blocked` char(1) NOT NULL DEFAULT '0' COMMENT '是否已拦截 0否 1是',
                                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                       PRIMARY KEY (`id`),
                                       UNIQUE KEY `uk_scene_identity_window`
                                           (`scene`,`identity_type`,`identity_value`,`window_start`),
                                       KEY `idx_scene_identity` (`scene`,`identity_type`,`identity_value`),
                                       KEY `idx_blocked` (`blocked`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='频率/防刷限制表';

DROP TABLE IF EXISTS `app_risk_blacklist`;
CREATE TABLE `app_risk_blacklist` (
                                      `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                      `type` varchar(16) NOT NULL COMMENT 'ip/device/country/mobile/email',
                                      `value` varchar(128) NOT NULL COMMENT '命中值',
                                      `risk_level` TINYINT NOT NULL DEFAULT 3 COMMENT '风险等级',
                                      `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '封禁原因',
                                      `status` char(1) NOT NULL DEFAULT '1' COMMENT '1生效 2失效',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `uk_type_value` (`type`,`value`),
                                      KEY `idx_status` (`status`),
                                      KEY `idx_type_status` (`type`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='风控黑名单';

DROP TABLE IF EXISTS `app_risk_appeal`;
CREATE TABLE `app_risk_appeal` (
                                   `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '申诉ID',

    -- 申诉主体
                                   `user_id` BIGINT UNSIGNED NOT NULL COMMENT '申诉用户ID',
                                   `risk_level` TINYINT NOT NULL DEFAULT 0 COMMENT '申诉时风险等级',
                                   `risk_reason` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '触发风控原因',

    -- 申诉内容
                                   `appeal_type` TINYINT NOT NULL DEFAULT 1 COMMENT '申诉类型 1账号封禁 2登录限制 3设备封禁',
                                   `appeal_reason` VARCHAR(500) NOT NULL COMMENT '用户申诉说明',
                                   `appeal_evidence` VARCHAR(1000) DEFAULT NULL COMMENT '申诉凭证(图片/链接)',

    -- 关联信息（关键）
                                   `ip` VARCHAR(45) DEFAULT NULL COMMENT '申诉时IP',
                                   `device_fp` VARCHAR(64) DEFAULT NULL COMMENT '申诉设备指纹',

    -- 审核信息
                                   `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态 1待处理 2通过 3拒绝',
                                   `reviewer_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '审核人ID',
                                   `review_remark` VARCHAR(255) DEFAULT NULL COMMENT '审核备注',

    -- 行为结果
                                   `action_result` TINYINT NOT NULL DEFAULT 0 COMMENT '处理结果 0无操作 1已解封账号 2已解封设备',

    -- 时间
                                   `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申诉时间',
                                   `reviewed_at` DATETIME DEFAULT NULL COMMENT '审核时间',

                                   PRIMARY KEY (`id`),

                                   KEY `idx_user_id` (`user_id`),
                                   KEY `idx_status` (`status`),
                                   KEY `idx_device_fp` (`device_fp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='风控申诉表';

DROP TABLE IF EXISTS `app_risk_strategy`;
CREATE TABLE `app_risk_strategy` (
                                     `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '策略ID',
                                     `scene` VARCHAR(32) NOT NULL COMMENT '场景：register/login/withdraw 等',
                                     `rule_code` VARCHAR(64) NOT NULL COMMENT '规则编码：如 LOGIN_FAIL_USER',
                                     `identity_type` VARCHAR(16) NOT NULL COMMENT '统计维度：user/ip/device 等',
                                     `window_seconds` INT NOT NULL COMMENT '统计窗口(秒)',
                                     `threshold` INT NOT NULL COMMENT '触发阈值（次数）',
                                     `action` VARCHAR(32) NOT NULL COMMENT '触发动作编码，关联 app_risk_action.code',
                                     `action_value` INT NOT NULL DEFAULT 0 COMMENT '动作值(秒/分数)，覆盖默认值时使用',
                                     `status` TINYINT NOT NULL DEFAULT 1 COMMENT '1启用 0停用',
                                     `priority` INT NOT NULL DEFAULT 100 COMMENT '优先级，数值越小越优先',
                                     `remark` VARCHAR(255) DEFAULT NULL COMMENT '说明',
                                     `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                     `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `uk_scene_rule` (`scene`,`rule_code`),
                                     KEY `idx_scene_status_priority` (`scene`,`status`,`priority`),
                                     KEY `idx_scene_identity` (`scene`,`identity_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='风控策略表：定义场景下的统计窗口、阈值与动作';

DROP TABLE IF EXISTS `app_risk_action`;
CREATE TABLE `app_risk_action` (
                                   `code` VARCHAR(32) NOT NULL COMMENT '动作编码，如 LOCK_5M/BAN/SCORE_50',
                                   `type` VARCHAR(16) NOT NULL COMMENT '动作类型：LOCK/BAN/SCORE',
                                   `default_value` INT NOT NULL DEFAULT 0 COMMENT '默认动作值：LOCK秒数/BAN=0/SCORE分值',
                                   `remark` VARCHAR(255) DEFAULT NULL COMMENT '说明',
                                   PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='风控动作字典：定义动作类型与默认值';

INSERT INTO app_risk_action VALUES
                                ('LOCK_5M','LOCK',300,'锁定5分钟'),
                                ('LOCK_30M','LOCK',1800,'锁定30分钟'),
                                ('BAN','BAN',0,'永久封禁'),
                                ('SCORE_50','SCORE',50,'增加50风险分'),
                                ('SCORE_80','SCORE',80,'增加80风险分');

DROP TABLE IF EXISTS `app_risk_strategy_cache`;
CREATE TABLE `app_risk_strategy_cache` (
                                           `scene` VARCHAR(32) NOT NULL COMMENT '场景：register/login/withdraw',
                                           `identity_type` VARCHAR(16) NOT NULL COMMENT '统计维度：user/ip/device',
                                           `rule_code` VARCHAR(64) NOT NULL COMMENT '规则编码',
                                           `window_seconds` INT NOT NULL COMMENT '统计窗口(秒)',
                                           `threshold` INT NOT NULL COMMENT '触发阈值（次数）',
                                           `action` VARCHAR(32) NOT NULL COMMENT '触发动作编码',
                                           `action_value` INT NOT NULL COMMENT '动作值(秒/分数)',
                                           PRIMARY KEY (`scene`,`identity_type`,`rule_code`),
                                           KEY `idx_scene_identity` (`scene`,`identity_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='风控策略运行缓存：策略下发后的本地缓存';

DROP TABLE IF EXISTS `app_message`;
CREATE TABLE `app_message` (
                               `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '消息ID',

                               `msg_type` TINYINT NOT NULL DEFAULT 1 COMMENT '消息类型：1系统通知 2站内信 3私信',

                               `title` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '消息标题',
                               `content` TEXT NOT NULL COMMENT '消息内容',

                               `sender_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送者ID（0=系统）',
                               `sender_type` TINYINT NOT NULL DEFAULT 0 COMMENT '发送者类型：0系统 1用户 2管理员',
                               `sender_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '发送者名称（冗余字段）',

                               `biz_type` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '业务类型：order / audit / activity',
                               `biz_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '业务ID',

                               `extra` JSON NULL COMMENT '扩展数据（跳转参数、按钮等）',

                               `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',

                               PRIMARY KEY (`id`),
                               KEY `idx_type_time` (`msg_type`, `created_at`),
                               KEY `idx_sender` (`sender_type`, `sender_id`),
                               KEY `idx_biz` (`biz_type`, `biz_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站内信-消息主体';

DROP TABLE IF EXISTS `app_message_user`;
CREATE TABLE `app_message_user` (
                                    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',

                                    `message_id` BIGINT UNSIGNED NOT NULL COMMENT '消息ID',
                                    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '接收用户ID',

                                    `is_read` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已读：0未读 1已读',
                                    `is_deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除：0否 1是',

                                    `read_at` DATETIME NULL COMMENT '阅读时间',
                                    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '接收时间',

                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `uk_user_message` (`user_id`, `message_id`),
                                    KEY `idx_user_read` (`user_id`, `is_read`),
                                    KEY `idx_user_time` (`user_id`, `created_at`),
                                    KEY `idx_user_deleted` (`user_id`, `is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站内信-用户收件箱';


DROP TABLE IF EXISTS `app_message_template`;
CREATE TABLE `app_message_template` (
                                        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '模板ID',

                                        `code` VARCHAR(50) NOT NULL COMMENT '模板编码（唯一，如：withdraw_fail）',
                                        `msg_type` TINYINT NOT NULL DEFAULT 1 COMMENT '消息类型：1系统通知 2站内信 3私信',

                                        `title_tpl` VARCHAR(100) NOT NULL COMMENT '标题模板',
                                        `content_tpl` TEXT NOT NULL COMMENT '内容模板（支持变量）',

                                        `sender_type` TINYINT NOT NULL DEFAULT 0 COMMENT '发送者类型：0系统 1用户 2管理员',
                                        `sender_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送者ID（0=系统）',
                                        `sender_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '发送者名称（冗余）',

                                        `biz_type` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '业务类型：order / withdraw / risk',
                                        `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1启用 0禁用',

                                        `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',

                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `uk_code` (`code`),
                                        KEY `idx_type_status` (`msg_type`, `status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站内信-消息模板';



DROP TABLE IF EXISTS `app_user_invite_code`;
CREATE TABLE `app_user_invite_code` (
                                   `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `code` varchar(64) NOT NULL COMMENT '邀请码',
                                   `owner_user_id` BIGINT UNSIGNED NOT NULL COMMENT '所属用户',
                                   `status` char(1) NOT NULL DEFAULT '1' COMMENT '1可用 2禁用',
                                   `total_limit` int NOT NULL DEFAULT 0 COMMENT '总次数 0不限制',
                                   `daily_limit` int NOT NULL DEFAULT 0 COMMENT '每日次数 0不限制',
                                   `used_total` int NOT NULL DEFAULT 0 COMMENT '已使用总次数',
                                   `used_today` int NOT NULL DEFAULT 0 COMMENT '今日已使用次数',
                                   `last_used_at` datetime DEFAULT NULL COMMENT '最后使用时间',
                                   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `uk_code` (`code`),
                                   KEY `idx_owner` (`owner_user_id`),
                                   KEY `idx_status` (`status`),
                                   KEY `idx_owner_status` (`owner_user_id`,`status`),
                                   KEY `idx_used_today` (`used_today`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邀请码表';

-- ----------------------------
-- Table structure for app_user_invite_relation
-- ----------------------------
DROP TABLE IF EXISTS `app_user_invite_relation`;
CREATE TABLE `app_user_invite_relation` (
                                            `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '邀请关系ID',
                                            `inviter_user_id` BIGINT UNSIGNED NOT NULL COMMENT '邀请人用户ID',
                                            `invitee_user_id` BIGINT UNSIGNED NOT NULL COMMENT '被邀请人用户ID',
                                            `invite_code` varchar(64) NOT NULL COMMENT '使用的邀请码',
                                            `invite_reward` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '邀请奖励',
                                            `invitee_reward` DECIMAL(30,2) NOT NULL DEFAULT 0.00 COMMENT '被邀请人奖励',
                                            `status` char(1) NOT NULL DEFAULT '1' COMMENT '状态 1有效 2无效',
                                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '邀请时间',
                                            PRIMARY KEY (`id`),
                                            UNIQUE KEY `uk_invitee` (`invitee_user_id`),
                                            KEY `idx_inviter` (`inviter_user_id`),
                                            KEY `idx_invite_code` (`invite_code`),
                                            KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户邀请关系表';


SET FOREIGN_KEY_CHECKS = 1;
