package uconfig

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/dy-gopkg/kit/util"
	"github.com/micro/go-config/source/consul"
	"github.com/micro/go-config/source/file"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)
type UserPassportConfig struct {
	UserPassportDB struct {
		ReadAddress string `json:"ReadAddress"`
		WriteAddress string `json:"WriteAddress"`
		User string `json:"User"`
		Password string `json:"Password"`
		DBName string `json:"DBName"`
		ReadTimeout string `json:"ReadTimeout"`
		WriteTimeout string `json:"WriteTimeout"`
	}

	UserPassportCache struct {
		Address string `json:"Address"`
		Password string `json:"Password"`
		DialTimeout time.Duration `json:"DialTimeout"`
		ReadTimeout time.Duration `json:"ReadTimeout"`
		WriteTimeout time.Duration `json:"WriteTimeout"`
	}
}

var (
	DefaultConfig UserPassportConfig
)

func InitBusinessConfig() {

	if len(util.BaseConf.BusinessConfig.Path) ==0 {
		logrus.Info("business config not set")
		return
	}
	err := config.Load(file.NewSource(file.WithPath("business_config.json")))
	if err != nil {
		logrus.Debugf("err:", err)

		// 从远程加载业务配置
		if len(util.BaseConf.BusinessConfig.Addr) == 0 {
			logrus.Warnf("remote business config address not set")
			return
		}
		err = config.Load(consul.NewSource(consul.WithAddress(util.BaseConf.ServiceConfig.Addr)))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	err = config.Get(util.BaseConf.ServiceConfig.Path).Scan(&DefaultConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// mysql redis 的配置不监听变更
}