// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISBenefitService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHRISBenefitService] method instead.
type HRISBenefitService struct {
	Options     []option.RequestOption
	Individuals *HRISBenefitIndividualService
}

// NewHRISBenefitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISBenefitService(opts ...option.RequestOption) (r *HRISBenefitService) {
	r = &HRISBenefitService{}
	r.Options = opts
	r.Individuals = NewHRISBenefitIndividualService(opts...)
	return
}

// **Availability: Automated and Assisted Benefits providers**
//
// Creates a new company-wide benefit. Please use the `/meta` endpoint to view
// available types for each provider.
func (r *HRISBenefitService) New(ctx context.Context, body HRISBenefitNewParams, opts ...option.RequestOption) (res *CreateCompanyBenefitsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "employer/benefits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// **Availability: Automated Benefits providers only**
//
// Lists benefit information for a given benefit
func (r *HRISBenefitService) Get(ctx context.Context, benefitID string, opts ...option.RequestOption) (res *CompanyBenefit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("employer/benefits/%s", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// **Availability: Automated and Assisted Benefits providers**
//
// Updates an existing company-wide benefit
func (r *HRISBenefitService) Update(ctx context.Context, benefitID string, body HRISBenefitUpdateParams, opts ...option.RequestOption) (res *UpdateCompanyBenefitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("employer/benefits/%s", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// **Availability: Automated Benefits providers only**
//
// List all company-wide benefits.
func (r *HRISBenefitService) List(ctx context.Context, opts ...option.RequestOption) (res *shared.SinglePage[CompanyBenefit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/benefits"
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

// **Availability: Automated Benefits providers only**
//
// List all company-wide benefits.
func (r *HRISBenefitService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *shared.SinglePageAutoPager[CompanyBenefit] {
	return shared.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// **Availability: Automated and Assisted Benefits providers**
//
// Lists available types and configurations for the provider associated with the
// access token.
func (r *HRISBenefitService) ListSupportedBenefits(ctx context.Context, opts ...option.RequestOption) (res *shared.SinglePage[SupportedBenefit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/benefits/meta"
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

// **Availability: Automated and Assisted Benefits providers**
//
// Lists available types and configurations for the provider associated with the
// access token.
func (r *HRISBenefitService) ListSupportedBenefitsAutoPaging(ctx context.Context, opts ...option.RequestOption) *shared.SinglePageAutoPager[SupportedBenefit] {
	return shared.NewSinglePageAutoPager(r.ListSupportedBenefits(ctx, opts...))
}

type BenefitContribution struct {
	// Contribution amount in cents (if `fixed`) or basis points (if `percent`).
	Amount int64 `json:"amount,nullable"`
	// Contribution type.
	Type BenefitContributionType `json:"type,nullable"`
	JSON benefitContributionJSON
}

// benefitContributionJSON contains the JSON metadata for the struct
// [BenefitContribution]
type benefitContributionJSON struct {
	Amount      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitContribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Contribution type.
type BenefitContributionType string

const (
	BenefitContributionTypeFixed   BenefitContributionType = "fixed"
	BenefitContributionTypePercent BenefitContributionType = "percent"
)

type BenefitFrequency string

const (
	BenefitFrequencyOneTime       BenefitFrequency = "one_time"
	BenefitFrequencyEveryPaycheck BenefitFrequency = "every_paycheck"
)

// Type of benefit.
type BenefitType string

const (
	BenefitType401k             BenefitType = "401k"
	BenefitType401kRoth         BenefitType = "401k_roth"
	BenefitType401kLoan         BenefitType = "401k_loan"
	BenefitType403b             BenefitType = "403b"
	BenefitType403bRoth         BenefitType = "403b_roth"
	BenefitType457              BenefitType = "457"
	BenefitType457Roth          BenefitType = "457_roth"
	BenefitTypeS125Medical      BenefitType = "s125_medical"
	BenefitTypeS125Dental       BenefitType = "s125_dental"
	BenefitTypeS125Vision       BenefitType = "s125_vision"
	BenefitTypeHsaPre           BenefitType = "hsa_pre"
	BenefitTypeHsaPost          BenefitType = "hsa_post"
	BenefitTypeFsaMedical       BenefitType = "fsa_medical"
	BenefitTypeFsaDependentCare BenefitType = "fsa_dependent_care"
	BenefitTypeSimpleIra        BenefitType = "simple_ira"
	BenefitTypeSimple           BenefitType = "simple"
	BenefitTypeCommuter         BenefitType = "commuter"
	BenefitTypeCustomPostTax    BenefitType = "custom_post_tax"
	BenefitTypeCustomPreTax     BenefitType = "custom_pre_tax"
)

type CompanyBenefit struct {
	BenefitID           string              `json:"benefit_id,required"`
	CompanyContribution BenefitContribution `json:"company_contribution,required,nullable"`
	Description         string              `json:"description,required,nullable"`
	EmployeeDeduction   BenefitContribution `json:"employee_deduction,required,nullable"`
	Frequency           BenefitFrequency    `json:"frequency,required,nullable"`
	// Type of benefit.
	Type BenefitType `json:"type,required,nullable"`
	JSON companyBenefitJSON
}

// companyBenefitJSON contains the JSON metadata for the struct [CompanyBenefit]
type companyBenefitJSON struct {
	BenefitID           apijson.Field
	CompanyContribution apijson.Field
	Description         apijson.Field
	EmployeeDeduction   apijson.Field
	Frequency           apijson.Field
	Type                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CompanyBenefit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateCompanyBenefitsResponse struct {
	BenefitID string `json:"benefit_id,required"`
	JSON      createCompanyBenefitsResponseJSON
}

// createCompanyBenefitsResponseJSON contains the JSON metadata for the struct
// [CreateCompanyBenefitsResponse]
type createCompanyBenefitsResponseJSON struct {
	BenefitID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateCompanyBenefitsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SupportedBenefit struct {
	// Whether the provider supports an annual maximum for this benefit.
	AnnualMaximum bool `json:"annual_maximum,nullable"`
	// Whether the provider supports catch up for this benefit. This field will only be
	// true for retirement benefits.
	CatchUp bool `json:"catch_up,nullable"`
	// Supported contribution types. An empty array indicates contributions are not
	// supported.
	CompanyContribution []SupportedBenefitCompanyContribution `json:"company_contribution,nullable"`
	Description         string                                `json:"description,nullable"`
	// Supported deduction types. An empty array indicates deductions are not
	// supported.
	EmployeeDeduction []SupportedBenefitEmployeeDeduction `json:"employee_deduction,nullable"`
	// The list of frequencies supported by the provider for this benefit
	Frequencies []BenefitFrequency `json:"frequencies"`
	// Whether the provider supports HSA contribution limits. Empty if this feature is
	// not supported for the benefit. This array only has values for HSA benefits.
	HsaContributionLimit []SupportedBenefitHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	// Type of benefit.
	Type BenefitType `json:"type,nullable"`
	JSON supportedBenefitJSON
}

// supportedBenefitJSON contains the JSON metadata for the struct
// [SupportedBenefit]
type supportedBenefitJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	CompanyContribution  apijson.Field
	Description          apijson.Field
	EmployeeDeduction    apijson.Field
	Frequencies          apijson.Field
	HsaContributionLimit apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SupportedBenefit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SupportedBenefitCompanyContribution string

const (
	SupportedBenefitCompanyContributionFixed   SupportedBenefitCompanyContribution = "fixed"
	SupportedBenefitCompanyContributionPercent SupportedBenefitCompanyContribution = "percent"
)

type SupportedBenefitEmployeeDeduction string

const (
	SupportedBenefitEmployeeDeductionFixed   SupportedBenefitEmployeeDeduction = "fixed"
	SupportedBenefitEmployeeDeductionPercent SupportedBenefitEmployeeDeduction = "percent"
)

type SupportedBenefitHsaContributionLimit string

const (
	SupportedBenefitHsaContributionLimitIndividual SupportedBenefitHsaContributionLimit = "individual"
	SupportedBenefitHsaContributionLimitFamily     SupportedBenefitHsaContributionLimit = "family"
)

type UpdateCompanyBenefitResponse struct {
	BenefitID string `json:"benefit_id,required"`
	JSON      updateCompanyBenefitResponseJSON
}

// updateCompanyBenefitResponseJSON contains the JSON metadata for the struct
// [UpdateCompanyBenefitResponse]
type updateCompanyBenefitResponseJSON struct {
	BenefitID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateCompanyBenefitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HRISBenefitNewParams struct {
	Description param.Field[string]           `json:"description"`
	Frequency   param.Field[BenefitFrequency] `json:"frequency"`
	// Type of benefit.
	Type param.Field[BenefitType] `json:"type"`
}

func (r HRISBenefitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISBenefitUpdateParams struct {
	// Updated name or description.
	Description param.Field[string] `json:"description"`
}

func (r HRISBenefitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
