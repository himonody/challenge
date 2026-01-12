package service

import (
	"challenge-admin/app/app/risk/models"
	"challenge-admin/core/utils/dateutils"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (e *RiskUser) ExportRiskUser(list []models.RiskUser) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{item.UserId, item.RiskLevel, item.RiskScore, item.Reason, dateutils.ConvertToStrByPrt(item.UpdatedAt, -1)})
	}
	return exportSimple("RiskUser", []interface{}{"用户ID", "风险等级", "风险分", "原因", "更新时间"}, rows)
}

func (e *RiskDevice) ExportRiskDevice(list []models.RiskDevice) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{item.Id, item.DeviceFp, item.UserId, dateutils.ConvertToStrByPrt(item.CreatedAt, -1)})
	}
	return exportSimple("RiskDevice", []interface{}{"ID", "设备指纹", "用户ID", "记录时间"}, rows)
}

func (e *RiskEvent) ExportRiskEvent(list []models.RiskEvent) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{item.Id, item.UserId, item.EventType, item.Detail, item.Score, dateutils.ConvertToStrByPrt(item.CreatedAt, -1)})
	}
	return exportSimple("RiskEvent", []interface{}{"ID", "用户ID", "事件类型", "详情", "分值", "时间"}, rows)
}

func (e *RiskRateLimit) ExportRiskRateLimit(list []models.RiskRateLimit) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{item.Id, item.Scene, item.IdentityType, item.IdentityValue, item.Count, dateutils.ConvertToStrByPrt(item.WindowStart, -1), dateutils.ConvertToStrByPrt(item.WindowEnd, -1), item.Blocked})
	}
	return exportSimple("RiskRateLimit", []interface{}{"ID", "场景", "标识类型", "标识值", "次数", "窗口起", "窗口止", "是否拦截"}, rows)
}

func (e *RiskBlacklist) ExportRiskBlacklist(list []models.RiskBlacklist) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{item.Id, item.Type, item.Value, item.RiskLevel, item.Reason, item.Status, dateutils.ConvertToStrByPrt(item.CreatedAt, -1)})
	}
	return exportSimple("RiskBlacklist", []interface{}{"ID", "类型", "命中值", "风险等级", "原因", "状态", "创建时间"}, rows)
}

func (e *RiskAction) ExportRiskAction(list []models.RiskAction) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{item.Code, item.Type, item.DefaultValue, item.Remark})
	}
	return exportSimple("RiskAction", []interface{}{"动作编码", "动作类型", "默认值", "说明"}, rows)
}

func (e *RiskStrategy) ExportRiskStrategy(list []models.RiskStrategy) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.Scene, item.RuleCode, item.IdentityType, item.WindowSeconds, item.Threshold,
			item.Action, item.ActionValue, item.Status, item.Priority, item.Remark,
			dateutils.ConvertToStrByPrt(item.CreatedAt, -1), dateutils.ConvertToStrByPrt(item.UpdatedAt, -1),
		})
	}
	return exportSimple("RiskStrategy", []interface{}{"ID", "场景", "规则编码", "维度", "窗口(秒)", "阈值", "动作", "动作值", "状态", "优先级", "说明", "创建时间", "更新时间"}, rows)
}

func (e *RiskStrategyCache) ExportRiskStrategyCache(list []models.RiskStrategyCache) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Scene, item.IdentityType, item.RuleCode, item.WindowSeconds, item.Threshold, item.Action, item.ActionValue,
		})
	}
	return exportSimple("RiskStrategyCache", []interface{}{"场景", "维度", "规则编码", "窗口(秒)", "阈值", "动作", "动作值"}, rows)
}

func (e *RiskAppeal) ExportRiskAppeal(list []models.RiskAppeal) ([]byte, error) {
	var rows [][]interface{}
	for _, item := range list {
		rows = append(rows, []interface{}{
			item.Id, item.UserId, item.RiskLevel, item.RiskReason, item.AppealType, item.AppealReason, item.AppealEvidence,
			item.Ip, item.DeviceFp, item.Status, item.ReviewerId, item.ReviewRemark, item.ActionResult,
			dateutils.ConvertToStrByPrt(item.CreatedAt, -1), dateutils.ConvertToStrByPrt(item.ReviewedAt, -1),
		})
	}
	return exportSimple("RiskAppeal", []interface{}{"ID", "用户ID", "风险等级", "风险原因", "申诉类型", "申诉说明", "凭证", "IP", "设备指纹", "状态", "审核人", "审核备注", "处理结果", "申诉时间", "审核时间"}, rows)
}
func exportSimple(sheetName string, header []interface{}, rows [][]interface{}) ([]byte, error) {
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "Z", 22)
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
