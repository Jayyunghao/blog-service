package api

import (
	"Practice/go-programming-tour-book/blog-service/global"
	"Practice/go-programming-tour-book/blog-service/internal/service"
	"Practice/go-programming-tour-book/blog-service/pkg/app"
	"Practice/go-programming-tour-book/blog-service/pkg/convert"
	"Practice/go-programming-tour-book/blog-service/pkg/errcode"
	"Practice/go-programming-tour-book/blog-service/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.Upload err :%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
