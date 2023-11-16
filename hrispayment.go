// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISPaymentService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHRISPaymentService] method instead.
type HRISPaymentService struct {
	Options []option.RequestOption
}

// NewHRISPaymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISPaymentService(opts ...option.RequestOption) (r *HRISPaymentService) {
	r = &HRISPaymentService{}
	r.Options = opts
	return
}

// Read payroll and contractor related payments by the company.
func (r *HRISPaymentService) List(ctx context.Context, query HRISPaymentListParams, opts ...option.RequestOption) (res *shared.SinglePage[Payment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/payment"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Read payroll and contractor related payments by the company.
func (r *HRISPaymentService) ListAutoPaging(ctx context.Context, query HRISPaymentListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[Payment] {
	return shared.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type Payment struct {
	// The unique id for the payment.
	ID            string `json:"id"`
	CompanyDebit  Money  `json:"company_debit,nullable"`
	DebitDate     string `json:"debit_date,nullable"`
	EmployeeTaxes Money  `json:"employee_taxes,nullable"`
	EmployerTaxes Money  `json:"employer_taxes,nullable"`
	GrossPay      Money  `json:"gross_pay,nullable"`
	// Array of every individual on this payment.
	IndividualIDs []string `json:"individual_ids,nullable"`
	NetPay        Money    `json:"net_pay,nullable"`
	PayDate       string   `json:"pay_date,nullable"`
	// The pay period object.
	PayPeriod PaymentPayPeriod `json:"pay_period,nullable"`
	JSON      paymentJSON      `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	ID            apijson.Field
	CompanyDebit  apijson.Field
	DebitDate     apijson.Field
	EmployeeTaxes apijson.Field
	EmployerTaxes apijson.Field
	GrossPay      apijson.Field
	IndividualIDs apijson.Field
	NetPay        apijson.Field
	PayDate       apijson.Field
	PayPeriod     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pay period object.
type PaymentPayPeriod struct {
	EndDate   string               `json:"end_date,nullable"`
	StartDate string               `json:"start_date,nullable"`
	JSON      paymentPayPeriodJSON `json:"-"`
}

// paymentPayPeriodJSON contains the JSON metadata for the struct
// [PaymentPayPeriod]
type paymentPayPeriodJSON struct {
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentPayPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HRISPaymentListParams struct {
	// The end date to retrieve payments by a company (inclusive) in `YYYY-MM-DD`
	// format.
	EndDate param.Field[time.Time] `query:"end_date,required" format:"date"`
	// The start date to retrieve payments by a company (inclusive) in `YYYY-MM-DD`
	// format.
	StartDate param.Field[time.Time] `query:"start_date,required" format:"date"`
}

// URLQuery serializes [HRISPaymentListParams]'s query parameters as `url.Values`.
func (r HRISPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
