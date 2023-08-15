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
