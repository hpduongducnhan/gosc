package elkclient

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type EslasticClient[MsgType any] struct {
	Addresses    []string
	AuthUser     string
	AuthPassword string
	client       *elasticsearch.Client
}

func (esc *EslasticClient[MsgType]) getDefaultConfig() *elasticsearch.Config {
	return &elasticsearch.Config{
		Addresses:     esc.Addresses,
		Username:      esc.AuthUser,
		Password:      esc.AuthPassword,
		RetryOnStatus: []int{502, 503, 504},
		MaxRetries:    3,
	}
}

func (esc *EslasticClient[MsgType]) Connect(config *elasticsearch.Config) error {
	if len(esc.Addresses) == 0 {
		log.Fatalf("No Elasticsearch addresses provided")
	}
	if config == nil {
		config = esc.getDefaultConfig()
	}
	es, err := elasticsearch.NewClient(*config)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return err
	}
	esc.client = es
	_, err = esc.client.Ping()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	return nil
}
