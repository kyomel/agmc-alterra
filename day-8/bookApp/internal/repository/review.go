package repository

import (
	"bookApp/internal/models"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibReview struct {
	MG *mongo.Client
}

type Config interface {
	Get(string) any
	GetString(string) string
}

type Env struct{}

func (e *Env) Get(env string) any {
	err := os.Getenv(env)

	return err
}

func (e *Env) GetString(env string) string {
	return e.Get(env).(string)
}

func LoadEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Env{}
}

type ReviewRepo interface {
	CreateReview(review models.Review) error
	GetReviews() (*[]models.Review, error)
	GetReview(id string) (*models.Review, error)
	DeleteReview(id string) error
	UpdateReview(review primitive.M, id string) error
}

func InitReview(MG *mongo.Client) ReviewRepo {
	return &LibReview{MG}
}

func (r *LibReview) CreateReview(review models.Review) error {
	configEnv := LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := r.MG.Database(configEnv.GetString("MONGODB_NAME")).Collection(configEnv.GetString("MONGODB_COLLECTION"))

	_, err := query.InsertOne(ctx, review)

	return err
}

func (r *LibReview) GetReviews() (*[]models.Review, error) {
	configEnv := LoadEnv()
	var reviews []models.Review

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	query := r.MG.Database(configEnv.GetString("MONGODB_NAME")).Collection(configEnv.GetString("MONGODB_COLLECTION"))

	res, err := query.Find(ctx, bson.M{})
	if err != nil {
		return &reviews, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var review models.Review
		if err = res.Decode(review); err != nil {
			return &reviews, err
		}

		reviews = append(reviews, review)
	}

	return &reviews, err
}

func (r *LibReview) GetReview(id string) (*models.Review, error) {
	configEnv := LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	query := r.MG.Database(configEnv.GetString("MONGODB_NAME")).Collection(configEnv.GetString("MONGODB_COLLECTION"))

	var review models.Review
	err := query.FindOne(ctx, bson.M{"id": objId}).Decode(&review)

	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *LibReview) DeleteReview(id string) error {
	configEnv := LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	query := r.MG.Database(configEnv.GetString("MONGODB_NAME")).Collection(configEnv.GetString("MONGODB_COLLECTION"))

	result, err := query.DeleteOne(ctx, bson.M{"id": objId})

	if result.DeletedCount < 1 {
		return err
	}

	return err
}

func (r *LibReview) UpdateReview(review primitive.M, id string) error {
	configEnv := LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	query := r.MG.Database(configEnv.GetString("MONGODB_NAME")).Collection(configEnv.GetString("MONGODB_COLLECTION"))

	result, err := query.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": review})
	if err != nil {
		return err
	}

	var updateReview models.Review
	if result.MatchedCount == 1 {
		err := query.FindOne(ctx, bson.M{"id": objId}).Decode(&updateReview)
		if err != nil {
			return err
		}
	}

	return err
}
