// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
	// An array of bank account objects associated with the payroll/HRIS system.
	Accounts []CompanyAccount `json:"accounts,required,nullable"`
	// The array of company departments.
	Departments []CompanyDepartment `json:"departments,required,nullable"`
	// The employer identification number.
	Ein string `json:"ein,required,nullable"`
	// The entity type object.
	Entity CompanyEntity `json:"entity,required,nullable"`
	// The legal name of the company.
	LegalName string     `json:"legal_name,required,nullable"`
	Locations []Location `json:"locations,required,nullable"`
	// The email of the main administrator on the account.
	PrimaryEmail string `json:"primary_email,required,nullable"`
	// The phone number of the main administrator on the account. Format: `XXXXXXXXXX`
	PrimaryPhoneNumber string      `json:"primary_phone_number,required,nullable"`
	JSON               companyJSON `json:"-"`
}

// companyJSON contains the JSON metadata for the struct [Company]
type companyJSON struct {
	ID                 apijson.Field
	Accounts           apijson.Field
	Departments        apijson.Field
	Ein                apijson.Field
	Entity             apijson.Field
	LegalName          apijson.Field
	Locations          apijson.Field
	PrimaryEmail       apijson.Field
	PrimaryPhoneNumber apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Company) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyJSON) RawJSON() string {
	return r.raw
}

type CompanyAccount struct {
	// The name of the bank associated in the payroll/HRIS system.
	AccountName string `json:"account_name,nullable"`
	// 10-12 digit number to specify the bank account
	AccountNumber string `json:"account_number,nullable"`
	// The type of bank account.
	AccountType CompanyAccountsAccountType `json:"account_type,nullable"`
	// Name of the banking institution.
	InstitutionName string `json:"institution_name,nullable"`
	// A nine-digit code that's based on the U.S. Bank location where your account was
	// opened.
	RoutingNumber string             `json:"routing_number,nullable"`
	JSON          companyAccountJSON `json:"-"`
}

// companyAccountJSON contains the JSON metadata for the struct [CompanyAccount]
type companyAccountJSON struct {
	AccountName     apijson.Field
	AccountNumber   apijson.Field
	AccountType     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CompanyAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyAccountJSON) RawJSON() string {
	return r.raw
}

// The type of bank account.
type CompanyAccountsAccountType string

const (
	CompanyAccountsAccountTypeChecking CompanyAccountsAccountType = "checking"
	CompanyAccountsAccountTypeSavings  CompanyAccountsAccountType = "savings"
)

func (r CompanyAccountsAccountType) IsKnown() bool {
	switch r {
	case CompanyAccountsAccountTypeChecking, CompanyAccountsAccountTypeSavings:
		return true
	}
	return false
}

type CompanyDepartment struct {
	// The department name.
	Name string `json:"name,nullable"`
	// The parent department, if present.
	Parent CompanyDepartmentsParent `json:"parent,nullable"`
	JSON   companyDepartmentJSON    `json:"-"`
}

// companyDepartmentJSON contains the JSON metadata for the struct
// [CompanyDepartment]
type companyDepartmentJSON struct {
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyDepartmentJSON) RawJSON() string {
	return r.raw
}

// The parent department, if present.
type CompanyDepartmentsParent struct {
	// The parent department's name.
	Name string                       `json:"name,nullable"`
	JSON companyDepartmentsParentJSON `json:"-"`
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

func (r companyDepartmentsParentJSON) RawJSON() string {
	return r.raw
}

// The entity type object.
type CompanyEntity struct {
	// The tax payer subtype of the company.
	Subtype CompanyEntitySubtype `json:"subtype,nullable"`
	// The tax payer type of the company.
	Type CompanyEntityType `json:"type,nullable"`
	JSON companyEntityJSON `json:"-"`
}

// companyEntityJSON contains the JSON metadata for the struct [CompanyEntity]
type companyEntityJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyEntityJSON) RawJSON() string {
	return r.raw
}

// The tax payer subtype of the company.
type CompanyEntitySubtype string

const (
	CompanyEntitySubtypeSCorporation CompanyEntitySubtype = "s_corporation"
	CompanyEntitySubtypeCCorporation CompanyEntitySubtype = "c_corporation"
	CompanyEntitySubtypeBCorporation CompanyEntitySubtype = "b_corporation"
)

func (r CompanyEntitySubtype) IsKnown() bool {
	switch r {
	case CompanyEntitySubtypeSCorporation, CompanyEntitySubtypeCCorporation, CompanyEntitySubtypeBCorporation:
		return true
	}
	return false
}

// The tax payer type of the company.
type CompanyEntityType string

const (
	CompanyEntityTypeLlc            CompanyEntityType = "llc"
	CompanyEntityTypeLp             CompanyEntityType = "lp"
	CompanyEntityTypeCorporation    CompanyEntityType = "corporation"
	CompanyEntityTypeSoleProprietor CompanyEntityType = "sole_proprietor"
	CompanyEntityTypeNonProfit      CompanyEntityType = "non_profit"
	CompanyEntityTypePartnership    CompanyEntityType = "partnership"
	CompanyEntityTypeCooperative    CompanyEntityType = "cooperative"
)

func (r CompanyEntityType) IsKnown() bool {
	switch r {
	case CompanyEntityTypeLlc, CompanyEntityTypeLp, CompanyEntityTypeCorporation, CompanyEntityTypeSoleProprietor, CompanyEntityTypeNonProfit, CompanyEntityTypePartnership, CompanyEntityTypeCooperative:
		return true
	}
	return false
}
