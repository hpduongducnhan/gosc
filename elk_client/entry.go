package elkclient

func NewElkCollector[MsgType any](addresses []string, username string, password string) (*EslasticClient[MsgType], error) {
	esClient := &EslasticClient[MsgType]{
		Addresses:    addresses,
		AuthUser:     username,
		AuthPassword: password,
	}
	logger.Info().Msg("Connecting to ELK")
	err := esClient.Connect(nil)
	if err != nil {
		return nil, err
	}
	return esClient, nil
}
