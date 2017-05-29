package main

type Action interface {
	Action() string
}

type ActionSessionInit struct {
	SessionID string `json:"session_id"`
}

func (a ActionSessionInit) Action() string {
	return "session_init"
}
