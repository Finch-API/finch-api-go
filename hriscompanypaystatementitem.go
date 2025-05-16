// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
)

// HRISCompanyPayStatementItemService contains methods and other services that help
// with interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISCompanyPayStatementItemService] method instead.
type HRISCompanyPayStatementItemService struct {
	Options []option.RequestOption
	Rules   *HRISCompanyPayStatementItemRuleService
}

// NewHRISCompanyPayStatementItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHRISCompanyPayStatementItemService(opts ...option.RequestOption) (r *HRISCompanyPayStatementItemService) {
	r = &HRISCompanyPayStatementItemService{}
	r.Options = opts
	r.Rules = NewHRISCompanyPayStatementItemRuleService(opts...)
	return
}

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon Retrieve a list of detailed pay statement
// items for the access token's connection account.
func (r *HRISCompanyPayStatementItemService) List(ctx context.Context, query HRISCompanyPayStatementItemListParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[HRISCompanyPayStatementItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon Retrieve a list of detailed pay statement
// items for the access token's connection account.
func (r *HRISCompanyPayStatementItemService) ListAutoPaging(ctx context.Context, query HRISCompanyPayStatementItemListParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[HRISCompanyPayStatementItemListResponse] {
	return pagination.NewResponsesPageAutoPager(r.List(ctx, query, opts...))
}

type HRISCompanyPayStatementItemListResponse struct {
	// The attributes of the pay statement item.
	Attributes HRISCompanyPayStatementItemListResponseAttributes `json:"attributes"`
	// The category of the pay statement item.
	Category HRISCompanyPayStatementItemListResponseCategory `json:"category"`
	// The name of the pay statement item.
	Name string                                      `json:"name"`
	JSON hrisCompanyPayStatementItemListResponseJSON `json:"-"`
}

// hrisCompanyPayStatementItemListResponseJSON contains the JSON metadata for the
// struct [HRISCompanyPayStatementItemListResponse]
type hrisCompanyPayStatementItemListResponseJSON struct {
	Attributes  apijson.Field
	Category    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemListResponseJSON) RawJSON() string {
	return r.raw
}

// The attributes of the pay statement item.
type HRISCompanyPayStatementItemListResponseAttributes struct {
	// `true` if the amount is paid by the employers. This field is only available for
	// taxes.
	Employer bool `json:"employer,nullable"`
	// The metadata of the pay statement item derived by the rules engine if available.
	// Each attribute will be a key-value pair defined by a rule.
	Metadata map[string]interface{} `json:"metadata,nullable"`
	// `true` if the pay statement item is pre-tax. This field is only available for
	// employee deductions.
	PreTax bool `json:"pre_tax,nullable"`
	// The type of the pay statement item.
	Type string                                                `json:"type,nullable"`
	JSON hrisCompanyPayStatementItemListResponseAttributesJSON `json:"-"`
}

// hrisCompanyPayStatementItemListResponseAttributesJSON contains the JSON metadata
// for the struct [HRISCompanyPayStatementItemListResponseAttributes]
type hrisCompanyPayStatementItemListResponseAttributesJSON struct {
	Employer    apijson.Field
	Metadata    apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemListResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemListResponseAttributesJSON) RawJSON() string {
	return r.raw
}

// The category of the pay statement item.
type HRISCompanyPayStatementItemListResponseCategory string

const (
	HRISCompanyPayStatementItemListResponseCategoryEarnings              HRISCompanyPayStatementItemListResponseCategory = "earnings"
	HRISCompanyPayStatementItemListResponseCategoryTaxes                 HRISCompanyPayStatementItemListResponseCategory = "taxes"
	HRISCompanyPayStatementItemListResponseCategoryEmployeeDeductions    HRISCompanyPayStatementItemListResponseCategory = "employee_deductions"
	HRISCompanyPayStatementItemListResponseCategoryEmployerContributions HRISCompanyPayStatementItemListResponseCategory = "employer_contributions"
)

func (r HRISCompanyPayStatementItemListResponseCategory) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemListResponseCategoryEarnings, HRISCompanyPayStatementItemListResponseCategoryTaxes, HRISCompanyPayStatementItemListResponseCategoryEmployeeDeductions, HRISCompanyPayStatementItemListResponseCategoryEmployerContributions:
		return true
	}
	return false
}

type HRISCompanyPayStatementItemListParams struct {
	// Comma-delimited list of pay statement item categories to filter on. If empty,
	// defaults to all categories.
	Categories param.Field[[]HRISCompanyPayStatementItemListParamsCategory] `query:"categories"`
	// The end date to retrieve pay statement items by via their last seen pay date in
	// `YYYY-MM-DD` format.
	EndDate param.Field[time.Time] `query:"end_date" format:"date"`
	// Case-insensitive partial match search by pay statement item name.
	Name param.Field[string] `query:"name"`
	// The start date to retrieve pay statement items by via their last seen pay date
	// (inclusive) in `YYYY-MM-DD` format.
	StartDate param.Field[time.Time] `query:"start_date" format:"date"`
	// String search by pay statement item type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [HRISCompanyPayStatementItemListParams]'s query parameters
// as `url.Values`.
func (r HRISCompanyPayStatementItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISCompanyPayStatementItemListParamsCategory string

const (
	HRISCompanyPayStatementItemListParamsCategoryEarnings              HRISCompanyPayStatementItemListParamsCategory = "earnings"
	HRISCompanyPayStatementItemListParamsCategoryTaxes                 HRISCompanyPayStatementItemListParamsCategory = "taxes"
	HRISCompanyPayStatementItemListParamsCategoryEmployeeDeductions    HRISCompanyPayStatementItemListParamsCategory = "employee_deductions"
	HRISCompanyPayStatementItemListParamsCategoryEmployerContributions HRISCompanyPayStatementItemListParamsCategory = "employer_contributions"
)

func (r HRISCompanyPayStatementItemListParamsCategory) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemListParamsCategoryEarnings, HRISCompanyPayStatementItemListParamsCategoryTaxes, HRISCompanyPayStatementItemListParamsCategoryEmployeeDeductions, HRISCompanyPayStatementItemListParamsCategoryEmployerContributions:
		return true
	}
	return false
}
