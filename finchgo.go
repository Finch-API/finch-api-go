// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
)

type FinchgoForwardResponse struct {
	// A string representation of the HTTP response body of the forwarded request’s
	// response received from the underlying integration’s API. This field may be null
	// in the case where the upstream system’s response is empty.
	Data string `json:"data,required,nullable"`
	// The HTTP headers of the forwarded request’s response, exactly as received from
	// the underlying integration’s API.
	Headers interface{} `json:"headers,required,nullable"`
	// An object containing details of your original forwarded request, for your ease
	// of reference.
	Request FinchgoForwardResponseRequest `json:"request,required"`
	// The HTTP status code of the forwarded request’s response, exactly received from
	// the underlying integration’s API. This value will be returned as an integer.
	StatusCode int64 `json:"statusCode,required"`
	JSON       finchgoForwardResponseJSON
}

// finchgoForwardResponseJSON contains the JSON metadata for the struct
// [FinchgoForwardResponse]
type finchgoForwardResponseJSON struct {
	Data        apijson.Field
	Headers     apijson.Field
	Request     apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinchgoForwardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing details of your original forwarded request, for your ease
// of reference.
type FinchgoForwardResponseRequest struct {
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
	Route string `json:"route,required"`
	JSON  finchgoForwardResponseRequestJSON
}

// finchgoForwardResponseRequestJSON contains the JSON metadata for the struct
// [FinchgoForwardResponseRequest]
type finchgoForwardResponseRequestJSON struct {
	Data        apijson.Field
	Headers     apijson.Field
	Method      apijson.Field
	Params      apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinchgoForwardResponseRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FinchgoForwardParams struct {
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

func (r FinchgoForwardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
