package omiux

import (
	"fmt"
	"testing"
)


type User struct {
	Name string
	Age int
}

func TestOmiux(t *testing.T) {
	api := &API{
		Name:        "Users API",
		Description: "Store and list your users here",
		Host:        "https://user-api.com/",
		Endpoints:   []*Endpoint{
			{
				Path:        "/v1/users",
				Name:        "Users",
				Description: "This is the endpoint to add or get users",
				CmdName:     "users",
				Actions:     []*Action{
					{
						Method:          "GET",
						Name:            "Get Users",
						Description:     "Returns a list of users",
						CmdName:         "get",
						Params:          []Param {
							&IntParam{
								Name: "age",
								Description: "age of the user",
								Example: "56",
							},
							&StringParam{
								Name: "name",
								Description: "name of the user",
								Example: "bob",
							},
						},
						Response:        User{
							Name: "patrick star",
							Age:  34,
						},
					},
					{
						Method:          "POST",
						Name:            "Add User",
						Description:     "Create a new user",
						CmdName:         "add",
						Request:          User{
							Name: "sandy cheeks",
							Age:  33,
						},
					},
				},
			},
		},
	}
	fmt.Println(api.GetBlueprint())
}