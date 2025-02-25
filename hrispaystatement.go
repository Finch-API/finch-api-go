// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
	"github.com/Finch-API/finch-api-go/shared"
)

// HRISPayStatementService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISPayStatementService] method instead.
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
// supports Benefits.
func (r *HRISPayStatementService) GetMany(ctx context.Context, body HRISPayStatementGetManyParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[PayStatementResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
// supports Benefits.
func (r *HRISPayStatementService) GetManyAutoPaging(ctx context.Context, body HRISPayStatementGetManyParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[PayStatementResponse] {
	return pagination.NewResponsesPageAutoPager(r.GetMany(ctx, body, opts...))
}

type PayStatement struct {
	// The array of earnings objects associated with this pay statement
	Earnings []PayStatementEarning `json:"earnings,nullable"`
	// The array of deductions objects associated with this pay statement.
	EmployeeDeductions    []PayStatementEmployeeDeduction    `json:"employee_deductions,nullable"`
	EmployerContributions []PayStatementEmployerContribution `json:"employer_contributions,nullable"`
	GrossPay              Money                              `json:"gross_pay,nullable"`
	// A stable Finch `id` (UUID v4) for an individual in the company
	IndividualID string `json:"individual_id"`
	NetPay       Money  `json:"net_pay,nullable"`
	// The payment method.
	PaymentMethod PayStatementPaymentMethod `json:"payment_method,nullable"`
	// The array of taxes objects associated with this pay statement.
	Taxes []PayStatementTax `json:"taxes,nullable"`
	// The number of hours worked for this pay period
	TotalHours float64 `json:"total_hours,nullable"`
	// The type of the payment associated with the pay statement.
	Type PayStatementType `json:"type,nullable"`
	JSON payStatementJSON `json:"-"`
}

// payStatementJSON contains the JSON metadata for the struct [PayStatement]
type payStatementJSON struct {
	Earnings              apijson.Field
	EmployeeDeductions    apijson.Field
	EmployerContributions apijson.Field
	GrossPay              apijson.Field
	IndividualID          apijson.Field
	NetPay                apijson.Field
	PaymentMethod         apijson.Field
	Taxes                 apijson.Field
	TotalHours            apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PayStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementJSON) RawJSON() string {
	return r.raw
}

type PayStatementEarning struct {
	// The earnings amount in cents.
	Amount     int64                          `json:"amount,nullable"`
	Attributes PayStatementEarningsAttributes `json:"attributes,nullable"`
	// The earnings currency code.
	Currency string `json:"currency,nullable"`
	// The number of hours associated with this earning. (For salaried employees, this
	// could be hours per pay period, `0` or `null`, depending on the provider).
	Hours float64 `json:"hours,nullable"`
	// The exact name of the deduction from the pay statement.
	Name string `json:"name,nullable"`
	// The type of earning.
	Type PayStatementEarningsType `json:"type,nullable"`
	JSON payStatementEarningJSON  `json:"-"`
}

// payStatementEarningJSON contains the JSON metadata for the struct
// [PayStatementEarning]
type payStatementEarningJSON struct {
	Amount      apijson.Field
	Attributes  apijson.Field
	Currency    apijson.Field
	Hours       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementEarningJSON) RawJSON() string {
	return r.raw
}

type PayStatementEarningsAttributes struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}             `json:"metadata"`
	JSON     payStatementEarningsAttributesJSON `json:"-"`
}

// payStatementEarningsAttributesJSON contains the JSON metadata for the struct
// [PayStatementEarningsAttributes]
type payStatementEarningsAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEarningsAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementEarningsAttributesJSON) RawJSON() string {
	return r.raw
}

// The type of earning.
type PayStatementEarningsType string

const (
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
	PayStatementEarningsType1099           PayStatementEarningsType = "1099"
	PayStatementEarningsTypeOther          PayStatementEarningsType = "other"
)

func (r PayStatementEarningsType) IsKnown() bool {
	switch r {
	case PayStatementEarningsTypeSalary, PayStatementEarningsTypeWage, PayStatementEarningsTypeReimbursement, PayStatementEarningsTypeOvertime, PayStatementEarningsTypeSeverance, PayStatementEarningsTypeDoubleOvertime, PayStatementEarningsTypePto, PayStatementEarningsTypeSick, PayStatementEarningsTypeBonus, PayStatementEarningsTypeCommission, PayStatementEarningsTypeTips, PayStatementEarningsType1099, PayStatementEarningsTypeOther:
		return true
	}
	return false
}

type PayStatementEmployeeDeduction struct {
	// The deduction amount in cents.
	Amount     int64                                    `json:"amount,nullable"`
	Attributes PayStatementEmployeeDeductionsAttributes `json:"attributes,nullable"`
	// The deduction currency.
	Currency string `json:"currency,nullable"`
	// The deduction name from the pay statement.
	Name string `json:"name,nullable"`
	// Boolean indicating if the deduction is pre-tax.
	PreTax bool `json:"pre_tax,nullable"`
	// Type of benefit.
	Type BenefitType                       `json:"type,nullable"`
	JSON payStatementEmployeeDeductionJSON `json:"-"`
}

// payStatementEmployeeDeductionJSON contains the JSON metadata for the struct
// [PayStatementEmployeeDeduction]
type payStatementEmployeeDeductionJSON struct {
	Amount      apijson.Field
	Attributes  apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEmployeeDeduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementEmployeeDeductionJSON) RawJSON() string {
	return r.raw
}

type PayStatementEmployeeDeductionsAttributes struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}                       `json:"metadata"`
	JSON     payStatementEmployeeDeductionsAttributesJSON `json:"-"`
}

// payStatementEmployeeDeductionsAttributesJSON contains the JSON metadata for the
// struct [PayStatementEmployeeDeductionsAttributes]
type payStatementEmployeeDeductionsAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEmployeeDeductionsAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementEmployeeDeductionsAttributesJSON) RawJSON() string {
	return r.raw
}

type PayStatementEmployerContribution struct {
	// The contribution amount in cents.
	Amount     int64                                       `json:"amount,nullable"`
	Attributes PayStatementEmployerContributionsAttributes `json:"attributes,nullable"`
	// The contribution currency.
	Currency string `json:"currency,nullable"`
	// The contribution name from the pay statement.
	Name string `json:"name,nullable"`
	// Type of benefit.
	Type BenefitType                          `json:"type,nullable"`
	JSON payStatementEmployerContributionJSON `json:"-"`
}

// payStatementEmployerContributionJSON contains the JSON metadata for the struct
// [PayStatementEmployerContribution]
type payStatementEmployerContributionJSON struct {
	Amount      apijson.Field
	Attributes  apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEmployerContribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementEmployerContributionJSON) RawJSON() string {
	return r.raw
}

type PayStatementEmployerContributionsAttributes struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}                          `json:"metadata"`
	JSON     payStatementEmployerContributionsAttributesJSON `json:"-"`
}

// payStatementEmployerContributionsAttributesJSON contains the JSON metadata for
// the struct [PayStatementEmployerContributionsAttributes]
type payStatementEmployerContributionsAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEmployerContributionsAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementEmployerContributionsAttributesJSON) RawJSON() string {
	return r.raw
}

// The payment method.
type PayStatementPaymentMethod string

const (
	PayStatementPaymentMethodCheck         PayStatementPaymentMethod = "check"
	PayStatementPaymentMethodDirectDeposit PayStatementPaymentMethod = "direct_deposit"
)

func (r PayStatementPaymentMethod) IsKnown() bool {
	switch r {
	case PayStatementPaymentMethodCheck, PayStatementPaymentMethodDirectDeposit:
		return true
	}
	return false
}

type PayStatementTax struct {
	// The tax amount in cents.
	Amount     int64                       `json:"amount,nullable"`
	Attributes PayStatementTaxesAttributes `json:"attributes,nullable"`
	// The currency code.
	Currency string `json:"currency,nullable"`
	// `true` if the amount is paid by the employers.
	Employer bool `json:"employer,nullable"`
	// The exact name of tax from the pay statement.
	Name string `json:"name,nullable"`
	// The type of taxes.
	Type PayStatementTaxesType `json:"type,nullable"`
	JSON payStatementTaxJSON   `json:"-"`
}

// payStatementTaxJSON contains the JSON metadata for the struct [PayStatementTax]
type payStatementTaxJSON struct {
	Amount      apijson.Field
	Attributes  apijson.Field
	Currency    apijson.Field
	Employer    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementTax) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementTaxJSON) RawJSON() string {
	return r.raw
}

type PayStatementTaxesAttributes struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}          `json:"metadata"`
	JSON     payStatementTaxesAttributesJSON `json:"-"`
}

// payStatementTaxesAttributesJSON contains the JSON metadata for the struct
// [PayStatementTaxesAttributes]
type payStatementTaxesAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementTaxesAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementTaxesAttributesJSON) RawJSON() string {
	return r.raw
}

// The type of taxes.
type PayStatementTaxesType string

const (
	PayStatementTaxesTypeState   PayStatementTaxesType = "state"
	PayStatementTaxesTypeFederal PayStatementTaxesType = "federal"
	PayStatementTaxesTypeLocal   PayStatementTaxesType = "local"
	PayStatementTaxesTypeFica    PayStatementTaxesType = "fica"
)

func (r PayStatementTaxesType) IsKnown() bool {
	switch r {
	case PayStatementTaxesTypeState, PayStatementTaxesTypeFederal, PayStatementTaxesTypeLocal, PayStatementTaxesTypeFica:
		return true
	}
	return false
}

// The type of the payment associated with the pay statement.
type PayStatementType string

const (
	PayStatementTypeRegularPayroll  PayStatementType = "regular_payroll"
	PayStatementTypeOffCyclePayroll PayStatementType = "off_cycle_payroll"
	PayStatementTypeOneTimePayment  PayStatementType = "one_time_payment"
)

func (r PayStatementType) IsKnown() bool {
	switch r {
	case PayStatementTypeRegularPayroll, PayStatementTypeOffCyclePayroll, PayStatementTypeOneTimePayment:
		return true
	}
	return false
}

type PayStatementResponse struct {
	Body      PayStatementResponseBody `json:"body"`
	Code      int64                    `json:"code"`
	PaymentID string                   `json:"payment_id"`
	JSON      payStatementResponseJSON `json:"-"`
}

// payStatementResponseJSON contains the JSON metadata for the struct
// [PayStatementResponse]
type payStatementResponseJSON struct {
	Body        apijson.Field
	Code        apijson.Field
	PaymentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBody struct {
	Paging shared.Paging `json:"paging"`
	// The array of pay statements for the current payment.
	PayStatements []PayStatement               `json:"pay_statements"`
	JSON          payStatementResponseBodyJSON `json:"-"`
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

func (r payStatementResponseBodyJSON) RawJSON() string {
	return r.raw
}

type HRISPayStatementGetManyParams struct {
	// The array of batch requests.
	Requests param.Field[[]HRISPayStatementGetManyParamsRequest] `json:"requests,required"`
}

func (r HRISPayStatementGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISPayStatementGetManyParamsRequest struct {
	// A stable Finch `id` (UUID v4) for a payment.
	PaymentID param.Field[string] `json:"payment_id,required"`
	// Number of pay statements to return (defaults to all).
	Limit param.Field[int64] `json:"limit"`
	// Index to start from.
	Offset param.Field[int64] `json:"offset"`
}

func (r HRISPayStatementGetManyParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
