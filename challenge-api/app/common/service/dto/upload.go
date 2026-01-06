package dto

import "mime/multipart"

type Upload struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
