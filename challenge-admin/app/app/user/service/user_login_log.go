package service

import (
	"challenge-admin/app/app/user/models"
	"challenge-admin/app/app/user/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/utils/dateutils"
	"fmt"

	"github.com/xuri/excelize/v2"
)

// UserLoginLog app-用户登录日志
type UserLoginLog struct {
	service.Service
}

// GetPage 获取用户登录日志分页列表
func (e *UserLoginLog) GetPage(c *dto.UserLoginLogQueryReq, p *middleware.DataPermission) ([]models.UserLoginLog, int64, int, error) {
	var data models.UserLoginLog
	var list []models.UserLoginLog
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Export 导出用户登录日志
func (e *UserLoginLog) Export(list []models.UserLoginLog) ([]byte, error) {
	sheetName := "UserLoginLog"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "J", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "用户ID", "登录时间", "登录IP", "设备指纹", "UA", "状态", "失败原因", "记录时间",
	})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.UserId, dateutils.ConvertToStrByPrt(item.LoginAt, -1), item.LoginIp, item.DeviceFp,
			item.UserAgent, item.Status, item.FailReason, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
