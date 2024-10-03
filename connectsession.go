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
	opts = append(r.Options[:], opts...)
	path := "connect/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new Connect session for reauthenticating an existing connection
func (r *ConnectSessionService) Reauthenticate(ctx context.Context, body ConnectSessionReauthenticateParams, opts ...option.RequestOption) (res *ConnectSessionReauthenticateResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	CustomerID    param.Field[string]                             `json:"customer_id,required"`
	CustomerName  param.Field[string]                             `json:"customer_name,required"`
	Products      param.Field[[]ConnectSessionNewParamsProduct]   `json:"products,required"`
	CustomerEmail param.Field[string]                             `json:"customer_email" format:"email"`
	Integration   param.Field[ConnectSessionNewParamsIntegration] `json:"integration"`
	Manual        param.Field[bool]                               `json:"manual"`
	// The number of minutes until the session expires (defaults to 20,160, which is 14
	// days)
	MinutesToExpire param.Field[float64]                        `json:"minutes_to_expire"`
	RedirectUri     param.Field[string]                         `json:"redirect_uri"`
	Sandbox         param.Field[ConnectSessionNewParamsSandbox] `json:"sandbox"`
}

func (r ConnectSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Finch products that can be requested during the Connect flow.
type ConnectSessionNewParamsProduct string

const (
	ConnectSessionNewParamsProductCompany      ConnectSessionNewParamsProduct = "company"
	ConnectSessionNewParamsProductDirectory    ConnectSessionNewParamsProduct = "directory"
	ConnectSessionNewParamsProductIndividual   ConnectSessionNewParamsProduct = "individual"
	ConnectSessionNewParamsProductEmployment   ConnectSessionNewParamsProduct = "employment"
	ConnectSessionNewParamsProductPayment      ConnectSessionNewParamsProduct = "payment"
	ConnectSessionNewParamsProductPayStatement ConnectSessionNewParamsProduct = "pay_statement"
	ConnectSessionNewParamsProductBenefits     ConnectSessionNewParamsProduct = "benefits"
	ConnectSessionNewParamsProductSsn          ConnectSessionNewParamsProduct = "ssn"
)

func (r ConnectSessionNewParamsProduct) IsKnown() bool {
	switch r {
	case ConnectSessionNewParamsProductCompany, ConnectSessionNewParamsProductDirectory, ConnectSessionNewParamsProductIndividual, ConnectSessionNewParamsProductEmployment, ConnectSessionNewParamsProductPayment, ConnectSessionNewParamsProductPayStatement, ConnectSessionNewParamsProductBenefits, ConnectSessionNewParamsProductSsn:
		return true
	}
	return false
}

type ConnectSessionNewParamsIntegration struct {
	AuthMethod param.Field[ConnectSessionNewParamsIntegrationAuthMethod] `json:"auth_method"`
	Provider   param.Field[string]                                       `json:"provider"`
}

func (r ConnectSessionNewParamsIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// The number of minutes until the session expires (defaults to 20,160, which is 14
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
	ConnectSessionReauthenticateParamsProductCompany      ConnectSessionReauthenticateParamsProduct = "company"
	ConnectSessionReauthenticateParamsProductDirectory    ConnectSessionReauthenticateParamsProduct = "directory"
	ConnectSessionReauthenticateParamsProductIndividual   ConnectSessionReauthenticateParamsProduct = "individual"
	ConnectSessionReauthenticateParamsProductEmployment   ConnectSessionReauthenticateParamsProduct = "employment"
	ConnectSessionReauthenticateParamsProductPayment      ConnectSessionReauthenticateParamsProduct = "payment"
	ConnectSessionReauthenticateParamsProductPayStatement ConnectSessionReauthenticateParamsProduct = "pay_statement"
	ConnectSessionReauthenticateParamsProductBenefits     ConnectSessionReauthenticateParamsProduct = "benefits"
	ConnectSessionReauthenticateParamsProductSsn          ConnectSessionReauthenticateParamsProduct = "ssn"
)

func (r ConnectSessionReauthenticateParamsProduct) IsKnown() bool {
	switch r {
	case ConnectSessionReauthenticateParamsProductCompany, ConnectSessionReauthenticateParamsProductDirectory, ConnectSessionReauthenticateParamsProductIndividual, ConnectSessionReauthenticateParamsProductEmployment, ConnectSessionReauthenticateParamsProductPayment, ConnectSessionReauthenticateParamsProductPayStatement, ConnectSessionReauthenticateParamsProductBenefits, ConnectSessionReauthenticateParamsProductSsn:
		return true
	}
	return false
}
