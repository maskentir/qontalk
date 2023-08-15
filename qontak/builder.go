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
//   builder := NewSendMessageInteractionsBuilder().
//       WithReceiveMessageFromAgent(true).
//       WithReceiveMessageFromCustomer(false).
//       WithStatusMessage(true).
//       WithURL("https://example.com").
//   interactions := builder.Build()
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
//   builder := NewSendInteractiveMessageBuilder().
//       WithRoomID("room123").
//       WithInteractiveData(interactiveData).
//   message := builder.Build()
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
//   builder := NewInteractiveDataBuilder().
//       WithHeader(header).
//       WithBody("Hello, this is interactive!").
//       WithButtons(buttons).
//       WithLists(lists).
//   data := builder.Build()
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
//   builder := NewInteractiveSectionBuilder().
//       WithTitle("Section Title").
//       WithRows(rows).
//   section := builder.Build()
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
//   builder := NewInteractiveRowBuilder().
//       WithID("row123").
//       WithTitle("Row Title").
//       WithDescription("Row Description").
//   row := builder.Build()
func (b *InteractiveRowBuilder) Build() InteractiveRow {
	return InteractiveRow{
		ID:          b.id,
		Title:       b.title,
		Description: b.description,
	}
}
