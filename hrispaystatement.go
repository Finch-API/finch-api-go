// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
	"github.com/tidwall/gjson"
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

type PayStatementResponse struct {
	Body      PayStatementResponseBody `json:"body,required"`
	Code      int64                    `json:"code,required"`
	PaymentID string                   `json:"payment_id,required"`
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
	Code      float64 `json:"code"`
	FinchCode string  `json:"finch_code"`
	Message   string  `json:"message"`
	Name      string  `json:"name"`
	// This field can have the runtime type of [PayStatementResponseBodyObjectPaging].
	Paging interface{} `json:"paging"`
	// This field can have the runtime type of
	// [[]PayStatementResponseBodyObjectPayStatement].
	PayStatements interface{}                  `json:"pay_statements"`
	JSON          payStatementResponseBodyJSON `json:"-"`
	union         PayStatementResponseBodyUnion
}

// payStatementResponseBodyJSON contains the JSON metadata for the struct
// [PayStatementResponseBody]
type payStatementResponseBodyJSON struct {
	Code          apijson.Field
	FinchCode     apijson.Field
	Message       apijson.Field
	Name          apijson.Field
	Paging        apijson.Field
	PayStatements apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r payStatementResponseBodyJSON) RawJSON() string {
	return r.raw
}

func (r *PayStatementResponseBody) UnmarshalJSON(data []byte) (err error) {
	*r = PayStatementResponseBody{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PayStatementResponseBodyUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [PayStatementResponseBodyObject],
// [PayStatementResponseBodyBatchError], [PayStatementResponseBodyObject].
func (r PayStatementResponseBody) AsUnion() PayStatementResponseBodyUnion {
	return r.union
}

// Union satisfied by [PayStatementResponseBodyObject],
// [PayStatementResponseBodyBatchError] or [PayStatementResponseBodyObject].
type PayStatementResponseBodyUnion interface {
	implementsPayStatementResponseBody()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PayStatementResponseBodyUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PayStatementResponseBodyObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PayStatementResponseBodyBatchError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PayStatementResponseBodyObject{}),
		},
	)
}

type PayStatementResponseBodyObject struct {
	Paging        PayStatementResponseBodyObjectPaging         `json:"paging,required"`
	PayStatements []PayStatementResponseBodyObjectPayStatement `json:"pay_statements,required"`
	JSON          payStatementResponseBodyObjectJSON           `json:"-"`
}

// payStatementResponseBodyObjectJSON contains the JSON metadata for the struct
// [PayStatementResponseBodyObject]
type payStatementResponseBodyObjectJSON struct {
	Paging        apijson.Field
	PayStatements apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PayStatementResponseBodyObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectJSON) RawJSON() string {
	return r.raw
}

func (r PayStatementResponseBodyObject) implementsPayStatementResponseBody() {}

type PayStatementResponseBodyObjectPaging struct {
	// The current start index of the returned list of elements
	Offset int64 `json:"offset,required"`
	// The total number of elements for the entire query (not just the given page)
	Count int64                                    `json:"count"`
	JSON  payStatementResponseBodyObjectPagingJSON `json:"-"`
}

// payStatementResponseBodyObjectPagingJSON contains the JSON metadata for the
// struct [PayStatementResponseBodyObjectPaging]
type payStatementResponseBodyObjectPagingJSON struct {
	Offset      apijson.Field
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPagingJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatement struct {
	// The array of earnings objects associated with this pay statement
	Earnings []PayStatementResponseBodyObjectPayStatementsEarning `json:"earnings,required,nullable"`
	// The array of deductions objects associated with this pay statement.
	EmployeeDeductions    []PayStatementResponseBodyObjectPayStatementsEmployeeDeduction    `json:"employee_deductions,required,nullable"`
	EmployerContributions []PayStatementResponseBodyObjectPayStatementsEmployerContribution `json:"employer_contributions,required,nullable"`
	GrossPay              Money                                                             `json:"gross_pay,required,nullable"`
	// A stable Finch `id` (UUID v4) for an individual in the company
	IndividualID string `json:"individual_id,required"`
	NetPay       Money  `json:"net_pay,required,nullable"`
	// The payment method.
	PaymentMethod PayStatementResponseBodyObjectPayStatementsPaymentMethod `json:"payment_method,required,nullable"`
	// The array of taxes objects associated with this pay statement.
	Taxes []PayStatementResponseBodyObjectPayStatementsTax `json:"taxes,required,nullable"`
	// The number of hours worked for this pay period
	TotalHours float64 `json:"total_hours,required,nullable"`
	// The type of the payment associated with the pay statement.
	Type PayStatementResponseBodyObjectPayStatementsType `json:"type,required,nullable"`
	JSON payStatementResponseBodyObjectPayStatementJSON  `json:"-"`
}

// payStatementResponseBodyObjectPayStatementJSON contains the JSON metadata for
// the struct [PayStatementResponseBodyObjectPayStatement]
type payStatementResponseBodyObjectPayStatementJSON struct {
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

func (r *PayStatementResponseBodyObjectPayStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEarning struct {
	// The earnings amount in cents.
	Amount int64 `json:"amount,required,nullable"`
	// The earnings currency code.
	Currency string `json:"currency,required,nullable"`
	// The number of hours associated with this earning. (For salaried employees, this
	// could be hours per pay period, `0` or `null`, depending on the provider).
	Hours float64 `json:"hours,required,nullable"`
	// The exact name of the deduction from the pay statement.
	Name string `json:"name,required,nullable"`
	// The type of earning.
	Type       PayStatementResponseBodyObjectPayStatementsEarningsType       `json:"type,required,nullable"`
	Attributes PayStatementResponseBodyObjectPayStatementsEarningsAttributes `json:"attributes,nullable"`
	JSON       payStatementResponseBodyObjectPayStatementsEarningJSON        `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEarningJSON contains the JSON
// metadata for the struct [PayStatementResponseBodyObjectPayStatementsEarning]
type payStatementResponseBodyObjectPayStatementsEarningJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Hours       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Attributes  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEarningJSON) RawJSON() string {
	return r.raw
}

// The type of earning.
type PayStatementResponseBodyObjectPayStatementsEarningsType string

const (
	PayStatementResponseBodyObjectPayStatementsEarningsTypeSalary         PayStatementResponseBodyObjectPayStatementsEarningsType = "salary"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeWage           PayStatementResponseBodyObjectPayStatementsEarningsType = "wage"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeReimbursement  PayStatementResponseBodyObjectPayStatementsEarningsType = "reimbursement"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeOvertime       PayStatementResponseBodyObjectPayStatementsEarningsType = "overtime"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeSeverance      PayStatementResponseBodyObjectPayStatementsEarningsType = "severance"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeDoubleOvertime PayStatementResponseBodyObjectPayStatementsEarningsType = "double_overtime"
	PayStatementResponseBodyObjectPayStatementsEarningsTypePto            PayStatementResponseBodyObjectPayStatementsEarningsType = "pto"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeSick           PayStatementResponseBodyObjectPayStatementsEarningsType = "sick"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeBonus          PayStatementResponseBodyObjectPayStatementsEarningsType = "bonus"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeCommission     PayStatementResponseBodyObjectPayStatementsEarningsType = "commission"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeTips           PayStatementResponseBodyObjectPayStatementsEarningsType = "tips"
	PayStatementResponseBodyObjectPayStatementsEarningsType1099           PayStatementResponseBodyObjectPayStatementsEarningsType = "1099"
	PayStatementResponseBodyObjectPayStatementsEarningsTypeOther          PayStatementResponseBodyObjectPayStatementsEarningsType = "other"
)

func (r PayStatementResponseBodyObjectPayStatementsEarningsType) IsKnown() bool {
	switch r {
	case PayStatementResponseBodyObjectPayStatementsEarningsTypeSalary, PayStatementResponseBodyObjectPayStatementsEarningsTypeWage, PayStatementResponseBodyObjectPayStatementsEarningsTypeReimbursement, PayStatementResponseBodyObjectPayStatementsEarningsTypeOvertime, PayStatementResponseBodyObjectPayStatementsEarningsTypeSeverance, PayStatementResponseBodyObjectPayStatementsEarningsTypeDoubleOvertime, PayStatementResponseBodyObjectPayStatementsEarningsTypePto, PayStatementResponseBodyObjectPayStatementsEarningsTypeSick, PayStatementResponseBodyObjectPayStatementsEarningsTypeBonus, PayStatementResponseBodyObjectPayStatementsEarningsTypeCommission, PayStatementResponseBodyObjectPayStatementsEarningsTypeTips, PayStatementResponseBodyObjectPayStatementsEarningsType1099, PayStatementResponseBodyObjectPayStatementsEarningsTypeOther:
		return true
	}
	return false
}

type PayStatementResponseBodyObjectPayStatementsEarningsAttributes struct {
	Metadata PayStatementResponseBodyObjectPayStatementsEarningsAttributesMetadata `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsEarningsAttributesJSON     `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEarningsAttributesJSON contains the
// JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEarningsAttributes]
type payStatementResponseBodyObjectPayStatementsEarningsAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEarningsAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEarningsAttributesJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEarningsAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}                                                    `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsEarningsAttributesMetadataJSON `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEarningsAttributesMetadataJSON
// contains the JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEarningsAttributesMetadata]
type payStatementResponseBodyObjectPayStatementsEarningsAttributesMetadataJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEarningsAttributesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEarningsAttributesMetadataJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEmployeeDeduction struct {
	// The deduction amount in cents.
	Amount int64 `json:"amount,required,nullable"`
	// The deduction currency.
	Currency string `json:"currency,required,nullable"`
	// The deduction name from the pay statement.
	Name string `json:"name,required,nullable"`
	// Boolean indicating if the deduction is pre-tax.
	PreTax bool `json:"pre_tax,required,nullable"`
	// Type of benefit.
	Type       BenefitType                                                             `json:"type,required,nullable"`
	Attributes PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributes `json:"attributes,nullable"`
	JSON       payStatementResponseBodyObjectPayStatementsEmployeeDeductionJSON        `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEmployeeDeductionJSON contains the
// JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEmployeeDeduction]
type payStatementResponseBodyObjectPayStatementsEmployeeDeductionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	Attributes  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEmployeeDeduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEmployeeDeductionJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributes struct {
	Metadata PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadata `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesJSON     `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesJSON
// contains the JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributes]
type payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}                                                              `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadataJSON `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadataJSON
// contains the JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadata]
type payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadataJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEmployeeDeductionsAttributesMetadataJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEmployerContribution struct {
	// The contribution currency.
	Currency string `json:"currency,required,nullable"`
	// The contribution name from the pay statement.
	Name string `json:"name,required,nullable"`
	// Type of benefit.
	Type BenefitType `json:"type,required,nullable"`
	// The contribution amount in cents.
	Amount     int64                                                                      `json:"amount,nullable"`
	Attributes PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributes `json:"attributes,nullable"`
	JSON       payStatementResponseBodyObjectPayStatementsEmployerContributionJSON        `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEmployerContributionJSON contains the
// JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEmployerContribution]
type payStatementResponseBodyObjectPayStatementsEmployerContributionJSON struct {
	Currency    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Amount      apijson.Field
	Attributes  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEmployerContribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEmployerContributionJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributes struct {
	Metadata PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadata `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesJSON     `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesJSON
// contains the JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributes]
type payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}                                                                 `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadataJSON `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadataJSON
// contains the JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadata]
type payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadataJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsEmployerContributionsAttributesMetadataJSON) RawJSON() string {
	return r.raw
}

// The payment method.
type PayStatementResponseBodyObjectPayStatementsPaymentMethod string

const (
	PayStatementResponseBodyObjectPayStatementsPaymentMethodCheck         PayStatementResponseBodyObjectPayStatementsPaymentMethod = "check"
	PayStatementResponseBodyObjectPayStatementsPaymentMethodDirectDeposit PayStatementResponseBodyObjectPayStatementsPaymentMethod = "direct_deposit"
	PayStatementResponseBodyObjectPayStatementsPaymentMethodOther         PayStatementResponseBodyObjectPayStatementsPaymentMethod = "other"
)

func (r PayStatementResponseBodyObjectPayStatementsPaymentMethod) IsKnown() bool {
	switch r {
	case PayStatementResponseBodyObjectPayStatementsPaymentMethodCheck, PayStatementResponseBodyObjectPayStatementsPaymentMethodDirectDeposit, PayStatementResponseBodyObjectPayStatementsPaymentMethodOther:
		return true
	}
	return false
}

type PayStatementResponseBodyObjectPayStatementsTax struct {
	// The currency code.
	Currency string `json:"currency,required,nullable"`
	// `true` if the amount is paid by the employers.
	Employer bool `json:"employer,required,nullable"`
	// The exact name of tax from the pay statement.
	Name string `json:"name,required,nullable"`
	// The type of taxes.
	Type PayStatementResponseBodyObjectPayStatementsTaxesType `json:"type,required,nullable"`
	// The tax amount in cents.
	Amount     int64                                                      `json:"amount,nullable"`
	Attributes PayStatementResponseBodyObjectPayStatementsTaxesAttributes `json:"attributes,nullable"`
	JSON       payStatementResponseBodyObjectPayStatementsTaxJSON         `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsTaxJSON contains the JSON metadata
// for the struct [PayStatementResponseBodyObjectPayStatementsTax]
type payStatementResponseBodyObjectPayStatementsTaxJSON struct {
	Currency    apijson.Field
	Employer    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Amount      apijson.Field
	Attributes  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsTax) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsTaxJSON) RawJSON() string {
	return r.raw
}

// The type of taxes.
type PayStatementResponseBodyObjectPayStatementsTaxesType string

const (
	PayStatementResponseBodyObjectPayStatementsTaxesTypeState   PayStatementResponseBodyObjectPayStatementsTaxesType = "state"
	PayStatementResponseBodyObjectPayStatementsTaxesTypeFederal PayStatementResponseBodyObjectPayStatementsTaxesType = "federal"
	PayStatementResponseBodyObjectPayStatementsTaxesTypeLocal   PayStatementResponseBodyObjectPayStatementsTaxesType = "local"
	PayStatementResponseBodyObjectPayStatementsTaxesTypeFica    PayStatementResponseBodyObjectPayStatementsTaxesType = "fica"
)

func (r PayStatementResponseBodyObjectPayStatementsTaxesType) IsKnown() bool {
	switch r {
	case PayStatementResponseBodyObjectPayStatementsTaxesTypeState, PayStatementResponseBodyObjectPayStatementsTaxesTypeFederal, PayStatementResponseBodyObjectPayStatementsTaxesTypeLocal, PayStatementResponseBodyObjectPayStatementsTaxesTypeFica:
		return true
	}
	return false
}

type PayStatementResponseBodyObjectPayStatementsTaxesAttributes struct {
	Metadata PayStatementResponseBodyObjectPayStatementsTaxesAttributesMetadata `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsTaxesAttributesJSON     `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsTaxesAttributesJSON contains the JSON
// metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsTaxesAttributes]
type payStatementResponseBodyObjectPayStatementsTaxesAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsTaxesAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsTaxesAttributesJSON) RawJSON() string {
	return r.raw
}

type PayStatementResponseBodyObjectPayStatementsTaxesAttributesMetadata struct {
	// The metadata to be attached to the entity by existing rules. It is a key-value
	// pairs where the values can be of any type (string, number, boolean, object,
	// array, etc.).
	Metadata map[string]interface{}                                                 `json:"metadata,required"`
	JSON     payStatementResponseBodyObjectPayStatementsTaxesAttributesMetadataJSON `json:"-"`
}

// payStatementResponseBodyObjectPayStatementsTaxesAttributesMetadataJSON contains
// the JSON metadata for the struct
// [PayStatementResponseBodyObjectPayStatementsTaxesAttributesMetadata]
type payStatementResponseBodyObjectPayStatementsTaxesAttributesMetadataJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyObjectPayStatementsTaxesAttributesMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyObjectPayStatementsTaxesAttributesMetadataJSON) RawJSON() string {
	return r.raw
}

// The type of the payment associated with the pay statement.
type PayStatementResponseBodyObjectPayStatementsType string

const (
	PayStatementResponseBodyObjectPayStatementsTypeOffCyclePayroll PayStatementResponseBodyObjectPayStatementsType = "off_cycle_payroll"
	PayStatementResponseBodyObjectPayStatementsTypeOneTimePayment  PayStatementResponseBodyObjectPayStatementsType = "one_time_payment"
	PayStatementResponseBodyObjectPayStatementsTypeRegularPayroll  PayStatementResponseBodyObjectPayStatementsType = "regular_payroll"
)

func (r PayStatementResponseBodyObjectPayStatementsType) IsKnown() bool {
	switch r {
	case PayStatementResponseBodyObjectPayStatementsTypeOffCyclePayroll, PayStatementResponseBodyObjectPayStatementsTypeOneTimePayment, PayStatementResponseBodyObjectPayStatementsTypeRegularPayroll:
		return true
	}
	return false
}

type PayStatementResponseBodyBatchError struct {
	Code      float64                                `json:"code,required"`
	Message   string                                 `json:"message,required"`
	Name      string                                 `json:"name,required"`
	FinchCode string                                 `json:"finch_code"`
	JSON      payStatementResponseBodyBatchErrorJSON `json:"-"`
}

// payStatementResponseBodyBatchErrorJSON contains the JSON metadata for the struct
// [PayStatementResponseBodyBatchError]
type payStatementResponseBodyBatchErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	FinchCode   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementResponseBodyBatchError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payStatementResponseBodyBatchErrorJSON) RawJSON() string {
	return r.raw
}

func (r PayStatementResponseBodyBatchError) implementsPayStatementResponseBody() {}

type HRISPayStatementGetManyParams struct {
	// The array of batch requests.
	Requests param.Field[[]HRISPayStatementGetManyParamsRequest] `json:"requests,required"`
}

func (r HRISPayStatementGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISPayStatementGetManyParamsRequest struct {
	// A stable Finch `id` (UUID v4) for a payment.
	PaymentID param.Field[string] `json:"payment_id,required" format:"uuid"`
	// Number of pay statements to return (defaults to all).
	Limit param.Field[int64] `json:"limit"`
	// Index to start from.
	Offset param.Field[int64] `json:"offset"`
}

func (r HRISPayStatementGetManyParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
