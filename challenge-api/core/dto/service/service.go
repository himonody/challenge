package service

import (
	"challenge/core/runtime"

	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logbase.Helper
	Lang  string //语言 en 英文 zh-cn中文
	Run   runtime.Runtime
	C     *gin.Context
}
