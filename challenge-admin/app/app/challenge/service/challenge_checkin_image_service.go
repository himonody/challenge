package service

import (
	"challenge-admin/app/app/challenge/models"
	"challenge-admin/app/app/challenge/service/dto"
	baseLang "challenge-admin/config/base/lang"
	cDto "challenge-admin/core/dto"
	"challenge-admin/core/dto/service"
	"challenge-admin/core/lang"
	"challenge-admin/core/middleware"
)

type ChallengeCheckinImage struct {
	service.Service
}

func (e *ChallengeCheckinImage) GetPage(c *dto.ChallengeCheckinImageQueryReq, p *middleware.DataPermission) ([]models.ChallengeCheckinImage, int64, int, error) {
	var list []models.ChallengeCheckinImage
	var data models.ChallengeCheckinImage
	var count int64
	err := e.Orm.Model(&data).
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

func (e *ChallengeCheckinImage) Create(c *dto.ChallengeCheckinImageCreateReq) (uint64, int, error) {
	if c.CurrUserId <= 0 || c.CheckinId == 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := models.ChallengeCheckinImage{
		CheckinId: c.CheckinId,
		UserId:    c.UserId,
		ImageUrl:  c.ImageUrl,
		ImageHash: c.ImageHash,
		SortNo:    c.SortNo,
		Status:    c.Status,
	}
	if err := e.Orm.Create(&data).Error; err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

func (e *ChallengeCheckinImage) Update(c *dto.ChallengeCheckinImageUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	updates := map[string]interface{}{}
	if c.ImageUrl != "" {
		updates["image_url"] = c.ImageUrl
	}
	if c.ImageHash != "" {
		updates["image_hash"] = c.ImageHash
	}
	if c.SortNo != 0 {
		updates["sort_no"] = c.SortNo
	}
	if c.Status != 0 {
		updates["status"] = c.Status
	}
	if len(updates) == 0 {
		return false, baseLang.SuccessCode, nil
	}
	if err := e.Orm.Model(&models.ChallengeCheckinImage{}).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}

func (e *ChallengeCheckinImage) Delete(id uint64, currUser int64) (bool, int, error) {
	if currUser <= 0 || id == 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if err := e.Orm.Delete(&models.ChallengeCheckinImage{}, id).Error; err != nil {
		return false, baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return true, baseLang.SuccessCode, nil
}
