// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/option"
)

// JobService contains methods and other services that help with interacting with
// the Finch API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewJobService] method instead.
type JobService struct {
	Options   []option.RequestOption
	Automated *JobAutomatedService
	Manual    *JobManualService
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r *JobService) {
	r = &JobService{}
	r.Options = opts
	r.Automated = NewJobAutomatedService(opts...)
	r.Manual = NewJobManualService(opts...)
	return
}
