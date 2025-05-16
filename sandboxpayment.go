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
	EndDate       param.Field[string]                                `json:"end_date"`
	PayStatements param.Field[[]SandboxPaymentNewParamsPayStatement] `json:"pay_statements"`
	StartDate     param.Field[string]                                `json:"start_date"`
}

func (r SandboxPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatement struct {
	// The array of earnings objects associated with this pay statement
	Earnings param.Field[[]SandboxPaymentNewParamsPayStatementsEarning] `json:"earnings"`
	// The array of deductions objects associated with this pay statement.
	EmployeeDeductions    param.Field[[]SandboxPaymentNewParamsPayStatementsEmployeeDeduction]    `json:"employee_deductions"`
	EmployerContributions param.Field[[]SandboxPaymentNewParamsPayStatementsEmployerContribution] `json:"employer_contributions"`
	GrossPay              param.Field[MoneyParam]                                                 `json:"gross_pay"`
	// A stable Finch `id` (UUID v4) for an individual in the company
	IndividualID param.Field[string]     `json:"individual_id"`
	NetPay       param.Field[MoneyParam] `json:"net_pay"`
	// The payment method.
	PaymentMethod param.Field[SandboxPaymentNewParamsPayStatementsPaymentMethod] `json:"payment_method"`
	// The array of taxes objects associated with this pay statement.
	Taxes param.Field[[]SandboxPaymentNewParamsPayStatementsTax] `json:"taxes"`
	// The number of hours worked for this pay period
	TotalHours param.Field[float64] `json:"total_hours"`
	// The type of the payment associated with the pay statement.
	Type param.Field[SandboxPaymentNewParamsPayStatementsType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEarning struct {
	// The earnings amount in cents.
	Amount     param.Field[int64]                                                  `json:"amount"`
	Attributes param.Field[SandboxPaymentNewParamsPayStatementsEarningsAttributes] `json:"attributes"`
	// The earnings currency code.
	Currency param.Field[string] `json:"currency"`
	// The number of hours associated with this earning. (For salaried employees, this
	// could be hours per pay period, `0` or `null`, depending on the provider).
	Hours param.Field[float64] `json:"hours"`
	// The exact name of the deduction from the pay statement.
	Name param.Field[string] `json:"name"`
	// The type of earning.
	Type param.Field[SandboxPaymentNewParamsPayStatementsEarningsType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsEarning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEarningsAttributes struct {
	Metadata param.Field[SandboxPaymentNewParamsPayStatementsEarningsAttributesMetadata] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsEarningsAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEarningsAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsEarningsAttributesMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of earning.
type SandboxPaymentNewParamsPayStatementsEarningsType string

const (
	SandboxPaymentNewParamsPayStatementsEarningsTypeSalary         SandboxPaymentNewParamsPayStatementsEarningsType = "salary"
	SandboxPaymentNewParamsPayStatementsEarningsTypeWage           SandboxPaymentNewParamsPayStatementsEarningsType = "wage"
	SandboxPaymentNewParamsPayStatementsEarningsTypeReimbursement  SandboxPaymentNewParamsPayStatementsEarningsType = "reimbursement"
	SandboxPaymentNewParamsPayStatementsEarningsTypeOvertime       SandboxPaymentNewParamsPayStatementsEarningsType = "overtime"
	SandboxPaymentNewParamsPayStatementsEarningsTypeSeverance      SandboxPaymentNewParamsPayStatementsEarningsType = "severance"
	SandboxPaymentNewParamsPayStatementsEarningsTypeDoubleOvertime SandboxPaymentNewParamsPayStatementsEarningsType = "double_overtime"
	SandboxPaymentNewParamsPayStatementsEarningsTypePto            SandboxPaymentNewParamsPayStatementsEarningsType = "pto"
	SandboxPaymentNewParamsPayStatementsEarningsTypeSick           SandboxPaymentNewParamsPayStatementsEarningsType = "sick"
	SandboxPaymentNewParamsPayStatementsEarningsTypeBonus          SandboxPaymentNewParamsPayStatementsEarningsType = "bonus"
	SandboxPaymentNewParamsPayStatementsEarningsTypeCommission     SandboxPaymentNewParamsPayStatementsEarningsType = "commission"
	SandboxPaymentNewParamsPayStatementsEarningsTypeTips           SandboxPaymentNewParamsPayStatementsEarningsType = "tips"
	SandboxPaymentNewParamsPayStatementsEarningsType1099           SandboxPaymentNewParamsPayStatementsEarningsType = "1099"
	SandboxPaymentNewParamsPayStatementsEarningsTypeOther          SandboxPaymentNewParamsPayStatementsEarningsType = "other"
)

func (r SandboxPaymentNewParamsPayStatementsEarningsType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsEarningsTypeSalary, SandboxPaymentNewParamsPayStatementsEarningsTypeWage, SandboxPaymentNewParamsPayStatementsEarningsTypeReimbursement, SandboxPaymentNewParamsPayStatementsEarningsTypeOvertime, SandboxPaymentNewParamsPayStatementsEarningsTypeSeverance, SandboxPaymentNewParamsPayStatementsEarningsTypeDoubleOvertime, SandboxPaymentNewParamsPayStatementsEarningsTypePto, SandboxPaymentNewParamsPayStatementsEarningsTypeSick, SandboxPaymentNewParamsPayStatementsEarningsTypeBonus, SandboxPaymentNewParamsPayStatementsEarningsTypeCommission, SandboxPaymentNewParamsPayStatementsEarningsTypeTips, SandboxPaymentNewParamsPayStatementsEarningsType1099, SandboxPaymentNewParamsPayStatementsEarningsTypeOther:
		return true
	}
	return false
}

type SandboxPaymentNewParamsPayStatementsEmployeeDeduction struct {
	// The deduction amount in cents.
	Amount     param.Field[int64]                                                            `json:"amount"`
	Attributes param.Field[SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributes] `json:"attributes"`
	// The deduction currency.
	Currency param.Field[string] `json:"currency"`
	// The deduction name from the pay statement.
	Name param.Field[string] `json:"name"`
	// Boolean indicating if the deduction is pre-tax.
	PreTax param.Field[bool] `json:"pre_tax"`
	// Type of benefit.
	Type param.Field[BenefitType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployeeDeduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributes struct {
	Metadata param.Field[SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributesMetadata] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributesMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployerContribution struct {
	// The contribution amount in cents.
	Amount     param.Field[int64]                                                               `json:"amount"`
	Attributes param.Field[SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributes] `json:"attributes"`
	// The contribution currency.
	Currency param.Field[string] `json:"currency"`
	// The contribution name from the pay statement.
	Name param.Field[string] `json:"name"`
	// Type of benefit.
	Type param.Field[BenefitType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployerContribution) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributes struct {
	Metadata param.Field[SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributesMetadata] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributesMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The payment method.
type SandboxPaymentNewParamsPayStatementsPaymentMethod string

const (
	SandboxPaymentNewParamsPayStatementsPaymentMethodCheck         SandboxPaymentNewParamsPayStatementsPaymentMethod = "check"
	SandboxPaymentNewParamsPayStatementsPaymentMethodDirectDeposit SandboxPaymentNewParamsPayStatementsPaymentMethod = "direct_deposit"
	SandboxPaymentNewParamsPayStatementsPaymentMethodOther         SandboxPaymentNewParamsPayStatementsPaymentMethod = "other"
)

func (r SandboxPaymentNewParamsPayStatementsPaymentMethod) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsPaymentMethodCheck, SandboxPaymentNewParamsPayStatementsPaymentMethodDirectDeposit, SandboxPaymentNewParamsPayStatementsPaymentMethodOther:
		return true
	}
	return false
}

type SandboxPaymentNewParamsPayStatementsTax struct {
	// The tax amount in cents.
	Amount     param.Field[int64]                                               `json:"amount"`
	Attributes param.Field[SandboxPaymentNewParamsPayStatementsTaxesAttributes] `json:"attributes"`
	// The currency code.
	Currency param.Field[string] `json:"currency"`
	// `true` if the amount is paid by the employers.
	Employer param.Field[bool] `json:"employer"`
	// The exact name of tax from the pay statement.
	Name param.Field[string] `json:"name"`
	// The type of taxes.
	Type param.Field[SandboxPaymentNewParamsPayStatementsTaxesType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsTax) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsTaxesAttributes struct {
	Metadata param.Field[SandboxPaymentNewParamsPayStatementsTaxesAttributesMetadata] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsTaxesAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsTaxesAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r SandboxPaymentNewParamsPayStatementsTaxesAttributesMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of taxes.
type SandboxPaymentNewParamsPayStatementsTaxesType string

const (
	SandboxPaymentNewParamsPayStatementsTaxesTypeState   SandboxPaymentNewParamsPayStatementsTaxesType = "state"
	SandboxPaymentNewParamsPayStatementsTaxesTypeFederal SandboxPaymentNewParamsPayStatementsTaxesType = "federal"
	SandboxPaymentNewParamsPayStatementsTaxesTypeLocal   SandboxPaymentNewParamsPayStatementsTaxesType = "local"
	SandboxPaymentNewParamsPayStatementsTaxesTypeFica    SandboxPaymentNewParamsPayStatementsTaxesType = "fica"
)

func (r SandboxPaymentNewParamsPayStatementsTaxesType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsTaxesTypeState, SandboxPaymentNewParamsPayStatementsTaxesTypeFederal, SandboxPaymentNewParamsPayStatementsTaxesTypeLocal, SandboxPaymentNewParamsPayStatementsTaxesTypeFica:
		return true
	}
	return false
}

// The type of the payment associated with the pay statement.
type SandboxPaymentNewParamsPayStatementsType string

const (
	SandboxPaymentNewParamsPayStatementsTypeRegularPayroll  SandboxPaymentNewParamsPayStatementsType = "regular_payroll"
	SandboxPaymentNewParamsPayStatementsTypeOffCyclePayroll SandboxPaymentNewParamsPayStatementsType = "off_cycle_payroll"
	SandboxPaymentNewParamsPayStatementsTypeOneTimePayment  SandboxPaymentNewParamsPayStatementsType = "one_time_payment"
)

func (r SandboxPaymentNewParamsPayStatementsType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsTypeRegularPayroll, SandboxPaymentNewParamsPayStatementsTypeOffCyclePayroll, SandboxPaymentNewParamsPayStatementsTypeOneTimePayment:
		return true
	}
	return false
}
