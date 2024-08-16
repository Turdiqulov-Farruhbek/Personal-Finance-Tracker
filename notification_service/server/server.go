package server

import (
	"fmt"
	"log"
	"net"

	"finance_tracker/notification_service/config"
	pb "finance_tracker/notification_service/genproto"
	"finance_tracker/notification_service/kafka"
	"finance_tracker/notification_service/service"
	"finance_tracker/notification_service/storage/postgres"

	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//
	notificationService := service.NewNotificationService(db)

	brokers := []string{cfg.KafkaHost + cfg.KafkaPort}

	kafkaProducer, err := kafka.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	kcm := kafka.NewKafkaConsumerManager(kafkaProducer)

	if err := kcm.RegisterConsumer(brokers, "notification-create", "notification", kafka.NotificationCreateHandler(notificationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notification-create' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "notify-all", "notification-all", kafka.NotifyAllHandler(notificationService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notify-all' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	server := grpc.NewServer()
	pb.RegisterNotificationServiceServer(server, notificationService)

	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on " + cfg.HTTPPort)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
