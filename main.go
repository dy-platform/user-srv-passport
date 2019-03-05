package main

import (
	"github.com/dy-platform/user-srv-passport/dal/db"
	"github.com/dy-platform/user-srv-passport/handler"
	"github.com/dy-platform/user-srv-passport/idl/platform/user/srv-passport"
	"github.com/dy-gopkg/kit"
	"github.com/dy-platform/user-srv-passport/util/config"
)

func main() {
	kit.Init()

	// 初始化配置
	uconfig.Init()

	// 初始化数据库
	db.Init()

	//TODO 初始化缓存
	//cache.CacheInit()

	platform_user_srv_passport.RegisterPassportHandler(kit.DefaultService.Server(), &handler.Handler{})

	kit.Run()
}