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

// CreateGoal creates a new goal
// @Summary 		Create Goal
// @Description 	Create a new goal
// @Tags 			Goal
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		goal body    pb.GoalCreate  true   "Goal data"
// @Success 		200  {string}  string "Goal created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/goals/create [post]
func (h *Handler) CreateGoal(c *gin.Context) {
	var req pb.GoalCreate
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
	if err := h.Producer.ProduceMessages("goal-create", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Goal created successfully")
}

// UpdateGoal updates an existing goal
// @Summary 		Update Goal
// @Description 	Update an existing goal
// @Tags 			Goal
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		goal_id path    string       true   "Goal ID"
// @Param     		goal    body    pb.GoalUpt  true   "Goal data"
// @Success 		200  {string}  string "Goal updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/goals/{goal_id} [put]
func (h *Handler) UpdateGoal(c *gin.Context) {
	goalID := c.Param("goal_id")

	var body pb.GoalUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	// Set the goal ID
	req := pb.GoalUpdate{
		Id:   goalID,
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
	if err := h.Producer.ProduceMessages("goal-update", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Goal updated successfully")
}

// GetGoal retrieves a goal by ID
// @Summary 		Get Goal
// @Description 	Retrieve a goal by its ID
// @Tags 			Goal
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		goal_id path    string       true   "Goal ID"
// @Success 		200  {object}  pb.GoalGet "Goal details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Goal not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/goals/{goal_id} [get]
func (h *Handler) GetGoal(c *gin.Context) {
	goalID := c.Param("goal_id")

	// Check cache first
	cached, err := h.Rdc.Get(goalID).Result()
	if err == nil {
		var goal pb.GoalGet
		if err := protojson.Unmarshal([]byte(cached), &goal); err == nil {
			c.JSON(200, &goal)
			return
		}

	}

	// If not in cache, retrieve from gRPC
	req := &pb.ById{Id: goalID}
	res, err := h.Clients.Goal.GetGoal(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get goal:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(goalID, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// ListGoals retrieves a list of goals
// @Summary List Goals
// @Description Retrieve a list of goals
// @Tags Goal
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "User ID"
// @Param status query string false "Goal Status"
// @Param name query string false "Goal Name"
// @Param target_from query float32 false "Target Amount From"
// @Param target_to query float32 false "Target Amount To"
// @Param deadline_from query string false "Deadline From"
// @Param deadline_to query string false "Deadline To"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.GoalList "List of goals"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /goals/list [get]
func (h *Handler) ListGoals(c *gin.Context) {
	// Initialize filter and set default values
	var filter pb.GoalFilter
	filter.Filter = &pb.Filter{}

	// Read each query parameter individually
	if userId := c.Query("user_id"); userId != "" {
		filter.UserId = userId
	}
	if status := c.Query("status"); status != "" {
		filter.Status = status
	}
	if name := c.Query("name"); name != "" {
		filter.Name = name
	}
	if targetFrom := c.Query("target_from"); targetFrom != "" {
		if t, err := strconv.ParseFloat(targetFrom, 32); err == nil {
			filter.TargetFrom = float32(t)
		}
	}
	if targetTo := c.Query("target_to"); targetTo != "" {
		if t, err := strconv.ParseFloat(targetTo, 32); err == nil {
			filter.TargetTo = float32(t)
		}
	}
	if deadlineFrom := c.Query("deadline_from"); deadlineFrom != "" {
		filter.DeadlineFrom = deadlineFrom
	}
	if deadlineTo := c.Query("deadline_to"); deadlineTo != "" {
		filter.DeadlineTo = deadlineTo
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
	cacheKey := fmt.Sprintf("goal-list:%v", filter.String())
	cached, err := h.Rdc.Get(cacheKey).Result()
	if err == nil {
		var goalList pb.GoalList
		if err := protojson.Unmarshal([]byte(cached), &goalList); err == nil {
			c.JSON(200, &goalList)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	res, err := h.Clients.Goal.ListGoals(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list goals:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(cacheKey, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// DeleteGoal deletes a goal by ID
// @Summary 		Delete Goal
// @Description 	Delete a goal by its ID
// @Tags 			Goal
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		goal_id path    string       true   "Goal ID"
// @Success 		200  {string}  string "Goal deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Goal not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/goals/{goal_id} [delete]
func (h *Handler) DeleteGoal(c *gin.Context) {
	goalID := c.Param("goal_id")

	// Call gRPC delete method
	req := &pb.ById{Id: goalID}
	_, err := h.Clients.Goal.DeleteGoal(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete goal:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Clear cache
	if err := h.Rdc.Del(goalID).Err(); err != nil {
		h.Logger.WARN.Println("Failed to delete cache entry:", err)
	}

	c.JSON(200, "Goal deleted successfully")
}
