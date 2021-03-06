// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/omarvides/goalearn/design
// --out=$(GOPATH)/src/github.com/omarvides/goalearn
// --version=v1.1.0-dirty
//
// API "user": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userID"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		if userID, err2 := strconv.Atoi(rawUserID); err2 == nil {
			rctx.UserID = userID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("userID", rawUserID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *UserAPI) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/user.api+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
