package repositories

import (
	"context"
	"errors"
	domain "task-manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection      *mongo.Collection
	passwordService domain.IPasswordService
}

func NewUserRepository(collection *mongo.Collection, passwordService domain.IPasswordService) *UserRepository {
	return &UserRepository{
		collection:      collection,
		passwordService: passwordService,
	}
}

func (r *UserRepository) Create(user domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	exists, err := r.Exists(user.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrDuplicateEntry
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = domain.RoleUser
	}

	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	var user domain.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Promote(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"role": domain.RoleAdmin,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"username": username}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *UserRepository) Exists(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
