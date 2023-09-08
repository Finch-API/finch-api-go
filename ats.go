// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/option"
)

// ATSService contains methods and other services that help with interacting with
// the Finch API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewATSService] method instead.
type ATSService struct {
	Options      []option.RequestOption
	Candidates   *ATSCandidateService
	Applications *ATSApplicationService
	Stages       *ATSStageService
	Jobs         *ATSJobService
	Offers       *ATSOfferService
}

// NewATSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewATSService(opts ...option.RequestOption) (r *ATSService) {
	r = &ATSService{}
	r.Options = opts
	r.Candidates = NewATSCandidateService(opts...)
	r.Applications = NewATSApplicationService(opts...)
	r.Stages = NewATSStageService(opts...)
	r.Jobs = NewATSJobService(opts...)
	r.Offers = NewATSOfferService(opts...)
	return
}
