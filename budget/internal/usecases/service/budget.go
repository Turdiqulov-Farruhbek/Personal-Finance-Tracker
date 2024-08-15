package service
import (
	"context"

	pb "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"
	"gitlab.com/saladin2098/finance_tracker1/budget/internal/storage"
)

type BudgetService struct {
	stg storage.StorageI
	pb.UnimplementedBudgetServiceServer
}

func NewBudgetService(stg storage.StorageI) *BudgetService {
    return &BudgetService{stg: stg}
}

func (s *BudgetService) CreateBudget(ctx context.Context, req *pb.BudgetCreate) (*pb.Void, error) {
    return s.stg.Budget().CreateBudget(req)
}

func (s *BudgetService) GetBudget(ctx context.Context, req *pb.ById) (*pb.BudgetGet, error) {
    return s.stg.Budget().GetBudget(req)
}

func (s *BudgetService) UpdateBudget(ctx context.Context, req *pb.BudgetUpdate) (*pb.Void, error) {
    return s.stg.Budget().UpdateBudget(req)
}

func (s *BudgetService) DeleteBudget(ctx context.Context, req *pb.ById) (*pb.Void, error) {
    return s.stg.Budget().DeleteBudget(req)
}

func (s *BudgetService) ListBudgets(ctx context.Context, req *pb.BudgetFilter) (*pb.BudgetList, error) {
    return s.stg.Budget().ListBudgets(req)
}