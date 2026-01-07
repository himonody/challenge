-- SSE 模块数据库迁移
-- 创建时间: 2026-01-07
-- 说明: 创建 SSE 消息记录表和订阅关系表

-- ========================================
-- 1. 消息记录表（用于持久化和重连恢复）
-- ========================================
CREATE TABLE IF NOT EXISTS `app_sse_message` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `event_id` varchar(64) NOT NULL COMMENT '事件ID',
  `event_type` varchar(64) NOT NULL COMMENT '事件类型',
  `receiver_id` varchar(64) NOT NULL COMMENT '接收者ID（用户ID或客户端ID）',
  `receiver_type` varchar(20) NOT NULL COMMENT '接收者类型（user/client/group）',
  `group_name` varchar(64) DEFAULT '' COMMENT '分组名称',
  `priority` tinyint DEFAULT 0 COMMENT '优先级（0-普通 1-高 2-紧急）',
  `data` text COMMENT '消息数据（JSON）',
  `status` tinyint DEFAULT 0 COMMENT '状态（0-待发送 1-已发送 2-已读）',
  `expire_at` datetime DEFAULT NULL COMMENT '过期时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_event_id` (`event_id`),
  KEY `idx_event_type` (`event_type`),
  KEY `idx_receiver_created` (`receiver_id`, `created_at`),
  KEY `idx_group_created` (`group_name`, `created_at`),
  KEY `idx_expire_at` (`expire_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='SSE消息记录表';

-- ========================================
-- 2. 订阅关系表
-- ========================================
CREATE TABLE IF NOT EXISTS `app_sse_subscription` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `group_name` varchar(64) NOT NULL COMMENT '订阅组名',
  `event_types` varchar(255) DEFAULT '' COMMENT '订阅的事件类型（逗号分隔）',
  `status` tinyint DEFAULT 1 COMMENT '状态（0-禁用 1-启用）',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_group` (`user_id`, `group_name`),
  KEY `idx_group_status` (`group_name`, `status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='SSE订阅关系表';

-- ========================================
-- 3. 示例数据（可选）
-- ========================================

-- 示例：创建一些测试订阅
-- INSERT INTO `app_sse_subscription` (`user_id`, `group_name`, `event_types`, `status`, `created_at`)
-- VALUES 
-- ('user001', 'notifications', ',notification,new_message,', 1, NOW()),
-- ('user001', 'challenge_100', ',challenge_update,rank_change,', 1, NOW());

-- 示例：创建一些测试消息
-- INSERT INTO `app_sse_message` (`event_id`, `event_type`, `receiver_id`, `receiver_type`, `group_name`, `priority`, `data`, `status`, `created_at`)
-- VALUES 
-- ('evt_001', 'notification', 'user001', 'user', '', 1, '{"title":"测试通知","content":"这是一条测试消息"}', 0, NOW()),
-- ('evt_002', 'challenge_update', 'user001', 'user', 'challenge_100', 0, '{"challenge_id":100,"rank":5}', 0, NOW());
