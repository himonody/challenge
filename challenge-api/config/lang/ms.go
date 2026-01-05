package lang

import "challenge/config/base/lang"

var ms = map[int]string{
	lang.SuccessCode:       "Berjaya",
	lang.RequestErr:        "Permintaan gagal",
	lang.AuthErr:           "Sesi tamat, sila log masuk semula",
	lang.ForbitErr:         "Kebenaran tidak mencukupi, sila hubungi pentadbir",
	lang.ServerErr:         "Ralat dalaman",
	lang.ParamErrCode:      "Parameter tidak sah",
	lang.OpErrCode:         "Ralat operasi, sila semak",
	lang.DataDecodeCode:    "Ralat nyahkod data",
	lang.DataDecodeLogCode: "Ralat nyahkod data: %s",
	lang.DataQueryCode:     "Kueri data gagal",
	lang.DataQueryLogCode:  "Kueri data gagal: %s",
	lang.DataInsertLogCode: "Gagal memasukkan data: %s",
	lang.DataInsertCode:    "Gagal memasukkan data",
	lang.DataNotUpdateCode: "Tiada data dikemas kini",
	lang.DataUpdateCode:    "Gagal mengemas kini data",
	lang.DataUpdateLogCode: "Gagal mengemas kini data: %s",
	lang.DataDeleteCode:    "Gagal memadam data",
	lang.DataDeleteLogCode: "Gagal memadam data: %s",
	lang.DataNotFoundCode:  "Data tidak ditemui",
	lang.ServerErrLogCode:  "Ralat dalaman: %s",

	lang.AuthUsernameErrorCode:           "Format nama pengguna tidak sah",
	lang.AuthPasswordErrorCode:           "Format kata laluan tidak sah",
	lang.AuthVerificationCodeErrorCode:   "Kod pengesahan tidak sah",
	lang.AuthUserAlreadyExistsCode:       "Pengguna sudah wujud",
	lang.AuthInviteCodeNotFoundErrorCode: "Kod jemputan tidak ditemui",
}
