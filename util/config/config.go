package uconfig

import (
	"github.com/Sirupsen/logrus"
	"github.com/micro/go-config"
)

type (
	MongoDBConfig struct {
		Addr []string
		Username string
		Password string
		PoolLimit int
	}

)

var (
	DefaultMgoConf MongoDBConfig
	DefaultWeChatOpenURL string
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


	DefaultWeChatOpenURL = config.Get("wechatOpenURL").String("https://api.weixin.qq.com/sns/jscode2session")
	if len(DefaultWeChatOpenURL) == 0 {
		logrus.Fatalf("wechatOpenURL is empty")
	}
}