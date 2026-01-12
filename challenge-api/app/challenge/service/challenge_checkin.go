package service

import (
	challengeStorage "challenge/app/challenge/storage"
	"challenge/config/base/constant"
	"challenge/core/middleware/auth/authdto"
	"strconv"

	"challenge/core/dto/service"
	"errors"
)

type ChallengeCheckin struct {
	service.Service
}

func NewChallengeCheckin(s *service.Service) *ChallengeCheckin {
	return &ChallengeCheckin{
		Service: *s,
	}
}

func (s *ChallengeCheckin) Start() error {
	userID := s.C.GetUint64(authdto.UserId)
	if userID == 0 {
		return errors.New("user id is 0")
	}

	locker := s.Run.GetLockerPrefix(challengeStorage.ChallengeLockPrefix)
	err := challengeStorage.WithCheckinLock(s.C.Request.Context(), locker, constant.ChallengeCheckinAction, strconv.Itoa(int(userID)), 10, func() error {

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
