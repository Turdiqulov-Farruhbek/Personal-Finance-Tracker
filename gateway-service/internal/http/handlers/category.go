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

// CreateCategory creates a new category
// @Summary 		Create Category
// @Description 	Create a new category
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		category body    pb.CategoryCreate  true   "Category data"
// @Success 		200  {string}  string "Category created successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/create [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	var req pb.CategoryCreate
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
	if err := h.Producer.ProduceMessages("category-create", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Category created successfully")
}

// UpdateCategory updates an existing category
// @Summary 		Update Category
// @Description 	Update an existing category
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Category ID"
// @Param     		category    body    pb.CategoryCreate  true   "Category data"
// @Success 		200  {string}  string "Category updated successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/{id} [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var body pb.CategoryCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, "Invalid request"+err.Error())
		return
	}

	// Set the category ID
	req := pb.CategoryUpdate{
		Id:       categoryID,
		Category: &body,
	}

	// Marshal the request to JSON
	data, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Produce message to Kafka
	if err := h.Producer.ProduceMessages("category-update", data); err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	c.JSON(200, "Category updated successfully")
}

// GetCategory retrieves a category by ID
// @Summary 		Get Category
// @Description 	Retrieve a category by its ID
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Category ID"
// @Success 		200  {object}  pb.CategoryGet "Category details"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Category not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/{id} [get]
func (h *Handler) GetCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Check cache first
	cached, err := h.Rdc.Get(categoryID).Result()
	if err == nil {
		var category pb.CategoryGet
		if err := protojson.Unmarshal([]byte(cached), &category); err == nil {
			h.Logger.ERROR.Println(err)
			c.JSON(200, &category)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	req := &pb.ById{Id: categoryID}
	res, err := h.Clients.Category.GetCategory(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get category:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(categoryID, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// ListCategories retrieves a list of categories
// @Summary List Categories
// @Description Retrieve a list of categories
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string false "User ID"
// @Param name query string false "Category Name"
// @Param type query string false "Category Type"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.CategoryList "List of categories"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /categories/list [get]
func (h *Handler) ListCategories(c *gin.Context) {
	// Initialize filter and set default values
	var filter pb.CategoryFilter
	filter.Filter = &pb.Filter{}

	// Read each query parameter individually
	if userId := c.Query("user_id"); userId != "" {
		filter.UserId = userId
	}
	if name := c.Query("name"); name != "" {
		filter.Name = name
	}
	if categoryType := c.Query("type"); categoryType != "" {
		filter.Type = categoryType
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
	cacheKey := fmt.Sprintf("category-list:%v", filter.String())
	cached, err := h.Rdc.Get(cacheKey).Result()
	if err == nil {
		var categoryList pb.CategoryList
		if err := protojson.Unmarshal([]byte(cached), &categoryList); err == nil {
			c.JSON(200, &categoryList)
			return
		}
	}

	// If not in cache, retrieve from gRPC
	res, err := h.Clients.Category.ListCategories(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list categories:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Cache the result
	data, _ := protojson.Marshal(res)
	_ = h.Rdc.Set(cacheKey, string(data), time.Minute*15).Err()

	c.JSON(200, res)
}

// DeleteCategory deletes a category by ID
// @Summary 		Delete Category
// @Description 	Delete a category by its ID
// @Tags 			Category
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id path    string       true   "Category ID"
// @Success 		200  {string}  string "Category deleted successfully"
// @Failure 		400  {string}  string "Invalid request"
// @Failure 		404  {string}  string "Category not found"
// @Failure 		500  {string}  string "Internal server error"
// @Router 			/categories/{id} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Call gRPC delete method
	req := &pb.ById{Id: categoryID}
	_, err := h.Clients.Category.DeleteCategory(c, req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete category:", err)
		c.JSON(500, "Internal server error"+err.Error())
		return
	}

	// Clear cache
	if err := h.Rdc.Del(categoryID).Err(); err != nil {
		h.Logger.WARN.Println("Failed to delete cache entry:", err)
	}

	c.JSON(200, "Category deleted successfully")
}
