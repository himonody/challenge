package apis

import (
	"net/http"

	"challenge/app/user/service"
	"challenge/app/user/service/dto"
	"challenge/core/dto/api"

	"github.com/gin-gonic/gin"
)

// User 用户API（参考 auth/apis 结构）
type User struct {
	api.Api
}

// GetProfile 获取用户资料
// @Summary 获取用户资料
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.GetProfileReq true "请求参数"
// @Success 200 {object} dto.GetProfileResp
// @Router /api/v1/user/profile [post]
func (a *User) GetProfile(c *gin.Context) {
	req := dto.GetProfileReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	resp, err := svc.GetProfile(&req)
	if err != nil {
		a.Error(http.StatusInternalServerError, err.Error())
		return
	}
	a.OK(resp, "success")
}

// ChangeLoginPassword 修改登录密码
// @Summary 修改登录密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.ChangeLoginPwdReq true "请求参数"
// @Success 200 {object} string
// @Router /api/v1/user/change-password [post]
func (a *User) ChangeLoginPassword(c *gin.Context) {
	req := dto.ChangeLoginPwdReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	if err := svc.ChangeLoginPassword(&req); err != nil {
		a.Error(http.StatusInternalServerError, err.Error())
		return
	}
	a.OK(nil, "修改登录密码成功")
}

// ChangePayPassword 修改支付密码
// @Summary 修改支付密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.ChangePayPwdReq true "请求参数"
// @Success 200 {object} string
// @Router /api/v1/user/change-pay-password [post]
func (a *User) ChangePayPassword(c *gin.Context) {
	req := dto.ChangePayPwdReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	if err := svc.ChangePayPassword(&req); err != nil {
		a.Error(http.StatusInternalServerError, err.Error())
		return
	}
	a.OK(nil, "修改支付密码成功")
}

// UpdateProfile 修改用户资料
// @Summary 修改用户资料
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.UpdateProfileReq true "请求参数"
// @Success 200 {object} string
// @Router /api/v1/user/update-profile [post]
func (a *User) UpdateProfile(c *gin.Context) {
	req := dto.UpdateProfileReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	if err := svc.UpdateProfile(&req); err != nil {
		a.Error(http.StatusInternalServerError, err.Error())
		return
	}
	a.OK(nil, "修改用户资料成功")
}

// GetInviteInfo 邀请好友
// @Summary 邀请好友
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.GetInviteInfoReq true "请求参数"
// @Success 200 {object} dto.GetInviteInfoResp
// @Router /api/v1/user/invite-info [post]
func (a *User) GetInviteInfo(c *gin.Context) {
	req := dto.GetInviteInfoReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	// 构建注册URL（可改为配置）
	registerURL := "https://your-domain.com/register"

	resp, err := svc.GetInviteInfo(&req, registerURL)
	if err != nil {
		a.Error(http.StatusInternalServerError, "获取邀请信息失败")
		return
	}
	a.OK(resp, "success")
}

// GetMyInvites 我的邀请
// @Summary 我的邀请
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.GetMyInvitesReq true "请求参数"
// @Success 200 {object} dto.GetMyInvitesResp
// @Router /api/v1/user/my-invites [post]
func (a *User) GetMyInvites(c *gin.Context) {
	req := dto.GetMyInvitesReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	resp, err := svc.GetMyInvites(&req)
	if err != nil {
		a.Error(http.StatusInternalServerError, "获取邀请列表失败")
		return
	}
	a.OK(resp, "success")
}

// GetStatistics 统计
// @Summary 统计
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.GetStatisticsReq true "请求参数"
// @Success 200 {object} dto.GetStatisticsResp
// @Router /api/v1/user/statistics [post]
func (a *User) GetStatistics(c *gin.Context) {
	req := dto.GetStatisticsReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	resp, err := svc.GetStatistics(&req)
	if err != nil {
		a.Error(http.StatusInternalServerError, "获取统计信息失败")
		return
	}
	a.OK(resp, "success")
}

// GetTodayStatistics 今日统计
// @Summary 今日统计
// @Tags 用户
// @Accept json
// @Produce json
// @Param body body dto.GetTodayStatReq true "请求参数"
// @Success 200 {object} dto.GetTodayStatResp
// @Router /api/v1/user/today-statistics [post]
func (a *User) GetTodayStatistics(c *gin.Context) {
	req := dto.GetTodayStatReq{}
	svc := service.UserService{}
	if err := a.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&svc.Service).
		MakeRuntime().
		Errors; err != nil {
		return
	}

	resp, err := svc.GetTodayStatistics(&req)
	if err != nil {
		a.Error(http.StatusInternalServerError, "获取今日统计信息失败")
		return
	}
	a.OK(resp, "success")
}
