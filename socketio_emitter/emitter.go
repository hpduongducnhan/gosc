package socketioemitter

import (
	"context"

	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	EVENT        = 2
	BINARY_EVENT = 5
)

type EmitterOpts struct {
	// Host means hostname like localhost
	RedisUrl string

	// Key means redis subscribe key
	Key string

	// Name
	Name string
}

type Emitter struct {
	Redis  *redis.Client
	Config *EmitterOpts
}

func (e *Emitter) Close() error {
	return e.Redis.Close()
}

func (e *Emitter) Publish(message WSMessage) (*redis.IntCmd, error) {
	message.SetName(e.Config.Name)
	packed, err := message.Pack()
	if err != nil {
		log.Error().Err(err).Msg("Failed to pack message")
		return nil, err
	}
	ctx := context.Background()

	if len(message.Rooms) == 0 {
		channel := e.Config.Key + "#" + message.Namespace + "#"
		log.Info().Str("channel", channel).Msg("Publishing message")
		return e.Redis.Publish(ctx, channel, packed), nil
	} else {
		for _, room := range message.Rooms {
			channel := e.Config.Key + "#" + message.Namespace + "#" + room + "#"
			log.Info().Str("channel", channel).Msg("Publishing message")
			_, err := e.Redis.Publish(ctx, channel, packed).Result()
			if err != nil {
				log.Error().Err(err).Msg("Failed to publish message")
				return nil, err
			}
		}
		return nil, nil
	}

}
