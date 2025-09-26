// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
)

// ProviderService contains methods and other services that help with interacting
// with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r *ProviderService) {
	r = &ProviderService{}
	r.Options = opts
	return
}

// Return details on all available payroll and HR systems.
func (r *ProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[ProviderListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "providers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Return details on all available payroll and HR systems.
func (r *ProviderService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProviderListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type ProviderListResponse struct {
	// The id of the payroll provider used in Connect.
	ID string `json:"id,required"`
	// The display name of the payroll provider.
	DisplayName string `json:"display_name,required"`
	// The list of Finch products supported on this payroll provider.
	Products []string `json:"products,required"`
	// The authentication methods supported by the provider.
	AuthenticationMethods []ProviderListResponseAuthenticationMethod `json:"authentication_methods"`
	// `true` if the integration is in a beta state, `false` otherwise
	Beta bool `json:"beta"`
	// The url to the official icon of the payroll provider.
	Icon string `json:"icon"`
	// The url to the official logo of the payroll provider.
	Logo string `json:"logo"`
	// [DEPRECATED] Whether the Finch integration with this provider uses the Assisted
	// Connect Flow by default. This field is now deprecated. Please check for a `type`
	// of `assisted` in the `authentication_methods` field instead.
	//
	// Deprecated: deprecated
	Manual bool `json:"manual"`
	// whether MFA is required for the provider.
	MfaRequired bool `json:"mfa_required"`
	// The hex code for the primary color of the payroll provider.
	PrimaryColor string                   `json:"primary_color"`
	JSON         providerListResponseJSON `json:"-"`
}

// providerListResponseJSON contains the JSON metadata for the struct
// [ProviderListResponse]
type providerListResponseJSON struct {
	ID                    apijson.Field
	DisplayName           apijson.Field
	Products              apijson.Field
	AuthenticationMethods apijson.Field
	Beta                  apijson.Field
	Icon                  apijson.Field
	Logo                  apijson.Field
	Manual                apijson.Field
	MfaRequired           apijson.Field
	PrimaryColor          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethod struct {
	// The type of authentication method
	Type ProviderListResponseAuthenticationMethodsType `json:"type,required"`
	// The supported benefit types and their configurations
	BenefitsSupport map[string]interface{} `json:"benefits_support"`
	// The supported fields for each Finch product
	SupportedFields map[string]interface{}                       `json:"supported_fields"`
	JSON            providerListResponseAuthenticationMethodJSON `json:"-"`
}

// providerListResponseAuthenticationMethodJSON contains the JSON metadata for the
// struct [ProviderListResponseAuthenticationMethod]
type providerListResponseAuthenticationMethodJSON struct {
	Type            apijson.Field
	BenefitsSupport apijson.Field
	SupportedFields apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodJSON) RawJSON() string {
	return r.raw
}

// The type of authentication method
type ProviderListResponseAuthenticationMethodsType string

const (
	ProviderListResponseAuthenticationMethodsTypeAssisted      ProviderListResponseAuthenticationMethodsType = "assisted"
	ProviderListResponseAuthenticationMethodsTypeCredential    ProviderListResponseAuthenticationMethodsType = "credential"
	ProviderListResponseAuthenticationMethodsTypeAPIToken      ProviderListResponseAuthenticationMethodsType = "api_token"
	ProviderListResponseAuthenticationMethodsTypeAPICredential ProviderListResponseAuthenticationMethodsType = "api_credential"
	ProviderListResponseAuthenticationMethodsTypeOAuth         ProviderListResponseAuthenticationMethodsType = "oauth"
	ProviderListResponseAuthenticationMethodsTypeAPI           ProviderListResponseAuthenticationMethodsType = "api"
)

func (r ProviderListResponseAuthenticationMethodsType) IsKnown() bool {
	switch r {
	case ProviderListResponseAuthenticationMethodsTypeAssisted, ProviderListResponseAuthenticationMethodsTypeCredential, ProviderListResponseAuthenticationMethodsTypeAPIToken, ProviderListResponseAuthenticationMethodsTypeAPICredential, ProviderListResponseAuthenticationMethodsTypeOAuth, ProviderListResponseAuthenticationMethodsTypeAPI:
		return true
	}
	return false
}
