package service

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"finance_tracker/budget/internal/storage"
)

type ReportService struct {
	stg storage.StorageI
	pb.UnimplementedReportServiceServer
}

func NewReportService(stg storage.StorageI) *ReportService {
	return &ReportService{stg: stg}
}

func (s *ReportService) GetSpendings(ctx context.Context, req *pb.SpendingReq) (*pb.SpendingGet, error) {
	transaction_filter := pb.TransactionFilter{
		UserId:     req.UserId,
		Type:       "credit",
		TimeFrom:   req.DateFrom,
		TimeTo:     req.DateTo,
		CategoryId: req.CategoryId,
		Filter:     &pb.Filter{},
	}
	transactions, err := s.stg.Transaction().ListTransactions(&transaction_filter)
	if err != nil {
		return nil, err
	}
	var total_spendings float32
	for _, transaction := range transactions.TransactionGet {
		total_spendings += transaction.Amount
	}
	return &pb.SpendingGet{
		UserId:       req.UserId,
		DateFrom:     req.DateFrom,
		DateTo:       req.DateTo,
		CategoryId:   req.CategoryId,
		TotalAmount:  total_spendings,
		Transactions: transactions.TransactionGet,
	}, nil
}
func (s *ReportService) GetIncomes(ctx context.Context, req *pb.IncomeReq) (*pb.IncomeGet, error) {
	transaction_filter := pb.TransactionFilter{
		UserId:     req.UserId,
		Type:       "debit",
		TimeFrom:   req.DateFrom,
		TimeTo:     req.DateTo,
		CategoryId: req.CategoryId,
		Filter:     &pb.Filter{},
	}
	transactions, err := s.stg.Transaction().ListTransactions(&transaction_filter)
	if err != nil {
		return nil, err
	}
	var total_incomes float32
	for _, transaction := range transactions.TransactionGet {
		total_incomes += transaction.Amount
	}
	return &pb.IncomeGet{
		UserId:      req.UserId,
		DateFrom:    req.DateFrom,
		DateTo:      req.DateTo,
		CategoryId:  req.CategoryId,
		TotalAmount: total_incomes,
	}, nil
}

func (s *ReportService) GoalProgress(ctx context.Context, req *pb.GoalProgresReq) (*pb.GoalProgresGet, error) {
	// Retrieve all goals for the user
	goals, err := s.stg.Goal().ListGoals(&pb.GoalFilter{
		UserId: req.UserId,
		Filter: &pb.Filter{},
	})
	if err != nil {
		return nil, err
	}
	log.Println("GoalS", goals)

	var goalProgresses []*pb.GoalProgress

	for _, goal := range goals.Goals {
		// Calculate total income within the goal's timeline
		log.Println("timelines in goal ", goal.CreatedAt, goal.Deadline)
		startDate := goal.CreatedAt
		deadlineStr, err := time.Parse("02-01-2006", goal.Deadline)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}
		deadline := deadlineStr.Format("2006-01-02 15:04:05 -0700 MST")

		trs, err := s.stg.Transaction().ListTransactions(&pb.TransactionFilter{
			UserId:   goal.UserId,
			Type:     "debit", // Credit indicates income
			TimeFrom: startDate,
			TimeTo:   deadline,
			Filter:   &pb.Filter{},
		})
		if err != nil {
			return nil, err
		}
		log.Println("transactions", trs)

		var totalIncome float32
		for _, tr := range trs.TransactionGet {
			totalIncome += tr.Amount
		}

		// Determine if the goal is reached
		var progress string
		if totalIncome >= goal.TargetAmount {
			progress = "Goal reached"
		} else {
			progress = fmt.Sprintf("Current progress: %.2f%%", (totalIncome/goal.TargetAmount)*100)
		}
		log.Println(totalIncome, "total income form report ")

		goalProgresses = append(goalProgresses, &pb.GoalProgress{
			Deadline:      goal.Deadline,
			TargetAmount:  goal.TargetAmount,
			CurrentAmount: float32(totalIncome),
			Progress:      progress,
			GoalName:      goal.Name,
		})
	}

	return &pb.GoalProgresGet{
		UserId: req.UserId,
		Goals:  goalProgresses,
	}, nil
}

func (s *ReportService) BudgetPerformance(ctx context.Context, req *pb.BudgetPerReq) (*pb.BudgetPerGet, error) {
	// Retrieve all budgets for the user
	budgets, err := s.stg.Budget().ListBudgets(&pb.BudgetFilter{
		UserId: req.UserId,
		Filter: &pb.Filter{},
	})
	if err != nil {
		return nil, err
	}

	var performances []*pb.PeriodBudgetPer

	for _, budget := range budgets.Budgets {
		var avgSpendings float32
		var progress string
		var totalPeriods int
		var periodsOverBudget int

		// Determine the budget's period type and calculate spendings accordingly
		switch budget.Period {
		case "daily":
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24)
		case "weekly":
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24*7)
		case "monthly":
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24*30)
		case "yearly":
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24*365)
		default:
			// Handle unexpected period types
			continue
		}

		if periodsOverBudget > 0 {
			progress = fmt.Sprintf("Over budget %d times out of %d periods", periodsOverBudget, totalPeriods)
		} else {
			progress = "Within budget for all periods"
		}

		performances = append(performances, &pb.PeriodBudgetPer{
			StartDate:    budget.StartDate,
			EndDate:      budget.EndDate,
			AvgSpendings: avgSpendings,
			TargetAmount: budget.Amount,
			Progress:     progress,
			Period:       budget.Period,
		})
	}

	return &pb.BudgetPerGet{
		UserId:       req.UserId,
		Performances: performances,
	}, nil
}

func (s *ReportService) calculatePeriodPerformance(budget *pb.BudgetGet, periodDuration time.Duration) (totalPeriods int, periodsOverBudget int, avgSpendings float32) {
	// Calculate the total number of periods in the budget's timeline
	startDate, _ := time.Parse("02-01-2006", budget.StartDate)
	endDate, _ := time.Parse("02-01-2006", budget.EndDate)
	duration := endDate.Sub(startDate)
	totalPeriods = int(duration / periodDuration)
	log.Println("total periods ", totalPeriods)

	var totalSpendings float32
	for i := 0; i <= totalPeriods; i++ {
		periodStart := startDate.Add(time.Duration(i) * periodDuration)
		periodEnd := periodStart.Add(periodDuration)

		// Retrieve transactions for the period
		trs, err := s.stg.Transaction().ListTransactions(&pb.TransactionFilter{
			UserId:   budget.UserId,
			Type:     "credit", // Debit indicates spending
			TimeFrom: periodStart.String(),
			TimeTo:   periodEnd.String(),
			Filter:   &pb.Filter{},
		})
		if err != nil {
			return 0, 0, 0 // Handle error appropriately in production code
		}

		var spendings float32
		for _, tr := range trs.TransactionGet {
			spendings += tr.Amount
		}

		totalSpendings += spendings

		if spendings > float32(budget.Amount) {
			periodsOverBudget++
		}
	}

	if totalPeriods > 0 {
		avgSpendings = totalSpendings / float32(totalPeriods)
	} else {
		avgSpendings = 0
	}

	log.Println(avgSpendings)
	return totalPeriods, periodsOverBudget, avgSpendings
}
