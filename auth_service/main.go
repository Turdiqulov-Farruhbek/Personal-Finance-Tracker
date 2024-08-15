package main

import (
	"log"
	"time"

	"gitlab.com/saladin2098/finance_tracker1/auth_service/api"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/api/handler"
	pb "gitlab.com/saladin2098/finance_tracker1/auth_service/genproto"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/kafka"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/server"

	// "gitlab.com/saladin2098/finance_tracker1/auth_service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	go server.Server()

	auth, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	//kafka\\//
	time.Sleep(10 * time.Second)
	kafka, err := kafka.NewKafkaProducer([]string{"localhost:9092"}) ////////////////////////////////////////
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
