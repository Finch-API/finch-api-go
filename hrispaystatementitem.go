// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
)

// HRISPayStatementItemService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISPayStatementItemService] method instead.
type HRISPayStatementItemService struct {
	Options []option.RequestOption
	Rules   *HRISPayStatementItemRuleService
}

// NewHRISPayStatementItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHRISPayStatementItemService(opts ...option.RequestOption) (r *HRISPayStatementItemService) {
	r = &HRISPayStatementItemService{}
	r.Options = opts
	r.Rules = NewHRISPayStatementItemRuleService(opts...)
	return
}

// Retrieve a list of detailed pay statement items for the access token's
// connection account.
func (r *HRISPayStatementItemService) List(ctx context.Context, query HRISPayStatementItemListParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[HRISPayStatementItemListResponse], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/pay-statement-item"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieve a list of detailed pay statement items for the access token's
// connection account.
func (r *HRISPayStatementItemService) ListAutoPaging(ctx context.Context, query HRISPayStatementItemListParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[HRISPayStatementItemListResponse] {
	return pagination.NewResponsesPageAutoPager(r.List(ctx, query, opts...))
}

type HRISPayStatementItemListResponse struct {
	// The attributes of the pay statement item.
	Attributes HRISPayStatementItemListResponseAttributes `json:"attributes" api:"required"`
	// The category of the pay statement item.
	Category HRISPayStatementItemListResponseCategory `json:"category" api:"required"`
	// The name of the pay statement item.
	Name string                               `json:"name" api:"required"`
	JSON hrisPayStatementItemListResponseJSON `json:"-"`
}

// hrisPayStatementItemListResponseJSON contains the JSON metadata for the struct
// [HRISPayStatementItemListResponse]
type hrisPayStatementItemListResponseJSON struct {
	Attributes  apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemListResponseJSON) RawJSON() string {
	return r.raw
}

// The attributes of the pay statement item.
type HRISPayStatementItemListResponseAttributes struct {
	// The metadata of the pay statement item derived by the rules engine if available.
	// Each attribute will be a key-value pair defined by a rule.
	Metadata map[string]interface{} `json:"metadata" api:"required,nullable"`
	// `true` if the amount is paid by the employers. This field is only available for
	// taxes.
	Employer bool `json:"employer" api:"nullable"`
	// `true` if the pay statement item is pre-tax. This field is only available for
	// employee deductions.
	PreTax bool `json:"pre_tax" api:"nullable"`
	// The type of the pay statement item.
	Type string                                         `json:"type" api:"nullable"`
	JSON hrisPayStatementItemListResponseAttributesJSON `json:"-"`
}

// hrisPayStatementItemListResponseAttributesJSON contains the JSON metadata for
// the struct [HRISPayStatementItemListResponseAttributes]
type hrisPayStatementItemListResponseAttributesJSON struct {
	Metadata    apijson.Field
	Employer    apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemListResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemListResponseAttributesJSON) RawJSON() string {
	return r.raw
}

// The category of the pay statement item.
type HRISPayStatementItemListResponseCategory string

const (
	HRISPayStatementItemListResponseCategoryEarnings              HRISPayStatementItemListResponseCategory = "earnings"
	HRISPayStatementItemListResponseCategoryTaxes                 HRISPayStatementItemListResponseCategory = "taxes"
	HRISPayStatementItemListResponseCategoryEmployeeDeductions    HRISPayStatementItemListResponseCategory = "employee_deductions"
	HRISPayStatementItemListResponseCategoryEmployerContributions HRISPayStatementItemListResponseCategory = "employer_contributions"
)

func (r HRISPayStatementItemListResponseCategory) IsKnown() bool {
	switch r {
	case HRISPayStatementItemListResponseCategoryEarnings, HRISPayStatementItemListResponseCategoryTaxes, HRISPayStatementItemListResponseCategoryEmployeeDeductions, HRISPayStatementItemListResponseCategoryEmployerContributions:
		return true
	}
	return false
}

type HRISPayStatementItemListParams struct {
	// Comma-delimited list of pay statement item categories to filter on. If empty,
	// defaults to all categories.
	Categories param.Field[[]HRISPayStatementItemListParamsCategory] `query:"categories"`
	// The end date to retrieve pay statement items by via their last seen pay date in
	// `YYYY-MM-DD` format.
	EndDate param.Field[time.Time] `query:"end_date" format:"date"`
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
	// Case-insensitive partial match search by pay statement item name.
	Name param.Field[string] `query:"name"`
	// The start date to retrieve pay statement items by via their last seen pay date
	// (inclusive) in `YYYY-MM-DD` format.
	StartDate param.Field[time.Time] `query:"start_date" format:"date"`
	// String search by pay statement item type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [HRISPayStatementItemListParams]'s query parameters as
// `url.Values`.
func (r HRISPayStatementItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISPayStatementItemListParamsCategory string

const (
	HRISPayStatementItemListParamsCategoryEarnings              HRISPayStatementItemListParamsCategory = "earnings"
	HRISPayStatementItemListParamsCategoryTaxes                 HRISPayStatementItemListParamsCategory = "taxes"
	HRISPayStatementItemListParamsCategoryEmployeeDeductions    HRISPayStatementItemListParamsCategory = "employee_deductions"
	HRISPayStatementItemListParamsCategoryEmployerContributions HRISPayStatementItemListParamsCategory = "employer_contributions"
)

func (r HRISPayStatementItemListParamsCategory) IsKnown() bool {
	switch r {
	case HRISPayStatementItemListParamsCategoryEarnings, HRISPayStatementItemListParamsCategoryTaxes, HRISPayStatementItemListParamsCategoryEmployeeDeductions, HRISPayStatementItemListParamsCategoryEmployerContributions:
		return true
	}
	return false
}
