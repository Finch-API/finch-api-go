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

// SandboxJobConfigurationService contains methods and other services that help
// with interacting with the Finch API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSandboxJobConfigurationService]
// method instead.
type SandboxJobConfigurationService struct {
	Options []option.RequestOption
}

// NewSandboxJobConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxJobConfigurationService(opts ...option.RequestOption) (r *SandboxJobConfigurationService) {
	r = &SandboxJobConfigurationService{}
	r.Options = opts
	return
}

// Get configurations for sandbox jobs
func (r *SandboxJobConfigurationService) Get(ctx context.Context, opts ...option.RequestOption) (res *[]SandboxJobConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := "sandbox/jobs/configuration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update configurations for sandbox jobs
func (r *SandboxJobConfigurationService) Update(ctx context.Context, body SandboxJobConfigurationUpdateParams, opts ...option.RequestOption) (res *SandboxJobConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := "sandbox/jobs/configuration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SandboxJobConfiguration struct {
	CompletionStatus SandboxJobConfigurationCompletionStatus `json:"completion_status,required"`
	Type             SandboxJobConfigurationType             `json:"type,required"`
	JSON             sandboxJobConfigurationJSON             `json:"-"`
}

// sandboxJobConfigurationJSON contains the JSON metadata for the struct
// [SandboxJobConfiguration]
type sandboxJobConfigurationJSON struct {
	CompletionStatus apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SandboxJobConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxJobConfigurationCompletionStatus string

const (
	SandboxJobConfigurationCompletionStatusComplete         SandboxJobConfigurationCompletionStatus = "complete"
	SandboxJobConfigurationCompletionStatusReauthError      SandboxJobConfigurationCompletionStatus = "reauth_error"
	SandboxJobConfigurationCompletionStatusPermissionsError SandboxJobConfigurationCompletionStatus = "permissions_error"
	SandboxJobConfigurationCompletionStatusError            SandboxJobConfigurationCompletionStatus = "error"
)

type SandboxJobConfigurationType string

const (
	SandboxJobConfigurationTypeDataSyncAll SandboxJobConfigurationType = "data_sync_all"
)

type SandboxJobConfigurationUpdateParams struct {
	CompletionStatus param.Field[SandboxJobConfigurationUpdateParamsCompletionStatus] `json:"completion_status,required"`
	Type             param.Field[SandboxJobConfigurationUpdateParamsType]             `json:"type,required"`
}

func (r SandboxJobConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxJobConfigurationUpdateParamsCompletionStatus string

const (
	SandboxJobConfigurationUpdateParamsCompletionStatusComplete         SandboxJobConfigurationUpdateParamsCompletionStatus = "complete"
	SandboxJobConfigurationUpdateParamsCompletionStatusReauthError      SandboxJobConfigurationUpdateParamsCompletionStatus = "reauth_error"
	SandboxJobConfigurationUpdateParamsCompletionStatusPermissionsError SandboxJobConfigurationUpdateParamsCompletionStatus = "permissions_error"
	SandboxJobConfigurationUpdateParamsCompletionStatusError            SandboxJobConfigurationUpdateParamsCompletionStatus = "error"
)

type SandboxJobConfigurationUpdateParamsType string

const (
	SandboxJobConfigurationUpdateParamsTypeDataSyncAll SandboxJobConfigurationUpdateParamsType = "data_sync_all"
)
