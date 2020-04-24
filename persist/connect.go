package persist

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bdkiran/traject/utils"
)

var collection *mongo.Collection

//InitializeMongo creates a connection to mongoDB
func InitializeMongo(uri string, username string, password string) error {
	//Creates a mongo connection string using passed in variables
	connectionString := fmt.Sprintf(uri, username, password)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		utils.DefaultLogger.Error.Println("Unable to make a connection to mongoDB")
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		utils.DefaultLogger.Error.Println("Error occureded when sending ping to mongoDB")
		return err
	}

	collection = client.Database("wcm_alpha").Collection("leads")
	utils.DefaultLogger.Info.Println("Connected to MongoDB!")
	return nil
}

//CreateLead adds a new lead to the db, takes in a byte array(json object) and stores it as a document.
func CreateLead(data []byte) error {
	var bdoc interface{}
	err := bson.UnmarshalExtJSON(data, false, &bdoc)
	if err != nil {
		utils.DefaultLogger.Error.Println("Error marshalling passed in data before performing insert.")
		return err
	}
	insertResult, err := collection.InsertOne(context.TODO(), bdoc)
	if err != nil {
		utils.DefaultLogger.Error.Println("An error occured when inserting data in mongoDB")
		return err
	}
	utils.DefaultLogger.Info.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}
