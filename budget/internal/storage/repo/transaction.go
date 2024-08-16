package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "finance_tracker/budget/internal/pkg/genproto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionRepo struct {
	mdb *mongo.Collection
}

func NewTransactionRepo(mdb *mongo.Database) *TransactionRepo {
	db := mdb.Collection("transaction")
	return &TransactionRepo{mdb: db}
}

func (r *TransactionRepo) CreateTransaction(req *pb.TransactionCreate) (*pb.Void, error) {
	now := time.Now().Format(time.RFC3339)

	transaction := bson.M{
		"user_id":     req.UserId,
		"account_id":  req.AccountId,
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
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
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

	filter := bson.M{"_id": obj_id, "deleted_at": 0}
	update := bson.M{"$set": updateFields}

	_, err = r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *TransactionRepo) DeleteTransaction(req *pb.ById) (*pb.Void, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	log.Println("DeleteTransaction id ", req.Id)
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
func (r *TransactionRepo) GetTransaction(req *pb.ById) (*pb.TransactionGet, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	var transaction pb.TransactionGet
	filter := bson.M{"_id": obj_id, "deleted_at": 0}
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

	var tra Transaction
	err = r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&tra)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	transaction.Id = req.Id
	transaction.UserId = tra.UserId
	transaction.AccountId = tra.AccountId
	transaction.CategoryId = tra.CategoryId
	transaction.Amount = tra.Amount
	transaction.Type = tra.Type
	transaction.Description = tra.Description
	transaction.Time = tra.Time

	return &transaction, nil
}
func (r *TransactionRepo) ListTransactions(req *pb.TransactionFilter) (*pb.TransactionList, error) {
	log.Println("transaction list req", req)
	log.Println("transaction user id ", req.UserId)
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
		var tra Transaction
		if err := cursor.Decode(&tra); err != nil {
			return nil, err
		}
		// Map MongoDB Object ID to proto Id field
		transaction.Id = cursor.Current.Lookup("_id").ObjectID().Hex()
		transaction.UserId = tra.UserId
		transaction.AccountId = tra.AccountId
		transaction.CategoryId = tra.CategoryId
		transaction.Amount = tra.Amount
		transaction.Type = tra.Type
		transaction.Description = tra.Description
		transaction.Time = tra.Time

		transactions = append(transactions, &transaction)
	}
	log.Println("transactions after tranaction list", transactions)

	totalCount := len(transactions)

	return &pb.TransactionList{
		TransactionGet: transactions,
		TotalCount:     int32(totalCount),
		Limit:          req.Filter.Limit,
		Offset:         req.Filter.Offset,
	}, nil
}

type Transaction struct {
	UserId      string  `bson:"user_id"`
	AccountId   string  `bson:"account_id"`
	CategoryId  string  `bson:"category_id"`
	Amount      float32 `bson:"amount"`
	Type        string  `bson:"type"`
	Description string  `bson:"description"`
	Time        string  `bson:"time"`
}
