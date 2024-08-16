package http

import (
	_ "finance_tracker/gateway/docs"
	"finance_tracker/gateway/internal/http/handlers"
	"finance_tracker/gateway/internal/http/middlerware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Finance Tracker API Documentation
// @version 1.0
// @description API for Instant Delivery resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	enforcer, err := casbin.NewEnforcer("./internal/http/casbin/model.conf", "./internal/http/casbin/policy.csv")
	if err != nil {
		panic(err)
	}
	router.Use(middlerware.NewAuth(enforcer))

	account := router.Group("/account")
	{
		account.POST("/create", h.CreateAcccount)
		account.GET("/:id", h.GetAccount)
		account.PUT("/update/:id", h.UpdateAccount)
		account.DELETE("/delete/:id", h.DeleteAccount)
		account.GET("/list", h.ListAccounts)
	}
	budget := router.Group("/budgets")
	{
		budget.POST("/create", h.CreateBudget)
		budget.GET("/:id", h.GetBudget)
		budget.PUT("/:id", h.UpdateBudget)
		budget.DELETE("/delete/:id", h.DeleteBudget)
		budget.GET("/list", h.ListBudgets)
	}
	category := router.Group("/categories")
	{
		category.POST("/create", h.CreateCategory)
		category.GET("/:id", h.GetCategory)
		category.PUT("/:id", h.UpdateCategory)
		category.DELETE("/:id", h.DeleteCategory)
		category.GET("/list", h.ListCategories)
	}
	goal := router.Group("/goals")
	{
		goal.POST("/create", h.CreateGoal)
		goal.GET("/:goal_id", h.GetGoal)
		goal.PUT("/:goal_id", h.UpdateGoal)
		goal.DELETE("/:goal_id", h.DeleteGoal)
		goal.GET("/list", h.ListGoals)
	}
	notification := router.Group("/notifications")
	{
		notification.POST("/create", h.CreateNotification)
		notification.GET("/:notification_id", h.GetNotification)
		notification.PUT("/:notification_id", h.UpdateNotification)
		notification.DELETE("/:notification_id", h.DeleteNotification)
		notification.GET("/list", h.ListNotifications)
		notification.POST("notify-all", h.NotifyAll)
	}
	report := router.Group("/report")
	{
		report.GET("/spendings", h.GetSpendings)
		report.GET("/incomes", h.GetIncomes)
		report.GET("/budget-performance", h.BudgetPerformance)
		report.GET("/goal-progress", h.GoalProgress)
	}
	transaction := router.Group("/transactions")
	{
		transaction.POST("/create", h.CreateTransaction)
		transaction.GET("/:id", h.GetTransaction)
		transaction.PUT("/:transaction_id", h.UpdateTransaction)
		transaction.DELETE("/:id", h.DeleteTransaction)
		transaction.GET("/list", h.ListTransactions)
	}

	return router
}
