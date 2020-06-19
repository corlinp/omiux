package main

import (
	"fmt"
	"net/http"
)

type DevicesResponse struct {
	DeviceType string `json:"deviceType"`
}

var limitParam = &IntParam{
	Name:        "limit",
	Description: "return up to this many records",
	Default:     100,
}

func main() {
	api := &API{
		Name: "ROC (Robot Operating Center) API",
		Description: "An API for all things Brain Corp",
		Endpoints: []*Endpoint{
			{
				Path:        "/v1/devices",
				Name:        "Devices",
				Description: "A device represents a physical thing - a big, a little, a rocbox, etc.",
				Actions: []*Action{
					{
						Method:      "GET",
						Name:        "List Devices",
						Description: "Return a JSON list of devices matching query filters",
						Params: []Param{
							&StringParam{
								Name:         "deviceType",
								Description:  "type of device, like BCM-scrubber, LBCM-whiz, etc.",
								Default: "",
							},
							limitParam,
						},
						Run: func(w http.ResponseWriter, r *http.Request) {
							fmt.Fprint(w, "hello!")
						},
						Response: DevicesResponse{},
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


type ActionRequest struct {
	action *Action
}

type ActionRunner func(a *ActionRequest, w http.ResponseWriter, r *http.Request)







