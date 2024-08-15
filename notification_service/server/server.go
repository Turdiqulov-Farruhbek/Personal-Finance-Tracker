package server

import (
	"fmt"
	"log"
	"net"

	"gitlab.com/saladin2098/finance_tracker1/notification_service/config"
	pb "gitlab.com/saladin2098/finance_tracker1/notification_service/genproto"
	"gitlab.com/saladin2098/finance_tracker1/notification_service/kafka"
	"gitlab.com/saladin2098/finance_tracker1/notification_service/service"
	"gitlab.com/saladin2098/finance_tracker1/notification_service/storage/postgres"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	//connect to db
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//
	notificationService := service.NewNotificationService(db)

	//kafka\\*//\\\
	brokers := []string{cfg.KafkaHost + cfg.KafkaPort}

	//kafka producer
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
