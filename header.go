package omiux

import "fmt"

type RequestHeader struct {
	Name string
	Example string
	Default string
	Required bool
}

func (h *RequestHeader) GetBlueprint() string {
	return fmt.Sprintf("            %s: %s", h.Name, h.Example)
}

type ResponseHeader struct {
	Name string
	Example string
}

func (h *ResponseHeader) GetBlueprint() string {
	return fmt.Sprintf("            %s: %s", h.Name, h.Example)
}

var headerRequestID = &ResponseHeader{
	Name:     "X-Request-ID",
	Example:  "01D78XYFJ1PRM1WPBCBT3VHMNV",
}

var headerContentLength = &ResponseHeader{
	Name:     "Content-Length",
	Example:  "3376",
}