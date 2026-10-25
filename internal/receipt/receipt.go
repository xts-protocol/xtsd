package receipt

type Receipt struct {
	MessageID string `json:"message_id"`
	APIKey    string `json:"api_key"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Timestamp string `json:"timestamp"`
}
