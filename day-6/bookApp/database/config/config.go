package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"bookApp/internal/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

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

func InitDB() *gorm.DB {
	configEnv := LoadEnv()

	loadConfig := map[string]string{
		"Username": configEnv.GetString("DB_USERNAME"),
		"Password": configEnv.GetString("DB_PASSWORD"),
		"Host":     configEnv.GetString("DB_HOST"),
		"Port":     configEnv.GetString("DB_PORT"),
		"DB":       configEnv.GetString("DB_NAME"),
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", loadConfig["Username"], loadConfig["Password"], loadConfig["Host"], loadConfig["Port"], loadConfig["DB"])

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err.Error())
	}
	MigrateDB()
	return db
}

func MigrateDB() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
}

func ConnectDB() *mongo.Client {
	configEnv := LoadEnv()
	client, err := mongo.NewClient(options.Client().ApplyURI(configEnv.GetString("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DBM *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	configEnv := LoadEnv()
	dbName := configEnv.GetString("MONGODB_NAME")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
