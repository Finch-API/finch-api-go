// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxDirectoryService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxDirectoryService] method instead.
type SandboxDirectoryService struct {
	Options []option.RequestOption
}

// NewSandboxDirectoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxDirectoryService(opts ...option.RequestOption) (r *SandboxDirectoryService) {
	r = &SandboxDirectoryService{}
	r.Options = opts
	return
}

// Add new individuals to a sandbox company
func (r *SandboxDirectoryService) New(ctx context.Context, body SandboxDirectoryNewParams, opts ...option.RequestOption) (res *[]SandboxDirectoryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sandbox/directory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SandboxDirectoryNewResponse = interface{}

type SandboxDirectoryNewParams struct {
	// Array of individuals to create. Takes all combined fields from `/individual` and
	// `/employment` endpoints. All fields are optional.
	Body []SandboxDirectoryNewParamsBody `json:"body,required"`
}

func (r SandboxDirectoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SandboxDirectoryNewParamsBody struct {
	// Worker's compensation classification code for this employee
	ClassCode param.Field[string] `json:"class_code"`
	// Custom fields for the individual. These are fields which are defined by the
	// employer in the system. Custom fields are not currently supported for assisted
	// connections.
	CustomFields param.Field[[]SandboxDirectoryNewParamsBodyCustomField] `json:"custom_fields"`
	// The department object.
	Department param.Field[SandboxDirectoryNewParamsBodyDepartment] `json:"department"`
	Dob        param.Field[string]                                  `json:"dob"`
	Emails     param.Field[[]SandboxDirectoryNewParamsBodyEmail]    `json:"emails"`
	// The employment object.
	Employment param.Field[SandboxDirectoryNewParamsBodyEmployment] `json:"employment"`
	// The detailed employment status of the individual.
	EmploymentStatus param.Field[SandboxDirectoryNewParamsBodyEmploymentStatus] `json:"employment_status"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn param.Field[string] `json:"encrypted_ssn"`
	EndDate      param.Field[string] `json:"end_date"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity param.Field[SandboxDirectoryNewParamsBodyEthnicity] `json:"ethnicity"`
	// The legal first name of the individual.
	FirstName param.Field[string] `json:"first_name"`
	// The gender of the individual.
	Gender param.Field[SandboxDirectoryNewParamsBodyGender] `json:"gender"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income param.Field[IncomeParam] `json:"income"`
	// The array of income history.
	IncomeHistory param.Field[[]IncomeParam] `json:"income_history"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive param.Field[bool] `json:"is_active"`
	// The legal last name of the individual.
	LastName         param.Field[string]        `json:"last_name"`
	LatestRehireDate param.Field[string]        `json:"latest_rehire_date"`
	Location         param.Field[LocationParam] `json:"location"`
	// The manager object representing the manager of the individual within the org.
	Manager param.Field[SandboxDirectoryNewParamsBodyManager] `json:"manager"`
	// The legal middle name of the individual.
	MiddleName   param.Field[string]                                     `json:"middle_name"`
	PhoneNumbers param.Field[[]SandboxDirectoryNewParamsBodyPhoneNumber] `json:"phone_numbers"`
	// The preferred name of the individual.
	PreferredName param.Field[string]        `json:"preferred_name"`
	Residence     param.Field[LocationParam] `json:"residence"`
	// The source system's unique employment identifier for this individual
	SourceID param.Field[string] `json:"source_id"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	// [Click here to learn more about enabling the SSN field](/developer-resources/Enable-SSN-Field).
	Ssn       param.Field[string] `json:"ssn"`
	StartDate param.Field[string] `json:"start_date"`
	// The current title of the individual.
	Title param.Field[string] `json:"title"`
}

func (r SandboxDirectoryNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxDirectoryNewParamsBodyCustomField struct {
	Name  param.Field[string]      `json:"name"`
	Value param.Field[interface{}] `json:"value"`
}

func (r SandboxDirectoryNewParamsBodyCustomField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The department object.
type SandboxDirectoryNewParamsBodyDepartment struct {
	// The name of the department associated with the individual.
	Name param.Field[string] `json:"name"`
}

func (r SandboxDirectoryNewParamsBodyDepartment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxDirectoryNewParamsBodyEmail struct {
	Data param.Field[string]                                  `json:"data"`
	Type param.Field[SandboxDirectoryNewParamsBodyEmailsType] `json:"type"`
}

func (r SandboxDirectoryNewParamsBodyEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxDirectoryNewParamsBodyEmailsType string

const (
	SandboxDirectoryNewParamsBodyEmailsTypeWork     SandboxDirectoryNewParamsBodyEmailsType = "work"
	SandboxDirectoryNewParamsBodyEmailsTypePersonal SandboxDirectoryNewParamsBodyEmailsType = "personal"
)

func (r SandboxDirectoryNewParamsBodyEmailsType) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyEmailsTypeWork, SandboxDirectoryNewParamsBodyEmailsTypePersonal:
		return true
	}
	return false
}

// The employment object.
type SandboxDirectoryNewParamsBodyEmployment struct {
	// The secondary employment type of the individual. Options: `full_time`, `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
	Subtype param.Field[SandboxDirectoryNewParamsBodyEmploymentSubtype] `json:"subtype"`
	// The main employment type of the individual.
	Type param.Field[SandboxDirectoryNewParamsBodyEmploymentType] `json:"type"`
}

func (r SandboxDirectoryNewParamsBodyEmployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The secondary employment type of the individual. Options: `full_time`, `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
type SandboxDirectoryNewParamsBodyEmploymentSubtype string

const (
	SandboxDirectoryNewParamsBodyEmploymentSubtypeFullTime             SandboxDirectoryNewParamsBodyEmploymentSubtype = "full_time"
	SandboxDirectoryNewParamsBodyEmploymentSubtypeIntern               SandboxDirectoryNewParamsBodyEmploymentSubtype = "intern"
	SandboxDirectoryNewParamsBodyEmploymentSubtypePartTime             SandboxDirectoryNewParamsBodyEmploymentSubtype = "part_time"
	SandboxDirectoryNewParamsBodyEmploymentSubtypeTemp                 SandboxDirectoryNewParamsBodyEmploymentSubtype = "temp"
	SandboxDirectoryNewParamsBodyEmploymentSubtypeSeasonal             SandboxDirectoryNewParamsBodyEmploymentSubtype = "seasonal"
	SandboxDirectoryNewParamsBodyEmploymentSubtypeIndividualContractor SandboxDirectoryNewParamsBodyEmploymentSubtype = "individual_contractor"
)

func (r SandboxDirectoryNewParamsBodyEmploymentSubtype) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyEmploymentSubtypeFullTime, SandboxDirectoryNewParamsBodyEmploymentSubtypeIntern, SandboxDirectoryNewParamsBodyEmploymentSubtypePartTime, SandboxDirectoryNewParamsBodyEmploymentSubtypeTemp, SandboxDirectoryNewParamsBodyEmploymentSubtypeSeasonal, SandboxDirectoryNewParamsBodyEmploymentSubtypeIndividualContractor:
		return true
	}
	return false
}

// The main employment type of the individual.
type SandboxDirectoryNewParamsBodyEmploymentType string

const (
	SandboxDirectoryNewParamsBodyEmploymentTypeEmployee   SandboxDirectoryNewParamsBodyEmploymentType = "employee"
	SandboxDirectoryNewParamsBodyEmploymentTypeContractor SandboxDirectoryNewParamsBodyEmploymentType = "contractor"
)

func (r SandboxDirectoryNewParamsBodyEmploymentType) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyEmploymentTypeEmployee, SandboxDirectoryNewParamsBodyEmploymentTypeContractor:
		return true
	}
	return false
}

// The detailed employment status of the individual.
type SandboxDirectoryNewParamsBodyEmploymentStatus string

const (
	SandboxDirectoryNewParamsBodyEmploymentStatusActive     SandboxDirectoryNewParamsBodyEmploymentStatus = "active"
	SandboxDirectoryNewParamsBodyEmploymentStatusDeceased   SandboxDirectoryNewParamsBodyEmploymentStatus = "deceased"
	SandboxDirectoryNewParamsBodyEmploymentStatusLeave      SandboxDirectoryNewParamsBodyEmploymentStatus = "leave"
	SandboxDirectoryNewParamsBodyEmploymentStatusOnboarding SandboxDirectoryNewParamsBodyEmploymentStatus = "onboarding"
	SandboxDirectoryNewParamsBodyEmploymentStatusPrehire    SandboxDirectoryNewParamsBodyEmploymentStatus = "prehire"
	SandboxDirectoryNewParamsBodyEmploymentStatusRetired    SandboxDirectoryNewParamsBodyEmploymentStatus = "retired"
	SandboxDirectoryNewParamsBodyEmploymentStatusTerminated SandboxDirectoryNewParamsBodyEmploymentStatus = "terminated"
)

func (r SandboxDirectoryNewParamsBodyEmploymentStatus) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyEmploymentStatusActive, SandboxDirectoryNewParamsBodyEmploymentStatusDeceased, SandboxDirectoryNewParamsBodyEmploymentStatusLeave, SandboxDirectoryNewParamsBodyEmploymentStatusOnboarding, SandboxDirectoryNewParamsBodyEmploymentStatusPrehire, SandboxDirectoryNewParamsBodyEmploymentStatusRetired, SandboxDirectoryNewParamsBodyEmploymentStatusTerminated:
		return true
	}
	return false
}

// The EEOC-defined ethnicity of the individual.
type SandboxDirectoryNewParamsBodyEthnicity string

const (
	SandboxDirectoryNewParamsBodyEthnicityAsian                           SandboxDirectoryNewParamsBodyEthnicity = "asian"
	SandboxDirectoryNewParamsBodyEthnicityWhite                           SandboxDirectoryNewParamsBodyEthnicity = "white"
	SandboxDirectoryNewParamsBodyEthnicityBlackOrAfricanAmerican          SandboxDirectoryNewParamsBodyEthnicity = "black_or_african_american"
	SandboxDirectoryNewParamsBodyEthnicityNativeHawaiianOrPacificIslander SandboxDirectoryNewParamsBodyEthnicity = "native_hawaiian_or_pacific_islander"
	SandboxDirectoryNewParamsBodyEthnicityAmericanIndianOrAlaskaNative    SandboxDirectoryNewParamsBodyEthnicity = "american_indian_or_alaska_native"
	SandboxDirectoryNewParamsBodyEthnicityHispanicOrLatino                SandboxDirectoryNewParamsBodyEthnicity = "hispanic_or_latino"
	SandboxDirectoryNewParamsBodyEthnicityTwoOrMoreRaces                  SandboxDirectoryNewParamsBodyEthnicity = "two_or_more_races"
	SandboxDirectoryNewParamsBodyEthnicityDeclineToSpecify                SandboxDirectoryNewParamsBodyEthnicity = "decline_to_specify"
)

func (r SandboxDirectoryNewParamsBodyEthnicity) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyEthnicityAsian, SandboxDirectoryNewParamsBodyEthnicityWhite, SandboxDirectoryNewParamsBodyEthnicityBlackOrAfricanAmerican, SandboxDirectoryNewParamsBodyEthnicityNativeHawaiianOrPacificIslander, SandboxDirectoryNewParamsBodyEthnicityAmericanIndianOrAlaskaNative, SandboxDirectoryNewParamsBodyEthnicityHispanicOrLatino, SandboxDirectoryNewParamsBodyEthnicityTwoOrMoreRaces, SandboxDirectoryNewParamsBodyEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type SandboxDirectoryNewParamsBodyGender string

const (
	SandboxDirectoryNewParamsBodyGenderFemale           SandboxDirectoryNewParamsBodyGender = "female"
	SandboxDirectoryNewParamsBodyGenderMale             SandboxDirectoryNewParamsBodyGender = "male"
	SandboxDirectoryNewParamsBodyGenderOther            SandboxDirectoryNewParamsBodyGender = "other"
	SandboxDirectoryNewParamsBodyGenderDeclineToSpecify SandboxDirectoryNewParamsBodyGender = "decline_to_specify"
)

func (r SandboxDirectoryNewParamsBodyGender) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyGenderFemale, SandboxDirectoryNewParamsBodyGenderMale, SandboxDirectoryNewParamsBodyGenderOther, SandboxDirectoryNewParamsBodyGenderDeclineToSpecify:
		return true
	}
	return false
}

// The manager object representing the manager of the individual within the org.
type SandboxDirectoryNewParamsBodyManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID param.Field[string] `json:"id"`
}

func (r SandboxDirectoryNewParamsBodyManager) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxDirectoryNewParamsBodyPhoneNumber struct {
	Data param.Field[string]                                        `json:"data"`
	Type param.Field[SandboxDirectoryNewParamsBodyPhoneNumbersType] `json:"type"`
}

func (r SandboxDirectoryNewParamsBodyPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxDirectoryNewParamsBodyPhoneNumbersType string

const (
	SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork     SandboxDirectoryNewParamsBodyPhoneNumbersType = "work"
	SandboxDirectoryNewParamsBodyPhoneNumbersTypePersonal SandboxDirectoryNewParamsBodyPhoneNumbersType = "personal"
)

func (r SandboxDirectoryNewParamsBodyPhoneNumbersType) IsKnown() bool {
	switch r {
	case SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork, SandboxDirectoryNewParamsBodyPhoneNumbersTypePersonal:
		return true
	}
	return false
}
