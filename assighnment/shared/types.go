package shared

// Message represents a single chat message
type Message struct {
	Sender  string
	Content string
}

// SendMessageArgs is the request for ChatServer.SendMessage
type SendMessageArgs struct {
	Sender  string
	Content string
}

// SendMessageReply is the response for ChatServer.SendMessage
type SendMessageReply struct {
	Success bool
	History []Message
}

// GetHistoryArgs is the request for ChatServer.GetHistory
type GetHistoryArgs struct{}

// GetHistoryReply is the response for ChatServer.GetHistory
type GetHistoryReply struct {
	History []Message
}

// ... more code ...