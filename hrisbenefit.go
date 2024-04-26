// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/pagination"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/shared"
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

// Creates a new company-wide deduction or contribution. Please use the
// `/providers` endpoint to view available types for each provider.
func (r *HRISBenefitService) New(ctx context.Context, body HRISBenefitNewParams, opts ...option.RequestOption) (res *CreateCompanyBenefitsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "employer/benefits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists deductions and contributions information for a given item
func (r *HRISBenefitService) Get(ctx context.Context, benefitID string, opts ...option.RequestOption) (res *CompanyBenefit, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("employer/benefits/%s", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing company-wide deduction or contribution
func (r *HRISBenefitService) Update(ctx context.Context, benefitID string, body HRISBenefitUpdateParams, opts ...option.RequestOption) (res *UpdateCompanyBenefitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("employer/benefits/%s", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all company-wide deductions and contributions.
func (r *HRISBenefitService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[CompanyBenefit], err error) {
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

// List all company-wide deductions and contributions.
func (r *HRISBenefitService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CompanyBenefit] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// Get deductions metadata
func (r *HRISBenefitService) ListSupportedBenefits(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[SupportedBenefit], err error) {
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

// Get deductions metadata
func (r *HRISBenefitService) ListSupportedBenefitsAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SupportedBenefit] {
	return pagination.NewSinglePageAutoPager(r.ListSupportedBenefits(ctx, opts...))
}

type BenefitContribution struct {
	// Contribution amount in cents (if `fixed`) or basis points (if `percent`).
	Amount int64 `json:"amount,nullable"`
	// Contribution type.
	Type BenefitContributionType `json:"type,nullable"`
	JSON benefitContributionJSON `json:"-"`
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

func (r benefitContributionJSON) RawJSON() string {
	return r.raw
}

// Contribution type.
type BenefitContributionType string

const (
	BenefitContributionTypeFixed   BenefitContributionType = "fixed"
	BenefitContributionTypePercent BenefitContributionType = "percent"
)

func (r BenefitContributionType) IsKnown() bool {
	switch r {
	case BenefitContributionTypeFixed, BenefitContributionTypePercent:
		return true
	}
	return false
}

type BenefitFeaturesAndOperations struct {
	SupportedFeatures   BenefitFeaturesAndOperationsSupportedFeatures `json:"supported_features"`
	SupportedOperations SupportPerBenefitType                         `json:"supported_operations"`
	JSON                benefitFeaturesAndOperationsJSON              `json:"-"`
}

// benefitFeaturesAndOperationsJSON contains the JSON metadata for the struct
// [BenefitFeaturesAndOperations]
type benefitFeaturesAndOperationsJSON struct {
	SupportedFeatures   apijson.Field
	SupportedOperations apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BenefitFeaturesAndOperations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitFeaturesAndOperationsJSON) RawJSON() string {
	return r.raw
}

type BenefitFeaturesAndOperationsSupportedFeatures struct {
	// Whether the provider supports an annual maximum for this benefit.
	AnnualMaximum bool `json:"annual_maximum,nullable"`
	// Whether the provider supports catch up for this benefit. This field will only be
	// true for retirement benefits.
	CatchUp bool `json:"catch_up,nullable"`
	// Supported contribution types. An empty array indicates contributions are not
	// supported.
	CompanyContribution []BenefitFeaturesAndOperationsSupportedFeaturesCompanyContribution `json:"company_contribution,nullable"`
	Description         string                                                             `json:"description,nullable"`
	// Supported deduction types. An empty array indicates deductions are not
	// supported.
	EmployeeDeduction []BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeduction `json:"employee_deduction,nullable"`
	// The list of frequencies supported by the provider for this benefit
	Frequencies []BenefitFrequency `json:"frequencies"`
	// Whether the provider supports HSA contribution limits. Empty if this feature is
	// not supported for the benefit. This array only has values for HSA benefits.
	HsaContributionLimit []BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	JSON                 benefitFeaturesAndOperationsSupportedFeaturesJSON                   `json:"-"`
}

// benefitFeaturesAndOperationsSupportedFeaturesJSON contains the JSON metadata for
// the struct [BenefitFeaturesAndOperationsSupportedFeatures]
type benefitFeaturesAndOperationsSupportedFeaturesJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	CompanyContribution  apijson.Field
	Description          apijson.Field
	EmployeeDeduction    apijson.Field
	Frequencies          apijson.Field
	HsaContributionLimit apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BenefitFeaturesAndOperationsSupportedFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitFeaturesAndOperationsSupportedFeaturesJSON) RawJSON() string {
	return r.raw
}

type BenefitFeaturesAndOperationsSupportedFeaturesCompanyContribution string

const (
	BenefitFeaturesAndOperationsSupportedFeaturesCompanyContributionFixed   BenefitFeaturesAndOperationsSupportedFeaturesCompanyContribution = "fixed"
	BenefitFeaturesAndOperationsSupportedFeaturesCompanyContributionPercent BenefitFeaturesAndOperationsSupportedFeaturesCompanyContribution = "percent"
)

func (r BenefitFeaturesAndOperationsSupportedFeaturesCompanyContribution) IsKnown() bool {
	switch r {
	case BenefitFeaturesAndOperationsSupportedFeaturesCompanyContributionFixed, BenefitFeaturesAndOperationsSupportedFeaturesCompanyContributionPercent:
		return true
	}
	return false
}

type BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeduction string

const (
	BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeductionFixed   BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeduction = "fixed"
	BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeductionPercent BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeduction = "percent"
)

func (r BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeduction) IsKnown() bool {
	switch r {
	case BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeductionFixed, BenefitFeaturesAndOperationsSupportedFeaturesEmployeeDeductionPercent:
		return true
	}
	return false
}

type BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimit string

const (
	BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimitIndividual BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimit = "individual"
	BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimitFamily     BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimit = "family"
)

func (r BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimit) IsKnown() bool {
	switch r {
	case BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimitIndividual, BenefitFeaturesAndOperationsSupportedFeaturesHsaContributionLimitFamily:
		return true
	}
	return false
}

type BenefitFrequency string

const (
	BenefitFrequencyOneTime       BenefitFrequency = "one_time"
	BenefitFrequencyEveryPaycheck BenefitFrequency = "every_paycheck"
	BenefitFrequencyMonthly       BenefitFrequency = "monthly"
)

func (r BenefitFrequency) IsKnown() bool {
	switch r {
	case BenefitFrequencyOneTime, BenefitFrequencyEveryPaycheck, BenefitFrequencyMonthly:
		return true
	}
	return false
}

// Type of benefit.
type BenefitType string

const (
	BenefitType_401k            BenefitType = "401k"
	BenefitType_401kRoth        BenefitType = "401k_roth"
	BenefitType_401kLoan        BenefitType = "401k_loan"
	BenefitType_403b            BenefitType = "403b"
	BenefitType_403bRoth        BenefitType = "403b_roth"
	BenefitType_457             BenefitType = "457"
	BenefitType_457Roth         BenefitType = "457_roth"
	BenefitTypeS125Medical      BenefitType = "s125_medical"
	BenefitTypeS125Dental       BenefitType = "s125_dental"
	BenefitTypeS125Vision       BenefitType = "s125_vision"
	BenefitTypeHsaPre           BenefitType = "hsa_pre"
	BenefitTypeHsaPost          BenefitType = "hsa_post"
	BenefitTypeFsaMedical       BenefitType = "fsa_medical"
	BenefitTypeFsaDependentCare BenefitType = "fsa_dependent_care"
	BenefitTypeSimpleIRA        BenefitType = "simple_ira"
	BenefitTypeSimple           BenefitType = "simple"
	BenefitTypeCommuter         BenefitType = "commuter"
	BenefitTypeCustomPostTax    BenefitType = "custom_post_tax"
	BenefitTypeCustomPreTax     BenefitType = "custom_pre_tax"
)

func (r BenefitType) IsKnown() bool {
	switch r {
	case BenefitType_401k, BenefitType_401kRoth, BenefitType_401kLoan, BenefitType_403b, BenefitType_403bRoth, BenefitType_457, BenefitType_457Roth, BenefitTypeS125Medical, BenefitTypeS125Dental, BenefitTypeS125Vision, BenefitTypeHsaPre, BenefitTypeHsaPost, BenefitTypeFsaMedical, BenefitTypeFsaDependentCare, BenefitTypeSimpleIRA, BenefitTypeSimple, BenefitTypeCommuter, BenefitTypeCustomPostTax, BenefitTypeCustomPreTax:
		return true
	}
	return false
}

// Each benefit type and their supported features. If the benefit type is not
// supported, the property will be null
type BenefitsSupport struct {
	Commuter         BenefitFeaturesAndOperations            `json:"commuter,nullable"`
	CustomPostTax    BenefitFeaturesAndOperations            `json:"custom_post_tax,nullable"`
	CustomPreTax     BenefitFeaturesAndOperations            `json:"custom_pre_tax,nullable"`
	FsaDependentCare BenefitFeaturesAndOperations            `json:"fsa_dependent_care,nullable"`
	FsaMedical       BenefitFeaturesAndOperations            `json:"fsa_medical,nullable"`
	HsaPost          BenefitFeaturesAndOperations            `json:"hsa_post,nullable"`
	HsaPre           BenefitFeaturesAndOperations            `json:"hsa_pre,nullable"`
	S125Dental       BenefitFeaturesAndOperations            `json:"s125_dental,nullable"`
	S125Medical      BenefitFeaturesAndOperations            `json:"s125_medical,nullable"`
	S125Vision       BenefitFeaturesAndOperations            `json:"s125_vision,nullable"`
	Simple           BenefitFeaturesAndOperations            `json:"simple,nullable"`
	SimpleIRA        BenefitFeaturesAndOperations            `json:"simple_ira,nullable"`
	ExtraFields      map[string]BenefitFeaturesAndOperations `json:"-,extras"`
	JSON             benefitsSupportJSON                     `json:"-"`
}

// benefitsSupportJSON contains the JSON metadata for the struct [BenefitsSupport]
type benefitsSupportJSON struct {
	Commuter         apijson.Field
	CustomPostTax    apijson.Field
	CustomPreTax     apijson.Field
	FsaDependentCare apijson.Field
	FsaMedical       apijson.Field
	HsaPost          apijson.Field
	HsaPre           apijson.Field
	S125Dental       apijson.Field
	S125Medical      apijson.Field
	S125Vision       apijson.Field
	Simple           apijson.Field
	SimpleIRA        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BenefitsSupport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r benefitsSupportJSON) RawJSON() string {
	return r.raw
}

type CompanyBenefit struct {
	BenefitID           string              `json:"benefit_id,required"`
	CompanyContribution BenefitContribution `json:"company_contribution,required,nullable"`
	Description         string              `json:"description,required,nullable"`
	EmployeeDeduction   BenefitContribution `json:"employee_deduction,required,nullable"`
	Frequency           BenefitFrequency    `json:"frequency,required,nullable"`
	// Type of benefit.
	Type BenefitType        `json:"type,required,nullable"`
	JSON companyBenefitJSON `json:"-"`
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

func (r companyBenefitJSON) RawJSON() string {
	return r.raw
}

type CreateCompanyBenefitsResponse struct {
	BenefitID string                            `json:"benefit_id,required"`
	JSON      createCompanyBenefitsResponseJSON `json:"-"`
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

func (r createCompanyBenefitsResponseJSON) RawJSON() string {
	return r.raw
}

type SupportPerBenefitType struct {
	CompanyBenefits    shared.OperationSupportMatrix `json:"company_benefits"`
	IndividualBenefits shared.OperationSupportMatrix `json:"individual_benefits"`
	JSON               supportPerBenefitTypeJSON     `json:"-"`
}

// supportPerBenefitTypeJSON contains the JSON metadata for the struct
// [SupportPerBenefitType]
type supportPerBenefitTypeJSON struct {
	CompanyBenefits    apijson.Field
	IndividualBenefits apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SupportPerBenefitType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r supportPerBenefitTypeJSON) RawJSON() string {
	return r.raw
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
	Type BenefitType          `json:"type,nullable"`
	JSON supportedBenefitJSON `json:"-"`
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

func (r supportedBenefitJSON) RawJSON() string {
	return r.raw
}

type SupportedBenefitCompanyContribution string

const (
	SupportedBenefitCompanyContributionFixed   SupportedBenefitCompanyContribution = "fixed"
	SupportedBenefitCompanyContributionPercent SupportedBenefitCompanyContribution = "percent"
)

func (r SupportedBenefitCompanyContribution) IsKnown() bool {
	switch r {
	case SupportedBenefitCompanyContributionFixed, SupportedBenefitCompanyContributionPercent:
		return true
	}
	return false
}

type SupportedBenefitEmployeeDeduction string

const (
	SupportedBenefitEmployeeDeductionFixed   SupportedBenefitEmployeeDeduction = "fixed"
	SupportedBenefitEmployeeDeductionPercent SupportedBenefitEmployeeDeduction = "percent"
)

func (r SupportedBenefitEmployeeDeduction) IsKnown() bool {
	switch r {
	case SupportedBenefitEmployeeDeductionFixed, SupportedBenefitEmployeeDeductionPercent:
		return true
	}
	return false
}

type SupportedBenefitHsaContributionLimit string

const (
	SupportedBenefitHsaContributionLimitIndividual SupportedBenefitHsaContributionLimit = "individual"
	SupportedBenefitHsaContributionLimitFamily     SupportedBenefitHsaContributionLimit = "family"
)

func (r SupportedBenefitHsaContributionLimit) IsKnown() bool {
	switch r {
	case SupportedBenefitHsaContributionLimitIndividual, SupportedBenefitHsaContributionLimitFamily:
		return true
	}
	return false
}

type UpdateCompanyBenefitResponse struct {
	BenefitID string                           `json:"benefit_id,required"`
	JSON      updateCompanyBenefitResponseJSON `json:"-"`
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

func (r updateCompanyBenefitResponseJSON) RawJSON() string {
	return r.raw
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
