package service

import (
	"context"
	"fmt"
	"time"

	pb "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"
	"gitlab.com/saladin2098/finance_tracker1/budget/internal/storage"
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
func (s *ReportService) BudgetPerformance(ctx context.Context, req *pb.BudgetPerReq) (*pb.BudgetPerGet, error) {
	// Retrieve all budgets for the user
	budgets, err := s.stg.Budget().ListBudgets(&pb.BudgetFilter{
		UserId: req.UserId,
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
			// Calculate daily spendings within the budget's timeline
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24)
		case "weekly":
			// Calculate weekly spendings within the budget's timeline
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24*7)
		case "monthly":
			// Calculate monthly spendings within the budget's timeline
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24*30)
		case "yearly":
			// Calculate yearly spendings within the budget's timeline
			totalPeriods, periodsOverBudget, avgSpendings = s.calculatePeriodPerformance(budget, time.Hour*24*365)
		}

		if periodsOverBudget > 0 {
			progress = fmt.Sprintf("Over budget %d times out of %d periods", periodsOverBudget, totalPeriods)
		} else {
			progress = "Within budget for all periods"
		}

		performances = append(performances, &pb.PeriodBudgetPer{
			StartDate:    budget.StartDate,
			EndDate:      budget.EndDate,
			AvgSpendings: float32(avgSpendings),
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

func (s *ReportService) GoalProgress(ctx context.Context, req *pb.GoalProgresReq) (*pb.GoalProgresGet, error) {
	// Retrieve all goals for the user
	goals, err := s.stg.Goal().ListGoals(&pb.GoalFilter{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	var goalProgresses []*pb.GoalProgress

	for _, goal := range goals.Goals {
		// Calculate total income within the goal's timeline
		startDate, _ := time.Parse(time.RFC3339, goal.CreatedAt)
		deadline, _ := time.Parse(time.RFC3339, goal.Deadline)

		trs, err := s.stg.Transaction().ListTransactions(&pb.TransactionFilter{
			UserId:   goal.UserId,
			Type:     "credit", // Credit indicates income
			TimeFrom: startDate.Format(time.RFC3339),
			TimeTo:   deadline.Format(time.RFC3339),
		})
		if err != nil {
			return nil, err
		}

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

func (s *ReportService) calculatePeriodPerformance(budget *pb.BudgetGet, periodDuration time.Duration) (totalPeriods int, periodsOverBudget int, avgSpendings float32) {
	// Calculate the total number of periods in the budget's timeline
	startDate, _ := time.Parse(time.RFC3339, budget.StartDate)
	endDate, _ := time.Parse(time.RFC3339, budget.EndDate)
	duration := endDate.Sub(startDate)
	totalPeriods = int(duration / periodDuration)

	var totalSpendings float32
	for i := 0; i < totalPeriods; i++ {
		periodStart := startDate.Add(time.Duration(i) * periodDuration)
		periodEnd := periodStart.Add(periodDuration)

		// Calculate the total spendings for this period
		trs, _ := s.stg.Transaction().ListTransactions(&pb.TransactionFilter{
			UserId:   budget.UserId,
            Type:     "debit", // Debit indicates spending
            TimeFrom: periodStart.Format(time.RFC3339),
            TimeTo:   periodEnd.Format(time.RFC3339),
		})

		var spendings float32
		for _, tr := range trs.TransactionGet {
			spendings += tr.Amount
		}

		totalSpendings += spendings

		if spendings > float32(budget.Amount) {
			periodsOverBudget++
		}
	}

	avgSpendings = totalSpendings / float32(totalPeriods)
	return
}

