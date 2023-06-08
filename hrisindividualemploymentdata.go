package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISIndividualEmploymentDataService contains methods and other services that
// help with interacting with the Finch API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewHRISIndividualEmploymentDataService] method instead.
type HRISIndividualEmploymentDataService struct {
	Options []option.RequestOption
}

// NewHRISIndividualEmploymentDataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHRISIndividualEmploymentDataService(opts ...option.RequestOption) (r *HRISIndividualEmploymentDataService) {
	r = &HRISIndividualEmploymentDataService{}
	r.Options = opts
	return
}

// Read individual employment and income data
//
// Note: Income information is returned as reported by the provider. This may not
// always be annualized income, but may be in units of bi-weekly, semi-monthly,
// daily, etc, depending on what information the provider returns.
func (r *HRISIndividualEmploymentDataService) GetMany(ctx context.Context, body HRISIndividualEmploymentDataGetManyParams, opts ...option.RequestOption) (res *shared.ResponsesPage[EmploymentDataResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
//
// Note: Income information is returned as reported by the provider. This may not
// always be annualized income, but may be in units of bi-weekly, semi-monthly,
// daily, etc, depending on what information the provider returns.
func (r *HRISIndividualEmploymentDataService) GetManyAutoPaging(ctx context.Context, body HRISIndividualEmploymentDataGetManyParams, opts ...option.RequestOption) *shared.ResponsesPageAutoPager[EmploymentDataResponse] {
	return shared.NewResponsesPageAutoPager(r.GetMany(ctx, body, opts...))
}

type EmploymentData struct {
	// string A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,nullable"`
	// The current title of the individual.
	Title string `json:"title,nullable"`
	// The manager object representing the manager of the individual within the org.
	Manager EmploymentDataManager `json:"manager,nullable"`
	// The department object.
	Department EmploymentDataDepartment `json:"department,nullable"`
	// The employment object.
	Employment EmploymentDataEmployment `json:"employment,nullable"`
	StartDate  string                   `json:"start_date,nullable"`
	EndDate    string                   `json:"end_date,nullable"`
	// `true` if the individual an an active employee or contractor at the company.
	IsActive bool `json:"is_active,nullable"`
	// Worker's compensation classification code for this employee
	ClassCode string   `json:"class_code,nullable"`
	Location  Location `json:"location,nullable"`
	// The employee's income as reported by the provider. This may not always be
	// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
	// depending on what information the provider returns.
	Income Income `json:"income,nullable"`
	// The array of income history.
	IncomeHistory []Income `json:"income_history,nullable"`
	// Note: This property is only available if enabled for your account. Please reach
	// out to your Finch representative if you would like access.
	WorkID string `json:"work_id,nullable"`
	// Note: This property is only available if enabled for your account. Please reach
	// out to your Finch representative if you would like access.
	WorkID2 string `json:"work_id_2,nullable"`
	// Note: This property is only available if enabled for your account. Please reach
	// out to your Finch representative if you would like access.
	PayGroupIDs []string `json:"pay_group_ids,nullable"`
	JSON        employmentDataJSON
}

// employmentDataJSON contains the JSON metadata for the struct [EmploymentData]
type employmentDataJSON struct {
	ID            apijson.Field
	FirstName     apijson.Field
	MiddleName    apijson.Field
	LastName      apijson.Field
	Title         apijson.Field
	Manager       apijson.Field
	Department    apijson.Field
	Employment    apijson.Field
	StartDate     apijson.Field
	EndDate       apijson.Field
	IsActive      apijson.Field
	ClassCode     apijson.Field
	Location      apijson.Field
	Income        apijson.Field
	IncomeHistory apijson.Field
	WorkID        apijson.Field
	WorkID2       apijson.Field
	PayGroupIDs   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EmploymentData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The manager object representing the manager of the individual within the org.
type EmploymentDataManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID   string `json:"id"`
	JSON employmentDataManagerJSON
}

// employmentDataManagerJSON contains the JSON metadata for the struct
// [EmploymentDataManager]
type employmentDataManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The department object.
type EmploymentDataDepartment struct {
	// The name of the department associated with the individual.
	Name string `json:"name,nullable"`
	JSON employmentDataDepartmentJSON
}

// employmentDataDepartmentJSON contains the JSON metadata for the struct
// [EmploymentDataDepartment]
type employmentDataDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The employment object.
type EmploymentDataEmployment struct {
	// The main employment type of the individual.
	Type EmploymentDataEmploymentType `json:"type,nullable"`
	// The secondary employment type of the individual. Options: `full_time`, `part_time`, `intern`, `temp`, `seasonal` and `individual_contractor`.
	Subtype EmploymentDataEmploymentSubtype `json:"subtype,nullable"`
	JSON    employmentDataEmploymentJSON
}

// employmentDataEmploymentJSON contains the JSON metadata for the struct
// [EmploymentDataEmployment]
type employmentDataEmploymentJSON struct {
	Type        apijson.Field
	Subtype     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentDataEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmploymentDataEmploymentType string

const (
	EmploymentDataEmploymentTypeEmployee   EmploymentDataEmploymentType = "employee"
	EmploymentDataEmploymentTypeContractor EmploymentDataEmploymentType = "contractor"
)

type EmploymentDataEmploymentSubtype string

const (
	EmploymentDataEmploymentSubtypeFullTime             EmploymentDataEmploymentSubtype = "full_time"
	EmploymentDataEmploymentSubtypeIntern               EmploymentDataEmploymentSubtype = "intern"
	EmploymentDataEmploymentSubtypePartTime             EmploymentDataEmploymentSubtype = "part_time"
	EmploymentDataEmploymentSubtypeTemp                 EmploymentDataEmploymentSubtype = "temp"
	EmploymentDataEmploymentSubtypeSeasonal             EmploymentDataEmploymentSubtype = "seasonal"
	EmploymentDataEmploymentSubtypeIndividualContractor EmploymentDataEmploymentSubtype = "individual_contractor"
)

type EmploymentDataResponse struct {
	IndividualID string         `json:"individual_id"`
	Code         int64          `json:"code"`
	Body         EmploymentData `json:"body"`
	JSON         employmentDataResponseJSON
}

// employmentDataResponseJSON contains the JSON metadata for the struct
// [EmploymentDataResponse]
type employmentDataResponseJSON struct {
	IndividualID apijson.Field
	Code         apijson.Field
	Body         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmploymentDataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HRISIndividualEmploymentDataGetManyParams struct {
	// The array of batch requests.
	Requests param.Field[[]HRISIndividualEmploymentDataGetManyParamsRequests] `json:"requests,required"`
}

func (r HRISIndividualEmploymentDataGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISIndividualEmploymentDataGetManyParamsRequests struct {
	// A stable Finch `id` (UUID v4) for an individual in the company. There is no
	// limit to the number of `individual_id` to send per request. It is preferantial
	// to send all ids in a single request for Finch to optimize provider rate-limits.
	IndividualID param.Field[string] `json:"individual_id,required"`
}

func (r HRISIndividualEmploymentDataGetManyParamsRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISIndividualEmploymentDataGetManyResponse struct {
	Responses []EmploymentDataResponse `json:"responses,required"`
	JSON      hrisIndividualEmploymentDataGetManyResponseJSON
}

// hrisIndividualEmploymentDataGetManyResponseJSON contains the JSON metadata for
// the struct [HRISIndividualEmploymentDataGetManyResponse]
type hrisIndividualEmploymentDataGetManyResponseJSON struct {
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISIndividualEmploymentDataGetManyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
