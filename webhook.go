// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// WebhookService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Validates that the given payload was sent by Finch and parses the payload.
func (r *WebhookService) Unwrap(payload []byte, headers http.Header, secret string, now time.Time) (res WebhookEvent, err error) {
	err = r.VerifySignature(payload, headers, secret, now)
	if err != nil {
		return nil, err
	}

	event := WebhookEvent(nil)
	err = apijson.UnmarshalRoot(payload, &event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// Validates whether or not the webhook payload was sent by Finch.
//
// An error will be raised if the webhook payload was not sent by Finch.
func (r *WebhookService) VerifySignature(payload []byte, headers http.Header, secret string, now time.Time) (err error) {
	parsedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return fmt.Errorf("invalid webhook secret: %s", err)
	}

	id := headers.Get("finch-event-id")
	if len(id) == 0 {
		return errors.New("could not find finch-event-id header")
	}
	sign := headers.Values("finch-signature")
	if len(sign) == 0 {
		return errors.New("could not find finch-signature header")
	}
	unixtime := headers.Get("finch-timestamp")
	if len(unixtime) == 0 {
		return errors.New("could not find finch-timestamp header")
	}

	timestamp, err := strconv.ParseInt(unixtime, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp header: %s, %s", unixtime, err)
	}

	if timestamp < now.Unix()-300 {
		return errors.New("webhook timestamp too old")
	}
	if timestamp > now.Unix()+300 {
		return errors.New("webhook timestamp too new")
	}

	mac := hmac.New(sha256.New, parsedSecret)
	mac.Write([]byte(id))
	mac.Write([]byte("."))
	mac.Write([]byte(unixtime))
	mac.Write([]byte("."))
	mac.Write(payload)
	expected := mac.Sum(nil)

	for _, part := range sign {
		parts := strings.Split(part, ",")
		if len(parts) != 2 {
			continue
		}
		if parts[0] != "v1" {
			continue
		}
		signature, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			continue
		}
		if hmac.Equal(signature, expected) {
			return nil
		}
	}

	return errors.New("None of the given webhook signatures match the expected signature")
}

type AccountUpdateEvent struct {
	Data      AccountUpdateEventData      `json:"data"`
	EventType AccountUpdateEventEventType `json:"event_type"`
	JSON      accountUpdateEventJSON      `json:"-"`
	BaseWebhookEvent
}

// accountUpdateEventJSON contains the JSON metadata for the struct
// [AccountUpdateEvent]
type accountUpdateEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountUpdateEvent) implementsWebhookEvent() {}

type AccountUpdateEventData struct {
	AuthenticationMethod AccountUpdateEventDataAuthenticationMethod `json:"authentication_method,required"`
	Status               shared.ConnectionStatusType                `json:"status,required"`
	JSON                 accountUpdateEventDataJSON                 `json:"-"`
}

// accountUpdateEventDataJSON contains the JSON metadata for the struct
// [AccountUpdateEventData]
type accountUpdateEventDataJSON struct {
	AuthenticationMethod apijson.Field
	Status               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountUpdateEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethod struct {
	// Each benefit type and their supported features. If the benefit type is not
	// supported, the property will be null
	BenefitsSupport BenefitsSupport `json:"benefits_support,nullable"`
	// The supported data fields returned by our HR and payroll endpoints
	SupportedFields AccountUpdateEventDataAuthenticationMethodSupportedFields `json:"supported_fields,nullable"`
	// The type of authentication method.
	Type AccountUpdateEventDataAuthenticationMethodType `json:"type"`
	JSON accountUpdateEventDataAuthenticationMethodJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodJSON contains the JSON metadata for
// the struct [AccountUpdateEventDataAuthenticationMethod]
type accountUpdateEventDataAuthenticationMethodJSON struct {
	BenefitsSupport apijson.Field
	SupportedFields apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The supported data fields returned by our HR and payroll endpoints
type AccountUpdateEventDataAuthenticationMethodSupportedFields struct {
	Company      AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompany      `json:"company"`
	Directory    AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectory    `json:"directory"`
	Employment   AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmployment   `json:"employment"`
	Individual   AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividual   `json:"individual"`
	PayStatement AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatement `json:"pay_statement"`
	Payment      AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayment      `json:"payment"`
	JSON         accountUpdateEventDataAuthenticationMethodSupportedFieldsJSON         `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsJSON contains the JSON
// metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFields]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsJSON struct {
	Company      apijson.Field
	Directory    apijson.Field
	Employment   apijson.Field
	Individual   apijson.Field
	PayStatement apijson.Field
	Payment      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompany struct {
	ID                 bool                                                                        `json:"id"`
	Accounts           AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccounts    `json:"accounts"`
	Departments        AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartments `json:"departments"`
	Ein                bool                                                                        `json:"ein"`
	Entity             AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntity      `json:"entity"`
	LegalName          bool                                                                        `json:"legal_name"`
	Locations          AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocations   `json:"locations"`
	PrimaryEmail       bool                                                                        `json:"primary_email"`
	PrimaryPhoneNumber bool                                                                        `json:"primary_phone_number"`
	JSON               accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyJSON        `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyJSON contains
// the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompany]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyJSON struct {
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

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompany) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccounts struct {
	AccountName     bool                                                                         `json:"account_name"`
	AccountNumber   bool                                                                         `json:"account_number"`
	AccountType     bool                                                                         `json:"account_type"`
	InstitutionName bool                                                                         `json:"institution_name"`
	RoutingNumber   bool                                                                         `json:"routing_number"`
	JSON            accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccountsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccountsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccounts]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccountsJSON struct {
	AccountName     apijson.Field
	AccountNumber   apijson.Field
	AccountType     apijson.Field
	InstitutionName apijson.Field
	RoutingNumber   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyAccounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartments struct {
	Name   bool                                                                              `json:"name"`
	Parent AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParent `json:"parent"`
	JSON   accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsJSON   `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartments]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsJSON struct {
	Name        apijson.Field
	Parent      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParent struct {
	Name bool                                                                                  `json:"name"`
	JSON accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParentJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParentJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParent]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyDepartmentsParent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntity struct {
	Subtype bool                                                                       `json:"subtype"`
	Type    bool                                                                       `json:"type"`
	JSON    accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntityJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntityJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntity]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntityJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocations struct {
	City       bool                                                                          `json:"city"`
	Country    bool                                                                          `json:"country"`
	Line1      bool                                                                          `json:"line1"`
	Line2      bool                                                                          `json:"line2"`
	PostalCode bool                                                                          `json:"postal_code"`
	State      bool                                                                          `json:"state"`
	JSON       accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocationsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocationsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocations]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocationsJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsCompanyLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectory struct {
	Individuals AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividuals `json:"individuals"`
	Paging      AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPaging      `json:"paging"`
	JSON        accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryJSON        `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryJSON contains
// the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectory]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryJSON struct {
	Individuals apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividuals struct {
	ID         bool                                                                                 `json:"id"`
	Department bool                                                                                 `json:"department"`
	FirstName  bool                                                                                 `json:"first_name"`
	IsActive   bool                                                                                 `json:"is_active"`
	LastName   bool                                                                                 `json:"last_name"`
	Manager    AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManager `json:"manager"`
	MiddleName bool                                                                                 `json:"middle_name"`
	JSON       accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsJSON    `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividuals]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsJSON struct {
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

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividuals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManager struct {
	ID   bool                                                                                     `json:"id"`
	JSON accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManagerJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManagerJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManager]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryIndividualsManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPaging struct {
	Count  bool                                                                         `json:"count"`
	Offset bool                                                                         `json:"offset"`
	JSON   accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPagingJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPagingJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPaging]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsDirectoryPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmployment struct {
	ID            bool                                                                          `json:"id"`
	ClassCode     bool                                                                          `json:"class_code"`
	CustomFields  bool                                                                          `json:"custom_fields"`
	Department    AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartment `json:"department"`
	Employment    AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmployment `json:"employment"`
	EndDate       bool                                                                          `json:"end_date"`
	FirstName     bool                                                                          `json:"first_name"`
	Income        AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncome     `json:"income"`
	IncomeHistory bool                                                                          `json:"income_history"`
	IsActive      bool                                                                          `json:"is_active"`
	LastName      bool                                                                          `json:"last_name"`
	Location      AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocation   `json:"location"`
	Manager       AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManager    `json:"manager"`
	MiddleName    bool                                                                          `json:"middle_name"`
	StartDate     bool                                                                          `json:"start_date"`
	Title         bool                                                                          `json:"title"`
	JSON          accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentJSON       `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentJSON contains
// the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmployment]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentJSON struct {
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

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartment struct {
	Name bool                                                                              `json:"name"`
	JSON accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartmentJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartmentJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartment]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmployment struct {
	Subtype bool                                                                              `json:"subtype"`
	Type    bool                                                                              `json:"type"`
	JSON    accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmploymentJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmploymentJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmployment]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmploymentJSON struct {
	Subtype     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentEmployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncome struct {
	Amount   bool                                                                          `json:"amount"`
	Currency bool                                                                          `json:"currency"`
	Unit     bool                                                                          `json:"unit"`
	JSON     accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncomeJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncomeJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncome]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncomeJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentIncome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocation struct {
	City       bool                                                                            `json:"city"`
	Country    bool                                                                            `json:"country"`
	Line1      bool                                                                            `json:"line1"`
	Line2      bool                                                                            `json:"line2"`
	PostalCode bool                                                                            `json:"postal_code"`
	State      bool                                                                            `json:"state"`
	JSON       accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocationJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocationJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocation]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocationJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManager struct {
	ID   bool                                                                           `json:"id"`
	JSON accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManagerJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManagerJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManager]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsEmploymentManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividual struct {
	ID            bool                                                                            `json:"id"`
	Dob           bool                                                                            `json:"dob"`
	Emails        AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmails       `json:"emails"`
	EncryptedSsn  bool                                                                            `json:"encrypted_ssn"`
	Ethnicity     bool                                                                            `json:"ethnicity"`
	FirstName     bool                                                                            `json:"first_name"`
	Gender        bool                                                                            `json:"gender"`
	LastName      bool                                                                            `json:"last_name"`
	MiddleName    bool                                                                            `json:"middle_name"`
	PhoneNumbers  AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbers `json:"phone_numbers"`
	PreferredName bool                                                                            `json:"preferred_name"`
	Residence     AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidence    `json:"residence"`
	Ssn           bool                                                                            `json:"ssn"`
	JSON          accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualJSON         `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualJSON contains
// the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividual]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualJSON struct {
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

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmails struct {
	Data bool                                                                          `json:"data"`
	Type bool                                                                          `json:"type"`
	JSON accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmailsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmailsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmails]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmailsJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualEmails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbers struct {
	Data bool                                                                                `json:"data"`
	Type bool                                                                                `json:"type"`
	JSON accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbersJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbersJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbers]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbersJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualPhoneNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidence struct {
	City       bool                                                                             `json:"city"`
	Country    bool                                                                             `json:"country"`
	Line1      bool                                                                             `json:"line1"`
	Line2      bool                                                                             `json:"line2"`
	PostalCode bool                                                                             `json:"postal_code"`
	State      bool                                                                             `json:"state"`
	JSON       accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidenceJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidenceJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidence]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidenceJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsIndividualResidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatement struct {
	Paging        AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPaging        `json:"paging"`
	PayStatements AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatements `json:"pay_statements"`
	JSON          accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementJSON          `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatement]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementJSON struct {
	Paging        apijson.Field
	PayStatements apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPaging struct {
	Count  bool                                                                            `json:"count,required"`
	Offset bool                                                                            `json:"offset,required"`
	JSON   accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPagingJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPagingJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPaging]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatements struct {
	Earnings              AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarnings              `json:"earnings"`
	EmployeeDeductions    AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductions    `json:"employee_deductions"`
	EmployerContributions AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributions `json:"employer_contributions"`
	// [DEPRECATED] Use `employer_contributions` instead
	EmployerDeductions AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductions `json:"employer_deductions"`
	GrossPay           bool                                                                                                 `json:"gross_pay"`
	IndividualID       bool                                                                                                 `json:"individual_id"`
	NetPay             bool                                                                                                 `json:"net_pay"`
	PaymentMethod      bool                                                                                                 `json:"payment_method"`
	Taxes              AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxes              `json:"taxes"`
	TotalHours         bool                                                                                                 `json:"total_hours"`
	Type               bool                                                                                                 `json:"type"`
	JSON               accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsJSON               `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatements]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsJSON struct {
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

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarnings struct {
	Amount   bool                                                                                           `json:"amount"`
	Currency bool                                                                                           `json:"currency"`
	Name     bool                                                                                           `json:"name"`
	Type     bool                                                                                           `json:"type"`
	JSON     accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarningsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarningsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarnings]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarningsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEarnings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductions struct {
	Amount   bool                                                                                                     `json:"amount"`
	Currency bool                                                                                                     `json:"currency"`
	Name     bool                                                                                                     `json:"name"`
	PreTax   bool                                                                                                     `json:"pre_tax"`
	Type     bool                                                                                                     `json:"type"`
	JSON     accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductions]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	PreTax      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployeeDeductions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributions struct {
	Amount   bool                                                                                                        `json:"amount"`
	Currency bool                                                                                                        `json:"currency"`
	Name     bool                                                                                                        `json:"name"`
	JSON     accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributions]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerContributions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// [DEPRECATED] Use `employer_contributions` instead
type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductions struct {
	Amount   bool                                                                                                     `json:"amount"`
	Currency bool                                                                                                     `json:"currency"`
	Name     bool                                                                                                     `json:"name"`
	JSON     accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductions]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsEmployerDeductions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxes struct {
	Amount   bool                                                                                        `json:"amount"`
	Currency bool                                                                                        `json:"currency"`
	Employer bool                                                                                        `json:"employer"`
	Name     bool                                                                                        `json:"name"`
	Type     bool                                                                                        `json:"type"`
	JSON     accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxesJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxesJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxes]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxesJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	Employer    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayStatementPayStatementsTaxes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayment struct {
	ID            bool                                                                      `json:"id"`
	CompanyDebit  bool                                                                      `json:"company_debit"`
	DebitDate     bool                                                                      `json:"debit_date"`
	EmployeeTaxes bool                                                                      `json:"employee_taxes"`
	EmployerTaxes bool                                                                      `json:"employer_taxes"`
	GrossPay      bool                                                                      `json:"gross_pay"`
	IndividualIDs bool                                                                      `json:"individual_ids"`
	NetPay        bool                                                                      `json:"net_pay"`
	PayDate       bool                                                                      `json:"pay_date"`
	PayPeriod     AccountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriod `json:"pay_period"`
	JSON          accountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentJSON      `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentJSON contains
// the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayment]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentJSON struct {
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

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriod struct {
	EndDate   bool                                                                          `json:"end_date"`
	StartDate bool                                                                          `json:"start_date"`
	JSON      accountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriodJSON `json:"-"`
}

// accountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriodJSON
// contains the JSON metadata for the struct
// [AccountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriod]
type accountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriodJSON struct {
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateEventDataAuthenticationMethodSupportedFieldsPaymentPayPeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of authentication method.
type AccountUpdateEventDataAuthenticationMethodType string

const (
	AccountUpdateEventDataAuthenticationMethodTypeAssisted      AccountUpdateEventDataAuthenticationMethodType = "assisted"
	AccountUpdateEventDataAuthenticationMethodTypeCredential    AccountUpdateEventDataAuthenticationMethodType = "credential"
	AccountUpdateEventDataAuthenticationMethodTypeAPIToken      AccountUpdateEventDataAuthenticationMethodType = "api_token"
	AccountUpdateEventDataAuthenticationMethodTypeAPICredential AccountUpdateEventDataAuthenticationMethodType = "api_credential"
	AccountUpdateEventDataAuthenticationMethodTypeOauth         AccountUpdateEventDataAuthenticationMethodType = "oauth"
)

type AccountUpdateEventEventType string

const (
	AccountUpdateEventEventTypeAccountUpdated AccountUpdateEventEventType = "account.updated"
)

type BaseWebhookEvent struct {
	// Unique Finch id of the employer account that was used to make this connection.
	AccountID string `json:"account_id,required"`
	// Unique Finch id of the company for which data has been updated.
	CompanyID string               `json:"company_id,required"`
	JSON      baseWebhookEventJSON `json:"-"`
}

// baseWebhookEventJSON contains the JSON metadata for the struct
// [BaseWebhookEvent]
type baseWebhookEventJSON struct {
	AccountID   apijson.Field
	CompanyID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaseWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyEvent struct {
	Data      map[string]interface{} `json:"data,nullable"`
	EventType CompanyEventEventType  `json:"event_type"`
	JSON      companyEventJSON       `json:"-"`
	BaseWebhookEvent
}

// companyEventJSON contains the JSON metadata for the struct [CompanyEvent]
type companyEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CompanyEvent) implementsWebhookEvent() {}

type CompanyEventEventType string

const (
	CompanyEventEventTypeCompanyUpdated CompanyEventEventType = "company.updated"
)

type DirectoryEvent struct {
	Data      DirectoryEventData      `json:"data"`
	EventType DirectoryEventEventType `json:"event_type"`
	JSON      directoryEventJSON      `json:"-"`
	BaseWebhookEvent
}

// directoryEventJSON contains the JSON metadata for the struct [DirectoryEvent]
type directoryEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DirectoryEvent) implementsWebhookEvent() {}

type DirectoryEventData struct {
	// The ID of the individual related to the event.
	IndividualID string                 `json:"individual_id"`
	JSON         directoryEventDataJSON `json:"-"`
}

// directoryEventDataJSON contains the JSON metadata for the struct
// [DirectoryEventData]
type directoryEventDataJSON struct {
	IndividualID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DirectoryEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryEventEventType string

const (
	DirectoryEventEventTypeDirectoryCreated DirectoryEventEventType = "directory.created"
	DirectoryEventEventTypeDirectoryUpdated DirectoryEventEventType = "directory.updated"
	DirectoryEventEventTypeDirectoryDeleted DirectoryEventEventType = "directory.deleted"
)

type EmploymentEvent struct {
	Data      EmploymentEventData      `json:"data"`
	EventType EmploymentEventEventType `json:"event_type"`
	JSON      employmentEventJSON      `json:"-"`
	BaseWebhookEvent
}

// employmentEventJSON contains the JSON metadata for the struct [EmploymentEvent]
type employmentEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmploymentEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r EmploymentEvent) implementsWebhookEvent() {}

type EmploymentEventData struct {
	// The ID of the individual related to the event.
	IndividualID string                  `json:"individual_id"`
	JSON         employmentEventDataJSON `json:"-"`
}

// employmentEventDataJSON contains the JSON metadata for the struct
// [EmploymentEventData]
type employmentEventDataJSON struct {
	IndividualID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmploymentEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmploymentEventEventType string

const (
	EmploymentEventEventTypeEmploymentCreated EmploymentEventEventType = "employment.created"
	EmploymentEventEventTypeEmploymentUpdated EmploymentEventEventType = "employment.updated"
	EmploymentEventEventTypeEmploymentDeleted EmploymentEventEventType = "employment.deleted"
)

type IndividualEvent struct {
	Data      IndividualEventData      `json:"data"`
	EventType IndividualEventEventType `json:"event_type"`
	JSON      individualEventJSON      `json:"-"`
	BaseWebhookEvent
}

// individualEventJSON contains the JSON metadata for the struct [IndividualEvent]
type individualEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IndividualEvent) implementsWebhookEvent() {}

type IndividualEventData struct {
	// The ID of the individual related to the event.
	IndividualID string                  `json:"individual_id"`
	JSON         individualEventDataJSON `json:"-"`
}

// individualEventDataJSON contains the JSON metadata for the struct
// [IndividualEventData]
type individualEventDataJSON struct {
	IndividualID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IndividualEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualEventEventType string

const (
	IndividualEventEventTypeIndividualCreated IndividualEventEventType = "individual.created"
	IndividualEventEventTypeIndividualUpdated IndividualEventEventType = "individual.updated"
	IndividualEventEventTypeIndividualDeleted IndividualEventEventType = "individual.deleted"
)

type JobCompletionEvent struct {
	Data      JobCompletionEventData      `json:"data"`
	EventType JobCompletionEventEventType `json:"event_type"`
	JSON      jobCompletionEventJSON      `json:"-"`
	BaseWebhookEvent
}

// jobCompletionEventJSON contains the JSON metadata for the struct
// [JobCompletionEvent]
type jobCompletionEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobCompletionEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r JobCompletionEvent) implementsWebhookEvent() {}

type JobCompletionEventData struct {
	// The id of the job which has completed.
	JobID string `json:"job_id,required"`
	// The url to query the result of the job.
	JobURL string                     `json:"job_url,required"`
	JSON   jobCompletionEventDataJSON `json:"-"`
}

// jobCompletionEventDataJSON contains the JSON metadata for the struct
// [JobCompletionEventData]
type jobCompletionEventDataJSON struct {
	JobID       apijson.Field
	JobURL      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobCompletionEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type JobCompletionEventEventType string

const (
	JobCompletionEventEventTypeJobBenefitCreateUpdated   JobCompletionEventEventType = "job.benefit_create.updated"
	JobCompletionEventEventTypeJobBenefitEnrollUpdated   JobCompletionEventEventType = "job.benefit_enroll.updated"
	JobCompletionEventEventTypeJobBenefitRegisterUpdated JobCompletionEventEventType = "job.benefit_register.updated"
	JobCompletionEventEventTypeJobBenefitUnenrollUpdated JobCompletionEventEventType = "job.benefit_unenroll.updated"
	JobCompletionEventEventTypeJobBenefitUpdateUpdated   JobCompletionEventEventType = "job.benefit_update.updated"
	JobCompletionEventEventTypeJobDataSyncAllUpdated     JobCompletionEventEventType = "job.data_sync_all.updated"
)

type PayStatementEvent struct {
	Data      PayStatementEventData      `json:"data"`
	EventType PayStatementEventEventType `json:"event_type"`
	JSON      payStatementEventJSON      `json:"-"`
	BaseWebhookEvent
}

// payStatementEventJSON contains the JSON metadata for the struct
// [PayStatementEvent]
type payStatementEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayStatementEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PayStatementEvent) implementsWebhookEvent() {}

type PayStatementEventData struct {
	// The ID of the individual associated with the pay statement.
	IndividualID string `json:"individual_id"`
	// The ID of the payment associated with the pay statement.
	PaymentID string                    `json:"payment_id"`
	JSON      payStatementEventDataJSON `json:"-"`
}

// payStatementEventDataJSON contains the JSON metadata for the struct
// [PayStatementEventData]
type payStatementEventDataJSON struct {
	IndividualID apijson.Field
	PaymentID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PayStatementEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PayStatementEventEventType string

const (
	PayStatementEventEventTypePayStatementCreated PayStatementEventEventType = "pay_statement.created"
	PayStatementEventEventTypePayStatementUpdated PayStatementEventEventType = "pay_statement.updated"
	PayStatementEventEventTypePayStatementDeleted PayStatementEventEventType = "pay_statement.deleted"
)

type PaymentEvent struct {
	Data      PaymentEventData      `json:"data"`
	EventType PaymentEventEventType `json:"event_type"`
	JSON      paymentEventJSON      `json:"-"`
	BaseWebhookEvent
}

// paymentEventJSON contains the JSON metadata for the struct [PaymentEvent]
type paymentEventJSON struct {
	Data        apijson.Field
	EventType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PaymentEvent) implementsWebhookEvent() {}

type PaymentEventData struct {
	// The date of the payment.
	PayDate string `json:"pay_date,required"`
	// The ID of the payment.
	PaymentID string               `json:"payment_id,required"`
	JSON      paymentEventDataJSON `json:"-"`
}

// paymentEventDataJSON contains the JSON metadata for the struct
// [PaymentEventData]
type paymentEventDataJSON struct {
	PayDate     apijson.Field
	PaymentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentEventEventType string

const (
	PaymentEventEventTypePaymentCreated PaymentEventEventType = "payment.created"
	PaymentEventEventTypePaymentUpdated PaymentEventEventType = "payment.updated"
	PaymentEventEventTypePaymentDeleted PaymentEventEventType = "payment.deleted"
)

// Union satisfied by [AccountUpdateEvent], [JobCompletionEvent], [CompanyEvent],
// [DirectoryEvent], [EmploymentEvent], [IndividualEvent], [PaymentEvent] or
// [PayStatementEvent].
type WebhookEvent interface {
	implementsWebhookEvent()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WebhookEvent)(nil)).Elem(), "")
}

type WebhookUnwrapParams struct {
}

type WebhookVerifySignatureParams struct {
}
