package main

// ActionOK represents success status.
type ActionOK struct{}

type Action interface {
	Action() string
}

type ActionSessionInit struct {
	SessionID string `json:"session_id"`
}

func (a ActionSessionInit) Action() string {
	return "session_init"
}

type ActionInvalidRequest struct {
	Hint string `json:"hint"`
}

func (a ActionInvalidRequest) Action() string {
	return "invalid_request"
}

type ActionInternalError struct {
}

func (a ActionInternalError) Action() string {
	return "internal_error"
}
