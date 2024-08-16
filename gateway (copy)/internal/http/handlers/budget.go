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

// CreateBudget creates a new budget
// @Summary 		Create Budget
// @Description 	Create a new budget
// @Tags 			Budget
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		budget body    pb.BudgetCreate  true   "Budget data"
// @Success 		200  {string}  string "Budget created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/budgets/create [post]
func (h *Handler) CreateBudget(c *gin.Context) {
	var req pb.BudgetCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	// Marshal the request to JSON
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Produce message to Kafka
	if err := h.Producer.ProduceMessages("budget-create", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Budget created successfully")
}

// UpdateBudget updates an existing budget
// @Summary 		Update Budget
// @Description 	Update an existing budget
// @Tags 			Budget
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Budget ID"
// @Param     		budget    body    pb.BudgetCreate  true   "Budget data"
// @Success 		200  {string}  string "Budget updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/budgets/{id} [put]
func (h *Handler) UpdateBudget(c *gin.Context) {
	budgetID := c.Param("id")

	var body pb.BudgetCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	// Set the budget ID
	req := pb.BudgetUpdate{
		Id:   budgetID,
		Body: &body,
	}

	// Marshal the request to JSON
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Produce message to Kafka
	if err := h.Producer.ProduceMessages("budget-update", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Budget updated successfully")
}

// GetBudget retrieves a budget by ID
// @Summary 		Get Budget
// @Description 	Retrieve a budget by its ID
// @Tags 			Budget
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Budget ID"
// @Success 		200  {object}  pb.BudgetGet "Budget details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Budget not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/budgets/{id} [get]
func (h *Handler) GetBudget(c *gin.Context) {
	budgetID := c.Param("id")

	// Check cache first
	cached, err := h.Rdc.Get(budgetID).Result()
	if err == nil {
		var budget pb.BudgetGet
		if err := protojson.Unmarshal([]byte(cached), &budget); err == nil {
			c.JSON(200, &budget)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	req := &pb.ById{Id: budgetID}
	res, err := h.Clients.Budget.GetBudget(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get budget:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(budgetID, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// ListBudgets retrieves a list of budgets
// @Summary List Budgets
// @Description Retrieve a list of budgets
// @Tags Budget
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "User ID"
// @Param category_id query string false "Category ID"
// @Param amount_from query float32 false "Minimum Amount"
// @Param amount_to query float32 false "Maximum Amount"
// @Param period query string false "Budget Period"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.BudgetList "List of budgets"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /budgets/list [get]
func (h *Handler) ListBudgets(c *gin.Context) {
	// Initialize filter and set default values
	var filter pb.BudgetFilter
	filter.Filter = &pb.Filter{}

	// Read each query parameter individually
	if userId := c.Query("user_id"); userId != "" {
		filter.UserId = userId
	}
	if categoryId := c.Query("category_id"); categoryId != "" {
		filter.CategoryId = categoryId
	}
	if amountFrom := c.Query("amount_from"); amountFrom != "" {
		if amount, err := strconv.ParseFloat(amountFrom, 32); err == nil {
			filter.AmountFrom = float32(amount)
		}
	}
	if amountTo := c.Query("amount_to"); amountTo != "" {
		if amount, err := strconv.ParseFloat(amountTo, 32); err == nil {
			filter.AmountTo = float32(amount)
		}
	}
	if period := c.Query("period"); period != "" {
		filter.Period = period
	}
	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Filter.Limit = int32(l)
		}
	}
	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Filter.Offset = int32(o)
		}
	}

	// Generate a cache key based on the filter
	cacheKey := fmt.Sprintf("budget-list:%v", filter.String())
	cached, err := h.Rdc.Get(cacheKey).Result()
	if err == nil {
		var budgetList pb.BudgetList
		if err := protojson.Unmarshal([]byte(cached), &budgetList); err == nil {
			c.JSON(200, &budgetList)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	res, err := h.Clients.Budget.ListBudgets(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list budgets:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(cacheKey, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// DeleteBudget deletes a budget by ID
// @Summary 		Delete Budget
// @Description 	Delete a budget by its ID
// @Tags 			Budget
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Budget ID"
// @Success 		200  {string}  string "Budget deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Budget not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/budgets/delete/{id} [delete]
func (h *Handler) DeleteBudget(c *gin.Context) {
	budgetID := c.Param("budget_id")

	// Call gRPC delete method
	req := &pb.ById{Id: budgetID}
	_, err := h.Clients.Budget.DeleteBudget(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete budget:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Clear cache
	if err := h.Rdc.Del(budgetID).Err(); err != nil {
		h.Logger.WARN.Println("Failed to delete cache entry:", err)
	}

	c.JSON(200, "Budget deleted successfully")
}
