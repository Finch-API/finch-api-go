// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
	"github.com/tidwall/gjson"
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
	opts = slices.Concat(r.Options, opts)
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
	ID   string  `json:"id" format:"uuid"`
	Code float64 `json:"code"`
	Dob  string  `json:"dob,nullable"`
	// This field can have the runtime type of [[]IndividualObjectEmail].
	Emails interface{} `json:"emails"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn,nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity IndividualEthnicity `json:"ethnicity,nullable"`
	FinchCode string              `json:"finch_code"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The gender of the individual.
	Gender IndividualGender `json:"gender,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,nullable"`
	Message  string `json:"message"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name,nullable"`
	Name       string `json:"name"`
	// This field can have the runtime type of [[]IndividualObjectPhoneNumber].
	PhoneNumbers interface{} `json:"phone_numbers"`
	// The preferred name of the individual.
	PreferredName string   `json:"preferred_name,nullable"`
	Residence     Location `json:"residence,nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	// [Click here to learn more about enabling the SSN field](/developer-resources/Enable-SSN-Field).
	Ssn   string         `json:"ssn,nullable"`
	JSON  individualJSON `json:"-"`
	union IndividualUnion
}

// individualJSON contains the JSON metadata for the struct [Individual]
type individualJSON struct {
	ID            apijson.Field
	Code          apijson.Field
	Dob           apijson.Field
	Emails        apijson.Field
	EncryptedSsn  apijson.Field
	Ethnicity     apijson.Field
	FinchCode     apijson.Field
	FirstName     apijson.Field
	Gender        apijson.Field
	LastName      apijson.Field
	Message       apijson.Field
	MiddleName    apijson.Field
	Name          apijson.Field
	PhoneNumbers  apijson.Field
	PreferredName apijson.Field
	Residence     apijson.Field
	Ssn           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r individualJSON) RawJSON() string {
	return r.raw
}

func (r *Individual) UnmarshalJSON(data []byte) (err error) {
	*r = Individual{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IndividualUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [IndividualObject],
// [IndividualBatchError].
func (r Individual) AsUnion() IndividualUnion {
	return r.union
}

// Union satisfied by [IndividualObject] or [IndividualBatchError].
type IndividualUnion interface {
	implementsIndividual()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBatchError{}),
		},
	)
}

type IndividualObject struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID  string `json:"id,required" format:"uuid"`
	Dob string `json:"dob,required,nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity IndividualObjectEthnicity `json:"ethnicity,required,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,required,nullable"`
	// The gender of the individual.
	Gender IndividualObjectGender `json:"gender,required,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,required,nullable"`
	// The legal middle name of the individual.
	MiddleName   string                        `json:"middle_name,required,nullable"`
	PhoneNumbers []IndividualObjectPhoneNumber `json:"phone_numbers,required,nullable"`
	// The preferred name of the individual.
	PreferredName string                  `json:"preferred_name,required,nullable"`
	Residence     Location                `json:"residence,required,nullable"`
	Emails        []IndividualObjectEmail `json:"emails,nullable"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn,nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	// [Click here to learn more about enabling the SSN field](/developer-resources/Enable-SSN-Field).
	Ssn  string               `json:"ssn,nullable"`
	JSON individualObjectJSON `json:"-"`
}

// individualObjectJSON contains the JSON metadata for the struct
// [IndividualObject]
type individualObjectJSON struct {
	ID            apijson.Field
	Dob           apijson.Field
	Ethnicity     apijson.Field
	FirstName     apijson.Field
	Gender        apijson.Field
	LastName      apijson.Field
	MiddleName    apijson.Field
	PhoneNumbers  apijson.Field
	PreferredName apijson.Field
	Residence     apijson.Field
	Emails        apijson.Field
	EncryptedSsn  apijson.Field
	Ssn           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IndividualObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualObjectJSON) RawJSON() string {
	return r.raw
}

func (r IndividualObject) implementsIndividual() {}

// The EEOC-defined ethnicity of the individual.
type IndividualObjectEthnicity string

const (
	IndividualObjectEthnicityAsian                           IndividualObjectEthnicity = "asian"
	IndividualObjectEthnicityWhite                           IndividualObjectEthnicity = "white"
	IndividualObjectEthnicityBlackOrAfricanAmerican          IndividualObjectEthnicity = "black_or_african_american"
	IndividualObjectEthnicityNativeHawaiianOrPacificIslander IndividualObjectEthnicity = "native_hawaiian_or_pacific_islander"
	IndividualObjectEthnicityAmericanIndianOrAlaskaNative    IndividualObjectEthnicity = "american_indian_or_alaska_native"
	IndividualObjectEthnicityHispanicOrLatino                IndividualObjectEthnicity = "hispanic_or_latino"
	IndividualObjectEthnicityTwoOrMoreRaces                  IndividualObjectEthnicity = "two_or_more_races"
	IndividualObjectEthnicityDeclineToSpecify                IndividualObjectEthnicity = "decline_to_specify"
)

func (r IndividualObjectEthnicity) IsKnown() bool {
	switch r {
	case IndividualObjectEthnicityAsian, IndividualObjectEthnicityWhite, IndividualObjectEthnicityBlackOrAfricanAmerican, IndividualObjectEthnicityNativeHawaiianOrPacificIslander, IndividualObjectEthnicityAmericanIndianOrAlaskaNative, IndividualObjectEthnicityHispanicOrLatino, IndividualObjectEthnicityTwoOrMoreRaces, IndividualObjectEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type IndividualObjectGender string

const (
	IndividualObjectGenderFemale           IndividualObjectGender = "female"
	IndividualObjectGenderMale             IndividualObjectGender = "male"
	IndividualObjectGenderOther            IndividualObjectGender = "other"
	IndividualObjectGenderDeclineToSpecify IndividualObjectGender = "decline_to_specify"
)

func (r IndividualObjectGender) IsKnown() bool {
	switch r {
	case IndividualObjectGenderFemale, IndividualObjectGenderMale, IndividualObjectGenderOther, IndividualObjectGenderDeclineToSpecify:
		return true
	}
	return false
}

type IndividualObjectPhoneNumber struct {
	Data string                           `json:"data,required,nullable"`
	Type IndividualObjectPhoneNumbersType `json:"type,required,nullable"`
	JSON individualObjectPhoneNumberJSON  `json:"-"`
}

// individualObjectPhoneNumberJSON contains the JSON metadata for the struct
// [IndividualObjectPhoneNumber]
type individualObjectPhoneNumberJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualObjectPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualObjectPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type IndividualObjectPhoneNumbersType string

const (
	IndividualObjectPhoneNumbersTypeWork     IndividualObjectPhoneNumbersType = "work"
	IndividualObjectPhoneNumbersTypePersonal IndividualObjectPhoneNumbersType = "personal"
)

func (r IndividualObjectPhoneNumbersType) IsKnown() bool {
	switch r {
	case IndividualObjectPhoneNumbersTypeWork, IndividualObjectPhoneNumbersTypePersonal:
		return true
	}
	return false
}

type IndividualObjectEmail struct {
	Data string                     `json:"data,required"`
	Type IndividualObjectEmailsType `json:"type,required,nullable"`
	JSON individualObjectEmailJSON  `json:"-"`
}

// individualObjectEmailJSON contains the JSON metadata for the struct
// [IndividualObjectEmail]
type individualObjectEmailJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualObjectEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualObjectEmailJSON) RawJSON() string {
	return r.raw
}

type IndividualObjectEmailsType string

const (
	IndividualObjectEmailsTypeWork     IndividualObjectEmailsType = "work"
	IndividualObjectEmailsTypePersonal IndividualObjectEmailsType = "personal"
)

func (r IndividualObjectEmailsType) IsKnown() bool {
	switch r {
	case IndividualObjectEmailsTypeWork, IndividualObjectEmailsTypePersonal:
		return true
	}
	return false
}

type IndividualBatchError struct {
	Code      float64                  `json:"code,required"`
	Message   string                   `json:"message,required"`
	Name      string                   `json:"name,required"`
	FinchCode string                   `json:"finch_code"`
	JSON      individualBatchErrorJSON `json:"-"`
}

// individualBatchErrorJSON contains the JSON metadata for the struct
// [IndividualBatchError]
type individualBatchErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	FinchCode   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualBatchError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualBatchErrorJSON) RawJSON() string {
	return r.raw
}

func (r IndividualBatchError) implementsIndividual() {}

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

type IndividualResponse struct {
	Body         Individual             `json:"body,required"`
	Code         int64                  `json:"code,required"`
	IndividualID string                 `json:"individual_id,required"`
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
