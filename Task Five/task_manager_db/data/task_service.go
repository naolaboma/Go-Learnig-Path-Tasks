package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	Collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		Collection: collection,
	}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := s.Collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	if err = cursor.All(c, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task id format")
	}
	var task models.Task
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = s.Collection.FindOne(c, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) CreateTask(task models.Task) (*models.Task, error) {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.Collection.InsertOne(c, task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")

	}
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
			"updated_at":  updatedTask.UpdatedAt,
		},
	}
	result, err := s.Collection.UpdateOne(c, filter, update)

	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("task not found")

	}

	return s.GetTaskByID(id)
}
func (s *TaskService) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")

	}

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.Collection.DeleteOne(c, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")

	}
	return nil
}
