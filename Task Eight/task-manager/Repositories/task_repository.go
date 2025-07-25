package repositories

import (
	"context"
	"errors"
	"task-manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{collection: collection}
}

func (r *TaskRepository) Create(task domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task.ID = "" // Let MongoDB generate the ID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	task.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return &task, nil
}

func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	if tasks == nil {
		return []domain.Task{}, nil // Return empty slice instead of nil
	}
	return tasks, nil
}

func (r *TaskRepository) GetByID(id string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	var task domain.Task
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(id string, task domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.DueDate,
			"status":      task.Status,
			"updated_at":  time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	return r.GetByID(id) // Return the updated document
}

func (r *TaskRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
