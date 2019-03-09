package util

import (
	"github.com/sirupsen/logrus"
	"github.com/micro/go-config"
)

type (
	MongoDBConfig struct {
		Addr []string `json:"addr"`
		Username string `json:"username"`
		Password string `json:"password"`
		PoolLimit int `json:"poolLimit"`
	}

	WeChatOpenConfig struct {
		URL string `json:"url""`
		Secrets map[string]string `json:"secrets"`
	}


)

var (
	DefaultMgoConf MongoDBConfig
	DefaultWeChatOpenConf WeChatOpenConfig
)

func Init() {
	// 加载mongo配置
	err := config.Get("mongodb").Scan(&DefaultMgoConf)
	if err != nil {
		logrus.Fatalf("get mgo config error: %s", err)
	}

	if len(DefaultMgoConf.Addr) == 0 {
		logrus.Fatalf("invalid mgo addr")
	}

	err = config.Get("wechatOpenConfig").Scan(&DefaultWeChatOpenConf)
	if len(DefaultWeChatOpenConf.URL) == 0 {
		DefaultWeChatOpenConf.URL = "https://api.weixin.qq.com/sns/jscode2session"
	}

	if len(DefaultWeChatOpenConf.Secrets) == 0 {
		logrus.Fatal("wechat secrets is empty")
	}

}