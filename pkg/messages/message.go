package messages

// Representation of an incoming message payload
type Message struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Subject string `json:"subject" form:"subject"`
	Text    string `json:"text" form:"text"`
}
