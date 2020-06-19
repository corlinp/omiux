package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type API struct {
	Name        string
	Description string
	Endpoints   []*Endpoint
}

func (api *API) Router() *mux.Router {
	r := mux.NewRouter()
	for _, ep := range api.Endpoints {
		for _, a := range ep.Actions {
			r.HandleFunc(ep.Path, a.contexter).Methods(a.Method)
		}
	}
	return r
}

func (api *API) GetBlueprint() string {
	out := &simpleWriter{}
	out.P("FORMAT: 1A\n")
	out.F("# %s", api.Name)
	out.P(api.Description)
	for _, ep := range api.Endpoints {
		out.P()
		out.P(ep.GetBlueprint())
	}
	return out.String()
}

func (api *API) GetCobra() *cobra.Command {
	out := &cobra.Command{
		Use: api.Name,
		Long: api.Description,
	}
	return out
}