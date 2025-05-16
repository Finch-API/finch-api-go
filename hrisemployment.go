// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
	"github.com/Finch-API/finch-api-go/shared"
	"github.com/tidwall/gjson"
)

// HRISEmploymentService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISEmploymentService] method instead.
type HRISEmploymentService struct {
	Options []option.RequestOption
}

// NewHRISEmploymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISEmploymentService(opts ...option.RequestOption) (r *HRISEmploymentService) {
	r = &HRISEmploymentService{}
	r.Options = opts
	return
}

// Read individual employment and income data
func (r *HRISEmploymentService) GetMany(ctx context.Context, body HRISEmploymentGetManyParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[EmploymentDataResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/employment"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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

// Read individual employment and income data
func (r *HRISEmploymentService) GetManyAutoPaging(ctx context.Context, body HRISEmploymentGetManyParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[EmploymentDataResponse] {
	return pagination.NewResponsesPageAutoPager(r.GetMany(ctx, body, opts...))
}

type EmploymentData struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id" format:"uuid"`
	// Worker's compensation classification code for this employee
	ClassCode string  `json:"class_code,nullable"`
	Code      float64 `json:"code"`
	// This field can have the runtime type of [[]EmploymentDataObjectCustomField].
	CustomFields interface{} `json:"custom_fields"`
	// This field can have the runtime type of [EmploymentDataObjectDepartment].
	Department interface{} `json:"department"`
	// This field can have the runtime type of [EmploymentDataObjectEmployment].
	Employment interface{} `json:"employment"`
	// The detailed employment status of the individual. Available options: `active`,
	// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
	EmploymentStatus EmploymentDataEmploymentStatus `json:"employment_status,nullable"`
	EndDate          string                         `json:"end_date,nullable"`
	FinchCode        string                         `json:"finch_code"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income Income `json:"income,nullable"`
	// This field can have the runtime type of [[]Income].
	IncomeHistory interface{} `json:"income_history"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive bool `json:"is_active,nullable"`
	// The legal last name of the individual.
	LastName         string   `json:"last_name,nullable"`
	LatestRehireDate string   `json:"latest_rehire_date,nullable"`
	Location         Location `json:"location,nullable"`
	// This field can have the runtime type of [EmploymentDataObjectManager].
	Manager interface{} `json:"manager"`
	Message string      `json:"message"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name,nullable"`
	Name       string `json:"name"`
	// The source system's unique employment identifier for this individual
	SourceID  string `json:"source_id,nullable"`
	StartDate string `json:"start_date,nullable"`
	// The current title of the individual.
	Title string `json:"title,nullable"`
	// This field is deprecated in favour of `source_id`
	//
	// Deprecated: deprecated
	WorkID string             `json:"work_id,nullable"`
	JSON   employmentDataJSON `json:"-"`
	union  EmploymentDataUnion
}

// employmentDataJSON contains the JSON metadata for the struct [EmploymentData]
type employmentDataJSON struct {
	ID               apijson.Field
	ClassCode        apijson.Field
	Code             apijson.Field
	CustomFields     apijson.Field
	Department       apijson.Field
	Employment       apijson.Field
	EmploymentStatus apijson.Field
	EndDate          apijson.Field
	FinchCode        apijson.Field
	FirstName        apijson.Field
	Income           apijson.Field
	IncomeHistory    apijson.Field
	IsActive         apijson.Field
	LastName         apijson.Field
	LatestRehireDate apijson.Field
	Location         apijson.Field
	Manager          apijson.Field
	Message          apijson.Field
	MiddleName       apijson.Field
	Name             apijson.Field
	SourceID         apijson.Field
	StartDate        apijson.Field
	Title            apijson.Field
	WorkID           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r employmentDataJSON) RawJSON() string {
	return r.raw
}

func (r *EmploymentData) UnmarshalJSON(data []byte) (err error) {
	*r = EmploymentData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EmploymentDataUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [EmploymentDataObject],
// [EmploymentDataBatchError].
func (r EmploymentData) AsUnion() EmploymentDataUnion {
	return r.union
}

// Union satisfied by [EmploymentDataObject] or [EmploymentDataBatchError].
type EmploymentDataUnion interface {
	implementsEmploymentData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmploymentDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmploymentDataObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmploymentDataBatchError{}),
		},
	)
}

type EmploymentDataObject struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id,required" format:"uuid"`
	// Worker's compensation classification code for this employee
	ClassCode string `json:"class_code,required,nullable"`
	// Custom fields for the individual. These are fields which are defined by the
	// employer in the system.
	CustomFields []EmploymentDataObjectCustomField `json:"custom_fields,required,nullable"`
	// The department object.
	Department EmploymentDataObjectDepartment `json:"department,required,nullable"`
	// The employment object.
	Employment EmploymentDataObjectEmployment `json:"employment,required,nullable"`
	// The detailed employment status of the individual. Available options: `active`,
	// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
	EmploymentStatus EmploymentDataObjectEmploymentStatus `json:"employment_status,required,nullable"`
	EndDate          string                               `json:"end_date,required,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,required,nullable"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive bool `json:"is_active,required,nullable"`
	// The legal last name of the individual.
	LastName         string   `json:"last_name,required,nullable"`
	LatestRehireDate string   `json:"latest_rehire_date,required,nullable"`
	Location         Location `json:"location,required,nullable"`
	// The manager object representing the manager of the individual within the org.
	Manager EmploymentDataObjectManager `json:"manager,required,nullable"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name,required,nullable"`
	StartDate  string `json:"start_date,required,nullable"`
	// The current title of the individual.
	Title string `json:"title,required,nullable"`
	// This field is deprecated in favour of `source_id`
	//
	// Deprecated: deprecated
	WorkID string `json:"work_id,required,nullable"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income Income `json:"income,nullable"`
	// The array of income history.
	IncomeHistory []Income `json:"income_history,nullable"`
	// The source system's unique employment identifier for this individual
	SourceID string                   `json:"source_id,nullable"`
	JSON     employmentDataObjectJSON `json:"-"`
}

// employmentDataObjectJSON contains the JSON metadata for the struct
// [EmploymentDataObject]
type employmentDataObjectJSON struct {
	ID               apijson.Field
	ClassCode        apijson.Field
	CustomFields     apijson.Field
	Department       apijson.Field
	Employment       apijson.Field
	EmploymentStatus apijson.Field
	EndDate          apijson.Field
	FirstName        apijson.Field
	IsActive         apijson.Field
	LastName         apijson.Field
	LatestRehireDate apijson.Field
	Location         apijson.Field
	Manager          apijson.Field
	MiddleName       apijson.Field
	StartDate        apijson.Field
	Title            apijson.Field
	WorkID           apijson.Field
	Income           apijson.Field
	IncomeHistory    apijson.Field
	SourceID         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmploymentDataObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataObjectJSON) RawJSON() string {
	return r.raw
}

func (r EmploymentDataObject) implementsEmploymentData() {}

type EmploymentDataObjectCustomField struct {
	Name  string                                     `json:"name"`
	Value EmploymentDataObjectCustomFieldsValueUnion `json:"value,nullable"`
	JSON  employmentDataObjectCustomFieldJSON        `json:"-"`
}

// employmentDataObjectCustomFieldJSON contains the JSON metadata for the struct
// [EmploymentDataObjectCustomField]
type employmentDataObjectCustomFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataObjectCustomField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataObjectCustomFieldJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString],
// [EmploymentDataObjectCustomFieldsValueArray], [shared.UnionFloat] or
// [shared.UnionBool].
type EmploymentDataObjectCustomFieldsValueUnion interface {
	ImplementsEmploymentDataObjectCustomFieldsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmploymentDataObjectCustomFieldsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmploymentDataObjectCustomFieldsValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type EmploymentDataObjectCustomFieldsValueArray []interface{}

func (r EmploymentDataObjectCustomFieldsValueArray) ImplementsEmploymentDataObjectCustomFieldsValueUnion() {
}

// The department object.
type EmploymentDataObjectDepartment struct {
	// The name of the department associated with the individual.
	Name string                             `json:"name,required,nullable"`
	JSON employmentDataObjectDepartmentJSON `json:"-"`
}

// employmentDataObjectDepartmentJSON contains the JSON metadata for the struct
// [EmploymentDataObjectDepartment]
type employmentDataObjectDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataObjectDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataObjectDepartmentJSON) RawJSON() string {
	return r.raw
}

// The employment object.
type EmploymentDataObjectEmployment struct {
	// The secondary employment type of the individual. Options: `full_time`,
	// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
	Subtype EmploymentDataObjectEmploymentSubtype `json:"subtype,required,nullable"`
	// The main employment type of the individual.
	Type EmploymentDataObjectEmploymentType `json:"type,required,nullable"`
	JSON employmentDataObjectEmploymentJSON `json:"-"`
}

// employmentDataObjectEmploymentJSON contains the JSON metadata for the struct
// [EmploymentDataObjectEmployment]
type employmentDataObjectEmploymentJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataObjectEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataObjectEmploymentJSON) RawJSON() string {
	return r.raw
}

// The secondary employment type of the individual. Options: `full_time`,
// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
type EmploymentDataObjectEmploymentSubtype string

const (
	EmploymentDataObjectEmploymentSubtypeFullTime             EmploymentDataObjectEmploymentSubtype = "full_time"
	EmploymentDataObjectEmploymentSubtypeIntern               EmploymentDataObjectEmploymentSubtype = "intern"
	EmploymentDataObjectEmploymentSubtypePartTime             EmploymentDataObjectEmploymentSubtype = "part_time"
	EmploymentDataObjectEmploymentSubtypeTemp                 EmploymentDataObjectEmploymentSubtype = "temp"
	EmploymentDataObjectEmploymentSubtypeSeasonal             EmploymentDataObjectEmploymentSubtype = "seasonal"
	EmploymentDataObjectEmploymentSubtypeIndividualContractor EmploymentDataObjectEmploymentSubtype = "individual_contractor"
)

func (r EmploymentDataObjectEmploymentSubtype) IsKnown() bool {
	switch r {
	case EmploymentDataObjectEmploymentSubtypeFullTime, EmploymentDataObjectEmploymentSubtypeIntern, EmploymentDataObjectEmploymentSubtypePartTime, EmploymentDataObjectEmploymentSubtypeTemp, EmploymentDataObjectEmploymentSubtypeSeasonal, EmploymentDataObjectEmploymentSubtypeIndividualContractor:
		return true
	}
	return false
}

// The main employment type of the individual.
type EmploymentDataObjectEmploymentType string

const (
	EmploymentDataObjectEmploymentTypeEmployee   EmploymentDataObjectEmploymentType = "employee"
	EmploymentDataObjectEmploymentTypeContractor EmploymentDataObjectEmploymentType = "contractor"
)

func (r EmploymentDataObjectEmploymentType) IsKnown() bool {
	switch r {
	case EmploymentDataObjectEmploymentTypeEmployee, EmploymentDataObjectEmploymentTypeContractor:
		return true
	}
	return false
}

// The detailed employment status of the individual. Available options: `active`,
// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
type EmploymentDataObjectEmploymentStatus string

const (
	EmploymentDataObjectEmploymentStatusActive     EmploymentDataObjectEmploymentStatus = "active"
	EmploymentDataObjectEmploymentStatusDeceased   EmploymentDataObjectEmploymentStatus = "deceased"
	EmploymentDataObjectEmploymentStatusLeave      EmploymentDataObjectEmploymentStatus = "leave"
	EmploymentDataObjectEmploymentStatusOnboarding EmploymentDataObjectEmploymentStatus = "onboarding"
	EmploymentDataObjectEmploymentStatusPrehire    EmploymentDataObjectEmploymentStatus = "prehire"
	EmploymentDataObjectEmploymentStatusRetired    EmploymentDataObjectEmploymentStatus = "retired"
	EmploymentDataObjectEmploymentStatusTerminated EmploymentDataObjectEmploymentStatus = "terminated"
)

func (r EmploymentDataObjectEmploymentStatus) IsKnown() bool {
	switch r {
	case EmploymentDataObjectEmploymentStatusActive, EmploymentDataObjectEmploymentStatusDeceased, EmploymentDataObjectEmploymentStatusLeave, EmploymentDataObjectEmploymentStatusOnboarding, EmploymentDataObjectEmploymentStatusPrehire, EmploymentDataObjectEmploymentStatusRetired, EmploymentDataObjectEmploymentStatusTerminated:
		return true
	}
	return false
}

// The manager object representing the manager of the individual within the org.
type EmploymentDataObjectManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID   string                          `json:"id,required" format:"uuid"`
	JSON employmentDataObjectManagerJSON `json:"-"`
}

// employmentDataObjectManagerJSON contains the JSON metadata for the struct
// [EmploymentDataObjectManager]
type employmentDataObjectManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataObjectManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataObjectManagerJSON) RawJSON() string {
	return r.raw
}

type EmploymentDataBatchError struct {
	Code      float64                      `json:"code,required"`
	Message   string                       `json:"message,required"`
	Name      string                       `json:"name,required"`
	FinchCode string                       `json:"finch_code"`
	JSON      employmentDataBatchErrorJSON `json:"-"`
}

// employmentDataBatchErrorJSON contains the JSON metadata for the struct
// [EmploymentDataBatchError]
type employmentDataBatchErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	FinchCode   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataBatchError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataBatchErrorJSON) RawJSON() string {
	return r.raw
}

func (r EmploymentDataBatchError) implementsEmploymentData() {}

// The detailed employment status of the individual. Available options: `active`,
// `deceased`, `leave`, `onboarding`, `prehire`, `retired`, `terminated`.
type EmploymentDataEmploymentStatus string

const (
	EmploymentDataEmploymentStatusActive     EmploymentDataEmploymentStatus = "active"
	EmploymentDataEmploymentStatusDeceased   EmploymentDataEmploymentStatus = "deceased"
	EmploymentDataEmploymentStatusLeave      EmploymentDataEmploymentStatus = "leave"
	EmploymentDataEmploymentStatusOnboarding EmploymentDataEmploymentStatus = "onboarding"
	EmploymentDataEmploymentStatusPrehire    EmploymentDataEmploymentStatus = "prehire"
	EmploymentDataEmploymentStatusRetired    EmploymentDataEmploymentStatus = "retired"
	EmploymentDataEmploymentStatusTerminated EmploymentDataEmploymentStatus = "terminated"
)

func (r EmploymentDataEmploymentStatus) IsKnown() bool {
	switch r {
	case EmploymentDataEmploymentStatusActive, EmploymentDataEmploymentStatusDeceased, EmploymentDataEmploymentStatusLeave, EmploymentDataEmploymentStatusOnboarding, EmploymentDataEmploymentStatusPrehire, EmploymentDataEmploymentStatusRetired, EmploymentDataEmploymentStatusTerminated:
		return true
	}
	return false
}

type EmploymentDataResponse struct {
	Body EmploymentData `json:"body,required"`
	Code int64          `json:"code,required"`
	// A stable Finch `id` (UUID v4) for an individual in the company.
	IndividualID string                     `json:"individual_id,required" format:"uuid"`
	JSON         employmentDataResponseJSON `json:"-"`
}

// employmentDataResponseJSON contains the JSON metadata for the struct
// [EmploymentDataResponse]
type employmentDataResponseJSON struct {
	Body         apijson.Field
	Code         apijson.Field
	IndividualID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmploymentDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataResponseJSON) RawJSON() string {
	return r.raw
}

type HRISEmploymentGetManyParams struct {
	// The array of batch requests.
	Requests param.Field[[]HRISEmploymentGetManyParamsRequest] `json:"requests,required"`
}

func (r HRISEmploymentGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISEmploymentGetManyParamsRequest struct {
	// A stable Finch `id` (UUID v4) for an individual in the company. There is no
	// limit to the number of `individual_id` to send per request. It is preferantial
	// to send all ids in a single request for Finch to optimize provider rate-limits.
	IndividualID param.Field[string] `json:"individual_id,required"`
}

func (r HRISEmploymentGetManyParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
