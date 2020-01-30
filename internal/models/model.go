package models

type Message struct {
	Body      Body `json:"body"`
	Headers   []Part
	HistoryID string `json:"historyId"`
	Parts     []struct {
		Body     Body `json:"body"`
		Headers  []Part
		MimeType string `json:"mimeType"`
		Filename string `json:"filename"`
	}
}

type Subscription struct {
	Sender string `json:"sender"`
	Link   string `json:"link"`
}

type Sub struct {
	ID     string `json:"id"`
	Sender string `json:"sender"`
	Link   string `json:"link"`
}

type GmailMsg struct {
	Messages []struct {
		ID       string `json:"id"`
		ThreadID string `json:"threadId"`
	}
	NextPageToken      string `json:"nextPageToken"`
	ResultSizeEstimate uint   `json:"resultSizeEstimate"`
}

type Body struct {
	Data string `json:"data"`
	Size int    `json:"size"`
}

type Part struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
