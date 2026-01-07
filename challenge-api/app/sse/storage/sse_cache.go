package storage

import (
	"challenge/core/utils/storage"
	"encoding/json"
	"fmt"
)

// 缓存Key前缀定义
const (
	SSEOnlineUsersPrefix  = "challenge:sse:online:users"  // 在线用户集合
	SSEUserClientsPrefix  = "challenge:sse:user:clients"  // 用户的客户端列表
	SSEClientInfoPrefix   = "challenge:sse:client:info"   // 客户端信息
	SSEGroupMembersPrefix = "challenge:sse:group:members" // 分组成员
	SSEUnreadCountPrefix  = "challenge:sse:unread:count"  // 未读消息计数
)

// 推荐的过期时间常量（秒）
const (
	DefaultOnlineExpire     = 300   // 在线状态：5分钟
	DefaultClientInfoExpire = 600   // 客户端信息：10分钟
	DefaultUnreadExpire     = 86400 // 未读计数：24小时
)

// ClientInfo 客户端信息
type ClientInfo struct {
	ClientID    string            `json:"client_id"`
	UserID      string            `json:"user_id"`
	Group       string            `json:"group"`
	ConnectedAt int64             `json:"connected_at"`
	Metadata    map[string]string `json:"metadata"`
}

// ========== 在线用户管理 ==========

// AddOnlineUser 添加在线用户
func AddOnlineUser(cache storage.AdapterCache, userID string) error {
	return cache.Set(SSEOnlineUsersPrefix, userID, "1", DefaultOnlineExpire)
}

// RemoveOnlineUser 移除在线用户
func RemoveOnlineUser(cache storage.AdapterCache, userID string) error {
	return cache.Del(SSEOnlineUsersPrefix, userID)
}

// IsUserOnline 检查用户是否在线
func IsUserOnline(cache storage.AdapterCache, userID string) bool {
	return cache.Exist(SSEOnlineUsersPrefix, userID)
}

// ========== 用户客户端管理 ==========

// AddUserClient 添加用户的客户端
func AddUserClient(cache storage.AdapterCache, userID string, clientID string, expire int) error {
	key := fmt.Sprintf("%s:%s", userID, clientID)
	return cache.Set(SSEUserClientsPrefix, key, clientID, expire)
}

// RemoveUserClient 移除用户的客户端
func RemoveUserClient(cache storage.AdapterCache, userID string, clientID string) error {
	key := fmt.Sprintf("%s:%s", userID, clientID)
	return cache.Del(SSEUserClientsPrefix, key)
}

// ========== 客户端信息管理 ==========

// SetClientInfo 设置客户端信息
func SetClientInfo(cache storage.AdapterCache, info *ClientInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return cache.Set(SSEClientInfoPrefix, info.ClientID, string(data), DefaultClientInfoExpire)
}

// GetClientInfo 获取客户端信息
func GetClientInfo(cache storage.AdapterCache, clientID string) (*ClientInfo, error) {
	data, err := cache.Get(SSEClientInfoPrefix, clientID)
	if err != nil {
		return nil, err
	}

	var info ClientInfo
	if err := json.Unmarshal([]byte(data), &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// DeleteClientInfo 删除客户端信息
func DeleteClientInfo(cache storage.AdapterCache, clientID string) error {
	return cache.Del(SSEClientInfoPrefix, clientID)
}

// ========== 分组成员管理 ==========

// AddGroupMember 添加分组成员
func AddGroupMember(cache storage.AdapterCache, group string, clientID string, expire int) error {
	key := fmt.Sprintf("%s:%s", group, clientID)
	return cache.Set(SSEGroupMembersPrefix, key, clientID, expire)
}

// RemoveGroupMember 移除分组成员
func RemoveGroupMember(cache storage.AdapterCache, group string, clientID string) error {
	key := fmt.Sprintf("%s:%s", group, clientID)
	return cache.Del(SSEGroupMembersPrefix, key)
}

// ========== 未读消息计数 ==========

// IncrUnreadCount 增加未读消息计数
func IncrUnreadCount(cache storage.AdapterCache, userID string) error {
	// 如果key不存在，会自动创建并设置为1
	err := cache.Increase(SSEUnreadCountPrefix, userID)
	if err != nil {
		return err
	}
	// 设置过期时间
	return cache.Expire(SSEUnreadCountPrefix, userID, DefaultUnreadExpire)
}

// GetUnreadCount 获取未读消息计数
func GetUnreadCount(cache storage.AdapterCache, userID string) (int, error) {
	count, err := cache.Get(SSEUnreadCountPrefix, userID)
	if err != nil {
		return 0, nil // 不存在返回0
	}

	var result int
	if _, err := fmt.Sscanf(count, "%d", &result); err != nil {
		return 0, err
	}

	return result, nil
}

// ResetUnreadCount 重置未读消息计数
func ResetUnreadCount(cache storage.AdapterCache, userID string) error {
	return cache.Del(SSEUnreadCountPrefix, userID)
}

// DecrUnreadCount 减少未读消息计数
func DecrUnreadCount(cache storage.AdapterCache, userID string, count int) error {
	for i := 0; i < count; i++ {
		if err := cache.Decrease(SSEUnreadCountPrefix, userID); err != nil {
			return err
		}
	}
	return nil
}
