package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Disconnect an employer from your application and invalidate all `access_token`s
// associated with the employer. We require applications to implement the
// Disconnect endpoint for billing and security purposes.
func (r *AccountService) Disconnect(ctx context.Context, opts ...option.RequestOption) (res *DisconnectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "disconnect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Read account information associated with an `access_token`
func (r *AccountService) Introspect(ctx context.Context, opts ...option.RequestOption) (res *Introspection, err error) {
	opts = append(r.Options[:], opts...)
	path := "introspect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DisconnectResponse struct {
	// If the request is successful, Finch will return “success” (HTTP 200 status).
	Status string `json:"status,required"`
	JSON   disconnectResponseJSON
}

// disconnectResponseJSON contains the JSON metadata for the struct
// [DisconnectResponse]
type disconnectResponseJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisconnectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Introspection struct {
	// The client id of the application associated with the `access_token`.
	ClientID string `json:"client_id,required"`
	// The Finch uuid of the company associated with the `access_token`.
	CompanyID string `json:"company_id,required"`
	// An array of the authorized products associated with the `access_token`.
	Products []string `json:"products,required"`
	// The account username used for login associated with the `access_token`.
	Username string `json:"username,required"`
	// The payroll provider associated with the `access_token`.
	PayrollProviderID string `json:"payroll_provider_id,required"`
	// Whether the connection associated with the `access_token` uses the Assisted
	// Connect Flow. (`true` if using Assisted Connect, `false` if connection is
	// automated)
	Manual bool `json:"manual,required"`
	JSON   introspectionJSON
}

// introspectionJSON contains the JSON metadata for the struct [Introspection]
type introspectionJSON struct {
	ClientID          apijson.Field
	CompanyID         apijson.Field
	Products          apijson.Field
	Username          apijson.Field
	PayrollProviderID apijson.Field
	Manual            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Introspection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
