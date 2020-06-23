package omiux

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type Action struct {
	Method      string
	Name        string
	Description string
	Params      []Param
	Headers      []*Header
	Request interface{}
	Response    interface{}
	Errors 	    []*Error
	Run         func(a *Context) (interface{}, *Error)
}

func (a *Action) GetBlueprint(ep *Endpoint) string {
	out := &simpleWriter{}
	out.F("### %s [%s]", a.Name, a.Method)
	out.P(a.Description)
	out.P()
	if len(a.Params) > 0 {
		out.P("+ Parameters")
		for _, p := range a.Params {
			out.S(GetParamBlueprint(p))
		}
		out.P()
	}
	if a.Request != nil {
		out.P("+ Request (application/json)")
	} else {
		out.P("+ Request")
	}
	out.P()
	if len(a.Headers) > 0 {
		out.P("    + Headers\n")
		for _, h := range a.Headers {
			out.P(h.GetBlueprint())
		}
		out.P()
	}
	if a.Request != nil {
		out.P("    + Body\n")
		j, _ := json.MarshalIndent(a.Request, "            ", "    ")
		out.S("            ")
		out.Write(j)
		out.P("\n")
	}

	if a.Response != nil {
		out.P("+ Response 200  (application/json)\n\n    + Body\n")
		j, _ := json.MarshalIndent(a.Response, "            ", "    ")
		out.S("            ")
		out.Write(j)
		out.P("\n")
	}
	if len(a.Errors) > 0 {
		for _, e := range a.Errors {
			out.F("+ Response %v  (application/json)\n\n    + Body\n", e.Status)
			j, _ := json.MarshalIndent(&ErrorResponse{
				Status:  e.Status,
				Code:    e.Code,
				Message: e.Message,
				Info:    "-",
				Path:    ep.Path,
			}, "            ", "    ")
			out.S("            ")
			out.Write(j)
			out.P("\n")
		}
	}
	return out.String()
}

type Context struct {
	Action *Action
	Writer http.ResponseWriter
	Request *http.Request
	params map[string]interface{}
}

func (a *Context) GetStringParam(name string) string {
	out, ok := a.params[name].(string)
	if !ok {
		panic("param is not of the right type!")
	}
	return out
}

func (a *Context) GetIntParam(name string) int64 {
	out, ok := a.params[name].(int64)
	if !ok {
		panic("param is not of the right type!")
	}
	return out
}

func (a *Action) parseRequest(w http.ResponseWriter, r *http.Request) (*Context, error) {
	out := &Context{
		Action: a,
		Writer: w,
		Request: r,
		params: make(map[string]interface{}),
	}
	for _, p := range a.Params {
		s := r.URL.Query().Get(p.Info().Name)
		v, err := p.Parse(s)
		if err != nil {
			return nil, errors.Wrap(err, "parsing " + p.Info().Name)
		}
		out.params[p.Info().Name] = v
	}
	return out, nil
}

func (a *Action) contexter(w http.ResponseWriter, r *http.Request) {
	ac, err := a.parseRequest(w, r)
	if err != nil {
		w.WriteHeader(ErrParsingParameter.Status)
		errResp := &ErrorResponse{
			Status:  ErrParsingParameter.Status,
			Code:    ErrParsingParameter.Code,
			Message: ErrParsingParameter.Message,
			Info:    err.Error(),
			Path:    r.RequestURI,
		}
		_ = json.NewEncoder(w).Encode(errResp)
		return
	}
	out, rerr := a.Run(ac)
	if rerr != nil {
		w.WriteHeader(rerr.Status)
		errResp := &ErrorResponse{
			Status:  rerr.Status,
			Code:    rerr.Code,
			Message: rerr.Message,
			Info:    rerr.info,
			Path:    r.RequestURI,
		}
		_ = json.NewEncoder(w).Encode(errResp)
		return
	}
	if out != nil {
		_ = json.NewEncoder(w).Encode(out)
		return
	}
	return
}
