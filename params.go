package omiux

import (
	"github.com/pkg/errors"
	"strconv"
)

type Param interface {
	Parse(s string) (interface{}, error)
	TypeName() string
	Info() *ParamInfo
}

type ParamInfo struct {
	Name        string
	Description string
	Example     string
	Default     string
	Required    bool
}


func GetParamBlueprint(p Param) string {
	out := simpleWriter{}
	reqString := "optional"
	if p.Info().Required {
		reqString = "Required"
	}
	out.F("    + %s: `%s` (%s, %s) - %s", p.Info().Name, p.Info().Example, p.TypeName(), reqString, p.Info().Description)
	if p.Info().Default != "" {
		out.F("        + Default: %s", p.Info().Default)
	}
	return out.String()
}

// STRING

type StringParam ParamInfo

func (p *StringParam) TypeName() string {
	return "string"
}

func (p *StringParam) Parse(s string) (interface{}, error) {
	if s == "" {
		if p.Required {
			return nil, errors.New("required string parameter is not present")
		}
		return p.Default, nil
	}
	return s, nil
}

func (p *StringParam) Info() *ParamInfo {
	return (*ParamInfo)(p)
}


// INT

type IntParam ParamInfo

func (p *IntParam) TypeName() string {
	return "number"
}

func (p *IntParam) Parse(s string) (interface{}, error) {
	if s == "" {
		if p.Required {
			return nil, errors.New("required integer parameter is not present")
		}
		return strconv.ParseInt(p.Default, 10, 32)
	}
	return strconv.ParseInt(s, 10, 32)
}

func (p *IntParam) Info() *ParamInfo {
	return (*ParamInfo)(p)
}

