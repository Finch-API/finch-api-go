// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxPaymentService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxPaymentService] method instead.
type SandboxPaymentService struct {
	Options []option.RequestOption
}

// NewSandboxPaymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxPaymentService(opts ...option.RequestOption) (r *SandboxPaymentService) {
	r = &SandboxPaymentService{}
	r.Options = opts
	return
}

// Add a new sandbox payment
func (r *SandboxPaymentService) New(ctx context.Context, body SandboxPaymentNewParams, opts ...option.RequestOption) (res *SandboxPaymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sandbox/payment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SandboxPaymentNewResponse struct {
	// The date of the payment.
	PayDate string `json:"pay_date,required"`
	// The ID of the payment.
	PaymentID string                        `json:"payment_id,required"`
	JSON      sandboxPaymentNewResponseJSON `json:"-"`
}

// sandboxPaymentNewResponseJSON contains the JSON metadata for the struct
// [SandboxPaymentNewResponse]
type sandboxPaymentNewResponseJSON struct {
	PayDate     apijson.Field
	PaymentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxPaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxPaymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxPaymentNewParams struct {
	EndDate       param.Field[string]              `json:"end_date"`
	PayStatements param.Field[[]PayStatementParam] `json:"pay_statements"`
	StartDate     param.Field[string]              `json:"start_date"`
}

func (r SandboxPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
