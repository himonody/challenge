package repo

import (
	userModels "challenge/app/user/models"
	"time"

	"gorm.io/gorm"
)

// CreateUser 创建用户
func CreateUser(db *gorm.DB, user *userModels.AppUser) error {
	return db.Table("app_user").Create(user).Error
}

// GetUserByID 根据ID获取用户信息
func GetUserByID(db *gorm.DB, userID uint64) (*userModels.AppUser, error) {
	var user userModels.AppUser
	err := db.Table("app_user").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户信息
func GetUserByUsername(db *gorm.DB, username string) (*userModels.AppUser, error) {
	var user userModels.AppUser
	err := db.Table("app_user").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(db *gorm.DB, userID uint64, updates map[string]interface{}) error {
	return db.Table("app_user").Where("id = ?", userID).Updates(updates).Error
}

// UpdateUserPassword 更新用户登录密码
func UpdateUserPassword(db *gorm.DB, userID uint64, newPassword string) error {
	return db.Table("app_user").Where("id = ?", userID).Update("pwd", newPassword).Error
}

// UpdateUserPayPassword 更新用户支付密码
func UpdateUserPayPassword(db *gorm.DB, userID uint64, newPayPassword string) error {
	return db.Table("app_user").Where("id = ?", userID).Update("pay_pwd", newPayPassword).Error
}

// GetInviteCodeByUserID 根据用户ID获取邀请码信息
func GetInviteCodeByUserID(db *gorm.DB, userID uint64) (*userModels.UserInviteCode, error) {
	var inviteCode userModels.UserInviteCode
	err := db.Table("app_user_invite_code").
		Where("owner_user_id = ? AND status = ?", userID, "1").
		First(&inviteCode).Error
	if err != nil {
		return nil, err
	}
	return &inviteCode, nil
}

// CreateInviteCode 创建邀请码
func CreateInviteCode(db *gorm.DB, inviteCode *userModels.UserInviteCode) error {
	return db.Table("app_user_invite_code").Create(inviteCode).Error
}

// GetInviteRelationsByInviter 获取用户的邀请记录
func GetInviteRelationsByInviter(db *gorm.DB, inviterUserID uint64, page, pageSize int) ([]userModels.AppUserInviteRelation, int64, error) {
	var relations []userModels.AppUserInviteRelation
	var total int64

	offset := (page - 1) * pageSize

	// 统计总数
	if err := db.Table("app_user_invite_relation").
		Where("inviter_user_id = ? AND status = ?", inviterUserID, "1").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表
	err := db.Table("app_user_invite_relation").
		Where("inviter_user_id = ? AND status = ?", inviterUserID, "1").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&relations).Error

	return relations, total, err
}

// CountTodayInvites 统计今日邀请人数
func CountTodayInvites(db *gorm.DB, inviterUserID uint64) (int64, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	err := db.Table("app_user_invite_relation").
		Where("inviter_user_id = ? AND DATE(created_at) = ? AND status = ?", inviterUserID, today, "1").
		Count(&count).Error
	return count, err
}

// CountTotalInvites 统计总邀请人数
func CountTotalInvites(db *gorm.DB, inviterUserID uint64) (int64, error) {
	var count int64
	err := db.Table("app_user_invite_relation").
		Where("inviter_user_id = ? AND status = ?", inviterUserID, "1").
		Count(&count).Error
	return count, err
}

// SumInviteRewardToday 统计今日邀请收益
func SumInviteRewardToday(db *gorm.DB, inviterUserID uint64) (float64, error) {
	var sum float64
	today := time.Now().Format("2006-01-02")
	err := db.Table("app_user_invite_relation").
		Where("inviter_user_id = ? AND DATE(created_at) = ? AND status = ?", inviterUserID, today, "1").
		Select("COALESCE(SUM(invite_reward), 0)").
		Scan(&sum).Error
	return sum, err
}

// SumInviteRewardTotal 统计总邀请收益
func SumInviteRewardTotal(db *gorm.DB, inviterUserID uint64) (float64, error) {
	var sum float64
	err := db.Table("app_user_invite_relation").
		Where("inviter_user_id = ? AND status = ?", inviterUserID, "1").
		Select("COALESCE(SUM(invite_reward), 0)").
		Scan(&sum).Error
	return sum, err
}
