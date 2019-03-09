package db

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/dy-platform/user-srv-passport/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	defaultMgo *mgo.Session
)

func Mgo() *mgo.Session {
	return defaultMgo
}

func Init() {
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

type UserPassport struct {
	UID       int64  `bson:"_id"`
	DeviceID  int64  `bson:"device_id"`
	Name      string `bson:"name"`
	Password  string `bson:"password"`
	Salt      string `bson:"salt"`
	Mobile    string `bson:"mobile"`
	Email     string `bson:"email"`
	WeChatID  string `bson:"wechat_id"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

var (
	DBUserPassport = "dayan_user_passport"
	CUserPassport  = "user_passport"
)

func InsertOneUserPassport(devID, uid int64, name, password, salt, mobile, email, weChatID string) error {
	u := &UserPassport{
		UID:        uid,
		DeviceID:  devID,
		Name:      name,
		Password:  password,
		Salt:      salt,
		Mobile:    mobile,
		Email:     email,
		WeChatID:  weChatID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	ses := defaultMgo.Copy()
	defer ses.Close()

	if ses == nil {
		//logrus.Warnf("mgo session is nil")
		return errors.New("mgo session is nil")
	}

	return ses.DB(DBUserPassport).C(CUserPassport).Insert(u)
}

func Count(weChatID string) (int, error) {
	query := bson.M{
		"wechat_id":weChatID,
	}
	ses := defaultMgo.Copy()
	defer ses.Close()

	if ses == nil {
		//logrus.Warnf("mgo session is nil")
		return 0, errors.New("mgo session is nil")
	}

	return ses.DB(DBUserPassport).C(CUserPassport).Find(query).Count()
}

func GetPassportByWeChatID(weChatID string) (*UserPassport, error) {
	query := bson.M{
		"wechat_id":weChatID,
	}

	ses := defaultMgo.Copy()
	defer ses.Close()

	if ses == nil {
		//logrus.Warnf("mgo session is nil")
		return nil, errors.New("mgo session is nil")
	}
	ret := &UserPassport{}
	err := ses.DB(DBUserPassport).C(CUserPassport).Find(query).One(ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}