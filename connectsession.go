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

// ConnectSessionService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectSessionService] method instead.
type ConnectSessionService struct {
	Options []option.RequestOption
}

// NewConnectSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectSessionService(opts ...option.RequestOption) (r *ConnectSessionService) {
	r = &ConnectSessionService{}
	r.Options = opts
	return
}

// Create a new connect session for an employer
func (r *ConnectSessionService) Connect(ctx context.Context, body ConnectSessionConnectParams, opts ...option.RequestOption) (res *ConnectSessionConnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connect/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new Connect session for reauthenticating an existing connection
func (r *ConnectSessionService) Reauthenticate(ctx context.Context, body ConnectSessionReauthenticateParams, opts ...option.RequestOption) (res *ConnectSessionReauthenticateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "connect/sessions/reauthenticate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ConnectSessionConnectResponse struct {
	// The Connect URL to redirect the user to for authentication
	ConnectURL string `json:"connect_url,required" format:"uri"`
	// The unique identifier for the created connect session
	SessionID string                            `json:"session_id,required"`
	JSON      connectSessionConnectResponseJSON `json:"-"`
}

// connectSessionConnectResponseJSON contains the JSON metadata for the struct
// [ConnectSessionConnectResponse]
type connectSessionConnectResponseJSON struct {
	ConnectURL  apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectSessionConnectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSessionConnectResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectSessionReauthenticateResponse struct {
	// The Connect URL to redirect the user to for reauthentication
	ConnectURL string `json:"connect_url,required" format:"uri"`
	// The unique identifier for the created connect session
	SessionID string                                   `json:"session_id,required"`
	JSON      connectSessionReauthenticateResponseJSON `json:"-"`
}

// connectSessionReauthenticateResponseJSON contains the JSON metadata for the
// struct [ConnectSessionReauthenticateResponse]
type connectSessionReauthenticateResponseJSON struct {
	ConnectURL  apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectSessionReauthenticateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSessionReauthenticateResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectSessionConnectParams struct {
	// Unique identifier for the customer
	CustomerID param.Field[string] `json:"customer_id,required"`
	// Name of the customer
	CustomerName param.Field[string] `json:"customer_name,required"`
	// The Finch products to request access to
	Products param.Field[[]ConnectSessionConnectParamsProduct] `json:"products,required"`
	// Email address of the customer
	CustomerEmail param.Field[string] `json:"customer_email" format:"email"`
	// Integration configuration for the connect session
	Integration param.Field[ConnectSessionConnectParamsIntegration] `json:"integration"`
	// Enable manual authentication mode
	Manual param.Field[bool] `json:"manual"`
	// The number of minutes until the session expires (defaults to 129,600, which is
	// 90 days)
	MinutesToExpire param.Field[float64] `json:"minutes_to_expire"`
	// The URI to redirect to after the Connect flow is completed
	RedirectUri param.Field[string] `json:"redirect_uri"`
	// Sandbox mode for testing
	Sandbox param.Field[ConnectSessionConnectParamsSandbox] `json:"sandbox"`
}

func (r ConnectSessionConnectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Finch products that can be requested during the Connect flow.
type ConnectSessionConnectParamsProduct string

const (
	ConnectSessionConnectParamsProductBenefits     ConnectSessionConnectParamsProduct = "benefits"
	ConnectSessionConnectParamsProductCompany      ConnectSessionConnectParamsProduct = "company"
	ConnectSessionConnectParamsProductDeduction    ConnectSessionConnectParamsProduct = "deduction"
	ConnectSessionConnectParamsProductDirectory    ConnectSessionConnectParamsProduct = "directory"
	ConnectSessionConnectParamsProductDocuments    ConnectSessionConnectParamsProduct = "documents"
	ConnectSessionConnectParamsProductEmployment   ConnectSessionConnectParamsProduct = "employment"
	ConnectSessionConnectParamsProductIndividual   ConnectSessionConnectParamsProduct = "individual"
	ConnectSessionConnectParamsProductPayment      ConnectSessionConnectParamsProduct = "payment"
	ConnectSessionConnectParamsProductPayStatement ConnectSessionConnectParamsProduct = "pay_statement"
	ConnectSessionConnectParamsProductSsn          ConnectSessionConnectParamsProduct = "ssn"
)

func (r ConnectSessionConnectParamsProduct) IsKnown() bool {
	switch r {
	case ConnectSessionConnectParamsProductBenefits, ConnectSessionConnectParamsProductCompany, ConnectSessionConnectParamsProductDeduction, ConnectSessionConnectParamsProductDirectory, ConnectSessionConnectParamsProductDocuments, ConnectSessionConnectParamsProductEmployment, ConnectSessionConnectParamsProductIndividual, ConnectSessionConnectParamsProductPayment, ConnectSessionConnectParamsProductPayStatement, ConnectSessionConnectParamsProductSsn:
		return true
	}
	return false
}

// Integration configuration for the connect session
type ConnectSessionConnectParamsIntegration struct {
	// The provider to integrate with
	Provider param.Field[string] `json:"provider,required"`
	// The authentication method to use
	AuthMethod param.Field[ConnectSessionConnectParamsIntegrationAuthMethod] `json:"auth_method"`
}

func (r ConnectSessionConnectParamsIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The authentication method to use
type ConnectSessionConnectParamsIntegrationAuthMethod string

const (
	ConnectSessionConnectParamsIntegrationAuthMethodAssisted   ConnectSessionConnectParamsIntegrationAuthMethod = "assisted"
	ConnectSessionConnectParamsIntegrationAuthMethodCredential ConnectSessionConnectParamsIntegrationAuthMethod = "credential"
	ConnectSessionConnectParamsIntegrationAuthMethodOAuth      ConnectSessionConnectParamsIntegrationAuthMethod = "oauth"
	ConnectSessionConnectParamsIntegrationAuthMethodAPIToken   ConnectSessionConnectParamsIntegrationAuthMethod = "api_token"
)

func (r ConnectSessionConnectParamsIntegrationAuthMethod) IsKnown() bool {
	switch r {
	case ConnectSessionConnectParamsIntegrationAuthMethodAssisted, ConnectSessionConnectParamsIntegrationAuthMethodCredential, ConnectSessionConnectParamsIntegrationAuthMethodOAuth, ConnectSessionConnectParamsIntegrationAuthMethodAPIToken:
		return true
	}
	return false
}

// Sandbox mode for testing
type ConnectSessionConnectParamsSandbox string

const (
	ConnectSessionConnectParamsSandboxFinch    ConnectSessionConnectParamsSandbox = "finch"
	ConnectSessionConnectParamsSandboxProvider ConnectSessionConnectParamsSandbox = "provider"
)

func (r ConnectSessionConnectParamsSandbox) IsKnown() bool {
	switch r {
	case ConnectSessionConnectParamsSandboxFinch, ConnectSessionConnectParamsSandboxProvider:
		return true
	}
	return false
}

type ConnectSessionReauthenticateParams struct {
	// The ID of the existing connection to reauthenticate
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The number of minutes until the session expires (defaults to 43,200, which is 30
	// days)
	MinutesToExpire param.Field[int64] `json:"minutes_to_expire"`
	// The products to request access to (optional for reauthentication)
	Products param.Field[[]ConnectSessionReauthenticateParamsProduct] `json:"products"`
	// The URI to redirect to after the Connect flow is completed
	RedirectUri param.Field[string] `json:"redirect_uri" format:"uri"`
}

func (r ConnectSessionReauthenticateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Finch products that can be requested during the Connect flow.
type ConnectSessionReauthenticateParamsProduct string

const (
	ConnectSessionReauthenticateParamsProductBenefits     ConnectSessionReauthenticateParamsProduct = "benefits"
	ConnectSessionReauthenticateParamsProductCompany      ConnectSessionReauthenticateParamsProduct = "company"
	ConnectSessionReauthenticateParamsProductDeduction    ConnectSessionReauthenticateParamsProduct = "deduction"
	ConnectSessionReauthenticateParamsProductDirectory    ConnectSessionReauthenticateParamsProduct = "directory"
	ConnectSessionReauthenticateParamsProductDocuments    ConnectSessionReauthenticateParamsProduct = "documents"
	ConnectSessionReauthenticateParamsProductEmployment   ConnectSessionReauthenticateParamsProduct = "employment"
	ConnectSessionReauthenticateParamsProductIndividual   ConnectSessionReauthenticateParamsProduct = "individual"
	ConnectSessionReauthenticateParamsProductPayment      ConnectSessionReauthenticateParamsProduct = "payment"
	ConnectSessionReauthenticateParamsProductPayStatement ConnectSessionReauthenticateParamsProduct = "pay_statement"
	ConnectSessionReauthenticateParamsProductSsn          ConnectSessionReauthenticateParamsProduct = "ssn"
)

func (r ConnectSessionReauthenticateParamsProduct) IsKnown() bool {
	switch r {
	case ConnectSessionReauthenticateParamsProductBenefits, ConnectSessionReauthenticateParamsProductCompany, ConnectSessionReauthenticateParamsProductDeduction, ConnectSessionReauthenticateParamsProductDirectory, ConnectSessionReauthenticateParamsProductDocuments, ConnectSessionReauthenticateParamsProductEmployment, ConnectSessionReauthenticateParamsProductIndividual, ConnectSessionReauthenticateParamsProductPayment, ConnectSessionReauthenticateParamsProductPayStatement, ConnectSessionReauthenticateParamsProductSsn:
		return true
	}
	return false
}
