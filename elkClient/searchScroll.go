package elkclient

import (
	"context"
	"log"
	"strings"
	"time"
)

var scrollTime time.Duration = 2 * time.Minute

func (esc *EslasticClient[MsgType]) clearScroll(scrollID string) {
	esClient := esc.client
	_, err := esClient.ClearScroll(
		esClient.ClearScroll.WithScrollID(scrollID),
	)
	if err != nil {
		log.Printf("Error clearing scroll: %s", err)
	}
}

func (esc *EslasticClient[MsgType]) SearchScroll(index string, query string, resultQueue chan MsgType, logParser func(InnerHit) MsgType) error {
	// Search with a scroll
	esClient := esc.client
	var scrollID string
	ctx := context.Background()

	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(strings.NewReader(query)),
		esClient.Search.WithScroll(scrollTime),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return err
	}
	parsedResp, err := parseElkResp(res)
	if err != nil {
		log.Printf("Error parsing response: %s", err)
		return err
	}

	scrollID = parsedResp.ScrollID
	// parse the hits and push to the result queue
	// log.Printf("got total hits: %d", len(parsedResp.Hits.Hits))
	for _, hit := range parsedResp.Hits.Hits {
		resultQueue <- logParser(hit)
	}

	// Iterate over the hits
	for {
		res, err := esClient.Scroll(
			esClient.Scroll.WithContext(ctx),
			esClient.Scroll.WithScrollID(scrollID),
			esClient.Scroll.WithScroll(scrollTime),
		)
		if err != nil {
			log.Printf("scroll get error: %s", err.Error())
			break
		}

		parsedResp, err := parseElkResp(res)
		if err != nil {
			log.Printf("Error parsing response: %s", err)
			return err
		}

		// exit if no more hits
		if len(parsedResp.Hits.Hits) == 0 {
			break
		}

		scrollID = parsedResp.ScrollID
		// parse the hits and push to the result queue
		// log.Printf("got total hits: %d", len(parsedResp.Hits.Hits))
		for _, hit := range parsedResp.Hits.Hits {
			resultQueue <- logParser(hit)
		}
	}

	// clear the scroll
	go esc.clearScroll(scrollID)

	return nil
}
