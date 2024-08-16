package grpc

import (
	"finance_tracker/gateway/internal/pkg/config"
	pb "finance_tracker/gateway/internal/pkg/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Auth         pb.AuthServiceClient
	Budget       pb.BudgetServiceClient
	Report       pb.ReportServiceClient
	Account      pb.AccountServiceClient
	Goal         pb.GoalServiceClient
	Transaction  pb.TransactionServiceClient
	Notification pb.NotificationServiceClient
	Category     pb.CategoryServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	notif_conn, err := grpc.NewClient(cfg.NotificationUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	auth_conn, err := grpc.NewClient(cfg.AuthUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	budget_conn, err := grpc.NewClient(cfg.BudgetUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authClient := pb.NewAuthServiceClient(auth_conn)
	budgetClient := pb.NewBudgetServiceClient(budget_conn)
	reportClient := pb.NewReportServiceClient(budget_conn)
	accountClient := pb.NewAccountServiceClient(budget_conn)
	goalClient := pb.NewGoalServiceClient(budget_conn)
	transactionClient := pb.NewTransactionServiceClient(budget_conn)
	notificationClient := pb.NewNotificationServiceClient(notif_conn)
	category := pb.NewCategoryServiceClient(budget_conn)

	return &Clients{
		Auth:         authClient,
		Budget:       budgetClient,
		Report:       reportClient,
		Account:      accountClient,
		Goal:         goalClient,
		Transaction:  transactionClient,
		Notification: notificationClient,
		Category:     category,
	}, nil
}
