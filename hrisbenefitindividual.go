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
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	if benefitID == "" {
		err = errors.New("missing required benefit_id parameter")
		return
	}
	path := fmt.Sprintf("employer/benefits/%s/individuals", benefitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type IndividualBenefit struct {
	Body         IndividualBenefitBody `json:"body"`
	Code         int64                 `json:"code"`
	IndividualID string                `json:"individual_id"`
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
	CatchUp             bool                `json:"catch_up,nullable"`
	CompanyContribution BenefitContribution `json:"company_contribution,nullable"`
	EmployeeDeduction   BenefitContribution `json:"employee_deduction,nullable"`
	// Type for HSA contribution limit if the benefit is a HSA.
	HsaContributionLimit IndividualBenefitBodyHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	JSON                 individualBenefitBodyJSON                 `json:"-"`
}

// individualBenefitBodyJSON contains the JSON metadata for the struct
// [IndividualBenefitBody]
type individualBenefitBodyJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	CompanyContribution  apijson.Field
	EmployeeDeduction    apijson.Field
	HsaContributionLimit apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IndividualBenefitBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBenefitBodyJSON) RawJSON() string {
	return r.raw
}

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
