package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{
		Collection: collection,
	}
}

func (s *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *UserService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) Register(user models.User) (*models.User, error) {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var existingUser models.User
	err := s.Collection.FindOne(c, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("username already exists")
	} else if err != mongo.ErrNoDocuments {
		return nil, err
	}

	hashedPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	count, err := s.Collection.CountDocuments(c, bson.M{})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		user.Role = models.RoleAdmin
	} else {
		user.Role = models.RoleAdmin
	}
	user.ID = primitive.NewObjectID()
	_, err = s.Collection.InsertOne(c, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err := s.Collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !s.checkPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credintials")
	}

	return &user, nil
}

func (s *UserService) PromoteUser(username string, promoterID primitive.ObjectID) error {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//checking if promoter is addmin
	var promoter models.User
	err := s.Collection.FindOne(c, bson.M{"_id": promoterID}).Decode(&promoter)
	if err != nil {
		return err
	}
	if promoter.Role != models.RoleAdmin {
		return errors.New("only admins can promote users")
	}

	// Updating user role

	result, err := s.Collection.UpdateOne(
		c, bson.M{"username": username},
		bson.M{"$set": bson.M{"role": models.RoleAdmin}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")

	}
	return nil
}

func (s *UserService) GetUserByID(id primitive.ObjectID) (*models.User, error) {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := s.Collection.FindOne(c, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
