package helper

import (
	"time"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePeriodFilter(budget *pb.BudgetGet, userId string) bson.M {
	now := time.Now()
	startDate := now
	switch budget.Period {
	case "daily":
		startDate = now.Truncate(24 * time.Hour)
	case "weekly":
		startDate = now.AddDate(0, 0, -int(now.Weekday()))
	case "monthly":
		startDate = now.AddDate(0, 0, -now.Day()+1)
	case "yearly":
		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	}

	return bson.M{
		"user_id":    userId,
		"time":       bson.M{"$gte": startDate.Format(time.RFC3339)},
		"deleted_at": 0,
	}
}
