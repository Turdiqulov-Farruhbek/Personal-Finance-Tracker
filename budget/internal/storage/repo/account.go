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

type AccountRepo struct {
	mdb *mongo.Collection
}

func NewAccountRepo(mdb *mongo.Database) *AccountRepo {
	db := mdb.Collection("account")
	return &AccountRepo{mdb: db}
}
func (r *AccountRepo) CreateAccount(req *pb.AccountCreate) (*pb.Void, error) {
	now := time.Now().Format(time.RFC3339)

	account := bson.M{
		"user_id":    req.UserId,
		"name":       req.Name,
		"type":       req.Type,
		"currency":   req.Currency,
		"balance":    0.0, // Assuming initial balance is 0
		"created_at": now,
		"updated_at": now,
		"deleted_at": 0,
	}

	_, err := r.mdb.InsertOne(context.Background(), account)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *AccountRepo) GetAccount(req *pb.ById) (*pb.AccountGet, error) {
	var account pb.AccountGet
	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	projection := bson.M{
		"_id":        1,
		"user_id":    1,
		"name":       1,
		"type":       1,
		"currency":   1,
		"balance":    1,
		"created_at": 1,
		"updated_at": 1,
	}

	err := r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	account.Id = req.Id

	return &account, nil
}

func (r *AccountRepo) UpdateAccount(req *pb.AccountUpdate) (*pb.Void, error) {
	updateFields := bson.M{}
	if req.Body.Name != "" {
		updateFields["name"] = req.Body.Name
	}
	if req.Body.Type != "" {
		updateFields["type"] = req.Body.Type
	}
	if req.Body.Currency != "" {
		updateFields["currency"] = req.Body.Currency
	}
	if req.Body.Balance != 0 {
		updateFields["balance"] = req.Body.Balance
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

func (r *AccountRepo) DeleteAccount(req *pb.ById) (*pb.Void, error) {
	deletedAt := time.Now().Unix()

	filter := bson.M{"_id": req.Id, "deleted_at": 0}
	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

	_, err := r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *AccountRepo) ListAccounts(req *pb.AccountFilter) (*pb.AccounList, error) {
	filter := bson.M{"deleted_at": 0}

	if req.Name != "" {
		filter["name"] = req.Name
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}
	if req.Currency != "" {
		filter["currency"] = req.Currency
	}
	if req.UserId!= "" {
        filter["user_id"] = req.UserId
    }
	if req.BalanceMin != 0 || req.BalanceMax != 0 {
		filter["balance"] = bson.M{}
		if req.BalanceMin != 0 {
			filter["balance"].(bson.M)["$gte"] = req.BalanceMin
		}
		if req.BalanceMax != 0 {
			filter["balance"].(bson.M)["$lte"] = req.BalanceMax
		}
	}

	options := options.Find()
	options.SetLimit(int64(req.Filter.Limit))
	options.SetSkip(int64(req.Filter.Offset))

	// Define projection to exclude 'deleted_at' and include only necessary fields
	projection := bson.M{
		"_id":        1,
		"user_id":    1,
		"name":       1,
		"type":       1,
		"currency":   1,
		"balance":    1,
		"created_at": 1,
		"updated_at": 1,
	}

	options.SetProjection(projection)

	cursor, err := r.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var accounts []*pb.AccountGet
	for cursor.Next(context.Background()) {
		var account pb.AccountGet
		if err := cursor.Decode(&account); err != nil {
			return nil, err
		}
		// Map MongoDB Object ID to proto Id field
		account.Id = cursor.Current.Lookup("_id").ObjectID().Hex()

		accounts = append(accounts, &account)
	}

	totalCount, err := r.mdb.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return &pb.AccounList{
		Accounts:   accounts,
		TotalCount: int32(totalCount),
		Limit:      req.Filter.Limit,
		Offset:     req.Filter.Offset,
	}, nil
}
