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
		"UserId":    req.UserId,
		"Name":      req.Name,
		"Type":      req.Type,
		"Currency":  req.Currency,
		"Balance":   0.0, // Assuming initial balance is 0
		"CreatedAt": now,
		"UpdatedAt": now,
		"DeletedAt": 0,
	}

	_, err := r.mdb.InsertOne(context.Background(), account)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *AccountRepo) GetAccount(req *pb.ById) (*pb.AccountGet, error) {
	var account pb.AccountGet
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	log.Println(obj_id)
	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	projection := bson.M{
		"_id":       1,
		"UserId":    1,
		"Name":      1,
		"Type":      1,
		"Currency":  1,
		"Balance":   1,
		"CreatedAt": 1,
		"UpdatedAt": 1,
	}
	var acc Account

	err = r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&acc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	account.Id = req.Id
	account.UserId = acc.UserId
	account.Name = acc.Name
	account.Type = acc.Type
	account.Currency = acc.Currency
	account.Balance = acc.Balance
	account.CreatedAt = acc.CreatedAt
	account.UpdatedAt = acc.UpdatedAt

	return &account, nil
}

func (r *AccountRepo) UpdateAccount(req *pb.AccountUpdate) (*pb.Void, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	updateFields := bson.M{}
	if req.Body.Name != "" {
		updateFields["Name"] = req.Body.Name
	}
	if req.Body.Type != "" {
		updateFields["Type"] = req.Body.Type
	}
	if req.Body.Currency != "" {
		updateFields["Currency"] = req.Body.Currency
	}
	if req.Body.Balance != 0 {
		updateFields["Balance"] = req.Body.Balance
	}
	updateFields["UpdatedAt"] = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": updateFields}

	_, err = r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	log.Println("updated account")

	return &pb.Void{}, nil
}

func (r *AccountRepo) DeleteAccount(req *pb.ById) (*pb.Void, error) {
	deletedAt := time.Now().Unix()
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": bson.M{"DeletedAt": deletedAt}}

	_, err = r.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	log.Println("deleted account")

	return &pb.Void{}, nil
}

func (r *AccountRepo) ListAccounts(req *pb.AccountFilter) (*pb.AccounList, error) {
	if r.mdb == nil {
		return nil, fmt.Errorf("mdb is not initialized")
	}

	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}

	if req.Filter == nil {
		return nil, fmt.Errorf("filter is nil")
	}

	filter := bson.M{"DeletedAt": 0}
	log.Println("000000000000000000000000000000000000000", req)

	if req.Name != "" {
		filter["Name"] = req.Name
	}
	if req.Type != "" {
		filter["Type"] = req.Type
	}
	if req.Currency != "" {
		filter["Currency"] = req.Currency
	}
	if req.UserId != "" {
		filter["UserId"] = req.UserId
	}
	if req.BalanceMin != 0 || req.BalanceMax != 0 {
		filter["Balance"] = bson.M{}
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

	projection := bson.M{
		"_id":       1,
		"UserId":    1,
		"Name":      1,
		"Type":      1,
		"Currency":  1,
		"Balance":   1,
		"CreatedAt": 1,
		"UpdatedAt": 1,
	}
	var acc Account

	log.Println(filter)
	options.SetProjection(projection)

	cursor, err := r.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	log.Println(filter)
	defer cursor.Close(context.Background())

	var accounts []*pb.AccountGet
	for cursor.Next(context.Background()) {
		var account pb.AccountGet
		if err := cursor.Decode(&acc); err != nil {
			return nil, err
		}

		account.Id = cursor.Current.Lookup("_id").ObjectID().Hex()
		account.UserId = acc.UserId
		account.Name = acc.Name
		account.Type = acc.Type
		account.Currency = acc.Currency
		account.Balance = acc.Balance
		account.CreatedAt = acc.CreatedAt
		account.UpdatedAt = acc.UpdatedAt

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

type Account struct {
	UserId    string  `bson:"UserId"`
	Name      string  `bson:"Name"`
	Type      string  `bson:"Type"`
	Currency  string  `bson:"Currency"`
	Balance   float32 `bson:"Balance"`
	CreatedAt string  `bson:"CreatedAt"`
	UpdatedAt string  `bson:"UpdatedAt"`
}
