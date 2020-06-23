package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"strings"
)

type API struct {
	Name        string
	Description string
	Host string
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
	out.P("FORMAT: 1A")
	out.P("HOST: "+api.Host)
	out.P()
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
		Short: api.Name,
		Long: api.Description,
	}
	for _, ep := range api.Endpoints {
		epCmd := &cobra.Command{
			Use: strings.ToLower(strings.ReplaceAll(ep.Name, " ", "-")),
			Short: ep.Description,
		}
		for _, a := range ep.Actions {
			aCmd := &cobra.Command{
				Use: strings.ToLower(strings.ReplaceAll(a.Name, " ", "-")),
				Short: a.Name,
				Long: a.Description,
				Run: func(cmd *cobra.Command, args []string) {
					v, err := cmd.Flags().GetString("deviceType")
					fmt.Println(v, err)
				},
			}
			for _, p := range a.Params {
				aCmd.Flags().StringP(p.Info().Name, string(p.Info().Name[0]), p.Info().Default, p.Info().Description)
			}
			epCmd.AddCommand(aCmd)
		}
		out.AddCommand(epCmd)
	}
	return out
}