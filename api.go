package omiux

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
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
					req, err := http.NewRequest(a.Method, api.Host + ep.Path, nil)
					if err != nil {
						panic(err)
					}
					q := req.URL.Query()
					for _, p := range a.Params {
						flag, err := cmd.Flags().GetString(p.Info().Name)
						if err == nil {
							q.Add(p.Info().Name, flag)
						}
					}
					req.URL.RawQuery = q.Encode()
					_, _ = fmt.Fprintf(os.Stderr, "%s %s", a.Method, q.Encode())
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						panic(err)
					}
					d, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						panic(err)
					}
					fmt.Println(string(d))
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