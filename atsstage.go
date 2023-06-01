package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// ATSStageService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewATSStageService] method instead.
type ATSStageService struct {
	Options []option.RequestOption
}

// NewATSStageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewATSStageService(opts ...option.RequestOption) (r *ATSStageService) {
	r = &ATSStageService{}
	r.Options = opts
	return
}

// Get all job stages for an organization. Depending on the system, some jobs may
// have stages specific to that job. Other job stages may apply broadly to all jobs
// in the system. Use the `job_id` to determine whether a job applies specifically
// to a job.
func (r *ATSStageService) List(ctx context.Context, opts ...option.RequestOption) (res *shared.SinglePage[Stage], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ats/stages"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Get all job stages for an organization. Depending on the system, some jobs may
// have stages specific to that job. Other job stages may apply broadly to all jobs
// in the system. Use the `job_id` to determine whether a job applies specifically
// to a job.
func (r *ATSStageService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *shared.SinglePageAutoPager[Stage] {
	return shared.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type Stage struct {
	ID string `json:"id" format:"uuid"`
	// The job id that this stage applies to, if applicable. Not all systems enumerate
	// stages specific to jobs.
	JobID string `json:"job_id,nullable" format:"uuid"`
	Name  string `json:"name,nullable"`
	JSON  stageJSON
}

// stageJSON contains the JSON metadata for the struct [Stage]
type stageJSON struct {
	ID          apijson.Field
	JobID       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Stage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
