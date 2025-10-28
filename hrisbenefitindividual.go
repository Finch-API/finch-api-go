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
	"github.com/Finch-API/finch-api-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// HRISBenefitIndividualService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISBenefitIndividualService] method instead.
type HRISBenefitIndividualService struct {
	Options []option.RequestOption
}

// NewHRISBenefitIndividualService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHRISBenefitIndividualService(opts ...option.RequestOption) (r *HRISBenefitIndividualService) {
	r = &HRISBenefitIndividualService{}
	r.Options = opts
	return
}

// Lists individuals currently enrolled in a given deduction.
func (r *HRISBenefitIndividualService) EnrolledIDs(ctx context.Context, benefitID string, opts ...option.RequestOption) (res *HRISBenefitIndividualEnrolledIDsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return
	}
	path := fmt.Sprintf("employer/benefits/%s/enrolled", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get enrollment information for the given individuals.
func (r *HRISBenefitIndividualService) GetManyBenefits(ctx context.Context, benefitID string, query HRISBenefitIndividualGetManyBenefitsParams, opts ...option.RequestOption) (res *pagination.SinglePage[IndividualBenefit], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return
	}
	path := fmt.Sprintf("employer/benefits/%s/individuals", benefitID)
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

// Get enrollment information for the given individuals.
func (r *HRISBenefitIndividualService) GetManyBenefitsAutoPaging(ctx context.Context, benefitID string, query HRISBenefitIndividualGetManyBenefitsParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[IndividualBenefit] {
	return pagination.NewSinglePageAutoPager(r.GetManyBenefits(ctx, benefitID, query, opts...))
}

// Unenroll individuals from a deduction or contribution
func (r *HRISBenefitIndividualService) UnenrollMany(ctx context.Context, benefitID string, body HRISBenefitIndividualUnenrollManyParams, opts ...option.RequestOption) (res *UnenrolledIndividualBenefitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return
	}
	path := fmt.Sprintf("employer/benefits/%s/individuals", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type IndividualBenefit struct {
	Body         IndividualBenefitBody `json:"body,required"`
	Code         int64                 `json:"code,required"`
	IndividualID string                `json:"individual_id,required"`
	JSON         individualBenefitJSON `json:"-"`
}

// individualBenefitJSON contains the JSON metadata for the struct
// [IndividualBenefit]
type individualBenefitJSON struct {
	Body         apijson.Field
	Code         apijson.Field
	IndividualID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IndividualBenefit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitJSON) RawJSON() string {
	return r.raw
}

type IndividualBenefitBody struct {
	// If the benefit supports annual maximum, the amount in cents for this individual.
	AnnualMaximum int64 `json:"annual_maximum,nullable"`
	// If the benefit supports catch up (401k, 403b, etc.), whether catch up is enabled
	// for this individual.
	CatchUp bool    `json:"catch_up,nullable"`
	Code    float64 `json:"code"`
	// This field can have the runtime type of
	// [IndividualBenefitBodyObjectCompanyContribution].
	CompanyContribution interface{} `json:"company_contribution"`
	// This field can have the runtime type of
	// [IndividualBenefitBodyObjectEmployeeDeduction].
	EmployeeDeduction interface{} `json:"employee_deduction"`
	FinchCode         string      `json:"finch_code"`
	// Type for HSA contribution limit if the benefit is a HSA.
	HsaContributionLimit IndividualBenefitBodyHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	Message              string                                    `json:"message"`
	Name                 string                                    `json:"name"`
	JSON                 individualBenefitBodyJSON                 `json:"-"`
	union                IndividualBenefitBodyUnion
}

// individualBenefitBodyJSON contains the JSON metadata for the struct
// [IndividualBenefitBody]
type individualBenefitBodyJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	Code                 apijson.Field
	CompanyContribution  apijson.Field
	EmployeeDeduction    apijson.Field
	FinchCode            apijson.Field
	HsaContributionLimit apijson.Field
	Message              apijson.Field
	Name                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r individualBenefitBodyJSON) RawJSON() string {
	return r.raw
}

func (r *IndividualBenefitBody) UnmarshalJSON(data []byte) (err error) {
	*r = IndividualBenefitBody{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IndividualBenefitBodyUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [IndividualBenefitBodyObject],
// [IndividualBenefitBodyBatchError].
func (r IndividualBenefitBody) AsUnion() IndividualBenefitBodyUnion {
	return r.union
}

// Union satisfied by [IndividualBenefitBodyObject] or
// [IndividualBenefitBodyBatchError].
type IndividualBenefitBodyUnion interface {
	implementsIndividualBenefitBody()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualBenefitBodyUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyBatchError{}),
		},
	)
}

type IndividualBenefitBodyObject struct {
	// If the benefit supports annual maximum, the amount in cents for this individual.
	AnnualMaximum int64 `json:"annual_maximum,required,nullable"`
	// If the benefit supports catch up (401k, 403b, etc.), whether catch up is enabled
	// for this individual.
	CatchUp             bool                                           `json:"catch_up,required,nullable"`
	CompanyContribution IndividualBenefitBodyObjectCompanyContribution `json:"company_contribution,required,nullable"`
	EmployeeDeduction   IndividualBenefitBodyObjectEmployeeDeduction   `json:"employee_deduction,required,nullable"`
	// Type for HSA contribution limit if the benefit is a HSA.
	HsaContributionLimit IndividualBenefitBodyObjectHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	JSON                 individualBenefitBodyObjectJSON                 `json:"-"`
}

// individualBenefitBodyObjectJSON contains the JSON metadata for the struct
// [IndividualBenefitBodyObject]
type individualBenefitBodyObjectJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	CompanyContribution  apijson.Field
	EmployeeDeduction    apijson.Field
	HsaContributionLimit apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IndividualBenefitBodyObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyObjectJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyObject) implementsIndividualBenefitBody() {}

type IndividualBenefitBodyObjectCompanyContribution struct {
	// Fixed contribution type.
	Type IndividualBenefitBodyObjectCompanyContributionType `json:"type,required"`
	// Contribution amount in cents.
	Amount int64 `json:"amount"`
	// This field can have the runtime type of
	// [[]IndividualBenefitBodyObjectCompanyContributionObjectTier].
	Tiers interface{}                                        `json:"tiers"`
	JSON  individualBenefitBodyObjectCompanyContributionJSON `json:"-"`
	union IndividualBenefitBodyObjectCompanyContributionUnion
}

// individualBenefitBodyObjectCompanyContributionJSON contains the JSON metadata
// for the struct [IndividualBenefitBodyObjectCompanyContribution]
type individualBenefitBodyObjectCompanyContributionJSON struct {
	Type        apijson.Field
	Amount      apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r individualBenefitBodyObjectCompanyContributionJSON) RawJSON() string {
	return r.raw
}

func (r *IndividualBenefitBodyObjectCompanyContribution) UnmarshalJSON(data []byte) (err error) {
	*r = IndividualBenefitBodyObjectCompanyContribution{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IndividualBenefitBodyObjectCompanyContributionUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [IndividualBenefitBodyObjectCompanyContributionObject],
// [IndividualBenefitBodyObjectCompanyContributionObject],
// [IndividualBenefitBodyObjectCompanyContributionObject].
func (r IndividualBenefitBodyObjectCompanyContribution) AsUnion() IndividualBenefitBodyObjectCompanyContributionUnion {
	return r.union
}

// Union satisfied by [IndividualBenefitBodyObjectCompanyContributionObject],
// [IndividualBenefitBodyObjectCompanyContributionObject] or
// [IndividualBenefitBodyObjectCompanyContributionObject].
type IndividualBenefitBodyObjectCompanyContributionUnion interface {
	implementsIndividualBenefitBodyObjectCompanyContribution()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualBenefitBodyObjectCompanyContributionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyObjectCompanyContributionObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyObjectCompanyContributionObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyObjectCompanyContributionObject{}),
		},
	)
}

type IndividualBenefitBodyObjectCompanyContributionObject struct {
	// Contribution amount in cents.
	Amount int64 `json:"amount,required"`
	// Fixed contribution type.
	Type IndividualBenefitBodyObjectCompanyContributionObjectType `json:"type,required"`
	JSON individualBenefitBodyObjectCompanyContributionObjectJSON `json:"-"`
}

// individualBenefitBodyObjectCompanyContributionObjectJSON contains the JSON
// metadata for the struct [IndividualBenefitBodyObjectCompanyContributionObject]
type individualBenefitBodyObjectCompanyContributionObjectJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyObjectCompanyContributionObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyObjectCompanyContributionObjectJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyObjectCompanyContributionObject) implementsIndividualBenefitBodyObjectCompanyContribution() {
}

// Fixed contribution type.
type IndividualBenefitBodyObjectCompanyContributionObjectType string

const (
	IndividualBenefitBodyObjectCompanyContributionObjectTypeFixed IndividualBenefitBodyObjectCompanyContributionObjectType = "fixed"
)

func (r IndividualBenefitBodyObjectCompanyContributionObjectType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyObjectCompanyContributionObjectTypeFixed:
		return true
	}
	return false
}

// Fixed contribution type.
type IndividualBenefitBodyObjectCompanyContributionType string

const (
	IndividualBenefitBodyObjectCompanyContributionTypeFixed   IndividualBenefitBodyObjectCompanyContributionType = "fixed"
	IndividualBenefitBodyObjectCompanyContributionTypePercent IndividualBenefitBodyObjectCompanyContributionType = "percent"
	IndividualBenefitBodyObjectCompanyContributionTypeTiered  IndividualBenefitBodyObjectCompanyContributionType = "tiered"
)

func (r IndividualBenefitBodyObjectCompanyContributionType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyObjectCompanyContributionTypeFixed, IndividualBenefitBodyObjectCompanyContributionTypePercent, IndividualBenefitBodyObjectCompanyContributionTypeTiered:
		return true
	}
	return false
}

type IndividualBenefitBodyObjectEmployeeDeduction struct {
	// Contribution amount in cents.
	Amount int64 `json:"amount,required"`
	// Fixed contribution type.
	Type  IndividualBenefitBodyObjectEmployeeDeductionType `json:"type,required"`
	JSON  individualBenefitBodyObjectEmployeeDeductionJSON `json:"-"`
	union IndividualBenefitBodyObjectEmployeeDeductionUnion
}

// individualBenefitBodyObjectEmployeeDeductionJSON contains the JSON metadata for
// the struct [IndividualBenefitBodyObjectEmployeeDeduction]
type individualBenefitBodyObjectEmployeeDeductionJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r individualBenefitBodyObjectEmployeeDeductionJSON) RawJSON() string {
	return r.raw
}

func (r *IndividualBenefitBodyObjectEmployeeDeduction) UnmarshalJSON(data []byte) (err error) {
	*r = IndividualBenefitBodyObjectEmployeeDeduction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IndividualBenefitBodyObjectEmployeeDeductionUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [IndividualBenefitBodyObjectEmployeeDeductionObject],
// [IndividualBenefitBodyObjectEmployeeDeductionObject].
func (r IndividualBenefitBodyObjectEmployeeDeduction) AsUnion() IndividualBenefitBodyObjectEmployeeDeductionUnion {
	return r.union
}

// Union satisfied by [IndividualBenefitBodyObjectEmployeeDeductionObject] or
// [IndividualBenefitBodyObjectEmployeeDeductionObject].
type IndividualBenefitBodyObjectEmployeeDeductionUnion interface {
	implementsIndividualBenefitBodyObjectEmployeeDeduction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualBenefitBodyObjectEmployeeDeductionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyObjectEmployeeDeductionObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyObjectEmployeeDeductionObject{}),
		},
	)
}

type IndividualBenefitBodyObjectEmployeeDeductionObject struct {
	// Contribution amount in cents.
	Amount int64 `json:"amount,required"`
	// Fixed contribution type.
	Type IndividualBenefitBodyObjectEmployeeDeductionObjectType `json:"type,required"`
	JSON individualBenefitBodyObjectEmployeeDeductionObjectJSON `json:"-"`
}

// individualBenefitBodyObjectEmployeeDeductionObjectJSON contains the JSON
// metadata for the struct [IndividualBenefitBodyObjectEmployeeDeductionObject]
type individualBenefitBodyObjectEmployeeDeductionObjectJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyObjectEmployeeDeductionObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyObjectEmployeeDeductionObjectJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyObjectEmployeeDeductionObject) implementsIndividualBenefitBodyObjectEmployeeDeduction() {
}

// Fixed contribution type.
type IndividualBenefitBodyObjectEmployeeDeductionObjectType string

const (
	IndividualBenefitBodyObjectEmployeeDeductionObjectTypeFixed IndividualBenefitBodyObjectEmployeeDeductionObjectType = "fixed"
)

func (r IndividualBenefitBodyObjectEmployeeDeductionObjectType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyObjectEmployeeDeductionObjectTypeFixed:
		return true
	}
	return false
}

// Fixed contribution type.
type IndividualBenefitBodyObjectEmployeeDeductionType string

const (
	IndividualBenefitBodyObjectEmployeeDeductionTypeFixed   IndividualBenefitBodyObjectEmployeeDeductionType = "fixed"
	IndividualBenefitBodyObjectEmployeeDeductionTypePercent IndividualBenefitBodyObjectEmployeeDeductionType = "percent"
)

func (r IndividualBenefitBodyObjectEmployeeDeductionType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyObjectEmployeeDeductionTypeFixed, IndividualBenefitBodyObjectEmployeeDeductionTypePercent:
		return true
	}
	return false
}

// Type for HSA contribution limit if the benefit is a HSA.
type IndividualBenefitBodyObjectHsaContributionLimit string

const (
	IndividualBenefitBodyObjectHsaContributionLimitIndividual IndividualBenefitBodyObjectHsaContributionLimit = "individual"
	IndividualBenefitBodyObjectHsaContributionLimitFamily     IndividualBenefitBodyObjectHsaContributionLimit = "family"
)

func (r IndividualBenefitBodyObjectHsaContributionLimit) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyObjectHsaContributionLimitIndividual, IndividualBenefitBodyObjectHsaContributionLimitFamily:
		return true
	}
	return false
}

type IndividualBenefitBodyBatchError struct {
	Code      float64                             `json:"code,required"`
	Message   string                              `json:"message,required"`
	Name      string                              `json:"name,required"`
	FinchCode string                              `json:"finch_code"`
	JSON      individualBenefitBodyBatchErrorJSON `json:"-"`
}

// individualBenefitBodyBatchErrorJSON contains the JSON metadata for the struct
// [IndividualBenefitBodyBatchError]
type individualBenefitBodyBatchErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	FinchCode   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyBatchError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyBatchErrorJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyBatchError) implementsIndividualBenefitBody() {}

// Type for HSA contribution limit if the benefit is a HSA.
type IndividualBenefitBodyHsaContributionLimit string

const (
	IndividualBenefitBodyHsaContributionLimitIndividual IndividualBenefitBodyHsaContributionLimit = "individual"
	IndividualBenefitBodyHsaContributionLimitFamily     IndividualBenefitBodyHsaContributionLimit = "family"
)

func (r IndividualBenefitBodyHsaContributionLimit) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyHsaContributionLimitIndividual, IndividualBenefitBodyHsaContributionLimitFamily:
		return true
	}
	return false
}

type UnenrolledIndividualBenefitResponse struct {
	JobID string                                  `json:"job_id,required" format:"uuid"`
	JSON  unenrolledIndividualBenefitResponseJSON `json:"-"`
}

// unenrolledIndividualBenefitResponseJSON contains the JSON metadata for the
// struct [UnenrolledIndividualBenefitResponse]
type unenrolledIndividualBenefitResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnenrolledIndividualBenefitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unenrolledIndividualBenefitResponseJSON) RawJSON() string {
	return r.raw
}

type HRISBenefitIndividualEnrolledIDsResponse struct {
	// The id of the benefit.
	BenefitID     string                                       `json:"benefit_id,required" format:"uuid"`
	IndividualIDs []string                                     `json:"individual_ids,required" format:"uuid"`
	JSON          hrisBenefitIndividualEnrolledIDsResponseJSON `json:"-"`
}

// hrisBenefitIndividualEnrolledIDsResponseJSON contains the JSON metadata for the
// struct [HRISBenefitIndividualEnrolledIDsResponse]
type hrisBenefitIndividualEnrolledIDsResponseJSON struct {
	BenefitID     apijson.Field
	IndividualIDs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *HRISBenefitIndividualEnrolledIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisBenefitIndividualEnrolledIDsResponseJSON) RawJSON() string {
	return r.raw
}

type HRISBenefitIndividualGetManyBenefitsParams struct {
	// comma-delimited list of stable Finch uuids for each individual. If empty,
	// defaults to all individuals
	IndividualIDs param.Field[string] `query:"individual_ids"`
}

// URLQuery serializes [HRISBenefitIndividualGetManyBenefitsParams]'s query
// parameters as `url.Values`.
func (r HRISBenefitIndividualGetManyBenefitsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISBenefitIndividualUnenrollManyParams struct {
	// Array of individual_ids to unenroll.
	IndividualIDs param.Field[[]string] `json:"individual_ids"`
}

func (r HRISBenefitIndividualUnenrollManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
