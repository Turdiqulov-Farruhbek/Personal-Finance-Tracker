package repo

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	pb "finance_tracker/budget/internal/pkg/genproto"
// 	"finance_tracker/budget/internal/usecases/kafka"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"google.golang.org/protobuf/encoding/protojson"
// )

// type TransactionRepo struct {
// 	mdb          *mongo.Collection
// 	accountDB    *mongo.Collection
// 	goalDB       *mongo.Collection
// 	budgetDB     *mongo.Collection
// 	notification *kafka.Producer
// }

// func NewTransactionRepo(mdb *mongo.Database, notif_client *kafka.Producer) *TransactionRepo {
// 	db := mdb.Collection("transaction")
// 	goal := mdb.Collection("goal")
// 	budget := mdb.Collection("budget")
// 	account := mdb.Collection("account")
// 	return &TransactionRepo{mdb: db,
// 		goalDB:       goal,
// 		budgetDB:     budget,
// 		accountDB:    account,
// 		notification: notif_client,
// 	}
// }

// func (r *TransactionRepo) CreateTransaction(req *pb.TransactionCreate) (*pb.Void, error) {
// 	// Check the current balance before proceeding
// 	account := &pb.AccountGet{}
// 	filter := bson.M{"_id": req.AccountId, "deleted_at": 0}
// 	err := r.accountDB.FindOne(context.TODO(), filter).Decode(account)
// 	if err != nil {
// 		return nil, err
// 	}

// 	currentBalance := account.Balance
// 	if req.Type == "debit" {
// 		currentBalance -= req.Amount
// 	} else {
// 		currentBalance += req.Amount
// 	}

// 	if currentBalance < 0 {
// 		return nil, fmt.Errorf("insufficient funds in the account")
// 	}

// 	// Update the account balance
// 	update := bson.M{"$set": bson.M{"balance": currentBalance}}
// 	_, err = r.accountDB.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Insert into transactions
// 	now := time.Now().Format(time.RFC3339)

// 	transaction := bson.M{
// 		"user_id":     req.UserId,
// 		"account_id":  req.AccountId,
// 		"category_id": req.CategoryId,
// 		"amount":      req.Amount,
// 		"type":        req.Type,
// 		"description": req.Description,
// 		"time":        now,
// 		"created_at":  now,
// 		"updated_at":  now,
// 		"deleted_at":  0,
// 	}

// 	_, err = r.mdb.InsertOne(context.Background(), transaction)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check for budget targets if the transaction is a debit
// 	if req.Type == "credit" {
// 		budgetFilter := bson.M{"user_id": req.UserId, "end_date": bson.M{"$gt": time.Now().Format("2006-01-02")}, "deleted_at": 0}
// 		cursor, err := r.budgetDB.Find(context.TODO(), budgetFilter)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer cursor.Close(context.Background())

// 		var budget pb.BudgetGet
// 		for cursor.Next(context.TODO()) {
// 			if err := cursor.Decode(&budget); err != nil {
// 				return nil, err
// 			}
// 			periodFilter := r.createPeriodFilter(&budget, req.UserId)
// 			var userSpendings float32
// 			cursor, err := r.mdb.Find(context.TODO(), periodFilter)
// 			if err != nil {
// 				return nil, err
// 			}
// 			defer cursor.Close(context.Background())

// 			for cursor.Next(context.Background()) {
// 				var txn pb.TransactionGet
// 				err := cursor.Decode(&txn)
// 				if err != nil {
// 					return nil, err
// 				}
// 				userSpendings += txn.Amount
// 			}

// 			if userSpendings > budget.Amount {
// 				message := fmt.Sprintf("Warning: Your spending in the %s budget category has exceeded the limit of $%v", budget.Id, budget.Amount)
// 				notification := pb.NotificationCreate{
// 					RecieverId: req.UserId,
// 					Message:    message,
// 				}
// 				input, err := protojson.Marshal(&notification)
// 				if err != nil {
// 					return nil, err
// 				}
// 				err = r.notification.ProduceMessages("notification-create", input)
// 				if err != nil {
// 					return nil, err
// 				}
// 			}

// 		}
// 	}
// 	goalFilter := bson.M{"user_id": req.UserId,
// 						"end_date":   bson.M{"$gt": time.Now()},
// 						"deleted_at": 0}
// 	var goal pb.GoalGet
// 	r.goalDB.FindOne(context.TODO(),goalFilter).Decode(&goal)
// 	if req.Type == "debit" {
// 		goalProgress := goal.CurrentAmount + req.Amount
//         if goalProgress >= goal.TargetAmount {
//             message := fmt.Sprintf("Congratulations! You have reached your goal of $%v", goal.TargetAmount)
//             notification := pb.NotificationCreate{
//                 RecieverId: req.UserId,
//                 Message:    message,
//             }
//             input, err := protojson.Marshal(&notification)
//             if err!= nil {
//                 return nil, err
//             }
//             err = r.notification.ProduceMessages("notification-create", input)
//             if err!= nil {
//                 return nil, err
//             }
//         }
//         update := bson.M{"$set": bson.M{"current_amount": goalProgress}}
//         _, err = r.goalDB.UpdateOne(context.Background(), goalFilter, update)
//         if err!= nil {
//             return nil, err
//         }
// 	}
// 	if req.Type == "credit" {
// 		goalProgress := goal.CurrentAmount - req.Amount
//         if goalProgress < 0 {
//             goalProgress = 0
//         }
//         update := bson.M{"$set": bson.M{"current_amount": goalProgress}}
//         _, err = r.goalDB.UpdateOne(context.Background(), goalFilter, update)
//         if err!= nil {
//             return nil, err
//         }
// 	}

// 	// Create a transaction notification message
// 	var message string
// 	if req.Type == "debit" {
// 		message = fmt.Sprintf("Debit: -$%v amount debited from your account, your current balance is $%v", req.Amount, currentBalance)
// 	} else {
// 		message = fmt.Sprintf("Credit: +$%v amount credited to your account, your current balance is $%v", req.Amount, currentBalance)
// 	}
// 	body := pb.NotificationCreate{
// 		RecieverId: req.UserId,
// 		Message:    message,
// 	}
// 	input, err := protojson.Marshal(&body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Write to Kafka topic
// 	err = r.notification.ProduceMessages("notification-create", input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Void{}, nil
// }
// func (r *TransactionRepo) UpdateTransaction(req *pb.TransactionUpdate) (*pb.Void, error) {
// 	updateFields := bson.M{}
// 	if req.Body.AccountId != "" {
// 		updateFields["account_id"] = req.Body.AccountId
// 	}
// 	if req.Body.CategoryId != "" {
// 		updateFields["category_id"] = req.Body.CategoryId
// 	}
// 	if req.Body.Amount != 0 {
// 		updateFields["amount"] = req.Body.Amount
// 	}
// 	if req.Body.Type != "" {
// 		updateFields["type"] = req.Body.Type
// 	}
// 	if req.Body.Description != "" {
// 		updateFields["description"] = req.Body.Description
// 	}
// 	updateFields["updated_at"] = time.Now().Format(time.RFC3339)

// 	filter := bson.M{"_id": req.Id, "deleted_at": 0}
// 	update := bson.M{"$set": updateFields}

// 	_, err := r.mdb.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Void{}, nil
// }

// func (r *TransactionRepo) DeleteTransaction(req *pb.ById) (*pb.Void, error) {
// 	deletedAt := time.Now().Unix()

// 	filter := bson.M{"_id": req.Id, "deleted_at": 0}
// 	update := bson.M{"$set": bson.M{"deleted_at": deletedAt}}

// 	_, err := r.mdb.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.Void{}, nil
// }
// func (r *TransactionRepo) GetTransaction(req *pb.ById) (*pb.TransactionGet, error) {
// 	var transaction pb.TransactionGet
// 	filter := bson.M{"_id": req.Id, "deleted_at": 0}
// 	projection := bson.M{
// 		"_id":         1,
// 		"user_id":     1,
// 		"account_id":  1,
// 		"category_id": 1,
// 		"amount":      1,
// 		"type":        1,
// 		"description": 1,
// 		"time":        1,
// 	}

// 	err := r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&transaction)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, fmt.Errorf("transaction not found")
// 		}
// 		return nil, err
// 	}

// 	// Map MongoDB Object ID to proto Id field
// 	transaction.Id = req.Id

// 	return &transaction, nil
// }
// func (r *TransactionRepo) ListTransactions(req *pb.TransactionFilter) (*pb.TransactionList, error) {
// 	filter := bson.M{"deleted_at": 0}

// 	if req.UserId != "" {
// 		filter["user_id"] = req.UserId
// 	}
// 	if req.AccountId != "" {
// 		filter["account_id"] = req.AccountId
// 	}
// 	if req.CategoryId != "" {
// 		filter["category_id"] = req.CategoryId
// 	}
// 	if req.Type != "" {
// 		filter["type"] = req.Type
// 	}
// 	if req.Description != "" {
// 		filter["description"] = req.Description
// 	}
// 	if req.TimeFrom != "" || req.TimeTo != "" {
// 		timeFilter := bson.M{}
// 		if req.TimeFrom != "" {
// 			timeFilter["$gte"] = req.TimeFrom
// 		}
// 		if req.TimeTo != "" {
// 			timeFilter["$lte"] = req.TimeTo
// 		}
// 		filter["time"] = timeFilter
// 	}
// 	if req.AmountFrom != 0 || req.AmountTo != 0 {
// 		amountFilter := bson.M{}
// 		if req.AmountFrom != 0 {
// 			amountFilter["$gte"] = req.AmountFrom
// 		}
// 		if req.AmountTo != 0 {
// 			amountFilter["$lte"] = req.AmountTo
// 		}
// 		filter["amount"] = amountFilter
// 	}

// 	options := options.Find()
// 	options.SetLimit(int64(req.Filter.Limit))
// 	options.SetSkip(int64(req.Filter.Offset))

// 	// Define projection to exclude 'deleted_at' and include only necessary fields
// 	projection := bson.M{
// 		"_id":         1,
// 		"user_id":     1,
// 		"account_id":  1,
// 		"category_id": 1,
// 		"amount":      1,
// 		"type":        1,
// 		"description": 1,
// 		"time":        1,
// 	}

// 	options.SetProjection(projection)

// 	cursor, err := r.mdb.Find(context.Background(), filter, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.Background())

// 	var transactions []*pb.TransactionGet
// 	for cursor.Next(context.Background()) {
// 		var transaction pb.TransactionGet
// 		if err := cursor.Decode(&transaction); err != nil {
// 			return nil, err
// 		}
// 		// Map MongoDB Object ID to proto Id field
// 		transaction.Id = cursor.Current.Lookup("_id").ObjectID().Hex()

// 		transactions = append(transactions, &transaction)
// 	}

// 	totalCount, err := r.mdb.CountDocuments(context.Background(), filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.TransactionList{
// 		TransactionGet: transactions,
// 		TotalCount:     int32(totalCount),
// 		Limit:          req.Filter.Limit,
// 		Offset:         req.Filter.Offset,
// 	}, nil
// }

// func (r *TransactionRepo) createPeriodFilter(budget *pb.BudgetGet, userId string) bson.M {
// 	now := time.Now()
// 	startDate := now
// 	switch budget.Period {
// 	case "daily":
// 		startDate = now.Truncate(24 * time.Hour)
// 	case "weekly":
// 		startDate = now.AddDate(0, 0, -int(now.Weekday()))
// 	case "monthly":
// 		startDate = now.AddDate(0, 0, -now.Day()+1)
// 	case "yearly":
// 		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
// 	}

// 	return bson.M{
// 		"user_id":    userId,
// 		"time":       bson.M{"$gte": startDate.Format(time.RFC3339)},
// 		"deleted_at": 0,
// 	}
// }
