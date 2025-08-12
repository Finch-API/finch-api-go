// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// RequestForwardingService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestForwardingService] method instead.
type RequestForwardingService struct {
	Options []option.RequestOption
}

// NewRequestForwardingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRequestForwardingService(opts ...option.RequestOption) (r *RequestForwardingService) {
	r = &RequestForwardingService{}
	r.Options = opts
	return
}

// The Forward API allows you to make direct requests to an employment system. If
// Finch's unified API doesn't have a data model that cleanly fits your needs, then
// Forward allows you to push or pull data models directly against an integration's
// API.
func (r *RequestForwardingService) Forward(ctx context.Context, body RequestForwardingForwardParams, opts ...option.RequestOption) (res *RequestForwardingForwardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "forward"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RequestForwardingForwardResponse struct {
	// A string representation of the HTTP response body of the forwarded request's
	// response received from the underlying integration's API. This field may be null
	// in the case where the upstream system's response is empty.
	Data string `json:"data,required,nullable"`
	// The HTTP headers of the forwarded request's response, exactly as received from
	// the underlying integration's API.
	Headers interface{} `json:"headers,required,nullable"`
	// An object containing details of your original forwarded request, for your ease
	// of reference.
	Request RequestForwardingForwardResponseRequest `json:"request,required"`
	// The HTTP status code of the forwarded request's response, exactly received from
	// the underlying integration's API. This value will be returned as an integer.
	StatusCode int64                                `json:"statusCode,required"`
	JSON       requestForwardingForwardResponseJSON `json:"-"`
}

// requestForwardingForwardResponseJSON contains the JSON metadata for the struct
// [RequestForwardingForwardResponse]
type requestForwardingForwardResponseJSON struct {
	Data        apijson.Field
	Headers     apijson.Field
	Request     apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestForwardingForwardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestForwardingForwardResponseJSON) RawJSON() string {
	return r.raw
}

// An object containing details of your original forwarded request, for your ease
// of reference.
type RequestForwardingForwardResponseRequest struct {
	// The body that was specified for the forwarded request. If a value was not
	// specified in the original request, this value will be returned as null ;
	// otherwise, this value will always be returned as a string.
	Data string `json:"data,required,nullable"`
	// The specified HTTP headers that were included in the forwarded request. If no
	// headers were specified, this will be returned as `null`.
	Headers interface{} `json:"headers,required,nullable"`
	// The HTTP method that was specified for the forwarded request. Valid values
	// include: `GET` , `POST` , `PUT` , `DELETE` , and `PATCH`.
	Method string `json:"method,required"`
	// The query parameters that were included in the forwarded request. If no query
	// parameters were specified, this will be returned as `null`.
	Params interface{} `json:"params,required,nullable"`
	// The URL route path that was specified for the forwarded request.
	Route string                                      `json:"route,required"`
	JSON  requestForwardingForwardResponseRequestJSON `json:"-"`
}

// requestForwardingForwardResponseRequestJSON contains the JSON metadata for the
// struct [RequestForwardingForwardResponseRequest]
type requestForwardingForwardResponseRequestJSON struct {
	Data        apijson.Field
	Headers     apijson.Field
	Method      apijson.Field
	Params      apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestForwardingForwardResponseRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestForwardingForwardResponseRequestJSON) RawJSON() string {
	return r.raw
}

type RequestForwardingForwardParams struct {
	// The HTTP method for the forwarded request. Valid values include: `GET` , `POST`
	// , `PUT` , `DELETE` , and `PATCH`.
	Method param.Field[string] `json:"method,required"`
	// The URL route path for the forwarded request. This value must begin with a
	// forward-slash ( / ) and may only contain alphanumeric characters, hyphens, and
	// underscores.
	Route param.Field[string] `json:"route,required"`
	// The body for the forwarded request. This value must be specified as either a
	// string or a valid JSON object.
	Data param.Field[string] `json:"data"`
	// The HTTP headers to include on the forwarded request. This value must be
	// specified as an object of key-value pairs. Example:
	// `{"Content-Type": "application/xml", "X-API-Version": "v1" }`
	Headers param.Field[interface{}] `json:"headers"`
	// The query parameters for the forwarded request. This value must be specified as
	// a valid JSON object rather than a query string.
	Params param.Field[interface{}] `json:"params"`
}

func (r RequestForwardingForwardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
