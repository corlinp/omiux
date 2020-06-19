package main

import "fmt"

type Param interface {
	GetName() string
	GetDescription() string
	TypeName() string
	Default() string
	Required() bool
}

// STRING

type StringParam struct {
	name string
	description string
	defaultValue string
	required bool
}

func (p *StringParam) GetName() string {
	return p.name
}

func (p *StringParam) GetDescription() string {
	return p.description
}

func (p *StringParam) Get() string {
	return p.defaultValue
}

func (p *StringParam) TypeName() string {
	return "string"
}

func (p *StringParam) Required() bool {
	return p.required
}

func (p *StringParam) Default() string {
	return fmt.Sprint(p.defaultValue)
}

func NewStringParam(name, description string, defaultValue string) *StringParam {
	return &StringParam{
		name:         name,
		description:  description,
		defaultValue: defaultValue,
	}
}

func GetParamBlueprint(p Param) string {
	out := blueprintWriter{}
	reqString := "optional"
	if p.Required() {
		reqString = "required"
	}
	out.F("\t+ %s: (%s, %s) - %s", p.GetName(), p.TypeName(), reqString, p.GetDescription())
	if p.Default() != "" {
		out.F("\t\t+ Default: %s\n", p.Default())
	}
	return out.String()
}

// INT

type IntParam struct {
	name string
	description string
	defaultValue int
	required bool
}

func (p *IntParam) GetName() string {
	return p.name
}

func (p *IntParam) GetDescription() string {
	return p.description
}

func (p *IntParam) Get() int {
	return p.defaultValue
}

func (p *IntParam) TypeName() string {
	return "number"
}

func (p *IntParam) Required() bool {
	return p.required
}

func (p *IntParam) Default() string {
	return fmt.Sprint(p.defaultValue)
}

func NewIntParam(name, description string, defaultValue int) *IntParam {
	return &IntParam{
		name:        name,
		description: description,
		defaultValue: defaultValue,
	}
}