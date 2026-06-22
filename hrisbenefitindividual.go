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
	"time"

	"github.com/Finch-API/finch-api-go/v2/internal/apijson"
	"github.com/Finch-API/finch-api-go/v2/internal/apiquery"
	"github.com/Finch-API/finch-api-go/v2/internal/param"
	"github.com/Finch-API/finch-api-go/v2/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/v2/option"
	"github.com/Finch-API/finch-api-go/v2/packages/pagination"
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

// Enroll an individual into a deduction or contribution. This is an overwrite
// operation. If the employee is already enrolled, the enrollment amounts will be
// adjusted. Making the same request multiple times will not create new
// enrollments, but will continue to set the state of the existing enrollment.
func (r *HRISBenefitIndividualService) EnrollMany(ctx context.Context, benefitID string, params HRISBenefitIndividualEnrollManyParams, opts ...option.RequestOption) (res *EnrolledIndividualBenefitResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("employer/benefits/%s/individuals", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Lists individuals currently enrolled in a given deduction.
func (r *HRISBenefitIndividualService) EnrolledIDs(ctx context.Context, benefitID string, query HRISBenefitIndividualEnrolledIDsParams, opts ...option.RequestOption) (res *HRISBenefitIndividualEnrolledIDsResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("employer/benefits/%s/enrolled", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get enrollment information for the given individuals.
func (r *HRISBenefitIndividualService) GetManyBenefits(ctx context.Context, benefitID string, query HRISBenefitIndividualGetManyBenefitsParams, opts ...option.RequestOption) (res *pagination.SinglePage[IndividualBenefit], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return nil, err
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
func (r *HRISBenefitIndividualService) UnenrollMany(ctx context.Context, benefitID string, params HRISBenefitIndividualUnenrollManyParams, opts ...option.RequestOption) (res *UnenrolledIndividualBenefitResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("employer/benefits/%s/individuals", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

type EnrolledIndividualBenefitResponse struct {
	JobID string                                `json:"job_id" api:"required" format:"uuid"`
	JSON  enrolledIndividualBenefitResponseJSON `json:"-"`
}

// enrolledIndividualBenefitResponseJSON contains the JSON metadata for the struct
// [EnrolledIndividualBenefitResponse]
type enrolledIndividualBenefitResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnrolledIndividualBenefitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enrolledIndividualBenefitResponseJSON) RawJSON() string {
	return r.raw
}

type IndividualBenefit struct {
	Body         IndividualBenefitBody `json:"body" api:"required"`
	Code         int64                 `json:"code" api:"required"`
	IndividualID string                `json:"individual_id" api:"required"`
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
	AnnualMaximum int64 `json:"annual_maximum" api:"nullable"`
	// If the benefit supports catch up (401k, 403b, etc.), whether catch up is enabled
	// for this individual.
	CatchUp bool    `json:"catch_up" api:"nullable"`
	Code    float64 `json:"code"`
	// This field can have the runtime type of
	// [IndividualBenefitBodyIndividualBenefitCompanyContribution].
	CompanyContribution interface{} `json:"company_contribution"`
	// This field can have the runtime type of
	// [IndividualBenefitBodyIndividualBenefitEmployeeDeduction].
	EmployeeDeduction interface{} `json:"employee_deduction"`
	FinchCode         string      `json:"finch_code"`
	// Type for HSA contribution limit if the benefit is a HSA.
	HsaContributionLimit IndividualBenefitBodyHsaContributionLimit `json:"hsa_contribution_limit" api:"nullable"`
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
// Possible runtime types of the union are
// [IndividualBenefitBodyIndividualBenefit], [IndividualBenefitBodyBatchError].
func (r IndividualBenefitBody) AsUnion() IndividualBenefitBodyUnion {
	return r.union
}

// Union satisfied by [IndividualBenefitBodyIndividualBenefit] or
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
			Type:       reflect.TypeOf(IndividualBenefitBodyIndividualBenefit{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyBatchError{}),
		},
	)
}

type IndividualBenefitBodyIndividualBenefit struct {
	// If the benefit supports annual maximum, the amount in cents for this individual.
	AnnualMaximum int64 `json:"annual_maximum" api:"required,nullable"`
	// If the benefit supports catch up (401k, 403b, etc.), whether catch up is enabled
	// for this individual.
	CatchUp bool `json:"catch_up" api:"required,nullable"`
	// Company contribution configuration. Supports fixed amounts (in cents),
	// percentage-based contributions (in basis points where 100 = 1%), or tiered
	// matching structures.
	CompanyContribution IndividualBenefitBodyIndividualBenefitCompanyContribution `json:"company_contribution" api:"required,nullable"`
	// Employee deduction configuration. Supports both fixed amounts (in cents) and
	// percentage-based contributions (in basis points where 100 = 1%).
	EmployeeDeduction IndividualBenefitBodyIndividualBenefitEmployeeDeduction `json:"employee_deduction" api:"required,nullable"`
	// Type for HSA contribution limit if the benefit is a HSA.
	HsaContributionLimit IndividualBenefitBodyIndividualBenefitHsaContributionLimit `json:"hsa_contribution_limit" api:"nullable"`
	JSON                 individualBenefitBodyIndividualBenefitJSON                 `json:"-"`
}

// individualBenefitBodyIndividualBenefitJSON contains the JSON metadata for the
// struct [IndividualBenefitBodyIndividualBenefit]
type individualBenefitBodyIndividualBenefitJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	CompanyContribution  apijson.Field
	EmployeeDeduction    apijson.Field
	HsaContributionLimit apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyIndividualBenefit) implementsIndividualBenefitBody() {}

// Company contribution configuration. Supports fixed amounts (in cents),
// percentage-based contributions (in basis points where 100 = 1%), or tiered
// matching structures.
type IndividualBenefitBodyIndividualBenefitCompanyContribution struct {
	// Contribution type. Supported values: "fixed" (amount in cents), "percent"
	// (amount in basis points), or "tiered" (multi-tier matching).
	Type IndividualBenefitBodyIndividualBenefitCompanyContributionType `json:"type" api:"required"`
	// Contribution amount in cents (for type=fixed) or basis points (for type=percent,
	// where 100 = 1%). Not used for type=tiered.
	Amount int64 `json:"amount"`
	// This field can have the runtime type of
	// [[]IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTier].
	Tiers interface{}                                                   `json:"tiers"`
	JSON  individualBenefitBodyIndividualBenefitCompanyContributionJSON `json:"-"`
	union IndividualBenefitBodyIndividualBenefitCompanyContributionUnion
}

// individualBenefitBodyIndividualBenefitCompanyContributionJSON contains the JSON
// metadata for the struct
// [IndividualBenefitBodyIndividualBenefitCompanyContribution]
type individualBenefitBodyIndividualBenefitCompanyContributionJSON struct {
	Type        apijson.Field
	Amount      apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r individualBenefitBodyIndividualBenefitCompanyContributionJSON) RawJSON() string {
	return r.raw
}

func (r *IndividualBenefitBodyIndividualBenefitCompanyContribution) UnmarshalJSON(data []byte) (err error) {
	*r = IndividualBenefitBodyIndividualBenefitCompanyContribution{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [IndividualBenefitBodyIndividualBenefitCompanyContributionUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed],
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent],
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered].
func (r IndividualBenefitBodyIndividualBenefitCompanyContribution) AsUnion() IndividualBenefitBodyIndividualBenefitCompanyContributionUnion {
	return r.union
}

// Company contribution configuration. Supports fixed amounts (in cents),
// percentage-based contributions (in basis points where 100 = 1%), or tiered
// matching structures.
//
// Union satisfied by
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed],
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent]
// or
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered].
type IndividualBenefitBodyIndividualBenefitCompanyContributionUnion interface {
	implementsIndividualBenefitBodyIndividualBenefitCompanyContribution()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualBenefitBodyIndividualBenefitCompanyContributionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered{}),
		},
	)
}

type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed struct {
	// Contribution amount in cents (for type=fixed) or basis points (for type=percent,
	// where 100 = 1%). Not used for type=tiered.
	Amount int64 `json:"amount" api:"required"`
	// Contribution type. Supported values: "fixed" (amount in cents), "percent"
	// (amount in basis points), or "tiered" (multi-tier matching).
	Type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedType `json:"type" api:"required"`
	JSON individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedJSON `json:"-"`
}

// individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedJSON
// contains the JSON metadata for the struct
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed]
type individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixed) implementsIndividualBenefitBodyIndividualBenefitCompanyContribution() {
}

// Contribution type. Supported values: "fixed" (amount in cents), "percent"
// (amount in basis points), or "tiered" (multi-tier matching).
type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedType string

const (
	IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedTypeFixed IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedType = "fixed"
)

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionFixedTypeFixed:
		return true
	}
	return false
}

type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent struct {
	// Contribution amount in cents (for type=fixed) or basis points (for type=percent,
	// where 100 = 1%). Not used for type=tiered.
	Amount int64 `json:"amount" api:"required"`
	// Contribution type. Supported values: "fixed" (amount in cents), "percent"
	// (amount in basis points), or "tiered" (multi-tier matching).
	Type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentType `json:"type" api:"required"`
	JSON individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentJSON `json:"-"`
}

// individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentJSON
// contains the JSON metadata for the struct
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent]
type individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercent) implementsIndividualBenefitBodyIndividualBenefitCompanyContribution() {
}

// Contribution type. Supported values: "fixed" (amount in cents), "percent"
// (amount in basis points), or "tiered" (multi-tier matching).
type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentType string

const (
	IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentTypePercent IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentType = "percent"
)

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionPercentTypePercent:
		return true
	}
	return false
}

type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered struct {
	// Array of tier objects defining employer match tiers based on employee
	// contribution thresholds. Required when type=tiered.
	Tiers []IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTier `json:"tiers" api:"required"`
	// Contribution type. Supported values: "fixed" (amount in cents), "percent"
	// (amount in basis points), or "tiered" (multi-tier matching).
	Type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredType `json:"type" api:"required"`
	JSON individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredJSON `json:"-"`
}

// individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredJSON
// contains the JSON metadata for the struct
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered]
type individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredJSON struct {
	Tiers       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTiered) implementsIndividualBenefitBodyIndividualBenefitCompanyContribution() {
}

type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTier struct {
	Match     int64                                                                                      `json:"match" api:"required"`
	Threshold int64                                                                                      `json:"threshold" api:"required"`
	JSON      individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTierJSON `json:"-"`
}

// individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTierJSON
// contains the JSON metadata for the struct
// [IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTier]
type individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTierJSON struct {
	Match       apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTierJSON) RawJSON() string {
	return r.raw
}

// Contribution type. Supported values: "fixed" (amount in cents), "percent"
// (amount in basis points), or "tiered" (multi-tier matching).
type IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredType string

const (
	IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTypeTiered IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredType = "tiered"
)

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitCompanyContributionCompanyContributionTieredTypeTiered:
		return true
	}
	return false
}

// Contribution type. Supported values: "fixed" (amount in cents), "percent"
// (amount in basis points), or "tiered" (multi-tier matching).
type IndividualBenefitBodyIndividualBenefitCompanyContributionType string

const (
	IndividualBenefitBodyIndividualBenefitCompanyContributionTypeFixed   IndividualBenefitBodyIndividualBenefitCompanyContributionType = "fixed"
	IndividualBenefitBodyIndividualBenefitCompanyContributionTypePercent IndividualBenefitBodyIndividualBenefitCompanyContributionType = "percent"
	IndividualBenefitBodyIndividualBenefitCompanyContributionTypeTiered  IndividualBenefitBodyIndividualBenefitCompanyContributionType = "tiered"
)

func (r IndividualBenefitBodyIndividualBenefitCompanyContributionType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitCompanyContributionTypeFixed, IndividualBenefitBodyIndividualBenefitCompanyContributionTypePercent, IndividualBenefitBodyIndividualBenefitCompanyContributionTypeTiered:
		return true
	}
	return false
}

// Employee deduction configuration. Supports both fixed amounts (in cents) and
// percentage-based contributions (in basis points where 100 = 1%).
type IndividualBenefitBodyIndividualBenefitEmployeeDeduction struct {
	// Contribution amount in cents (for type=fixed) or basis points (for type=percent,
	// where 100 = 1%).
	Amount int64 `json:"amount" api:"required"`
	// Contribution type. Supported values: "fixed" (amount in cents) or "percent"
	// (amount in basis points).
	Type  IndividualBenefitBodyIndividualBenefitEmployeeDeductionType `json:"type" api:"required"`
	JSON  individualBenefitBodyIndividualBenefitEmployeeDeductionJSON `json:"-"`
	union IndividualBenefitBodyIndividualBenefitEmployeeDeductionUnion
}

// individualBenefitBodyIndividualBenefitEmployeeDeductionJSON contains the JSON
// metadata for the struct
// [IndividualBenefitBodyIndividualBenefitEmployeeDeduction]
type individualBenefitBodyIndividualBenefitEmployeeDeductionJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r individualBenefitBodyIndividualBenefitEmployeeDeductionJSON) RawJSON() string {
	return r.raw
}

func (r *IndividualBenefitBodyIndividualBenefitEmployeeDeduction) UnmarshalJSON(data []byte) (err error) {
	*r = IndividualBenefitBodyIndividualBenefitEmployeeDeduction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IndividualBenefitBodyIndividualBenefitEmployeeDeductionUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed],
// [IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent].
func (r IndividualBenefitBodyIndividualBenefitEmployeeDeduction) AsUnion() IndividualBenefitBodyIndividualBenefitEmployeeDeductionUnion {
	return r.union
}

// Employee deduction configuration. Supports both fixed amounts (in cents) and
// percentage-based contributions (in basis points where 100 = 1%).
//
// Union satisfied by
// [IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed]
// or
// [IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent].
type IndividualBenefitBodyIndividualBenefitEmployeeDeductionUnion interface {
	implementsIndividualBenefitBodyIndividualBenefitEmployeeDeduction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualBenefitBodyIndividualBenefitEmployeeDeductionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent{}),
		},
	)
}

type IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed struct {
	// Contribution amount in cents (for type=fixed) or basis points (for type=percent,
	// where 100 = 1%).
	Amount int64 `json:"amount" api:"required"`
	// Contribution type. Supported values: "fixed" (amount in cents) or "percent"
	// (amount in basis points).
	Type IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedType `json:"type" api:"required"`
	JSON individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedJSON `json:"-"`
}

// individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedJSON
// contains the JSON metadata for the struct
// [IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed]
type individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixed) implementsIndividualBenefitBodyIndividualBenefitEmployeeDeduction() {
}

// Contribution type. Supported values: "fixed" (amount in cents) or "percent"
// (amount in basis points).
type IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedType string

const (
	IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedTypeFixed IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedType = "fixed"
)

func (r IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionFixedTypeFixed:
		return true
	}
	return false
}

type IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent struct {
	// Contribution amount in cents (for type=fixed) or basis points (for type=percent,
	// where 100 = 1%).
	Amount int64 `json:"amount" api:"required"`
	// Contribution type. Supported values: "fixed" (amount in cents) or "percent"
	// (amount in basis points).
	Type IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentType `json:"type" api:"required"`
	JSON individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentJSON `json:"-"`
}

// individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentJSON
// contains the JSON metadata for the struct
// [IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent]
type individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercent) implementsIndividualBenefitBodyIndividualBenefitEmployeeDeduction() {
}

// Contribution type. Supported values: "fixed" (amount in cents) or "percent"
// (amount in basis points).
type IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentType string

const (
	IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentTypePercent IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentType = "percent"
)

func (r IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitEmployeeDeductionEmployeeDeductionContributionPercentTypePercent:
		return true
	}
	return false
}

// Contribution type. Supported values: "fixed" (amount in cents) or "percent"
// (amount in basis points).
type IndividualBenefitBodyIndividualBenefitEmployeeDeductionType string

const (
	IndividualBenefitBodyIndividualBenefitEmployeeDeductionTypeFixed   IndividualBenefitBodyIndividualBenefitEmployeeDeductionType = "fixed"
	IndividualBenefitBodyIndividualBenefitEmployeeDeductionTypePercent IndividualBenefitBodyIndividualBenefitEmployeeDeductionType = "percent"
)

func (r IndividualBenefitBodyIndividualBenefitEmployeeDeductionType) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitEmployeeDeductionTypeFixed, IndividualBenefitBodyIndividualBenefitEmployeeDeductionTypePercent:
		return true
	}
	return false
}

// Type for HSA contribution limit if the benefit is a HSA.
type IndividualBenefitBodyIndividualBenefitHsaContributionLimit string

const (
	IndividualBenefitBodyIndividualBenefitHsaContributionLimitIndividual IndividualBenefitBodyIndividualBenefitHsaContributionLimit = "individual"
	IndividualBenefitBodyIndividualBenefitHsaContributionLimitFamily     IndividualBenefitBodyIndividualBenefitHsaContributionLimit = "family"
)

func (r IndividualBenefitBodyIndividualBenefitHsaContributionLimit) IsKnown() bool {
	switch r {
	case IndividualBenefitBodyIndividualBenefitHsaContributionLimitIndividual, IndividualBenefitBodyIndividualBenefitHsaContributionLimitFamily:
		return true
	}
	return false
}

type IndividualBenefitBodyBatchError struct {
	Code      float64                             `json:"code" api:"required"`
	Message   string                              `json:"message" api:"required"`
	Name      string                              `json:"name" api:"required"`
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
	JobID string                                  `json:"job_id" api:"required" format:"uuid"`
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
	BenefitID     string                                       `json:"benefit_id" api:"required" format:"uuid"`
	IndividualIDs []string                                     `json:"individual_ids" api:"required" format:"uuid"`
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

type HRISBenefitIndividualEnrollManyParams struct {
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
	// Array of the individual_id to enroll and a configuration object.
	Individuals []HRISBenefitIndividualEnrollManyParamsIndividual `json:"individuals"`
}

func (r HRISBenefitIndividualEnrollManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Individuals)
}

// URLQuery serializes [HRISBenefitIndividualEnrollManyParams]'s query parameters
// as `url.Values`.
func (r HRISBenefitIndividualEnrollManyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISBenefitIndividualEnrollManyParamsIndividual struct {
	Configuration param.Field[HRISBenefitIndividualEnrollManyParamsIndividualsConfiguration] `json:"configuration"`
	// Finch id (uuidv4) for the individual to enroll
	IndividualID param.Field[string] `json:"individual_id"`
}

func (r HRISBenefitIndividualEnrollManyParamsIndividual) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISBenefitIndividualEnrollManyParamsIndividualsConfiguration struct {
	// For HSA benefits only - whether the contribution limit is for an individual or
	// family
	AnnualContributionLimit param.Field[HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimit] `json:"annual_contribution_limit"`
	// Maximum annual amount in cents
	AnnualMaximum param.Field[int64] `json:"annual_maximum"`
	// For retirement benefits only - whether catch up contributions are enabled
	CatchUp             param.Field[bool]                                                                             `json:"catch_up"`
	CompanyContribution param.Field[HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContribution] `json:"company_contribution"`
	// The date the enrollment will take effect
	EffectiveDate     param.Field[time.Time]                                                                      `json:"effective_date" format:"date"`
	EmployeeDeduction param.Field[HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeduction] `json:"employee_deduction"`
}

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For HSA benefits only - whether the contribution limit is for an individual or
// family
type HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimit string

const (
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimitIndividual HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimit = "individual"
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimitFamily     HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimit = "family"
)

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimit) IsKnown() bool {
	switch r {
	case HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimitIndividual, HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimitFamily:
		return true
	}
	return false
}

type HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContribution struct {
	// Amount in cents for fixed type or basis points (1/100th of a percent) for
	// percent type
	Amount param.Field[int64] `json:"amount"`
	// Array of tier objects for tiered contribution matching (required when type is
	// tiered)
	Tiers param.Field[[]HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTier] `json:"tiers"`
	Type  param.Field[HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionType]   `json:"type"`
}

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContribution) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTier struct {
	// The employer match percentage in basis points (0-10000 = 0-100%)
	Match param.Field[int64] `json:"match" api:"required"`
	// The employee contribution threshold in basis points (0-10000 = 0-100%)
	Threshold param.Field[int64] `json:"threshold" api:"required"`
}

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionType string

const (
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypeFixed   HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionType = "fixed"
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypePercent HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionType = "percent"
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypeTiered  HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionType = "tiered"
)

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionType) IsKnown() bool {
	switch r {
	case HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypeFixed, HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypePercent, HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypeTiered:
		return true
	}
	return false
}

type HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeduction struct {
	// Amount in cents for fixed type or basis points (1/100th of a percent) for
	// percent type
	Amount param.Field[int64]                                                                              `json:"amount"`
	Type   param.Field[HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionType] `json:"type"`
}

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionType string

const (
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionTypeFixed   HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionType = "fixed"
	HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionTypePercent HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionType = "percent"
)

func (r HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionType) IsKnown() bool {
	switch r {
	case HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionTypeFixed, HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionTypePercent:
		return true
	}
	return false
}

type HRISBenefitIndividualEnrolledIDsParams struct {
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
}

// URLQuery serializes [HRISBenefitIndividualEnrolledIDsParams]'s query parameters
// as `url.Values`.
func (r HRISBenefitIndividualEnrolledIDsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISBenefitIndividualGetManyBenefitsParams struct {
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
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
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
	// Array of individual_ids to unenroll.
	IndividualIDs param.Field[[]string] `json:"individual_ids"`
}

func (r HRISBenefitIndividualUnenrollManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISBenefitIndividualUnenrollManyParams]'s query parameters
// as `url.Values`.
func (r HRISBenefitIndividualUnenrollManyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
