// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/pagination"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// PayrollPayGroupService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayrollPayGroupService] method instead.
type PayrollPayGroupService struct {
	Options []option.RequestOption
}

// NewPayrollPayGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPayrollPayGroupService(opts ...option.RequestOption) (r *PayrollPayGroupService) {
	r = &PayrollPayGroupService{}
	r.Options = opts
	return
}

// Read information from a single pay group
func (r *PayrollPayGroupService) Get(ctx context.Context, payGroupID string, opts ...option.RequestOption) (res *PayrollPayGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if payGroupID == "" {
		err = errors.New("missing required pay_group_id parameter")
		return
	}
	path := fmt.Sprintf("employer/pay-groups/%s", payGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Read company pay groups and frequencies
func (r *PayrollPayGroupService) List(ctx context.Context, query PayrollPayGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PayrollPayGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/pay-groups"
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

// Read company pay groups and frequencies
func (r *PayrollPayGroupService) ListAutoPaging(ctx context.Context, query PayrollPayGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PayrollPayGroupListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type PayrollPayGroupGetResponse struct {
	// Finch id (uuidv4) for the pay group
	ID            string   `json:"id,required" format:"uuid"`
	IndividualIDs []string `json:"individual_ids,required" format:"uuid"`
	// Name of the pay group
	Name string `json:"name,required"`
	// List of pay frequencies associated with this pay group
	PayFrequencies []PayrollPayGroupGetResponsePayFrequency `json:"pay_frequencies,required"`
	JSON           payrollPayGroupGetResponseJSON           `json:"-"`
}

// payrollPayGroupGetResponseJSON contains the JSON metadata for the struct
// [PayrollPayGroupGetResponse]
type payrollPayGroupGetResponseJSON struct {
	ID             apijson.Field
	IndividualIDs  apijson.Field
	Name           apijson.Field
	PayFrequencies apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PayrollPayGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payrollPayGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type PayrollPayGroupGetResponsePayFrequency string

const (
	PayrollPayGroupGetResponsePayFrequencyAnnually     PayrollPayGroupGetResponsePayFrequency = "annually"
	PayrollPayGroupGetResponsePayFrequencySemiAnnually PayrollPayGroupGetResponsePayFrequency = "semi_annually"
	PayrollPayGroupGetResponsePayFrequencyQuarterly    PayrollPayGroupGetResponsePayFrequency = "quarterly"
	PayrollPayGroupGetResponsePayFrequencyMonthly      PayrollPayGroupGetResponsePayFrequency = "monthly"
	PayrollPayGroupGetResponsePayFrequencySemiMonthly  PayrollPayGroupGetResponsePayFrequency = "semi_monthly"
	PayrollPayGroupGetResponsePayFrequencyBiWeekly     PayrollPayGroupGetResponsePayFrequency = "bi_weekly"
	PayrollPayGroupGetResponsePayFrequencyWeekly       PayrollPayGroupGetResponsePayFrequency = "weekly"
	PayrollPayGroupGetResponsePayFrequencyDaily        PayrollPayGroupGetResponsePayFrequency = "daily"
	PayrollPayGroupGetResponsePayFrequencyOther        PayrollPayGroupGetResponsePayFrequency = "other"
)

func (r PayrollPayGroupGetResponsePayFrequency) IsKnown() bool {
	switch r {
	case PayrollPayGroupGetResponsePayFrequencyAnnually, PayrollPayGroupGetResponsePayFrequencySemiAnnually, PayrollPayGroupGetResponsePayFrequencyQuarterly, PayrollPayGroupGetResponsePayFrequencyMonthly, PayrollPayGroupGetResponsePayFrequencySemiMonthly, PayrollPayGroupGetResponsePayFrequencyBiWeekly, PayrollPayGroupGetResponsePayFrequencyWeekly, PayrollPayGroupGetResponsePayFrequencyDaily, PayrollPayGroupGetResponsePayFrequencyOther:
		return true
	}
	return false
}

type PayrollPayGroupListResponse struct {
	// Finch id (uuidv4) for the pay group
	ID string `json:"id" format:"uuid"`
	// Name of the pay group
	Name string `json:"name"`
	// List of pay frequencies associated with this pay group
	PayFrequencies []PayrollPayGroupListResponsePayFrequency `json:"pay_frequencies"`
	JSON           payrollPayGroupListResponseJSON           `json:"-"`
}

// payrollPayGroupListResponseJSON contains the JSON metadata for the struct
// [PayrollPayGroupListResponse]
type payrollPayGroupListResponseJSON struct {
	ID             apijson.Field
	Name           apijson.Field
	PayFrequencies apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PayrollPayGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payrollPayGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type PayrollPayGroupListResponsePayFrequency string

const (
	PayrollPayGroupListResponsePayFrequencyAnnually     PayrollPayGroupListResponsePayFrequency = "annually"
	PayrollPayGroupListResponsePayFrequencySemiAnnually PayrollPayGroupListResponsePayFrequency = "semi_annually"
	PayrollPayGroupListResponsePayFrequencyQuarterly    PayrollPayGroupListResponsePayFrequency = "quarterly"
	PayrollPayGroupListResponsePayFrequencyMonthly      PayrollPayGroupListResponsePayFrequency = "monthly"
	PayrollPayGroupListResponsePayFrequencySemiMonthly  PayrollPayGroupListResponsePayFrequency = "semi_monthly"
	PayrollPayGroupListResponsePayFrequencyBiWeekly     PayrollPayGroupListResponsePayFrequency = "bi_weekly"
	PayrollPayGroupListResponsePayFrequencyWeekly       PayrollPayGroupListResponsePayFrequency = "weekly"
	PayrollPayGroupListResponsePayFrequencyDaily        PayrollPayGroupListResponsePayFrequency = "daily"
	PayrollPayGroupListResponsePayFrequencyOther        PayrollPayGroupListResponsePayFrequency = "other"
)

func (r PayrollPayGroupListResponsePayFrequency) IsKnown() bool {
	switch r {
	case PayrollPayGroupListResponsePayFrequencyAnnually, PayrollPayGroupListResponsePayFrequencySemiAnnually, PayrollPayGroupListResponsePayFrequencyQuarterly, PayrollPayGroupListResponsePayFrequencyMonthly, PayrollPayGroupListResponsePayFrequencySemiMonthly, PayrollPayGroupListResponsePayFrequencyBiWeekly, PayrollPayGroupListResponsePayFrequencyWeekly, PayrollPayGroupListResponsePayFrequencyDaily, PayrollPayGroupListResponsePayFrequencyOther:
		return true
	}
	return false
}

type PayrollPayGroupListParams struct {
	IndividualID   param.Field[string]   `query:"individual_id" format:"uuid"`
	PayFrequencies param.Field[[]string] `query:"pay_frequencies"`
}

// URLQuery serializes [PayrollPayGroupListParams]'s query parameters as
// `url.Values`.
func (r PayrollPayGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
