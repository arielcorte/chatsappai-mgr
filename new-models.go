package main

// -- WEBHOOKS

// -- conversation_status_changed

// {"additional_attributes":{},"can_reply":false,"channel":"Channel::Whatsapp","contact_inbox":{"id":547,"contact_id":89,"inbox_id":17,"source_id":"5493513055609","created_at":"2023-10-04T21:32:48.070Z","updated_at":"2023-10-04T21:32:48.070Z","hmac_verified":false,"pubsub_token":"yA27U2GouFMCXBBTob12617o"},"id":24,"inbox_id":17,"messages":[{"id":15001,"content":"Tu conversación ha sido asignada a un agente.\n\nEn la brevedad se contactarán contigo para ayudarte.\n\nMuchas Gracias 😊","account_id":1,"inbox_id":17,"conversation_id":24,"message_type":1,"created_at":1696896263,"updated_at":"2023-10-10T00:04:25.515Z","private":false,"status":"delivered","source_id":"wamid.HBgNNTQ5MzUxMzA1NTYwORUCABEYEjA0QkE0NTM1NTRGMzlGNUQyOAA=","content_type":"text","content_attributes":{},"sender_type":"AgentBot","sender_id":1,"external_source_ids":{},"additional_attributes":{},"processed_message_content":"Tu conversación ha sido asignada a un agente.\n\nEn la brevedad se contactarán contigo para ayudarte.\n\nMuchas Gracias 😊","sentiment":{},"conversation":{"assignee_id":3,"unread_count":0,"last_activity_at":1696896269,"contact_inbox":{"source_id":"5493513055609"}},"sender":{"id":1,"name":"ChatsappAI-Bot","avatar_url":"","type":"agent_bot"}}],"labels":[],"meta":{"sender":{"additional_attributes":{},"custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":"","type":"contact"},"assignee":{"id":3,"name":"Ariel Corte","available_name":"Ariel (Tecnichal Staff)","avatar_url":"https://app.chatsappai.com/rails/active_storage/representations/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBDdz09IiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--3f3bd5ba433dd789761c51a0ac6288b5e35e8d8a/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdCem9MWm05eWJXRjBTU0lJYW5CbkJqb0dSVlE2RTNKbGMybDZaVjkwYjE5bWFXeHNXd2RwQWZvdyIsImV4cCI6bnVsbCwicHVyIjoidmFyaWF0aW9uIn19--f591e7a4b8661c5e78b1e49f0b4fffe1eafb2258/arielprofile.jpg","type":"user","availability_status":null,"thumbnail":"https://app.chatsappai.com/rails/active_storage/representations/redirect/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaHBDdz09IiwiZXhwIjpudWxsLCJwdXIiOiJibG9iX2lkIn19--3f3bd5ba433dd789761c51a0ac6288b5e35e8d8a/eyJfcmFpbHMiOnsibWVzc2FnZSI6IkJBaDdCem9MWm05eWJXRjBTU0lJYW5CbkJqb0dSVlE2RTNKbGMybDZaVjkwYjE5bWFXeHNXd2RwQWZvdyIsImV4cCI6bnVsbCwicHVyIjoidmFyaWF0aW9uIn19--f591e7a4b8661c5e78b1e49f0b4fffe1eafb2258/arielprofile.jpg"},"team":null,"hmac_verified":false},"status":"open","custom_attributes":{},"snoozed_until":null,"unread_count":0,"first_reply_created_at":null,"priority":null,"waiting_since":1696896133,"agent_last_seen_at":1696997787,"contact_last_seen_at":0,"timestamp":1696896269,"created_at":1696896133,"event":"conversation_status_changed","changed_attributes":null}

type ConversationStatusChangedEvent struct {
	AdditionalAttributes struct{}         `json:"additional_attributes"`
	CanReply             bool             `json:"can_reply"`
	Channel              string           `json:"channel"`
	ContactInbox         CSC_ContactInbox `json:"contact_inbox"`
	Id                   int              `json:"id"`
	InboxId              int              `json:"inbox_id"`
	Messages             []CSC_Message    `json:"messages"`
	Labels               []struct{}       `json:"labels"`
	Meta                 CSC_Meta         `json:"meta"`
	Status               string           `json:"status"`
	CustomAttributes     struct{}         `json:"custom_attributes"`
	SnoozedUntil         interface{}      `json:"snoozed_until"`
	UnreadCount          int              `json:"unread_count"`
	FirstReplyCreatedAt  interface{}      `json:"first_reply_created_at"`
	Priority             interface{}      `json:"priority"`
	WaitingSince         int              `json:"waiting_since"`
	AgentLastSeenAt      int              `json:"agent_last_seen_at"`
	ContactLastSeenAt    int              `json:"contact_last_seen_at"`
	Timestamp            int              `json:"timestamp"`
	CreatedAt            int              `json:"created_at"`
	Event                string           `json:"event"`
	ChangedAttributes    interface{}      `json:"changed_attributes"`
}

type CSC_ContactInbox struct {
	Id           int    `json:"id"`
	ContactId    int    `json:"contact_id"`
	InboxId      int    `json:"inbox_id"`
	SourceId     string `json:"source_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	HmacVerified bool   `json:"hmac_verified"`
	PubsubToken  string `json:"pubsub_token"`
}

type CSC_Message struct {
	Id                      int                      `json:"id"`
	Content                 string                   `json:"content"`
	AccountId               int                      `json:"account_id"`
	InboxId                 int                      `json:"inbox_id"`
	ConversationId          int                      `json:"conversation_id"`
	MessageType             int                      `json:"message_type"`
	CreatedAt               int                      `json:"created_at"`
	UpdatedAt               string                   `json:"updated_at"`
	Private                 bool                     `json:"private"`
	Status                  string                   `json:"status"`
	SourceId                string                   `json:"source_id"`
	ContentType             string                   `json:"content_type"`
	ContentAttributes       struct{}                 `json:"content_attributes"`
	SenderType              string                   `json:"sender_type"`
	SenderId                int                      `json:"sender_id"`
	ExternalSourceIds       struct{}                 `json:"external_source_ids"`
	AdditionalAttributes    struct{}                 `json:"additional_attributes"`
	ProcessedMessageContent string                   `json:"processed_message_content"`
	Sentiment               struct{}                 `json:"sentiment"`
	Conversation            CSC_Message_Conversation `json:"conversation"`
	Sender                  CSC_Message_Sender       `json:"sender"`
}

type CSC_Message_Conversation struct {
	AssigneeId     int              `json:"assignee_id"`
	UnreadCount    int              `json:"unread_count"`
	LastActivityAt int              `json:"last_activity_at"`
	ContactInbox   CSC_ContactInbox `json:"contact_inbox"`
}

type CSC_Message_Sender struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
}

type CSC_Meta struct {
	Sender       CSC_Meta_Sender   `json:"sender"`
	Assignee     CSC_Meta_Assignee `json:"assignee"`
	Team         interface{}       `json:"team"`
	HmacVerified bool              `json:"hmac_verified"`
}

type CSC_Meta_Sender struct {
	AdditionalAttributes struct{} `json:"additional_attributes"`
	CustomAttributes     struct{} `json:"custom_attributes"`
	Email                string   `json:"email"`
	Id                   int      `json:"id"`
	Identifier           string   `json:"identifier"`
	Name                 string   `json:"name"`
	PhoneNumber          string   `json:"phone_number"`
	Thumbnail            string   `json:"thumbnail"`
	Type                 string   `json:"type"`
}

type CSC_Meta_Assignee struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	AvailableName      string `json:"available_name"`
	AvatarUrl          string `json:"avatar_url"`
	Type               string `json:"type"`
	AvailabilityStatus string `json:"availability_status"`
	Thumbnail          string `json:"thumbnail"`
}

// -- message_created

// {"account":{"id":1,"name":"ChatsappAI.com"},"additional_attributes":{},"content_attributes":{},"content_type":"text","content":"hola","conversation":{"additional_attributes":{},"can_reply":true,"channel":"Channel::Whatsapp","contact_inbox":{"id":547,"contact_id":89,"inbox_id":17,"source_id":"5493513055609","created_at":"2023-10-04T21:32:48.070Z","updated_at":"2023-10-04T21:32:48.070Z","hmac_verified":false,"pubsub_token":"yA27U2GouFMCXBBTob12617o"},"id":25,"inbox_id":17,"messages":[{"id":17547,"content":"hola","account_id":1,"inbox_id":17,"conversation_id":25,"message_type":0,"created_at":1696996971,"updated_at":"2023-10-11T04:02:51.785Z","private":false,"status":"sent","source_id":"wamid.HBgNNTQ5MzUxMzA1NTYwORUCABIYFjNFQjA3NDk2RDIxN0E2OENCNTA1NjkA","content_type":"text","content_attributes":{},"sender_type":"Contact","sender_id":89,"external_source_ids":{},"additional_attributes":{},"processed_message_content":"hola","sentiment":{"label":"positive","score":0.8144278719272577,"value":1},"conversation":{"assignee_id":null,"unread_count":10,"last_activity_at":1696996971,"contact_inbox":{"source_id":"5493513055609"}},"sender":{"additional_attributes":{},"custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":"","type":"contact"}}],"labels":[],"meta":{"sender":{"additional_attributes":{},"custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":"","type":"contact"},"assignee":null,"team":null,"hmac_verified":false},"status":"pending","custom_attributes":{},"snoozed_until":null,"unread_count":10,"first_reply_created_at":null,"priority":null,"waiting_since":1696896274,"agent_last_seen_at":0,"contact_last_seen_at":0,"timestamp":1696996971,"created_at":1696896274},"created_at":"2023-10-11T04:02:51.319Z","id":17547,"inbox":{"id":17,"name":"ChatsappAI"},"message_type":"incoming","private":false,"sender":{"account":{"id":1,"name":"ChatsappAI.com"},"additional_attributes":{},"avatar":"","custom_attributes":{},"email":null,"id":89,"identifier":null,"name":"Ariel Corte","phone_number":"+5493513055609","thumbnail":""},"source_id":"wamid.HBgNNTQ5MzUxMzA1NTYwORUCABIYFjNFQjA3NDk2RDIxN0E2OENCNTA1NjkA","event":"message_created"}

type MessageCreatedEvent struct {
	Account              MC_Account      `json:"account"`
	AdditionalAttributes struct{}        `json:"additional_attributes"`
	ContentAttributes    struct{}        `json:"content_attributes"`
	ContentType          string          `json:"content_type"`
	Content              string          `json:"content"`
	Conversation         MC_Conversation `json:"conversation"`
	CreatedAt            string          `json:"created_at"`
	Id                   int             `json:"id"`
	Inbox                MC_Inbox        `json:"inbox"`
	MessageType          string          `json:"message_type"`
	Private              bool            `json:"private"`
	Sender               MC_Sender       `json:"sender"`
	SourceId             string          `json:"source_id"`
	Event                string          `json:"event"`
}

type MC_Account struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MC_Conversation struct {
	AdditionalAttributes struct{}                  `json:"additional_attributes"`
	CanReply             bool                      `json:"can_reply"`
	Channel              string                    `json:"channel"`
	ContactInbox         MC_ContactInbox           `json:"contact_inbox"`
	Id                   int                       `json:"id"`
	InboxId              int                       `json:"inbox_id"`
	Messages             []MC_Conversation_Message `json:"messages"`
	Labels               []struct{}                `json:"labels"`
	Meta                 MC_Meta                   `json:"meta"`
	Status               string                    `json:"status"`
	CustomAttributes     struct{}                  `json:"custom_attributes"`
	SnoozedUntil         interface{}               `json:"snoozed_until"`
	UnreadCount          int                       `json:"unread_count"`
	FirstReplyCreatedAt  interface{}               `json:"first_reply_created_at"`
	Priority             interface{}               `json:"priority"`
	WaitingSince         int                       `json:"waiting_since"`
	AgentLastSeenAt      int                       `json:"agent_last_seen_at"`
	ContactLastSeenAt    int                       `json:"contact_last_seen_at"`
	Timestamp            int                       `json:"timestamp"`
	CreatedAt            int                       `json:"created_at"`
}

type MC_ContactInbox struct {
	Id           int    `json:"id"`
	ContactId    int    `json:"contact_id"`
	InboxId      int    `json:"inbox_id"`
	SourceId     string `json:"source_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	HmacVerified bool   `json:"hmac_verified"`
	PubsubToken  string `json:"pubsub_token"`
}

type MC_Conversation_Message struct {
	Id                      int                                  `json:"id"`
	Content                 string                               `json:"content"`
	AccountId               int                                  `json:"account_id"`
	InboxId                 int                                  `json:"inbox_id"`
	ConversationId          int                                  `json:"conversation_id"`
	MessageType             int                                  `json:"message_type"`
	CreatedAt               int                                  `json:"created_at"`
	UpdatedAt               string                               `json:"updated_at"`
	Private                 bool                                 `json:"private"`
	Status                  string                               `json:"status"`
	SourceId                string                               `json:"source_id"`
	ContentType             string                               `json:"content_type"`
	ContentAttributes       struct{}                             `json:"content_attributes"`
	SenderType              string                               `json:"sender_type"`
	SenderId                int                                  `json:"sender_id"`
	ExternalSourceIds       struct{}                             `json:"external_source_ids"`
	AdditionalAttributes    struct{}                             `json:"additional_attributes"`
	ProcessedMessageContent string                               `json:"processed_message_content"`
	Sentiment               MC_Conversation_Message_Sentiment    `json:"sentiment"`
	Conversation            MC_Conversation_Message_Conversation `json:"conversation"`
	Sender                  MC_Conversation_Message_Sender       `json:"sender"`
}

type MC_Conversation_Message_Sentiment struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
	Value int     `json:"value"`
}

type MC_Conversation_Message_Conversation struct {
	AssigneeId     interface{}     `json:"assignee_id"`
	UnreadCount    int             `json:"unread_count"`
	LastActivityAt int             `json:"last_activity_at"`
	ContactInbox   MC_ContactInbox `json:"contact_inbox"`
}

type MC_Conversation_Message_Sender struct {
	AdditionalAttributes struct{} `json:"additional_attributes"`
	CustomAttributes     struct{} `json:"custom_attributes"`
	Email                string   `json:"email"`
	Id                   int      `json:"id"`
	Identifier           string   `json:"identifier"`
	Name                 string   `json:"name"`
	PhoneNumber          string   `json:"phone_number"`
	Thumbnail            string   `json:"thumbnail"`
	Type                 string   `json:"type"`
}

type MC_Meta struct {
	Sender       MC_Meta_Sender `json:"sender"`
	Assignee     interface{}    `json:"assignee"`
	Team         interface{}    `json:"team"`
	HmacVerified bool           `json:"hmac_verified"`
}

type MC_Meta_Sender struct {
	AdditionalAttributes struct{} `json:"additional_attributes"`
	CustomAttributes     struct{} `json:"custom_attributes"`
	Email                string   `json:"email"`
	Id                   int      `json:"id"`
	Identifier           string   `json:"identifier"`
	Name                 string   `json:"name"`
	PhoneNumber          string   `json:"phone_number"`
	Thumbnail            string   `json:"thumbnail"`
	Type                 string   `json:"type"`
}

type MC_Inbox struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MC_Sender struct {
	Account              MC_Account  `json:"account"`
	AdditionalAttributes struct{}    `json:"additional_attributes"`
	Avatar               string      `json:"avatar"`
	CustomAttributes     struct{}    `json:"custom_attributes"`
	Email                interface{} `json:"email"`
	Id                   int         `json:"id"`
	Identifier           interface{} `json:"identifier"`
	Name                 string      `json:"name"`
	PhoneNumber          interface{} `json:"phone_number"`
	Thumbnail            string      `json:"thumbnail"`
}

// -- API

/*
id
number
content
string
The text content of the message

content_type
string
Enum: "text" "input_select" "cards" "form"
The type of the template message

content_attributes
object
The content attributes for each content_type

message_type
string
Enum: "incoming" "outgoing" "activity" "template"
The type of the message

created_at
integer
The time at which message was created

private
boolean
The flags which shows whether the message is private or not

attachment
object
The file object attached to the image

sender
object
User/Agent/AgentBot object

conversation_id
number
ID of the conversation
*/

type Message struct {
	ID                int         `json:"id"`
	Content           string      `json:"content"`
	ContentType       string      `json:"content_type"`
	ContentAttributes interface{} `json:"content_attributes"`
	MessageType       int         `json:"message_type"`
	CreatedAt         int         `json:"created_at"`
	Private           bool        `json:"private"`
	Attachment        interface{} `json:"attachment"`
	Sender            Sender      `json:"sender"`
	ConversationID    int         `json:"conversation_id"`
}

/*
account_id
integer
The ID of the user

user_id
integer
The ID of the user

role
string
whether user is an administrator or agent
*/

type User struct {
	AccountID int    `json:"account_id"`
	UserID    int    `json:"user_id"`
	Role      string `json:"role"`
}

/*
id
integer
uid
string
name
string
available_name
string
display_name
string
email
string
account_id
integer
role
string
Enum: "agent" "administrator"
confirmed
boolean
availability_status
string
Enum: "available" "busy" "offline"
The availability status of the agent computed by Chatwoot.

auto_offline
boolean
Whether the availability status of agent is configured to go offline automatically when away.

custom_attributes
object
Available for users who are created through platform APIs and has custom attributes associated.
*/

type Agent struct {
	ID                 int         `json:"id"`
	UID                string      `json:"uid"`
	Name               string      `json:"name"`
	AvailableName      string      `json:"available_name"`
	DisplayName        string      `json:"display_name"`
	Email              string      `json:"email"`
	AccountID          int         `json:"account_id"`
	Role               string      `json:"role"`
	Confirmed          bool        `json:"confirmed"`
	AvailabilityStatus string      `json:"availability_status"`
	AutoOffline        bool        `json:"auto_offline"`
	CustomAttributes   interface{} `json:"custom_attributes"`
}

/*
id
number
ID of the agent bot

name
string
The name of the agent bot

description
string
The description about the agent bot

account_id
number
Account ID if it's an account specific bot

outgoing_url
string
The webhook URL for the bot
*/

type AgentBot struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccountID   int    `json:"account_id"`
	OutgoingURL string `json:"outgoing_url"`
}

type Sender struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Type      string `json:"type"`
}

/*
data
object
meta
object
mine_count
number
unassigned_count
number
assigned_count
number
all_count
number
payload
Array of objects
array of conversations

Array
id
number
ID of the conversation

messages
Array of objects (message)
Array
content
string
The text content of the message

content_type
string
Enum: "text" "input_select" "cards" "form"
The type of the template message

content_attributes
object
The content attributes for each content_type

message_type
string
Enum: "incoming" "outgoing" "activity" "template"
The type of the message

created_at
integer
The time at which message was created

private
boolean
The flags which shows whether the message is private or not

attachment
object
The file object attached to the image

sender
object
User/Agent/AgentBot object

conversation_id
number
ID of the conversation

account_id
number
Account Id

inbox_id
number
ID of the inbox

status
string
Enum: "open" "resolved" "pending"
The status of the conversation

timestamp
string
The time at which conversation was created

contact_last_seen_at
string
agent_last_seen_at
string
unread_count
number
The number of unread messages

additional_attributes
object
The object containing additional attributes related to the conversation

custom_attributes
object
The object to save custom attributes for conversation, accepts custom attributes key and value

meta
object
sender
object
id
number
ID fo the sender

name
string
The name of the sender

thumbnail
string
Avatar URL of the contact

channel
string
Channel Type

assignee
object (user)
id
number
uid
string
name
string
available_name
string
display_name
string
email
string
account_id
number
role
string
Enum: "agent" "administrator"
confirmed
boolean
custom_attributes
object
Available for users who are created through platform APIs and has custom attributes associated.

accounts
Array of objects (account)
Array
id
number
Account ID

name
string
Name of the account

role
string
Enum: "administrator" "agent"
The user role in the account


*/

type Conversation struct {
	ID                   int         `json:"id"`
	Messages             []Message   `json:"messages"`
	AccountID            int         `json:"account_id"`
	InboxID              int         `json:"inbox_id"`
	Status               string      `json:"status"`
	Timestamp            int         `json:"timestamp"`
	ContactLastSeenAt    int         `json:"contact_last_seen_at"`
	AgentLastSeenAt      int         `json:"agent_last_seen_at"`
	UnreadCount          int         `json:"unread_count"`
	AdditionalAttributes interface{} `json:"additional_attributes"`
	CustomAttributes     interface{} `json:"custom_attributes"`
	Meta                 struct {
		Sender   Sender `json:"sender"`
		Assignee User   `json:"assignee"`
	} `json:"meta"`
}

/*
id
number
ID of the inbox

name
string
The name of the inbox

website_url
string
Website URL

channel_type
string
The type of the inbox

avatar_url
string
The avatar image of the inbox

widget_color
string
Widget Color used for customization of the widget

website_token
string
Website Token

enable_auto_assignment
boolean
The flag which shows whether Auto Assignment is enabled or not

web_widget_script
string
Script used to load the website widget

welcome_title
string
Welcome title to be displayed on the widget

welcome_tagline
string
Welcome tagline to be displayed on the widget

greeting_enabled
boolean
The flag which shows whether greeting is enabled

greeting_message
string
A greeting message when the user starts the conversation


*/

type Inbox struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	WebsiteURL           string `json:"website_url"`
	ChannelType          string `json:"channel_type"`
	AvatarURL            string `json:"avatar_url"`
	WidgetColor          string `json:"widget_color"`
	WebsiteToken         string `json:"website_token"`
	EnableAutoAssignment bool   `json:"enable_auto_assignment"`
	WebWidgetScript      string `json:"web_widget_script"`
	WelcomeTitle         string `json:"welcome_title"`
	WelcomeTagline       string `json:"welcome_tagline"`
	GreetingEnabled      bool   `json:"greeting_enabled"`
	GreetingMessage      string `json:"greeting_message"`
}

/*
contact
object
email
string
Email address of the contact

name
string
The name of the contact

phone_number
string
Phone number of the contact

thumbnail
string
Avatar URL of the contact

additional_attributes
object
The object containing additional attributes related to the contact

custom_attributes
object
The object to save custom attributes for contact, accepts custom attributes key and value

contact_inboxes
Array of objects (contact_inboxes)
Array
source_id
string
Contact Inbox Source Id

inbox
object (inbox)
id
number
ID of the inbox

name
string
The name of the inbox

website_url
string
Website URL

channel_type
string
The type of the inbox

avatar_url
string
The avatar image of the inbox

widget_color
string
Widget Color used for customization of the widget

website_token
string
Website Token

enable_auto_assignment
boolean
The flag which shows whether Auto Assignment is enabled or not

web_widget_script
string
Script used to load the website widget

welcome_title
string
Welcome title to be displayed on the widget

welcome_tagline
string
Welcome tagline to be displayed on the widget

greeting_enabled
boolean
The flag which shows whether greeting is enabled

greeting_message
string
A greeting message when the user starts the conversation


*/

type Contact struct {
	Email                string      `json:"email"`
	Name                 string      `json:"name"`
	PhoneNumber          string      `json:"phone_number"`
	Thumbnail            string      `json:"thumbnail"`
	AdditionalAttributes interface{} `json:"additional_attributes"`
	CustomAttributes     interface{} `json:"custom_attributes"`
	ContactInboxes       []struct {
		SourceID string `json:"source_id"`
		Inbox    Inbox  `json:"inbox"`
	} `json:"contact_inboxes"`
}

/*
id
integer
ID of the canned response

content
string
Message content for canned response

short_code
string
Short Code for quick access of the canned response

account_id
integer
Account Id
*/

type CannedResponse struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	ShortCode string `json:"short_code"`
	AccountID int    `json:"account_id"`
}
