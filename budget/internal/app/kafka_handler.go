package app

import (
	"context"
	"log"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"finance_tracker/budget/internal/usecases/service"
	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	account     *service.AccountService
	budget      *service.BudgetService
	category    *service.CategoryService
	goal        *service.GoalService
	transaction *service.TransactionService
}

func (h *KafkaHandler) AccountCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.AccountCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.account.CreateAccount(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create Account via Kafka: %v", err)
			return
		}
		log.Printf("Created Account: %+v", res)
	}
}
func (h *KafkaHandler) UpdateCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.AccountUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.account.UpdateAccount(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update account via Kafka: %v", err)
			return
		}
		log.Printf("Updated Account: %+v", res)
	}
}
func (h *KafkaHandler) BudgetCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BudgetCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.budget.CreateBudget(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create budget via Kafka: %v", err)
			return
		}
		log.Printf("Created Budget: %+v", res)
	}
}
func (h *KafkaHandler) BudgetUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BudgetUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.budget.UpdateBudget(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update budget via Kafka: %v", err)
			return
		}
		log.Printf("Updated Budget: %+v", res)
	}
}
func (h *KafkaHandler) CategoryCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.CategoryCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.category.CreateCategory(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create category via Kafka: %v", err)
			return
		}
		log.Printf("Created Category: %+v", res)
	}
}
func (h *KafkaHandler) CategoryUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.CategoryUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.category.UpdateCategory(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update category via Kafka: %v", err)
			return
		}
		log.Printf("Updated Category: %+v", res)
	}
}
func (h *KafkaHandler) GoalCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.GoalCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.goal.CreateGoal(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create goal via Kafka: %v", err)
			return
		}
		log.Printf("Created Goal: %+v", res)
	}
}
func (h *KafkaHandler) GoalUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.GoalUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.goal.UpdateGoal(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update goal via Kafka: %v", err)
			return
		}
		log.Printf("Updated Goal: %+v", res)
	}
}
func (h *KafkaHandler) TransactionCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.TransactionCreate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.transaction.CreateTransaction(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create transaction via Kafka: %v", err)
			return
		}
		log.Printf("Created Transaction: %+v", res)
	}
}
func (h *KafkaHandler) TransactionUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.TransactionUpdate
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.transaction.UpdateTransaction(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update transaction via Kafka: %v", err)
			return
		}
		log.Printf("Updated Transaction: %+v", res)
	}
}
