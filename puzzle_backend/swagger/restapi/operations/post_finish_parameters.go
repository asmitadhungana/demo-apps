// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostFinishParams creates a new PostFinishParams object
// no default values defined in spec.
func NewPostFinishParams() PostFinishParams {

	return PostFinishParams{}
}

// PostFinishParams contains all the bound params for the post finish operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostFinish
type PostFinishParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the solution's height (the same number found in all squares)
	  In: query
	*/
	Height *int64
	/*user's account private key, hex-encoded
	  In: query
	*/
	Key *string
	/*where the cursor was after completing the last move in sequence, in telephone keypad notation (1-9)
	  In: query
	*/
	LastPos *int64
	/*level number (1-based)
	  In: query
	*/
	Level *int64
	/*user's moves from first to last; [udlr]* in regex
	  In: query
	*/
	Sequence *string
	/*the game ID (staking transaction ID) returned by POST /play
	  In: query
	*/
	Txid *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostFinishParams() beforehand.
func (o *PostFinishParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qHeight, qhkHeight, _ := qs.GetOK("height")
	if err := o.bindHeight(qHeight, qhkHeight, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qLastPos, qhkLastPos, _ := qs.GetOK("last_pos")
	if err := o.bindLastPos(qLastPos, qhkLastPos, route.Formats); err != nil {
		res = append(res, err)
	}

	qLevel, qhkLevel, _ := qs.GetOK("level")
	if err := o.bindLevel(qLevel, qhkLevel, route.Formats); err != nil {
		res = append(res, err)
	}

	qSequence, qhkSequence, _ := qs.GetOK("sequence")
	if err := o.bindSequence(qSequence, qhkSequence, route.Formats); err != nil {
		res = append(res, err)
	}

	qTxid, qhkTxid, _ := qs.GetOK("txid")
	if err := o.bindTxid(qTxid, qhkTxid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHeight binds and validates parameter Height from query.
func (o *PostFinishParams) bindHeight(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("height", "query", "int64", raw)
	}
	o.Height = &value

	return nil
}

// bindKey binds and validates parameter Key from query.
func (o *PostFinishParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}

// bindLastPos binds and validates parameter LastPos from query.
func (o *PostFinishParams) bindLastPos(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("last_pos", "query", "int64", raw)
	}
	o.LastPos = &value

	return nil
}

// bindLevel binds and validates parameter Level from query.
func (o *PostFinishParams) bindLevel(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("level", "query", "int64", raw)
	}
	o.Level = &value

	return nil
}

// bindSequence binds and validates parameter Sequence from query.
func (o *PostFinishParams) bindSequence(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Sequence = &raw

	return nil
}

// bindTxid binds and validates parameter Txid from query.
func (o *PostFinishParams) bindTxid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Txid = &raw

	return nil
}