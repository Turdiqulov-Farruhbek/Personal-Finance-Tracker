package grpc

import (
	"finance_tracker/budget/internal/pkg/config"
	pb "finance_tracker/budget/internal/pkg/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectNotification(cfg *config.Config) (*pb.NotificationServiceClient, error) {
	notif_conn, err := grpc.NewClient(cfg.NotificationUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	output := pb.NewNotificationServiceClient(notif_conn)
	return &output, nil // Create a new client
}
