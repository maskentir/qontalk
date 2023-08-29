package qontak

// SendMessageInteractions is a struct representing the parameters for sending message interactions.
type SendMessageInteractions struct {
	ReceiveMessageFromAgent    bool
	ReceiveMessageFromCustomer bool
	StatusMessage              bool
	URL                        string
}

// Button represents an interactive message button.
type Button struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// InteractiveRow represents a row in an interactive message section.
type InteractiveRow struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// InteractiveSection represents a section in an interactive message.
type InteractiveSection struct {
	Title string           `json:"title"`
	Rows  []InteractiveRow `json:"rows"`
}

// InteractiveLists represents a list of interactive sections with a common button.
type InteractiveLists struct {
	Button   string               `json:"button"`
	Sections []InteractiveSection `json:"sections"`
}

// InteractiveHeader represents the header of an interactive message.
type InteractiveHeader struct {
	Format   string `json:"format"`
	Text     string `json:"text"`
	Link     string `json:"link"`
	Filename string `json:"filename"`
}

// SendInteractiveMessage is a struct representing a message to be sent interactively.
type SendInteractiveMessage struct {
	RoomID      string          `json:"room_id"`
	Type        string          `json:"type"`
	Interactive InteractiveData `json:"interactive"`
}

// InteractiveData represents the data for an interactive message.
type InteractiveData struct {
	Header  *InteractiveHeader `json:"header,omitempty"`
	Body    string             `json:"body"`
	Buttons []Button           `json:"buttons"`
	Lists   *InteractiveLists  `json:"lists,omitempty"`
}

// WhatsAppMessage represents the parameters for sending a WhatsApp message.
type WhatsAppMessage struct {
	RoomID  string
	Message string
}

// ButtonMessage represents a button in a message.
type ButtonMessage struct {
	Index string `json:"index"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// KeyValue represents a key-value pair.
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// KeyValueText represents a key-value pair with text value.
type KeyValueText struct {
	Key       string `json:"key"`
	ValueText string `json:"value_text"`
	Value     string `json:"value"`
}

// DirectWhatsAppBroadcast is a builder for creating parameters for sending direct WhatsApp broadcast.
type DirectWhatsAppBroadcast struct {
	ToName               string            `json:"to_name"`
	ToNumber             string            `json:"to_number"`
	MessageTemplateID    string            `json:"message_template_id"`
	ChannelIntegrationID string            `json:"channel_integration_id"`
	Language             map[string]string `json:"language"`
	HeaderParams         []KeyValue        `json:"header"`
	BodyParams           []KeyValueText    `json:"body"`
	Buttons              []ButtonMessage   `json:"buttons"`
}
