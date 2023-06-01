package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISCompanyService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHRISCompanyService] method instead.
type HRISCompanyService struct {
	Options []option.RequestOption
}

// NewHRISCompanyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISCompanyService(opts ...option.RequestOption) (r *HRISCompanyService) {
	r = &HRISCompanyService{}
	r.Options = opts
	return
}

// Read basic company data
func (r *HRISCompanyService) Get(ctx context.Context, opts ...option.RequestOption) (res *Company, err error) {
	opts = append(r.Options[:], opts...)
	path := "employer/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Company struct {
	// A stable Finch `id` (UUID v4) for the company.
	ID string `json:"id,required"`
	// The legal name of the company.
	LegalName string `json:"legal_name,required,nullable"`
	// The entity type object.
	Entity CompanyEntity `json:"entity,required,nullable"`
	// The email of the main administrator on the account.
	PrimaryEmail string `json:"primary_email,required,nullable"`
	// The phone number of the main administrator on the account. Format: `XXXXXXXXXX`
	PrimaryPhoneNumber string `json:"primary_phone_number,required,nullable"`
	// The array of company departments.
	Departments []CompanyDepartments `json:"departments,required,nullable"`
	// The employer identification number.
	Ein       string     `json:"ein,required,nullable"`
	Locations []Location `json:"locations,required,nullable"`
	// An array of bank account objects associated with the payroll/HRIS system.
	Accounts []CompanyAccounts `json:"accounts,required,nullable"`
	JSON     companyJSON
}

// companyJSON contains the JSON metadata for the struct [Company]
type companyJSON struct {
	ID                 apijson.Field
	LegalName          apijson.Field
	Entity             apijson.Field
	PrimaryEmail       apijson.Field
	PrimaryPhoneNumber apijson.Field
	Departments        apijson.Field
	Ein                apijson.Field
	Locations          apijson.Field
	Accounts           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Company) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The entity type object.
type CompanyEntity struct {
	// The tax payer type of the company.
	Type CompanyEntityType `json:"type,nullable"`
	// The tax payer subtype of the company.
	Subtype CompanyEntitySubtype `json:"subtype,nullable"`
	JSON    companyEntityJSON
}

// companyEntityJSON contains the JSON metadata for the struct [CompanyEntity]
type companyEntityJSON struct {
	Type        apijson.Field
	Subtype     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyEntityType string

const (
	CompanyEntityTypeLlc            CompanyEntityType = "llc"
	CompanyEntityTypeCorporation    CompanyEntityType = "corporation"
	CompanyEntityTypeSoleProprietor CompanyEntityType = "sole_proprietor"
	CompanyEntityTypeNonProfit      CompanyEntityType = "non_profit"
	CompanyEntityTypePartnership    CompanyEntityType = "partnership"
	CompanyEntityTypeCooperative    CompanyEntityType = "cooperative"
)

type CompanyEntitySubtype string

const (
	CompanyEntitySubtypeSCorporation CompanyEntitySubtype = "s_corporation"
	CompanyEntitySubtypeCCorporation CompanyEntitySubtype = "c_corporation"
	CompanyEntitySubtypeBCorporation CompanyEntitySubtype = "b_corporation"
)

type CompanyDepartments struct {
	// The department name.
	Name string `json:"name,nullable"`
	// The parent department, if present.
	Parent CompanyDepartmentsParent `json:"parent,nullable"`
	JSON   companyDepartmentsJSON
}

// companyDepartmentsJSON contains the JSON metadata for the struct
// [CompanyDepartments]
type companyDepartmentsJSON struct {
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyDepartments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parent department, if present.
type CompanyDepartmentsParent struct {
	// The parent department's name.
	Name string `json:"name,nullable"`
	JSON companyDepartmentsParentJSON
}

// companyDepartmentsParentJSON contains the JSON metadata for the struct
// [CompanyDepartmentsParent]
type companyDepartmentsParentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyDepartmentsParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyAccounts struct {
	// A nine-digit code that's based on the U.S. Bank location where your account was
	// opened.
	RoutingNumber string `json:"routing_number,nullable"`
	// The name of the bank associated in the payroll/HRIS system.
	AccountName string `json:"account_name,nullable"`
	// Name of the banking institution.
	InstitutionName string `json:"institution_name,nullable"`
	// The type of bank account.
	AccountType CompanyAccountsAccountType `json:"account_type,nullable"`
	// 10-12 digit number to specify the bank account
	AccountNumber string `json:"account_number,nullable"`
	JSON          companyAccountsJSON
}

// companyAccountsJSON contains the JSON metadata for the struct [CompanyAccounts]
type companyAccountsJSON struct {
	RoutingNumber   apijson.Field
	AccountName     apijson.Field
	InstitutionName apijson.Field
	AccountType     apijson.Field
	AccountNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CompanyAccounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyAccountsAccountType string

const (
	CompanyAccountsAccountTypeChecking CompanyAccountsAccountType = "checking"
	CompanyAccountsAccountTypeSavings  CompanyAccountsAccountType = "savings"
)
