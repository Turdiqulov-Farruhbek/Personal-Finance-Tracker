package server

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "gitlab.com/saladin2098/finance_tracker1/auth_service/genproto"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/kafka"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/service"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/storage/postgres"
	"google.golang.org/grpc"
)

func Server() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error with dbconnection", err)
		return
	}

	userService := service.NewUserService(db)

	//*kafka\\//\\
	time.Sleep(time.Second * 20)
	brokers := []string{"localhost:9092"} //////////////////////////////////////////////////////

	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "user-create", "user", kafka.UserCreateHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'user-create' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	if err := kcm.RegisterConsumer(brokers, "forgot_password", "forgot_password_id", kafka.ForgotPasswordHandler(userService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'forgot_password' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}
	// *kafka\\//\\

	newServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(newServer, service.NewUserService(db))

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on :50050")
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
