// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/shared"
	"github.com/tidwall/gjson"
)

// HRISDocumentService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISDocumentService] method instead.
type HRISDocumentService struct {
	Options []option.RequestOption
}

// NewHRISDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISDocumentService(opts ...option.RequestOption) (r *HRISDocumentService) {
	r = &HRISDocumentService{}
	r.Options = opts
	return
}

// **Beta:** This endpoint is in beta and may change. Retrieve a list of
// company-wide documents.
func (r *HRISDocumentService) List(ctx context.Context, query HRISDocumentListParams, opts ...option.RequestOption) (res *HRISDocumentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "employer/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// **Beta:** This endpoint is in beta and may change. Retrieve details of a
// specific document by its ID.
func (r *HRISDocumentService) Retreive(ctx context.Context, documentID string, query HRISDocumentRetreiveParams, opts ...option.RequestOption) (res *HRISDocumentRetreiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return
	}
	path := fmt.Sprintf("employer/documents/%s", documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DocumentResponse struct {
	// A stable Finch id for the document.
	ID string `json:"id,required" format:"uuid"`
	// The ID of the individual associated with the document. This will be null for
	// employer-level documents.
	IndividualID string `json:"individual_id,required,nullable"`
	// The type of document.
	Type DocumentResponseType `json:"type,required"`
	// A URL to access the document. Format:
	// `https://api.tryfinch.com/employer/documents/:document_id`.
	URL string `json:"url,required" format:"uri"`
	// The year the document applies to, if available.
	Year float64              `json:"year,required"`
	JSON documentResponseJSON `json:"-"`
}

// documentResponseJSON contains the JSON metadata for the struct
// [DocumentResponse]
type documentResponseJSON struct {
	ID           apijson.Field
	IndividualID apijson.Field
	Type         apijson.Field
	URL          apijson.Field
	Year         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DocumentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentResponseJSON) RawJSON() string {
	return r.raw
}

// The type of document.
type DocumentResponseType string

const (
	DocumentResponseTypeW4_2020 DocumentResponseType = "w4_2020"
	DocumentResponseTypeW4_2005 DocumentResponseType = "w4_2005"
)

func (r DocumentResponseType) IsKnown() bool {
	switch r {
	case DocumentResponseTypeW4_2020, DocumentResponseTypeW4_2005:
		return true
	}
	return false
}

// A 2005 version of the W-4 tax form containing information on an individual's
// filing status, dependents, and withholding details.
type W42005 struct {
	// Detailed information specific to the 2005 W4 form.
	Data W42005Data `json:"data,required"`
	// Specifies the form type, indicating that this document is a 2005 W4 form.
	Type W42005Type `json:"type,required"`
	// The tax year this W4 document applies to.
	Year float64    `json:"year,required"`
	JSON w42005JSON `json:"-"`
}

// w42005JSON contains the JSON metadata for the struct [W42005]
type w42005JSON struct {
	Data        apijson.Field
	Type        apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *W42005) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r w42005JSON) RawJSON() string {
	return r.raw
}

func (r W42005) implementsHRISDocumentRetreiveResponse() {}

// Detailed information specific to the 2005 W4 form.
type W42005Data struct {
	// Additional withholding amount (in cents).
	AdditionalWithholding int64 `json:"additional_withholding,required"`
	// Indicates exemption status from federal tax withholding.
	Exemption W42005DataExemption `json:"exemption,required,nullable"`
	// The individual's filing status for tax purposes.
	FilingStatus W42005DataFilingStatus `json:"filing_status,required,nullable"`
	// The unique identifier for the individual associated with this 2005 W4 form.
	IndividualID string `json:"individual_id,required" format:"uuid"`
	// Total number of allowances claimed (in cents).
	TotalNumberOfAllowances int64          `json:"total_number_of_allowances,required"`
	JSON                    w42005DataJSON `json:"-"`
}

// w42005DataJSON contains the JSON metadata for the struct [W42005Data]
type w42005DataJSON struct {
	AdditionalWithholding   apijson.Field
	Exemption               apijson.Field
	FilingStatus            apijson.Field
	IndividualID            apijson.Field
	TotalNumberOfAllowances apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *W42005Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r w42005DataJSON) RawJSON() string {
	return r.raw
}

// Indicates exemption status from federal tax withholding.
type W42005DataExemption string

const (
	W42005DataExemptionExempt    W42005DataExemption = "exempt"
	W42005DataExemptionNonExempt W42005DataExemption = "non_exempt"
)

func (r W42005DataExemption) IsKnown() bool {
	switch r {
	case W42005DataExemptionExempt, W42005DataExemptionNonExempt:
		return true
	}
	return false
}

// The individual's filing status for tax purposes.
type W42005DataFilingStatus string

const (
	W42005DataFilingStatusMarried                              W42005DataFilingStatus = "married"
	W42005DataFilingStatusMarriedButWithholdAtHigherSingleRate W42005DataFilingStatus = "married_but_withhold_at_higher_single_rate"
	W42005DataFilingStatusSingle                               W42005DataFilingStatus = "single"
)

func (r W42005DataFilingStatus) IsKnown() bool {
	switch r {
	case W42005DataFilingStatusMarried, W42005DataFilingStatusMarriedButWithholdAtHigherSingleRate, W42005DataFilingStatusSingle:
		return true
	}
	return false
}

// Specifies the form type, indicating that this document is a 2005 W4 form.
type W42005Type string

const (
	W42005TypeW4_2005 W42005Type = "w4_2005"
)

func (r W42005Type) IsKnown() bool {
	switch r {
	case W42005TypeW4_2005:
		return true
	}
	return false
}

// A 2020 version of the W-4 tax form containing information on an individual's
// filing status, dependents, and withholding details.
type W42020 struct {
	// Detailed information specific to the 2020 W4 form.
	Data W42020Data `json:"data,required"`
	// Specifies the form type, indicating that this document is a 2020 W4 form.
	Type W42020Type `json:"type,required"`
	// The tax year this W4 document applies to.
	Year float64    `json:"year,required"`
	JSON w42020JSON `json:"-"`
}

// w42020JSON contains the JSON metadata for the struct [W42020]
type w42020JSON struct {
	Data        apijson.Field
	Type        apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *W42020) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r w42020JSON) RawJSON() string {
	return r.raw
}

func (r W42020) implementsHRISDocumentRetreiveResponse() {}

// Detailed information specific to the 2020 W4 form.
type W42020Data struct {
	// Amount claimed for dependents other than qualifying children under 17 (in
	// cents).
	AmountForOtherDependents int64 `json:"amount_for_other_dependents,required"`
	// Amount claimed for dependents under 17 years old (in cents).
	AmountForQualifyingChildrenUnder17 int64 `json:"amount_for_qualifying_children_under_17,required"`
	// Deductible expenses (in cents).
	Deductions int64 `json:"deductions,required"`
	// Additional withholding amount (in cents).
	ExtraWithholding int64 `json:"extra_withholding,required"`
	// The individual's filing status for tax purposes.
	FilingStatus W42020DataFilingStatus `json:"filing_status,required,nullable"`
	// The unique identifier for the individual associated with this document.
	IndividualID string `json:"individual_id,required" format:"uuid"`
	// Additional income from sources outside of primary employment (in cents).
	OtherIncome int64 `json:"other_income,required"`
	// Total amount claimed for dependents and other credits (in cents).
	TotalClaimDependentAndOtherCredits int64          `json:"total_claim_dependent_and_other_credits,required"`
	JSON                               w42020DataJSON `json:"-"`
}

// w42020DataJSON contains the JSON metadata for the struct [W42020Data]
type w42020DataJSON struct {
	AmountForOtherDependents           apijson.Field
	AmountForQualifyingChildrenUnder17 apijson.Field
	Deductions                         apijson.Field
	ExtraWithholding                   apijson.Field
	FilingStatus                       apijson.Field
	IndividualID                       apijson.Field
	OtherIncome                        apijson.Field
	TotalClaimDependentAndOtherCredits apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *W42020Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r w42020DataJSON) RawJSON() string {
	return r.raw
}

// The individual's filing status for tax purposes.
type W42020DataFilingStatus string

const (
	W42020DataFilingStatusHeadOfHousehold                                 W42020DataFilingStatus = "head_of_household"
	W42020DataFilingStatusMarriedFilingJointlyOrQualifyingSurvivingSpouse W42020DataFilingStatus = "married_filing_jointly_or_qualifying_surviving_spouse"
	W42020DataFilingStatusSingleOrMarriedFilingSeparately                 W42020DataFilingStatus = "single_or_married_filing_separately"
)

func (r W42020DataFilingStatus) IsKnown() bool {
	switch r {
	case W42020DataFilingStatusHeadOfHousehold, W42020DataFilingStatusMarriedFilingJointlyOrQualifyingSurvivingSpouse, W42020DataFilingStatusSingleOrMarriedFilingSeparately:
		return true
	}
	return false
}

// Specifies the form type, indicating that this document is a 2020 W4 form.
type W42020Type string

const (
	W42020TypeW4_2020 W42020Type = "w4_2020"
)

func (r W42020Type) IsKnown() bool {
	switch r {
	case W42020TypeW4_2020:
		return true
	}
	return false
}

type HRISDocumentListResponse struct {
	Documents []DocumentResponse           `json:"documents,required"`
	Paging    shared.Paging                `json:"paging,required"`
	JSON      hrisDocumentListResponseJSON `json:"-"`
}

// hrisDocumentListResponseJSON contains the JSON metadata for the struct
// [HRISDocumentListResponse]
type hrisDocumentListResponseJSON struct {
	Documents   apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISDocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisDocumentListResponseJSON) RawJSON() string {
	return r.raw
}

// A 2020 version of the W-4 tax form containing information on an individual's
// filing status, dependents, and withholding details.
type HRISDocumentRetreiveResponse struct {
	// This field can have the runtime type of [W42020Data], [W42005Data].
	Data interface{} `json:"data,required"`
	// Specifies the form type, indicating that this document is a 2020 W4 form.
	Type HRISDocumentRetreiveResponseType `json:"type,required"`
	// The tax year this W4 document applies to.
	Year  float64                          `json:"year,required"`
	JSON  hrisDocumentRetreiveResponseJSON `json:"-"`
	union HRISDocumentRetreiveResponseUnion
}

// hrisDocumentRetreiveResponseJSON contains the JSON metadata for the struct
// [HRISDocumentRetreiveResponse]
type hrisDocumentRetreiveResponseJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r hrisDocumentRetreiveResponseJSON) RawJSON() string {
	return r.raw
}

func (r *HRISDocumentRetreiveResponse) UnmarshalJSON(data []byte) (err error) {
	*r = HRISDocumentRetreiveResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [HRISDocumentRetreiveResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [W42020], [W42005].
func (r HRISDocumentRetreiveResponse) AsUnion() HRISDocumentRetreiveResponseUnion {
	return r.union
}

// A 2020 version of the W-4 tax form containing information on an individual's
// filing status, dependents, and withholding details.
//
// Union satisfied by [W42020] or [W42005].
type HRISDocumentRetreiveResponseUnion interface {
	implementsHRISDocumentRetreiveResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*HRISDocumentRetreiveResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(W42020{}),
			DiscriminatorValue: "w4_2020",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(W42005{}),
			DiscriminatorValue: "w4_2005",
		},
	)
}

// Specifies the form type, indicating that this document is a 2020 W4 form.
type HRISDocumentRetreiveResponseType string

const (
	HRISDocumentRetreiveResponseTypeW4_2020 HRISDocumentRetreiveResponseType = "w4_2020"
	HRISDocumentRetreiveResponseTypeW4_2005 HRISDocumentRetreiveResponseType = "w4_2005"
)

func (r HRISDocumentRetreiveResponseType) IsKnown() bool {
	switch r {
	case HRISDocumentRetreiveResponseTypeW4_2020, HRISDocumentRetreiveResponseTypeW4_2005:
		return true
	}
	return false
}

type HRISDocumentListParams struct {
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids,required" format:"uuid"`
	// Comma-delimited list of stable Finch uuids for each individual. If empty,
	// defaults to all individuals
	IndividualIDs param.Field[[]string] `query:"individual_ids"`
	// Number of documents to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
	// Comma-delimited list of document types to filter on. If empty, defaults to all
	// types
	Types param.Field[[]HRISDocumentListParamsType] `query:"types"`
}

// URLQuery serializes [HRISDocumentListParams]'s query parameters as `url.Values`.
func (r HRISDocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISDocumentListParamsType string

const (
	HRISDocumentListParamsTypeW4_2020 HRISDocumentListParamsType = "w4_2020"
	HRISDocumentListParamsTypeW4_2005 HRISDocumentListParamsType = "w4_2005"
)

func (r HRISDocumentListParamsType) IsKnown() bool {
	switch r {
	case HRISDocumentListParamsTypeW4_2020, HRISDocumentListParamsTypeW4_2005:
		return true
	}
	return false
}

type HRISDocumentRetreiveParams struct {
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids,required" format:"uuid"`
}

// URLQuery serializes [HRISDocumentRetreiveParams]'s query parameters as
// `url.Values`.
func (r HRISDocumentRetreiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
