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
