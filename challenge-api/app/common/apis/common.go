package apis

import (
	"challenge/app/common/service"
	"challenge/app/common/service/dto"
	l "challenge/config/base/lang"
	"challenge/config/lang"
	"challenge/core/config"
	"challenge/core/dto/api"
	"challenge/core/storage/oss"
	"io"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Common struct {
	api.Api
}

func (co *Common) Captcha(c *gin.Context) {
	s := &service.Common{}
	err := co.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		co.Error(http.StatusBadRequest, "captcha generate failed")
		return
	}
	resp, err := s.Captcha()
	if err != nil {
		co.Error(http.StatusInternalServerError, "captcha generate failed")
		return
	}
	co.OK(resp, lang.Get(co.Lang, l.SuccessCode))
}
func (co *Common) Upload(c *gin.Context) {
	req := new(dto.Upload)
	s := service.Common{}
	err := co.MakeContext(c).
		MakeOrm().
		Bind(req, binding.FormMultipart).
		MakeService(&s.Service).
		Errors
	if err != nil {
		co.Error(http.StatusBadRequest, "captcha generate failed")
		return
	}

	fh := req.File
	if fh == nil {
		co.Error(http.StatusBadRequest, "missing file")
		return
	}
	if fh.Size > 5*1024*1024 { // 5MB，按需调整
		co.Error(http.StatusBadRequest, "file too large")
		return
	}

	// 打开流并校验 MIME
	f, err := fh.Open()
	if err != nil {
		co.Error(http.StatusBadRequest, "open file failed")
		return
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, _ := f.Read(buf)
	ct := http.DetectContentType(buf[:n])
	allowed := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}
	if !allowed[ct] {
		co.Error(http.StatusBadRequest, "unsupported image type")
		return
	}
	// 重置读取指针
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		co.Error(http.StatusInternalServerError, "seek file failed")
		return
	}

	store, err := oss.New(oss.MinIO, *config.OssConfig)
	if err != nil {
		co.Error(http.StatusInternalServerError, "init storage failed")
		return
	}

	objectKey := path.Join("uploads/images", fh.Filename)

	if err := store.UploadReader(objectKey, f, fh.Size, ct); err != nil {
		co.Error(http.StatusInternalServerError, "upload failed")
		return
	}
	//写入数据库

	//url, err := store.GeneratePresignedURL(objectKey, 10*time.Minute)
	//if err != nil {
	//	co.Error(http.StatusInternalServerError, "generate url failed")
	//	return
	//}
}
