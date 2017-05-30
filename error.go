package main

type ErrorDuplicatedEmail struct { // TODO: ONLY for testing
}

func (e *ErrorDuplicatedEmail) Error() string {
	return "Duplicated email"
}
