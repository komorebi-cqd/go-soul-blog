package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-soul-blog/global"
	"github.com/go-soul-blog/internal/model"
	"github.com/go-soul-blog/internal/routers"
	"github.com/go-soul-blog/pkg/setting"
)

func init() {
	// 初始化设置
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// 连接数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// 设置配置信息
func setupSetting() error {
	setting, err := setting.NewSetting()
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

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// 连接数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	fmt.Println("global.DBEngine \n", global.DBEngine)
	if err != nil {
		return err
	}
	return nil
}

// @title 博客系统
// @version 1.0
// @description 博客系统
func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
	}

	s.ListenAndServe()
}
