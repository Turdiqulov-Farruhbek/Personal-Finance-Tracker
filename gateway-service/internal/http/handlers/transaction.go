package handlers

import (
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"strconv"
	"time"

	pb "finance_tracker/gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateTransaction godoc
// @Summary 		Create a new transaction
// @Description 	Create a new transaction for a user account and write to Kafka
// @Tags 			Transaction
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			body body pb.TransactionCreate true "Transaction creation details"
// @Success 		200 {object} string "Transaction created successfully"
// @Failure 		400 {object} string "Invalid request"
// @Failure 		500 {object} string "Internal server error"
// @Router 			/transactions/create [post]
func (h *Handler) CreateTransaction(c *gin.Context) {
	var req pb.TransactionCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Produce message to Kafka
	messageData, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal Kafka message:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}
	h.Producer.ProduceMessages("transaction-create", messageData)

	c.JSON(200, gin.H{"message": "Transaction created successfully"})
}

// GetTransaction godoc
// @Summary 		Get a transaction by ID
// @Description 	Retrieve a transaction by its ID
// @Tags 			Transaction
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			id path string true "Transaction ID"
// @Success 		200 {object} pb.TransactionGet "Transaction retrieved successfully"
// @Failure 		400 {object} string "Invalid request"
// @Failure 		500 {object} string "Internal server error"
// @Router 			/transactions/{id} [get]
func (h *Handler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	cacheKey := fmt.Sprintf("transaction:%s", id)

	// Check cache
	if cached, err := h.Rdc.Get(cacheKey).Result(); err == nil {
		var res pb.TransactionGet
		if err := json.Unmarshal([]byte(cached), &res); err == nil {
			c.JSON(http.StatusOK, &res)
			return
		}
	}

	req := &pb.ById{Id: id}
	res, err := h.Clients.Transaction.GetTransaction(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Println("failed to get transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get transaction"})
		return
	}

	// Set cache with 15 minutes expiration
	data, _ := json.Marshal(res)
	h.Rdc.Set(cacheKey, data, time.Minute*15)

	c.JSON(http.StatusOK, res)
}

// UpdateTransaction godoc
// @Summary 		Update a transaction
// @Description 	Update a transaction by its ID and write to Kafka
// @Tags 			Transaction
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			body path pb.TransactionUpt true "Transaction update details"
// @Success 		200 {object} string "Transaction updated successfully"
// @Failure 		400 {object} string "Invalid request"
// @Failure 		500 {object} string "Internal server error"
// @Router 			/transactions/{transaction_id} [put]
func (h *Handler) UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var body pb.TransactionUpt
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := pb.TransactionUpdate{
		Id:   id,
		Body: &body,
	}
	_, err := h.Clients.Transaction.UpdateTransaction(c.Request.Context(), &req)
	if err != nil {
		h.Logger.ERROR.Println("failed to update transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update transaction"})
		return
	}

	// Invalidate cache for this transaction and user's transactions
	cacheKey := fmt.Sprintf("transaction:%s", id)
	if err := h.Rdc.Del(cacheKey).Err(); err != nil {
		h.Logger.ERROR.Println("failed to invalidate cache:", err)
	}
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal Kafka message:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("transaction-update", data)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce message to Kafka:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "transaction updated successfully"})
}

// DeleteTransaction godoc
// @Summary 		Delete a transaction
// @Description 	Delete a transaction by its ID
// @Tags 			Transaction
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			id path string true "Transaction ID"
// @Success 		200 {object} string "Transaction deleted successfully"
// @Failure 		400 {object} string "Invalid request"
// @Failure 		500 {object} string "Internal server error"
// @Router 			/transactions/{id} [delete]
func (h *Handler) DeleteTransaction(c *gin.Context) {
	req := &pb.ById{
		Id: c.Param("id"),
	}

	_, err := h.Clients.Transaction.DeleteTransaction(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete transaction:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}
	cacheKey := fmt.Sprintf("transaction:%s", req.Id)
	if err := h.Rdc.Del(cacheKey).Err(); err != nil {
		h.Logger.ERROR.Println("Failed to invalidate cache:", err)
	}

	c.JSON(200, gin.H{"message": "Transaction deleted successfully"})
}

// ListTransactions godoc
// @Summary 		List transactions with filters
// @Description 	Retrieve a list of transactions based on filters
// @Tags 			Transaction
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			user_id query    string  false  "User ID"
// @Param 			account_id query string  false  "Account ID"
// @Param 			category_id query string  false  "Category ID"
// @Param 			type query      string  false  "Transaction Type"
// @Param 			description query string false "Transaction Description"
// @Param 			time_from query string false "Time From"
// @Param 			time_to query   string false "Time To"
// @Param 			amount_from query float32  false "Amount From"
// @Param 			amount_to query float32   false "Amount To"
// @Param 			limit query     int32   false "Limit"
// @Param 			offset query    int32   false "Offset"
// @Success 		200 {object} pb.TransactionList "Transactions retrieved successfully"
// @Failure 		400 {object} string "Invalid request"
// @Failure 		500 {object} string "Internal server error"
// @Router 			/transactions/list [get]
func (h *Handler) ListTransactions(c *gin.Context) {
	var req pb.TransactionFilter

	// Convert query parameters to appropriate types
	req.UserId = c.Query("user_id")
	req.AccountId = c.Query("account_id")
	req.CategoryId = c.Query("category_id")
	req.Type = c.Query("type")
	req.Description = c.Query("description")
	req.TimeFrom = c.Query("time_from")
	req.TimeTo = c.Query("time_to")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	// Convert amount parameters from string to float
	if amountFrom := c.Query("amount_from"); amountFrom != "" {
		if value, err := strconv.ParseFloat(amountFrom, 32); err == nil {
			req.AmountFrom = float32(value)
		}
	}

	if amountTo := c.Query("amount_to"); amountTo != "" {
		if value, err := strconv.ParseFloat(amountTo, 32); err == nil {
			req.AmountTo = float32(value)
		}
	}
	req.Filter = &pb.Filter{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	// Call the service method
	res, err := h.Clients.Transaction.ListTransactions(c.Request.Context(), &req)
	if err != nil {
		h.Logger.ERROR.Println("failed to list transactions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list transactions" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
