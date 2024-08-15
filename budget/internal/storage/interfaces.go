package storage

import pb "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"

type StorageI interface {
	Account() AccountI
	Budget() BudgetI
	Transaction() TransactionI
	Category() CategoryI
	Goal() GoalI
}

type AccountI interface {
	CreateAccount(req *pb.AccountCreate) (*pb.Void, error)
	GetAccount(req *pb.ById) (*pb.AccountGet, error)
	UpdateAccount(req *pb.AccountUpdate) (*pb.Void, error)
	DeleteAccount(req *pb.ById) (*pb.Void, error)
	ListAccounts(req *pb.AccountFilter) (*pb.AccounList ,error)
}

type BudgetI interface {
	CreateBudget(req *pb.BudgetCreate) (*pb.Void, error)
	UpdateBudget(req *pb.BudgetUpdate) (*pb.Void, error) 
	DeleteBudget(req *pb.ById) (*pb.Void, error) 
	GetBudget(req *pb.ById) (*pb.BudgetGet, error) 
	ListBudgets(req *pb.BudgetFilter) (*pb.BudgetList, error) 
}
type CategoryI interface {
	CreateCategory(req *pb.CategoryCreate) (*pb.Void, error) 
	UpdateCategory(req *pb.CategoryUpdate) (*pb.Void, error) 
	DeleteCategory(req *pb.ById) (*pb.Void, error) 
	GetCategory(req *pb.ById) (*pb.CategoryGet, error)
	ListCategories(req *pb.CategoryFilter) (*pb.CategoryList, error) 
}
type TransactionI interface {
	CreateTransaction(req *pb.TransactionCreate) (*pb.Void, error) 
	UpdateTransaction(req *pb.TransactionUpdate) (*pb.Void, error) 
	DeleteTransaction(req *pb.ById) (*pb.Void, error) 
	GetTransaction(req *pb.ById) (*pb.TransactionGet, error)
	ListTransactions(req *pb.TransactionFilter) (*pb.TransactionList, error) 
}
type GoalI interface {
	CreateGoal(req *pb.GoalCreate) (*pb.Void, error) 
	UpdateGoal(req *pb.GoalUpdate) (*pb.Void, error) 
	DeleteGoal(req *pb.ById) (*pb.Void, error) 
	GetGoal(req *pb.ById) (*pb.GoalGet, error) 
	ListGoals(req *pb.GoalFilter) (*pb.GoalList, error) 
}