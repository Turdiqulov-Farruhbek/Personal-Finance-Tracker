package main

import (
	"log"
	"time"

	"finance_tracker/auth_service/api"
	"finance_tracker/auth_service/api/handler"
	pb "finance_tracker/auth_service/genproto"
	"finance_tracker/auth_service/kafka"
	"finance_tracker/auth_service/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	go server.Server()

	auth, err := grpc.NewClient("auth_service:40040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	//kafka\\//
	time.Sleep(10 * time.Second)
	kafka, err := kafka.NewKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		log.Fatal(err)
		return
	}
	//kafka\\//

	client := pb.NewAuthServiceClient(auth)
	h := handler.NewHandler(client, kafka)
	r := api.NewGin(h)
	r.Run(":8080")
}
