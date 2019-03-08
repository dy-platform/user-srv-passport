package main

import (
	"github.com/dy-gopkg/kit/micro"
	"github.com/dy-platform/user-srv-passport/dal/db"
	"github.com/dy-platform/user-srv-passport/handler"
	"github.com/dy-platform/user-srv-passport/idl/platform/user/srv-passport"
	"github.com/dy-platform/user-srv-passport/util/config"
	"github.com/sirupsen/logrus"
)

func main() {
	micro.Init()

	// 初始化配置
	uconfig.Init()

	// 初始化数据库
	db.Init()

	//TODO 初始化缓存
	//cache.CacheInit()

	err := platform_user_srv_passport.RegisterPassportHandler(micro.DefaultService.Server(), &handler.Handler{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}

	micro.Run()
}