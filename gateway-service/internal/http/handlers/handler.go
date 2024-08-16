package handlers

import (
	"finance_tracker/gateway/internal/grpc"
	"finance_tracker/gateway/internal/pkg/kafka"
	"finance_tracker/gateway/internal/pkg/logger"

	"github.com/go-redis/redis"
)

type Handler struct {
	Clients  grpc.Clients
	Producer kafka.KafkaProducer
	Logger   *logger.Logger
	Rdc      *redis.Client
}

func NewHandler(clients grpc.Clients, producer kafka.KafkaProducer, logger *logger.Logger, rdc *redis.Client) *Handler {
	return &Handler{Clients: clients, Producer: producer, Logger: logger, Rdc: rdc}
}
