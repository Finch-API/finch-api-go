// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"
	"time"

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
	opts = slices.Concat(r.Options, opts)
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
	EndDate param.Field[time.Time] `json:"end_date" format:"date"`
	// Array of pay statements to include in the payment.
	PayStatements param.Field[[]SandboxPaymentNewParamsPayStatement] `json:"pay_statements"`
	StartDate     param.Field[time.Time]                             `json:"start_date" format:"date"`
}

func (r SandboxPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatement struct {
	IndividualID          param.Field[string]                                                     `json:"individual_id,required" format:"uuid"`
	Earnings              param.Field[[]SandboxPaymentNewParamsPayStatementsEarning]              `json:"earnings"`
	EmployeeDeductions    param.Field[[]SandboxPaymentNewParamsPayStatementsEmployeeDeduction]    `json:"employee_deductions"`
	EmployerContributions param.Field[[]SandboxPaymentNewParamsPayStatementsEmployerContribution] `json:"employer_contributions"`
	GrossPay              param.Field[int64]                                                      `json:"gross_pay"`
	NetPay                param.Field[int64]                                                      `json:"net_pay"`
	PaymentMethod         param.Field[SandboxPaymentNewParamsPayStatementsPaymentMethod]          `json:"payment_method"`
	Taxes                 param.Field[[]SandboxPaymentNewParamsPayStatementsTax]                  `json:"taxes"`
	TotalHours            param.Field[float64]                                                    `json:"total_hours"`
	Type                  param.Field[SandboxPaymentNewParamsPayStatementsType]                   `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEarning struct {
	Amount param.Field[int64]                                            `json:"amount"`
	Hours  param.Field[float64]                                          `json:"hours"`
	Name   param.Field[string]                                           `json:"name"`
	Type   param.Field[SandboxPaymentNewParamsPayStatementsEarningsType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsEarning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEarningsType string

const (
	SandboxPaymentNewParamsPayStatementsEarningsTypeBonus          SandboxPaymentNewParamsPayStatementsEarningsType = "bonus"
	SandboxPaymentNewParamsPayStatementsEarningsTypeCommission     SandboxPaymentNewParamsPayStatementsEarningsType = "commission"
	SandboxPaymentNewParamsPayStatementsEarningsTypeDoubleOvertime SandboxPaymentNewParamsPayStatementsEarningsType = "double_overtime"
	SandboxPaymentNewParamsPayStatementsEarningsTypeOther          SandboxPaymentNewParamsPayStatementsEarningsType = "other"
	SandboxPaymentNewParamsPayStatementsEarningsTypeOvertime       SandboxPaymentNewParamsPayStatementsEarningsType = "overtime"
	SandboxPaymentNewParamsPayStatementsEarningsTypePto            SandboxPaymentNewParamsPayStatementsEarningsType = "pto"
	SandboxPaymentNewParamsPayStatementsEarningsTypeReimbursement  SandboxPaymentNewParamsPayStatementsEarningsType = "reimbursement"
	SandboxPaymentNewParamsPayStatementsEarningsTypeSalary         SandboxPaymentNewParamsPayStatementsEarningsType = "salary"
	SandboxPaymentNewParamsPayStatementsEarningsTypeSeverance      SandboxPaymentNewParamsPayStatementsEarningsType = "severance"
	SandboxPaymentNewParamsPayStatementsEarningsTypeSick           SandboxPaymentNewParamsPayStatementsEarningsType = "sick"
	SandboxPaymentNewParamsPayStatementsEarningsTypeTips           SandboxPaymentNewParamsPayStatementsEarningsType = "tips"
	SandboxPaymentNewParamsPayStatementsEarningsTypeWage           SandboxPaymentNewParamsPayStatementsEarningsType = "wage"
	SandboxPaymentNewParamsPayStatementsEarningsType1099           SandboxPaymentNewParamsPayStatementsEarningsType = "1099"
)

func (r SandboxPaymentNewParamsPayStatementsEarningsType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsEarningsTypeBonus, SandboxPaymentNewParamsPayStatementsEarningsTypeCommission, SandboxPaymentNewParamsPayStatementsEarningsTypeDoubleOvertime, SandboxPaymentNewParamsPayStatementsEarningsTypeOther, SandboxPaymentNewParamsPayStatementsEarningsTypeOvertime, SandboxPaymentNewParamsPayStatementsEarningsTypePto, SandboxPaymentNewParamsPayStatementsEarningsTypeReimbursement, SandboxPaymentNewParamsPayStatementsEarningsTypeSalary, SandboxPaymentNewParamsPayStatementsEarningsTypeSeverance, SandboxPaymentNewParamsPayStatementsEarningsTypeSick, SandboxPaymentNewParamsPayStatementsEarningsTypeTips, SandboxPaymentNewParamsPayStatementsEarningsTypeWage, SandboxPaymentNewParamsPayStatementsEarningsType1099:
		return true
	}
	return false
}

type SandboxPaymentNewParamsPayStatementsEmployeeDeduction struct {
	Amount param.Field[int64]                                                      `json:"amount"`
	Name   param.Field[string]                                                     `json:"name"`
	PreTax param.Field[bool]                                                       `json:"pre_tax"`
	Type   param.Field[SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployeeDeduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType string

const (
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_457             SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "457"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_401k            SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "401k"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_401kRoth        SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "401k_roth"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_401kLoan        SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "401k_loan"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_403b            SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "403b"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_403bRoth        SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "403b_roth"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_457Roth         SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "457_roth"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeCommuter         SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "commuter"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeCustomPostTax    SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "custom_post_tax"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeCustomPreTax     SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "custom_pre_tax"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeFsaDependentCare SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "fsa_dependent_care"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeFsaMedical       SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "fsa_medical"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeHsaPost          SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "hsa_post"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeHsaPre           SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "hsa_pre"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeS125Dental       SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "s125_dental"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeS125Medical      SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "s125_medical"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeS125Vision       SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "s125_vision"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeSimple           SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "simple"
	SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeSimpleIRA        SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType = "simple_ira"
)

func (r SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_457, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_401k, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_401kRoth, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_401kLoan, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_403b, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_403bRoth, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_457Roth, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeCommuter, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeCustomPostTax, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeCustomPreTax, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeFsaDependentCare, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeFsaMedical, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeHsaPost, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeHsaPre, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeS125Dental, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeS125Medical, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeS125Vision, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeSimple, SandboxPaymentNewParamsPayStatementsEmployeeDeductionsTypeSimpleIRA:
		return true
	}
	return false
}

type SandboxPaymentNewParamsPayStatementsEmployerContribution struct {
	Amount param.Field[int64]                                                         `json:"amount"`
	Name   param.Field[string]                                                        `json:"name"`
	Type   param.Field[SandboxPaymentNewParamsPayStatementsEmployerContributionsType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsEmployerContribution) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsEmployerContributionsType string

const (
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_457             SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "457"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_401k            SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "401k"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_401kRoth        SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "401k_roth"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_401kLoan        SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "401k_loan"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_403b            SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "403b"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_403bRoth        SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "403b_roth"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsType_457Roth         SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "457_roth"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeCommuter         SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "commuter"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeCustomPostTax    SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "custom_post_tax"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeCustomPreTax     SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "custom_pre_tax"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeFsaDependentCare SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "fsa_dependent_care"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeFsaMedical       SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "fsa_medical"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeHsaPost          SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "hsa_post"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeHsaPre           SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "hsa_pre"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeS125Dental       SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "s125_dental"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeS125Medical      SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "s125_medical"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeS125Vision       SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "s125_vision"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeSimple           SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "simple"
	SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeSimpleIRA        SandboxPaymentNewParamsPayStatementsEmployerContributionsType = "simple_ira"
)

func (r SandboxPaymentNewParamsPayStatementsEmployerContributionsType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsEmployerContributionsType_457, SandboxPaymentNewParamsPayStatementsEmployerContributionsType_401k, SandboxPaymentNewParamsPayStatementsEmployerContributionsType_401kRoth, SandboxPaymentNewParamsPayStatementsEmployerContributionsType_401kLoan, SandboxPaymentNewParamsPayStatementsEmployerContributionsType_403b, SandboxPaymentNewParamsPayStatementsEmployerContributionsType_403bRoth, SandboxPaymentNewParamsPayStatementsEmployerContributionsType_457Roth, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeCommuter, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeCustomPostTax, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeCustomPreTax, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeFsaDependentCare, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeFsaMedical, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeHsaPost, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeHsaPre, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeS125Dental, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeS125Medical, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeS125Vision, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeSimple, SandboxPaymentNewParamsPayStatementsEmployerContributionsTypeSimpleIRA:
		return true
	}
	return false
}

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
	Amount   param.Field[int64]                                         `json:"amount"`
	Employer param.Field[bool]                                          `json:"employer"`
	Name     param.Field[string]                                        `json:"name"`
	Type     param.Field[SandboxPaymentNewParamsPayStatementsTaxesType] `json:"type"`
}

func (r SandboxPaymentNewParamsPayStatementsTax) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxPaymentNewParamsPayStatementsTaxesType string

const (
	SandboxPaymentNewParamsPayStatementsTaxesTypeFederal SandboxPaymentNewParamsPayStatementsTaxesType = "federal"
	SandboxPaymentNewParamsPayStatementsTaxesTypeFica    SandboxPaymentNewParamsPayStatementsTaxesType = "fica"
	SandboxPaymentNewParamsPayStatementsTaxesTypeLocal   SandboxPaymentNewParamsPayStatementsTaxesType = "local"
	SandboxPaymentNewParamsPayStatementsTaxesTypeState   SandboxPaymentNewParamsPayStatementsTaxesType = "state"
)

func (r SandboxPaymentNewParamsPayStatementsTaxesType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsTaxesTypeFederal, SandboxPaymentNewParamsPayStatementsTaxesTypeFica, SandboxPaymentNewParamsPayStatementsTaxesTypeLocal, SandboxPaymentNewParamsPayStatementsTaxesTypeState:
		return true
	}
	return false
}

type SandboxPaymentNewParamsPayStatementsType string

const (
	SandboxPaymentNewParamsPayStatementsTypeOffCyclePayroll SandboxPaymentNewParamsPayStatementsType = "off_cycle_payroll"
	SandboxPaymentNewParamsPayStatementsTypeOneTimePayment  SandboxPaymentNewParamsPayStatementsType = "one_time_payment"
	SandboxPaymentNewParamsPayStatementsTypeRegularPayroll  SandboxPaymentNewParamsPayStatementsType = "regular_payroll"
)

func (r SandboxPaymentNewParamsPayStatementsType) IsKnown() bool {
	switch r {
	case SandboxPaymentNewParamsPayStatementsTypeOffCyclePayroll, SandboxPaymentNewParamsPayStatementsTypeOneTimePayment, SandboxPaymentNewParamsPayStatementsTypeRegularPayroll:
		return true
	}
	return false
}
