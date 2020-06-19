package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type Param interface {
	GetName() string
	GetDescription() string
	TypeName() string
	GetDefault() string
	IsRequired() bool
	Parse(s string) (interface{}, error)
}

// STRING

type StringParam struct {
	Name string
	Description string
	Default string
	required bool
}

func (p *StringParam) GetName() string {
	return p.Name
}

func (p *StringParam) GetDescription() string {
	return p.Description
}

func (p *StringParam) Get() string {
	return p.Default
}

func (p *StringParam) TypeName() string {
	return "string"
}

func (p *StringParam) IsRequired() bool {
	return p.required
}

func (p *StringParam) GetDefault() string {
	return fmt.Sprint(p.Default)
}

func (p *StringParam) Parse(s string) (interface{}, error) {
	if s == "" {
		if p.required {
			return nil, errors.New("this should be a nice error")
		}
		return p.Default, nil
	}
	return s, nil
}

func GetParamBlueprint(p Param) string {
	out := simpleWriter{}
	reqString := "optional"
	if p.IsRequired() {
		reqString = "Required"
	}
	out.F("\t+ %s: (%s, %s) - %s", p.GetName(), p.TypeName(), reqString, p.GetDescription())
	if p.GetDefault() != "" {
		out.F("\t\t+ Default: %s\n", p.GetDefault())
	}
	return out.String()
}

// INT

type IntParam struct {
	Name        string
	Description string
	Default     int
	Required    bool
}

func (p *IntParam) GetName() string {
	return p.Name
}

func (p *IntParam) GetDescription() string {
	return p.Description
}

func (p *IntParam) Get() int {
	return p.Default
}

func (p *IntParam) TypeName() string {
	return "number"
}

func (p *IntParam) IsRequired() bool {
	return p.Required
}

func (p *IntParam) GetDefault() string {
	return fmt.Sprint(p.Default)
}

func (p *IntParam) Parse(s string) (interface{}, error) {
	if s == "" {
		if p.Required {
			return nil, errors.New("this should be a nice int error")
		}
		return p.Default, nil
	}
	return strconv.ParseInt(s, 10, 32)
}

























