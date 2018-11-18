package handler

import (
	"errors"
	"github.com/dy-gopkg/util/format"
	srv "github.com/dy-platform/user-srv-passport/idl/platform/user/srv-passport"
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

	if !format.IsDigit(req.Code) {
		logrus.Debugf("bad code %v", req.Code)
		return errors.New("bad code")
	}

	if err = format.Password(req.Password); err != nil {
		logrus.Debugf("bad password. %v", err)
		return err
	}

	// TODO 监测 短信验证码


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
