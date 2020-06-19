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
	Default:     "100",
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
							//&StringParam{
							//	Name:         "deviceType",
							//	Description:  "type of device as found in /v1/device-types",
							//	Example: "BCM-scrubber",
							//},
							limitParam,
						},
						Run: func(a *Context) (interface{}, *Error) {
							fmt.Println("deviceType:", a.GetStringParam("deviceType"))
							fmt.Println("limit:", a.GetIntParam("limit"))
							return &DevicesResponse{DeviceType: "catdog"}, nil
						},
						Response: DevicesResponse{},
						Errors: []*Error{
							ErrUnauthorized,
						},
					},
				},
			},
		},
	}

	bp := api.GetBlueprint()
	fmt.Println(bp)

	r := api.Router()
	http.ListenAndServe(":8080", r)

	//mk := markdown.ToHTML([]byte(bp), nil, nil)
	//http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request){
	//	w.Header().Set("Content-Type", "text/html")
	//	w.Write(mk)
	//})
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))
}



