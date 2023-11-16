// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISService contains methods and other services that help with interacting with
// the Finch API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewHRISService] method instead.
type HRISService struct {
	Options       []option.RequestOption
	Company       *HRISCompanyService
	Directory     *HRISDirectoryService
	Individuals   *HRISIndividualService
	Employments   *HRISEmploymentService
	Payments      *HRISPaymentService
	PayStatements *HRISPayStatementService
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
	r.Benefits = NewHRISBenefitService(opts...)
	return
}

// The employee's income as reported by the provider. This may not always be
// annualized income, but may be in units of bi-weekly, semi-monthly, daily, etc,
// depending on what information the provider returns.
type Income struct {
	// The income amount in cents.
	Amount int64 `json:"amount,nullable"`
	// The currency code.
	Currency string `json:"currency,nullable"`
	// The date the income amount went into effect.
	EffectiveDate string `json:"effective_date,nullable"`
	// The income unit of payment. Options: `yearly`, `quarterly`, `monthly`,
	// `semi_monthly`, `bi_weekly`, `weekly`, `daily`, `hourly`, and `fixed`.
	Unit IncomeUnit `json:"unit,nullable"`
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

type Location struct {
	// City, district, suburb, town, or village.
	City string `json:"city,nullable"`
	// The 2-letter ISO 3166 country code.
	Country string `json:"country,nullable"`
	// Street address or PO box.
	Line1 string `json:"line1,nullable"`
	// Apartment, suite, unit, or building.
	Line2 string `json:"line2,nullable"`
	Name  string `json:"name,nullable"`
	// The postal code or zip code.
	PostalCode string `json:"postal_code,nullable"`
	SourceID   string `json:"source_id,nullable"`
	// The state code.
	State string       `json:"state,nullable"`
	JSON  locationJSON `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Name        apijson.Field
	PostalCode  apijson.Field
	SourceID    apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Money struct {
	// Amount for money object (in cents)
	Amount   int64     `json:"amount,nullable"`
	Currency string    `json:"currency"`
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

type Paging struct {
	// The total number of elements for the entire query (not just the given page)
	Count int64 `json:"count"`
	// The current start index of the returned list of elements
	Offset int64      `json:"offset"`
	JSON   pagingJSON `json:"-"`
}

// pagingJSON contains the JSON metadata for the struct [Paging]
type pagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Paging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
