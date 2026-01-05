package lang

import "challenge/config/base/lang"

var id = map[int]string{
	lang.SuccessCode:       "Berhasil",
	lang.RequestErr:        "Permintaan gagal",
	lang.AuthErr:           "Sesi kedaluwarsa, silakan masuk lagi",
	lang.ForbitErr:         "Izin tidak cukup, silakan hubungi administrator",
	lang.ServerErr:         "Kesalahan internal",
	lang.ParamErrCode:      "Parameter tidak valid",
	lang.OpErrCode:         "Kesalahan operasi, silakan periksa",
	lang.DataDecodeCode:    "Kesalahan parsing data",
	lang.DataDecodeLogCode: "Kesalahan parsing data: %s",
	lang.DataQueryCode:     "Kueri data gagal",
	lang.DataQueryLogCode:  "Kueri data gagal: %s",
	lang.DataInsertLogCode: "Gagal memasukkan data: %s",
	lang.DataInsertCode:    "Gagal memasukkan data",
	lang.DataNotUpdateCode: "Tidak ada data yang diperbarui",
	lang.DataUpdateCode:    "Gagal memperbarui data",
	lang.DataUpdateLogCode: "Gagal memperbarui data: %s",
	lang.DataDeleteCode:    "Gagal menghapus data",
	lang.DataDeleteLogCode: "Gagal menghapus data: %s",
	lang.DataNotFoundCode:  "Data tidak ditemukan",
	lang.ServerErrLogCode:  "Kesalahan internal: %s",

	lang.AuthUsernameErrorCode:           "Format nama pengguna tidak valid",
	lang.AuthPasswordErrorCode:           "Format kata sandi tidak valid",
	lang.AuthVerificationCodeErrorCode:   "Kode verifikasi tidak valid",
	lang.AuthUserAlreadyExistsCode:       "Pengguna sudah ada",
	lang.AuthInviteCodeNotFoundErrorCode: "Kode undangan tidak ditemukan",
}
