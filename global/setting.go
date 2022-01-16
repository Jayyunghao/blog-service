package global

import (
	"Practice/go-programming-tour-book/blog-service/pkg/logger"
	"Practice/go-programming-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettinsS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)
