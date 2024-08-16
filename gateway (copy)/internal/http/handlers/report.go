package handlers

import (
	pb "finance_tracker/gateway/internal/pkg/genproto"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// GetSpendings godoc
// @Summary 		Get Spendings
// @Description 	Retrieve spendings for a user within a date range and optional category filter
// @Tags 			Report
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param 			user_id  query    string  true   "User ID"
// @Param 			date_from query   string  false  "Start Date"
// @Param 			date_to   query   string  false  "End Date"
// @Param 			category_id query string  false  "Category ID"
// @Success 		200 {object} pb.SpendingGet
// @Failure 		400 {object} string  "Invalid request"
// @Failure 		500 {object} string  "Internal server error"
// @Router 			/report/spendings [get]
func (h *Handler) GetSpendings(c *gin.Context) {
	req := &pb.SpendingReq{
		UserId:     c.Query("user_id"),
		DateFrom:   c.Query("date_from"),
		DateTo:     c.Query("date_to"),
		CategoryId: c.Query("category_id"),
	}

	cacheKey := "spendings:" + req.UserId + ":" + req.DateFrom + ":" + req.DateTo + ":" + req.CategoryId
	if cached, err := h.Rdc.Get(cacheKey).Result(); err == nil && cached != "" {
		var resp pb.SpendingGet
		if err := protojson.Unmarshal([]byte(cached), &resp); err == nil {
			h.Logger.ERROR.Println(err)
			c.JSON(200, &resp)
			return
		}
	}

	resp, err := h.Clients.Report.GetSpendings(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get spendings:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	data, err := protojson.Marshal(resp)
	if err == nil {
		h.Rdc.Set(cacheKey, string(data), time.Minute*15)
	}
	c.JSON(200, resp)
}

// GetIncomes godoc
// @Summary 		Get Incomes
// @Description 	Retrieve incomes for a user within a date range and optional category filter
// @Tags 			Report
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			user_id  query    string  true   "User ID"
// @Param 			date_from query   string  false  "Start Date"
// @Param 			date_to   query   string  false  "End Date"
// @Param 			category_id query string  false  "Category ID"
// @Success 		200 {object} pb.IncomeGet "Incomes retrieved successfully"
// @Failure 		400 {object} string  "Invalid request"
// @Failure 		500 {object} string  "Internal server error"
// @Router 			/report/incomes [get]
func (h *Handler) GetIncomes(c *gin.Context) {
	req := &pb.IncomeReq{
		UserId:     c.Query("user_id"),
		DateFrom:   c.Query("date_from"),
		DateTo:     c.Query("date_to"),
		CategoryId: c.Query("category_id"),
	}

	cacheKey := "incomes:" + req.UserId + ":" + req.DateFrom + ":" + req.DateTo + ":" + req.CategoryId
	if cached, err := h.Rdc.Get(cacheKey).Result(); err == nil && cached != "" {
		var resp pb.IncomeGet
		if err := protojson.Unmarshal([]byte(cached), &resp); err == nil {
			c.JSON(200, &resp)
			return
		}
	}

	resp, err := h.Clients.Report.GetIncomes(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get incomes:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	data, err := protojson.Marshal(resp)
	if err == nil {
		h.Rdc.Set(cacheKey, string(data), time.Minute*15)
	}
	c.JSON(200, resp)
}

// BudgetPerformance godoc
// @Summary 		Get Budget Performance
// @Description 	Retrieve budget performance metrics for a user
// @Tags 			Report
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			user_id  query    string  true   "User ID"
// @Success 		200 {object} pb.BudgetPerGet "Budget performance retrieved successfully"
// @Failure 		400 {object} string  "Invalid request"
// @Failure 		500 {object} string  "Internal server error"
// @Router 			/report/budget-performance [get]
func (h *Handler) BudgetPerformance(c *gin.Context) {
	req := &pb.BudgetPerReq{
		UserId: c.Query("user_id"),
	}

	cacheKey := "budget_performance:" + req.UserId
	if cached, err := h.Rdc.Get(cacheKey).Result(); err == nil && cached != "" {
		var resp pb.BudgetPerGet
		if err := protojson.Unmarshal([]byte(cached), &resp); err == nil {
			c.JSON(200, &resp)
			return
		}
	}

	resp, err := h.Clients.Report.BudgetPerformance(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get budget performance:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	data, err := protojson.Marshal(resp)
	if err == nil {
		h.Rdc.Set(cacheKey, string(data), time.Minute*15)
	}
	c.JSON(200, resp)
}

// GoalProgress godoc
// @Summary 		Get Goal Progress
// @Description 	Retrieve goal progress metrics for a user
// @Tags 			Report
// @Security  		BearerAuth
// @Accept  		json
// @Produce  		json
// @Param 			user_id  query    string  true   "User ID"
// @Success 		200 {object} pb.GoalProgresGet "Goal progress retrieved successfully"
// @Failure 		400 {object} string  "Invalid request"
// @Failure 		500 {object} string  "Internal server error"
// @Router 			/report/goal-progress [get]
func (h *Handler) GoalProgress(c *gin.Context) {
	req := &pb.GoalProgresReq{
		UserId: c.Query("user_id"),
	}

	cacheKey := "goal_progress:" + req.UserId
	if cached, err := h.Rdc.Get(cacheKey).Result(); err == nil && cached != "" {
		var resp pb.GoalProgresGet
		if err := protojson.Unmarshal([]byte(cached), &resp); err == nil {
			c.JSON(200, &resp)
			return
		}
	}

	resp, err := h.Clients.Report.GoalProgress(c.Request.Context(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get goal progress:", err)
		c.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	data, err := protojson.Marshal(resp)
	if err == nil {
		h.Rdc.Set(cacheKey, string(data), time.Minute*15)
	}
	c.JSON(200, resp)
}
