# omiux

Omiux is a Go API schema system. Define your API and Omiux will turn it into a router, CLI, and docs. Keep everything in one struct!

```go
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
```

`api.GetBlueprint()`:

=== RUN   TestOmiux
FORMAT: 1A
HOST: https://user-api.com/

# Users API
Store and list your users here

## Users [/v1/users{?age,name}]


This is the endpoint to add or get users

### Get Users [GET]
Returns a list of users

+ Parameters
    + age: `56` (number, optional) - age of the user
    + name: `bob` (string, optional) - name of the user

+ Request

    + Headers


+ Response 200  (application/json)

    + Body

            {
                "Name": "patrick star",
                "Age": 34
            }



### Add User [POST]
Create a new user

+ Request (application/json)

    + Body

            {
                "Name": "sandy cheeks",
                "Age": 33
            }

+ Response 200  (application/json)

    + Headers

            X-Request-ID: 01D78XYFJ1PRM1WPBCBT3VHMNV
