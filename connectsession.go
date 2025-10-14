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
func (r *ConnectSessionService) New(ctx context.Context, body ConnectSessionNewParams, opts ...option.RequestOption) (res *ConnectSessionNewResponse, err error) {
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

type ConnectSessionNewResponse struct {
	// The Connect URL to redirect the user to for authentication
	ConnectURL string `json:"connect_url,required" format:"uri"`
	// The unique identifier for the created connect session
	SessionID string                        `json:"session_id,required"`
	JSON      connectSessionNewResponseJSON `json:"-"`
}

// connectSessionNewResponseJSON contains the JSON metadata for the struct
// [ConnectSessionNewResponse]
type connectSessionNewResponseJSON struct {
	ConnectURL  apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectSessionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectSessionNewResponseJSON) RawJSON() string {
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

type ConnectSessionNewParams struct {
	// Email address of the customer
	CustomerEmail param.Field[string] `json:"customer_email,required" format:"email"`
	// Unique identifier for the customer
	CustomerID param.Field[string] `json:"customer_id,required"`
	// Name of the customer
	CustomerName param.Field[string] `json:"customer_name,required"`
	// Integration configuration for the connect session
	Integration param.Field[ConnectSessionNewParamsIntegration] `json:"integration,required"`
	// Enable manual authentication mode
	Manual param.Field[bool] `json:"manual,required"`
	// The number of minutes until the session expires (defaults to 129,600, which is
	// 90 days)
	MinutesToExpire param.Field[float64] `json:"minutes_to_expire,required"`
	// The Finch products to request access to
	Products param.Field[[]ConnectSessionNewParamsProduct] `json:"products,required"`
	// The URI to redirect to after the Connect flow is completed
	RedirectUri param.Field[string] `json:"redirect_uri,required"`
	// Sandbox mode for testing
	Sandbox param.Field[ConnectSessionNewParamsSandbox] `json:"sandbox,required"`
}

func (r ConnectSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Integration configuration for the connect session
type ConnectSessionNewParamsIntegration struct {
	// The authentication method to use
	AuthMethod param.Field[ConnectSessionNewParamsIntegrationAuthMethod] `json:"auth_method,required"`
	// The provider to integrate with
	Provider param.Field[string] `json:"provider,required"`
}

func (r ConnectSessionNewParamsIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The authentication method to use
type ConnectSessionNewParamsIntegrationAuthMethod string

const (
	ConnectSessionNewParamsIntegrationAuthMethodAssisted   ConnectSessionNewParamsIntegrationAuthMethod = "assisted"
	ConnectSessionNewParamsIntegrationAuthMethodCredential ConnectSessionNewParamsIntegrationAuthMethod = "credential"
	ConnectSessionNewParamsIntegrationAuthMethodOAuth      ConnectSessionNewParamsIntegrationAuthMethod = "oauth"
	ConnectSessionNewParamsIntegrationAuthMethodAPIToken   ConnectSessionNewParamsIntegrationAuthMethod = "api_token"
)

func (r ConnectSessionNewParamsIntegrationAuthMethod) IsKnown() bool {
	switch r {
	case ConnectSessionNewParamsIntegrationAuthMethodAssisted, ConnectSessionNewParamsIntegrationAuthMethodCredential, ConnectSessionNewParamsIntegrationAuthMethodOAuth, ConnectSessionNewParamsIntegrationAuthMethodAPIToken:
		return true
	}
	return false
}

// The Finch products that can be requested during the Connect flow.
type ConnectSessionNewParamsProduct string

const (
	ConnectSessionNewParamsProductBenefits     ConnectSessionNewParamsProduct = "benefits"
	ConnectSessionNewParamsProductCompany      ConnectSessionNewParamsProduct = "company"
	ConnectSessionNewParamsProductDeduction    ConnectSessionNewParamsProduct = "deduction"
	ConnectSessionNewParamsProductDirectory    ConnectSessionNewParamsProduct = "directory"
	ConnectSessionNewParamsProductDocuments    ConnectSessionNewParamsProduct = "documents"
	ConnectSessionNewParamsProductEmployment   ConnectSessionNewParamsProduct = "employment"
	ConnectSessionNewParamsProductIndividual   ConnectSessionNewParamsProduct = "individual"
	ConnectSessionNewParamsProductPayment      ConnectSessionNewParamsProduct = "payment"
	ConnectSessionNewParamsProductPayStatement ConnectSessionNewParamsProduct = "pay_statement"
	ConnectSessionNewParamsProductSsn          ConnectSessionNewParamsProduct = "ssn"
)

func (r ConnectSessionNewParamsProduct) IsKnown() bool {
	switch r {
	case ConnectSessionNewParamsProductBenefits, ConnectSessionNewParamsProductCompany, ConnectSessionNewParamsProductDeduction, ConnectSessionNewParamsProductDirectory, ConnectSessionNewParamsProductDocuments, ConnectSessionNewParamsProductEmployment, ConnectSessionNewParamsProductIndividual, ConnectSessionNewParamsProductPayment, ConnectSessionNewParamsProductPayStatement, ConnectSessionNewParamsProductSsn:
		return true
	}
	return false
}

// Sandbox mode for testing
type ConnectSessionNewParamsSandbox string

const (
	ConnectSessionNewParamsSandboxFinch    ConnectSessionNewParamsSandbox = "finch"
	ConnectSessionNewParamsSandboxProvider ConnectSessionNewParamsSandbox = "provider"
)

func (r ConnectSessionNewParamsSandbox) IsKnown() bool {
	switch r {
	case ConnectSessionNewParamsSandboxFinch, ConnectSessionNewParamsSandboxProvider:
		return true
	}
	return false
}

type ConnectSessionReauthenticateParams struct {
	// The ID of the existing connection to reauthenticate
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The number of minutes until the session expires (defaults to 43,200, which is 30
	// days)
	MinutesToExpire param.Field[int64] `json:"minutes_to_expire,required"`
	// The products to request access to (optional for reauthentication)
	Products param.Field[[]ConnectSessionReauthenticateParamsProduct] `json:"products,required"`
	// The URI to redirect to after the Connect flow is completed
	RedirectUri param.Field[string] `json:"redirect_uri,required" format:"uri"`
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
