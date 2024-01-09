// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxConnectionAccountService contains methods and other services that help
// with interacting with the Finch API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSandboxConnectionAccountService]
// method instead.
type SandboxConnectionAccountService struct {
	Options []option.RequestOption
}

// NewSandboxConnectionAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSandboxConnectionAccountService(opts ...option.RequestOption) (r *SandboxConnectionAccountService) {
	r = &SandboxConnectionAccountService{}
	r.Options = opts
	return
}

// Create a new account for an existing connection (company/provider pair)
func (r *SandboxConnectionAccountService) New(ctx context.Context, body SandboxConnectionAccountNewParams, opts ...option.RequestOption) (res *SandboxConnectionAccountNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sandbox/connections/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing sandbox account. Change the connection status to understand
// how the Finch API responds.
func (r *SandboxConnectionAccountService) Update(ctx context.Context, body SandboxConnectionAccountUpdateParams, opts ...option.RequestOption) (res *SandboxConnectionAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sandbox/connections/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SandboxConnectionAccountNewResponse struct {
	AccessToken        string                                                `json:"access_token,required" format:"uuid"`
	AccountID          string                                                `json:"account_id,required" format:"uuid"`
	AuthenticationType SandboxConnectionAccountNewResponseAuthenticationType `json:"authentication_type,required"`
	CompanyID          string                                                `json:"company_id,required" format:"uuid"`
	Products           []string                                              `json:"products,required"`
	ProviderID         string                                                `json:"provider_id,required"`
	JSON               sandboxConnectionAccountNewResponseJSON               `json:"-"`
}

// sandboxConnectionAccountNewResponseJSON contains the JSON metadata for the
// struct [SandboxConnectionAccountNewResponse]
type sandboxConnectionAccountNewResponseJSON struct {
	AccessToken        apijson.Field
	AccountID          apijson.Field
	AuthenticationType apijson.Field
	CompanyID          apijson.Field
	Products           apijson.Field
	ProviderID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SandboxConnectionAccountNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxConnectionAccountNewResponseAuthenticationType string

const (
	SandboxConnectionAccountNewResponseAuthenticationTypeCredentials SandboxConnectionAccountNewResponseAuthenticationType = "credentials"
	SandboxConnectionAccountNewResponseAuthenticationTypeAPIToken    SandboxConnectionAccountNewResponseAuthenticationType = "api_token"
	SandboxConnectionAccountNewResponseAuthenticationTypeOauth       SandboxConnectionAccountNewResponseAuthenticationType = "oauth"
	SandboxConnectionAccountNewResponseAuthenticationTypeAssisted    SandboxConnectionAccountNewResponseAuthenticationType = "assisted"
)

type SandboxConnectionAccountUpdateResponse struct {
	AccountID          string                                                   `json:"account_id,required"`
	AuthenticationType SandboxConnectionAccountUpdateResponseAuthenticationType `json:"authentication_type,required"`
	CompanyID          string                                                   `json:"company_id,required" format:"uuid"`
	Products           []string                                                 `json:"products,required"`
	ProviderID         string                                                   `json:"provider_id,required"`
	JSON               sandboxConnectionAccountUpdateResponseJSON               `json:"-"`
}

// sandboxConnectionAccountUpdateResponseJSON contains the JSON metadata for the
// struct [SandboxConnectionAccountUpdateResponse]
type sandboxConnectionAccountUpdateResponseJSON struct {
	AccountID          apijson.Field
	AuthenticationType apijson.Field
	CompanyID          apijson.Field
	Products           apijson.Field
	ProviderID         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SandboxConnectionAccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxConnectionAccountUpdateResponseAuthenticationType string

const (
	SandboxConnectionAccountUpdateResponseAuthenticationTypeCredentials SandboxConnectionAccountUpdateResponseAuthenticationType = "credentials"
	SandboxConnectionAccountUpdateResponseAuthenticationTypeAPIToken    SandboxConnectionAccountUpdateResponseAuthenticationType = "api_token"
	SandboxConnectionAccountUpdateResponseAuthenticationTypeOauth       SandboxConnectionAccountUpdateResponseAuthenticationType = "oauth"
	SandboxConnectionAccountUpdateResponseAuthenticationTypeAssisted    SandboxConnectionAccountUpdateResponseAuthenticationType = "assisted"
)

type SandboxConnectionAccountNewParams struct {
	CompanyID          param.Field[string]                                              `json:"company_id,required" format:"uuid"`
	ProviderID         param.Field[string]                                              `json:"provider_id,required"`
	AuthenticationType param.Field[SandboxConnectionAccountNewParamsAuthenticationType] `json:"authentication_type"`
	// Optional, defaults to Organization products (`company`, `directory`,
	// `employment`, `individual`)
	Products param.Field[[]string] `json:"products"`
}

func (r SandboxConnectionAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxConnectionAccountNewParamsAuthenticationType string

const (
	SandboxConnectionAccountNewParamsAuthenticationTypeCredentials SandboxConnectionAccountNewParamsAuthenticationType = "credentials"
	SandboxConnectionAccountNewParamsAuthenticationTypeAPIToken    SandboxConnectionAccountNewParamsAuthenticationType = "api_token"
	SandboxConnectionAccountNewParamsAuthenticationTypeOauth       SandboxConnectionAccountNewParamsAuthenticationType = "oauth"
	SandboxConnectionAccountNewParamsAuthenticationTypeAssisted    SandboxConnectionAccountNewParamsAuthenticationType = "assisted"
)

type SandboxConnectionAccountUpdateParams struct {
	ConnectionStatus param.Field[shared.ConnectionStatusType] `json:"connection_status"`
}

func (r SandboxConnectionAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
