// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/pagination"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISIndividualService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISIndividualService] method instead.
type HRISIndividualService struct {
	Options []option.RequestOption
}

// NewHRISIndividualService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISIndividualService(opts ...option.RequestOption) (r *HRISIndividualService) {
	r = &HRISIndividualService{}
	r.Options = opts
	return
}

// Read individual data, excluding income and employment data
func (r *HRISIndividualService) GetMany(ctx context.Context, body HRISIndividualGetManyParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[IndividualResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/individual"
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

// Read individual data, excluding income and employment data
func (r *HRISIndividualService) GetManyAutoPaging(ctx context.Context, body HRISIndividualGetManyParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[IndividualResponse] {
	return pagination.NewResponsesPageAutoPager(r.GetMany(ctx, body, opts...))
}

type Individual struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID     string            `json:"id"`
	Dob    string            `json:"dob,nullable"`
	Emails []IndividualEmail `json:"emails,nullable"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn,nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity IndividualEthnicity `json:"ethnicity,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The gender of the individual.
	Gender IndividualGender `json:"gender,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,nullable"`
	// The legal middle name of the individual.
	MiddleName   string                  `json:"middle_name,nullable"`
	PhoneNumbers []IndividualPhoneNumber `json:"phone_numbers,nullable"`
	// The preferred name of the individual.
	PreferredName string   `json:"preferred_name,nullable"`
	Residence     Location `json:"residence,nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	Ssn  string         `json:"ssn,nullable"`
	JSON individualJSON `json:"-"`
}

// individualJSON contains the JSON metadata for the struct [Individual]
type individualJSON struct {
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

func (r *Individual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualJSON) RawJSON() string {
	return r.raw
}

type IndividualEmail struct {
	Data string               `json:"data"`
	Type IndividualEmailsType `json:"type,nullable"`
	JSON individualEmailJSON  `json:"-"`
}

// individualEmailJSON contains the JSON metadata for the struct [IndividualEmail]
type individualEmailJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualEmailJSON) RawJSON() string {
	return r.raw
}

type IndividualEmailsType string

const (
	IndividualEmailsTypeWork     IndividualEmailsType = "work"
	IndividualEmailsTypePersonal IndividualEmailsType = "personal"
)

func (r IndividualEmailsType) IsKnown() bool {
	switch r {
	case IndividualEmailsTypeWork, IndividualEmailsTypePersonal:
		return true
	}
	return false
}

// The EEOC-defined ethnicity of the individual.
type IndividualEthnicity string

const (
	IndividualEthnicityAsian                           IndividualEthnicity = "asian"
	IndividualEthnicityWhite                           IndividualEthnicity = "white"
	IndividualEthnicityBlackOrAfricanAmerican          IndividualEthnicity = "black_or_african_american"
	IndividualEthnicityNativeHawaiianOrPacificIslander IndividualEthnicity = "native_hawaiian_or_pacific_islander"
	IndividualEthnicityAmericanIndianOrAlaskaNative    IndividualEthnicity = "american_indian_or_alaska_native"
	IndividualEthnicityHispanicOrLatino                IndividualEthnicity = "hispanic_or_latino"
	IndividualEthnicityTwoOrMoreRaces                  IndividualEthnicity = "two_or_more_races"
	IndividualEthnicityDeclineToSpecify                IndividualEthnicity = "decline_to_specify"
)

func (r IndividualEthnicity) IsKnown() bool {
	switch r {
	case IndividualEthnicityAsian, IndividualEthnicityWhite, IndividualEthnicityBlackOrAfricanAmerican, IndividualEthnicityNativeHawaiianOrPacificIslander, IndividualEthnicityAmericanIndianOrAlaskaNative, IndividualEthnicityHispanicOrLatino, IndividualEthnicityTwoOrMoreRaces, IndividualEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type IndividualGender string

const (
	IndividualGenderFemale           IndividualGender = "female"
	IndividualGenderMale             IndividualGender = "male"
	IndividualGenderOther            IndividualGender = "other"
	IndividualGenderDeclineToSpecify IndividualGender = "decline_to_specify"
)

func (r IndividualGender) IsKnown() bool {
	switch r {
	case IndividualGenderFemale, IndividualGenderMale, IndividualGenderOther, IndividualGenderDeclineToSpecify:
		return true
	}
	return false
}

type IndividualPhoneNumber struct {
	Data string                     `json:"data,nullable"`
	Type IndividualPhoneNumbersType `json:"type,nullable"`
	JSON individualPhoneNumberJSON  `json:"-"`
}

// individualPhoneNumberJSON contains the JSON metadata for the struct
// [IndividualPhoneNumber]
type individualPhoneNumberJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type IndividualPhoneNumbersType string

const (
	IndividualPhoneNumbersTypeWork     IndividualPhoneNumbersType = "work"
	IndividualPhoneNumbersTypePersonal IndividualPhoneNumbersType = "personal"
)

func (r IndividualPhoneNumbersType) IsKnown() bool {
	switch r {
	case IndividualPhoneNumbersTypeWork, IndividualPhoneNumbersTypePersonal:
		return true
	}
	return false
}

type IndividualResponse struct {
	Body         Individual             `json:"body"`
	Code         int64                  `json:"code"`
	IndividualID string                 `json:"individual_id"`
	JSON         individualResponseJSON `json:"-"`
}

// individualResponseJSON contains the JSON metadata for the struct
// [IndividualResponse]
type individualResponseJSON struct {
	Body         apijson.Field
	Code         apijson.Field
	IndividualID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IndividualResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualResponseJSON) RawJSON() string {
	return r.raw
}

type HRISIndividualGetManyParams struct {
	Options  param.Field[HRISIndividualGetManyParamsOptions]   `json:"options"`
	Requests param.Field[[]HRISIndividualGetManyParamsRequest] `json:"requests"`
}

func (r HRISIndividualGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISIndividualGetManyParamsOptions struct {
	Include param.Field[[]string] `json:"include"`
}

func (r HRISIndividualGetManyParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISIndividualGetManyParamsRequest struct {
	IndividualID param.Field[string] `json:"individual_id"`
}

func (r HRISIndividualGetManyParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
