package app

import (
	"log"
	"path/filepath"
	"runtime"

	"finance_tracker/gateway/internal/grpc"
	"finance_tracker/gateway/internal/http"
	"finance_tracker/gateway/internal/http/handlers"
	"finance_tracker/gateway/internal/pkg/config"
	"finance_tracker/gateway/internal/pkg/kafka"
	"finance_tracker/gateway/internal/pkg/logger"

	"github.com/go-redis/redis"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Run(cfg config.Config) {
	logger := logger.NewLogger(basepath, cfg.LogPath)
	clients, err := grpc.NewClients(&cfg)
	if err != nil {
		logger.ERROR.Println("Failed to create gRPC clients", err)
		log.Fatal(err)
		return
	}

	//connect to kafka
	broker := []string{cfg.KafkaUrl}
	kafka, err := kafka.NewKafkaProducer(broker)
	if err != nil {
		logger.ERROR.Println("Failed to connect to Kafka", err)
		log.Fatal(err)
		return
	}
	defer kafka.Close()

	//redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisUrl,
	})

	// make handler
	h := handlers.NewHandler(*clients, kafka, logger, rdb)

	// make gin
	router := http.NewGin(h)

	// start server
	router.Run(":5050")
}
