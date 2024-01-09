// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxJobService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSandboxJobService] method instead.
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
