// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// JobManualService contains methods and other services that help with interacting
// with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobManualService] method instead.
type JobManualService struct {
	Options []option.RequestOption
}

// NewJobManualService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobManualService(opts ...option.RequestOption) (r *JobManualService) {
	r = &JobManualService{}
	r.Options = opts
	return
}

// Get a manual job by `job_id`. Manual jobs are completed by a human and include
// Assisted Benefits jobs.
func (r *JobManualService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *ManualAsyncJob, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("jobs/manual/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ManualAsyncJob struct {
	// Specific information about the job, such as individual statuses for batch jobs.
	Body   []interface{}        `json:"body,required,nullable"`
	JobID  string               `json:"job_id,required" format:"uuid"`
	Status ManualAsyncJobStatus `json:"status,required"`
	JSON   manualAsyncJobJSON   `json:"-"`
}

// manualAsyncJobJSON contains the JSON metadata for the struct [ManualAsyncJob]
type manualAsyncJobJSON struct {
	Body        apijson.Field
	JobID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManualAsyncJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r manualAsyncJobJSON) RawJSON() string {
	return r.raw
}

type ManualAsyncJobStatus string

const (
	ManualAsyncJobStatusPending    ManualAsyncJobStatus = "pending"
	ManualAsyncJobStatusInProgress ManualAsyncJobStatus = "in_progress"
	ManualAsyncJobStatusError      ManualAsyncJobStatus = "error"
	ManualAsyncJobStatusComplete   ManualAsyncJobStatus = "complete"
)

func (r ManualAsyncJobStatus) IsKnown() bool {
	switch r {
	case ManualAsyncJobStatusPending, ManualAsyncJobStatusInProgress, ManualAsyncJobStatusError, ManualAsyncJobStatusComplete:
		return true
	}
	return false
}
