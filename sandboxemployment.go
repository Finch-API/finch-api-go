// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxEmploymentService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxEmploymentService] method instead.
type SandboxEmploymentService struct {
	Options []option.RequestOption
}

// NewSandboxEmploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxEmploymentService(opts ...option.RequestOption) (r *SandboxEmploymentService) {
	r = &SandboxEmploymentService{}
	r.Options = opts
	return
}

// Update sandbox employment
func (r *SandboxEmploymentService) Update(ctx context.Context, individualID string, body SandboxEmploymentUpdateParams, opts ...option.RequestOption) (res *SandboxEmploymentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if individualID == "" {
		err = errors.New("missing required individual_id parameter")
		return
	}
	path := fmt.Sprintf("sandbox/employment/%s", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SandboxEmploymentUpdateResponse struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id" format:"uuid"`
	// Worker's compensation classification code for this employee
	ClassCode string `json:"class_code,nullable"`
	// Custom fields for the individual. These are fields which are defined by the
	// employer in the system. Custom fields are not currently supported for assisted
	// connections.
	CustomFields []SandboxEmploymentUpdateResponseCustomField `json:"custom_fields,nullable"`
	// The department object.
	Department SandboxEmploymentUpdateResponseDepartment `json:"department,nullable"`
	// The employment object.
	Employment SandboxEmploymentUpdateResponseEmployment `json:"employment,nullable"`
	// The detailed employment status of the individual. Available options: `active`,
	// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
	EmploymentStatus SandboxEmploymentUpdateResponseEmploymentStatus `json:"employment_status,nullable"`
	EndDate          string                                          `json:"end_date,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income Income `json:"income,nullable"`
	// The array of income history.
	IncomeHistory []Income `json:"income_history,nullable"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive bool `json:"is_active,nullable"`
	// The legal last name of the individual.
	LastName         string   `json:"last_name,nullable"`
	LatestRehireDate string   `json:"latest_rehire_date,nullable"`
	Location         Location `json:"location,nullable"`
	// The manager object representing the manager of the individual within the org.
	Manager SandboxEmploymentUpdateResponseManager `json:"manager,nullable"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name,nullable"`
	// The source system's unique employment identifier for this individual
	SourceID  string `json:"source_id,nullable"`
	StartDate string `json:"start_date,nullable"`
	// The current title of the individual.
	Title string                              `json:"title,nullable"`
	JSON  sandboxEmploymentUpdateResponseJSON `json:"-"`
}

// sandboxEmploymentUpdateResponseJSON contains the JSON metadata for the struct
// [SandboxEmploymentUpdateResponse]
type sandboxEmploymentUpdateResponseJSON struct {
	ID               apijson.Field
	ClassCode        apijson.Field
	CustomFields     apijson.Field
	Department       apijson.Field
	Employment       apijson.Field
	EmploymentStatus apijson.Field
	EndDate          apijson.Field
	FirstName        apijson.Field
	Income           apijson.Field
	IncomeHistory    apijson.Field
	IsActive         apijson.Field
	LastName         apijson.Field
	LatestRehireDate apijson.Field
	Location         apijson.Field
	Manager          apijson.Field
	MiddleName       apijson.Field
	SourceID         apijson.Field
	StartDate        apijson.Field
	Title            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SandboxEmploymentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxEmploymentUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxEmploymentUpdateResponseCustomField struct {
	Name  string                                         `json:"name,nullable"`
	Value interface{}                                    `json:"value"`
	JSON  sandboxEmploymentUpdateResponseCustomFieldJSON `json:"-"`
}

// sandboxEmploymentUpdateResponseCustomFieldJSON contains the JSON metadata for
// the struct [SandboxEmploymentUpdateResponseCustomField]
type sandboxEmploymentUpdateResponseCustomFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxEmploymentUpdateResponseCustomField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxEmploymentUpdateResponseCustomFieldJSON) RawJSON() string {
	return r.raw
}

// The department object.
type SandboxEmploymentUpdateResponseDepartment struct {
	// The name of the department associated with the individual.
	Name string                                        `json:"name,nullable"`
	JSON sandboxEmploymentUpdateResponseDepartmentJSON `json:"-"`
}

// sandboxEmploymentUpdateResponseDepartmentJSON contains the JSON metadata for the
// struct [SandboxEmploymentUpdateResponseDepartment]
type sandboxEmploymentUpdateResponseDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxEmploymentUpdateResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxEmploymentUpdateResponseDepartmentJSON) RawJSON() string {
	return r.raw
}

// The employment object.
type SandboxEmploymentUpdateResponseEmployment struct {
	// The secondary employment type of the individual. Options: `full_time`,
	// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
	Subtype SandboxEmploymentUpdateResponseEmploymentSubtype `json:"subtype,nullable"`
	// The main employment type of the individual.
	Type SandboxEmploymentUpdateResponseEmploymentType `json:"type,nullable"`
	JSON sandboxEmploymentUpdateResponseEmploymentJSON `json:"-"`
}

// sandboxEmploymentUpdateResponseEmploymentJSON contains the JSON metadata for the
// struct [SandboxEmploymentUpdateResponseEmployment]
type sandboxEmploymentUpdateResponseEmploymentJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxEmploymentUpdateResponseEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxEmploymentUpdateResponseEmploymentJSON) RawJSON() string {
	return r.raw
}

// The secondary employment type of the individual. Options: `full_time`,
// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
type SandboxEmploymentUpdateResponseEmploymentSubtype string

const (
	SandboxEmploymentUpdateResponseEmploymentSubtypeFullTime             SandboxEmploymentUpdateResponseEmploymentSubtype = "full_time"
	SandboxEmploymentUpdateResponseEmploymentSubtypeIntern               SandboxEmploymentUpdateResponseEmploymentSubtype = "intern"
	SandboxEmploymentUpdateResponseEmploymentSubtypePartTime             SandboxEmploymentUpdateResponseEmploymentSubtype = "part_time"
	SandboxEmploymentUpdateResponseEmploymentSubtypeTemp                 SandboxEmploymentUpdateResponseEmploymentSubtype = "temp"
	SandboxEmploymentUpdateResponseEmploymentSubtypeSeasonal             SandboxEmploymentUpdateResponseEmploymentSubtype = "seasonal"
	SandboxEmploymentUpdateResponseEmploymentSubtypeIndividualContractor SandboxEmploymentUpdateResponseEmploymentSubtype = "individual_contractor"
)

func (r SandboxEmploymentUpdateResponseEmploymentSubtype) IsKnown() bool {
	switch r {
	case SandboxEmploymentUpdateResponseEmploymentSubtypeFullTime, SandboxEmploymentUpdateResponseEmploymentSubtypeIntern, SandboxEmploymentUpdateResponseEmploymentSubtypePartTime, SandboxEmploymentUpdateResponseEmploymentSubtypeTemp, SandboxEmploymentUpdateResponseEmploymentSubtypeSeasonal, SandboxEmploymentUpdateResponseEmploymentSubtypeIndividualContractor:
		return true
	}
	return false
}

// The main employment type of the individual.
type SandboxEmploymentUpdateResponseEmploymentType string

const (
	SandboxEmploymentUpdateResponseEmploymentTypeEmployee   SandboxEmploymentUpdateResponseEmploymentType = "employee"
	SandboxEmploymentUpdateResponseEmploymentTypeContractor SandboxEmploymentUpdateResponseEmploymentType = "contractor"
)

func (r SandboxEmploymentUpdateResponseEmploymentType) IsKnown() bool {
	switch r {
	case SandboxEmploymentUpdateResponseEmploymentTypeEmployee, SandboxEmploymentUpdateResponseEmploymentTypeContractor:
		return true
	}
	return false
}

// The detailed employment status of the individual. Available options: `active`,
// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
type SandboxEmploymentUpdateResponseEmploymentStatus string

const (
	SandboxEmploymentUpdateResponseEmploymentStatusActive     SandboxEmploymentUpdateResponseEmploymentStatus = "active"
	SandboxEmploymentUpdateResponseEmploymentStatusDeceased   SandboxEmploymentUpdateResponseEmploymentStatus = "deceased"
	SandboxEmploymentUpdateResponseEmploymentStatusLeave      SandboxEmploymentUpdateResponseEmploymentStatus = "leave"
	SandboxEmploymentUpdateResponseEmploymentStatusOnboarding SandboxEmploymentUpdateResponseEmploymentStatus = "onboarding"
	SandboxEmploymentUpdateResponseEmploymentStatusPrehire    SandboxEmploymentUpdateResponseEmploymentStatus = "prehire"
	SandboxEmploymentUpdateResponseEmploymentStatusRetired    SandboxEmploymentUpdateResponseEmploymentStatus = "retired"
	SandboxEmploymentUpdateResponseEmploymentStatusTerminated SandboxEmploymentUpdateResponseEmploymentStatus = "terminated"
)

func (r SandboxEmploymentUpdateResponseEmploymentStatus) IsKnown() bool {
	switch r {
	case SandboxEmploymentUpdateResponseEmploymentStatusActive, SandboxEmploymentUpdateResponseEmploymentStatusDeceased, SandboxEmploymentUpdateResponseEmploymentStatusLeave, SandboxEmploymentUpdateResponseEmploymentStatusOnboarding, SandboxEmploymentUpdateResponseEmploymentStatusPrehire, SandboxEmploymentUpdateResponseEmploymentStatusRetired, SandboxEmploymentUpdateResponseEmploymentStatusTerminated:
		return true
	}
	return false
}

// The manager object representing the manager of the individual within the org.
type SandboxEmploymentUpdateResponseManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID   string                                     `json:"id" format:"uuid"`
	JSON sandboxEmploymentUpdateResponseManagerJSON `json:"-"`
}

// sandboxEmploymentUpdateResponseManagerJSON contains the JSON metadata for the
// struct [SandboxEmploymentUpdateResponseManager]
type sandboxEmploymentUpdateResponseManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxEmploymentUpdateResponseManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxEmploymentUpdateResponseManagerJSON) RawJSON() string {
	return r.raw
}

type SandboxEmploymentUpdateParams struct {
	// Worker's compensation classification code for this employee
	ClassCode param.Field[string] `json:"class_code"`
	// Custom fields for the individual. These are fields which are defined by the
	// employer in the system. Custom fields are not currently supported for assisted
	// connections.
	CustomFields param.Field[[]SandboxEmploymentUpdateParamsCustomField] `json:"custom_fields"`
	// The department object.
	Department param.Field[SandboxEmploymentUpdateParamsDepartment] `json:"department"`
	// The employment object.
	Employment param.Field[SandboxEmploymentUpdateParamsEmployment] `json:"employment"`
	// The detailed employment status of the individual. Available options: `active`,
	// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
	EmploymentStatus param.Field[SandboxEmploymentUpdateParamsEmploymentStatus] `json:"employment_status"`
	EndDate          param.Field[string]                                        `json:"end_date"`
	// The legal first name of the individual.
	FirstName param.Field[string] `json:"first_name"`
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
	Manager param.Field[SandboxEmploymentUpdateParamsManager] `json:"manager"`
	// The legal middle name of the individual.
	MiddleName param.Field[string] `json:"middle_name"`
	// The source system's unique employment identifier for this individual
	SourceID  param.Field[string] `json:"source_id"`
	StartDate param.Field[string] `json:"start_date"`
	// The current title of the individual.
	Title param.Field[string] `json:"title"`
}

func (r SandboxEmploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxEmploymentUpdateParamsCustomField struct {
	Name  param.Field[string]      `json:"name"`
	Value param.Field[interface{}] `json:"value"`
}

func (r SandboxEmploymentUpdateParamsCustomField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The department object.
type SandboxEmploymentUpdateParamsDepartment struct {
	// The name of the department associated with the individual.
	Name param.Field[string] `json:"name"`
}

func (r SandboxEmploymentUpdateParamsDepartment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The employment object.
type SandboxEmploymentUpdateParamsEmployment struct {
	// The secondary employment type of the individual. Options: `full_time`,
	// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
	Subtype param.Field[SandboxEmploymentUpdateParamsEmploymentSubtype] `json:"subtype"`
	// The main employment type of the individual.
	Type param.Field[SandboxEmploymentUpdateParamsEmploymentType] `json:"type"`
}

func (r SandboxEmploymentUpdateParamsEmployment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The secondary employment type of the individual. Options: `full_time`,
// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
type SandboxEmploymentUpdateParamsEmploymentSubtype string

const (
	SandboxEmploymentUpdateParamsEmploymentSubtypeFullTime             SandboxEmploymentUpdateParamsEmploymentSubtype = "full_time"
	SandboxEmploymentUpdateParamsEmploymentSubtypeIntern               SandboxEmploymentUpdateParamsEmploymentSubtype = "intern"
	SandboxEmploymentUpdateParamsEmploymentSubtypePartTime             SandboxEmploymentUpdateParamsEmploymentSubtype = "part_time"
	SandboxEmploymentUpdateParamsEmploymentSubtypeTemp                 SandboxEmploymentUpdateParamsEmploymentSubtype = "temp"
	SandboxEmploymentUpdateParamsEmploymentSubtypeSeasonal             SandboxEmploymentUpdateParamsEmploymentSubtype = "seasonal"
	SandboxEmploymentUpdateParamsEmploymentSubtypeIndividualContractor SandboxEmploymentUpdateParamsEmploymentSubtype = "individual_contractor"
)

func (r SandboxEmploymentUpdateParamsEmploymentSubtype) IsKnown() bool {
	switch r {
	case SandboxEmploymentUpdateParamsEmploymentSubtypeFullTime, SandboxEmploymentUpdateParamsEmploymentSubtypeIntern, SandboxEmploymentUpdateParamsEmploymentSubtypePartTime, SandboxEmploymentUpdateParamsEmploymentSubtypeTemp, SandboxEmploymentUpdateParamsEmploymentSubtypeSeasonal, SandboxEmploymentUpdateParamsEmploymentSubtypeIndividualContractor:
		return true
	}
	return false
}

// The main employment type of the individual.
type SandboxEmploymentUpdateParamsEmploymentType string

const (
	SandboxEmploymentUpdateParamsEmploymentTypeEmployee   SandboxEmploymentUpdateParamsEmploymentType = "employee"
	SandboxEmploymentUpdateParamsEmploymentTypeContractor SandboxEmploymentUpdateParamsEmploymentType = "contractor"
)

func (r SandboxEmploymentUpdateParamsEmploymentType) IsKnown() bool {
	switch r {
	case SandboxEmploymentUpdateParamsEmploymentTypeEmployee, SandboxEmploymentUpdateParamsEmploymentTypeContractor:
		return true
	}
	return false
}

// The detailed employment status of the individual. Available options: `active`,
// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
type SandboxEmploymentUpdateParamsEmploymentStatus string

const (
	SandboxEmploymentUpdateParamsEmploymentStatusActive     SandboxEmploymentUpdateParamsEmploymentStatus = "active"
	SandboxEmploymentUpdateParamsEmploymentStatusDeceased   SandboxEmploymentUpdateParamsEmploymentStatus = "deceased"
	SandboxEmploymentUpdateParamsEmploymentStatusLeave      SandboxEmploymentUpdateParamsEmploymentStatus = "leave"
	SandboxEmploymentUpdateParamsEmploymentStatusOnboarding SandboxEmploymentUpdateParamsEmploymentStatus = "onboarding"
	SandboxEmploymentUpdateParamsEmploymentStatusPrehire    SandboxEmploymentUpdateParamsEmploymentStatus = "prehire"
	SandboxEmploymentUpdateParamsEmploymentStatusRetired    SandboxEmploymentUpdateParamsEmploymentStatus = "retired"
	SandboxEmploymentUpdateParamsEmploymentStatusTerminated SandboxEmploymentUpdateParamsEmploymentStatus = "terminated"
)

func (r SandboxEmploymentUpdateParamsEmploymentStatus) IsKnown() bool {
	switch r {
	case SandboxEmploymentUpdateParamsEmploymentStatusActive, SandboxEmploymentUpdateParamsEmploymentStatusDeceased, SandboxEmploymentUpdateParamsEmploymentStatusLeave, SandboxEmploymentUpdateParamsEmploymentStatusOnboarding, SandboxEmploymentUpdateParamsEmploymentStatusPrehire, SandboxEmploymentUpdateParamsEmploymentStatusRetired, SandboxEmploymentUpdateParamsEmploymentStatusTerminated:
		return true
	}
	return false
}

// The manager object representing the manager of the individual within the org.
type SandboxEmploymentUpdateParamsManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r SandboxEmploymentUpdateParamsManager) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
