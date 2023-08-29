package qontak

// SendMessageInteractionsBuilder is a builder for creating message interactions.
type SendMessageInteractionsBuilder struct {
	ReceiveMessageFromAgent    bool
	ReceiveMessageFromCustomer bool
	StatusMessage              bool
	URL                        string
}

// NewSendMessageInteractionsBuilder creates a new instance of SendMessageInteractionsBuilder.
func NewSendMessageInteractionsBuilder() *SendMessageInteractionsBuilder {
	return &SendMessageInteractionsBuilder{}
}

// WithReceiveMessageFromAgent sets the ReceiveMessageFromAgent flag.
func (b *SendMessageInteractionsBuilder) WithReceiveMessageFromAgent(
	receive bool,
) *SendMessageInteractionsBuilder {
	b.ReceiveMessageFromAgent = receive
	return b
}

// WithReceiveMessageFromCustomer sets the ReceiveMessageFromCustomer flag.
func (b *SendMessageInteractionsBuilder) WithReceiveMessageFromCustomer(
	receive bool,
) *SendMessageInteractionsBuilder {
	b.ReceiveMessageFromCustomer = receive
	return b
}

// WithStatusMessage sets the StatusMessage flag.
func (b *SendMessageInteractionsBuilder) WithStatusMessage(
	status bool,
) *SendMessageInteractionsBuilder {
	b.StatusMessage = status
	return b
}

// WithURL sets the URL for message interactions.
func (b *SendMessageInteractionsBuilder) WithURL(url string) *SendMessageInteractionsBuilder {
	b.URL = url
	return b
}

// Build builds the SendMessageInteractions using the configuration from the builder.
// Example:
//
//	builder := NewSendMessageInteractionsBuilder().
//	    WithReceiveMessageFromAgent(true).
//	    WithReceiveMessageFromCustomer(false).
//	    WithStatusMessage(true).
//	    WithURL("https://example.com").
//	interactions := builder.Build()
func (b *SendMessageInteractionsBuilder) Build() SendMessageInteractions {
	return SendMessageInteractions{
		ReceiveMessageFromAgent:    b.ReceiveMessageFromAgent,
		ReceiveMessageFromCustomer: b.ReceiveMessageFromCustomer,
		StatusMessage:              b.StatusMessage,
		URL:                        b.URL,
	}
}

// SendInteractiveMessageBuilder is a builder for creating interactive messages.
type SendInteractiveMessageBuilder struct {
	RoomID          string
	InteractiveData InteractiveData
}

// NewSendInteractiveMessageBuilder creates a new instance of SendInteractiveMessageBuilder.
func NewSendInteractiveMessageBuilder() *SendInteractiveMessageBuilder {
	return &SendInteractiveMessageBuilder{}
}

// WithRoomID sets the RoomID for interactive messages.
func (b *SendInteractiveMessageBuilder) WithRoomID(roomID string) *SendInteractiveMessageBuilder {
	b.RoomID = roomID
	return b
}

// WithInteractiveData sets the interactive data for messages.
func (b *SendInteractiveMessageBuilder) WithInteractiveData(
	data InteractiveData,
) *SendInteractiveMessageBuilder {
	b.InteractiveData = data
	return b
}

// Build builds the SendInteractiveMessage using the configuration from the builder.
// Example:
//
//	builder := NewSendInteractiveMessageBuilder().
//	    WithRoomID("room123").
//	    WithInteractiveData(interactiveData).
//	message := builder.Build()
func (b *SendInteractiveMessageBuilder) Build() SendInteractiveMessage {
	return SendInteractiveMessage{
		RoomID:      b.RoomID,
		Type:        "string",
		Interactive: b.InteractiveData,
	}
}

// InteractiveDataBuilder is a builder for creating interactive message data.
type InteractiveDataBuilder struct {
	header  *InteractiveHeader
	body    string
	buttons []Button
	lists   *InteractiveLists
}

// NewInteractiveDataBuilder creates a new instance of InteractiveDataBuilder.
func NewInteractiveDataBuilder() *InteractiveDataBuilder {
	return &InteractiveDataBuilder{}
}

// WithHeader sets the header of the interactive message.
func (b *InteractiveDataBuilder) WithHeader(
	header *InteractiveHeader,
) *InteractiveDataBuilder {
	b.header = header
	return b
}

// WithBody sets the body of the interactive message.
func (b *InteractiveDataBuilder) WithBody(body string) *InteractiveDataBuilder {
	b.body = body
	return b
}

// WithButtons sets the buttons of the interactive message.
func (b *InteractiveDataBuilder) WithButtons(buttons []Button) *InteractiveDataBuilder {
	b.buttons = buttons
	return b
}

// WithLists sets the lists of the interactive message.
func (b *InteractiveDataBuilder) WithLists(lists *InteractiveLists) *InteractiveDataBuilder {
	b.lists = lists
	return b
}

// Build builds the InteractiveData using the configuration from the builder.
// Example:
//
//	builder := NewInteractiveDataBuilder().
//	    WithHeader(header).
//	    WithBody("Hello, this is interactive!").
//	    WithButtons(buttons).
//	    WithLists(lists).
//	data := builder.Build()
func (b *InteractiveDataBuilder) Build() InteractiveData {
	interactiveData := InteractiveData{}

	if b.header != nil {
		interactiveData.Header = b.header
	}
	interactiveData.Body = b.body
	interactiveData.Buttons = b.buttons

	if b.lists != nil {
		interactiveData.Lists = b.lists
	}

	return interactiveData
}

// InteractiveSectionBuilder is a builder for interactive message sections
type InteractiveSectionBuilder struct {
	title string
	rows  []InteractiveRow
}

// NewInteractiveSectionBuilder creates a new instance of InteractiveSectionBuilder
func NewInteractiveSectionBuilder() *InteractiveSectionBuilder {
	return &InteractiveSectionBuilder{}
}

// WithTitle sets the title of the interactive section
func (b *InteractiveSectionBuilder) WithTitle(title string) *InteractiveSectionBuilder {
	b.title = title
	return b
}

// WithRows sets the rows of the interactive section
func (b *InteractiveSectionBuilder) WithRows(rows []InteractiveRow) *InteractiveSectionBuilder {
	b.rows = rows
	return b
}

// Build builds the InteractiveSection using the configuration from the builder.
// Example:
//
//	builder := NewInteractiveSectionBuilder().
//	    WithTitle("Section Title").
//	    WithRows(rows).
//	section := builder.Build()
func (b *InteractiveSectionBuilder) Build() InteractiveSection {
	return InteractiveSection{
		Title: b.title,
		Rows:  b.rows,
	}
}

// InteractiveRowBuilder is a builder for interactive message rows
type InteractiveRowBuilder struct {
	id          string
	title       string
	description string
}

// NewInteractiveRowBuilder creates a new instance of InteractiveRowBuilder
func NewInteractiveRowBuilder() *InteractiveRowBuilder {
	return &InteractiveRowBuilder{}
}

// WithID sets the ID of the interactive row
func (b *InteractiveRowBuilder) WithID(id string) *InteractiveRowBuilder {
	b.id = id
	return b
}

// WithTitle sets the title of the interactive row
func (b *InteractiveRowBuilder) WithTitle(title string) *InteractiveRowBuilder {
	b.title = title
	return b
}

// WithDescription sets the description of the interactive row
func (b *InteractiveRowBuilder) WithDescription(description string) *InteractiveRowBuilder {
	b.description = description
	return b
}

// Build builds the InteractiveRow using the configuration from the builder.
// Example:
//
//	builder := NewInteractiveRowBuilder().
//	    WithID("row123").
//	    WithTitle("Row Title").
//	    WithDescription("Row Description").
//	row := builder.Build()
func (b *InteractiveRowBuilder) Build() InteractiveRow {
	return InteractiveRow{
		ID:          b.id,
		Title:       b.title,
		Description: b.description,
	}
}

// WhatsAppMessageBuilder is a builder for creating WhatsApp message parameters.
type WhatsAppMessageBuilder struct {
	roomID  string
	message string
}

// NewWhatsAppMessageBuilder creates a new instance of WhatsAppMessageBuilder.
func NewWhatsAppMessageBuilder() *WhatsAppMessageBuilder {
	return &WhatsAppMessageBuilder{}
}

// WithRoomID sets the room ID for the WhatsApp message.
func (b *WhatsAppMessageBuilder) WithRoomID(roomID string) *WhatsAppMessageBuilder {
	b.roomID = roomID
	return b
}

// WithMessage sets the text message for the WhatsApp message.
func (b *WhatsAppMessageBuilder) WithMessage(message string) *WhatsAppMessageBuilder {
	b.message = message
	return b
}

// Build constructs WhatsApp message parameters using the configurations set in the builder.
// Example:
//
//	messageBuilder := NewWhatsAppMessageBuilder().
//	    WithRoomID("room123").
//	    WithMessage("Hello, this is a message!")
//	messageParams := messageBuilder.Build()
func (b *WhatsAppMessageBuilder) Build() WhatsAppMessage {
	return WhatsAppMessage{
		RoomID:  b.roomID,
		Message: b.message,
	}
}

// NewDirectWhatsAppBroadcastBuilder creates a new instance of DirectWhatsAppBroadcastBuilder.
func NewDirectWhatsAppBroadcastBuilder() *DirectWhatsAppBroadcastBuilder {
	return &DirectWhatsAppBroadcastBuilder{
		headerParams: make([]KeyValue, 0),
		bodyParams:   make([]KeyValueText, 0),
		buttons:      make([]ButtonMessage, 0),
		language:     make(map[string]string),
	}
}

// DirectWhatsAppBroadcastBuilder is a builder for creating parameters for sending direct WhatsApp broadcast.
type DirectWhatsAppBroadcastBuilder struct {
	toName               string
	toNumber             string
	messageTemplateID    string
	channelIntegrationID string
	headerParams         []KeyValue
	bodyParams           []KeyValueText
	buttons              []ButtonMessage
	language             map[string]string
}

// WithToName sets the recipient's name.
func (b *DirectWhatsAppBroadcastBuilder) WithToName(toName string) *DirectWhatsAppBroadcastBuilder {
	b.toName = toName
	return b
}

// WithToNumber sets the recipient's WhatsApp number.
func (b *DirectWhatsAppBroadcastBuilder) WithToNumber(toNumber string) *DirectWhatsAppBroadcastBuilder {
	b.toNumber = toNumber
	return b
}

// WithMessageTemplateID sets the ID of the message template to be used.
func (b *DirectWhatsAppBroadcastBuilder) WithMessageTemplateID(messageTemplateID string) *DirectWhatsAppBroadcastBuilder {
	b.messageTemplateID = messageTemplateID
	return b
}

// WithChannelIntegrationID sets the ID of the channel integration to be used.
func (b *DirectWhatsAppBroadcastBuilder) WithChannelIntegrationID(channelIntegrationID string) *DirectWhatsAppBroadcastBuilder {
	b.channelIntegrationID = channelIntegrationID
	return b
}

// WithLanguage sets the language for the message.
func (b *DirectWhatsAppBroadcastBuilder) WithLanguage(languageCode string) *DirectWhatsAppBroadcastBuilder {
	b.language["code"] = languageCode
	return b
}

// AddHeaderParam adds a key-value pair to the header parameters.
func (b *DirectWhatsAppBroadcastBuilder) AddHeaderParam(key, value string) *DirectWhatsAppBroadcastBuilder {
	b.headerParams = append(b.headerParams, KeyValue{Key: key, Value: value})
	return b
}

// AddBodyParam adds a key-value pair to the body parameters.
func (b *DirectWhatsAppBroadcastBuilder) AddBodyParam(key, valueText, value string) *DirectWhatsAppBroadcastBuilder {
	b.bodyParams = append(b.bodyParams, KeyValueText{Key: key, ValueText: valueText, Value: value})
	return b
}

// AddButton adds a button to the list of buttons.
func (b *DirectWhatsAppBroadcastBuilder) AddButton(button ButtonMessage) *DirectWhatsAppBroadcastBuilder {
	b.buttons = append(b.buttons, button)
	return b
}

// Build constructs a DirectWhatsAppBroadcastParams using the configurations set in the builder.
func (b *DirectWhatsAppBroadcastBuilder) Build() DirectWhatsAppBroadcast {
	return DirectWhatsAppBroadcast{
		ToName:               b.toName,
		ToNumber:             b.toNumber,
		MessageTemplateID:    b.messageTemplateID,
		ChannelIntegrationID: b.channelIntegrationID,
		Language:             b.language,
		HeaderParams:         b.headerParams,
		BodyParams:           b.bodyParams,
		Buttons:              b.buttons,
	}
}
