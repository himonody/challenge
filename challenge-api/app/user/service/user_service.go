package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	userModels "challenge/app/user/models"
	"challenge/app/user/repo"
	"challenge/app/user/service/dto"
	userStorage "challenge/app/user/storage"
	baseConstant "challenge/config/base/constant"
	baseLang "challenge/config/base/lang"
	coreService "challenge/core/dto/service"
	"challenge/core/lang"
	"challenge/core/utils/encrypt"

	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// UserService 用户服务（参考 auth/service，嵌入通用 Service）
type UserService struct {
	coreService.Service
}

// NewUserService 基于通用 Service 初始化用户服务
func NewUserService(s *coreService.Service) *UserService {
	return &UserService{Service: *s}
}

// 封装金额保留两位小数（截断）
func formatDec2(d decimal.Decimal) decimal.Decimal {
	return d.Truncate(2)
}

func formatFloatToDec2(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f).Truncate(2)
}

// 异步记录操作日志（最佳努力）
func (s *UserService) asyncOperLog(userID uint64, actionType, remark string, status string) {
	ctx := context.Background()
	if s.C != nil && s.C.Request != nil {
		ctx = s.C.Request.Context()
	}
	ip := ""
	if s.C != nil {
		ip = s.C.ClientIP()
	}
	now := time.Now()
	uid := int(userID)
	log := &userModels.AppUserOperLog{
		UserID:     uid,
		ActionType: actionType,
		OperateIP:  ip,
		ByType:     baseConstant.UserOperByTypeApp,
		Status:     status,
		CreateBy:   uid,
		UpdateBy:   uid,
		CreatedAt:  now,
		UpdatedAt:  now,
		Remark:     remark,
	}
	go func() {
		_ = repo.CreateUserOperLog(s.Orm.WithContext(ctx), log)
	}()
}

// GetProfile 获取用户资料
func (s *UserService) GetProfile(req *dto.GetProfileReq) (*dto.GetProfileResp, error) {
	user, err := repo.GetUserByID(s.Orm, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(lang.MsgInfo[baseLang.UserNotFoundCode])
		}
		return nil, err
	}

	// 将 decimal.Decimal 转为 float64
	money := formatDec2(user.Money)
	freezeMoney := formatDec2(user.FreezeMoney)

	resp := &dto.GetProfileResp{
		ID:          uint64(user.ID),
		Username:    user.Username,
		Nickname:    user.Nickname,
		TrueName:    user.TrueName,
		Money:       money,
		FreezeMoney: freezeMoney,
		Email:       user.Email,
		MobileTitle: user.MobileTitle,
		Mobile:      user.Mobile,
		Avatar:      user.Avatar,
		RefCode:     user.RefCode,
		LevelID:     user.LevelID,
		Status:      user.Status,
		RegisterAt:  user.RegisterAt,
		RegisterIP:  user.RegisterIp,
		LastLoginAt: user.LastLoginAt,
		LastLoginIP: user.LastLoginIp,
	}

	return resp, nil
}

// ChangeLoginPassword 修改登录密码
func (s *UserService) ChangeLoginPassword(req *dto.ChangeLoginPwdReq) error {
	locker := s.Run.GetLockerPrefix(userStorage.UserLockPrefix)
	return userStorage.WithUserLock(s.C.Request.Context(), locker, req.UserID, "change_login_pwd", 10, func() error {
		// 获取用户信息
		user, err := repo.GetUserByID(s.Orm, req.UserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New(lang.MsgInfo[baseLang.UserNotFoundCode])
			}
			return err
		}

		// 验证旧密码
		if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(req.OldPassword)); err != nil {
			return errors.New(lang.MsgInfo[baseLang.PasswordErrorCode])
		}

		// 加密新密码
		newHashPwd, _ := encrypt.HashEncrypt(req.NewPassword)

		// 更新密码
		if err := repo.UpdateUserPassword(s.Orm, req.UserID, newHashPwd); err != nil {
			return err
		}
		s.asyncOperLog(req.UserID, baseConstant.UserActionSecurityLoginPw, baseConstant.OperLogRemarkChangeLoginPwd, baseConstant.GeneralStatusOk)
		return nil
	})
}

// ChangePayPassword 修改支付密码
func (s *UserService) ChangePayPassword(req *dto.ChangePayPwdReq) error {
	locker := s.Run.GetLockerPrefix(userStorage.UserLockPrefix)
	return userStorage.WithUserLock(s.C.Request.Context(), locker, req.UserID, "change_pay_pwd", 10, func() error {
		// 获取用户信息
		user, err := repo.GetUserByID(s.Orm, req.UserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New(lang.MsgInfo[baseLang.UserNotFoundCode])
			}
			return err
		}

		// 如果用户有旧支付密码，验证旧支付密码
		if user.PayPwd != "" {
			if err := bcrypt.CompareHashAndPassword([]byte(user.PayPwd), []byte(req.OldPayPwd)); err != nil {
				return errors.New(lang.MsgInfo[baseLang.PayPasswordErrorCode])
			}
		}

		// 加密新支付密码
		newHashPayPwd, _ := encrypt.HashEncrypt(req.NewPayPwd)

		// 更新支付密码
		if err := repo.UpdateUserPayPassword(s.Orm, req.UserID, newHashPayPwd); err != nil {
			return err
		}
		s.asyncOperLog(req.UserID, baseConstant.UserActionSecurityPayPw, baseConstant.OperLogRemarkChangePayPwd, baseConstant.GeneralStatusOk)
		return nil
	})
}

// UpdateProfile 修改用户资料（除密码）
func (s *UserService) UpdateProfile(req *dto.UpdateProfileReq) error {
	locker := s.Run.GetLockerPrefix(userStorage.UserLockPrefix)
	return userStorage.WithUserLock(s.C.Request.Context(), locker, req.UserID, "update_profile", 10, func() error {
		// 构建更新数据
		updates := make(map[string]interface{})

		if req.Nickname != "" {
			updates["nickname"] = req.Nickname
		}
		if req.TrueName != "" {
			updates["true_name"] = req.TrueName
		}
		if req.Email != "" {
			updates["email"] = req.Email
		}
		if req.MobileTitle != "" {
			updates["mobile_title"] = req.MobileTitle
		}
		if req.Mobile != "" {
			updates["mobile"] = req.Mobile
		}
		if req.Avatar != "" {
			updates["avatar"] = req.Avatar
		}

		if len(updates) == 0 {
			return errors.New(lang.MsgInfo[baseLang.UserNoFieldsToUpdateCode])
		}

		updates["update_by"] = req.UserID
		updates["updated_at"] = time.Now()

		// 更新用户信息
		if err := repo.UpdateUser(s.Orm, req.UserID, updates); err != nil {
			return err
		}
		s.asyncOperLog(req.UserID, baseConstant.UserActionProfileAvatar, baseConstant.OperLogRemarkUpdateProfile, baseConstant.GeneralStatusOk)
		return nil
	})
}

// GetInviteInfo 获取邀请信息
func (s *UserService) GetInviteInfo(userID uint64, registerURL string) (*dto.GetInviteInfoResp, error) {
	var inviteCodeRes *dto.GetInviteInfoResp
	locker := s.Run.GetLockerPrefix(userStorage.UserLockPrefix)
	err := userStorage.WithUserLock(s.C.Request.Context(), locker, userID, "invite_info", 10, func() error {
		// 获取用户邀请码
		inviteCode, err := repo.GetInviteCodeByUserID(s.Orm, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 如果不存在，创建邀请码
				inviteCode, err = s.createInviteCode(userID)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		// 构建邀请链接
		inviteURL := fmt.Sprintf("%s?invite_code=%s", registerURL, inviteCode.Code)

		inviteCodeRes = &dto.GetInviteInfoResp{
			InviteCode: inviteCode.Code,
			InviteURL:  inviteURL,
			UsedTotal:  inviteCode.UsedTotal,
			TotalLimit: inviteCode.TotalLimit,
			DailyLimit: inviteCode.DailyLimit,
			UsedToday:  inviteCode.UsedToday,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return inviteCodeRes, nil
}

// createInviteCode 创建邀请码
func (s *UserService) createInviteCode(userID uint64) (*userModels.UserInviteCode, error) {
	// 生成随机邀请码
	code := generateRandomCode(8)

	now := time.Now()
	inviteCode := &userModels.UserInviteCode{
		Code:        code,
		OwnerUserID: userID,
		Status:      "1", // 1表示可用
		TotalLimit:  0,   // 0表示不限制
		DailyLimit:  0,   // 0表示不限制
		UsedTotal:   0,
		UsedToday:   0,
		CreatedAt:   &now,
	}

	if err := repo.CreateInviteCode(s.Orm, inviteCode); err != nil {
		return nil, err
	}

	return inviteCode, nil
}

// generateRandomCode 生成随机邀请码
func generateRandomCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GetMyInvites 获取我的邀请列表
func (s *UserService) GetMyInvites(req *dto.GetMyInvitesReq) (*dto.GetMyInvitesResp, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 获取邀请记录
	relations, total, err := repo.GetInviteRelationsByInviter(s.Orm, req.UserID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 构建响应
	list := make([]dto.InviteeInfo, 0, len(relations))
	for _, relation := range relations {
		// 获取被邀请人信息
		inviteeUser, err := repo.GetUserByID(s.Orm, relation.InviteeUserID)
		if err != nil {
			continue
		}

		list = append(list, dto.InviteeInfo{
			UserID:       uint64(inviteeUser.ID),
			Username:     inviteeUser.Username,
			Nickname:     inviteeUser.Nickname,
			Avatar:       inviteeUser.Avatar,
			InviteReward: formatDec2(relation.InviteReward),
			CreatedAt:    relation.CreatedAt,
		})
	}

	resp := &dto.GetMyInvitesResp{
		Total:    total,
		List:     list,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	return resp, nil
}

// GetStatistics 获取用户统计信息
func (s *UserService) GetStatistics(userID uint64) (*dto.GetStatisticsResp, error) {
	resp := &dto.GetStatisticsResp{
		ExperienceAmount: formatFloatToDec2(0),
		PlatformBonus:    formatFloatToDec2(0),
		WanfenIncome:     formatFloatToDec2(0),
	}

	g, ctx := errgroup.WithContext(s.C.Request.Context())

	g.Go(func() error {
		totalCheckin, err := repo.CountUserTotalCheckin(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.TotalCheckin = int(totalCheckin)
		}
		return err
	})

	g.Go(func() error {
		totalMissCheckin, err := repo.CountUserTotalMissCheckin(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.TotalMissCheckin = int(totalMissCheckin)
		}
		return err
	})

	g.Go(func() error {
		continuousCheckin, err := repo.GetUserContinuousCheckin(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.ContinuousCheckin = continuousCheckin
		}
		return err
	})

	g.Go(func() error {
		challengeAmount, err := repo.GetUserChallengeAmount(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.ChallengeAmount = formatFloatToDec2(challengeAmount)
		}
		return err
	})

	g.Go(func() error {
		todayIncome, err := repo.SumUserTodaySettlement(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.TodayIncome = formatFloatToDec2(todayIncome)
		}
		return err
	})

	g.Go(func() error {
		totalIncome, err := repo.SumUserTotalSettlement(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.TotalIncome = formatFloatToDec2(totalIncome)
		}
		return err
	})

	g.Go(func() error {
		todayInvite, err := repo.CountTodayInvites(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.TodayInvite = int(todayInvite)
		}
		return err
	})

	g.Go(func() error {
		totalInvite, err := repo.CountTotalInvites(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.TotalInvite = int(totalInvite)
		}
		return err
	})

	g.Go(func() error {
		inviteRewardToday, err := repo.SumInviteRewardToday(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.InviteRewardToday = formatFloatToDec2(inviteRewardToday)
		}
		return err
	})

	g.Go(func() error {
		inviteRewardTotal, err := repo.SumInviteRewardTotal(s.Orm.WithContext(ctx), userID)
		if err == nil {
			resp.InviteRewardTotal = formatFloatToDec2(inviteRewardTotal)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTodayStatistics 获取今日统计信息
func (s *UserService) GetTodayStatistics(req *dto.GetTodayStatReq) (*dto.GetTodayStatResp, error) {
	resp := &dto.GetTodayStatResp{}

	g, ctx := errgroup.WithContext(s.C.Request.Context())

	g.Go(func() error {
		todayCheckin, err := repo.CheckTodayCheckin(s.Orm.WithContext(ctx), req.UserID)
		if err == nil {
			resp.TodayCheckin = todayCheckin
		}
		return err
	})

	g.Go(func() error {
		todayIncome, err := repo.SumUserTodaySettlement(s.Orm.WithContext(ctx), req.UserID)
		if err == nil {
			resp.TodayIncome = formatFloatToDec2(todayIncome)
		}
		return err
	})

	g.Go(func() error {
		todayInvite, err := repo.CountTodayInvites(s.Orm.WithContext(ctx), req.UserID)
		if err == nil {
			resp.TodayInvite = int(todayInvite)
		}
		return err
	})

	g.Go(func() error {
		todayInviteReward, err := repo.SumInviteRewardToday(s.Orm.WithContext(ctx), req.UserID)
		if err == nil {
			resp.TodayInviteReward = formatFloatToDec2(todayInviteReward)
		}
		return err
	})

	g.Go(func() error {
		continuousCheckin, err := repo.GetUserContinuousCheckin(s.Orm.WithContext(ctx), req.UserID)
		if err == nil {
			resp.ContinuousCheckin = continuousCheckin
		}
		return err
	})

	// 挑战状态单独处理，保持原有错误语义
	var activeChallengeErr error
	g.Go(func() error {
		activeChallenge, err := repo.GetUserActiveChallenge(s.Orm.WithContext(ctx), req.UserID)
		if err != nil {
			activeChallengeErr = err
			return nil
		}
		switch activeChallenge.Status {
		case 1:
			resp.ChallengeStatus = baseConstant.StatusDescChallengeDoing
		case 2:
			resp.ChallengeStatus = baseConstant.StatusDescChallengeSuccess
		case 3:
			resp.ChallengeStatus = baseConstant.StatusDescChallengeFail
		default:
			resp.ChallengeStatus = baseConstant.StatusDescChallengeUnknown
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	if activeChallengeErr != nil {
		if errors.Is(activeChallengeErr, gorm.ErrRecordNotFound) {
			resp.ChallengeStatus = baseConstant.StatusDescChallengeNone
		} else {
			resp.ChallengeStatus = baseConstant.StatusDescChallengeUnknown
		}
	}

	return resp, nil
}
