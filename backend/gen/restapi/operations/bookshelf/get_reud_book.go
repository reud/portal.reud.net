// Code generated by go-swagger; DO NOT EDIT.

package bookshelf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReudBookHandlerFunc turns a function with the right signature into a get reud book handler
type GetReudBookHandlerFunc func(GetReudBookParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReudBookHandlerFunc) Handle(params GetReudBookParams) middleware.Responder {
	return fn(params)
}

// GetReudBookHandler interface for that can handle valid get reud book params
type GetReudBookHandler interface {
	Handle(GetReudBookParams) middleware.Responder
}

// NewGetReudBook creates a new http.Handler for the get reud book operation
func NewGetReudBook(ctx *middleware.Context, handler GetReudBookHandler) *GetReudBook {
	return &GetReudBook{Context: ctx, Handler: handler}
}

/*GetReudBook swagger:route GET /bookshelf bookshelf getReudBook

Get books I read, from DB

DBに保存された本の情報の取得

*/
type GetReudBook struct {
	Context *middleware.Context
	Handler GetReudBookHandler
}

func (o *GetReudBook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReudBookParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}