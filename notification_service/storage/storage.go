package storage

import pb "gitlab.com/saladin2098/finance_tracker1/notification_service/genproto"

type StorageI interface {
	Notification() NotificationI
}
type NotificationI interface {
	CreateNotification(req *pb.NotificationCreate) (*pb.Void, error)
	NotifyAll(req *pb.NotificationMessage) (*pb.Void, error)
	DeleteNotificcation(id *pb.ById) (*pb.Void, error)
	UpdateNotificcation(req *pb.NotificationUpdate) (*pb.Void, error)
	GetNotifications(req *pb.NotifFilter) (*pb.NotificationList, error)
	GetNotification(id *pb.ById) (*pb.NotificationGet, error)
}
