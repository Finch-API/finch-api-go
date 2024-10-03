// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/shared"
)

// AccountService contains methods and other services that help with interacting
// with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
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

// Disconnect one or more `access_token`s from your application.
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
	Status string                 `json:"status,required"`
	JSON   disconnectResponseJSON `json:"-"`
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

func (r disconnectResponseJSON) RawJSON() string {
	return r.raw
}

type Introspection struct {
	// [DEPRECATED] Use `connection_id` to associate tokens with a Finch connection
	// instead of this account ID.
	AccountID             string                              `json:"account_id,required"`
	AuthenticationMethods []IntrospectionAuthenticationMethod `json:"authentication_methods,required"`
	// The client ID of the application associated with the `access_token`.
	ClientID string `json:"client_id,required"`
	// The type of application associated with a token.
	ClientType IntrospectionClientType `json:"client_type,required"`
	// [DEPRECATED] Use `connection_id` to associate tokens with a Finch connection
	// instead of this company ID.
	CompanyID string `json:"company_id,required"`
	// The Finch UUID of the connection associated with the `access_token`.
	ConnectionID     string                        `json:"connection_id,required"`
	ConnectionStatus IntrospectionConnectionStatus `json:"connection_status,required"`
	// The type of the connection associated with the token.
	//
	// - `provider` - connection to an external provider
	// - `finch` - finch-generated data.
	ConnectionType IntrospectionConnectionType `json:"connection_type,required"`
	// The email of your customer you provided to Finch when a connect session was
	// created for this connection.
	CustomerEmail string `json:"customer_email,required,nullable"`
	// The ID of your customer you provided to Finch when a connect session was created
	// for this connection.
	CustomerID string `json:"customer_id,required,nullable"`
	// The name of your customer you provided to Finch when a connect session was
	// created for this connection.
	CustomerName string `json:"customer_name,required,nullable"`
	// Whether the connection associated with the `access_token` uses the Assisted
	// Connect Flow. (`true` if using Assisted Connect, `false` if connection is
	// automated)
	Manual bool `json:"manual,required"`
	// [DEPRECATED] Use `provider_id` to identify the provider instead of this payroll
	// provider ID.
	PayrollProviderID string `json:"payroll_provider_id,required"`
	// An array of the authorized products associated with the `access_token`.
	Products []string `json:"products,required"`
	// The ID of the provider associated with the `access_token`.
	ProviderID string `json:"provider_id,required"`
	// The account username used for login associated with the `access_token`.
	Username string            `json:"username,required"`
	JSON     introspectionJSON `json:"-"`
}

// introspectionJSON contains the JSON metadata for the struct [Introspection]
type introspectionJSON struct {
	AccountID             apijson.Field
	AuthenticationMethods apijson.Field
	ClientID              apijson.Field
	ClientType            apijson.Field
	CompanyID             apijson.Field
	ConnectionID          apijson.Field
	ConnectionStatus      apijson.Field
	ConnectionType        apijson.Field
	CustomerEmail         apijson.Field
	CustomerID            apijson.Field
	CustomerName          apijson.Field
	Manual                apijson.Field
	PayrollProviderID     apijson.Field
	Products              apijson.Field
	ProviderID            apijson.Field
	Username              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Introspection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r introspectionJSON) RawJSON() string {
	return r.raw
}

type IntrospectionAuthenticationMethod struct {
	ConnectionStatus IntrospectionAuthenticationMethodsConnectionStatus `json:"connection_status"`
	// An array of the authorized products associated with the `access_token`.
	Products []string `json:"products"`
	// The type of authentication method.
	Type IntrospectionAuthenticationMethodsType `json:"type"`
	JSON introspectionAuthenticationMethodJSON  `json:"-"`
}

// introspectionAuthenticationMethodJSON contains the JSON metadata for the struct
// [IntrospectionAuthenticationMethod]
type introspectionAuthenticationMethodJSON struct {
	ConnectionStatus apijson.Field
	Products         apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IntrospectionAuthenticationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r introspectionAuthenticationMethodJSON) RawJSON() string {
	return r.raw
}

type IntrospectionAuthenticationMethodsConnectionStatus struct {
	Message string                                                 `json:"message"`
	Status  shared.ConnectionStatusType                            `json:"status"`
	JSON    introspectionAuthenticationMethodsConnectionStatusJSON `json:"-"`
}

// introspectionAuthenticationMethodsConnectionStatusJSON contains the JSON
// metadata for the struct [IntrospectionAuthenticationMethodsConnectionStatus]
type introspectionAuthenticationMethodsConnectionStatusJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntrospectionAuthenticationMethodsConnectionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r introspectionAuthenticationMethodsConnectionStatusJSON) RawJSON() string {
	return r.raw
}

// The type of authentication method.
type IntrospectionAuthenticationMethodsType string

const (
	IntrospectionAuthenticationMethodsTypeAssisted      IntrospectionAuthenticationMethodsType = "assisted"
	IntrospectionAuthenticationMethodsTypeCredential    IntrospectionAuthenticationMethodsType = "credential"
	IntrospectionAuthenticationMethodsTypeAPIToken      IntrospectionAuthenticationMethodsType = "api_token"
	IntrospectionAuthenticationMethodsTypeAPICredential IntrospectionAuthenticationMethodsType = "api_credential"
	IntrospectionAuthenticationMethodsTypeOAuth         IntrospectionAuthenticationMethodsType = "oauth"
)

func (r IntrospectionAuthenticationMethodsType) IsKnown() bool {
	switch r {
	case IntrospectionAuthenticationMethodsTypeAssisted, IntrospectionAuthenticationMethodsTypeCredential, IntrospectionAuthenticationMethodsTypeAPIToken, IntrospectionAuthenticationMethodsTypeAPICredential, IntrospectionAuthenticationMethodsTypeOAuth:
		return true
	}
	return false
}

// The type of application associated with a token.
type IntrospectionClientType string

const (
	IntrospectionClientTypeProduction  IntrospectionClientType = "production"
	IntrospectionClientTypeDevelopment IntrospectionClientType = "development"
	IntrospectionClientTypeSandbox     IntrospectionClientType = "sandbox"
)

func (r IntrospectionClientType) IsKnown() bool {
	switch r {
	case IntrospectionClientTypeProduction, IntrospectionClientTypeDevelopment, IntrospectionClientTypeSandbox:
		return true
	}
	return false
}

type IntrospectionConnectionStatus struct {
	Message string                            `json:"message"`
	Status  shared.ConnectionStatusType       `json:"status"`
	JSON    introspectionConnectionStatusJSON `json:"-"`
}

// introspectionConnectionStatusJSON contains the JSON metadata for the struct
// [IntrospectionConnectionStatus]
type introspectionConnectionStatusJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntrospectionConnectionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r introspectionConnectionStatusJSON) RawJSON() string {
	return r.raw
}

// The type of the connection associated with the token.
//
// - `provider` - connection to an external provider
// - `finch` - finch-generated data.
type IntrospectionConnectionType string

const (
	IntrospectionConnectionTypeProvider IntrospectionConnectionType = "provider"
	IntrospectionConnectionTypeFinch    IntrospectionConnectionType = "finch"
)

func (r IntrospectionConnectionType) IsKnown() bool {
	switch r {
	case IntrospectionConnectionTypeProvider, IntrospectionConnectionTypeFinch:
		return true
	}
	return false
}
