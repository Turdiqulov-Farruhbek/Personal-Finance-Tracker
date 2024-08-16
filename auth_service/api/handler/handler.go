package handler

import (
	pb "finance_tracker/auth_service/genproto"
	"finance_tracker/auth_service/kafka"
)

type Handler struct {
	stg      pb.AuthServiceClient
	Producer kafka.KafkaProducer
}

func NewHandler(auth pb.AuthServiceClient, kafka kafka.KafkaProducer) *Handler {
	return &Handler{stg: auth, Producer: kafka}
}
