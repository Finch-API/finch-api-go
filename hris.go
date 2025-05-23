// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISService contains methods and other services that help with interacting with
// the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISService] method instead.
type HRISService struct {
	Options       []option.RequestOption
	Company       *HRISCompanyService
	Directory     *HRISDirectoryService
	Individuals   *HRISIndividualService
	Employments   *HRISEmploymentService
	Payments      *HRISPaymentService
	PayStatements *HRISPayStatementService
	Documents     *HRISDocumentService
	Benefits      *HRISBenefitService
}

// NewHRISService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHRISService(opts ...option.RequestOption) (r *HRISService) {
	r = &HRISService{}
	r.Options = opts
	r.Company = NewHRISCompanyService(opts...)
	r.Directory = NewHRISDirectoryService(opts...)
	r.Individuals = NewHRISIndividualService(opts...)
	r.Employments = NewHRISEmploymentService(opts...)
	r.Payments = NewHRISPaymentService(opts...)
	r.PayStatements = NewHRISPayStatementService(opts...)
	r.Documents = NewHRISDocumentService(opts...)
	r.Benefits = NewHRISBenefitService(opts...)
	return
}

// The employee's income as reported by the provider. This may not always be
// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
// depending on what information the provider returns.
type Income struct {
	// The income amount in cents.
	Amount int64 `json:"amount,required,nullable"`
	// The currency code.
	Currency string `json:"currency,required,nullable"`
	// The date the income amount went into effect.
	EffectiveDate string `json:"effective_date,required,nullable"`
	// The income unit of payment. Options: `yearly`, `quarterly`, `monthly`,
	// `semi_monthly`, `bi_weekly`, `weekly`, `daily`, `hourly`, and `fixed`.
	Unit IncomeUnit `json:"unit,required,nullable"`
	JSON incomeJSON `json:"-"`
}

// incomeJSON contains the JSON metadata for the struct [Income]
type incomeJSON struct {
	Amount        apijson.Field
	Currency      apijson.Field
	EffectiveDate apijson.Field
	Unit          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Income) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomeJSON) RawJSON() string {
	return r.raw
}

// The income unit of payment. Options: `yearly`, `quarterly`, `monthly`,
// `semi_monthly`, `bi_weekly`, `weekly`, `daily`, `hourly`, and `fixed`.
type IncomeUnit string

const (
	IncomeUnitYearly      IncomeUnit = "yearly"
	IncomeUnitQuarterly   IncomeUnit = "quarterly"
	IncomeUnitMonthly     IncomeUnit = "monthly"
	IncomeUnitSemiMonthly IncomeUnit = "semi_monthly"
	IncomeUnitBiWeekly    IncomeUnit = "bi_weekly"
	IncomeUnitWeekly      IncomeUnit = "weekly"
	IncomeUnitDaily       IncomeUnit = "daily"
	IncomeUnitHourly      IncomeUnit = "hourly"
	IncomeUnitFixed       IncomeUnit = "fixed"
)

func (r IncomeUnit) IsKnown() bool {
	switch r {
	case IncomeUnitYearly, IncomeUnitQuarterly, IncomeUnitMonthly, IncomeUnitSemiMonthly, IncomeUnitBiWeekly, IncomeUnitWeekly, IncomeUnitDaily, IncomeUnitHourly, IncomeUnitFixed:
		return true
	}
	return false
}

// The employee's income as reported by the provider. This may not always be
// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
// depending on what information the provider returns.
type IncomeParam struct {
	// The income amount in cents.
	Amount param.Field[int64] `json:"amount,required"`
	// The currency code.
	Currency param.Field[string] `json:"currency,required"`
	// The date the income amount went into effect.
	EffectiveDate param.Field[string] `json:"effective_date,required"`
	// The income unit of payment. Options: `yearly`, `quarterly`, `monthly`,
	// `semi_monthly`, `bi_weekly`, `weekly`, `daily`, `hourly`, and `fixed`.
	Unit param.Field[IncomeUnit] `json:"unit,required"`
}

func (r IncomeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Location struct {
	// City, district, suburb, town, or village.
	City string `json:"city,required,nullable"`
	// The 2-letter ISO 3166 country code.
	Country string `json:"country,required,nullable"`
	// Street address or PO box.
	Line1 string `json:"line1,required,nullable"`
	// Apartment, suite, unit, or building.
	Line2 string `json:"line2,required,nullable"`
	// The postal code or zip code.
	PostalCode string `json:"postal_code,required,nullable"`
	// The state code.
	State    string       `json:"state,required,nullable"`
	Name     string       `json:"name,nullable"`
	SourceID string       `json:"source_id,nullable"`
	JSON     locationJSON `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Name        apijson.Field
	SourceID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationJSON) RawJSON() string {
	return r.raw
}

type LocationParam struct {
	// City, district, suburb, town, or village.
	City param.Field[string] `json:"city,required"`
	// The 2-letter ISO 3166 country code.
	Country param.Field[string] `json:"country,required"`
	// Street address or PO box.
	Line1 param.Field[string] `json:"line1,required"`
	// Apartment, suite, unit, or building.
	Line2 param.Field[string] `json:"line2,required"`
	// The postal code or zip code.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// The state code.
	State    param.Field[string] `json:"state,required"`
	Name     param.Field[string] `json:"name"`
	SourceID param.Field[string] `json:"source_id"`
}

func (r LocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Money struct {
	// Amount for money object (in cents)
	Amount   int64     `json:"amount,required,nullable"`
	Currency string    `json:"currency,required"`
	JSON     moneyJSON `json:"-"`
}

// moneyJSON contains the JSON metadata for the struct [Money]
type moneyJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Money) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moneyJSON) RawJSON() string {
	return r.raw
}

type MoneyParam struct {
	// Amount for money object (in cents)
	Amount   param.Field[int64]  `json:"amount,required"`
	Currency param.Field[string] `json:"currency,required"`
}

func (r MoneyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
