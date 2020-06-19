package main

import (
	"encoding/json"
	"net/http"
)

type Action struct {
	Method      string
	Name        string
	Description string
	Params      []Param
	Response    interface{}
	Run         func(w http.ResponseWriter, r *http.Request)
}

func (a *Action) GetBlueprint() string {
	out := &blueprintWriter{}
	out.F("### %s [%s]\n", a.Name, a.Method)
	out.P(a.Description)
	out.P("\n")
	if len(a.Params) > 0 {
		out.P("+ Parameters\n")
		for _, p := range a.Params {
			out.P(GetParamBlueprint(p))
		}
	}
	if a.Response != nil {
		out.P("+ Response 200  (application/json)\n\n\t+ Body\n")
		j, _ := json.MarshalIndent(a.Response, "\t\t", "    ")
		out.S("\t\t")
		out.Write(j)
		out.P("\n")
	}
	return out.String()
}
