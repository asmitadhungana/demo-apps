// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostRegHandlerFunc turns a function with the right signature into a post reg handler
type PostRegHandlerFunc func(PostRegParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRegHandlerFunc) Handle(params PostRegParams) middleware.Responder {
	return fn(params)
}

// PostRegHandler interface for that can handle valid post reg params
type PostRegHandler interface {
	Handle(PostRegParams) middleware.Responder
}

// NewPostReg creates a new http.Handler for the post reg operation
func NewPostReg(ctx *middleware.Context, handler PostRegHandler) *PostReg {
	return &PostReg{Context: ctx, Handler: handler}
}

/*PostReg swagger:route POST /reg postReg

FE & ContentOS calls this API when the Harmony game is loaded.

*/
type PostReg struct {
	Context *middleware.Context
	Handler PostRegHandler
}

func (o *PostReg) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRegParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostRegCreatedBody post reg created body
// swagger:model PostRegCreatedBody
type PostRegCreatedBody struct {

	// account
	Account string `json:"account,omitempty"`

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this post reg created body
func (o *PostRegCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRegCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRegCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostRegCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRegOKBody post reg o k body
// swagger:model PostRegOKBody
type PostRegOKBody struct {

	// account
	Account string `json:"account,omitempty"`

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this post reg o k body
func (o *PostRegOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRegOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRegOKBody) UnmarshalBinary(b []byte) error {
	var res PostRegOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}