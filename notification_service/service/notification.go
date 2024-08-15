package service

import (
	"context"

	pb "gitlab.com/saladin2098/finance_tracker1/notification_service/genproto"
	"gitlab.com/saladin2098/finance_tracker1/notification_service/storage"
)

type NotificationService struct {
	stg storage.StorageI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(stg storage.StorageI) *NotificationService {
	return &NotificationService{stg: stg}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.NotificationCreate) (*pb.Void, error) {
	return s.stg.Notification().CreateNotification(req)
}
func (s *NotificationService) NotifyAll(ctx context.Context, req *pb.NotificationMessage) (*pb.Void, error) {
	return s.stg.Notification().NotifyAll(req)
}
func (s *NotificationService) DeleteNotificcation(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return s.stg.Notification().DeleteNotificcation(req)
}
func (s *NotificationService) UpdateNotificcation(ctx context.Context, req *pb.NotificationUpdate) (*pb.Void, error) {
	return s.stg.Notification().UpdateNotificcation(req)
}
func (s *NotificationService) GetNotifications(ctx context.Context, req *pb.NotifFilter) (*pb.NotificationList, error) {
	return s.stg.Notification().GetNotifications(req)
}
func (s *NotificationService) GetNotification(ctx context.Context, req *pb.ById) (*pb.NotificationGet, error) {
	return s.stg.Notification().GetNotification(req)
}
