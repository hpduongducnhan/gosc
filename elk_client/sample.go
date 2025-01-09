package elkclient

type sampleLog struct {
	Message    string                 `mapstructure:"message" json:"message"`
	TicketID   string                 `json:"ticketID"`
	TicketCode string                 `json:"ticketCode"`
	GroupID    string                 `json:"groupID"`
	Raw        map[string]interface{} `json:"logSource"`

	Finished bool `json:"finished"`
}

func (sl *sampleLog) IsValid() bool {
	if sl.TicketID != "" && sl.TicketCode != "" && sl.GroupID != "" {
		return true
	}
	return false
}

func (sl *sampleLog) IsFinished() bool {
	return sl.Finished
}

func newEmptySampleLog() sampleLog {
	new := sampleLog{}
	new.Finished = true
	return new
}

func parseSearchScrollResponse(hit InnerHit) sampleLog {
	newLog := newEmptySampleLog()
	newLog.Raw = hit.Source
	newLog.Message = hit.Source["message"].(string)
	newLog.TicketID = hit.Source["ticketID"].(string)
	newLog.TicketCode = hit.Source["ticketCode"].(string)
	newLog.GroupID = hit.Source["groupID"].(string)
	return newLog
}

func RunExample() {
	index := "filebeat-em-8.8.2*"
	address := []string{"http://localhost:9200"}
	username := "elastic"
	password := "changeme"
	resultQueue := make(chan sampleLog)
	query := ""
	esClient, err := NewElkCollector[sampleLog](address, username, password)
	if err != nil {
		return
	}
	esClient.SearchScroll(index, query, resultQueue, parseSearchScrollResponse)
	esClient.Search(index, query, resultQueue, parseSearchScrollResponse)
}
