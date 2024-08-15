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

type TransactionRepo struct {
	mdb *mongo.Collection

}

func NewTransactionRepo(mdb *mongo.Database) *TransactionRepo {
    db := mdb.Collection("transaction")
    return &TransactionRepo{mdb: db,
    }
}

func (r *TransactionRepo) CreateTransaction(req *pb.TransactionCreate) (*pb.Void, error) {
	now := time.Now().Format(time.RFC3339)

	transaction := bson.M{
		"user_id":    req.UserId,
		"account_id": req.AccountId,
		"category_id": req.CategoryId,
		"amount":      req.Amount,
		"type":        req.Type,
		"description": req.Description,
		"time":        now,
		"created_at":  now,
		"updated_at":  now,
		"deleted_at":  0,
	}

	_, err := r.mdb.InsertOne(context.Background(), transaction)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *TransactionRepo) UpdateTransaction(req *pb.TransactionUpdate) (*pb.Void, error) {
	updateFields := bson.M{}
	if req.Body.AccountId != "" {
		updateFields["account_id"] = req.Body.AccountId
	}
	if req.Body.CategoryId != "" {
		updateFields["category_id"] = req.Body.CategoryId
	}
	if req.Body.Amount != 0 {
		updateFields["amount"] = req.Body.Amount
	}
	if req.Body.Type != "" {
		updateFields["type"] = req.Body.Type
	}
	if req.Body.Description != "" {
		updateFields["description"] = req.Body.Description
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

func (r *TransactionRepo) DeleteTransaction(req *pb.ById) (*pb.Void, error) {
	deletedAt := time.Now().Unix()

	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *TransactionRepo) GetTransaction(req *pb.ById) (*pb.TransactionGet, error) {
	var transaction pb.TransactionGet
	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	projection := bson.M{
		"_id":         1,
		"user_id":     1,
		"account_id":  1,
		"category_id": 1,
		"amount":      1,
		"type":        1,
		"description": 1,
		"time":        1,
	}

	err := r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	transaction.Id = req.Id

	return &transaction, nil
}
func (r *TransactionRepo) ListTransactions(req *pb.TransactionFilter) (*pb.TransactionList, error) {
	filter := bson.M{"deleted_at": 0}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.AccountId != "" {
		filter["account_id"] = req.AccountId
	}
	if req.CategoryId != "" {
		filter["category_id"] = req.CategoryId
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}
	if req.Description != "" {
		filter["description"] = req.Description
	}
	if req.TimeFrom != "" || req.TimeTo != "" {
		timeFilter := bson.M{}
		if req.TimeFrom != "" {
			timeFilter["$gte"] = req.TimeFrom
		}
		if req.TimeTo != "" {
			timeFilter["$lte"] = req.TimeTo
		}
		filter["time"] = timeFilter
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

	options := options.Find()
	options.SetLimit(int64(req.Filter.Limit))
	options.SetSkip(int64(req.Filter.Offset))

	// Define projection to exclude 'deleted_at' and include only necessary fields
	projection := bson.M{
		"_id":         1,
		"user_id":     1,
		"account_id":  1,
		"category_id": 1,
		"amount":      1,
		"type":        1,
		"description": 1,
		"time":        1,
	}

	options.SetProjection(projection)

	cursor, err := r.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var transactions []*pb.TransactionGet
	for cursor.Next(context.Background()) {
		var transaction pb.TransactionGet
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		// Map MongoDB Object ID to proto Id field
		transaction.Id = cursor.Current.Lookup("_id").ObjectID().Hex()

		transactions = append(transactions, &transaction)
	}

	totalCount, err := r.mdb.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionList{
		TransactionGet: transactions,
		TotalCount:     int32(totalCount),
		Limit:          req.Filter.Limit,
		Offset:         req.Filter.Offset,
	}, nil
}
