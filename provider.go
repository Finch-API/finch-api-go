// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// ProviderService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewProviderService] method instead.
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
func (r *ProviderService) List(ctx context.Context, opts ...option.RequestOption) (res *shared.SinglePage[Provider], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *ProviderService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *shared.SinglePageAutoPager[Provider] {
	return shared.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type Provider struct {
	// The id of the payroll provider used in Connect.
	ID string `json:"id"`
	// The list of authentication methods supported by the provider.
	AuthenticationMethods []ProviderAuthenticationMethod `json:"authentication_methods"`
	// The display name of the payroll provider.
	DisplayName string `json:"display_name"`
	// The url to the official icon of the payroll provider.
	Icon string `json:"icon"`
	// The url to the official logo of the payroll provider.
	Logo string `json:"logo"`
	// [DEPRECATED] Whether the Finch integration with this provider uses the Assisted
	// Connect Flow by default. This field is now deprecated. Please check for a `type`
	// of `assisted` in the `authentication_methods` field instead.
	Manual bool `json:"manual"`
	// whether MFA is required for the provider.
	MfaRequired bool `json:"mfa_required"`
	// The hex code for the primary color of the payroll provider.
	PrimaryColor string `json:"primary_color"`
	// The list of Finch products supported on this payroll provider.
	Products []string     `json:"products"`
	JSON     providerJSON `json:"-"`
}

// providerJSON contains the JSON metadata for the struct [Provider]
type providerJSON struct {
	ID                    apijson.Field
	AuthenticationMethods apijson.Field
	DisplayName           apijson.Field
	Icon                  apijson.Field
	Logo                  apijson.Field
	Manual                apijson.Field
	MfaRequired           apijson.Field
	PrimaryColor          apijson.Field
	Products              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Provider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethod struct {
	// Each benefit type and their supported features. If the benefit type is not
	// supported, the property will be null
	BenefitsSupport BenefitsSupport `json:"benefits_support,nullable"`
	// The supported data fields returned by our HR and payroll endpoints
	SupportedFields ProviderAuthenticationMethodsSupportedFields `json:"supported_fields,nullable"`
	// The type of authentication method.
	Type ProviderAuthenticationMethodsType `json:"type"`
	JSON providerAuthenticationMethodJSON  `json:"-"`
}

// providerAuthenticationMethodJSON contains the JSON metadata for the struct
// [ProviderAuthenticationMethod]
type providerAuthenticationMethodJSON struct {
	BenefitsSupport apijson.Field
	SupportedFields apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderAuthenticationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodJSON) RawJSON() string {
	return r.raw
}

// The supported data fields returned by our HR and payroll endpoints
type ProviderAuthenticationMethodsSupportedFields struct {
	Company      ProviderAuthenticationMethodsSupportedFieldsCompany      `json:"company"`
	Directory    ProviderAuthenticationMethodsSupportedFieldsDirectory    `json:"directory"`
	Employment   ProviderAuthenticationMethodsSupportedFieldsEmployment   `json:"employment"`
	Individual   ProviderAuthenticationMethodsSupportedFieldsIndividual   `json:"individual"`
	PayStatement ProviderAuthenticationMethodsSupportedFieldsPayStatement `json:"pay_statement"`
	Payment      ProviderAuthenticationMethodsSupportedFieldsPayment      `json:"payment"`
	JSON         providerAuthenticationMethodsSupportedFieldsJSON         `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsJSON contains the JSON metadata for
// the struct [ProviderAuthenticationMethodsSupportedFields]
type providerAuthenticationMethodsSupportedFieldsJSON struct {
	Company      apijson.Field
	Directory    apijson.Field
	Employment   apijson.Field
	Individual   apijson.Field
	PayStatement apijson.Field
	Payment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsCompany struct {
	ID                 bool                                                           `json:"id"`
	Accounts           ProviderAuthenticationMethodsSupportedFieldsCompanyAccounts    `json:"accounts"`
	Departments        ProviderAuthenticationMethodsSupportedFieldsCompanyDepartments `json:"departments"`
	Ein                bool                                                           `json:"ein"`
	Entity             ProviderAuthenticationMethodsSupportedFieldsCompanyEntity      `json:"entity"`
	LegalName          bool                                                           `json:"legal_name"`
	Locations          ProviderAuthenticationMethodsSupportedFieldsCompanyLocations   `json:"locations"`
	PrimaryEmail       bool                                                           `json:"primary_email"`
	PrimaryPhoneNumber bool                                                           `json:"primary_phone_number"`
	JSON               providerAuthenticationMethodsSupportedFieldsCompanyJSON        `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsCompanyJSON contains the JSON
// metadata for the struct [ProviderAuthenticationMethodsSupportedFieldsCompany]
type providerAuthenticationMethodsSupportedFieldsCompanyJSON struct {
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

func (r *ProviderAuthenticationMethodsSupportedFieldsCompany) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsCompanyJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsCompanyAccounts struct {
	AccountName     bool                                                            `json:"account_name"`
	AccountNumber   bool                                                            `json:"account_number"`
	AccountType     bool                                                            `json:"account_type"`
	InstitutionName bool                                                            `json:"institution_name"`
	RoutingNumber   bool                                                            `json:"routing_number"`
	JSON            providerAuthenticationMethodsSupportedFieldsCompanyAccountsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsCompanyAccountsJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsCompanyAccounts]
type providerAuthenticationMethodsSupportedFieldsCompanyAccountsJSON struct {
	AccountName     apijson.Field
	AccountNumber   apijson.Field
	AccountType     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsCompanyAccounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsCompanyAccountsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsCompanyDepartments struct {
	Name   bool                                                                 `json:"name"`
	Parent ProviderAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent `json:"parent"`
	JSON   providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON   `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsCompanyDepartments]
type providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON struct {
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsCompanyDepartments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent struct {
	Name bool                                                                     `json:"name"`
	JSON providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent]
type providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsCompanyEntity struct {
	Subtype bool                                                          `json:"subtype"`
	Type    bool                                                          `json:"type"`
	JSON    providerAuthenticationMethodsSupportedFieldsCompanyEntityJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsCompanyEntityJSON contains the JSON
// metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsCompanyEntity]
type providerAuthenticationMethodsSupportedFieldsCompanyEntityJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsCompanyEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsCompanyEntityJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsCompanyLocations struct {
	City       bool                                                             `json:"city"`
	Country    bool                                                             `json:"country"`
	Line1      bool                                                             `json:"line1"`
	Line2      bool                                                             `json:"line2"`
	PostalCode bool                                                             `json:"postal_code"`
	State      bool                                                             `json:"state"`
	JSON       providerAuthenticationMethodsSupportedFieldsCompanyLocationsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsCompanyLocationsJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsCompanyLocations]
type providerAuthenticationMethodsSupportedFieldsCompanyLocationsJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsCompanyLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsCompanyLocationsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsDirectory struct {
	Individuals ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividuals `json:"individuals"`
	Paging      ProviderAuthenticationMethodsSupportedFieldsDirectoryPaging      `json:"paging"`
	JSON        providerAuthenticationMethodsSupportedFieldsDirectoryJSON        `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsDirectoryJSON contains the JSON
// metadata for the struct [ProviderAuthenticationMethodsSupportedFieldsDirectory]
type providerAuthenticationMethodsSupportedFieldsDirectoryJSON struct {
	Individuals apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsDirectory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsDirectoryJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividuals struct {
	ID         bool                                                                    `json:"id"`
	Department bool                                                                    `json:"department"`
	FirstName  bool                                                                    `json:"first_name"`
	IsActive   bool                                                                    `json:"is_active"`
	LastName   bool                                                                    `json:"last_name"`
	Manager    ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager `json:"manager"`
	MiddleName bool                                                                    `json:"middle_name"`
	JSON       providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON    `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON contains
// the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividuals]
type providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON struct {
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

func (r *ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividuals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager struct {
	ID   bool                                                                        `json:"id"`
	JSON providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager]
type providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsDirectoryPaging struct {
	Count  bool                                                            `json:"count"`
	Offset bool                                                            `json:"offset"`
	JSON   providerAuthenticationMethodsSupportedFieldsDirectoryPagingJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsDirectoryPagingJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsDirectoryPaging]
type providerAuthenticationMethodsSupportedFieldsDirectoryPagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsDirectoryPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsDirectoryPagingJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsEmployment struct {
	ID            bool                                                             `json:"id"`
	ClassCode     bool                                                             `json:"class_code"`
	CustomFields  bool                                                             `json:"custom_fields"`
	Department    ProviderAuthenticationMethodsSupportedFieldsEmploymentDepartment `json:"department"`
	Employment    ProviderAuthenticationMethodsSupportedFieldsEmploymentEmployment `json:"employment"`
	EndDate       bool                                                             `json:"end_date"`
	FirstName     bool                                                             `json:"first_name"`
	Income        ProviderAuthenticationMethodsSupportedFieldsEmploymentIncome     `json:"income"`
	IncomeHistory bool                                                             `json:"income_history"`
	IsActive      bool                                                             `json:"is_active"`
	LastName      bool                                                             `json:"last_name"`
	Location      ProviderAuthenticationMethodsSupportedFieldsEmploymentLocation   `json:"location"`
	Manager       ProviderAuthenticationMethodsSupportedFieldsEmploymentManager    `json:"manager"`
	MiddleName    bool                                                             `json:"middle_name"`
	StartDate     bool                                                             `json:"start_date"`
	Title         bool                                                             `json:"title"`
	JSON          providerAuthenticationMethodsSupportedFieldsEmploymentJSON       `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsEmploymentJSON contains the JSON
// metadata for the struct [ProviderAuthenticationMethodsSupportedFieldsEmployment]
type providerAuthenticationMethodsSupportedFieldsEmploymentJSON struct {
	ID            apijson.Field
	ClassCode     apijson.Field
	CustomFields  apijson.Field
	Department    apijson.Field
	Employment    apijson.Field
	EndDate       apijson.Field
	FirstName     apijson.Field
	Income        apijson.Field
	IncomeHistory apijson.Field
	IsActive      apijson.Field
	LastName      apijson.Field
	Location      apijson.Field
	Manager       apijson.Field
	MiddleName    apijson.Field
	StartDate     apijson.Field
	Title         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsEmploymentJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsEmploymentDepartment struct {
	Name bool                                                                 `json:"name"`
	JSON providerAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON contains
// the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsEmploymentDepartment]
type providerAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsEmploymentDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsEmploymentEmployment struct {
	Subtype bool                                                                 `json:"subtype"`
	Type    bool                                                                 `json:"type"`
	JSON    providerAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON contains
// the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsEmploymentEmployment]
type providerAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsEmploymentEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsEmploymentIncome struct {
	Amount   bool                                                             `json:"amount"`
	Currency bool                                                             `json:"currency"`
	Unit     bool                                                             `json:"unit"`
	JSON     providerAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsEmploymentIncome]
type providerAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsEmploymentIncome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsEmploymentLocation struct {
	City       bool                                                               `json:"city"`
	Country    bool                                                               `json:"country"`
	Line1      bool                                                               `json:"line1"`
	Line2      bool                                                               `json:"line2"`
	PostalCode bool                                                               `json:"postal_code"`
	State      bool                                                               `json:"state"`
	JSON       providerAuthenticationMethodsSupportedFieldsEmploymentLocationJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsEmploymentLocationJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsEmploymentLocation]
type providerAuthenticationMethodsSupportedFieldsEmploymentLocationJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsEmploymentLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsEmploymentLocationJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsEmploymentManager struct {
	ID   bool                                                              `json:"id"`
	JSON providerAuthenticationMethodsSupportedFieldsEmploymentManagerJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsEmploymentManagerJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsEmploymentManager]
type providerAuthenticationMethodsSupportedFieldsEmploymentManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsEmploymentManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsEmploymentManagerJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsIndividual struct {
	ID            bool                                                               `json:"id"`
	Dob           bool                                                               `json:"dob"`
	Emails        ProviderAuthenticationMethodsSupportedFieldsIndividualEmails       `json:"emails"`
	EncryptedSsn  bool                                                               `json:"encrypted_ssn"`
	Ethnicity     bool                                                               `json:"ethnicity"`
	FirstName     bool                                                               `json:"first_name"`
	Gender        bool                                                               `json:"gender"`
	LastName      bool                                                               `json:"last_name"`
	MiddleName    bool                                                               `json:"middle_name"`
	PhoneNumbers  ProviderAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers `json:"phone_numbers"`
	PreferredName bool                                                               `json:"preferred_name"`
	Residence     ProviderAuthenticationMethodsSupportedFieldsIndividualResidence    `json:"residence"`
	Ssn           bool                                                               `json:"ssn"`
	JSON          providerAuthenticationMethodsSupportedFieldsIndividualJSON         `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsIndividualJSON contains the JSON
// metadata for the struct [ProviderAuthenticationMethodsSupportedFieldsIndividual]
type providerAuthenticationMethodsSupportedFieldsIndividualJSON struct {
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

func (r *ProviderAuthenticationMethodsSupportedFieldsIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsIndividualJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsIndividualEmails struct {
	Data bool                                                             `json:"data"`
	Type bool                                                             `json:"type"`
	JSON providerAuthenticationMethodsSupportedFieldsIndividualEmailsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsIndividualEmailsJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsIndividualEmails]
type providerAuthenticationMethodsSupportedFieldsIndividualEmailsJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsIndividualEmails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsIndividualEmailsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers struct {
	Data bool                                                                   `json:"data"`
	Type bool                                                                   `json:"type"`
	JSON providerAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON contains
// the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers]
type providerAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsIndividualResidence struct {
	City       bool                                                                `json:"city"`
	Country    bool                                                                `json:"country"`
	Line1      bool                                                                `json:"line1"`
	Line2      bool                                                                `json:"line2"`
	PostalCode bool                                                                `json:"postal_code"`
	State      bool                                                                `json:"state"`
	JSON       providerAuthenticationMethodsSupportedFieldsIndividualResidenceJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsIndividualResidenceJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsIndividualResidence]
type providerAuthenticationMethodsSupportedFieldsIndividualResidenceJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsIndividualResidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsIndividualResidenceJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatement struct {
	Paging        ProviderAuthenticationMethodsSupportedFieldsPayStatementPaging        `json:"paging"`
	PayStatements ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements `json:"pay_statements"`
	JSON          providerAuthenticationMethodsSupportedFieldsPayStatementJSON          `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementJSON contains the JSON
// metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatement]
type providerAuthenticationMethodsSupportedFieldsPayStatementJSON struct {
	Paging        apijson.Field
	PayStatements apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPaging struct {
	Count  bool                                                               `json:"count,required"`
	Offset bool                                                               `json:"offset,required"`
	JSON   providerAuthenticationMethodsSupportedFieldsPayStatementPagingJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPagingJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPaging]
type providerAuthenticationMethodsSupportedFieldsPayStatementPagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPagingJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements struct {
	Earnings              ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings              `json:"earnings"`
	EmployeeDeductions    ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions    `json:"employee_deductions"`
	EmployerContributions ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions `json:"employer_contributions"`
	// [DEPRECATED] Use `employer_contributions` instead
	EmployerDeductions ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductions `json:"employer_deductions"`
	GrossPay           bool                                                                                    `json:"gross_pay"`
	IndividualID       bool                                                                                    `json:"individual_id"`
	NetPay             bool                                                                                    `json:"net_pay"`
	PaymentMethod      bool                                                                                    `json:"payment_method"`
	Taxes              ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes              `json:"taxes"`
	TotalHours         bool                                                                                    `json:"total_hours"`
	Type               bool                                                                                    `json:"type"`
	JSON               providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON               `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON struct {
	Earnings              apijson.Field
	EmployeeDeductions    apijson.Field
	EmployerContributions apijson.Field
	EmployerDeductions    apijson.Field
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

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings struct {
	Amount   bool                                                                              `json:"amount"`
	Currency bool                                                                              `json:"currency"`
	Name     bool                                                                              `json:"name"`
	Type     bool                                                                              `json:"type"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions struct {
	Amount   bool                                                                                        `json:"amount"`
	Currency bool                                                                                        `json:"currency"`
	Name     bool                                                                                        `json:"name"`
	PreTax   bool                                                                                        `json:"pre_tax"`
	Type     bool                                                                                        `json:"type"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions struct {
	Amount   bool                                                                                           `json:"amount"`
	Currency bool                                                                                           `json:"currency"`
	Name     bool                                                                                           `json:"name"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON) RawJSON() string {
	return r.raw
}

// [DEPRECATED] Use `employer_contributions` instead
type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductions struct {
	Amount   bool                                                                                        `json:"amount"`
	Currency bool                                                                                        `json:"currency"`
	Name     bool                                                                                        `json:"name"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductions]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes struct {
	Amount   bool                                                                           `json:"amount"`
	Currency bool                                                                           `json:"currency"`
	Employer bool                                                                           `json:"employer"`
	Name     bool                                                                           `json:"name"`
	Type     bool                                                                           `json:"type"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Employer    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPayment struct {
	ID            bool                                                         `json:"id"`
	CompanyDebit  bool                                                         `json:"company_debit"`
	DebitDate     bool                                                         `json:"debit_date"`
	EmployeeTaxes bool                                                         `json:"employee_taxes"`
	EmployerTaxes bool                                                         `json:"employer_taxes"`
	GrossPay      bool                                                         `json:"gross_pay"`
	IndividualIDs bool                                                         `json:"individual_ids"`
	NetPay        bool                                                         `json:"net_pay"`
	PayDate       bool                                                         `json:"pay_date"`
	PayPeriod     ProviderAuthenticationMethodsSupportedFieldsPaymentPayPeriod `json:"pay_period"`
	JSON          providerAuthenticationMethodsSupportedFieldsPaymentJSON      `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPaymentJSON contains the JSON
// metadata for the struct [ProviderAuthenticationMethodsSupportedFieldsPayment]
type providerAuthenticationMethodsSupportedFieldsPaymentJSON struct {
	ID            apijson.Field
	CompanyDebit  apijson.Field
	DebitDate     apijson.Field
	EmployeeTaxes apijson.Field
	EmployerTaxes apijson.Field
	GrossPay      apijson.Field
	IndividualIDs apijson.Field
	NetPay        apijson.Field
	PayDate       apijson.Field
	PayPeriod     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPaymentJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthenticationMethodsSupportedFieldsPaymentPayPeriod struct {
	EndDate   bool                                                             `json:"end_date"`
	StartDate bool                                                             `json:"start_date"`
	JSON      providerAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON `json:"-"`
}

// providerAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON contains the
// JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPaymentPayPeriod]
type providerAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON struct {
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPaymentPayPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON) RawJSON() string {
	return r.raw
}

// The type of authentication method.
type ProviderAuthenticationMethodsType string

const (
	ProviderAuthenticationMethodsTypeAssisted      ProviderAuthenticationMethodsType = "assisted"
	ProviderAuthenticationMethodsTypeCredential    ProviderAuthenticationMethodsType = "credential"
	ProviderAuthenticationMethodsTypeAPIToken      ProviderAuthenticationMethodsType = "api_token"
	ProviderAuthenticationMethodsTypeAPICredential ProviderAuthenticationMethodsType = "api_credential"
	ProviderAuthenticationMethodsTypeOAuth         ProviderAuthenticationMethodsType = "oauth"
)

func (r ProviderAuthenticationMethodsType) IsKnown() bool {
	switch r {
	case ProviderAuthenticationMethodsTypeAssisted, ProviderAuthenticationMethodsTypeCredential, ProviderAuthenticationMethodsTypeAPIToken, ProviderAuthenticationMethodsTypeAPICredential, ProviderAuthenticationMethodsTypeOAuth:
		return true
	}
	return false
}
