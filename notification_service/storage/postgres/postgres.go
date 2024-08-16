package postgres

import (
	"database/sql"
	"fmt"

	"finance_tracker/notification_service/config"
	pb "finance_tracker/notification_service/genproto"
	"finance_tracker/notification_service/storage"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Storage struct {
	db            *sql.DB
	NotificationS storage.NotificationI
}

func ConnectDB(cfg *config.Config) (*Storage, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//connecting to user service
	auth_path := cfg.AuthHost + cfg.AuthPort
	conn_auth, err := grpc.NewClient(auth_path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	authC := pb.NewAuthServiceClient(conn_auth)

	notifS := NewNotificationRepo(db, authC)
	return &Storage{
		db:            db,
		NotificationS: notifS,
	}, nil
}
func (s *Storage) Notification() storage.NotificationI {
	return s.NotificationS
}
