package main

import "fmt"

var HeaderBearerToken = &Header{
	Name:     "Authorization",
	Example:  "Bearer XXX.YYY.ZZZ",
	Required: true,
}

type Header struct {
	Name string
	Example string
	Default string
	Required bool
}

func (h *Header) GetBlueprint() string {
	return fmt.Sprintf("            %s: %s", h.Name, h.Example)
}