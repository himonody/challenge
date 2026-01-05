package lang

import "challenge/config/base/lang"

var ta = map[int]string{
	lang.SuccessCode:       "செயல் வெற்றிகரமாக முடிந்தது",
	lang.RequestErr:        "கோரிக்கை தோல்வியடைந்தது",
	lang.AuthErr:           "அமர்வு காலாவதியானது, தயவுசெய்து மீண்டும் உள்நுழைக",
	lang.ForbitErr:         "அனுமதி போதவில்லை, நிர்வாகியை தொடர்பு கொள்ளவும்",
	lang.ServerErr:         "உள் பிழை",
	lang.ParamErrCode:      "தவறான அளவுரு",
	lang.OpErrCode:         "செயல்பாட்டுப் பிழை, தயவுசெய்து சரிபார்க்கவும்",
	lang.DataDecodeCode:    "தரவைப் பகுப்பாய்வு செய்யும்போது பிழை",
	lang.DataDecodeLogCode: "தரவைப் பகுப்பாய்வு செய்யும்போது பிழை: %s",
	lang.DataQueryCode:     "தரவு விசாரணை தோல்வியடைந்தது",
	lang.DataQueryLogCode:  "தரவு விசாரணை தோல்வியடைந்தது: %s",
	lang.DataInsertLogCode: "தரவைச் சேர்க்க முடியவில்லை: %s",
	lang.DataInsertCode:    "தரவைச் சேர்க்க முடியவில்லை",
	lang.DataNotUpdateCode: "தரவு புதுப்பிக்கப்படவில்லை",
	lang.DataUpdateCode:    "தரவைப் புதுப்பிப்பதில் தோல்வி",
	lang.DataUpdateLogCode: "தரவைப் புதுப்பிப்பதில் தோல்வி: %s",
	lang.DataDeleteCode:    "தரவை நீக்க முடியவில்லை",
	lang.DataDeleteLogCode: "தரவை நீக்க முடியவில்லை: %s",
	lang.DataNotFoundCode:  "தரவு கிடைக்கவில்லை",
	lang.ServerErrLogCode:  "உள் பிழை: %s",

	lang.AuthUsernameErrorCode:           "பயனர் பெயர் வடிவம் தவறானது",
	lang.AuthPasswordErrorCode:           "கடவுச்சொல் வடிவம் தவறானது",
	lang.AuthVerificationCodeErrorCode:   "சரிபார்ப்பு குறியீடு தவறானது",
	lang.AuthUserAlreadyExistsCode:       "பயனர் ஏற்கனவே உள்ளது",
	lang.AuthInviteCodeNotFoundErrorCode: "அழைப்புக் குறியீடு கிடைக்கவில்லை",
}
