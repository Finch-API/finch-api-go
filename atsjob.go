// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// ATSJobService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewATSJobService] method instead.
type ATSJobService struct {
	Options []option.RequestOption
}

// NewATSJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewATSJobService(opts ...option.RequestOption) (r *ATSJobService) {
	r = &ATSJobService{}
	r.Options = opts
	return
}

// Gets a job from an organization.
func (r *ATSJobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *Job, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ats/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all of an organization's jobs.
func (r *ATSJobService) List(ctx context.Context, query ATSJobListParams, opts ...option.RequestOption) (res *JobsPage, err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ats/jobs"
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

// Gets all of an organization's jobs.
func (r *ATSJobService) ListAutoPaging(ctx context.Context, query ATSJobListParams, opts ...option.RequestOption) *JobsPageAutoPager {
	return NewJobsPageAutoPager(r.List(ctx, query, opts...))
}

type JobsPage struct {
	Jobs   []Job  `json:"jobs,required"`
	Paging Paging `json:"paging,required"`
	JSON   jobsPageJSON
	cfg    *requestconfig.RequestConfig
	res    *http.Response
}

// jobsPageJSON contains the JSON metadata for the struct [JobsPage]
type jobsPageJSON struct {
	Jobs        apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *JobsPage) GetNextPage() (res *JobsPage, err error) {
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *JobsPage) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type JobsPageAutoPager struct {
	page *JobsPage
	cur  Job
	idx  int
	run  int
	err  error
}

func NewJobsPageAutoPager(page *JobsPage, err error) *JobsPageAutoPager {
	return &JobsPageAutoPager{
		page: page,
		err:  err,
	}
}

func (r *JobsPageAutoPager) Next() bool {
	if r.page == nil || len(r.page.Jobs) == 0 {
		return false
	}
	if r.idx >= len(r.page.Jobs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Jobs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *JobsPageAutoPager) Current() Job {
	return r.cur
}

func (r *JobsPageAutoPager) Err() error {
	return r.err
}

func (r *JobsPageAutoPager) Index() int {
	return r.run
}

type Job struct {
	ID         string        `json:"id,required" format:"uuid"`
	ClosedAt   time.Time     `json:"closed_at,required,nullable" format:"date-time"`
	CreatedAt  time.Time     `json:"created_at,required,nullable" format:"date-time"`
	Department JobDepartment `json:"department,required"`
	HiringTeam JobHiringTeam `json:"hiring_team,required"`
	Name       string        `json:"name,required,nullable"`
	Status     JobStatus     `json:"status,required,nullable"`
	JSON       jobJSON
}

// jobJSON contains the JSON metadata for the struct [Job]
type jobJSON struct {
	ID          apijson.Field
	ClosedAt    apijson.Field
	CreatedAt   apijson.Field
	Department  apijson.Field
	HiringTeam  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Job) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type JobDepartment struct {
	Name string `json:"name,nullable"`
	JSON jobDepartmentJSON
}

// jobDepartmentJSON contains the JSON metadata for the struct [JobDepartment]
type jobDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type JobHiringTeam struct {
	HiringManagers []JobHiringTeamHiringManager `json:"hiring_managers,nullable"`
	Recruiters     []JobHiringTeamRecruiter     `json:"recruiters,nullable"`
	JSON           jobHiringTeamJSON
}

// jobHiringTeamJSON contains the JSON metadata for the struct [JobHiringTeam]
type jobHiringTeamJSON struct {
	HiringManagers apijson.Field
	Recruiters     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *JobHiringTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type JobHiringTeamHiringManager struct {
	Name string `json:"name"`
	JSON jobHiringTeamHiringManagerJSON
}

// jobHiringTeamHiringManagerJSON contains the JSON metadata for the struct
// [JobHiringTeamHiringManager]
type jobHiringTeamHiringManagerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobHiringTeamHiringManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type JobHiringTeamRecruiter struct {
	Name string `json:"name"`
	JSON jobHiringTeamRecruiterJSON
}

// jobHiringTeamRecruiterJSON contains the JSON metadata for the struct
// [JobHiringTeamRecruiter]
type jobHiringTeamRecruiterJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobHiringTeamRecruiter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type JobStatus string

const (
	JobStatusOpen     JobStatus = "open"
	JobStatusClosed   JobStatus = "closed"
	JobStatusOnHold   JobStatus = "on_hold"
	JobStatusDraft    JobStatus = "draft"
	JobStatusArchived JobStatus = "archived"
)

type ATSJobListParams struct {
	// Number of jobs to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [ATSJobListParams]'s query parameters as `url.Values`.
func (r ATSJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
