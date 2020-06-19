package main


type API struct {
	Name        string
	Description string
	Endpoints   []*Endpoint
}

func (api *API) GetBlueprint() string {
	out := &blueprintWriter{}
	out.P("FORMAT: 1A\n")
	out.F("# %s", api.Name)
	out.P(api.Description)
	for _, ep := range api.Endpoints {
		out.P()
		out.P(ep.GetBlueprint())
	}
	return out.String()
}

