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
	// The access token for the connection
	AccessToken string `json:"access_token,required"`
	// The type of application associated with a token.
	ClientType CreateAccessTokenResponseClientType `json:"client_type,required"`
	// The Finch UUID of the connection associated with the `access_token`
	ConnectionID string `json:"connection_id,required"`
	// The type of the connection associated with the token.
	//
	// - `provider` - connection to an external provider
	// - `finch` - finch-generated data.
	ConnectionType CreateAccessTokenResponseConnectionType `json:"connection_type,required"`
	// An array of the authorized products associated with the `access_token`
	Products []string `json:"products,required"`
	// The ID of the provider associated with the `access_token`
	ProviderID string `json:"provider_id,required"`
	// The RFC 8693 token type (Finch uses `bearer` tokens)
	TokenType string `json:"token_type,required"`
	// [DEPRECATED] Use `connection_id` to identify the connection instead of this
	// account ID
	//
	// Deprecated: deprecated
	AccountID string `json:"account_id"`
	// [DEPRECATED] Use `connection_id` to identify the connection instead of this
	// company ID
	//
	// Deprecated: deprecated
	CompanyID string `json:"company_id"`
	// The ID of your customer you provided to Finch when a connect session was created
	// for this connection
	CustomerID string                        `json:"customer_id,nullable"`
	JSON       createAccessTokenResponseJSON `json:"-"`
}

// createAccessTokenResponseJSON contains the JSON metadata for the struct
// [CreateAccessTokenResponse]
type createAccessTokenResponseJSON struct {
	AccessToken    apijson.Field
	ClientType     apijson.Field
	ConnectionID   apijson.Field
	ConnectionType apijson.Field
	Products       apijson.Field
	ProviderID     apijson.Field
	TokenType      apijson.Field
	AccountID      apijson.Field
	CompanyID      apijson.Field
	CustomerID     apijson.Field
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
	CreateAccessTokenResponseClientTypeDevelopment CreateAccessTokenResponseClientType = "development"
	CreateAccessTokenResponseClientTypeProduction  CreateAccessTokenResponseClientType = "production"
	CreateAccessTokenResponseClientTypeSandbox     CreateAccessTokenResponseClientType = "sandbox"
)

func (r CreateAccessTokenResponseClientType) IsKnown() bool {
	switch r {
	case CreateAccessTokenResponseClientTypeDevelopment, CreateAccessTokenResponseClientTypeProduction, CreateAccessTokenResponseClientTypeSandbox:
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
	CreateAccessTokenResponseConnectionTypeFinch    CreateAccessTokenResponseConnectionType = "finch"
	CreateAccessTokenResponseConnectionTypeProvider CreateAccessTokenResponseConnectionType = "provider"
)

func (r CreateAccessTokenResponseConnectionType) IsKnown() bool {
	switch r {
	case CreateAccessTokenResponseConnectionTypeFinch, CreateAccessTokenResponseConnectionTypeProvider:
		return true
	}
	return false
}

type AccessTokenNewParams struct {
	// The client ID for your application
	ClientID param.Field[string] `json:"client_id,required" format:"uuid"`
	// The client secret for your application
	ClientSecret param.Field[string] `json:"client_secret,required"`
	// The authorization code received from the authorization server
	Code param.Field[string] `json:"code,required"`
	// The redirect URI used in the authorization request (optional)
	RedirectUri param.Field[string] `json:"redirect_uri"`
}

func (r AccessTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
