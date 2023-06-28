// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// ProviderService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewProviderService] method instead.
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
func (r *ProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *shared.SinglePage[Provider], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *ProviderService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *shared.SinglePageAutoPager[Provider] {
	return shared.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type Provider struct {
	// The id of the payroll provider used in Connect.
	ID string `json:"id"`
	// The display name of the payroll provider.
	DisplayName string `json:"display_name"`
	// The url to the official icon of the payroll provider.
	Icon string `json:"icon"`
	// The url to the official logo of the payroll provider.
	Logo string `json:"logo"`
	// Whether the Finch integration with this provider uses the Assisted Connect Flow
	// by default.
	Manual bool `json:"manual"`
	// whether MFA is required for the provider.
	MfaRequired bool `json:"mfa_required"`
	// The hex code for the primary color of the payroll provider.
	PrimaryColor string `json:"primary_color"`
	// The list of Finch products supported on this payroll provider.
	Products []string `json:"products"`
	JSON     providerJSON
}

// providerJSON contains the JSON metadata for the struct [Provider]
type providerJSON struct {
	ID           apijson.Field
	DisplayName  apijson.Field
	Icon         apijson.Field
	Logo         apijson.Field
	Manual       apijson.Field
	MfaRequired  apijson.Field
	PrimaryColor apijson.Field
	Products     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Provider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
