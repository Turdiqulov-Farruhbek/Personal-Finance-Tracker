package service

import (
	"context"

	pb "finance_tracker/auth_service/genproto"
	"finance_tracker/auth_service/storage"
)

type UserService struct {
	stg storage.StorageI
	pb.UnimplementedAuthServiceServer
}

func NewUserService(stg storage.StorageI) *UserService {
	return &UserService{stg: stg}
}
func (s *UserService) Register(ctx context.Context, req *pb.UserCreateReq) (*pb.Void, error) {
	return s.stg.User().RegisterUser(req)
}
func (s *UserService) RegisterCourier(ctx context.Context, req *pb.UserCreateReq) (*pb.Void, error) {
	return s.stg.User().RegisterUser(req)
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.Token, error) {
	return s.stg.User().Login(req)
}
func (s *UserService) RefreshToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	return s.stg.User().RefreshToken(req)
}
func (s *UserService) UpdateProfile(ctx context.Context, req *pb.UserUpdate) (*pb.Void, error) {
	return s.stg.User().UpdateProfile(req)
}
func (s *UserService) GetUserProfile(ctx context.Context, req *pb.ById) (*pb.UserCreateRes, error) {
	return s.stg.User().GetUserProfile(req)
}
func (s *UserService) ChangePassword(ctx context.Context, req *pb.PasswordChangeReq) (*pb.Void, error) {
	return s.stg.User().ChangePassword(req)
}
func (s *UserService) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.Void, error) {
	return s.stg.User().ForgotPassword(req)
}
func (s *UserService) ResetPassword(ctx context.Context, req *pb.PasswordResetReq) (*pb.Void, error) {
	return s.stg.User().ResetPassword(req)
}
func (s *UserService) ConfirmEmail(ctx context.Context, req *pb.EmailConfirm) (*pb.Void, error) {
	return s.stg.User().ConfirmEmail(req)
}
func (s *UserService) ResendCode(ctx context.Context, req *pb.ResendReq) (*pb.Void, error) {
	return s.stg.User().ResendCode(req)
}
