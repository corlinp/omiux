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
	Default string
	Required bool
}

func (h *ResponseHeader) GetBlueprint() string {
	return fmt.Sprintf("            %s: %s", h.Name, h.Example)
}