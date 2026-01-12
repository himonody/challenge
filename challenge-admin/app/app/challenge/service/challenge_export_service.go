package service

import (
	"challenge-admin/app/app/challenge/models"
	"challenge-admin/core/utils/dateutils"
	"fmt"

	"github.com/xuri/excelize/v2"
)

// ChallengeConfig 导出
func (e *ChallengeConfig) ExportConfig(list []models.ChallengeConfig) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.DayCount, item.Amount, item.CheckinStart, item.CheckinEnd, item.PlatformBonus,
			item.Status, item.Sort, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	return exportSimple("ChallengeConfig", []interface{}{"ID", "天数", "金额", "开始", "结束", "平台补贴", "状态", "排序", "创建时间"}, rows)
}

// Checkin
func (e *ChallengeCheckin) ExportCheckin(list []models.ChallengeCheckin) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.ChallengeId, item.UserId, dateutils.ConvertToStr(item.CheckinDate, 0), dateutils.ConvertToStrByPrt(item.CheckinTime, -1),
			item.MoodCode, item.MoodText, item.ContentType, item.Status, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	return exportSimple("ChallengeCheckin", []interface{}{"ID", "挑战ID", "用户ID", "打卡日", "打卡时间", "心情", "心情描述", "内容类型", "状态", "创建时间"}, rows)
}

func (e *ChallengeCheckinImage) ExportCheckinImage(list []models.ChallengeCheckinImage) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.CheckinId, item.UserId, item.ImageUrl, item.ImageHash, item.SortNo, item.Status, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	return exportSimple("ChallengeCheckinImage", []interface{}{"ID", "打卡ID", "用户ID", "图片URL", "图片Hash", "排序", "状态", "创建时间"}, rows)
}

func (e *ChallengeCheckinVideoAd) ExportCheckinVideoAd(list []models.ChallengeCheckinVideoAd) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.CheckinId, item.UserId, item.AdPlatform, item.AdUnitId, item.AdOrderNo,
			item.VideoDuration, item.WatchDuration, item.RewardAmount, item.VerifyStatus,
			dateutils.ConvertToStrByPrt(item.CreatedAt, -1), dateutils.ConvertToStrByPrt(item.VerifiedAt, -1),
		})
	}
	return exportSimple("ChallengeCheckinVideoAd", []interface{}{"ID", "打卡ID", "用户ID", "平台", "广告位", "订单号", "视频时长", "观看时长", "收益", "校验状态", "完成时间", "校验时间"}, rows)
}

// User
func (e *ChallengeUser) ExportChallengeUser(list []models.ChallengeUser) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.UserId, item.ConfigId, item.PoolId, item.ChallengeAmount, item.StartDate, item.EndDate,
			item.Status, item.FailReason, dateutils.ConvertToStrByPrt(item.CreatedAt, -1), dateutils.ConvertToStrByPrt(item.FinishedAt, -1),
		})
	}
	return exportSimple("ChallengeUser", []interface{}{"ID", "用户ID", "配置ID", "奖池ID", "挑战金额", "开始日", "结束日", "状态", "失败原因", "报名时间", "完成时间"}, rows)
}

// Pool
func (e *ChallengePool) ExportChallengePool(list []models.ChallengePool) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.ConfigId, dateutils.ConvertToStrByPrt(item.StartDate, -1), dateutils.ConvertToStrByPrt(item.EndDate, -1),
			item.TotalAmount, item.Settled, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	return exportSimple("ChallengePool", []interface{}{"ID", "配置ID", "开始时间", "结束时间", "奖池金额", "已结算", "创建时间"}, rows)
}

func (e *ChallengePoolFlow) ExportChallengePoolFlow(list []models.ChallengePoolFlow) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.PoolId, item.UserId, item.Amount, item.Type, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	return exportSimple("ChallengePoolFlow", []interface{}{"ID", "奖池ID", "用户ID", "金额", "类型", "创建时间"}, rows)
}

// Settlement
func (e *ChallengeSettlement) ExportChallengeSettlement(list []models.ChallengeSettlement) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.ChallengeId, item.UserId, item.Reward, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	return exportSimple("ChallengeSettlement", []interface{}{"ID", "挑战ID", "用户ID", "奖励金额", "结算时间"}, rows)
}

// Stats
func (e *ChallengeDailyStat) ExportChallengeDailyStat(list []models.ChallengeDailyStat) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			dateutils.ConvertToStr(item.StatDate, 0), item.JoinUserCnt, item.SuccessUserCnt, item.FailUserCnt,
			item.JoinAmount, item.SuccessAmount, item.FailAmount, item.PlatformBonus, item.PoolAmount,
			dateutils.ConvertToStrByPrt(item.CreatedAt, -1), dateutils.ConvertToStrByPrt(item.UpdatedAt, -1),
		})
	}
	return exportSimple("ChallengeDailyStat", []interface{}{"日期", "参与人数", "成功人数", "失败人数", "参与金额", "成功金额", "失败金额", "平台补贴", "奖池金额", "创建时间", "更新时间"}, rows)
}

func (e *ChallengeTotalStat) ExportChallengeTotalStat(list []models.ChallengeTotalStat) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.TotalUserCnt, item.TotalJoinCnt, item.TotalSuccessCnt, item.TotalFailCnt,
			item.TotalJoinAmount, item.TotalSuccessAmount, item.TotalFailAmount, item.TotalPlatformBonus, item.TotalPoolAmount,
			item.UpdatedAt,
		})
	}
	return exportSimple("ChallengeTotalStat", []interface{}{"ID", "累计用户", "累计参与", "累计成功", "累计失败", "参与金额", "成功金额", "失败金额", "平台补贴", "奖池金额", "更新时间"}, rows)
}

func exportSimple(sheetName string, header []interface{}, rows [][]interface{}) ([]byte, error) {
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "Z", 20)
	_ = xlsx.SetSheetRow(sheetName, "A1", &header)
	for i, row := range rows {
		axis := fmt.Sprintf("A%d", i+2)
		_ = xlsx.SetSheetRow(sheetName, axis, &row)
	}
	xlsx.SetActiveSheet(no)
	buf, _ := xlsx.WriteToBuffer()
	_ = xlsx.Close()
	return buf.Bytes(), nil
}
