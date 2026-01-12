package repo

import (
	challengeModels "challenge/app/challenge/models"
	"time"

	"gorm.io/gorm"
)

// GetUserActiveChallenge 获取用户进行中的挑战
func GetUserActiveChallenge(db *gorm.DB, userID uint64) (*challengeModels.AppChallengeUser, error) {
	var challenge challengeModels.AppChallengeUser
	err := db.Table("app_challenge_user").
		Where("user_id = ? AND status = ?", userID, 1). // 状态1=进行中
		First(&challenge).Error
	if err != nil {
		return nil, err
	}
	return &challenge, nil
}

// CountUserTotalCheckin 统计用户总打卡天数
func CountUserTotalCheckin(db *gorm.DB, userID uint64) (int64, error) {
	var count int64
	err := db.Table("app_challenge_checkin").
		Where("user_id = ? AND status = ?", userID, 1). // 状态1=成功
		Count(&count).Error
	return count, err
}

// CountUserTotalMissCheckin 统计用户总未打卡天数（失败的挑战）
func CountUserTotalMissCheckin(db *gorm.DB, userID uint64) (int64, error) {
	var count int64
	err := db.Table("app_challenge_user").
		Where("user_id = ? AND status = ? AND fail_reason = ?", userID, 3, 1). // 状态3=失败，fail_reason=1未打卡
		Count(&count).Error
	return count, err
}

// GetUserContinuousCheckin 获取用户连续打卡天数
func GetUserContinuousCheckin(db *gorm.DB, userID uint64) (int, error) {
	// 获取用户最近的打卡记录
	var checkins []challengeModels.AppChallengeUser
	err := db.Table("app_challenge_checkin").
		Where("user_id = ? AND status = ?", userID, 1).
		Order("checkin_date DESC").
		Limit(365). // 最多查询一年
		Find(&checkins).Error
	if err != nil {
		return 0, err
	}

	if len(checkins) == 0 {
		return 0, nil
	}

	// 计算连续天数
	continuous := 0
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	// 检查今天或昨天是否打卡
	firstDate := checkins[0].FinishedAt.Format("2006-01-02")
	if firstDate != today && firstDate != yesterday {
		return 0, nil
	}

	// 从第一个记录开始往前推算连续天数
	lastDate := checkins[0].FinishedAt
	continuous = 1

	for i := 1; i < len(checkins); i++ {
		expectedDate := lastDate.AddDate(0, 0, -1)
		if checkins[i].FinishedAt.Format("2006-01-02") == expectedDate.Format("2006-01-02") {
			continuous++
			lastDate = checkins[i].FinishedAt
		} else {
			break
		}
	}

	return continuous, nil
}

// CheckTodayCheckin 检查今天是否打卡
func CheckTodayCheckin(db *gorm.DB, userID uint64) (bool, error) {
	var count int64
	today := time.Now().Format("2006-01-02")
	err := db.Table("app_challenge_checkin").
		Where("user_id = ? AND DATE(checkin_date) = ? AND status = ?", userID, today, 1).
		Count(&count).Error
	return count > 0, err
}

// SumUserTotalSettlement 统计用户总收益（结算金额）
func SumUserTotalSettlement(db *gorm.DB, userID uint64) (float64, error) {
	var sum float64
	err := db.Table("app_challenge_settlement").
		Where("user_id = ?", userID).
		Select("COALESCE(SUM(reward), 0)").
		Scan(&sum).Error
	return sum, err
}

// SumUserTodaySettlement 统计用户今日收益
func SumUserTodaySettlement(db *gorm.DB, userID uint64) (float64, error) {
	var sum float64
	today := time.Now().Format("2006-01-02")
	err := db.Table("app_challenge_settlement").
		Where("user_id = ? AND DATE(created_at) = ?", userID, today).
		Select("COALESCE(SUM(reward), 0)").
		Scan(&sum).Error
	return sum, err
}

// GetUserChallengeAmount 获取用户挑战金额
func GetUserChallengeAmount(db *gorm.DB, userID uint64) (float64, error) {
	var sum float64
	err := db.Table("app_challenge_user").
		Where("user_id = ? AND status = ?", userID, 1). // 进行中的挑战
		Select("COALESCE(SUM(challenge_amount), 0)").
		Scan(&sum).Error
	return sum, err
}
