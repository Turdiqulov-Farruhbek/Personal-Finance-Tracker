package app

import (
	"log"
	"net"

	"finance_tracker/budget/internal/pkg/config"
	pb "finance_tracker/budget/internal/pkg/genproto"
	mg "finance_tracker/budget/internal/pkg/mongo"
	"finance_tracker/budget/internal/storage/repo"
	"finance_tracker/budget/internal/usecases/kafka"
	"finance_tracker/budget/internal/usecases/service"
	"google.golang.org/grpc"
)

func Run(cfg config.Config) {

	// connect to mongo
	mg_db, err := mg.New(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	// connect to kafka
	kf_m, err := kafka.NewKafkaProducer([]string{cfg.KafkaUrl})
	if err != nil {
		log.Fatal(err)
	}

	// create storage
	stg := repo.NewStorage(mg_db)

	// create services

	account_service := service.NewAccountService(stg)
	budget_service := service.NewBudgetService(stg)
	transaction_service := service.NewTransactionService(stg, kf_m)
	category_service := service.NewCategoryService(stg)
	goal_service := service.NewGoalService(stg)
	report_service := service.NewReportService(stg)

	//register kafka handlers
	kafka_handler := &KafkaHandler{
		account:     account_service,
		budget:      budget_service,
		category:    category_service,
		goal:        goal_service,
		transaction: transaction_service,
	}

	// register kafka handlers
	if err := Register(kafka_handler, &cfg); err != nil {
		log.Fatal("Error registering kafka handlers: ", err)
	}

	server := grpc.NewServer()

	pb.RegisterAccountServiceServer(server, account_service)
	pb.RegisterBudgetServiceServer(server, budget_service)
	pb.RegisterTransactionServiceServer(server, transaction_service)
	pb.RegisterCategoryServiceServer(server, category_service)
	pb.RegisterGoalServiceServer(server, goal_service)
	pb.RegisterReportServiceServer(server, report_service)

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	log.Println("Starting gRPC server on port", cfg.GRPCPort)
	if err := server.Serve(listener); err != nil {
		log.Fatal("gRPC server failed to start: ", err)
	}

}
