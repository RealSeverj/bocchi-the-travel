package service

import (
	"bocchi/kitex_gen/user"
	"bocchi/pkg/errno"
	"bocchi/pkg/utils/otp2fa"
	"bocchi/pkg/utils/pwd"
	"bocchi/rpc/user/dal/db"
)

func (s *UserService) Register(req *user.RegisterRequest) (*db.User, error) {
	if len(req.Username) < 4 /*||len(req.Password)<8*/ {
		return nil, errno.ParamError
	}

	PwdDigest := pwd.SetPassword(req.Password)
	userModel := &db.User{
		UserName: req.Username,
		Email:    req.Email,
		Password: PwdDigest,
	}
	return db.Register(s.ctx, userModel)
}

func (s *UserService) Login(req *user.LoginRequest) (*db.User, error) {
	userModel := &db.User{
		UserName: req.Username,
		Password: req.Password,
	}
	userResp, err := db.Login(s.ctx, userModel)
	if err != nil {
		return nil, err
	}
	if userResp.Type2fa == 1 && !otp2fa.CheckTotp(*req.Otp, userResp.Otp) {
		return nil, errno.Verify2FAError
	}
	return userResp, nil
}

func (s *UserService) Info(req *user.InfoRequest) (*db.User, error) {
	userModel := &db.User{
		ID: req.UserId,
	}
	return db.QueryUserByID(userModel)
}