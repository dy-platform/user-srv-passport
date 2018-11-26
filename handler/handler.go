package handler

import (
	"github.com/dy-gopkg/kit"
	"github.com/dy-gopkg/util/format"
	"github.com/dy-platform/user-srv-passport/dal/db"
	"github.com/dy-platform/user-srv-passport/idl"
	snowflake "github.com/dy-platform/user-srv-passport/idl/platform/id/srv-snowflake"
	srv "github.com/dy-platform/user-srv-passport/idl/platform/user/srv-passport"
	"github.com/dy-ss/crypto/password"
	"github.com/sirupsen/logrus"
	"context"
)

type Handler struct {

}

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

	cl := snowflake.NewSnowFlakeService("platform.id.srv.snowflake", kit.DefaultService.Client())
	idReq := &snowflake.GetIDReq{Num:1}
	idRsp, err := cl.GetID(ctx, idReq)
	if err != nil {
		logrus.Warnf("platform.id.srv.snowflake GetID error: %v", err)
		rsp.BaseResp = &base.Resp{Code:uint32(base.CODE_SERVICE_EXCEPTION)}
		return nil
	}

	ua := db.UserAuth{
		Uid:idRsp.IDs[0],
		Name:         req.Name,
		Mobile:       mobile,
		Email:        req.Email,
		Password:     string(passwd),
		Salt:         string(salt),
		UserStatus:   0,
		AppID:        req.AppID,
		WechatOpenID: "",
		QQOpenID:     "",
		WeiboOpenID:  "",
	}
	db.UpdateUserAuth(&ua)

	rsp.UserID = idRsp.IDs[0]
	rsp.Mobile = mobile
	rsp.BaseResp = &base.Resp{Code:uint32(base.CODE_SUCESS)}
	return nil
}

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
