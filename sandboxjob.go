// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxJobService contains methods and other services that help with interacting
// with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxJobService] method instead.
type SandboxJobService struct {
	Options       []option.RequestOption
	Configuration *SandboxJobConfigurationService
}

// NewSandboxJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxJobService(opts ...option.RequestOption) (r *SandboxJobService) {
	r = &SandboxJobService{}
	r.Options = opts
	r.Configuration = NewSandboxJobConfigurationService(opts...)
	return
}

// Enqueue a new sandbox job
func (r *SandboxJobService) New(ctx context.Context, body SandboxJobNewParams, opts ...option.RequestOption) (res *SandboxJobNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandbox/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SandboxJobNewResponse struct {
	// The number of allowed refreshes per hour (per hour, fixed window)
	AllowedRefreshes int64 `json:"allowed_refreshes,required"`
	// The id of the job that has been created.
	JobID string `json:"job_id,required" format:"uuid"`
	// The url that can be used to retrieve the job status
	JobURL string `json:"job_url,required"`
	// The number of remaining refreshes available (per hour, fixed window)
	RemainingRefreshes int64                     `json:"remaining_refreshes,required"`
	JSON               sandboxJobNewResponseJSON `json:"-"`
}

// sandboxJobNewResponseJSON contains the JSON metadata for the struct
// [SandboxJobNewResponse]
type sandboxJobNewResponseJSON struct {
	AllowedRefreshes   apijson.Field
	JobID              apijson.Field
	JobURL             apijson.Field
	RemainingRefreshes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SandboxJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxJobNewResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxJobNewParams struct {
	// The type of job to start. Currently the only supported type is `data_sync_all`
	Type param.Field[SandboxJobNewParamsType] `json:"type,required"`
}

func (r SandboxJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of job to start. Currently the only supported type is `data_sync_all`
type SandboxJobNewParamsType string

const (
	SandboxJobNewParamsTypeDataSyncAll SandboxJobNewParamsType = "data_sync_all"
)

func (r SandboxJobNewParamsType) IsKnown() bool {
	switch r {
	case SandboxJobNewParamsTypeDataSyncAll:
		return true
	}
	return false
}
