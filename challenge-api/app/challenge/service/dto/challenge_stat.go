package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type ChallengeStatResp struct {
	DayStat   ChallengeDayStat   `json:"day_stat"`
	TotalStat ChallengeTotalStat `json:"total_stat"`
}
type ChallengeDayStat struct {
	StatDate time.Time `json:"statDate"` //统计日期 YYYYMMDD

	JoinUserCnt    int `json:"joinUserCnt"`    //参与人数
	SuccessUserCnt int `json:"successUserCnt"` //成功人数
	FailUserCnt    int `json:"failUserCnt"`    //失败人数

	JoinAmount    decimal.Decimal `json:"joinAmount"`    //参与总金额
	SuccessAmount decimal.Decimal `json:"successAmount"` //成功金额
	FailAmount    decimal.Decimal `json:"failAmount"`    //失败金额

	PlatformBonus decimal.Decimal `json:"platformBonus"` //平台补贴
	PoolAmount    decimal.Decimal `json:"poolAmount"`    //奖池金额
}
type ChallengeTotalStat struct {
	ID int `json:"id"`

	TotalUserCnt    int `json:"totalUserCnt"`    //累计用户数
	TotalJoinCnt    int `json:"totalJoinCnt"`    //累计参与人次
	TotalSuccessCnt int `json:"totalSuccessCnt"` //累计成功人次
	TotalFailCnt    int `json:"totalFailCnt"`    //累计失败人次

	TotalJoinAmount    decimal.Decimal `json:"totalJoinAmount"`    //累计参与金额
	TotalSuccessAmount decimal.Decimal `json:"totalSuccessAmount"` //累计成功金额
	TotalFailAmount    decimal.Decimal `json:"totalFailAmount"`    //累计失败金额

	TotalPlatformBonus decimal.Decimal `json:"totalPlatformBonus"` //累计平台补贴
	TotalPoolAmount    decimal.Decimal `json:"totalPoolAmount"`    //累计奖池金额
}
