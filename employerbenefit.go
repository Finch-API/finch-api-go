// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// EmployerBenefitService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewEmployerBenefitService] method
// instead.
type EmployerBenefitService struct {
	Options []option.RequestOption
}

// NewEmployerBenefitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmployerBenefitService(opts ...option.RequestOption) (r *EmployerBenefitService) {
	r = &EmployerBenefitService{}
	r.Options = opts
	return
}

// **Availability: Assisted Benefits providers only**
//
// Register existing benefits from the customer on the provider, on Finch's end.
// Please use the `/provider` endpoint to view available types for each provider.
func (r *EmployerBenefitService) Register(ctx context.Context, body EmployerBenefitRegisterParams, opts ...option.RequestOption) (res *RegisterCompanyBenefitsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "employer/benefits/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RegisterCompanyBenefitsResponse struct {
	BenefitID string `json:"benefit_id,required"`
	JobID     string `json:"job_id,required"`
	JSON      registerCompanyBenefitsResponseJSON
}

// registerCompanyBenefitsResponseJSON contains the JSON metadata for the struct
// [RegisterCompanyBenefitsResponse]
type registerCompanyBenefitsResponseJSON struct {
	BenefitID   apijson.Field
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegisterCompanyBenefitsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmployerBenefitRegisterParams struct {
	Description param.Field[string]           `json:"description"`
	Frequency   param.Field[BenefitFrequency] `json:"frequency"`
	// Type of benefit.
	Type param.Field[BenefitType] `json:"type"`
}

func (r EmployerBenefitRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
