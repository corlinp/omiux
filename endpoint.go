package main

import "github.com/gorilla/mux"

type Endpoint struct {
	Path        string
	Name        string
	Description string
	Actions     []*Action
}

func (ep *Endpoint) Router() *mux.Router {
	r := mux.NewRouter()
	for _, a := range ep.Actions {
		r.HandleFunc(ep.Path, a.Run).Methods(a.Method)
	}
	return r
}

func (ep *Endpoint) GetBlueprint() string {
	out := &blueprintWriter{}
	out.F("## %s [%s]\n", ep.Name, ep.Path)
	out.P(ep.Description)
	for _, a := range ep.Actions {
		out.P()
		out.P(a.GetBlueprint())
	}
	out.P()
	return out.String()
}

