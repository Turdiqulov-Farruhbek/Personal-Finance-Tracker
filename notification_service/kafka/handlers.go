package kafka

import (
	"context"
	// "encoding/json"
	"log"

	pb "finance_tracker/notification_service/genproto"
	"finance_tracker/notification_service/service"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/encoding/protojson"
)

func NotificationCreateHandler(notifService *service.NotificationService) func(producer KafkaProducer, message kafka.Message) {
	return func(producer KafkaProducer, message kafka.Message) {

		//unmarshal the message
		var notif pb.NotificationCreate
		if err := protojson.Unmarshal([]byte(message.Value), &notif); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := notifService.CreateNotification(context.Background(), &notif)
		if err != nil {
			log.Fatal("error while creating notification through kafka" + err.Error())
			return
		}
		log.Printf("Created notification: %+v", res)
	}
}
func NotifyAllHandler(notifService *service.NotificationService) func(producer KafkaProducer, message kafka.Message) {
	return func(producer KafkaProducer, message kafka.Message) {

		//unmarshal the message
		var req pb.NotificationMessage
		if err := protojson.Unmarshal([]byte(message.Value), &req); err != nil {
			log.Fatalf("Failed to unmarshal JSON to Protobuf message: %v", err)
			return
		}

		res, err := notifService.NotifyAll(context.Background(), &req)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Notified All: %+v", res)
	}
}
