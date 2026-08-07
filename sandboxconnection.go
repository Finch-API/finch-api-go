// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/v2/internal/apijson"
	"github.com/Finch-API/finch-api-go/v2/internal/param"
	"github.com/Finch-API/finch-api-go/v2/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/v2/option"
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
	var preClientOpts = []option.RequestOption{requestconfig.WithBasicAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "sandbox/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SandboxConnectionNewResponse struct {
	AccessToken        string                                         `json:"access_token" api:"required"`
	AccountID          string                                         `json:"account_id" api:"required" format:"uuid"`
	AuthenticationType SandboxConnectionNewResponseAuthenticationType `json:"authentication_type" api:"required"`
	ConnectionID       string                                         `json:"connection_id" api:"required"`
	EntityID           string                                         `json:"entity_id" api:"required"`
	Products           []string                                       `json:"products" api:"required"`
	ProviderID         string                                         `json:"provider_id" api:"required"`
	TokenType          string                                         `json:"token_type" api:"required"`
	CompanyID          string                                         `json:"company_id" api:"nullable" format:"uuid"`
	JSON               sandboxConnectionNewResponseJSON               `json:"-"`
}

// sandboxConnectionNewResponseJSON contains the JSON metadata for the struct
// [SandboxConnectionNewResponse]
type sandboxConnectionNewResponseJSON struct {
	AccessToken        apijson.Field
	AccountID          apijson.Field
	AuthenticationType apijson.Field
	ConnectionID       apijson.Field
	EntityID           apijson.Field
	Products           apijson.Field
	ProviderID         apijson.Field
	TokenType          apijson.Field
	CompanyID          apijson.Field
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
	SandboxConnectionNewResponseAuthenticationTypeAPIToken   SandboxConnectionNewResponseAuthenticationType = "api_token"
	SandboxConnectionNewResponseAuthenticationTypeAssisted   SandboxConnectionNewResponseAuthenticationType = "assisted"
	SandboxConnectionNewResponseAuthenticationTypeCredential SandboxConnectionNewResponseAuthenticationType = "credential"
	SandboxConnectionNewResponseAuthenticationTypeOAuth      SandboxConnectionNewResponseAuthenticationType = "oauth"
)

func (r SandboxConnectionNewResponseAuthenticationType) IsKnown() bool {
	switch r {
	case SandboxConnectionNewResponseAuthenticationTypeAPIToken, SandboxConnectionNewResponseAuthenticationTypeAssisted, SandboxConnectionNewResponseAuthenticationTypeCredential, SandboxConnectionNewResponseAuthenticationTypeOAuth:
		return true
	}
	return false
}

type SandboxConnectionNewParams struct {
	// The provider associated with the connection
	ProviderID         param.Field[string]                                       `json:"provider_id" api:"required"`
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
	SandboxConnectionNewParamsAuthenticationTypeAPIToken   SandboxConnectionNewParamsAuthenticationType = "api_token"
	SandboxConnectionNewParamsAuthenticationTypeAssisted   SandboxConnectionNewParamsAuthenticationType = "assisted"
	SandboxConnectionNewParamsAuthenticationTypeCredential SandboxConnectionNewParamsAuthenticationType = "credential"
	SandboxConnectionNewParamsAuthenticationTypeOAuth      SandboxConnectionNewParamsAuthenticationType = "oauth"
)

func (r SandboxConnectionNewParamsAuthenticationType) IsKnown() bool {
	switch r {
	case SandboxConnectionNewParamsAuthenticationTypeAPIToken, SandboxConnectionNewParamsAuthenticationTypeAssisted, SandboxConnectionNewParamsAuthenticationTypeCredential, SandboxConnectionNewParamsAuthenticationTypeOAuth:
		return true
	}
	return false
}
