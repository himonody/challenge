package lang

import "challenge/config/base/lang"

var en = map[int]string{
	lang.SuccessCode:                     "Operation succeeded",
	lang.RequestErr:                      "Request failed",
	lang.AuthErr:                         "Session expired, please log in again",
	lang.ForbitErr:                       "Insufficient permissions, please contact the administrator",
	lang.ServerErr:                       "Internal error",
	lang.ParamErrCode:                    "Invalid parameters",
	lang.OpErrCode:                       "Operation error, please check",
	lang.DataDecodeCode:                  "Data parsing error",
	lang.DataDecodeLogCode:               "Data parsing error: %s",
	lang.DataQueryCode:                   "Data query failed",
	lang.DataQueryLogCode:                "Data query failed: %s",
	lang.DataInsertLogCode:               "Data insert failed: %s",
	lang.DataInsertCode:                  "Data insert failed",
	lang.DataNotUpdateCode:               "No data updated",
	lang.DataUpdateCode:                  "Data update failed",
	lang.DataUpdateLogCode:               "Data update failed: %s",
	lang.DataDeleteCode:                  "Data delete failed",
	lang.DataDeleteLogCode:               "Data delete failed: %s",
	lang.DataNotFoundCode:                "Data not found",
	lang.ServerErrLogCode:                "Internal error: %s",
	lang.AuthUsernameErrorCode:           "Invalid username format",
	lang.AuthPasswordErrorCode:           "Invalid password format",
	lang.AuthVerificationCodeErrorCode:   "Invalid verification code",
	lang.AuthUserAlreadyExistsCode:       "User already exists",
	lang.AuthInviteCodeNotFoundErrorCode: "Invite code not found",

	// risk
	lang.RiskStrategyNotFoundCode: "No risk strategy available",
	lang.RiskBlacklistHitCode:     "Request blocked by risk blacklist",
}
