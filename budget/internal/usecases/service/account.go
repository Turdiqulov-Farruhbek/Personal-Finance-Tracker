package service

import (
	"context"

	pb "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"
	"gitlab.com/saladin2098/finance_tracker1/budget/internal/storage"
)

type AccountService struct {
	stg storage.StorageI
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(stg storage.StorageI) *AccountService {
	return &AccountService{stg: stg}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.AccountCreate) (*pb.Void, error) {
	return s.stg.Account().CreateAccount(req)
}

func (s *AccountService) GetAccount(ctx context.Context, req *pb.ById) (*pb.AccountGet, error) {
    return s.stg.Account().GetAccount(req)
}

func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.AccountUpdate) (*pb.Void, error) {
    return s.stg.Account().UpdateAccount(req)
}

func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.ById) (*pb.Void, error) {
    return s.stg.Account().DeleteAccount(req)
}

func (s *AccountService) ListAccounts(ctx context.Context, req *pb.AccountFilter) (*pb.AccounList, error) {
    return s.stg.Account().ListAccounts(req)
}