package service

import (
	"challenge-admin/app/app/withdraw/models"
	"challenge-admin/app/app/withdraw/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
	"challenge-admin/core/utils/dateutils"
	"errors"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type WithdrawOrder struct {
	service.Service
}

// GetPage 获取提现订单分页
func (e *WithdrawOrder) GetPage(c *dto.WithdrawOrderQueryReq, p *middleware.DataPermission) ([]models.WithdrawOrder, int64, int, error) {
	var list []models.WithdrawOrder
	var data models.WithdrawOrder
	var count int64

	err := e.Orm.Table(data.TableName()).
		Select("app_withdraw_order.*, u.username as username").
		Joins("LEFT JOIN app_user u ON app_withdraw_order.user_id = u.id").
		Order("app_withdraw_order.id desc").
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

// Get 详情
func (e *WithdrawOrder) Get(id int64, p *middleware.DataPermission) (*models.WithdrawOrder, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.WithdrawOrder{}
	err := e.Orm.Table(data.TableName()).
		Select("app_withdraw_order.*, u.username as user_name").
		Joins("LEFT JOIN app_user u ON app_withdraw_order.user_id = u.id").
		Scopes(
			middleware.Permission(data.TableName(), p),
		).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// UpdateStatus 审核/更新状态
func (e *WithdrawOrder) UpdateStatus(c *dto.WithdrawOrderUpdateStatusReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.ReviewerId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	updates := map[string]interface{}{}
	if c.Status > 0 && data.Status != c.Status {
		updates["status"] = c.Status
		updates["reviewed_at"] = time.Now()
	}
	if c.RejectReason != "" {
		updates["reject_reason"] = c.RejectReason
	}
	if c.ReviewIp != "" {
		updates["review_ip"] = c.ReviewIp
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	updates["update_by"] = c.ReviewerId
	updates["updated_at"] = time.Now()
	err = e.Orm.Model(&models.WithdrawOrder{}).Where("id=?", c.Id).Updates(updates).Error
	if err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

// Export 导出提现订单
func (e *WithdrawOrder) Export(list []models.WithdrawOrder) ([]byte, error) {
	sheetName := "WithdrawOrder"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "J", 24)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "用户ID", "金额", "手续费", "状态", "地址", "申请IP", "申请时间", "审核时间", "审核IP", "拒绝原因",
	})
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id,
			item.UserId,
			item.Amount.StringFixedBank(2),
			item.Free.StringFixedBank(2),
			item.Status,
			item.Address,
			item.ApplyIp,
			dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
			dateutils.ConvertToStrByPrt(item.ReviewedAt, -1),
			item.ReviewIp,
			item.RejectReason,
		})
	}
	xlsx.SetActiveSheet(no)
	buf, _ := xlsx.WriteToBuffer()
	_ = xlsx.Close()
	return buf.Bytes(), nil
}
