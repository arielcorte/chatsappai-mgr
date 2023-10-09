package main

// Account represents the payload structure for an account.
type Account struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Inbox represents the payload structure for an inbox.
type Inbox struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Contact represents the payload structure for a contact.
type Contact struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Avatar  string  `json:"avatar"`
	Type    string  `json:"type"`
	Account Account `json:"account"`
}

// User represents the payload structure for a user (agent/admin).
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

// Conversation represents the payload structure for a conversation.
type Conversation struct {
	AdditionalAttributes struct {
		Browser struct {
			DeviceName      string `json:"device_name"`
			BrowserName     string `json:"browser_name"`
			PlatformName    string `json:"platform_name"`
			BrowserVersion  string `json:"browser_version"`
			PlatformVersion string `json:"platform_version"`
		} `json:"browser"`
		Referer     string `json:"referer"`
		InitiatedAt struct {
			Timestamp string `json:"timestamp"`
		} `json:"initiated_at"`
	} `json:"additional_attributes"`
	CanReply     bool   `json:"can_reply"`
	Channel      string `json:"channel"`
	ID           int    `json:"id"`
	InboxID      int    `json:"inbox_id"`
	ContactInbox struct {
		ID           int    `json:"id"`
		ContactID    int    `json:"contact_id"`
		InboxID      int    `json:"inbox_id"`
		SourceID     string `json:"source_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		HMACVerified bool   `json:"hmac_verified"`
	} `json:"contact_inbox"`
	Meta struct {
		Sender   Contact `json:"sender"`
		Assignee User    `json:"assignee"`
	} `json:"meta"`
	Status            string `json:"status"`
	UnreadCount       int    `json:"unread_count"`
	AgentLastSeenAt   int64  `json:"agent_last_seen_at"`
	ContactLastSeenAt int64  `json:"contact_last_seen_at"`
	Timestamp         int64  `json:"timestamp"`
	AccountID         int    `json:"account_id"`
	AssigneeID        int    `json:"assignee_id"`
}

// Message represents the payload structure for a message.
type Message struct {
	ID           int          `json:"id"`
	Content      string       `json:"content"`
	MessageType  string       `json:"message_type"`
	ContentType  string       `json:"content_type"`
	Private      bool         `json:"private"`
	Account      Account      `json:"account"`
	Conversation Conversation `json:"conversation"`
}

type MessageID struct {
	ID int `json:"id"`
}

// WebhookEvent represents a generic webhook event.
type WebhookEvent struct {
	Event string `json:"event"`
}

type ConversationUpdatedEvent struct {
	ChangedAttributes []map[string]struct {
		CurrentValue  string `json:"current_value"`
		PreviousValue string `json:"previous_value"`
	} `json:"changed_attributes"`
	Conversation Conversation `json:"conversation"`
}

type WebwidgetTriggeredEvent struct {
	ID                  int          `json:"id"`
	Contact             Contact      `json:"contact"`
	Inbox               Inbox        `json:"inbox"`
	Account             Account      `json:"account"`
	CurrentConversation Conversation `json:"current_conversation"`
	SourceID            string       `json:"source_id"`
	EventInfo           struct {
		InitiatedAt struct {
			Timestamp string `json:"timestamp"`
		} `json:"initiated_at"`
		Referer         string `json:"referer"`
		WidgetLanguage  string `json:"widget_language"`
		BrowserLanguage string `json:"browser_language"`
		Browser         struct {
			BrowserName     string `json:"browser_name"`
			BrowserVersion  string `json:"browser_version"`
			DeviceName      string `json:"device_name"`
			PlatformName    string `json:"platform_name"`
			PlatformVersion string `json:"platform_version"`
		} `json:"browser"`
	} `json:"event_info"`
}

type AgentBot struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Desciption  string `json:"description"`
	AccountID   int    `json:"account_id"`
	OutgoingURL string `json:"outgoing_url"`
	AccessToken string `json:"access_token"`
}

// WebhookPayload represents the payload structure for a message_created webhook.
// {"account":{"id":1,"name":"ChatsappAI.com"},"additional_attributes":{},"content_attributes":{},"content_type":"text","content":"test","conversation":{"additional_attributes":{},"can_reply":true,"channel":"Channel::Whatsapp","contact_inbox":{"id":84,"contact_id":89,"inbox_id":10,"source_id":"5493513055609","created_at":"2023-09-28T06:23:46.462Z","updated_at":"2023-09-28T06:23:46.462Z","hmac_verified":false,"pubsub_token":"e3nzQ74DmyT7qSH6WgcpBLm5"},"id":12,"inbox_id":10,"messages":[{"id":1646,"content":"test","account_id":1,"inbox_id":10,"conversation_id":12,"message_type":0,"created_at":1696014860,"updated_at":"2023-09-29T19:14:20.162Z","private":false,"status":"sent","source_id":"wamid.HBgNNTQ5MzUxMzA1NTYwORUCABIYIEI1MTgwMjQwMTkzN0YyMzBGMTEzQUY4MkYyRjJFRTJBAA==","content_type":"text","content_attributes":{},"sender_type":"Contact","sender_id":89,"external_source_ids":{},"additional_attributes":{},"processed_message_content":"test","sentiment":{},"conversation":{"assignee_id":3,"unread_count":1,"last_activity_at":1696014860,"contact_inbox":{"source_id":"5493513055609"}},"sender":{"additional_attributes":{},"custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":"","type":"contact"}}],"labels":[],"meta":{"sender":{"additional_attributes":{},"custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":"","type":"contact"},"assignee":{"id":3,"name":"Ariel Corte","available_name":"Ariel (Tecnichal Staff)","avatar_url":"https://app.chatsappai.com/rails/active_storage/representations/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBDdz09IiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--3f3bd5ba433dd789761c51a0ac6288b5e35e8d8a/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdCem9MWm05eWJXRjBTU0lJYW5CbkJqb0dSVlE2RTNKbGMybDZaVjkwYjE5bWFXeHNXd2RwQWZvdyIsImV4cCI6bnVsbCwicHVyIjoidmFyaWF0aW9uIn19--f591e7a4b8661c5e78b1e49f0b4fffe1eafb2258/arielprofile.jpg","type":"user","availability_status":null,"thumbnail":"https://app.chatsappai.com/rails/active_storage/representations/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBDdz09IiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--3f3bd5ba433dd789761c51a0ac6288b5e35e8d8a/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdCem9MWm05eWJXRjBTU0lJYW5CbkJqb0dSVlE2RTNKbGMybDZaVjkwYjE5bWFXeHNXd2RwQWZvdyIsImV4cCI6bnVsbCwicHVyIjoidmFyaWF0aW9uIn19--f591e7a4b8661c5e78b1e49f0b4fffe1eafb2258/arielprofile.jpg"},"team":null,"hmac_verified":false},"status":"open","custom_attributes":{},"snoozed_until":null,"unread_count":1,"first_reply_created_at":"2023-09-28T08:26:30.978Z","priority":null,"waiting_since":1696013510,"agent_last_seen_at":1696014734,"contact_last_seen_at":0,"timestamp":1696014860,"created_at":1695882226},"created_at":"2023-09-29T19:14:20.162Z","id":1646,"inbox":{"id":10,"name":"test"},"message_type":"incoming","private":false,"sender":{"account":{"id":1,"name":"ChatsappAI.com"},"additional_attributes":{},"avatar":"","custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":""},"source_id":"wamid.HBgNNTQ5MzUxMzA1NTYwORUCABIYIEI1MTgwMjQwMTkzN0YyMzBGMTEzQUY4MkYyRjJFRTJBAA==","event":"message_created"}
