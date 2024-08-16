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

type CategoryRepo struct {
	mdb *mongo.Collection
}

func NewCategoryRepo(mdb *mongo.Database) *CategoryRepo {
	db := mdb.Collection("category")
	return &CategoryRepo{mdb: db}
}

func (r *CategoryRepo) CreateCategory(req *pb.CategoryCreate) (*pb.Void, error) {
	now := time.Now().Format(time.RFC3339)

	category := bson.M{
		"user_id":    req.UserId,
		"name":       req.Name,
		"type":       req.Type,
		"created_at": now,
		"updated_at": now,
		"deleted_at": 0,
	}

	_, err := r.mdb.InsertOne(context.Background(), category)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *CategoryRepo) UpdateCategory(req *pb.CategoryUpdate) (*pb.Void, error) {
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	updateFields := bson.M{}
	if req.Category.Name != "" {
		updateFields["name"] = req.Category.Name
	}
	if req.Category.Type != "" {
		updateFields["type"] = req.Category.Type
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
func (r *CategoryRepo) DeleteCategory(req *pb.ById) (*pb.Void, error) {
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
func (r *CategoryRepo) GetCategory(req *pb.ById) (*pb.CategoryGet, error) {
	log.Println(req.Id)
	obj_id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		log.Println(err)
	}
	log.Println(obj_id)
	var category pb.CategoryGet

	filter := bson.M{"_id": obj_id, "deleted_at": 0}
	projection := bson.M{
		"_id":        1,
		"user_id":    1,
		"name":       1,
		"type":       1,
		"created_at": 1,
		"updated_at": 1,
	}
	var cat Category

	err = r.mdb.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&cat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("category not found")
		}
		return nil, err
	}

	// Map MongoDB Object ID to proto Id field
	category.Id = req.Id
	category.UserId = cat.UserId
	category.Name = cat.Name
	category.Type = cat.Type
	category.CreatedAt = cat.CreatedAt
	category.UpdatedAt = cat.UpdatedAt

	return &category, nil
}

func (r *CategoryRepo) ListCategories(req *pb.CategoryFilter) (*pb.CategoryList, error) {
	filter := bson.M{"deleted_at": 0}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.Name != "" {
		filter["name"] = req.Name
	}
	if req.Type != "" {
		filter["type"] = req.Type
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
		"created_at": 1,
		"updated_at": 1,
	}

	options.SetProjection(projection)

	cursor, err := r.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var categories []*pb.CategoryGet
	for cursor.Next(context.Background()) {
		var category pb.CategoryGet
		var cat Category
		if err := cursor.Decode(&cat); err != nil {
			return nil, err
		}
		// Map MongoDB Object ID to proto Id field
		category.Id = cursor.Current.Lookup("_id").ObjectID().Hex()
		category.UserId = cat.UserId
		category.Name = cat.Name
		category.Type = cat.Type
		category.CreatedAt = cat.CreatedAt
		category.UpdatedAt = cat.UpdatedAt

		categories = append(categories, &category)
	}

	totalCount := len(categories)
	if err != nil {
		return nil, err
	}

	return &pb.CategoryList{
		Get:    categories,
		Total:  int32(totalCount),
		Limit:  req.Filter.Limit,
		Offset: req.Filter.Offset,
	}, nil
}

type Category struct {
	UserId    string `bson:"user_id"`
	Name      string `bson:"name"`
	Type      string `bson:"type"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}
