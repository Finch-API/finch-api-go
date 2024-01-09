// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSandboxService] method instead.
type SandboxService struct {
	Options     []option.RequestOption
	Connections *SandboxConnectionService
	Company     *SandboxCompanyService
	Directory   *SandboxDirectoryService
	Individual  *SandboxIndividualService
	Employment  *SandboxEmploymentService
	Payment     *SandboxPaymentService
	Jobs        *SandboxJobService
}

// NewSandboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSandboxService(opts ...option.RequestOption) (r *SandboxService) {
	r = &SandboxService{}
	r.Options = opts
	r.Connections = NewSandboxConnectionService(opts...)
	r.Company = NewSandboxCompanyService(opts...)
	r.Directory = NewSandboxDirectoryService(opts...)
	r.Individual = NewSandboxIndividualService(opts...)
	r.Employment = NewSandboxEmploymentService(opts...)
	r.Payment = NewSandboxPaymentService(opts...)
	r.Jobs = NewSandboxJobService(opts...)
	return
}
