package elkclient

import "log"

func NewElkCollector[MsgType any](addresses []string, username string, password string) (*EslasticClient[MsgType], error) {
	esClient := &EslasticClient[MsgType]{
		Addresses:    addresses,
		AuthUser:     username,
		AuthPassword: password,
	}
	log.Printf("Connecting to ELK: %v", esClient)
	err := esClient.Connect(nil)
	if err != nil {
		log.Printf("Error connecting to ELK: %s", err)
		return nil, err
	}
	return esClient, nil
}
