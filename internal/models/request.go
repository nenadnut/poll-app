package models

// OptionRequest holds a request data for ent.Option creation
type OptionRequest struct {
	Text string `json:"text" validate:"required"`
}

// OptionsBulkRequest represents a batch of OptionRequest
type OptionsBulkRequest struct {
	Options []OptionRequest `json:"options" validate:"required"`
}

// QuestionRequest holds a request data for ent.Question creation
type QuestionRequest struct {
	Text         string          `json:"text" validate:"required"`
	Options      []OptionRequest `json:"options" validate:"required"`
	NumOfAnswers int             `json:"num_of_answers"`
	Required     bool            `json:"required"`
}

// QuestionsBulkRequest represents a batch of QuestionRequest
type QuestionsBulkRequest struct {
	Questions []QuestionRequest `json:"questions"`
}
