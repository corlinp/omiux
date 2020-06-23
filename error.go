package main

import "net/http"

var ErrUnauthorized = &Error {
	Status: http.StatusUnauthorized,
	Code: "Unauthorized",
	Message: "you do not have permission to perform this action",
}

var ErrParsingParameter = &Error{
	Status:  http.StatusBadRequest,
	Code:    "ErrParsingParam",
	Message: "Could not parse query parameter",
}

/*
This is an attempt to make debugging the API more usable by providing standard info in an error
Let's take a look at a standard AWS error response:
<Error>
  <Code>NoSuchKey</Code>
  <Message>The resource you requested does not exist</Message>
  <Resource>/mybucket/myfoto.jpg</Resource>
  <RequestId>4442587FB7D0A2F9</RequestId>
</Error>
- CODE allows clients to programmatically identify the error type
- MESSAGE is helpful for users trying to debug their own issues
- RESOURCE and REQUESTID help the AWS team track down reported issues by tracing the request
We're going to mimic that behavior with our error structs here
*/

// Error is an instance of that error
type Error struct {
	Status    int         `json:"status,omitempty"`
	Code      string      `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	info string
}

func (e *Error) WithInfo(info string) *Error {
	return &Error{
		Status:  e.Status,
		Code:    e.Code,
		Message: e.Message,
		info:    info,
	}
}

type ErrorResponse struct {
	Status    int         `json:"status,omitempty"`
	Code      string      `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Info      string      `json:"info,omitempty"`
	Path      string      `json:"resource,omitempty"`
	//RequestId string `json:"request_id,omitempty"`
}