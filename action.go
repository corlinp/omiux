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
	out := &simpleWriter{}
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

type ActionContext struct {
	params map[string]interface{}
}

func (a *ActionContext) GetStringParam(name string) string {
	out, ok := a.params[name].(string)
	if !ok {
		panic("param is not of the right type!")
	}
	return out
}

func (a *Action) parseRequest(w http.ResponseWriter, r *http.Request) (*ActionContext, error) {
	out := &ActionContext{
		params: make(map[string]interface{}),
	}
	for _, p := range a.Params {
		s := r.URL.Query().Get(p.GetName())
		v, err := p.Parse(s)
		if err != nil {
			return nil, err
		}
		out.params[p.GetName()] = v
	}
	return out, nil
}