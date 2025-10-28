// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxCompanyService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxCompanyService] method instead.
type SandboxCompanyService struct {
	Options []option.RequestOption
}

// NewSandboxCompanyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxCompanyService(opts ...option.RequestOption) (r *SandboxCompanyService) {
	r = &SandboxCompanyService{}
	r.Options = opts
	return
}

// Update a sandbox company's data
func (r *SandboxCompanyService) Update(ctx context.Context, body SandboxCompanyUpdateParams, opts ...option.RequestOption) (res *SandboxCompanyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandbox/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SandboxCompanyUpdateResponse struct {
	// An array of bank account objects associated with the payroll/HRIS system.
	Accounts []SandboxCompanyUpdateResponseAccount `json:"accounts,required,nullable"`
	// The array of company departments.
	Departments []SandboxCompanyUpdateResponseDepartment `json:"departments,required,nullable"`
	// The employer identification number.
	Ein string `json:"ein,required,nullable"`
	// The entity type object.
	Entity SandboxCompanyUpdateResponseEntity `json:"entity,required,nullable"`
	// The legal name of the company.
	LegalName string     `json:"legal_name,required,nullable"`
	Locations []Location `json:"locations,required,nullable"`
	// The email of the main administrator on the account.
	PrimaryEmail string `json:"primary_email,required,nullable" format:"email"`
	// The phone number of the main administrator on the account. Format: E.164, with
	// extension where applicable, e.g. `+NNNNNNNNNNN xExtension`
	PrimaryPhoneNumber string                           `json:"primary_phone_number,required,nullable"`
	JSON               sandboxCompanyUpdateResponseJSON `json:"-"`
}

// sandboxCompanyUpdateResponseJSON contains the JSON metadata for the struct
// [SandboxCompanyUpdateResponse]
type sandboxCompanyUpdateResponseJSON struct {
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

func (r *SandboxCompanyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxCompanyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxCompanyUpdateResponseAccount struct {
	// The name of the bank associated in the payroll/HRIS system.
	AccountName string `json:"account_name,nullable"`
	// 10-12 digit number to specify the bank account
	AccountNumber string `json:"account_number,nullable"`
	// The type of bank account.
	AccountType SandboxCompanyUpdateResponseAccountsAccountType `json:"account_type,nullable"`
	// Name of the banking institution.
	InstitutionName string `json:"institution_name,nullable"`
	// A nine-digit code that's based on the U.S. Bank location where your account was
	// opened.
	RoutingNumber string                                  `json:"routing_number,nullable"`
	JSON          sandboxCompanyUpdateResponseAccountJSON `json:"-"`
}

// sandboxCompanyUpdateResponseAccountJSON contains the JSON metadata for the
// struct [SandboxCompanyUpdateResponseAccount]
type sandboxCompanyUpdateResponseAccountJSON struct {
	AccountName     apijson.Field
	AccountNumber   apijson.Field
	AccountType     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SandboxCompanyUpdateResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxCompanyUpdateResponseAccountJSON) RawJSON() string {
	return r.raw
}

// The type of bank account.
type SandboxCompanyUpdateResponseAccountsAccountType string

const (
	SandboxCompanyUpdateResponseAccountsAccountTypeChecking SandboxCompanyUpdateResponseAccountsAccountType = "checking"
	SandboxCompanyUpdateResponseAccountsAccountTypeSavings  SandboxCompanyUpdateResponseAccountsAccountType = "savings"
)

func (r SandboxCompanyUpdateResponseAccountsAccountType) IsKnown() bool {
	switch r {
	case SandboxCompanyUpdateResponseAccountsAccountTypeChecking, SandboxCompanyUpdateResponseAccountsAccountTypeSavings:
		return true
	}
	return false
}

type SandboxCompanyUpdateResponseDepartment struct {
	// The department name.
	Name string `json:"name,nullable"`
	// The parent department, if present.
	Parent SandboxCompanyUpdateResponseDepartmentsParent `json:"parent,nullable"`
	JSON   sandboxCompanyUpdateResponseDepartmentJSON    `json:"-"`
}

// sandboxCompanyUpdateResponseDepartmentJSON contains the JSON metadata for the
// struct [SandboxCompanyUpdateResponseDepartment]
type sandboxCompanyUpdateResponseDepartmentJSON struct {
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxCompanyUpdateResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxCompanyUpdateResponseDepartmentJSON) RawJSON() string {
	return r.raw
}

// The parent department, if present.
type SandboxCompanyUpdateResponseDepartmentsParent struct {
	// The parent department's name.
	Name string                                            `json:"name,nullable"`
	JSON sandboxCompanyUpdateResponseDepartmentsParentJSON `json:"-"`
}

// sandboxCompanyUpdateResponseDepartmentsParentJSON contains the JSON metadata for
// the struct [SandboxCompanyUpdateResponseDepartmentsParent]
type sandboxCompanyUpdateResponseDepartmentsParentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxCompanyUpdateResponseDepartmentsParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxCompanyUpdateResponseDepartmentsParentJSON) RawJSON() string {
	return r.raw
}

// The entity type object.
type SandboxCompanyUpdateResponseEntity struct {
	// The tax payer subtype of the company.
	Subtype SandboxCompanyUpdateResponseEntitySubtype `json:"subtype,nullable"`
	// The tax payer type of the company.
	Type SandboxCompanyUpdateResponseEntityType `json:"type,nullable"`
	JSON sandboxCompanyUpdateResponseEntityJSON `json:"-"`
}

// sandboxCompanyUpdateResponseEntityJSON contains the JSON metadata for the struct
// [SandboxCompanyUpdateResponseEntity]
type sandboxCompanyUpdateResponseEntityJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxCompanyUpdateResponseEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxCompanyUpdateResponseEntityJSON) RawJSON() string {
	return r.raw
}

// The tax payer subtype of the company.
type SandboxCompanyUpdateResponseEntitySubtype string

const (
	SandboxCompanyUpdateResponseEntitySubtypeSCorporation SandboxCompanyUpdateResponseEntitySubtype = "s_corporation"
	SandboxCompanyUpdateResponseEntitySubtypeCCorporation SandboxCompanyUpdateResponseEntitySubtype = "c_corporation"
	SandboxCompanyUpdateResponseEntitySubtypeBCorporation SandboxCompanyUpdateResponseEntitySubtype = "b_corporation"
)

func (r SandboxCompanyUpdateResponseEntitySubtype) IsKnown() bool {
	switch r {
	case SandboxCompanyUpdateResponseEntitySubtypeSCorporation, SandboxCompanyUpdateResponseEntitySubtypeCCorporation, SandboxCompanyUpdateResponseEntitySubtypeBCorporation:
		return true
	}
	return false
}

// The tax payer type of the company.
type SandboxCompanyUpdateResponseEntityType string

const (
	SandboxCompanyUpdateResponseEntityTypeLlc            SandboxCompanyUpdateResponseEntityType = "llc"
	SandboxCompanyUpdateResponseEntityTypeLp             SandboxCompanyUpdateResponseEntityType = "lp"
	SandboxCompanyUpdateResponseEntityTypeCorporation    SandboxCompanyUpdateResponseEntityType = "corporation"
	SandboxCompanyUpdateResponseEntityTypeSoleProprietor SandboxCompanyUpdateResponseEntityType = "sole_proprietor"
	SandboxCompanyUpdateResponseEntityTypeNonProfit      SandboxCompanyUpdateResponseEntityType = "non_profit"
	SandboxCompanyUpdateResponseEntityTypePartnership    SandboxCompanyUpdateResponseEntityType = "partnership"
	SandboxCompanyUpdateResponseEntityTypeCooperative    SandboxCompanyUpdateResponseEntityType = "cooperative"
)

func (r SandboxCompanyUpdateResponseEntityType) IsKnown() bool {
	switch r {
	case SandboxCompanyUpdateResponseEntityTypeLlc, SandboxCompanyUpdateResponseEntityTypeLp, SandboxCompanyUpdateResponseEntityTypeCorporation, SandboxCompanyUpdateResponseEntityTypeSoleProprietor, SandboxCompanyUpdateResponseEntityTypeNonProfit, SandboxCompanyUpdateResponseEntityTypePartnership, SandboxCompanyUpdateResponseEntityTypeCooperative:
		return true
	}
	return false
}

type SandboxCompanyUpdateParams struct {
	// An array of bank account objects associated with the payroll/HRIS system.
	Accounts param.Field[[]SandboxCompanyUpdateParamsAccount] `json:"accounts,required"`
	// The array of company departments.
	Departments param.Field[[]SandboxCompanyUpdateParamsDepartment] `json:"departments,required"`
	// The employer identification number.
	Ein param.Field[string] `json:"ein,required"`
	// The entity type object.
	Entity param.Field[SandboxCompanyUpdateParamsEntity] `json:"entity,required"`
	// The legal name of the company.
	LegalName param.Field[string]          `json:"legal_name,required"`
	Locations param.Field[[]LocationParam] `json:"locations,required"`
	// The email of the main administrator on the account.
	PrimaryEmail param.Field[string] `json:"primary_email,required" format:"email"`
	// The phone number of the main administrator on the account. Format: E.164, with
	// extension where applicable, e.g. `+NNNNNNNNNNN xExtension`
	PrimaryPhoneNumber param.Field[string] `json:"primary_phone_number,required"`
}

func (r SandboxCompanyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxCompanyUpdateParamsAccount struct {
	// The name of the bank associated in the payroll/HRIS system.
	AccountName param.Field[string] `json:"account_name"`
	// 10-12 digit number to specify the bank account
	AccountNumber param.Field[string] `json:"account_number"`
	// The type of bank account.
	AccountType param.Field[SandboxCompanyUpdateParamsAccountsAccountType] `json:"account_type"`
	// Name of the banking institution.
	InstitutionName param.Field[string] `json:"institution_name"`
	// A nine-digit code that's based on the U.S. Bank location where your account was
	// opened.
	RoutingNumber param.Field[string] `json:"routing_number"`
}

func (r SandboxCompanyUpdateParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of bank account.
type SandboxCompanyUpdateParamsAccountsAccountType string

const (
	SandboxCompanyUpdateParamsAccountsAccountTypeChecking SandboxCompanyUpdateParamsAccountsAccountType = "checking"
	SandboxCompanyUpdateParamsAccountsAccountTypeSavings  SandboxCompanyUpdateParamsAccountsAccountType = "savings"
)

func (r SandboxCompanyUpdateParamsAccountsAccountType) IsKnown() bool {
	switch r {
	case SandboxCompanyUpdateParamsAccountsAccountTypeChecking, SandboxCompanyUpdateParamsAccountsAccountTypeSavings:
		return true
	}
	return false
}

type SandboxCompanyUpdateParamsDepartment struct {
	// The department name.
	Name param.Field[string] `json:"name"`
	// The parent department, if present.
	Parent param.Field[SandboxCompanyUpdateParamsDepartmentsParent] `json:"parent"`
}

func (r SandboxCompanyUpdateParamsDepartment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The parent department, if present.
type SandboxCompanyUpdateParamsDepartmentsParent struct {
	// The parent department's name.
	Name param.Field[string] `json:"name"`
}

func (r SandboxCompanyUpdateParamsDepartmentsParent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity type object.
type SandboxCompanyUpdateParamsEntity struct {
	// The tax payer subtype of the company.
	Subtype param.Field[SandboxCompanyUpdateParamsEntitySubtype] `json:"subtype"`
	// The tax payer type of the company.
	Type param.Field[SandboxCompanyUpdateParamsEntityType] `json:"type"`
}

func (r SandboxCompanyUpdateParamsEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tax payer subtype of the company.
type SandboxCompanyUpdateParamsEntitySubtype string

const (
	SandboxCompanyUpdateParamsEntitySubtypeSCorporation SandboxCompanyUpdateParamsEntitySubtype = "s_corporation"
	SandboxCompanyUpdateParamsEntitySubtypeCCorporation SandboxCompanyUpdateParamsEntitySubtype = "c_corporation"
	SandboxCompanyUpdateParamsEntitySubtypeBCorporation SandboxCompanyUpdateParamsEntitySubtype = "b_corporation"
)

func (r SandboxCompanyUpdateParamsEntitySubtype) IsKnown() bool {
	switch r {
	case SandboxCompanyUpdateParamsEntitySubtypeSCorporation, SandboxCompanyUpdateParamsEntitySubtypeCCorporation, SandboxCompanyUpdateParamsEntitySubtypeBCorporation:
		return true
	}
	return false
}

// The tax payer type of the company.
type SandboxCompanyUpdateParamsEntityType string

const (
	SandboxCompanyUpdateParamsEntityTypeLlc            SandboxCompanyUpdateParamsEntityType = "llc"
	SandboxCompanyUpdateParamsEntityTypeLp             SandboxCompanyUpdateParamsEntityType = "lp"
	SandboxCompanyUpdateParamsEntityTypeCorporation    SandboxCompanyUpdateParamsEntityType = "corporation"
	SandboxCompanyUpdateParamsEntityTypeSoleProprietor SandboxCompanyUpdateParamsEntityType = "sole_proprietor"
	SandboxCompanyUpdateParamsEntityTypeNonProfit      SandboxCompanyUpdateParamsEntityType = "non_profit"
	SandboxCompanyUpdateParamsEntityTypePartnership    SandboxCompanyUpdateParamsEntityType = "partnership"
	SandboxCompanyUpdateParamsEntityTypeCooperative    SandboxCompanyUpdateParamsEntityType = "cooperative"
)

func (r SandboxCompanyUpdateParamsEntityType) IsKnown() bool {
	switch r {
	case SandboxCompanyUpdateParamsEntityTypeLlc, SandboxCompanyUpdateParamsEntityTypeLp, SandboxCompanyUpdateParamsEntityTypeCorporation, SandboxCompanyUpdateParamsEntityTypeSoleProprietor, SandboxCompanyUpdateParamsEntityTypeNonProfit, SandboxCompanyUpdateParamsEntityTypePartnership, SandboxCompanyUpdateParamsEntityTypeCooperative:
		return true
	}
	return false
}
