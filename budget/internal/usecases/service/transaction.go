package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"finance_tracker/budget/internal/storage"
	"finance_tracker/budget/internal/usecases/kafka"
	"google.golang.org/protobuf/encoding/protojson"
)

type TransactionService struct {
	stg storage.StorageI
	pb.UnimplementedTransactionServiceServer
	producer kafka.KafkaProducer
}

func NewTransactionService(stg storage.StorageI, kafka kafka.KafkaProducer) *TransactionService {
	return &TransactionService{
		stg:      stg,
		producer: kafka,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *pb.TransactionCreate) (*pb.Void, error) {
	//check the account balance
	account_id := pb.ById{Id: req.AccountId}
	account, err := s.stg.Account().GetAccount(&account_id)
	if err != nil {
		return nil, err
	}
	currentBalance := account.Balance
	if req.Type == "credit" {
		currentBalance -= req.Amount
	} else {
		currentBalance += req.Amount
	}

	if currentBalance < 0 {
		return nil, errors.New("insufficient funds in the account with id: " + req.AccountId)
	}
	//update the account balance
	account_update := pb.AccountUpdate{
		Id:   req.AccountId,
		Body: &pb.AccountUpt{Balance: currentBalance},
	}
	_, err = s.stg.Account().UpdateAccount(&account_update)
	if err != nil {
		return nil, err
	}
	// Check for budget targets if the transaction is a credit
	if req.Type == "credit" {
		budgetFilter := pb.BudgetFilter{
			UserId: req.UserId,
			Status: "ongoing",
			Filter: &pb.Filter{},
		}
		budgets, err := s.stg.Budget().ListBudgets(&budgetFilter)
		if err != nil {
			return nil, err
		}
		for _, budget := range budgets.Budgets {
			var totalSpendings float32
			var timeFrom, timeTo string

			// Determine the time range based on the budget period
			switch budget.Period {
			case "daily":
				timeFrom = time.Now().Format("2006-01-02")
				timeTo = time.Now().AddDate(0, 0, 1).Format("2006-01-02")

			case "weekly":
				timeFrom = time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Format("2006-01-02")
				timeTo = time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())).Format("2006-01-02")

			case "monthly":
				firstDayOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
				timeFrom = firstDayOfMonth.Format("2006-01-02")
				timeTo = firstDayOfMonth.AddDate(0, 1, 0).Format("2006-01-02")

			case "yearly":
				firstDayOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Now().Location())
				timeFrom = firstDayOfYear.Format("2006-01-02")
				timeTo = firstDayOfYear.AddDate(1, 0, 0).Format("2006-01-02")

			default:
				return nil, fmt.Errorf("invalid budget period: %s", budget.Period)
			}

			cursor, err := s.stg.Transaction().ListTransactions(&pb.TransactionFilter{
				AccountId: req.AccountId,
				Type:      "credit",
				TimeFrom:  timeFrom,
				TimeTo:    timeTo,
				Filter:    &pb.Filter{},
			})
			if err != nil {
				return nil, err
			}

			// Calculate the total spendings within the selected time range
			for _, transaction := range cursor.TransactionGet {
				totalSpendings += transaction.Amount
			}

			// Check if the total spendings exceed the budget amount
			if totalSpendings >= budget.Amount {
				// Create a notification message
				message := fmt.Sprintf("You have exceeded your %s budget by %.2f as of %v",
					budget.Period, totalSpendings-budget.Amount, time.Now())
				notification := pb.NotificationCreate{
					RecieverId: req.UserId,
					Message:    message,
				}

				// Marshal the notification into JSON
				input, err := protojson.Marshal(&notification)
				if err != nil {
					return nil, err
				}

				// Produce the notification message to the Kafka topic
				if err := s.producer.ProduceMessages("notification-create", input); err != nil {
					return nil, err
				}
			}

		}
	}

	//update goal
	filter := pb.GoalFilter{
		UserId: req.UserId,
		Status: "active",
	}
	filter.Filter = &pb.Filter{}
	goals, err := s.stg.Goal().ListGoals(&filter)
	if err != nil {
		return nil, err
	}
	goal := goals.Goals[0]
	if req.Type == "debit" {
		goalProgress := goal.CurrentAmount + req.Amount
		if goalProgress >= goal.TargetAmount {
			message := fmt.Sprintf("Congratulations! You have reached your goal of $%v", goal.TargetAmount)
			notification := pb.NotificationCreate{
				RecieverId: req.UserId,
				Message:    message,
			}
			input, err := protojson.Marshal(&notification)
			if err != nil {
				return nil, err
			}
			err = s.producer.ProduceMessages("notification-create", input)
			if err != nil {
				return nil, err
			}
		}
		_, err = s.stg.Goal().UpdateGoal(&pb.GoalUpdate{
			Id:   goal.Id,
			Body: &pb.GoalUpt{CurrentAmount: goalProgress},
		})
		if err != nil {
			return nil, err
		}

	} else {
		goalProgress := goal.CurrentAmount - req.Amount
		_, err = s.stg.Goal().UpdateGoal(&pb.GoalUpdate{
			Id:   goal.Id,
			Body: &pb.GoalUpt{CurrentAmount: goalProgress},
		})
		if err != nil {
			return nil, err
		}
	}

	// send notification to kafka
	message := fmt.Sprintf("A new amount %.2f has been %ved from your account with id: %v", req.Amount, req.Type, req.AccountId)
	notification := pb.NotificationCreate{
		RecieverId: req.UserId,
		Message:    message,
	}

	input, err := protojson.Marshal(&notification)
	if err != nil {
		return nil, err
	}

	if err := s.producer.ProduceMessages("notification-create", input); err != nil {
		return nil, err
	}

	return s.stg.Transaction().CreateTransaction(req)
}

func (s *TransactionService) GetTransaction(ctx context.Context, req *pb.ById) (*pb.TransactionGet, error) {
	return s.stg.Transaction().GetTransaction(req)
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, req *pb.TransactionUpdate) (*pb.Void, error) {
	//get the current tarnaction to compare
	currentTransaction, err := s.stg.Transaction().GetTransaction(&pb.ById{Id: req.Id})
	if err != nil {
		return nil, err
	}

	difference := req.Body.Amount - currentTransaction.Amount

	//check the account balance
	account_id := pb.ById{Id: currentTransaction.AccountId}
	account, err := s.stg.Account().GetAccount(&account_id)
	if err != nil {
		return nil, err
	}
	currentBalance := account.Balance
	if currentTransaction.Type == "credit" {
		currentBalance -= difference
	} else {
		currentBalance += difference
	}

	if currentBalance < 0 {
		return nil, errors.New("insufficient funds in the account with id: " + currentTransaction.AccountId)
	}

	//update the account balance
	account_update := pb.AccountUpdate{
		Id:   currentTransaction.AccountId,
		Body: &pb.AccountUpt{Balance: currentBalance},
	}
	_, err = s.stg.Account().UpdateAccount(&account_update)
	if err != nil {
		return nil, err
	}

	// Check for budget targets if the transaction is a credit
	if currentTransaction.Type == "credit" {
		budgetFilter := pb.BudgetFilter{UserId: currentTransaction.UserId, Status: "ongoing"}
		budgets, err := s.stg.Budget().ListBudgets(&budgetFilter)
		if err != nil {
			return nil, err
		}
		for _, budget := range budgets.Budgets {
			var totalSpendings float32
			var timeFrom, timeTo string

			// Determine the time range based on the budget period
			switch budget.Period {
			case "daily":
				timeFrom = time.Now().Format("2006-01-02")
				timeTo = time.Now().AddDate(0, 0, 1).Format("2006-01-02")

			case "weekly":
				timeFrom = time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Format("2006-01-02")
				timeTo = time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())).Format("2006-01-02")

			case "monthly":
				firstDayOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
				timeFrom = firstDayOfMonth.Format("2006-01-02")
				timeTo = firstDayOfMonth.AddDate(0, 1, 0).Format("2006-01-02")

			case "yearly":
				firstDayOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Now().Location())
				timeFrom = firstDayOfYear.Format("2006-01-02")
				timeTo = firstDayOfYear.AddDate(1, 0, 0).Format("2006-01-02")

			default:
				return nil, fmt.Errorf("invalid budget period: %s", budget.Period)
			}

			cursor, err := s.stg.Transaction().ListTransactions(&pb.TransactionFilter{
				AccountId: currentTransaction.AccountId,
				Type:      "credit",
				TimeFrom:  timeFrom,
				TimeTo:    timeTo,
			})
			if err != nil {
				return nil, err
			}

			// Calculate the total spendings within the selected time range
			for _, transaction := range cursor.TransactionGet {
				totalSpendings += transaction.Amount
			}
			totalSpendings += difference

			// Check if the total spendings exceed the budget amount
			if totalSpendings >= budget.Amount {
				// Create a notification message
				message := fmt.Sprintf("You have exceeded your %s budget by %.2f as of %v",
					budget.Period, totalSpendings-budget.Amount, time.Now())
				notification := pb.NotificationCreate{
					RecieverId: currentTransaction.UserId,
					Message:    message,
				}

				// Marshal the notification into JSON
				input, err := protojson.Marshal(&notification)
				if err != nil {
					return nil, err
				}

				// Produce the notification message to the Kafka topic
				if err := s.producer.ProduceMessages("notification-create", input); err != nil {
					return nil, err
				}
			}

		}
	}

	//update goal
	filter := pb.GoalFilter{
		UserId: currentTransaction.UserId,
		Status: "ongoing",
	}
	goals, err := s.stg.Goal().ListGoals(&filter)
	if err != nil {
		return nil, err
	}
	goal := goals.Goals[0]
	if currentTransaction.Type == "debit" {
		goalProgress := goal.CurrentAmount + difference
		if goalProgress >= goal.TargetAmount {
			message := fmt.Sprintf("Congratulations! You have reached your goal of $%v", goal.TargetAmount)
			notification := pb.NotificationCreate{
				RecieverId: currentTransaction.UserId,
				Message:    message,
			}
			input, err := protojson.Marshal(&notification)
			if err != nil {
				return nil, err
			}
			err = s.producer.ProduceMessages("notification-create", input)
			if err != nil {
				return nil, err
			}
		}
		_, err = s.stg.Goal().UpdateGoal(&pb.GoalUpdate{
			Id:   goal.Id,
			Body: &pb.GoalUpt{CurrentAmount: goalProgress},
		})
		if err != nil {
			return nil, err
		}

	} else {
		goalProgress := goal.CurrentAmount - difference
		_, err = s.stg.Goal().UpdateGoal(&pb.GoalUpdate{
			Id:   goal.Id,
			Body: &pb.GoalUpt{CurrentAmount: goalProgress},
		})
		if err != nil {
			return nil, err
		}
	}

	// send notification to kafka
	message := fmt.Sprintf("A new amount %.2f has been %ved from your account with id: %v", difference, req.Body.Type, currentTransaction.AccountId)
	notification := pb.NotificationCreate{
		RecieverId: currentTransaction.UserId,
		Message:    message,
	}

	input, err := protojson.Marshal(&notification)
	if err != nil {
		return nil, err
	}

	if err := s.producer.ProduceMessages("notification-create", input); err != nil {
		return nil, err
	}

	return s.stg.Transaction().UpdateTransaction(req)
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	//get account id from transaction
	transaction, err := s.stg.Transaction().GetTransaction(req)
	if err != nil {
		return nil, err
	}
	account_id := pb.ById{Id: transaction.AccountId}
	account, err := s.stg.Account().GetAccount(&account_id)
	if err != nil {
		return nil, err
	}
	currentBalance := account.Balance
	if transaction.Type == "credit" {
		currentBalance += transaction.Amount
	} else {
		currentBalance -= transaction.Amount
	}

	//update the account balance
	account_update := pb.AccountUpdate{
		Id:   transaction.AccountId,
		Body: &pb.AccountUpt{Balance: currentBalance},
	}
	_, err = s.stg.Account().UpdateAccount(&account_update)
	if err != nil {
		return nil, err
	}
	//update goal
	filter := pb.GoalFilter{
		UserId: transaction.UserId,
		Status: "ongoing",
	}
	goals, err := s.stg.Goal().ListGoals(&filter)
	if err != nil {
		return nil, err
	}
	goal := goals.Goals[0]
	if transaction.Type == "debit" {
		goalProgress := goal.CurrentAmount - transaction.Amount

		_, err = s.stg.Goal().UpdateGoal(&pb.GoalUpdate{
			Id:   goal.Id,
			Body: &pb.GoalUpt{CurrentAmount: goalProgress},
		})
		if err != nil {
			return nil, err
		}

	} else {
		goalProgress := goal.CurrentAmount + transaction.Amount
		_, err = s.stg.Goal().UpdateGoal(&pb.GoalUpdate{
			Id:   goal.Id,
			Body: &pb.GoalUpt{CurrentAmount: goalProgress},
		})
		if err != nil {
			return nil, err
		}
	}

	// send notification to kafka
	message := fmt.Sprintf("transaction with %v id has been deleted successfully", req.Id)
	notification := pb.NotificationCreate{
		RecieverId: transaction.UserId,
		Message:    message,
	}

	input, err := protojson.Marshal(&notification)
	if err != nil {
		return nil, err
	}

	if err := s.producer.ProduceMessages("notification-create", input); err != nil {
		return nil, err
	}

	return s.stg.Transaction().DeleteTransaction(req)
}

func (s *TransactionService) ListTransactions(ctx context.Context, req *pb.TransactionFilter) (*pb.TransactionList, error) {
	return s.stg.Transaction().ListTransactions(req)
}
