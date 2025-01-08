package mongoclient

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoClient(uri string, config *options.ClientOptions) *MongoDbClient {
	client := &MongoDbClient{Uri: uri}
	err := client.Connect(nil)
	if err != nil {
		log.Fatal(`connect to mongodb failed: %w`, err)
	}

	return client
}
