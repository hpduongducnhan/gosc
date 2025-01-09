package elkclient

import (
	"context"
	"strings"
)

func (esc *EslasticClient[MsgType]) Search(index string, query string, resultQueue chan MsgType, logParser func(InnerHit) MsgType) error {
	// Search with a scroll
	esClient := esc.client
	ctx := context.Background()

	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return err
	}
	parsedResp, err := parseElkResp(res)
	if err != nil {
		return err
	}

	// parse the hits and push to the result queue
	// log.Printf("got total hits: %d", len(parsedResp.Hits.Hits))
	for _, hit := range parsedResp.Hits.Hits {
		// log.Printf("hit: %v", hit)
		resultQueue <- logParser(hit)
	}
	return nil
}
