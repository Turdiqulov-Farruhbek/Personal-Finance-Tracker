package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	pb "finance_tracker/gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateNotification creates a new notification
// @Summary 		Create Notification
// @Description 	Create a new notification
// @Tags 			Notification
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		notification body    pb.NotificationCreate  true   "Notification data"
// @Success 		200  {string}  string "Notification created successfully"
// @Failure 		400  {object}  string  "Invalid request"
// @Failure 		500  {object}  string  "Internal server error"
// @Router 			/notifications/create [post]
func (h *Handler) CreateNotification(c *gin.Context) {
	var req pb.NotificationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Marshal the request to JSON
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	// Produce message to Kafka
	if err := h.Producer.ProduceMessages("notification-create", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(200, "Notification created successfully")
}

// UpdateNotification updates an existing notification
// @Summary 		Update Notification
// @Description 	Update an existing notification
// @Tags 			Notification
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		notification_id path    string                   true   "Notification ID"
// @Param     		notification    body    pb.NotificationUpt  true   "Notification data"
// @Success 		200  {string}  string "Notification updated successfully"
// @Failure 		400  {object}  string  "Invalid request"
// @Failure 		500  {object}  string  "Internal server error"
// @Router 			/notifications/{notification_id} [put]
func (h *Handler) UpdateNotification(c *gin.Context) {
	notificationID := c.Param("notification_id")

	var body pb.NotificationUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Set the notification ID
	req := pb.NotificationUpdate{
		NotificationId: notificationID,
		Body:           &body,
	}

	_, err := h.Clients.Notification.UpdateNotification(c, &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to update notification:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}
	// delete from cache
	if err := h.Rdc.Del(notificationID).Err(); err != nil {
		h.Logger.WARN.Println("Failed to delete cache entry:", err)
	}

	c.JSON(200, "Notification updated successfully")
}

// GetNotification retrieves a notification by ID
// @Summary 		Get Notification
// @Description 	Retrieve a notification by its ID
// @Tags 			Notification
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		notification_id path    string       true   "Notification ID"
// @Success 		200  {object}  pb.NotificationGet "Notification details"
// @Failure 		400  {object}  string  "Invalid request"
// @Failure 		404  {object}  string  "Notification not found"
// @Failure 		500  {object}  string  "Internal server error"
// @Router 			/notifications/{notification_id} [get]
func (h *Handler) GetNotification(c *gin.Context) {
	notificationID := c.Param("notification_id")

	// Check cache first
	cached, err := h.Rdc.Get(notificationID).Result()
	if err == nil {
		var notification pb.NotificationGet
		if err := protojson.Unmarshal([]byte(cached), &notification); err == nil {
			c.JSON(200, &notification)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	req := &pb.ById{Id: notificationID}
	res, err := h.Clients.Notification.GetNotification(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get notification:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(notificationID, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// ListNotifications retrieves a list of notifications
// @Summary List Notifications
// @Description Retrieve a list of notifications
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param reciever_id query string false "Receiver ID"
// @Param status query string false "Notification Status"
// @Param sender_id query string false "Sender ID"
// @Param limit query int32 false "Limit"
// @Param offset query int32 false "Offset"
// @Success 200 {object} pb.NotificationList "List of notifications"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /notifications/list [get]
func (h *Handler) ListNotifications(c *gin.Context) {
	// Initialize filter and set default values
	var filter pb.NotifFilter
	filter.Filter = &pb.Filter{}

	// Read each query parameter individually
	if receiverID := c.Query("reciever_id"); receiverID != "" {
		filter.RecieverId = receiverID
	}
	if status := c.Query("status"); status != "" {
		filter.Status = status
	}
	if senderID := c.Query("sender_id"); senderID != "" {
		filter.SenderId = senderID
	}
	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.ParseInt(limit, 10, 32); err == nil {
			filter.Filter.Limit = int32(l)
		}
	}
	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.ParseInt(offset, 10, 32); err == nil {
			filter.Filter.Offset = int32(o)
		}
	}

	// Generate a cache key based on the filter
	cacheKey := fmt.Sprintf("notification-list:%v", filter.String())
	cached, err := h.Rdc.Get(cacheKey).Result()
	if err == nil {
		var notifList pb.NotificationList
		if err := protojson.Unmarshal([]byte(cached), &notifList); err == nil {
			c.JSON(200, &notifList)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	res, err := h.Clients.Notification.GetNotifications(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list notifications:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(cacheKey, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// DeleteNotification deletes a notification by ID
// @Summary 		Delete Notification
// @Description 	Delete a notification by its ID
// @Tags 			Notification
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		notification_id path    string       true   "Notification ID"
// @Success 		200  {string}  string "Notification deleted successfully"
// @Failure 		400  {object}  string  "Invalid request"
// @Failure 		404  {object}  string  "Notification not found"
// @Failure 		500  {object}  string  "Internal server error"
// @Router 			/notifications/{notification_id} [delete]
func (h *Handler) DeleteNotification(c *gin.Context) {
	notificationID := c.Param("notification_id")

	// Call gRPC delete method
	req := &pb.ById{Id: notificationID}
	_, err := h.Clients.Notification.DeleteNotification(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete notification:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	// Clear cache
	if err := h.Rdc.Del(notificationID).Err(); err != nil {
		h.Logger.WARN.Println("Failed to delete cache entry:", err)
	}

	c.JSON(200, "Notification deleted successfully")
}

// NotifyAll sends a notification to all users in the target group
// @Summary 		Notify All
// @Description 	Send a notification to all users in the target group
// @Tags 			Notification
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		notification body    pb.NotificationMessage  true   "Notification data"
// @Success 		200  {string}  string "Notification sent successfully"
// @Failure 		400  {object}  string  "Invalid request"
// @Failure 		500  {object}  string  "Internal server error"
// @Router 			/notifications/notify-all [post]
func (h *Handler) NotifyAll(c *gin.Context) {
	var req pb.NotificationMessage
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Marshal the request to JSON
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	// Produce message to Kafka
	if err := h.Producer.ProduceMessages("notification-broadcast", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	c.JSON(200, "Notifications sent successfully")
}
