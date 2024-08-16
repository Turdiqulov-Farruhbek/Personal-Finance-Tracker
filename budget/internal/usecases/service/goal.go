package service

import (
	"context"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"finance_tracker/budget/internal/storage"
)

type GoalService struct {
	stg storage.StorageI
	pb.UnimplementedGoalServiceServer
}

func NewGoalService(stg storage.StorageI) *GoalService {
	return &GoalService{stg: stg}
}
func (s *GoalService) CreateGoal(ctx context.Context, req *pb.GoalCreate) (*pb.Void, error) {
	return s.stg.Goal().CreateGoal(req)
}

func (s *GoalService) GetGoal(ctx context.Context, req *pb.ById) (*pb.GoalGet, error) {
	return s.stg.Goal().GetGoal(req)
}

func (s *GoalService) UpdateGoal(ctx context.Context, req *pb.GoalUpdate) (*pb.Void, error) {
	return s.stg.Goal().UpdateGoal(req)
}

func (s *GoalService) DeleteGoal(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return s.stg.Goal().DeleteGoal(req)
}

func (s *GoalService) ListGoals(ctx context.Context, req *pb.GoalFilter) (*pb.GoalList, error) {
	return s.stg.Goal().ListGoals(req)
}
