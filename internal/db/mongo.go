package db

package db

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo() {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI is not set in environment variables")
    }

    var client *mongo.Client
    var err error

    for i := 0; i < 10; i++ {
        client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
        if err == nil {
            break
        }
        log.Println("MongoDB not ready, retrying in 3 seconds...")
        time.Sleep(3 * time.Second)
    }

    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    log.Println("Connected to MongoDB")
    Client = client
}
