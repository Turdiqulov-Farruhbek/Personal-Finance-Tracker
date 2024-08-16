package repo

import (
	"finance_tracker/budget/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	AccountS     storage.AccountI
	BudgetS      storage.BudgetI
	TransactionS storage.TransactionI
	CategoryS    storage.CategoryI
	GoalS        storage.GoalI
	mdb          *mongo.Database
}

func NewStorage(mdb *mongo.Database) *Storage {
	return &Storage{
		AccountS:     NewAccountRepo(mdb),
		BudgetS:      NewBudgetRepo(mdb),
		TransactionS: NewTransactionRepo(mdb),
		CategoryS:    NewCategoryRepo(mdb),
		GoalS:        NewGoalRepo(mdb),
		mdb:          mdb,
	}
}
func (s *Storage) Account() storage.AccountI {
	if s.AccountS == nil {
		s.AccountS = NewAccountRepo(s.mdb)
	}
	return s.AccountS
}
func (s *Storage) Budget() storage.BudgetI {
	if s.BudgetS == nil {
		s.BudgetS = NewBudgetRepo(s.mdb)
	}
	return s.BudgetS
}
func (s *Storage) Transaction() storage.TransactionI {
	if s.TransactionS == nil {
		s.TransactionS = NewTransactionRepo(s.mdb)
	}
	return s.TransactionS
}
func (s *Storage) Category() storage.CategoryI {
	if s.CategoryS == nil {
		s.CategoryS = NewCategoryRepo(s.mdb)
	}
	return s.CategoryS
}
func (s *Storage) Goal() storage.GoalI {
	if s.GoalS == nil {
		s.GoalS = NewGoalRepo(s.mdb)
	}
	return s.GoalS
}
