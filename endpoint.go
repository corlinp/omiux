package omiux

import (
	"fmt"
	"strings"
)

type Endpoint struct {
	Path        string
	Name        string
	Description string
	CmdName 	string
	Actions     []*Action
}


func (ep *Endpoint) GetBlueprint() string {
	allParams := map[string]string{}
	for _, a := range ep.Actions {
		for _, p := range a.Params {
			allParams[p.Info().Name] = p.Info().Name
		}
	}
	allParamsList := []string{}
	for n := range allParams {
		allParamsList = append(allParamsList, n)
	}
	out := &simpleWriter{}
	paramString := ""
	if len(allParamsList) > 0 {
		paramString = fmt.Sprintf("{?%s}", strings.Join(allParamsList, ","))
	}
	out.FS("## %s [%s%s]\n", ep.Name, ep.Path, paramString)
	out.P()
	out.P()
	out.P(ep.Description)
	for _, a := range ep.Actions {
		out.P()
		out.P(a.GetBlueprint(ep))
	}
	out.P()
	return out.String()
}

