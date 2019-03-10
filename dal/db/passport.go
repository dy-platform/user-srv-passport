package db

import (
	"errors"
	"github.com/dy-platform/user-srv-passport/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	CUserPassport  = "user_passport"
)

func InsertUserPassport(u *model.UserPassport) error {
	ses := defaultMgo.Copy()
	defer ses.Close()

	if ses == nil {
		return errors.New("mgo session is nil")
	}

	now := time.Now().Unix()
	u.UpdatedAt = now
	u.CreatedAt = now

	return ses.DB(DBUserPassport).C(CUserPassport).Insert(u)
}

func Count(weChatID string) (int, error) {
	ses := defaultMgo.Copy()
	defer ses.Close()

	if ses == nil {
		return 0, errors.New("mgo session is nil")
	}

	query := bson.M{
		"wechat_id":weChatID,
	}

	return ses.DB(DBUserPassport).C(CUserPassport).Find(query).Count()
}

func GetPassportByWeChatID(weChatID string) (*model.UserPassport, error) {
	query := bson.M{
		"wechat_id":weChatID,
	}

	ses := defaultMgo.Copy()
	defer ses.Close()

	if ses == nil {
		return nil, errors.New("mgo session is nil")
	}
	ret := &model.UserPassport{}
	err := ses.DB(DBUserPassport).C(CUserPassport).Find(query).One(ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}