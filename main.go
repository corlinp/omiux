package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type DevicesResponse struct {
	DeviceType string `json:"deviceType"`
}

var limitParam = NewIntParam("limit", "return up to this many records", 100)

func main() {

	//devices := NewEndpoint("v1/devices", "Devices",
	//	"A device represents a physical thing - a big, a little, a rocbox, etc.")
	//
	//devicesGet := devices.NewAction("GET", "List Devices",
	//			"Return a JSON list of devices matching query filters")
	//devicesGet.AddParams(limitParam,
	//	NewStringParam("deviceType", "this is the type of device", ""))

	api := &API{
			name: "ROC (Robot Operating Center) API",
			description: "An API for all things Brain Corp",
			endpoints: []*Endpoint{
				{
					path:        "/v1/devices",
					name:        "Devices",
					description: "A device represents a physical thing - a big, a little, a rocbox, etc.",
					actions: []*Action{
						{
							method:      "GET",
							name:        "List Devices",
							description: "Return a JSON list of devices matching query filters",
							params: []Param{
								&StringParam{
									name:         "deviceType",
									description:  "type of device, like BCM-scrubber, LBCM-whiz, etc.",
									defaultValue: "",
								},
								limitParam,
							},
							run: func(w http.ResponseWriter, r *http.Request) {
								fmt.Fprint(w, "hello!")
							},
							response: DevicesResponse{},
						},
					},
				},
			},
	}

	bp := api.GetBlueprint()
	fmt.Println(bp)

	//r := devices.Router()
	//http.ListenAndServe(":8080", r)

	//mk := markdown.ToHTML([]byte(bp), nil, nil)
	//http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request){
	//	w.Header().Set("Content-Type", "text/html")
	//	w.Write(mk)
	//})
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

type API struct {
	name string
	description string
	endpoints []*Endpoint
}

func (api *API) GetBlueprint() string {
	out := &blueprintWriter{}
	out.P("FORMAT: 1A\n")
	out.F("# %s", api.name)
	out.P(api.description)
	for _, ep := range api.endpoints {
		out.P()
		out.P(ep.GetBlueprint())
	}
	return out.String()
}

type Action struct {
	method string
	name string
	description string
	params []Param
	response interface{}
	run func(w http.ResponseWriter, r *http.Request)
}

func (ep *Endpoint) NewAction(m, n, d string) *Action {
	a := &Action{
		method:   m,
		name: n,
		description: d,
		params: []Param{},
	}
	ep.actions = append(ep.actions, a)
	return a
}

func (a *Action) AddParams(p ...Param) *Action {
	a.params = append(a.params, p...)
	return a
}

func (a *Action) GetBlueprint() string {
	out := &blueprintWriter{}
	out.F("### %s [%s]\n", a.name, a.method)
	out.P(a.description)
	out.P("\n")
	if len(a.params) > 0 {
		out.P("+ Parameters\n")
		for _, p := range a.params {
			out.P(GetParamBlueprint(p))
		}
	}
	if a.response != nil {
		out.P("+ Response 200  (application/json)\n\n\t+ Body\n")
		j, _ := json.MarshalIndent(a.response, "\t\t", "    ")
		out.S("\t\t")
		out.Write(j)
		out.P("\n")
	}
	return out.String()
}

type Endpoint struct {
	path   string
	name string
	description string
	actions []*Action
}

func NewEndpoint(p, n, d string) *Endpoint {
	return &Endpoint{
		path:   p,
		name: n,
		description: d,
		actions: []*Action{},
	}
}

func (ep *Endpoint) AddActions(a ...*Action) *Endpoint {
	ep.actions = append(ep.actions, a...)
	return ep
}

func (ep *Endpoint) Router() *mux.Router {
	r := mux.NewRouter()
	for _, a := range ep.actions {
		r.HandleFunc(ep.path, a.run).Methods(a.method)
	}
	return r
}

func (ep *Endpoint) GetBlueprint() string {
	out := &blueprintWriter{}
	out.F("## %s [%s]\n", ep.name, ep.path)
	out.P(ep.description)
	for _, a := range ep.actions {
		out.P()
		out.P(a.GetBlueprint())
	}
	out.P()
	return out.String()
}


type ActionRequest struct {
	action *Action
}

type ActionRunner func(a *ActionRequest, w http.ResponseWriter, r *http.Request)







