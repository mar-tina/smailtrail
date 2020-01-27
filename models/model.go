package models

type Message struct {
	Body struct {
		Size string `json:"size"`
		Data string `json:"data"`
	} `json:"body"`
	Headers []Part

	HistoryID string `json:"historyId"`
	Parts     []struct {
		Body     Body `json:"body"`
		Headers  []Part
		MimeType string `json:"mimeType"`
		Filename string `json:"filename"`
	}
}

type Body struct {
	Data string `json:"data"`
	Size int    `json:"size"`
}

type Part struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
