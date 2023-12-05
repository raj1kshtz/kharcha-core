package db

import (
	"bytes"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	onceMongo       sync.Once
	mongoDBClient   *mongo.Client
	mongoDBDatabase *mongo.Database
	ConnProps       *ConnectionProperties
)

func buildDBURI(props *ConnectionProperties) string {
	var buffer bytes.Buffer
	buffer.WriteString("mongodb://")
	buffer.WriteString(props.Host)
	buffer.WriteString(":")
	buffer.WriteString(props.Port)
	buffer.WriteString("/")
	buffer.WriteString(props.Database) // Include the database name in the connection string
	buffer.WriteString("?maxIdleTimeMS=")
	buffer.WriteString(strconv.Itoa(props.MaxConnLifeTime * 1000)) // Convert seconds to milliseconds

	log.Println("Database connection URI:", buffer.String())

	return buffer.String()
}

// GetMongoSession gets the MongoDB session.
func GetMongoSession(props *ConnectionProperties) (*mongo.Client, *mongo.Database, error) {
	onceMongo.Do(func() {
		initializeMongoClient(props)
	})

	return mongoDBClient, mongoDBDatabase, nil
}

// initializeMongoClient initializes the MongoDB client.
func initializeMongoClient(props *ConnectionProperties) {
	uri := buildDBURI(props)

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetMaxPoolSize(uint64(props.MaxPoolSize))
	clientOptions.SetMinPoolSize(uint64(props.MinPoolSize))
	clientOptions.SetMaxConnIdleTime(time.Duration(props.MaxConnIdleTime) * time.Second)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	mongoDBClient = client
	mongoDBDatabase = client.Database(props.Database)
}
