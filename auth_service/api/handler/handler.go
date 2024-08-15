package handler

import (
	pb "gitlab.com/saladin2098/finance_tracker1/auth_service/genproto"
	"gitlab.com/saladin2098/finance_tracker1/auth_service/kafka"
)

type Handler struct {
	stg      pb.AuthServiceClient
	Producer kafka.KafkaProducer
}

func NewHandler(auth pb.AuthServiceClient, kafka kafka.KafkaProducer) *Handler {
	return &Handler{stg: auth, Producer: kafka}
}
