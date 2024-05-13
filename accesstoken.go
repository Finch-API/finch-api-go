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

// AccessTokenService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessTokenService] method instead.
type AccessTokenService struct {
	Options []option.RequestOption
}

// NewAccessTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessTokenService(opts ...option.RequestOption) (r *AccessTokenService) {
	r = &AccessTokenService{}
	r.Options = opts
	return
}

// Exchange the authorization code for an access token
func (r *AccessTokenService) New(ctx context.Context, body AccessTokenNewParams, opts ...option.RequestOption) (res *CreateAccessTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateAccessTokenResponse struct {
	AccessToken string `json:"access_token,required"`
	// The Finch uuid of the account used to connect this company.
	AccountID string `json:"account_id,required"`
	// The type of application associated with a token.
	ClientType CreateAccessTokenResponseClientType `json:"client_type,required"`
	// The Finch uuid of the company associated with the `access_token`.
	CompanyID string `json:"company_id,required"`
	// The type of the connection associated with the token.
	//
	// - `provider` - connection to an external provider
	// - `finch` - finch-generated data.
	ConnectionType CreateAccessTokenResponseConnectionType `json:"connection_type,required"`
	// An array of the authorized products associated with the `access_token`.
	Products []string `json:"products,required"`
	// The payroll provider associated with the `access_token`.
	ProviderID string                        `json:"provider_id,required"`
	JSON       createAccessTokenResponseJSON `json:"-"`
}

// createAccessTokenResponseJSON contains the JSON metadata for the struct
// [CreateAccessTokenResponse]
type createAccessTokenResponseJSON struct {
	AccessToken    apijson.Field
	AccountID      apijson.Field
	ClientType     apijson.Field
	CompanyID      apijson.Field
	ConnectionType apijson.Field
	Products       apijson.Field
	ProviderID     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreateAccessTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createAccessTokenResponseJSON) RawJSON() string {
	return r.raw
}

// The type of application associated with a token.
type CreateAccessTokenResponseClientType string

const (
	CreateAccessTokenResponseClientTypeProduction  CreateAccessTokenResponseClientType = "production"
	CreateAccessTokenResponseClientTypeDevelopment CreateAccessTokenResponseClientType = "development"
	CreateAccessTokenResponseClientTypeSandbox     CreateAccessTokenResponseClientType = "sandbox"
)

func (r CreateAccessTokenResponseClientType) IsKnown() bool {
	switch r {
	case CreateAccessTokenResponseClientTypeProduction, CreateAccessTokenResponseClientTypeDevelopment, CreateAccessTokenResponseClientTypeSandbox:
		return true
	}
	return false
}

// The type of the connection associated with the token.
//
// - `provider` - connection to an external provider
// - `finch` - finch-generated data.
type CreateAccessTokenResponseConnectionType string

const (
	CreateAccessTokenResponseConnectionTypeProvider CreateAccessTokenResponseConnectionType = "provider"
	CreateAccessTokenResponseConnectionTypeFinch    CreateAccessTokenResponseConnectionType = "finch"
)

func (r CreateAccessTokenResponseConnectionType) IsKnown() bool {
	switch r {
	case CreateAccessTokenResponseConnectionTypeProvider, CreateAccessTokenResponseConnectionTypeFinch:
		return true
	}
	return false
}

type AccessTokenNewParams struct {
	Code         param.Field[string] `json:"code,required"`
	ClientID     param.Field[string] `json:"client_id" format:"uuid"`
	ClientSecret param.Field[string] `json:"client_secret"`
	RedirectUri  param.Field[string] `json:"redirect_uri"`
}

func (r AccessTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
