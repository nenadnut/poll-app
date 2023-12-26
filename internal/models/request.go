package models

// OptionRequest holds a request data for ent.Option creation
type OptionRequest struct {
	Text string `json:"text" validate:"required"`
}

// QuestionRequest holds a request data for ent.Question creation
type QuestionRequest struct {
	Text         string          `json:"text" validate:"required"`
	Options      []OptionRequest `json:"options" validate:"required"`
	NumOfAnswers int             `json:"num_of_answers"`
}

// QuestionsBulkRequest represents a batch of QuestionRequest
type QuestionsBulkRequest struct {
	Questions []QuestionRequest `json:"questions"`
}
