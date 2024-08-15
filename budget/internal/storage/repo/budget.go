package repo

import (
	"context"
	"fmt"
	"time"

	pb "gitlab.com/saladin2098/finance_tracker1/budget/internal/pkg/genproto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BudgetRepo struct {
	mdb *mongo.Collection
}

func NewBudgetRepo(mdb *mongo.Database) *BudgetRepo {
    db := mdb.Collection("budget")
    return &BudgetRepo{mdb: db}
}
func (r *BudgetRepo) CreateBudget(req *pb.BudgetCreate) (*pb.Void, error) {
	now := time.Now().Format(time.RFC3339)

	budget := bson.M{
		"user_id":     req.UserId,
		"category_id": req.CategoryId,
		"amount":      req.Amount,
		"period":      req.Period,
		"start_date":  req.StartDate,
		"end_date":    req.EndDate,
		"created_at":  now,
		"updated_at":  now,
		"deleted_at":  0,
	}

	_, err := r.mdb.InsertOne(context.Background(), budget)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *BudgetRepo) UpdateBudget(req *pb.BudgetUpdate) (*pb.Void, error) {
	updateFields := bson.M{}
	if req.Body.UserId != "" {
		updateFields["user_id"] = req.Body.UserId
	}
	if req.Body.CategoryId != "" {
		updateFields["category_id"] = req.Body.CategoryId
	}
	if req.Body.Amount != 0 {
		updateFields["amount"] = req.Body.Amount
	}
	if req.Body.Period != "" {
		updateFields["period"] = req.Body.Period
	}
	if req.Body.StartDate != "" {
		updateFields["start_date"] = req.Body.StartDate
	}
	if req.Body.EndDate != "" {
		updateFields["end_date"] = req.Body.EndDate
	}
	updateFields["updated_at"] = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	update := bson.M{"$set": updateFields}

	_, err := r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *BudgetRepo) DeleteBudget(req *pb.ById) (*pb.Void, error) {
	deletedAt := time.Now().Unix()

	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *BudgetRepo) GetBudget(req *pb.ById) (*pb.BudgetGet, error) {
	var budget pb.BudgetGet
	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	projection := bson.M{
		"_id":         1,
		"user_id":     1,
		"category_id": 1,
		"amount":      1,
		"period":      1,
		"start_date":  1,
		"end_date":    1,
	}

	err := r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&budget)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("budget not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	budget.Id = req.Id

	return &budget, nil
}
func (r *BudgetRepo) ListBudgets(req *pb.BudgetFilter) (*pb.BudgetList, error) {
	filter := bson.M{"deleted_at": 0}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.CategoryId != "" {
		filter["category_id"] = req.CategoryId
	}
	if req.Status != "" && req.Status == "completed" {
		filter["end_date"] = bson.M{"$lte": time.Now().Format(time.RFC3339)}
	}
	if req.Status!= "" && req.Status == "ongoing" {
        filter["end_date"] = bson.M{"$gt": time.Now().Format(time.RFC3339)}
    }
	if req.AmountFrom != 0 || req.AmountTo != 0 {
		amountFilter := bson.M{}
		if req.AmountFrom != 0 {
			amountFilter["$gte"] = req.AmountFrom
		}
		if req.AmountTo != 0 {
			amountFilter["$lte"] = req.AmountTo
		}
		filter["amount"] = amountFilter
	}
	if req.Period != "" {
		filter["period"] = req.Period
	}

	options := options.Find()
	options.SetLimit(int64(req.Filter.Limit))
	options.SetSkip(int64(req.Filter.Offset))

	// Define projection to exclude 'deleted_at' and include only necessary fields
	projection := bson.M{
		"_id":         1,
		"user_id":     1,
		"category_id": 1,
		"amount":      1,
		"period":      1,
		"start_date":  1,
		"end_date":    1,
	}

	options.SetProjection(projection)

	cursor, err := r.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var budgets []*pb.BudgetGet
	for cursor.Next(context.Background()) {
		var budget pb.BudgetGet
		if err := cursor.Decode(&budget); err != nil {
			return nil, err
		}
		// Map MongoDB Object ID to proto Id field
		budget.Id = cursor.Current.Lookup("_id").ObjectID().Hex()

		budgets = append(budgets, &budget)
	}

	totalCount, err := r.mdb.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return &pb.BudgetList{
		Budgets:    budgets,
		TotalCount: int32(totalCount),
		Limit:      req.Filter.Limit,
		Offset:     req.Filter.Offset,
	}, nil
}