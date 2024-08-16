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

// @Summary Create Account
// @Description Create a new account
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param account body pb.AccountCreate true "Account data"
// @Success 200 {string} string "Account created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /account/create [post]
func (h *Handler) CreateAcccount(c *gin.Context) {
	var req pb.AccountCreate
	if err := c.ShouldBindJSON(&req); err != nil {

		h.Logger.ERROR.Println("Failed to bind request", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("account-create", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Account created successfully")
}

// @Summary Get Account
// @Description Get an account by ID
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 200 {object} pb.AccountGet "Account data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Account not found"
// @Failure 500 {string} string "Internal server error"
// @Router /account/{id} [get]
func (h *Handler) GetAccount(c *gin.Context) {
	id := c.Param("id")

	cacheKey := fmt.Sprintf("account_%s", id)
	var account pb.AccountGet

	val, err := h.Rdc.Get(cacheKey).Result()
	if err == nil && val != "" {
		if err := protojson.Unmarshal([]byte(val), &account); err == nil {
			c.JSON(200, &account)
			return
		}
	}

	req := &pb.ById{Id: id}
	resp, err := h.Clients.Account.GetAccount(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get account", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	data, err := protojson.Marshal(resp)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Rdc.Set(cacheKey, data, time.Minute*15).Err(); err != nil {
		h.Logger.ERROR.Println("Failed to cache account data", err)
	}

	c.JSON(200, resp)
}

// @Summary Update Account
// @Description Update an existing account by ID
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Param account body pb.AccountUpt true "Account update data"
// @Success 200 {string} string "Account updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /account/update/{id} [put]
func (h *Handler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")

	var body pb.AccountUpt
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	req := pb.AccountUpdate{
		Id:   id,
		Body: &body,
	}
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Producer.ProduceMessages("account-update", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce message to Kafka", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Account updated successfully")
}

// @Summary List Accounts
// @Description List accounts with filters
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "Account Name"
// @Param type query string false "Account Type"
// @Param currency query string false "Currency"
// @Param balanceMin query int false "Minimum Balance"
// @Param balanceMax query int false "Maximum Balance"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param userId query string false "User ID"
// @Success 200 {object} pb.AccounList "List of accounts"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /account/list [get]
func (h *Handler) ListAccounts(c *gin.Context) {
	var filter pb.AccountFilter

	// Manually extract and convert query parameters
	filter.Name = c.Query("name")
	filter.Type = c.Query("type")
	filter.Currency = c.Query("currency")
	filter.UserId = c.Query("userId")
	f := pb.Filter{}
	filter.Filter = &f

	// Convert balanceMin and balanceMax to float
	if balanceMin := c.Query("balanceMin"); balanceMin != "" {
		if value, err := strconv.ParseFloat(balanceMin, 32); err == nil {
			filter.BalanceMin = float32(value)
		} else {
			h.Logger.ERROR.Println("Invalid balanceMin", err)
			c.JSON(400, "Invalid balanceMin value")
			return
		}
	}

	if balanceMax := c.Query("balanceMax"); balanceMax != "" {
		if value, err := strconv.ParseFloat(balanceMax, 32); err == nil {
			filter.BalanceMax = float32(value)
		} else {
			h.Logger.ERROR.Println("Invalid balanceMax", err)
			c.JSON(400, "Invalid balanceMax value")
			return
		}
	}

	// Convert limit and offset to int32 and assign to the Filter field
	if limit := c.Query("limit"); limit != "" {
		if value, err := strconv.Atoi(limit); err == nil {
			filter.Filter.Limit = int32(value)
		} else {
			h.Logger.ERROR.Println("Invalid limit", err)
			c.JSON(400, "Invalid limit value")
			return
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if value, err := strconv.Atoi(offset); err == nil {
			filter.Filter.Offset = int32(value)
		} else {
			h.Logger.ERROR.Println("Invalid offset", err)
			c.JSON(400, "Invalid offset value")
			return
		}
	}

	// Generate cache key
	cacheKey := fmt.Sprintf("account_list_%s", filter.String())

	// Check cache
	val, err := h.Rdc.Get(cacheKey).Result()
	if err == nil && val != "" {
		var list pb.AccounList
		if err := protojson.Unmarshal([]byte(val), &list); err == nil {
			c.JSON(200, &list)
			return
		}
	}

	// Call the service method
	resp, err := h.Clients.Account.ListAccounts(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list accounts", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Marshal response for caching
	data, err := protojson.Marshal(resp)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal proto message", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the response
	if err := h.Rdc.Set(cacheKey, data, time.Minute*15).Err(); err != nil {
		h.Logger.ERROR.Println("Failed to cache account list", err)
	}

	c.JSON(200, resp)
}

// @Summary Delete Account
// @Description Delete an account by ID
// @Tags Account
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 200 {string} string "Account deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /account/delete/{id} [delete]
func (h *Handler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")

	req := &pb.ById{Id: id}
	_, err := h.Clients.Account.DeleteAccount(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete account:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	if err := h.Rdc.Del(fmt.Sprintf("account_%s", id)).Err(); err != nil {
		h.Logger.ERROR.Println("Failed to remove cache entry:", err)
	}

	c.JSON(200, "Account deleted successfully")
}
