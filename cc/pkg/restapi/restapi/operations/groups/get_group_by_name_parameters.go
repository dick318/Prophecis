// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetGroupByNameParams creates a new GetGroupByNameParams object
// no default values defined in spec.
func NewGetGroupByNameParams() GetGroupByNameParams {

	return GetGroupByNameParams{}
}

// GetGroupByNameParams contains all the bound params for the get group by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetGroupByName
type GetGroupByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the id of group
	  Required: true
	  In: path
	*/
	GroupName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetGroupByNameParams() beforehand.
func (o *GetGroupByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rGroupName, rhkGroupName, _ := route.Params.GetOK("groupName")
	if err := o.bindGroupName(rGroupName, rhkGroupName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindGroupName binds and validates parameter GroupName from path.
func (o *GetGroupByNameParams) bindGroupName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.GroupName = raw

	return nil
}