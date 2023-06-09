// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISBenefitIndividualService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHRISBenefitIndividualService] method
// instead.
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

// **Availability: Automated Benefits providers only**
//
// Lists individuals currently enrolled in a given benefit.
func (r *HRISBenefitIndividualService) EnrolledIDs(ctx context.Context, benefitID string, opts ...option.RequestOption) (res *IndividualEnrolledIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("employer/benefits/%s/enrolled", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// **Availability: Automated Benefits providers only**
//
// Get enrolled benefit information for the given individuals.
func (r *HRISBenefitIndividualService) GetManyBenefits(ctx context.Context, benefitID string, query HRISBenefitIndividualGetManyBenefitsParams, opts ...option.RequestOption) (res *shared.SinglePage[IndividualBenefit], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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

// **Availability: Automated Benefits providers only**
//
// Get enrolled benefit information for the given individuals.
func (r *HRISBenefitIndividualService) GetManyBenefitsAutoPaging(ctx context.Context, benefitID string, query HRISBenefitIndividualGetManyBenefitsParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[IndividualBenefit] {
	return shared.NewSinglePageAutoPager(r.GetManyBenefits(ctx, benefitID, query, opts...))
}

// **Availability: Automated and Assisted Benefits providers**
//
// Unenroll individuals from a benefit
func (r *HRISBenefitIndividualService) Unenroll(ctx context.Context, benefitID string, body HRISBenefitIndividualUnenrollParams, opts ...option.RequestOption) (res *shared.SinglePage[UnenrolledIndividual], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("employer/benefits/%s/individuals", benefitID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, body, &res, opts...)
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
// Unenroll individuals from a benefit
func (r *HRISBenefitIndividualService) UnenrollAutoPaging(ctx context.Context, benefitID string, body HRISBenefitIndividualUnenrollParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[UnenrolledIndividual] {
	return shared.NewSinglePageAutoPager(r.Unenroll(ctx, benefitID, body, opts...))
}

type IndividualBenefit struct {
	IndividualID string                `json:"individual_id"`
	Code         int64                 `json:"code"`
	Body         IndividualBenefitBody `json:"body"`
	JSON         individualBenefitJSON
}

// individualBenefitJSON contains the JSON metadata for the struct
// [IndividualBenefit]
type individualBenefitJSON struct {
	IndividualID apijson.Field
	Code         apijson.Field
	Body         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IndividualBenefit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualBenefitBody struct {
	EmployeeDeduction   BenfitContribution `json:"employee_deduction,nullable"`
	CompanyContribution BenfitContribution `json:"company_contribution,nullable"`
	// If the benefit supports annual maximum, the amount in cents for this individual.
	AnnualMaximum int64 `json:"annual_maximum,nullable"`
	// If the benefit supports catch up (401k, 403b, etc.), whether catch up is enabled
	// for this individual.
	CatchUp bool `json:"catch_up,nullable"`
	// Type for HSA contribution limit if the benefit is a HSA.
	HsaContributionLimit IndividualBenefitBodyHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	JSON                 individualBenefitBodyJSON
}

// individualBenefitBodyJSON contains the JSON metadata for the struct
// [IndividualBenefitBody]
type individualBenefitBodyJSON struct {
	EmployeeDeduction    apijson.Field
	CompanyContribution  apijson.Field
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	HsaContributionLimit apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IndividualBenefitBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualBenefitBodyHsaContributionLimit string

const (
	IndividualBenefitBodyHsaContributionLimitIndividual IndividualBenefitBodyHsaContributionLimit = "individual"
	IndividualBenefitBodyHsaContributionLimitFamily     IndividualBenefitBodyHsaContributionLimit = "family"
)

type UnenrolledIndividual struct {
	IndividualID string `json:"individual_id"`
	// HTTP status code
	Code int64                    `json:"code"`
	Body UnenrolledIndividualBody `json:"body"`
	JSON unenrolledIndividualJSON
}

// unenrolledIndividualJSON contains the JSON metadata for the struct
// [UnenrolledIndividual]
type unenrolledIndividualJSON struct {
	IndividualID apijson.Field
	Code         apijson.Field
	Body         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UnenrolledIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UnenrolledIndividualBody struct {
	// Identifier indicating whether the benefit was newly enrolled or updated.
	Name string `json:"name,nullable"`
	// A descriptive identifier for the response.
	FinchCode string `json:"finch_code,nullable"`
	// Short description in English that provides more information about the response.
	Message string `json:"message,nullable"`
	JSON    unenrolledIndividualBodyJSON
}

// unenrolledIndividualBodyJSON contains the JSON metadata for the struct
// [UnenrolledIndividualBody]
type unenrolledIndividualBodyJSON struct {
	Name        apijson.Field
	FinchCode   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnenrolledIndividualBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualEnrolledIDsResponse struct {
	BenefitID     string   `json:"benefit_id,required"`
	IndividualIDs []string `json:"individual_ids,required"`
	JSON          individualEnrolledIDsResponseJSON
}

// individualEnrolledIDsResponseJSON contains the JSON metadata for the struct
// [IndividualEnrolledIDsResponse]
type individualEnrolledIDsResponseJSON struct {
	BenefitID     apijson.Field
	IndividualIDs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IndividualEnrolledIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISBenefitIndividualUnenrollParams struct {
	// Array of individual_ids to unenroll.
	IndividualIDs param.Field[[]string] `json:"individual_ids"`
}

func (r HRISBenefitIndividualUnenrollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
