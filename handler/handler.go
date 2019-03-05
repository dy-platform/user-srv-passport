package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dy-gopkg/kit"
	"github.com/dy-gopkg/util/format"
	"github.com/dy-platform/user-srv-passport/dal/db"
	"github.com/dy-platform/user-srv-passport/idl"
	snowflake "github.com/dy-platform/user-srv-passport/idl/platform/id/srv-snowflake"
	srv "github.com/dy-platform/user-srv-passport/idl/platform/user/srv-passport"
	"github.com/dy-platform/user-srv-passport/util/config"
	"github.com/dy-ss/crypto/password"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type Handler struct {
}

// 注册
func (h *Handler) SignUp(ctx context.Context, req *srv.SignUpReq, rsp *srv.SignUpResp) error {
	mobile, err := format.Mobile(req.Mobile)
	if err != nil {
		logrus.Debugf("bad mobile, err:%v", err)
		return err
	}

	// 检验手机号是否已被注册

	// TODO 后期再验证短信验证码
	//if !format.IsDigit(req.Code) {
	//	logrus.Debugf("bad code %v", req.Code)
	//	return errors.New("bad code")
	//}

	if err = format.Password(req.Password); err != nil {
		logrus.Debugf("bad password. %v", err)
		return err
	}

	// 生成密码
	passwd, salt := password.Make([]byte(req.Password))

	// 产生一个UID
	cl := snowflake.NewSnowFlakeService("platform.id.srv.snowflake", kit.Client())
	idReq := &snowflake.GetIDReq{Num: 1}
	idRsp, err := cl.GetID(ctx, idReq)
	if err != nil {
		logrus.Warnf("platform.id.srv.snowflake GetID error: %v", err)
		rsp.BaseResp = &base.Resp{Code: int32(base.CODE_SERVICE_EXCEPTION)}
		return nil
	}

	// 插入一条用户通行证信息
	err = db.InsertOneUserPassport(req.DeviceID, idRsp.IDs[0], req.Name, string(passwd), string(salt), mobile, req.Email, "")
	if err != nil {
		logrus.Warnf("db.InsertOneUserPassport error:%v", err)
		rsp.BaseResp = &base.Resp{
			Code: int32(base.CODE_DATA_EXCEPTION),
			Msg:  err.Error(),
		}
		return nil
	}

	rsp.UserID = idRsp.IDs[0]
	rsp.Mobile = mobile
	rsp.BaseResp = &base.Resp{Code: int32(base.CODE_OK)}
	return nil
}

type WeChatOpenCode2SessionResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 微信登陆
func (h *Handler) WeChatSignIn(ctx context.Context, req *srv.WeChatSignInReq, rsp *srv.WeChatSignInResp) error {
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	if len(req.AppID) == 0 || len(req.Secret) == 0 || len(req.Code) == 0 {
		logrus.Warnf("invalid parameter")
		rsp.BaseResp.Code = int32(base.CODE_INVALID_PARAMETER)
		rsp.BaseResp.Msg = "invalid parameter"
		return nil
	}

	secret, ok := uconfig.DefaultWeChatOpenConf.Secrets[req.AppID]
	if !ok {
		logrus.Warnf("invalid appid")
		rsp.BaseResp.Code = int32(base.CODE_INVALID_PARAMETER)
		rsp.BaseResp.Msg = "invalid parameter"
		return
	}

	reqStr := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		uconfig.DefaultWeChatOpenConf.URL, req.AppID, secret, req.Code)
	resp1, err := http.Get(reqStr)
	if err != nil {
		logrus.Warnf("get %s error. %s", reqStr, err)
		rsp.BaseResp.Code = int32(base.CODE_SERVICE_EXCEPTION)
		rsp.BaseResp.Msg = "service exception"
		return nil
	}
	defer resp1.Body.Close()

	body, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		logrus.Warnf("read body from resp error. %s", err)
		rsp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		rsp.BaseResp.Msg = "data exception"
		return nil
	}

	jsonResp := &WeChatOpenCode2SessionResp{}
	err = json.Unmarshal(body, jsonResp)

	if jsonResp.ErrCode != 0 {
		logrus.Warnf("request wechat open api return errcode:%d", jsonResp.ErrCode)
		rsp.BaseResp.Code = int32(base.CODE_FAILED)
		return nil
	}

	u, err := db.GetPassportByWeChatID(jsonResp.UnionID)
	if err != nil {
		// 没有找到
		if err == mgo.ErrNotFound {
			// 请求一个uid
			cl := snowflake.NewSnowFlakeService("platform.id.srv.snowflake", kit.Client())
			idReq := &snowflake.GetIDReq{Num: 1}
			idRsp, err := cl.GetID(ctx, idReq)
			if err != nil {
				logrus.Warnf("platform.id.srv.snowflake GetID error: %v", err)
				rsp.BaseResp = &base.Resp{Code: int32(base.CODE_SERVICE_EXCEPTION)}
				return nil
			}

			u.UID = idRsp.IDs[0]
			// 插入一条用户通行证信息
			err = db.InsertOneUserPassport(req.DeviceID, idRsp.IDs[0], "", "", "", "", "", jsonResp.UnionID)
			if err != nil {
				logrus.Warnf("db.InsertOneUserPassport error:%v", err)
				rsp.BaseResp = &base.Resp{
					Code: int32(base.CODE_DATA_EXCEPTION),
					Msg:  err.Error(),
				}
				return nil
			}
		} else {
			logrus.Warnf("mgo error:%v", err)
			rsp.BaseResp.Code = int32(base.CODE_FAILED)
			rsp.BaseResp.Msg = err.Error()
			return nil
		}
	}

	rsp.UserID = u.UID

	return nil
}

// 登陆
func (h *Handler) MobileSignIn(ctx context.Context, req *srv.MobileSignInReq, rsp *srv.MobileSignInResp) error {

	return nil
}

func (h *Handler) EmailSignIn(ctx context.Context, req *srv.EmailSignInReq, rsp *srv.EmailSignInResp) error {

	return nil
}

func (h *Handler) UserNameSignIn(ctx context.Context, req *srv.UserNameSignInReq, rsp *srv.UserNameSignInResp) error {

	return nil
}

func (h *Handler) TokenSignIn(ctx context.Context, req *srv.TokenSignInReq, rsp *srv.TokenSignInResp) error {

	return nil
}

func (h *Handler) SignOut(ctx context.Context, req *srv.SignOutReq, rsp *srv.SignOutResp) error {

	return nil
}

func (h *Handler) ChangePassword(ctx context.Context, req *srv.ChangePasswordReq, rsp *srv.ChangePasswordResp) error {

	return nil
}

func (h *Handler) ResetPassword(ctx context.Context, req *srv.ResetPasswordReq, rsp *srv.ResetPasswordResp) error {

	return nil
}

func (h *Handler) BindMobile(ctx context.Context, req *srv.BindMobileReq, rsp *srv.BindMobileResp) error {

	return nil
}

func (h *Handler) UnbindMobile(ctx context.Context, req *srv.UnbindMobileReq, rsp *srv.UnbindMobileResp) error {

	return nil
}
