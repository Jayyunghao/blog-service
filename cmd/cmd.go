package cmd

import (
	"Practice/go-programming-tour-book/blog-service/global"
	"Practice/go-programming-tour-book/blog-service/internal/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
var (
	port         string
	runMode      string
	config       string
	BuildTime    string = "unknown"
	BuildVersion string = "unknown"
	GitCommitID  string = "unknown"
	isVersion    bool
)

var rootCmd = &cobra.Command{
	Use:   "blog_service",
	Short: "support blog service",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := blogService()
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&port, "port", "", "启动端口")
	rootCmd.PersistentFlags().StringVar(&runMode, "mode", "", "启动模式")
	rootCmd.PersistentFlags().StringVar(&config, "config", "configs/", "指定要使用的配置文件目录")
	rootCmd.PersistentFlags().BoolVar(&isVersion, "version", false, "版本信息")
}

func blogService() error {
	if isVersion {
		printVersion()
		return nil
	}

	err := globalInit()
	if err != nil {
		return err
	}
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	//http server
	httpServer := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//https server
	httpsServer := &http.Server{
		Addr:    ":" + global.ServerSetting.HttpsPort,
		Handler: router,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog_service")
	var eg errgroup.Group
	eg.Go(func() error {
		global.Logger.Infof("Start to listening the incoming requests on http address: %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
			return err
		}
		global.Logger.Infof("Server on %s stopped", httpServer.Addr)
		return nil
	})
	eg.Go(func() error {
		if global.ServerSetting.HttpsPort == "" || global.ServerSetting.CertFile == "" || global.ServerSetting.KeyFile == "" {
			return nil
		}
		global.Logger.Infof("Start to listening the incoming requests on http address: %s", httpServer.Addr)
		if err := httpsServer.ListenAndServeTLS(global.ServerSetting.CertFile, global.ServerSetting.KeyFile); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
			return err
		}
		global.Logger.Infof("Server on %s stopped", httpServer.Addr)
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}
	//等待中断信号
	quit := make(chan os.Signal)
	//接收syscall.SIGINT和syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shut down server...")
	//最大时间控制，通知该服务端它有5s的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	if err := httpsServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func printVersion() {
	fmt.Printf("build_time: %s\n", BuildTime)
	fmt.Printf("build_version: %s\n", BuildVersion)
	fmt.Printf("git_commit_id: %s\n", GitCommitID)
}
