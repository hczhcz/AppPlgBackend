package main

func actionSessionInit() (string, interface{}) {
	type response struct {
		SessionID string `json:"session_id"`
	}

	return "session_init", response{"0000"} // TODO: mock
}

func actionInvalidRequest(hint string) (string, interface{}) {
	type response struct {
		Hint string `json:"hint"`
	}

	return "invalid_request", response{hint}
}

func actionInternalError() (string, interface{}) {
	type response struct {
	}

	return "internal_error", response{}
}

func actionDuplicatedEmail() (string, interface{}) { // TODO: ONLY for testing
	type response struct {
	}

	return "duplicated_email", response{}
}

type Action interface {
	Action() (string, interface{})
}

func (e ErrorDuplicatedEmail) Action() (string, interface{}) {
	return actionDuplicatedEmail()
}

func GetActionFromError(e error) (string, interface{}) {
	if a, ok := e.(Action); ok {
		return a.Action()
	} else {
		return actionInternalError()
	}
}
