package cmd

import (
	"Practice/go-programming-tour-book/blog-service/global"
	"Practice/go-programming-tour-book/blog-service/internal/model"
	"Practice/go-programming-tour-book/blog-service/pkg/logger"
	"Practice/go-programming-tour-book/blog-service/pkg/setting"
	"Practice/go-programming-tour-book/blog-service/pkg/tracer"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"strings"
	"time"
)


func globalInit() error {
	err := setupSetting()
	if err != nil {
		return fmt.Errorf("init.setupSetting err :%v", err)
	}
	err = setupLogger()
	if err != nil {
		return fmt.Errorf("init.setupLogger err :%v", err)
	}
	err = setupDBEngine()
	if err != nil {
		return fmt.Errorf("init.setupDBEngine err:%v", err)
	}
	err = setupTracer()
	if err != nil {
		return fmt.Errorf("init.setupTracer err:%v", err)
	}
	//setupFlag()
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config,",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	fmt.Printf("\nserverSetting : %v\nappSetting : %v\nDatabase: %v\nJWT: %v\nEmail:%v\n", global.ServerSetting, global.AppSetting, global.DatabaseSetting, global.AppSetting, global.EmailSetting)
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   60000,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2) //第二层调用栈信息是什么意思？
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service","127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}