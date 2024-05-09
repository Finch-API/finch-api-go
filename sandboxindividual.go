// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// SandboxIndividualService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxIndividualService] method instead.
type SandboxIndividualService struct {
	Options []option.RequestOption
}

// NewSandboxIndividualService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxIndividualService(opts ...option.RequestOption) (r *SandboxIndividualService) {
	r = &SandboxIndividualService{}
	r.Options = opts
	return
}

// Update sandbox individual
func (r *SandboxIndividualService) Update(ctx context.Context, individualID string, body SandboxIndividualUpdateParams, opts ...option.RequestOption) (res *SandboxIndividualUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("sandbox/individual/%s", individualID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SandboxIndividualUpdateResponse struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID     string                                 `json:"id"`
	Dob    string                                 `json:"dob,nullable"`
	Emails []SandboxIndividualUpdateResponseEmail `json:"emails,nullable"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn,nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity SandboxIndividualUpdateResponseEthnicity `json:"ethnicity,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The gender of the individual.
	Gender SandboxIndividualUpdateResponseGender `json:"gender,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,nullable"`
	// The legal middle name of the individual.
	MiddleName   string                                       `json:"middle_name,nullable"`
	PhoneNumbers []SandboxIndividualUpdateResponsePhoneNumber `json:"phone_numbers,nullable"`
	// The preferred name of the individual.
	PreferredName string `json:"preferred_name,nullable"`
	Residence Location `json:"residence,nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	Ssn  string                              `json:"ssn,nullable"`
	JSON sandboxIndividualUpdateResponseJSON `json:"-"`
}

// sandboxIndividualUpdateResponseJSON contains the JSON metadata for the struct
// [SandboxIndividualUpdateResponse]
type sandboxIndividualUpdateResponseJSON struct {
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

func (r *SandboxIndividualUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxIndividualUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxIndividualUpdateResponseEmail struct {
	Data string                                    `json:"data"`
	Type SandboxIndividualUpdateResponseEmailsType `json:"type,nullable"`
	JSON sandboxIndividualUpdateResponseEmailJSON  `json:"-"`
}

// sandboxIndividualUpdateResponseEmailJSON contains the JSON metadata for the
// struct [SandboxIndividualUpdateResponseEmail]
type sandboxIndividualUpdateResponseEmailJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxIndividualUpdateResponseEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxIndividualUpdateResponseEmailJSON) RawJSON() string {
	return r.raw
}

type SandboxIndividualUpdateResponseEmailsType string

const (
	SandboxIndividualUpdateResponseEmailsTypeWork     SandboxIndividualUpdateResponseEmailsType = "work"
	SandboxIndividualUpdateResponseEmailsTypePersonal SandboxIndividualUpdateResponseEmailsType = "personal"
)

func (r SandboxIndividualUpdateResponseEmailsType) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateResponseEmailsTypeWork, SandboxIndividualUpdateResponseEmailsTypePersonal:
		return true
	}
	return false
}

// The EEOC-defined ethnicity of the individual.
type SandboxIndividualUpdateResponseEthnicity string

const (
	SandboxIndividualUpdateResponseEthnicityAsian                           SandboxIndividualUpdateResponseEthnicity = "asian"
	SandboxIndividualUpdateResponseEthnicityWhite                           SandboxIndividualUpdateResponseEthnicity = "white"
	SandboxIndividualUpdateResponseEthnicityBlackOrAfricanAmerican          SandboxIndividualUpdateResponseEthnicity = "black_or_african_american"
	SandboxIndividualUpdateResponseEthnicityNativeHawaiianOrPacificIslander SandboxIndividualUpdateResponseEthnicity = "native_hawaiian_or_pacific_islander"
	SandboxIndividualUpdateResponseEthnicityAmericanIndianOrAlaskaNative    SandboxIndividualUpdateResponseEthnicity = "american_indian_or_alaska_native"
	SandboxIndividualUpdateResponseEthnicityHispanicOrLatino                SandboxIndividualUpdateResponseEthnicity = "hispanic_or_latino"
	SandboxIndividualUpdateResponseEthnicityTwoOrMoreRaces                  SandboxIndividualUpdateResponseEthnicity = "two_or_more_races"
	SandboxIndividualUpdateResponseEthnicityDeclineToSpecify                SandboxIndividualUpdateResponseEthnicity = "decline_to_specify"
)

func (r SandboxIndividualUpdateResponseEthnicity) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateResponseEthnicityAsian, SandboxIndividualUpdateResponseEthnicityWhite, SandboxIndividualUpdateResponseEthnicityBlackOrAfricanAmerican, SandboxIndividualUpdateResponseEthnicityNativeHawaiianOrPacificIslander, SandboxIndividualUpdateResponseEthnicityAmericanIndianOrAlaskaNative, SandboxIndividualUpdateResponseEthnicityHispanicOrLatino, SandboxIndividualUpdateResponseEthnicityTwoOrMoreRaces, SandboxIndividualUpdateResponseEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type SandboxIndividualUpdateResponseGender string

const (
	SandboxIndividualUpdateResponseGenderFemale           SandboxIndividualUpdateResponseGender = "female"
	SandboxIndividualUpdateResponseGenderMale             SandboxIndividualUpdateResponseGender = "male"
	SandboxIndividualUpdateResponseGenderOther            SandboxIndividualUpdateResponseGender = "other"
	SandboxIndividualUpdateResponseGenderDeclineToSpecify SandboxIndividualUpdateResponseGender = "decline_to_specify"
)

func (r SandboxIndividualUpdateResponseGender) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateResponseGenderFemale, SandboxIndividualUpdateResponseGenderMale, SandboxIndividualUpdateResponseGenderOther, SandboxIndividualUpdateResponseGenderDeclineToSpecify:
		return true
	}
	return false
}

type SandboxIndividualUpdateResponsePhoneNumber struct {
	Data string                                          `json:"data"`
	Type SandboxIndividualUpdateResponsePhoneNumbersType `json:"type,nullable"`
	JSON sandboxIndividualUpdateResponsePhoneNumberJSON  `json:"-"`
}

// sandboxIndividualUpdateResponsePhoneNumberJSON contains the JSON metadata for
// the struct [SandboxIndividualUpdateResponsePhoneNumber]
type sandboxIndividualUpdateResponsePhoneNumberJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxIndividualUpdateResponsePhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxIndividualUpdateResponsePhoneNumberJSON) RawJSON() string {
	return r.raw
}

type SandboxIndividualUpdateResponsePhoneNumbersType string

const (
	SandboxIndividualUpdateResponsePhoneNumbersTypeWork     SandboxIndividualUpdateResponsePhoneNumbersType = "work"
	SandboxIndividualUpdateResponsePhoneNumbersTypePersonal SandboxIndividualUpdateResponsePhoneNumbersType = "personal"
)

func (r SandboxIndividualUpdateResponsePhoneNumbersType) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateResponsePhoneNumbersTypeWork, SandboxIndividualUpdateResponsePhoneNumbersTypePersonal:
		return true
	}
	return false
}

type SandboxIndividualUpdateParams struct {
	Dob    param.Field[string]                               `json:"dob"`
	Emails param.Field[[]SandboxIndividualUpdateParamsEmail] `json:"emails"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn param.Field[string] `json:"encrypted_ssn"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity param.Field[SandboxIndividualUpdateParamsEthnicity] `json:"ethnicity"`
	// The legal first name of the individual.
	FirstName param.Field[string] `json:"first_name"`
	// The gender of the individual.
	Gender param.Field[SandboxIndividualUpdateParamsGender] `json:"gender"`
	// The legal last name of the individual.
	LastName param.Field[string] `json:"last_name"`
	// The legal middle name of the individual.
	MiddleName   param.Field[string]                                     `json:"middle_name"`
	PhoneNumbers param.Field[[]SandboxIndividualUpdateParamsPhoneNumber] `json:"phone_numbers"`
	// The preferred name of the individual.
	PreferredName param.Field[string] `json:"preferred_name"`
	Residence param.Field[LocationParam] `json:"residence"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	Ssn param.Field[string] `json:"ssn"`
}

func (r SandboxIndividualUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxIndividualUpdateParamsEmail struct {
	Data param.Field[string]                                  `json:"data"`
	Type param.Field[SandboxIndividualUpdateParamsEmailsType] `json:"type"`
}

func (r SandboxIndividualUpdateParamsEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxIndividualUpdateParamsEmailsType string

const (
	SandboxIndividualUpdateParamsEmailsTypeWork     SandboxIndividualUpdateParamsEmailsType = "work"
	SandboxIndividualUpdateParamsEmailsTypePersonal SandboxIndividualUpdateParamsEmailsType = "personal"
)

func (r SandboxIndividualUpdateParamsEmailsType) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateParamsEmailsTypeWork, SandboxIndividualUpdateParamsEmailsTypePersonal:
		return true
	}
	return false
}

// The EEOC-defined ethnicity of the individual.
type SandboxIndividualUpdateParamsEthnicity string

const (
	SandboxIndividualUpdateParamsEthnicityAsian                           SandboxIndividualUpdateParamsEthnicity = "asian"
	SandboxIndividualUpdateParamsEthnicityWhite                           SandboxIndividualUpdateParamsEthnicity = "white"
	SandboxIndividualUpdateParamsEthnicityBlackOrAfricanAmerican          SandboxIndividualUpdateParamsEthnicity = "black_or_african_american"
	SandboxIndividualUpdateParamsEthnicityNativeHawaiianOrPacificIslander SandboxIndividualUpdateParamsEthnicity = "native_hawaiian_or_pacific_islander"
	SandboxIndividualUpdateParamsEthnicityAmericanIndianOrAlaskaNative    SandboxIndividualUpdateParamsEthnicity = "american_indian_or_alaska_native"
	SandboxIndividualUpdateParamsEthnicityHispanicOrLatino                SandboxIndividualUpdateParamsEthnicity = "hispanic_or_latino"
	SandboxIndividualUpdateParamsEthnicityTwoOrMoreRaces                  SandboxIndividualUpdateParamsEthnicity = "two_or_more_races"
	SandboxIndividualUpdateParamsEthnicityDeclineToSpecify                SandboxIndividualUpdateParamsEthnicity = "decline_to_specify"
)

func (r SandboxIndividualUpdateParamsEthnicity) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateParamsEthnicityAsian, SandboxIndividualUpdateParamsEthnicityWhite, SandboxIndividualUpdateParamsEthnicityBlackOrAfricanAmerican, SandboxIndividualUpdateParamsEthnicityNativeHawaiianOrPacificIslander, SandboxIndividualUpdateParamsEthnicityAmericanIndianOrAlaskaNative, SandboxIndividualUpdateParamsEthnicityHispanicOrLatino, SandboxIndividualUpdateParamsEthnicityTwoOrMoreRaces, SandboxIndividualUpdateParamsEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type SandboxIndividualUpdateParamsGender string

const (
	SandboxIndividualUpdateParamsGenderFemale           SandboxIndividualUpdateParamsGender = "female"
	SandboxIndividualUpdateParamsGenderMale             SandboxIndividualUpdateParamsGender = "male"
	SandboxIndividualUpdateParamsGenderOther            SandboxIndividualUpdateParamsGender = "other"
	SandboxIndividualUpdateParamsGenderDeclineToSpecify SandboxIndividualUpdateParamsGender = "decline_to_specify"
)

func (r SandboxIndividualUpdateParamsGender) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateParamsGenderFemale, SandboxIndividualUpdateParamsGenderMale, SandboxIndividualUpdateParamsGenderOther, SandboxIndividualUpdateParamsGenderDeclineToSpecify:
		return true
	}
	return false
}

type SandboxIndividualUpdateParamsPhoneNumber struct {
	Data param.Field[string]                                        `json:"data"`
	Type param.Field[SandboxIndividualUpdateParamsPhoneNumbersType] `json:"type"`
}

func (r SandboxIndividualUpdateParamsPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SandboxIndividualUpdateParamsPhoneNumbersType string

const (
	SandboxIndividualUpdateParamsPhoneNumbersTypeWork     SandboxIndividualUpdateParamsPhoneNumbersType = "work"
	SandboxIndividualUpdateParamsPhoneNumbersTypePersonal SandboxIndividualUpdateParamsPhoneNumbersType = "personal"
)

func (r SandboxIndividualUpdateParamsPhoneNumbersType) IsKnown() bool {
	switch r {
	case SandboxIndividualUpdateParamsPhoneNumbersTypeWork, SandboxIndividualUpdateParamsPhoneNumbersTypePersonal:
		return true
	}
	return false
}
