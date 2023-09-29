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
	Event                string `json:"event"`
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
	Messages []Message `json:"messages"`
	Meta     struct {
		Sender   Contact `json:"sender"`
		Assignee User    `json:"assignee"`
	} `json:"meta"`
	Status            string `json:"status"`
	UnreadCount       int    `json:"unread_count"`
	AgentLastSeenAt   int64  `json:"agent_last_seen_at"`
	ContactLastSeenAt int64  `json:"contact_last_seen_at"`
	Timestamp         int64  `json:"timestamp"`
	AccountID         int    `json:"account_id"`
}

// Message represents the payload structure for a message.
type Message struct {
	Event             string      `json:"event"`
	ID                int         `json:"id"`
	Content           string      `json:"content"`
	MessageType       string      `json:"message_type"`
	CreatedAt         int64       `json:"created_at"`
	Private           bool        `json:"private"`
	SourceID          string      `json:"source_id"`
	ContentType       string      `json:"content_type"`
	ContentAttributes interface{} `json:"content_attributes"`
	Sender            struct {
		Type    string  `json:"type"`
		User    User    `json:"user,omitempty"`
		Contact Contact `json:"contact,omitempty"`
	} `json:"sender"`
	Account      Account      `json:"account"`
	Conversation Conversation `json:"conversation"`
	Inbox        Inbox        `json:"inbox"`
}

// WebhookEvent represents a generic webhook event.
type WebhookEvent interface {
}

type ConversationUpdatedEvent struct {
	Event             string `json:"event"`
	ChangedAttributes []map[string]struct {
		CurrentValue  string `json:"current_value"`
		PreviousValue string `json:"previous_value"`
	} `json:"changed_attributes"`
	Conversation Conversation `json:"conversation"`
}

type WebwidgetTriggeredEvent struct {
	Event               string       `json:"event"`
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
