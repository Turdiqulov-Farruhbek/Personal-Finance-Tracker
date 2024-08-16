package repo

import (
	"context"
	"fmt"
	"time"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GoalRepo struct {
	mdb *mongo.Collection
}

func NewGoalRepo(mdb *mongo.Database) *GoalRepo {
	db := mdb.Collection("goal")
	return &GoalRepo{mdb: db}
}

func (r *GoalRepo) CreateGoal(req *pb.GoalCreate) (*pb.Void, error) {
	now := time.Now().Format(time.RFC3339)

	goal := bson.M{
		"user_id":        req.UserId,
		"name":           req.Name,
		"target_amount":  req.TargetAmount,
		"current_amount": req.CurrentAmount,
		"deadline":       req.Deadline,
		"status":         "active", // default status
		"created_at":     now,
		"updated_at":     now,
		"deleted_at":     0,
	}

	_, err := r.mdb.InsertOne(context.Background(), goal)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *GoalRepo) UpdateGoal(req *pb.GoalUpdate) (*pb.Void, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	updateFields := bson.M{}
	if req.Body.Name != "" {
		updateFields["name"] = req.Body.Name
	}
	if req.Body.TargetAmount != 0 {
		updateFields["target_amount"] = req.Body.TargetAmount
	}
	if req.Body.CurrentAmount != 0 {
		updateFields["current_amount"] = req.Body.CurrentAmount
	}
	if req.Body.Deadline != "" {
		updateFields["deadline"] = req.Body.Deadline
	}
	if req.Body.Status != "" {
		updateFields["status"] = req.Body.Status
	}
	updateFields["updated_at"] = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": obj_id, "deleted_at": 0}
	update := bson.M{"$set": updateFields}

	_, err = r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *GoalRepo) DeleteGoal(req *pb.ById) (*pb.Void, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	deletedAt := time.Now().Unix()

	filter := bson.M{"_id": obj_id, "deleted_at": 0}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err = r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *GoalRepo) GetGoal(req *pb.ById) (*pb.GoalGet, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	var goal pb.GoalGet
	filter := bson.M{"_id": obj_id, "deleted_at": 0}
	projection := bson.M{
		"_id":            1,
		"user_id":        1,
		"name":           1,
		"target_amount":  1,
		"current_amount": 1,
		"deadline":       1,
		"status":         1,
		"created_at":     1,
		"updated_at":     1,
	}

	var goal_model Goal
	err = r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&goal_model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("goal not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	goal.Id = req.Id
	goal.UserId = goal_model.UserId
	goal.Name = goal_model.Name
	goal.TargetAmount = goal_model.TargetAmount
	goal.CurrentAmount = goal_model.CurrentAmount
	goal.Deadline = goal_model.Deadline
	goal.Status = goal_model.Status
	goal.CreatedAt = goal_model.CreatedAt.String()
	goal.UpdatedAt = goal_model.UpdatedAt.String()

	return &goal, nil
}
func (r *GoalRepo) ListGoals(req *pb.GoalFilter) (*pb.GoalList, error) {
	filter := bson.M{"deleted_at": 0}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.Status != "" {
		filter["status"] = req.Status
	}
	if req.Name != "" {
		filter["name"] = bson.M{"$regex": req.Name, "$options": "i"}
	}
	if req.TargetFrom != 0 || req.TargetTo != 0 {
		targetFilter := bson.M{}
		if req.TargetFrom != 0 {
			targetFilter["$gte"] = req.TargetFrom
		}
		if req.TargetTo != 0 {
			targetFilter["$lte"] = req.TargetTo
		}
		filter["target_amount"] = targetFilter
	}
	if req.DeadlineFrom != "" || req.DeadlineTo != "" {
		deadlineFilter := bson.M{}
		if req.DeadlineFrom != "" {
			deadlineFilter["$gte"] = req.DeadlineFrom
		}
		if req.DeadlineTo != "" {
			deadlineFilter["$lte"] = req.DeadlineTo
		}
		filter["deadline"] = deadlineFilter
	}

	options := options.Find()
	options.SetLimit(int64(req.Filter.Limit))
	options.SetSkip(int64(req.Filter.Offset))

	projection := bson.M{
		"_id":            1,
		"user_id":        1,
		"name":           1,
		"target_amount":  1,
		"current_amount": 1,
		"deadline":       1,
		"status":         1,
		"created_at":     1,
		"updated_at":     1,
	}

	options.SetProjection(projection)

	cursor, err := r.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var goals []*pb.GoalGet
	for cursor.Next(context.Background()) {
		var goal pb.GoalGet
		var goal_model Goal
		if err := cursor.Decode(&goal_model); err != nil {
			return nil, err
		}
		goal.Id = cursor.Current.Lookup("_id").ObjectID().Hex()
		goal.UserId = goal_model.UserId
		goal.Name = goal_model.Name
		goal.TargetAmount = goal_model.TargetAmount
		goal.CurrentAmount = goal_model.CurrentAmount
		goal.Deadline = goal_model.Deadline
		goal.Status = goal_model.Status
		goal.CreatedAt = goal_model.CreatedAt.String()
		goal.UpdatedAt = goal_model.UpdatedAt.String()

		goals = append(goals, &goal)
	}

	totalCount, err := r.mdb.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return &pb.GoalList{
		Goals:      goals,
		TotalCount: int32(totalCount),
		Limit:      req.Filter.Limit,
		Offset:     req.Filter.Offset,
	}, nil
}

type Goal struct {
	UserId        string    `bson:"user_id"`
	Name          string    `bson:"name"`
	TargetAmount  float32   `bson:"target_amount"`
	CurrentAmount float32   `bson:"current_amount"`
	Deadline      string    `bson:"deadline"`
	Status        string    `bson:"status"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}
