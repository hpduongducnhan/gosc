package redisclient

import (
	"context"

	"github.com/hpduongducnhan/gosc/utils"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type WebhookRequest struct {
	ReqURL    string  `json:"req_url"`
	ReqMethod string  `json:"req_method"`
	ReqBody   ReqBody `json:"req_body"`
}

type ReqBody struct {
	TicketCode  string     `json:"TicketCode"`
	Data        TicketData `json:"Data"`
	API         string     `json:"API"`
	Content     string     `json:"Content"`
	CreatedDate string     `json:"CreatedDate"`
}

type TicketData struct {
	StepName   string `json:"StepName"`
	StepKey    string `json:"StepKey"`
	ActionName string `json:"ActionName"`
	ActionKey  string `json:"ActionKey"`
	TicketID   int    `json:"TicketID"`
	UpdatedBy  int    `json:"UpdatedBy"`
}

func RunSubsribeTicketEvents() {
	ctx := context.Background()
	rclient, err := ConnectRedis(ctx, "redis://:password@localhost:6379/0?protocol=3")
	if err != nil {
		panic(err)
	}

	SubscribeWithReconnect(context.Background(), rclient, "ticket:events", func(msg *redis.Message) {
		payload := msg.Payload
		whReq, err := utils.JsonString2Struct[WebhookRequest](payload)
		if err != nil {
			log.Error().Err(err).Str("payload", payload).Msg("Error parsing payload")
		} else {
			log.Info().Interface("webhookRequest", whReq).Msg("Webhook request received")
		}
	}, 0, 0)
}
