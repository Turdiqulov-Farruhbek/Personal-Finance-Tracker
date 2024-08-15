// tests/goal_repo_test.go

package tests

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"
// 	"gitlab.com/saladin2098/finance_tracker1/budget/internal/storage/repo"

// 	// "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/repo"
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
// )

// func setupTestDB(t *testing.T) *mongo.Database {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	db := mt.DB
// 	return db
// }

// func TestGoalRepo_CreateGoal(t *testing.T) {
// 	collection := setupTestDB(t)
// 	repo := repo.NewGoalRepo(collection)

// 	req := &genproto.GoalCreate{
// 		UserId:        "user123",
// 		Name:          "New Goal",
// 		TargetAmount:  1000.0,
// 		CurrentAmount: 100.0,
// 		Deadline:      "2024-12-31",
// 	}

// 	_, err := repo.CreateGoal(req)
// 	assert.NoError(t, err)

// 	var result genproto.GoalGet
// 	err = collection.FindOne(nil, bson.M{"user_id": req.UserId}).Decode(&result)
// 	assert.NoError(t, err)
// 	assert.Equal(t, req.UserId, result.UserId)
// 	assert.Equal(t, req.Name, result.Name)
// }

// func TestGoalRepo_GetGoal(t *testing.T) {
// 	collection := setupTestDB(t)
// 	repo := repo.NewGoalRepo(collection)

// 	// Insert a document to test the GetGoal method
// 	goal := bson.M{
// 		"user_id":        "user123",
// 		"name":           "Test Goal",
// 		"target_amount":  1000.0,
// 		"current_amount": 200.0,
// 		"deadline":       "2024-12-31",
// 		"created_at":     time.Now().Format(time.RFC3339),
// 		"updated_at":     time.Now().Format(time.RFC3339),
// 	}
// 	res, err := collection.InsertOne(nil, goal)
// 	assert.NoError(t, err)

// 	// Convert the inserted ID to string
// 	id := res.InsertedID.(primitive.ObjectID).Hex()

// 	// Test the GetGoal method
// 	req := &genproto.ById{Id: id}
// 	goalRes, err := repo.GetGoal(req)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Test Goal", goalRes.Name)
// 	assert.Equal(t, "user123", goalRes.UserId)
// }

// func TestGoalRepo_DeleteGoal(t *testing.T) {
// 	collection := setupTestDB(t)
// 	repo := repo.NewGoalRepo(collection)

// 	// Insert a document to test the DeleteGoal method
// 	goal := bson.M{
// 		"user_id":        "user123",
// 		"name":           "Test Goal",
// 		"target_amount":  1000.0,
// 		"current_amount": 200.0,
// 		"deadline":       "2024-12-31",
// 		"created_at":     time.Now().Format(time.RFC3339),
// 		"updated_at":     time.Now().Format(time.RFC3339),
// 	}
// 	res, err := collection.InsertOne(context.TODO(), goal)
// 	assert.NoError(t, err)

// 	// Convert the inserted ID to string
// 	id := res.InsertedID.(primitive.ObjectID).Hex()

// 	// Test the DeleteGoal method
// 	_, err = repo.DeleteGoal(&genproto.ById{Id: id})
// 	assert.NoError(t, err)

// 	var deletedGoal genproto.GoalGet
// 	err = collection.FindOne(nil, bson.M{"_id": res.InsertedID}).Decode(&deletedGoal)
// 	assert.NoError(t, err)
// 	assert.True(t, deletedGoal.DeletedAt > 0) // Check if DeletedAt is set
// }
