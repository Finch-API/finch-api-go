// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxConnectionService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxConnectionService] method instead.
type SandboxConnectionService struct {
	Options  []option.RequestOption
	Accounts *SandboxConnectionAccountService
}

// NewSandboxConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxConnectionService(opts ...option.RequestOption) (r *SandboxConnectionService) {
	r = &SandboxConnectionService{}
	r.Options = opts
	r.Accounts = NewSandboxConnectionAccountService(opts...)
	return
}

// Create a new connection (new company/provider pair) with a new account
func (r *SandboxConnectionService) New(ctx context.Context, body SandboxConnectionNewParams, opts ...option.RequestOption) (res *SandboxConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandbox/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SandboxConnectionNewResponse struct {
	AccessToken string `json:"access_token,required" format:"uuid"`
	// [DEPRECATED] Use `connection_id` to associate a connection with an access token
	//
	// Deprecated: deprecated
	AccountID          string                                         `json:"account_id,required" format:"uuid"`
	AuthenticationType SandboxConnectionNewResponseAuthenticationType `json:"authentication_type,required"`
	// [DEPRECATED] Use `connection_id` to associate a connection with an access token
	//
	// Deprecated: deprecated
	CompanyID string `json:"company_id,required" format:"uuid"`
	// The ID of the new connection
	ConnectionID string   `json:"connection_id,required" format:"uuid"`
	Products     []string `json:"products,required"`
	// The ID of the provider associated with the `access_token`.
	ProviderID string                           `json:"provider_id,required" format:"uuid"`
	TokenType  string                           `json:"token_type"`
	JSON       sandboxConnectionNewResponseJSON `json:"-"`
}

// sandboxConnectionNewResponseJSON contains the JSON metadata for the struct
// [SandboxConnectionNewResponse]
type sandboxConnectionNewResponseJSON struct {
	AccessToken        apijson.Field
	AccountID          apijson.Field
	AuthenticationType apijson.Field
	CompanyID          apijson.Field
	ConnectionID       apijson.Field
	Products           apijson.Field
	ProviderID         apijson.Field
	TokenType          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SandboxConnectionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxConnectionNewResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxConnectionNewResponseAuthenticationType string

const (
	SandboxConnectionNewResponseAuthenticationTypeCredential SandboxConnectionNewResponseAuthenticationType = "credential"
	SandboxConnectionNewResponseAuthenticationTypeAPIToken   SandboxConnectionNewResponseAuthenticationType = "api_token"
	SandboxConnectionNewResponseAuthenticationTypeOAuth      SandboxConnectionNewResponseAuthenticationType = "oauth"
	SandboxConnectionNewResponseAuthenticationTypeAssisted   SandboxConnectionNewResponseAuthenticationType = "assisted"
)

func (r SandboxConnectionNewResponseAuthenticationType) IsKnown() bool {
	switch r {
	case SandboxConnectionNewResponseAuthenticationTypeCredential, SandboxConnectionNewResponseAuthenticationTypeAPIToken, SandboxConnectionNewResponseAuthenticationTypeOAuth, SandboxConnectionNewResponseAuthenticationTypeAssisted:
		return true
	}
	return false
}

type SandboxConnectionNewParams struct {
	// The provider associated with the connection
	ProviderID         param.Field[string]                                       `json:"provider_id,required"`
	AuthenticationType param.Field[SandboxConnectionNewParamsAuthenticationType] `json:"authentication_type"`
	// Optional: the size of the employer to be created with this connection. Defaults
	// to 20. Note that if this is higher than 100, historical payroll data will not be
	// generated, and instead only one pay period will be created.
	EmployeeSize param.Field[int64]    `json:"employee_size"`
	Products     param.Field[[]string] `json:"products"`
}

func (r SandboxConnectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxConnectionNewParamsAuthenticationType string

const (
	SandboxConnectionNewParamsAuthenticationTypeCredential SandboxConnectionNewParamsAuthenticationType = "credential"
	SandboxConnectionNewParamsAuthenticationTypeAPIToken   SandboxConnectionNewParamsAuthenticationType = "api_token"
	SandboxConnectionNewParamsAuthenticationTypeOAuth      SandboxConnectionNewParamsAuthenticationType = "oauth"
	SandboxConnectionNewParamsAuthenticationTypeAssisted   SandboxConnectionNewParamsAuthenticationType = "assisted"
)

func (r SandboxConnectionNewParamsAuthenticationType) IsKnown() bool {
	switch r {
	case SandboxConnectionNewParamsAuthenticationTypeCredential, SandboxConnectionNewParamsAuthenticationTypeAPIToken, SandboxConnectionNewParamsAuthenticationTypeOAuth, SandboxConnectionNewParamsAuthenticationTypeAssisted:
		return true
	}
	return false
}
