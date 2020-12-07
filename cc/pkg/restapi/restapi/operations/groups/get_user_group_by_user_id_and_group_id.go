// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserGroupByUserIDAndGroupIDHandlerFunc turns a function with the right signature into a get user group by user Id and group Id handler
type GetUserGroupByUserIDAndGroupIDHandlerFunc func(GetUserGroupByUserIDAndGroupIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserGroupByUserIDAndGroupIDHandlerFunc) Handle(params GetUserGroupByUserIDAndGroupIDParams) middleware.Responder {
	return fn(params)
}

// GetUserGroupByUserIDAndGroupIDHandler interface for that can handle valid get user group by user Id and group Id params
type GetUserGroupByUserIDAndGroupIDHandler interface {
	Handle(GetUserGroupByUserIDAndGroupIDParams) middleware.Responder
}

// NewGetUserGroupByUserIDAndGroupID creates a new http.Handler for the get user group by user Id and group Id operation
func NewGetUserGroupByUserIDAndGroupID(ctx *middleware.Context, handler GetUserGroupByUserIDAndGroupIDHandler) *GetUserGroupByUserIDAndGroupID {
	return &GetUserGroupByUserIDAndGroupID{Context: ctx, Handler: handler}
}

/*GetUserGroupByUserIDAndGroupID swagger:route GET /cc/v1/groups/userGroup/user/{userId}/group/{groupId} Groups getUserGroupByUserIdAndGroupId

Returns a userGroup.

Optional extended description in Markdown.

*/
type GetUserGroupByUserIDAndGroupID struct {
	Context *middleware.Context
	Handler GetUserGroupByUserIDAndGroupIDHandler
}

func (o *GetUserGroupByUserIDAndGroupID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserGroupByUserIDAndGroupIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}