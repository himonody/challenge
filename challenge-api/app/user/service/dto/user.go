package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// GetProfileReq 获取用户资料请求
type GetProfileReq struct {
	UserID uint64 `json:"userId" binding:"required"` // 用户ID
}

// GetProfileResp 获取用户资料响应
type GetProfileResp struct {
	ID          uint64          `json:"id"`          // 用户ID
	Username    string          `json:"username"`    // 用户名
	Nickname    string          `json:"nickname"`    // 昵称
	TrueName    string          `json:"trueName"`    // 真实姓名
	Money       decimal.Decimal `json:"money"`       // 余额
	FreezeMoney decimal.Decimal `json:"freezeMoney"` // 冻结金额
	Email       string          `json:"email"`       // 邮箱
	MobileTitle string          `json:"mobileTitle"` // 手机号国家前缀
	Mobile      string          `json:"mobile"`      // 手机号
	Avatar      string          `json:"avatar"`      // 头像
	RefCode     string          `json:"refCode"`     // 推荐码
	LevelID     int             `json:"levelId"`     // 等级ID
	Status      string          `json:"status"`      // 状态
	RegisterAt  time.Time       `json:"registerAt"`  // 注册时间
	RegisterIP  string          `json:"registerIp"`  // 注册IP
	LastLoginAt *time.Time      `json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string          `json:"lastLoginIp"` // 最后登录IP
}

// ChangeLoginPwdReq 修改登录密码请求
type ChangeLoginPwdReq struct {
	UserID      uint64 `json:"userId" binding:"required"`      // 用户ID
	OldPassword string `json:"oldPassword" binding:"required"` // 旧密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

// ChangePayPwdReq 修改支付密码请求
type ChangePayPwdReq struct {
	UserID    uint64 `json:"userId" binding:"required"`    // 用户ID
	OldPayPwd string `json:"oldPayPwd" binding:"required"` // 旧支付密码
	NewPayPwd string `json:"newPayPwd" binding:"required"` // 新支付密码
}

// UpdateProfileReq 修改用户资料请求（除密码）
type UpdateProfileReq struct {
	UserID      uint64 `json:"userId" binding:"required"` // 用户ID
	Nickname    string `json:"nickname"`                  // 昵称
	TrueName    string `json:"trueName"`                  // 真实姓名
	Email       string `json:"email"`                     // 邮箱
	MobileTitle string `json:"mobileTitle"`               // 手机号国家前缀
	Mobile      string `json:"mobile"`                    // 手机号
	Avatar      string `json:"avatar"`                    // 头像
}

// GetInviteInfoReq 邀请好友请求
type GetInviteInfoReq struct {
	UserID uint64 `json:"userId" binding:"required"` // 用户ID
}

// GetInviteInfoResp 邀请好友响应
type GetInviteInfoResp struct {
	InviteCode string `json:"inviteCode"` // 邀请码
	InviteURL  string `json:"inviteUrl"`  // 邀请链接
	UsedTotal  int    `json:"usedTotal"`  // 已使用总次数
	TotalLimit int    `json:"totalLimit"` // 总次数限制
	DailyLimit int    `json:"dailyLimit"` // 每日次数限制
	UsedToday  int    `json:"usedToday"`  // 今日已使用次数
}

// GetMyInvitesReq 我的邀请请求
type GetMyInvitesReq struct {
	UserID   uint64 `json:"userId" binding:"required"` // 用户ID
	Page     int    `json:"page"`                      // 页码
	PageSize int    `json:"pageSize"`                  // 每页数量
}

// InviteeInfo 被邀请人信息
type InviteeInfo struct {
	UserID       uint64          `json:"userId"`       // 用户ID
	Username     string          `json:"username"`     // 用户名
	Nickname     string          `json:"nickname"`     // 昵称
	Avatar       string          `json:"avatar"`       // 头像
	InviteReward decimal.Decimal `json:"inviteReward"` // 邀请奖励
	CreatedAt    time.Time       `json:"createdAt"`    // 邀请时间
}

// GetMyInvitesResp 我的邀请响应
type GetMyInvitesResp struct {
	Total    int64         `json:"total"`    // 总数
	List     []InviteeInfo `json:"list"`     // 邀请列表
	Page     int           `json:"page"`     // 当前页
	PageSize int           `json:"pageSize"` // 每页数量
}

// GetStatisticsReq 统计请求
type GetStatisticsReq struct {
	UserID uint64 `json:"userId" binding:"required"` // 用户ID
}

// GetStatisticsResp 统计响应
type GetStatisticsResp struct {
	// 打卡相关
	TotalCheckin      int `json:"totalCheckin"`      // 总打卡天数
	TotalMissCheckin  int `json:"totalMissCheckin"`  // 总未打卡天数
	ContinuousCheckin int `json:"continuousCheckin"` // 连续打卡天数

	// 挑战相关
	ChallengeAmount  decimal.Decimal `json:"challengeAmount"`  // 挑战金
	ExperienceAmount decimal.Decimal `json:"experienceAmount"` // 体验金

	// 收益相关
	PlatformBonus decimal.Decimal `json:"platformBonus"` // 平台补贴
	WanfenIncome  decimal.Decimal `json:"wanfenIncome"`  // 万份收益
	TodayIncome   decimal.Decimal `json:"todayIncome"`   // 今日收益
	TotalIncome   decimal.Decimal `json:"totalIncome"`   // 总收益

	// 邀请相关
	TodayInvite       int             `json:"todayInvite"`       // 今日邀请人数
	TotalInvite       int             `json:"totalInvite"`       // 总邀请人数
	InviteRewardToday decimal.Decimal `json:"inviteRewardToday"` // 今日邀请收益
	InviteRewardTotal decimal.Decimal `json:"inviteRewardTotal"` // 总邀请收益
}

// GetTodayStatReq 今日统计请求
type GetTodayStatReq struct {
	UserID uint64 `json:"userId" binding:"required"` // 用户ID
}

// GetTodayStatResp 今日统计响应
type GetTodayStatResp struct {
	TodayCheckin      bool            `json:"todayCheckin"`      // 今日是否打卡
	TodayIncome       decimal.Decimal `json:"todayIncome"`       // 今日收益
	TodayInvite       int             `json:"todayInvite"`       // 今日邀请人数
	TodayInviteReward decimal.Decimal `json:"todayInviteReward"` // 今日邀请收益
	ContinuousCheckin int             `json:"continuousCheckin"` // 连续打卡天数
	ChallengeStatus   string          `json:"challengeStatus"`   // 挑战状态（进行中/成功/失败）
}
