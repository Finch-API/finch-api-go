// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"slices"

	"github.com/Finch-API/finch-api-go/v2/internal/apijson"
	"github.com/Finch-API/finch-api-go/v2/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/v2/option"
	"github.com/Finch-API/finch-api-go/v2/packages/pagination"
)

// ProviderService contains methods and other services that help with interacting
// with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r *ProviderService) {
	r = &ProviderService{}
	r.Options = opts
	return
}

// Return details on all available payroll and HR systems.
func (r *ProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[ProviderListResponse], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "providers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Return details on all available payroll and HR systems.
func (r *ProviderService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProviderListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type ProviderListResponse struct {
	// The id of the payroll provider used in Connect.
	ID string `json:"id" api:"required"`
	// The display name of the payroll provider.
	DisplayName string `json:"display_name" api:"required"`
	// The list of Finch products supported on this payroll provider.
	Products []string `json:"products" api:"required"`
	// The authentication methods supported by the provider.
	AuthenticationMethods []ProviderListResponseAuthenticationMethod `json:"authentication_methods"`
	// `true` if the integration is in a beta state, `false` otherwise
	Beta bool `json:"beta"`
	// The url to the official icon of the payroll provider.
	Icon string `json:"icon"`
	// The url to the official logo of the payroll provider.
	Logo string `json:"logo"`
	// [DEPRECATED] Whether the Finch integration with this provider uses the Assisted
	// Connect Flow by default. This field is now deprecated. Please check for a `type`
	// of `assisted` in the `authentication_methods` field instead.
	//
	// Deprecated: deprecated
	Manual bool `json:"manual"`
	// whether MFA is required for the provider.
	MfaRequired bool `json:"mfa_required"`
	// The hex code for the primary color of the payroll provider.
	PrimaryColor string                   `json:"primary_color"`
	JSON         providerListResponseJSON `json:"-"`
}

// providerListResponseJSON contains the JSON metadata for the struct
// [ProviderListResponse]
type providerListResponseJSON struct {
	ID                    apijson.Field
	DisplayName           apijson.Field
	Products              apijson.Field
	AuthenticationMethods apijson.Field
	Beta                  apijson.Field
	Icon                  apijson.Field
	Logo                  apijson.Field
	Manual                apijson.Field
	MfaRequired           apijson.Field
	PrimaryColor          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethod struct {
	// The type of authentication method
	Type ProviderListResponseAuthenticationMethodsType `json:"type" api:"required"`
	// The supported benefit types and their configurations
	BenefitsSupport map[string]interface{} `json:"benefits_support"`
	// The supported data fields returned by our HR, payroll, and benefits endpoints
	SupportedFields ProviderListResponseAuthenticationMethodsSupportedFields `json:"supported_fields" api:"nullable"`
	JSON            providerListResponseAuthenticationMethodJSON             `json:"-"`
}

// providerListResponseAuthenticationMethodJSON contains the JSON metadata for the
// struct [ProviderListResponseAuthenticationMethod]
type providerListResponseAuthenticationMethodJSON struct {
	Type            apijson.Field
	BenefitsSupport apijson.Field
	SupportedFields apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodJSON) RawJSON() string {
	return r.raw
}

// The type of authentication method
type ProviderListResponseAuthenticationMethodsType string

const (
	ProviderListResponseAuthenticationMethodsTypeAssisted      ProviderListResponseAuthenticationMethodsType = "assisted"
	ProviderListResponseAuthenticationMethodsTypeCredential    ProviderListResponseAuthenticationMethodsType = "credential"
	ProviderListResponseAuthenticationMethodsTypeAPIToken      ProviderListResponseAuthenticationMethodsType = "api_token"
	ProviderListResponseAuthenticationMethodsTypeAPICredential ProviderListResponseAuthenticationMethodsType = "api_credential"
	ProviderListResponseAuthenticationMethodsTypeOAuth         ProviderListResponseAuthenticationMethodsType = "oauth"
	ProviderListResponseAuthenticationMethodsTypeAPI           ProviderListResponseAuthenticationMethodsType = "api"
)

func (r ProviderListResponseAuthenticationMethodsType) IsKnown() bool {
	switch r {
	case ProviderListResponseAuthenticationMethodsTypeAssisted, ProviderListResponseAuthenticationMethodsTypeCredential, ProviderListResponseAuthenticationMethodsTypeAPIToken, ProviderListResponseAuthenticationMethodsTypeAPICredential, ProviderListResponseAuthenticationMethodsTypeOAuth, ProviderListResponseAuthenticationMethodsTypeAPI:
		return true
	}
	return false
}

// The supported data fields returned by our HR, payroll, and benefits endpoints
type ProviderListResponseAuthenticationMethodsSupportedFields struct {
	Company         ProviderListResponseAuthenticationMethodsSupportedFieldsCompany         `json:"company"`
	Directory       ProviderListResponseAuthenticationMethodsSupportedFieldsDirectory       `json:"directory"`
	Employment      ProviderListResponseAuthenticationMethodsSupportedFieldsEmployment      `json:"employment"`
	Individual      ProviderListResponseAuthenticationMethodsSupportedFieldsIndividual      `json:"individual"`
	PayGroup        ProviderListResponseAuthenticationMethodsSupportedFieldsPayGroup        `json:"pay_group"`
	PayStatement    ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatement    `json:"pay_statement"`
	Payment         ProviderListResponseAuthenticationMethodsSupportedFieldsPayment         `json:"payment"`
	PlanDependents  ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependents  `json:"plan_dependents"`
	PlanEnrollments ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollments `json:"plan_enrollments"`
	Plans           ProviderListResponseAuthenticationMethodsSupportedFieldsPlans           `json:"plans"`
	JSON            providerListResponseAuthenticationMethodsSupportedFieldsJSON            `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsJSON contains the JSON
// metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFields]
type providerListResponseAuthenticationMethodsSupportedFieldsJSON struct {
	Company         apijson.Field
	Directory       apijson.Field
	Employment      apijson.Field
	Individual      apijson.Field
	PayGroup        apijson.Field
	PayStatement    apijson.Field
	Payment         apijson.Field
	PlanDependents  apijson.Field
	PlanEnrollments apijson.Field
	Plans           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsCompany struct {
	ID                 bool                                                                       `json:"id"`
	Accounts           ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyAccounts    `json:"accounts"`
	Departments        ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartments `json:"departments"`
	Ein                bool                                                                       `json:"ein"`
	Entity             ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyEntity      `json:"entity"`
	LegalName          bool                                                                       `json:"legal_name"`
	Locations          ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyLocations   `json:"locations"`
	PrimaryEmail       bool                                                                       `json:"primary_email"`
	PrimaryPhoneNumber bool                                                                       `json:"primary_phone_number"`
	JSON               providerListResponseAuthenticationMethodsSupportedFieldsCompanyJSON        `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsCompanyJSON contains the
// JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsCompany]
type providerListResponseAuthenticationMethodsSupportedFieldsCompanyJSON struct {
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

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsCompany) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsCompanyJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyAccounts struct {
	AccountName     bool                                                                        `json:"account_name"`
	AccountNumber   bool                                                                        `json:"account_number"`
	AccountType     bool                                                                        `json:"account_type"`
	InstitutionName bool                                                                        `json:"institution_name"`
	RoutingNumber   bool                                                                        `json:"routing_number"`
	JSON            providerListResponseAuthenticationMethodsSupportedFieldsCompanyAccountsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsCompanyAccountsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyAccounts]
type providerListResponseAuthenticationMethodsSupportedFieldsCompanyAccountsJSON struct {
	AccountName     apijson.Field
	AccountNumber   apijson.Field
	AccountType     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyAccounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsCompanyAccountsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartments struct {
	Name   bool                                                                             `json:"name"`
	Parent ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent `json:"parent"`
	JSON   providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON   `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartments]
type providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON struct {
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent struct {
	Name bool                                                                                 `json:"name"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent]
type providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyEntity struct {
	Subtype bool                                                                      `json:"subtype"`
	Type    bool                                                                      `json:"type"`
	JSON    providerListResponseAuthenticationMethodsSupportedFieldsCompanyEntityJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsCompanyEntityJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyEntity]
type providerListResponseAuthenticationMethodsSupportedFieldsCompanyEntityJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsCompanyEntityJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyLocations struct {
	City       bool                                                                         `json:"city"`
	Country    bool                                                                         `json:"country"`
	Line1      bool                                                                         `json:"line1"`
	Line2      bool                                                                         `json:"line2"`
	PostalCode bool                                                                         `json:"postal_code"`
	State      bool                                                                         `json:"state"`
	JSON       providerListResponseAuthenticationMethodsSupportedFieldsCompanyLocationsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsCompanyLocationsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyLocations]
type providerListResponseAuthenticationMethodsSupportedFieldsCompanyLocationsJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsCompanyLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsCompanyLocationsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsDirectory struct {
	Individuals ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividuals `json:"individuals"`
	Paging      ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryPaging      `json:"paging"`
	JSON        providerListResponseAuthenticationMethodsSupportedFieldsDirectoryJSON        `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsDirectoryJSON contains
// the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsDirectory]
type providerListResponseAuthenticationMethodsSupportedFieldsDirectoryJSON struct {
	Individuals apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsDirectory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsDirectoryJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividuals struct {
	ID         bool                                                                                `json:"id"`
	Department bool                                                                                `json:"department" api:"nullable"`
	FirstName  bool                                                                                `json:"first_name"`
	IsActive   bool                                                                                `json:"is_active"`
	LastName   bool                                                                                `json:"last_name"`
	Manager    ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager `json:"manager"`
	MiddleName bool                                                                                `json:"middle_name"`
	JSON       providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON    `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividuals]
type providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON struct {
	ID          apijson.Field
	Department  apijson.Field
	FirstName   apijson.Field
	IsActive    apijson.Field
	LastName    apijson.Field
	Manager     apijson.Field
	MiddleName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividuals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager struct {
	ID   bool                                                                                    `json:"id"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager]
type providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryPaging struct {
	Count  bool                                                                        `json:"count"`
	Offset bool                                                                        `json:"offset"`
	JSON   providerListResponseAuthenticationMethodsSupportedFieldsDirectoryPagingJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsDirectoryPagingJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryPaging]
type providerListResponseAuthenticationMethodsSupportedFieldsDirectoryPagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsDirectoryPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsDirectoryPagingJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsEmployment struct {
	ID               bool                                                                         `json:"id"`
	ClassCode        bool                                                                         `json:"class_code"`
	CustomFields     bool                                                                         `json:"custom_fields"`
	Department       ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartment `json:"department"`
	Employment       ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentEmployment `json:"employment"`
	EmploymentStatus bool                                                                         `json:"employment_status"`
	EndDate          bool                                                                         `json:"end_date"`
	FirstName        bool                                                                         `json:"first_name"`
	Income           ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentIncome     `json:"income"`
	IncomeHistory    bool                                                                         `json:"income_history"`
	IsActive         bool                                                                         `json:"is_active"`
	LastName         bool                                                                         `json:"last_name"`
	Location         ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentLocation   `json:"location"`
	Manager          ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentManager    `json:"manager" api:"nullable"`
	MiddleName       bool                                                                         `json:"middle_name"`
	StartDate        bool                                                                         `json:"start_date"`
	Title            bool                                                                         `json:"title"`
	JSON             providerListResponseAuthenticationMethodsSupportedFieldsEmploymentJSON       `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsEmploymentJSON contains
// the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsEmployment]
type providerListResponseAuthenticationMethodsSupportedFieldsEmploymentJSON struct {
	ID               apijson.Field
	ClassCode        apijson.Field
	CustomFields     apijson.Field
	Department       apijson.Field
	Employment       apijson.Field
	EmploymentStatus apijson.Field
	EndDate          apijson.Field
	FirstName        apijson.Field
	Income           apijson.Field
	IncomeHistory    apijson.Field
	IsActive         apijson.Field
	LastName         apijson.Field
	Location         apijson.Field
	Manager          apijson.Field
	MiddleName       apijson.Field
	StartDate        apijson.Field
	Title            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsEmploymentJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartment struct {
	Name bool                                                                             `json:"name"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartment]
type providerListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentEmployment struct {
	Subtype bool                                                                             `json:"subtype"`
	Type    bool                                                                             `json:"type"`
	JSON    providerListResponseAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentEmployment]
type providerListResponseAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentIncome struct {
	Amount   bool                                                                         `json:"amount"`
	Currency bool                                                                         `json:"currency"`
	Unit     bool                                                                         `json:"unit"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentIncome]
type providerListResponseAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentIncome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentLocation struct {
	City       bool                                                                           `json:"city"`
	Country    bool                                                                           `json:"country"`
	Line1      bool                                                                           `json:"line1"`
	Line2      bool                                                                           `json:"line2"`
	PostalCode bool                                                                           `json:"postal_code"`
	State      bool                                                                           `json:"state"`
	JSON       providerListResponseAuthenticationMethodsSupportedFieldsEmploymentLocationJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsEmploymentLocationJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentLocation]
type providerListResponseAuthenticationMethodsSupportedFieldsEmploymentLocationJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsEmploymentLocationJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentManager struct {
	ID   bool                                                                          `json:"id"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsEmploymentManagerJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsEmploymentManagerJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentManager]
type providerListResponseAuthenticationMethodsSupportedFieldsEmploymentManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsEmploymentManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsEmploymentManagerJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsIndividual struct {
	ID            bool                                                                           `json:"id"`
	Dob           bool                                                                           `json:"dob"`
	Emails        ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualEmails       `json:"emails"`
	EncryptedSsn  bool                                                                           `json:"encrypted_ssn"`
	Ethnicity     bool                                                                           `json:"ethnicity"`
	FirstName     bool                                                                           `json:"first_name"`
	Gender        bool                                                                           `json:"gender"`
	LastName      bool                                                                           `json:"last_name"`
	MiddleName    bool                                                                           `json:"middle_name"`
	PhoneNumbers  ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers `json:"phone_numbers"`
	PreferredName bool                                                                           `json:"preferred_name"`
	Residence     ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualResidence    `json:"residence"`
	Ssn           bool                                                                           `json:"ssn"`
	JSON          providerListResponseAuthenticationMethodsSupportedFieldsIndividualJSON         `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsIndividualJSON contains
// the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsIndividual]
type providerListResponseAuthenticationMethodsSupportedFieldsIndividualJSON struct {
	ID            apijson.Field
	Dob           apijson.Field
	Emails        apijson.Field
	EncryptedSsn  apijson.Field
	Ethnicity     apijson.Field
	FirstName     apijson.Field
	Gender        apijson.Field
	LastName      apijson.Field
	MiddleName    apijson.Field
	PhoneNumbers  apijson.Field
	PreferredName apijson.Field
	Residence     apijson.Field
	Ssn           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsIndividualJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualEmails struct {
	Data bool                                                                         `json:"data"`
	Type bool                                                                         `json:"type"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsIndividualEmailsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsIndividualEmailsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualEmails]
type providerListResponseAuthenticationMethodsSupportedFieldsIndividualEmailsJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualEmails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsIndividualEmailsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers struct {
	Data bool                                                                               `json:"data"`
	Type bool                                                                               `json:"type"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers]
type providerListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualResidence struct {
	City       bool                                                                            `json:"city"`
	Country    bool                                                                            `json:"country"`
	Line1      bool                                                                            `json:"line1"`
	Line2      bool                                                                            `json:"line2"`
	PostalCode bool                                                                            `json:"postal_code"`
	State      bool                                                                            `json:"state"`
	JSON       providerListResponseAuthenticationMethodsSupportedFieldsIndividualResidenceJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsIndividualResidenceJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualResidence]
type providerListResponseAuthenticationMethodsSupportedFieldsIndividualResidenceJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsIndividualResidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsIndividualResidenceJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayGroup struct {
	ID             bool                                                                 `json:"id"`
	IndividualIDs  bool                                                                 `json:"individual_ids"`
	Name           bool                                                                 `json:"name"`
	PayFrequencies bool                                                                 `json:"pay_frequencies"`
	JSON           providerListResponseAuthenticationMethodsSupportedFieldsPayGroupJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayGroupJSON contains
// the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayGroup]
type providerListResponseAuthenticationMethodsSupportedFieldsPayGroupJSON struct {
	ID             apijson.Field
	IndividualIDs  apijson.Field
	Name           apijson.Field
	PayFrequencies apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayGroupJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatement struct {
	Paging        ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPaging        `json:"paging"`
	PayStatements ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatements `json:"pay_statements"`
	JSON          providerListResponseAuthenticationMethodsSupportedFieldsPayStatementJSON          `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatement]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementJSON struct {
	Paging        apijson.Field
	PayStatements apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPaging struct {
	Count  bool                                                                           `json:"count" api:"required"`
	Offset bool                                                                           `json:"offset" api:"required"`
	JSON   providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPagingJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPagingJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPaging]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPagingJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatements struct {
	Earnings              ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings              `json:"earnings"`
	EmployeeDeductions    ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions    `json:"employee_deductions"`
	EmployerContributions ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions `json:"employer_contributions"`
	GrossPay              bool                                                                                                   `json:"gross_pay"`
	IndividualID          bool                                                                                                   `json:"individual_id"`
	NetPay                bool                                                                                                   `json:"net_pay"`
	PaymentMethod         bool                                                                                                   `json:"payment_method"`
	Taxes                 ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes                 `json:"taxes"`
	TotalHours            bool                                                                                                   `json:"total_hours"`
	Type                  bool                                                                                                   `json:"type"`
	JSON                  providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON                  `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatements]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON struct {
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

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings struct {
	Amount   bool                                                                                          `json:"amount"`
	Currency bool                                                                                          `json:"currency"`
	Name     bool                                                                                          `json:"name"`
	Type     bool                                                                                          `json:"type"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions struct {
	Amount   bool                                                                                                    `json:"amount"`
	Currency bool                                                                                                    `json:"currency"`
	Name     bool                                                                                                    `json:"name"`
	PreTax   bool                                                                                                    `json:"pre_tax"`
	Type     bool                                                                                                    `json:"type"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions struct {
	Amount   bool                                                                                                       `json:"amount"`
	Currency bool                                                                                                       `json:"currency"`
	Name     bool                                                                                                       `json:"name"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes struct {
	Amount   bool                                                                                       `json:"amount"`
	Currency bool                                                                                       `json:"currency"`
	Employer bool                                                                                       `json:"employer"`
	Name     bool                                                                                       `json:"name"`
	Type     bool                                                                                       `json:"type"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes]
type providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Employer    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPayment struct {
	ID             bool                                                                     `json:"id"`
	CompanyDebit   bool                                                                     `json:"company_debit"`
	DebitDate      bool                                                                     `json:"debit_date"`
	EmployeeTaxes  bool                                                                     `json:"employee_taxes"`
	EmployerTaxes  bool                                                                     `json:"employer_taxes"`
	GrossPay       bool                                                                     `json:"gross_pay"`
	IndividualIDs  bool                                                                     `json:"individual_ids"`
	NetPay         bool                                                                     `json:"net_pay"`
	PayDate        bool                                                                     `json:"pay_date"`
	PayFrequencies bool                                                                     `json:"pay_frequencies"`
	PayGroupIDs    bool                                                                     `json:"pay_group_ids"`
	PayPeriod      ProviderListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriod `json:"pay_period"`
	JSON           providerListResponseAuthenticationMethodsSupportedFieldsPaymentJSON      `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPaymentJSON contains the
// JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPayment]
type providerListResponseAuthenticationMethodsSupportedFieldsPaymentJSON struct {
	ID             apijson.Field
	CompanyDebit   apijson.Field
	DebitDate      apijson.Field
	EmployeeTaxes  apijson.Field
	EmployerTaxes  apijson.Field
	GrossPay       apijson.Field
	IndividualIDs  apijson.Field
	NetPay         apijson.Field
	PayDate        apijson.Field
	PayFrequencies apijson.Field
	PayGroupIDs    apijson.Field
	PayPeriod      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPaymentJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriod struct {
	EndDate   bool                                                                         `json:"end_date"`
	StartDate bool                                                                         `json:"start_date"`
	JSON      providerListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriod]
type providerListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON struct {
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependents struct {
	Coverage    ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverage `json:"coverage"`
	DateOfBirth bool                                                                           `json:"date_of_birth"`
	DependentID bool                                                                           `json:"dependent_id"`
	FirstName   bool                                                                           `json:"first_name"`
	Gender      bool                                                                           `json:"gender"`
	LastName    bool                                                                           `json:"last_name"`
	MiddleName  bool                                                                           `json:"middle_name"`
	Ssn         bool                                                                           `json:"ssn"`
	JSON        providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsJSON     `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependents]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsJSON struct {
	Coverage    apijson.Field
	DateOfBirth apijson.Field
	DependentID apijson.Field
	FirstName   apijson.Field
	Gender      apijson.Field
	LastName    apijson.Field
	MiddleName  apijson.Field
	Ssn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverage struct {
	Enrollments              ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollments `json:"enrollments"`
	IndividualID             bool                                                                                      `json:"individual_id"`
	RelationshipToIndividual bool                                                                                      `json:"relationship_to_individual"`
	JSON                     providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageJSON        `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverage]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageJSON struct {
	Enrollments              apijson.Field
	IndividualID             apijson.Field
	RelationshipToIndividual apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollments struct {
	ID   bool                                                                                          `json:"id"`
	Type bool                                                                                          `json:"type"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollmentsJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollmentsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollments]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollmentsJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanDependentsCoverageEnrollmentsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollments struct {
	ID                bool                                                                                 `json:"id"`
	Contributions     ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributions `json:"contributions"`
	CoverageEndDate   bool                                                                                 `json:"coverage_end_date"`
	CoverageStartDate bool                                                                                 `json:"coverage_start_date"`
	CoverageTier      bool                                                                                 `json:"coverage_tier"`
	DependentIDs      bool                                                                                 `json:"dependent_ids"`
	IndividualID      bool                                                                                 `json:"individual_id"`
	PlanID            bool                                                                                 `json:"plan_id"`
	Status            bool                                                                                 `json:"status"`
	JSON              providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsJSON          `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollments]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsJSON struct {
	ID                apijson.Field
	Contributions     apijson.Field
	CoverageEndDate   apijson.Field
	CoverageStartDate apijson.Field
	CoverageTier      apijson.Field
	DependentIDs      apijson.Field
	IndividualID      apijson.Field
	PlanID            apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributions struct {
	EmployeeContribution ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContribution `json:"employee_contribution"`
	EmployerContribution ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContribution `json:"employer_contribution"`
	Frequency            bool                                                                                                     `json:"frequency"`
	JSON                 providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsJSON                 `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributions]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsJSON struct {
	EmployeeContribution apijson.Field
	EmployerContribution apijson.Field
	Frequency            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContribution struct {
	Amount   bool                                                                                                         `json:"amount"`
	Currency bool                                                                                                         `json:"currency"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContributionJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContributionJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContribution]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContributionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployeeContributionJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContribution struct {
	Amount   bool                                                                                                         `json:"amount"`
	Currency bool                                                                                                         `json:"currency"`
	JSON     providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContributionJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContributionJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContribution]
type providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContributionJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlanEnrollmentsContributionsEmployerContributionJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlans struct {
	ID             bool                                                                 `json:"id"`
	Carrier        ProviderListResponseAuthenticationMethodsSupportedFieldsPlansCarrier `json:"carrier"`
	CoverageTiers  bool                                                                 `json:"coverage_tiers"`
	DeductionCodes bool                                                                 `json:"deduction_codes"`
	Description    bool                                                                 `json:"description"`
	EndDate        bool                                                                 `json:"end_date"`
	Name           bool                                                                 `json:"name"`
	NetworkType    bool                                                                 `json:"network_type"`
	StartDate      bool                                                                 `json:"start_date"`
	Type           bool                                                                 `json:"type"`
	JSON           providerListResponseAuthenticationMethodsSupportedFieldsPlansJSON    `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlansJSON contains the
// JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlans]
type providerListResponseAuthenticationMethodsSupportedFieldsPlansJSON struct {
	ID             apijson.Field
	Carrier        apijson.Field
	CoverageTiers  apijson.Field
	DeductionCodes apijson.Field
	Description    apijson.Field
	EndDate        apijson.Field
	Name           apijson.Field
	NetworkType    apijson.Field
	StartDate      apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlans) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlansJSON) RawJSON() string {
	return r.raw
}

type ProviderListResponseAuthenticationMethodsSupportedFieldsPlansCarrier struct {
	ID   bool                                                                     `json:"id"`
	Name bool                                                                     `json:"name"`
	JSON providerListResponseAuthenticationMethodsSupportedFieldsPlansCarrierJSON `json:"-"`
}

// providerListResponseAuthenticationMethodsSupportedFieldsPlansCarrierJSON
// contains the JSON metadata for the struct
// [ProviderListResponseAuthenticationMethodsSupportedFieldsPlansCarrier]
type providerListResponseAuthenticationMethodsSupportedFieldsPlansCarrierJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponseAuthenticationMethodsSupportedFieldsPlansCarrier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseAuthenticationMethodsSupportedFieldsPlansCarrierJSON) RawJSON() string {
	return r.raw
}
