package persist

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

//ConnectToMongo creates a connection to mongoDB
func ConnectToMongo(uri string, username string, password string) {
	//Creates a mongo connection string using passed in variables
	connectionString := fmt.Sprintf(uri, username, password)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	collection = client.Database("wcm_alpha").Collection("leads")
}

//CreateLead adds a new lead to the db.
func CreateLead(data []byte) {
	var bdoc interface{}
	err := bson.UnmarshalExtJSON(data, false, &bdoc)
	if err != nil {
		log.Println("Error marshalling for storage.")
		return
	}
	insertResult, err := collection.InsertOne(context.TODO(), bdoc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
}
