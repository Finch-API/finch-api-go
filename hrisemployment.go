// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/v2/internal/apijson"
	"github.com/Finch-API/finch-api-go/v2/internal/apiquery"
	"github.com/Finch-API/finch-api-go/v2/internal/param"
	"github.com/Finch-API/finch-api-go/v2/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/v2/option"
	"github.com/Finch-API/finch-api-go/v2/packages/pagination"
	"github.com/Finch-API/finch-api-go/v2/shared"
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
func (r *HRISEmploymentService) GetMany(ctx context.Context, params HRISEmploymentGetManyParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[EmploymentDataResponse], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/employment"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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
func (r *HRISEmploymentService) GetManyAutoPaging(ctx context.Context, params HRISEmploymentGetManyParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[EmploymentDataResponse] {
	return pagination.NewResponsesPageAutoPager(r.GetMany(ctx, params, opts...))
}

type EmploymentData struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id" format:"uuid"`
	// Worker's compensation classification code for this employee
	ClassCode string  `json:"class_code" api:"nullable"`
	Code      float64 `json:"code"`
	// This field can have the runtime type of
	// [[]EmploymentDataEmploymentDataResponseBodyCustomField].
	CustomFields interface{} `json:"custom_fields"`
	// This field can have the runtime type of
	// [EmploymentDataEmploymentDataResponseBodyDepartment].
	Department interface{} `json:"department"`
	// This field can have the runtime type of
	// [EmploymentDataEmploymentDataResponseBodyEmployment].
	Employment interface{} `json:"employment"`
	// The detailed employment status of the individual.
	EmploymentStatus EmploymentDataEmploymentStatus `json:"employment_status" api:"nullable"`
	EndDate          string                         `json:"end_date" api:"nullable"`
	FinchCode        string                         `json:"finch_code"`
	// The legal first name of the individual.
	FirstName string `json:"first_name" api:"nullable"`
	// The FLSA status of the individual. Available options: `exempt`, `non_exempt`,
	// `unknown`.
	FlsaStatus EmploymentDataFlsaStatus `json:"flsa_status" api:"nullable"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income Income `json:"income" api:"nullable"`
	// This field can have the runtime type of [[]Income].
	IncomeHistory interface{} `json:"income_history"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive bool `json:"is_active" api:"nullable"`
	// The legal last name of the individual.
	LastName         string   `json:"last_name" api:"nullable"`
	LatestRehireDate string   `json:"latest_rehire_date" api:"nullable"`
	Location         Location `json:"location" api:"nullable"`
	// This field can have the runtime type of
	// [EmploymentDataEmploymentDataResponseBodyManager].
	Manager interface{} `json:"manager"`
	Message string      `json:"message"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name" api:"nullable"`
	Name       string `json:"name"`
	// The source system's unique employment identifier for this individual
	SourceID  string `json:"source_id" api:"nullable"`
	StartDate string `json:"start_date" api:"nullable"`
	// The current title of the individual.
	Title string `json:"title" api:"nullable"`
	// This field is deprecated in favour of `source_id`
	//
	// Deprecated: deprecated
	WorkID string             `json:"work_id" api:"nullable"`
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
	FlsaStatus       apijson.Field
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
// Possible runtime types of the union are
// [EmploymentDataEmploymentDataResponseBody], [EmploymentDataBatchError].
func (r EmploymentData) AsUnion() EmploymentDataUnion {
	return r.union
}

// Union satisfied by [EmploymentDataEmploymentDataResponseBody] or
// [EmploymentDataBatchError].
type EmploymentDataUnion interface {
	implementsEmploymentData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmploymentDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmploymentDataEmploymentDataResponseBody{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmploymentDataBatchError{}),
		},
	)
}

type EmploymentDataEmploymentDataResponseBody struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id" api:"required" format:"uuid"`
	// Worker's compensation classification code for this employee
	ClassCode string `json:"class_code" api:"required,nullable"`
	// The department object.
	Department EmploymentDataEmploymentDataResponseBodyDepartment `json:"department" api:"required,nullable"`
	// The employment object.
	Employment EmploymentDataEmploymentDataResponseBodyEmployment `json:"employment" api:"required,nullable"`
	// The detailed employment status of the individual.
	EmploymentStatus EmploymentDataEmploymentDataResponseBodyEmploymentStatus `json:"employment_status" api:"required,nullable"`
	EndDate          string                                                   `json:"end_date" api:"required,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name" api:"required,nullable"`
	// The FLSA status of the individual. Available options: `exempt`, `non_exempt`,
	// `unknown`.
	FlsaStatus EmploymentDataEmploymentDataResponseBodyFlsaStatus `json:"flsa_status" api:"required,nullable"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive bool `json:"is_active" api:"required,nullable"`
	// The legal last name of the individual.
	LastName         string   `json:"last_name" api:"required,nullable"`
	LatestRehireDate string   `json:"latest_rehire_date" api:"required,nullable"`
	Location         Location `json:"location" api:"required,nullable"`
	// The manager object representing the manager of the individual within the org.
	Manager EmploymentDataEmploymentDataResponseBodyManager `json:"manager" api:"required,nullable"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name" api:"required,nullable"`
	StartDate  string `json:"start_date" api:"required,nullable"`
	// The current title of the individual.
	Title string `json:"title" api:"required,nullable"`
	// Custom fields for the individual. These are fields which are defined by the
	// employer in the system. Custom fields are not currently supported for assisted
	// connections.
	CustomFields []EmploymentDataEmploymentDataResponseBodyCustomField `json:"custom_fields" api:"nullable"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income Income `json:"income" api:"nullable"`
	// The array of income history.
	IncomeHistory []Income `json:"income_history" api:"nullable"`
	// The source system's unique employment identifier for this individual
	SourceID string `json:"source_id" api:"nullable"`
	// This field is deprecated in favour of `source_id`
	//
	// Deprecated: deprecated
	WorkID string                                       `json:"work_id" api:"nullable"`
	JSON   employmentDataEmploymentDataResponseBodyJSON `json:"-"`
}

// employmentDataEmploymentDataResponseBodyJSON contains the JSON metadata for the
// struct [EmploymentDataEmploymentDataResponseBody]
type employmentDataEmploymentDataResponseBodyJSON struct {
	ID               apijson.Field
	ClassCode        apijson.Field
	Department       apijson.Field
	Employment       apijson.Field
	EmploymentStatus apijson.Field
	EndDate          apijson.Field
	FirstName        apijson.Field
	FlsaStatus       apijson.Field
	IsActive         apijson.Field
	LastName         apijson.Field
	LatestRehireDate apijson.Field
	Location         apijson.Field
	Manager          apijson.Field
	MiddleName       apijson.Field
	StartDate        apijson.Field
	Title            apijson.Field
	CustomFields     apijson.Field
	Income           apijson.Field
	IncomeHistory    apijson.Field
	SourceID         apijson.Field
	WorkID           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmploymentDataEmploymentDataResponseBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataEmploymentDataResponseBodyJSON) RawJSON() string {
	return r.raw
}

func (r EmploymentDataEmploymentDataResponseBody) implementsEmploymentData() {}

// The department object.
type EmploymentDataEmploymentDataResponseBodyDepartment struct {
	// The name of the department associated with the individual.
	Name string                                                 `json:"name" api:"required,nullable"`
	JSON employmentDataEmploymentDataResponseBodyDepartmentJSON `json:"-"`
}

// employmentDataEmploymentDataResponseBodyDepartmentJSON contains the JSON
// metadata for the struct [EmploymentDataEmploymentDataResponseBodyDepartment]
type employmentDataEmploymentDataResponseBodyDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataEmploymentDataResponseBodyDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataEmploymentDataResponseBodyDepartmentJSON) RawJSON() string {
	return r.raw
}

// The employment object.
type EmploymentDataEmploymentDataResponseBodyEmployment struct {
	// The secondary employment type of the individual. Options: `full_time`,
	// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
	Subtype EmploymentDataEmploymentDataResponseBodyEmploymentSubtype `json:"subtype" api:"required,nullable"`
	// The main employment type of the individual.
	Type EmploymentDataEmploymentDataResponseBodyEmploymentType `json:"type" api:"required,nullable"`
	JSON employmentDataEmploymentDataResponseBodyEmploymentJSON `json:"-"`
}

// employmentDataEmploymentDataResponseBodyEmploymentJSON contains the JSON
// metadata for the struct [EmploymentDataEmploymentDataResponseBodyEmployment]
type employmentDataEmploymentDataResponseBodyEmploymentJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataEmploymentDataResponseBodyEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataEmploymentDataResponseBodyEmploymentJSON) RawJSON() string {
	return r.raw
}

// The secondary employment type of the individual. Options: `full_time`,
// `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
type EmploymentDataEmploymentDataResponseBodyEmploymentSubtype string

const (
	EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeFullTime             EmploymentDataEmploymentDataResponseBodyEmploymentSubtype = "full_time"
	EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeIntern               EmploymentDataEmploymentDataResponseBodyEmploymentSubtype = "intern"
	EmploymentDataEmploymentDataResponseBodyEmploymentSubtypePartTime             EmploymentDataEmploymentDataResponseBodyEmploymentSubtype = "part_time"
	EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeTemp                 EmploymentDataEmploymentDataResponseBodyEmploymentSubtype = "temp"
	EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeSeasonal             EmploymentDataEmploymentDataResponseBodyEmploymentSubtype = "seasonal"
	EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeIndividualContractor EmploymentDataEmploymentDataResponseBodyEmploymentSubtype = "individual_contractor"
)

func (r EmploymentDataEmploymentDataResponseBodyEmploymentSubtype) IsKnown() bool {
	switch r {
	case EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeFullTime, EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeIntern, EmploymentDataEmploymentDataResponseBodyEmploymentSubtypePartTime, EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeTemp, EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeSeasonal, EmploymentDataEmploymentDataResponseBodyEmploymentSubtypeIndividualContractor:
		return true
	}
	return false
}

// The main employment type of the individual.
type EmploymentDataEmploymentDataResponseBodyEmploymentType string

const (
	EmploymentDataEmploymentDataResponseBodyEmploymentTypeEmployee   EmploymentDataEmploymentDataResponseBodyEmploymentType = "employee"
	EmploymentDataEmploymentDataResponseBodyEmploymentTypeContractor EmploymentDataEmploymentDataResponseBodyEmploymentType = "contractor"
)

func (r EmploymentDataEmploymentDataResponseBodyEmploymentType) IsKnown() bool {
	switch r {
	case EmploymentDataEmploymentDataResponseBodyEmploymentTypeEmployee, EmploymentDataEmploymentDataResponseBodyEmploymentTypeContractor:
		return true
	}
	return false
}

// The detailed employment status of the individual.
type EmploymentDataEmploymentDataResponseBodyEmploymentStatus string

const (
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusActive     EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "active"
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusDeceased   EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "deceased"
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusLeave      EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "leave"
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusOnboarding EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "onboarding"
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusPrehire    EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "prehire"
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusRetired    EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "retired"
	EmploymentDataEmploymentDataResponseBodyEmploymentStatusTerminated EmploymentDataEmploymentDataResponseBodyEmploymentStatus = "terminated"
)

func (r EmploymentDataEmploymentDataResponseBodyEmploymentStatus) IsKnown() bool {
	switch r {
	case EmploymentDataEmploymentDataResponseBodyEmploymentStatusActive, EmploymentDataEmploymentDataResponseBodyEmploymentStatusDeceased, EmploymentDataEmploymentDataResponseBodyEmploymentStatusLeave, EmploymentDataEmploymentDataResponseBodyEmploymentStatusOnboarding, EmploymentDataEmploymentDataResponseBodyEmploymentStatusPrehire, EmploymentDataEmploymentDataResponseBodyEmploymentStatusRetired, EmploymentDataEmploymentDataResponseBodyEmploymentStatusTerminated:
		return true
	}
	return false
}

// The FLSA status of the individual. Available options: `exempt`, `non_exempt`,
// `unknown`.
type EmploymentDataEmploymentDataResponseBodyFlsaStatus string

const (
	EmploymentDataEmploymentDataResponseBodyFlsaStatusExempt    EmploymentDataEmploymentDataResponseBodyFlsaStatus = "exempt"
	EmploymentDataEmploymentDataResponseBodyFlsaStatusNonExempt EmploymentDataEmploymentDataResponseBodyFlsaStatus = "non_exempt"
	EmploymentDataEmploymentDataResponseBodyFlsaStatusUnknown   EmploymentDataEmploymentDataResponseBodyFlsaStatus = "unknown"
)

func (r EmploymentDataEmploymentDataResponseBodyFlsaStatus) IsKnown() bool {
	switch r {
	case EmploymentDataEmploymentDataResponseBodyFlsaStatusExempt, EmploymentDataEmploymentDataResponseBodyFlsaStatusNonExempt, EmploymentDataEmploymentDataResponseBodyFlsaStatusUnknown:
		return true
	}
	return false
}

// The manager object representing the manager of the individual within the org.
type EmploymentDataEmploymentDataResponseBodyManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID   string                                              `json:"id" api:"required" format:"uuid"`
	JSON employmentDataEmploymentDataResponseBodyManagerJSON `json:"-"`
}

// employmentDataEmploymentDataResponseBodyManagerJSON contains the JSON metadata
// for the struct [EmploymentDataEmploymentDataResponseBodyManager]
type employmentDataEmploymentDataResponseBodyManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataEmploymentDataResponseBodyManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataEmploymentDataResponseBodyManagerJSON) RawJSON() string {
	return r.raw
}

type EmploymentDataEmploymentDataResponseBodyCustomField struct {
	Name  string                                                         `json:"name" api:"nullable"`
	Value EmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion `json:"value" api:"nullable"`
	JSON  employmentDataEmploymentDataResponseBodyCustomFieldJSON        `json:"-"`
}

// employmentDataEmploymentDataResponseBodyCustomFieldJSON contains the JSON
// metadata for the struct [EmploymentDataEmploymentDataResponseBodyCustomField]
type employmentDataEmploymentDataResponseBodyCustomFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataEmploymentDataResponseBodyCustomField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r employmentDataEmploymentDataResponseBodyCustomFieldJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString],
// [EmploymentDataEmploymentDataResponseBodyCustomFieldsValueArray],
// [shared.UnionFloat] or [shared.UnionBool].
type EmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion interface {
	ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmploymentDataEmploymentDataResponseBodyCustomFieldsValueArray{}),
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

type EmploymentDataEmploymentDataResponseBodyCustomFieldsValueArray []interface{}

func (r EmploymentDataEmploymentDataResponseBodyCustomFieldsValueArray) ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion() {
}

type EmploymentDataBatchError struct {
	Code      float64                      `json:"code" api:"required"`
	Message   string                       `json:"message" api:"required"`
	Name      string                       `json:"name" api:"required"`
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

// The detailed employment status of the individual.
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

// The FLSA status of the individual. Available options: `exempt`, `non_exempt`,
// `unknown`.
type EmploymentDataFlsaStatus string

const (
	EmploymentDataFlsaStatusExempt    EmploymentDataFlsaStatus = "exempt"
	EmploymentDataFlsaStatusNonExempt EmploymentDataFlsaStatus = "non_exempt"
	EmploymentDataFlsaStatusUnknown   EmploymentDataFlsaStatus = "unknown"
)

func (r EmploymentDataFlsaStatus) IsKnown() bool {
	switch r {
	case EmploymentDataFlsaStatusExempt, EmploymentDataFlsaStatusNonExempt, EmploymentDataFlsaStatusUnknown:
		return true
	}
	return false
}

type EmploymentDataResponse struct {
	Body EmploymentData `json:"body" api:"required"`
	Code int64          `json:"code" api:"required"`
	// A stable Finch `id` (UUID v4) for an individual in the company.
	IndividualID string                     `json:"individual_id" api:"required" format:"uuid"`
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
	// The array of batch requests. Maximum 10000 items per request.
	Requests param.Field[[]HRISEmploymentGetManyParamsRequest] `json:"requests" api:"required"`
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
}

func (r HRISEmploymentGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISEmploymentGetManyParams]'s query parameters as
// `url.Values`.
func (r HRISEmploymentGetManyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISEmploymentGetManyParamsRequest struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	IndividualID param.Field[string] `json:"individual_id" api:"required"`
}

func (r HRISEmploymentGetManyParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
