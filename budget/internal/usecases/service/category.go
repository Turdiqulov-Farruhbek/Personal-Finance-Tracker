package service
import (
	"context"

	pb "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"
	"gitlab.com/saladin2098/finance_tracker1/budget/internal/storage"
)

type CategoryService struct {
    stg storage.StorageI
    pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(stg storage.StorageI) *CategoryService {
    return &CategoryService{stg: stg}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CategoryCreate) (*pb.Void, error) {
    return s.stg.Category().CreateCategory(req)
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.ById) (*pb.CategoryGet, error) {
    return s.stg.Category().GetCategory(req)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.CategoryUpdate) (*pb.Void, error) {
    return s.stg.Category().UpdateCategory(req)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.ById) (*pb.Void, error) {
    return s.stg.Category().DeleteCategory(req)
}

func (s *CategoryService) ListCategories(ctx context.Context, req *pb.CategoryFilter) (*pb.CategoryList, error) {
    return s.stg.Category().ListCategories(req)
}