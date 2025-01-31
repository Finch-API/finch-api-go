// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
)

// JobAutomatedService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobAutomatedService] method instead.
type JobAutomatedService struct {
	Options []option.RequestOption
}

// NewJobAutomatedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobAutomatedService(opts ...option.RequestOption) (r *JobAutomatedService) {
	r = &JobAutomatedService{}
	r.Options = opts
	return
}

// Enqueue an automated job.
//
// `data_sync_all`: Enqueue a job to re-sync all data for a connection.
// `data_sync_all` has a concurrency limit of 1 job at a time per connection. This
// means that if this endpoint is called while a job is already in progress for
// this connection, Finch will return the `job_id` of the job that is currently in
// progress. Finch allows a fixed window rate limit of 1 forced refresh per hour
// per connection.
//
// `w4_form_employee_sync`: Enqueues a job for sync W-4 data for a particular
// individual, identified by `individual_id`. This feature is currently in beta.
//
// This endpoint is available for _Scale_ tier customers as an add-on. To request
// access to this endpoint, please contact your Finch account manager.
func (r *JobAutomatedService) New(ctx context.Context, body JobAutomatedNewParams, opts ...option.RequestOption) (res *JobAutomatedNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "jobs/automated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an automated job by `job_id`.
func (r *JobAutomatedService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *AutomatedAsyncJob, err error) {
	opts = append(r.Options[:], opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("jobs/automated/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all automated jobs. Automated jobs are completed by a machine. By default,
// jobs are sorted in descending order by submission time. For scheduled jobs such
// as data syncs, only the next scheduled job is shown.
func (r *JobAutomatedService) List(ctx context.Context, query JobAutomatedListParams, opts ...option.RequestOption) (res *pagination.Page[AutomatedAsyncJob], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "jobs/automated"
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

// Get all automated jobs. Automated jobs are completed by a machine. By default,
// jobs are sorted in descending order by submission time. For scheduled jobs such
// as data syncs, only the next scheduled job is shown.
func (r *JobAutomatedService) ListAutoPaging(ctx context.Context, query JobAutomatedListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AutomatedAsyncJob] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type AutomatedAsyncJob struct {
	// The datetime the job completed.
	CompletedAt time.Time `json:"completed_at,required,nullable" format:"date-time"`
	// The datetime when the job was created. for scheduled jobs, this will be the
	// initial connection time. For ad-hoc jobs, this will be the time the creation
	// request was received.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The id of the job that has been created.
	JobID string `json:"job_id,required" format:"uuid"`
	// The url that can be used to retrieve the job status
	JobURL string `json:"job_url,required"`
	// The datetime a job is scheduled to be run. For scheduled jobs, this datetime can
	// be in the future if the job has not yet been enqueued. For ad-hoc jobs, this
	// field will be null.
	ScheduledAt time.Time `json:"scheduled_at,required,nullable" format:"date-time"`
	// The datetime a job entered into the job queue.
	StartedAt time.Time               `json:"started_at,required,nullable" format:"date-time"`
	Status    AutomatedAsyncJobStatus `json:"status,required"`
	// Only `data_sync_all` currently supported
	Type AutomatedAsyncJobType `json:"type,required"`
	JSON automatedAsyncJobJSON `json:"-"`
}

// automatedAsyncJobJSON contains the JSON metadata for the struct
// [AutomatedAsyncJob]
type automatedAsyncJobJSON struct {
	CompletedAt apijson.Field
	CreatedAt   apijson.Field
	JobID       apijson.Field
	JobURL      apijson.Field
	ScheduledAt apijson.Field
	StartedAt   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomatedAsyncJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automatedAsyncJobJSON) RawJSON() string {
	return r.raw
}

type AutomatedAsyncJobStatus string

const (
	AutomatedAsyncJobStatusPending          AutomatedAsyncJobStatus = "pending"
	AutomatedAsyncJobStatusInProgress       AutomatedAsyncJobStatus = "in_progress"
	AutomatedAsyncJobStatusComplete         AutomatedAsyncJobStatus = "complete"
	AutomatedAsyncJobStatusError            AutomatedAsyncJobStatus = "error"
	AutomatedAsyncJobStatusReauthError      AutomatedAsyncJobStatus = "reauth_error"
	AutomatedAsyncJobStatusPermissionsError AutomatedAsyncJobStatus = "permissions_error"
)

func (r AutomatedAsyncJobStatus) IsKnown() bool {
	switch r {
	case AutomatedAsyncJobStatusPending, AutomatedAsyncJobStatusInProgress, AutomatedAsyncJobStatusComplete, AutomatedAsyncJobStatusError, AutomatedAsyncJobStatusReauthError, AutomatedAsyncJobStatusPermissionsError:
		return true
	}
	return false
}

// Only `data_sync_all` currently supported
type AutomatedAsyncJobType string

const (
	AutomatedAsyncJobTypeDataSyncAll AutomatedAsyncJobType = "data_sync_all"
)

func (r AutomatedAsyncJobType) IsKnown() bool {
	switch r {
	case AutomatedAsyncJobTypeDataSyncAll:
		return true
	}
	return false
}

type JobAutomatedNewResponse struct {
	// The number of allowed refreshes per hour (per hour, fixed window)
	AllowedRefreshes int64 `json:"allowed_refreshes,required"`
	// The id of the job that has been created.
	JobID string `json:"job_id,required" format:"uuid"`
	// The url that can be used to retrieve the job status
	JobURL string `json:"job_url,required"`
	// The number of remaining refreshes available (per hour, fixed window)
	RemainingRefreshes int64                       `json:"remaining_refreshes,required"`
	JSON               jobAutomatedNewResponseJSON `json:"-"`
}

// jobAutomatedNewResponseJSON contains the JSON metadata for the struct
// [JobAutomatedNewResponse]
type jobAutomatedNewResponseJSON struct {
	AllowedRefreshes   apijson.Field
	JobID              apijson.Field
	JobURL             apijson.Field
	RemainingRefreshes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *JobAutomatedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobAutomatedNewResponseJSON) RawJSON() string {
	return r.raw
}

// This interface is a union satisfied by one of the following:
// [JobAutomatedNewParamsDataSyncAll], [JobAutomatedNewParamsW4FormEmployeeSync].
type JobAutomatedNewParams interface {
	ImplementsJobAutomatedNewParams()
}

type JobAutomatedNewParamsDataSyncAll struct {
	// The type of job to start.
	Type param.Field[JobAutomatedNewParamsDataSyncAllType] `json:"type,required"`
}

func (r JobAutomatedNewParamsDataSyncAll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (JobAutomatedNewParamsDataSyncAll) ImplementsJobAutomatedNewParams() {

}

// The type of job to start.
type JobAutomatedNewParamsDataSyncAllType string

const (
	JobAutomatedNewParamsDataSyncAllTypeDataSyncAll JobAutomatedNewParamsDataSyncAllType = "data_sync_all"
)

func (r JobAutomatedNewParamsDataSyncAllType) IsKnown() bool {
	switch r {
	case JobAutomatedNewParamsDataSyncAllTypeDataSyncAll:
		return true
	}
	return false
}

type JobAutomatedNewParamsW4FormEmployeeSync struct {
	Params param.Field[JobAutomatedNewParamsW4FormEmployeeSyncParams] `json:"params,required"`
	// The type of job to start.
	Type param.Field[JobAutomatedNewParamsW4FormEmployeeSyncType] `json:"type,required"`
}

func (r JobAutomatedNewParamsW4FormEmployeeSync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (JobAutomatedNewParamsW4FormEmployeeSync) ImplementsJobAutomatedNewParams() {

}

type JobAutomatedNewParamsW4FormEmployeeSyncParams struct {
	// The unique ID of the individual for W-4 data sync.
	IndividualID param.Field[string] `json:"individual_id,required"`
}

func (r JobAutomatedNewParamsW4FormEmployeeSyncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of job to start.
type JobAutomatedNewParamsW4FormEmployeeSyncType string

const (
	JobAutomatedNewParamsW4FormEmployeeSyncTypeW4FormEmployeeSync JobAutomatedNewParamsW4FormEmployeeSyncType = "w4_form_employee_sync"
)

func (r JobAutomatedNewParamsW4FormEmployeeSyncType) IsKnown() bool {
	switch r {
	case JobAutomatedNewParamsW4FormEmployeeSyncTypeW4FormEmployeeSync:
		return true
	}
	return false
}

type JobAutomatedListParams struct {
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [JobAutomatedListParams]'s query parameters as `url.Values`.
func (r JobAutomatedListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
