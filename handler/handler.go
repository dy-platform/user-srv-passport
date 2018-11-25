package handler

import (
	"github.com/dy-gopkg/util/format"
	srv "github.com/dy-platform/user-srv-passport/idl/platform/user/srv-passport"
	"github.com/dy-ss/crypto/password"
	"github.com/sirupsen/logrus"
	"context"
	"github.com/jinzhu/gorm"
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

	// TODO  后期再检测短信验证码

	// 生成密码
	passwd, salt := password.Make([]byte(req.Password))





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
