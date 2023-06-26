// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISPayStatementService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHRISPayStatementService] method
// instead.
type HRISPayStatementService struct {
	Options []option.RequestOption
}

// NewHRISPayStatementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHRISPayStatementService(opts ...option.RequestOption) (r *HRISPayStatementService) {
	r = &HRISPayStatementService{}
	r.Options = opts
	return
}

// Read detailed pay statements for each individual.
//
// Deduction and contribution types are supported by the payroll systems that
// support Benefits.
func (r *HRISPayStatementService) GetMany(ctx context.Context, body HRISPayStatementGetManyParams, opts ...option.RequestOption) (res *shared.ResponsesPage[PayStatementResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/pay-statement"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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

// Read detailed pay statements for each individual.
//
// Deduction and contribution types are supported by the payroll systems that
// support Benefits.
func (r *HRISPayStatementService) GetManyAutoPaging(ctx context.Context, body HRISPayStatementGetManyParams, opts ...option.RequestOption) *shared.ResponsesPageAutoPager[PayStatementResponse] {
	return shared.NewResponsesPageAutoPager(r.GetMany(ctx, body, opts...))
}

type PayStatement struct {
	// A stable Finch `id` (UUID v4) for an individual in the company
	IndividualID string `json:"individual_id"`
	// The type of the payment associated with the pay statement.
	Type PayStatementType `json:"type,nullable"`
	// The payment method.
	PaymentMethod PayStatementPaymentMethod `json:"payment_method,nullable"`
	// The number of hours worked for this pay period
	TotalHours int64 `json:"total_hours,nullable"`
	GrossPay   Money `json:"gross_pay,nullable"`
	NetPay     Money `json:"net_pay,nullable"`
	// The array of earnings objects associated with this pay statement
	Earnings []PayStatementEarnings `json:"earnings,nullable"`
	// The array of taxes objects associated with this pay statement.
	Taxes []PayStatementTaxes `json:"taxes,nullable"`
	// The array of deductions objects associated with this pay statement.
	EmployeeDeductions    []PayStatementEmployeeDeductions    `json:"employee_deductions,nullable"`
	EmployerContributions []PayStatementEmployerContributions `json:"employer_contributions,nullable"`
	JSON                  payStatementJSON
}

// payStatementJSON contains the JSON metadata for the struct [PayStatement]
type payStatementJSON struct {
	IndividualID          apijson.Field
	Type                  apijson.Field
	PaymentMethod         apijson.Field
	TotalHours            apijson.Field
	GrossPay              apijson.Field
	NetPay                apijson.Field
	Earnings              apijson.Field
	Taxes                 apijson.Field
	EmployeeDeductions    apijson.Field
	EmployerContributions apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PayStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the payment associated with the pay statement.
type PayStatementType string

const (
	PayStatementTypeRegularPayroll  PayStatementType = "regular_payroll"
	PayStatementTypeOffCyclePayroll PayStatementType = "off_cycle_payroll"
	PayStatementTypeOneTimePayment  PayStatementType = "one_time_payment"
)

// The payment method.
type PayStatementPaymentMethod string

const (
	PayStatementPaymentMethodCheck         PayStatementPaymentMethod = "check"
	PayStatementPaymentMethodDirectDeposit PayStatementPaymentMethod = "direct_deposit"
)

type PayStatementEarnings struct {
	// The type of earning.
	Type PayStatementEarningsType `json:"type,nullable"`
	// The exact name of the deduction from the pay statement.
	Name string `json:"name,nullable"`
	// The earnings amount in cents.
	Amount int64 `json:"amount,nullable"`
	// The earnings currency code.
	Currency string `json:"currency,nullable"`
	// The number of hours associated with this earning. (For salaried employees, this
	// could be hours per pay period, `0` or `null`, depending on the provider).
	Hours float64 `json:"hours,nullable"`
	JSON  payStatementEarningsJSON
}

// payStatementEarningsJSON contains the JSON metadata for the struct
// [PayStatementEarnings]
type payStatementEarningsJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Hours       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEarnings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of earning.
type PayStatementEarningsType string

const (
	PayStatementEarningsType1099           PayStatementEarningsType = "1099"
	PayStatementEarningsTypeSalary         PayStatementEarningsType = "salary"
	PayStatementEarningsTypeWage           PayStatementEarningsType = "wage"
	PayStatementEarningsTypeReimbursement  PayStatementEarningsType = "reimbursement"
	PayStatementEarningsTypeOvertime       PayStatementEarningsType = "overtime"
	PayStatementEarningsTypeSeverance      PayStatementEarningsType = "severance"
	PayStatementEarningsTypeDoubleOvertime PayStatementEarningsType = "double_overtime"
	PayStatementEarningsTypePto            PayStatementEarningsType = "pto"
	PayStatementEarningsTypeSick           PayStatementEarningsType = "sick"
	PayStatementEarningsTypeBonus          PayStatementEarningsType = "bonus"
	PayStatementEarningsTypeCommission     PayStatementEarningsType = "commission"
	PayStatementEarningsTypeTips           PayStatementEarningsType = "tips"
	PayStatementEarningsTypeOther          PayStatementEarningsType = "other"
)

type PayStatementTaxes struct {
	// The type of taxes.
	Type PayStatementTaxesType `json:"type,nullable"`
	// The exact name of tax from the pay statement.
	Name string `json:"name,nullable"`
	// `true` if the amount is paid by the employers.
	Employer bool `json:"employer,nullable"`
	// The tax amount in cents.
	Amount int64 `json:"amount,nullable"`
	// The currency code.
	Currency string `json:"currency,nullable"`
	JSON     payStatementTaxesJSON
}

// payStatementTaxesJSON contains the JSON metadata for the struct
// [PayStatementTaxes]
type payStatementTaxesJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	Employer    apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementTaxes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of taxes.
type PayStatementTaxesType string

const (
	PayStatementTaxesTypeState   PayStatementTaxesType = "state"
	PayStatementTaxesTypeFederal PayStatementTaxesType = "federal"
	PayStatementTaxesTypeLocal   PayStatementTaxesType = "local"
	PayStatementTaxesTypeFica    PayStatementTaxesType = "fica"
)

type PayStatementEmployeeDeductions struct {
	// The deduction name from the pay statement.
	Name string `json:"name,nullable"`
	// The deduction amount in cents.
	Amount int64 `json:"amount,nullable"`
	// The deduction currency.
	Currency string `json:"currency,nullable"`
	// Boolean indicating if the deduction is pre-tax.
	PreTax bool `json:"pre_tax,nullable"`
	// Type of benefit.
	Type BenefitType `json:"type,nullable"`
	JSON payStatementEmployeeDeductionsJSON
}

// payStatementEmployeeDeductionsJSON contains the JSON metadata for the struct
// [PayStatementEmployeeDeductions]
type payStatementEmployeeDeductionsJSON struct {
	Name        apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEmployeeDeductions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PayStatementEmployerContributions struct {
	// The contribution name from the pay statement.
	Name string `json:"name,nullable"`
	// The contribution amount in cents.
	Amount int64 `json:"amount,nullable"`
	// The contribution currency.
	Currency string `json:"currency,nullable"`
	// Type of benefit.
	Type BenefitType `json:"type,nullable"`
	JSON payStatementEmployerContributionsJSON
}

// payStatementEmployerContributionsJSON contains the JSON metadata for the struct
// [PayStatementEmployerContributions]
type payStatementEmployerContributionsJSON struct {
	Name        apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEmployerContributions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PayStatementResponse struct {
	PaymentID string                   `json:"payment_id"`
	Code      int64                    `json:"code"`
	Body      PayStatementResponseBody `json:"body"`
	JSON      payStatementResponseJSON
}

// payStatementResponseJSON contains the JSON metadata for the struct
// [PayStatementResponse]
type payStatementResponseJSON struct {
	PaymentID   apijson.Field
	Code        apijson.Field
	Body        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PayStatementResponseBody struct {
	Paging Paging `json:"paging"`
	// The array of pay statements for the current payment.
	PayStatements []PayStatement `json:"pay_statements"`
	JSON          payStatementResponseBodyJSON
}

// payStatementResponseBodyJSON contains the JSON metadata for the struct
// [PayStatementResponseBody]
type payStatementResponseBodyJSON struct {
	Paging        apijson.Field
	PayStatements apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PayStatementResponseBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HRISPayStatementGetManyParams struct {
	// The array of batch requests.
	Requests param.Field[[]HRISPayStatementGetManyParamsRequests] `json:"requests,required"`
}

func (r HRISPayStatementGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISPayStatementGetManyParamsRequests struct {
	// A stable Finch `id` (UUID v4) for a payment.
	PaymentID param.Field[string] `json:"payment_id,required"`
	// Number of pay statements to return (defaults to all).
	Limit param.Field[int64] `json:"limit"`
	// Index to start from.
	Offset param.Field[int64] `json:"offset"`
}

func (r HRISPayStatementGetManyParamsRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
