package db

import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/dy-platform/user-srv-passport/util"
	"github.com/sirupsen/logrus"
)

const (
	DBUserPassport = "dayan_user"
)

var (
	defaultMgo *mgo.Session
)

func Mgo() *mgo.Session {
	return defaultMgo
}

func InitMgo() {
	dialInfo := &mgo.DialInfo{
		Addrs:     util.DefaultMgoConf.Addr,
		Direct:    false,
		Timeout:   time.Second * 3,
		PoolLimit: util.DefaultMgoConf.PoolLimit,
		Username:  util.DefaultMgoConf.Username,
		Password:  util.DefaultMgoConf.Password,
	}

	ses, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatalf("dail mgo server error:%v", err)
	}

	ses.SetMode(mgo.Monotonic, true)
	defaultMgo = ses
}
