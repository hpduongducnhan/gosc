package socketioemitter

import (
	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var socketIoEmitter *Emitter

const WsNormalEvent = 2
const WsBinaryEvent = 5

func NewEmitter(redisUrl string, key string) (*Emitter, error) {
	// set default key
	if key == "" {
		key = "socket.io"
	}

	eName, err := generateRandomName(6)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate random name")
		return nil, err
	}

	// save emitter options
	emitterOptions := &EmitterOpts{
		RedisUrl: redisUrl,
		Key:      key,
		Name:     eName,
	}

	// connect to redis
	redisOpts, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse redis url")
		return nil, err
	}
	redisClient := redis.NewClient(redisOpts)

	socketIoEmitter = &Emitter{
		Redis:  redisClient,
		Config: emitterOptions,
	}
	return socketIoEmitter, nil
}

func NewWsMessage(event string, data interface{}, namespace string) *WSMessage {
	return &WSMessage{
		Event:     event,
		Data:      data,
		Namespace: namespace,
		Type:      WsNormalEvent, // EVENT type
		Flags:     make(map[string]interface{}),
	}
}
