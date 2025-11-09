// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/shared"
	"github.com/tidwall/gjson"
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
	opts = slices.Concat(r.Options, opts)
	path := "forward"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RequestForwardingForwardResponse struct {
	// An object containing details of your original forwarded request, for your ease
	// of reference.
	Request RequestForwardingForwardResponseRequest `json:"request,required"`
	// The HTTP status code of the forwarded request's response, exactly received from
	// the underlying integration's API. This value will be returned as an integer.
	StatusCode int64 `json:"statusCode,required"`
	// A string representation of the HTTP response body of the forwarded request's
	// response received from the underlying integration's API. This field may be null
	// in the case where the upstream system's response is empty.
	Data string `json:"data,nullable"`
	// The HTTP headers of the forwarded request's response, exactly as received from
	// the underlying integration's API.
	Headers map[string]interface{}               `json:"headers,nullable"`
	JSON    requestForwardingForwardResponseJSON `json:"-"`
}

// requestForwardingForwardResponseJSON contains the JSON metadata for the struct
// [RequestForwardingForwardResponse]
type requestForwardingForwardResponseJSON struct {
	Request     apijson.Field
	StatusCode  apijson.Field
	Data        apijson.Field
	Headers     apijson.Field
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
	// The HTTP method that was specified for the forwarded request. Valid values
	// include: `GET` , `POST` , `PUT` , `DELETE` , and `PATCH`.
	Method string `json:"method,required"`
	// The URL route path that was specified for the forwarded request.
	Route string `json:"route,required"`
	// The body that was specified for the forwarded request.
	Data RequestForwardingForwardResponseRequestDataUnion `json:"data,nullable"`
	// The HTTP headers that were specified for the forwarded request.
	Headers map[string]string `json:"headers,nullable"`
	// The query parameters that were specified for the forwarded request.
	Params map[string]interface{}                      `json:"params,nullable"`
	JSON   requestForwardingForwardResponseRequestJSON `json:"-"`
}

// requestForwardingForwardResponseRequestJSON contains the JSON metadata for the
// struct [RequestForwardingForwardResponseRequest]
type requestForwardingForwardResponseRequestJSON struct {
	Method      apijson.Field
	Route       apijson.Field
	Data        apijson.Field
	Headers     apijson.Field
	Params      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestForwardingForwardResponseRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestForwardingForwardResponseRequestJSON) RawJSON() string {
	return r.raw
}

// The body that was specified for the forwarded request.
//
// Union satisfied by [shared.UnionString] or
// [RequestForwardingForwardResponseRequestDataMap].
type RequestForwardingForwardResponseRequestDataUnion interface {
	ImplementsRequestForwardingForwardResponseRequestDataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RequestForwardingForwardResponseRequestDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RequestForwardingForwardResponseRequestDataMap{}),
		},
	)
}

type RequestForwardingForwardResponseRequestDataMap map[string]interface{}

func (r RequestForwardingForwardResponseRequestDataMap) ImplementsRequestForwardingForwardResponseRequestDataUnion() {
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
	Headers param.Field[map[string]interface{}] `json:"headers"`
	// The query parameters for the forwarded request. This value must be specified as
	// a valid JSON object rather than a query string.
	Params param.Field[map[string]interface{}] `json:"params"`
}

func (r RequestForwardingForwardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
