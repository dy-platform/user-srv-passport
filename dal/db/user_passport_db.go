package db

import (
	"fmt"
	"github.com/dy-platform/user-srv-passport/util/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

type UserPassportDB struct {
	rClient *gorm.DB
	wClient *gorm.DB

	sync.RWMutex
}

var (
	db           = &UserPassportDB{}
	dbArgsFormat = "%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&maxAllowedPacket=536870912"
)

func DBInit() {
	//logrus.Infof()
	var err error
	rArgs := fmt.Sprintf(dbArgsFormat,
		uconfig.DefaultConfig.UserPassportDB.User,
		uconfig.DefaultConfig.UserPassportDB.Password,
		uconfig.DefaultConfig.UserPassportDB.ReadAddress,
		uconfig.DefaultConfig.UserPassportDB.DBName,
		uconfig.DefaultConfig.UserPassportDB.ReadTimeout,
		uconfig.DefaultConfig.UserPassportDB.WriteTimeout)

	db.rClient, err = gorm.Open("mysql", rArgs)
	if err != nil {
		logrus.Warnf("open read-mysqldb failed. args:%s", rArgs)
		os.Exit(1)
	}

	wArgs := fmt.Sprintf(dbArgsFormat,
		uconfig.DefaultConfig.UserPassportDB.User,
		uconfig.DefaultConfig.UserPassportDB.Password,
		uconfig.DefaultConfig.UserPassportDB.ReadAddress,
		uconfig.DefaultConfig.UserPassportDB.DBName,
		uconfig.DefaultConfig.UserPassportDB.ReadTimeout,
		uconfig.DefaultConfig.UserPassportDB.WriteTimeout)

	db.wClient, err = gorm.Open("mysql", wArgs)
	if err != nil {
		logrus.Warnf("open write-mysqldb failed. args:%s", wArgs)
		os.Exit(1)
	}

	// TODO PING

	db.wClient.DB().SetMaxIdleConns(10)
	db.wClient.DB().SetMaxOpenConns(20)
}

type UserAuth struct {
	Uid int64 `gorm:"column:uid"`
	Name         string `gorm:"column:name"`
	Mobile       string `gorm:"column:mobile"`
	Email        string `gorm:"column:email"`
	Password     string `gorm:"column:password"`
	Salt         string `gorm:"column:salt"`
	UserStatus   int32  `gorm:"column:user_status"`
	AppID        int32  `gorm:"column:appid"`
	WechatOpenID string `gorm:"column:wechat_openid"`
	QQOpenID     string `gorm:"column:qq_openid"`
	WeiboOpenID  string `gorm:"column:weibo_openid"`
}

func UpdateUserAuth(ua *UserAuth) {
	db.wClient.Create(ua)
}
