// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/v2/internal/apijson"
	"github.com/Finch-API/finch-api-go/v2/internal/apiquery"
	"github.com/Finch-API/finch-api-go/v2/internal/param"
	"github.com/Finch-API/finch-api-go/v2/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/v2/option"
	"github.com/Finch-API/finch-api-go/v2/packages/pagination"
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
func (r *HRISIndividualService) GetMany(ctx context.Context, params HRISIndividualGetManyParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[IndividualResponse], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/individual"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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
func (r *HRISIndividualService) GetManyAutoPaging(ctx context.Context, params HRISIndividualGetManyParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[IndividualResponse] {
	return pagination.NewResponsesPageAutoPager(r.GetMany(ctx, params, opts...))
}

type Individual struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID   string  `json:"id" format:"uuid"`
	Code float64 `json:"code"`
	Dob  string  `json:"dob" api:"nullable"`
	// This field can have the runtime type of
	// [[]IndividualIndividualResponseBodyEmail].
	Emails interface{} `json:"emails"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn" api:"nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity IndividualEthnicity `json:"ethnicity" api:"nullable"`
	FinchCode string              `json:"finch_code"`
	// The legal first name of the individual.
	FirstName string `json:"first_name" api:"nullable"`
	// The gender of the individual.
	Gender IndividualGender `json:"gender" api:"nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name" api:"nullable"`
	Message  string `json:"message"`
	// The legal middle name of the individual.
	MiddleName string `json:"middle_name" api:"nullable"`
	Name       string `json:"name"`
	// This field can have the runtime type of
	// [[]IndividualIndividualResponseBodyPhoneNumber].
	PhoneNumbers interface{} `json:"phone_numbers"`
	// The preferred name of the individual.
	PreferredName string   `json:"preferred_name" api:"nullable"`
	Residence     Location `json:"residence" api:"nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	// [Click here to learn more about enabling the SSN field](/developer-resources/Enable-SSN-Field).
	Ssn   string         `json:"ssn" api:"nullable"`
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
// Possible runtime types of the union are [IndividualIndividualResponseBody],
// [IndividualBatchError].
func (r Individual) AsUnion() IndividualUnion {
	return r.union
}

// Union satisfied by [IndividualIndividualResponseBody] or [IndividualBatchError].
type IndividualUnion interface {
	implementsIndividual()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualIndividualResponseBody{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBatchError{}),
		},
	)
}

type IndividualIndividualResponseBody struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID  string `json:"id" api:"required" format:"uuid"`
	Dob string `json:"dob" api:"required,nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity IndividualIndividualResponseBodyEthnicity `json:"ethnicity" api:"required,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name" api:"required,nullable"`
	// The gender of the individual.
	Gender IndividualIndividualResponseBodyGender `json:"gender" api:"required,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name" api:"required,nullable"`
	// The legal middle name of the individual.
	MiddleName   string                                        `json:"middle_name" api:"required,nullable"`
	PhoneNumbers []IndividualIndividualResponseBodyPhoneNumber `json:"phone_numbers" api:"required,nullable"`
	// The preferred name of the individual.
	PreferredName string                                  `json:"preferred_name" api:"required,nullable"`
	Residence     Location                                `json:"residence" api:"required,nullable"`
	Emails        []IndividualIndividualResponseBodyEmail `json:"emails" api:"nullable"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn" api:"nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	// [Click here to learn more about enabling the SSN field](/developer-resources/Enable-SSN-Field).
	Ssn  string                               `json:"ssn" api:"nullable"`
	JSON individualIndividualResponseBodyJSON `json:"-"`
}

// individualIndividualResponseBodyJSON contains the JSON metadata for the struct
// [IndividualIndividualResponseBody]
type individualIndividualResponseBodyJSON struct {
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

func (r *IndividualIndividualResponseBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualIndividualResponseBodyJSON) RawJSON() string {
	return r.raw
}

func (r IndividualIndividualResponseBody) implementsIndividual() {}

// The EEOC-defined ethnicity of the individual.
type IndividualIndividualResponseBodyEthnicity string

const (
	IndividualIndividualResponseBodyEthnicityAsian                           IndividualIndividualResponseBodyEthnicity = "asian"
	IndividualIndividualResponseBodyEthnicityWhite                           IndividualIndividualResponseBodyEthnicity = "white"
	IndividualIndividualResponseBodyEthnicityBlackOrAfricanAmerican          IndividualIndividualResponseBodyEthnicity = "black_or_african_american"
	IndividualIndividualResponseBodyEthnicityNativeHawaiianOrPacificIslander IndividualIndividualResponseBodyEthnicity = "native_hawaiian_or_pacific_islander"
	IndividualIndividualResponseBodyEthnicityAmericanIndianOrAlaskaNative    IndividualIndividualResponseBodyEthnicity = "american_indian_or_alaska_native"
	IndividualIndividualResponseBodyEthnicityHispanicOrLatino                IndividualIndividualResponseBodyEthnicity = "hispanic_or_latino"
	IndividualIndividualResponseBodyEthnicityTwoOrMoreRaces                  IndividualIndividualResponseBodyEthnicity = "two_or_more_races"
	IndividualIndividualResponseBodyEthnicityDeclineToSpecify                IndividualIndividualResponseBodyEthnicity = "decline_to_specify"
)

func (r IndividualIndividualResponseBodyEthnicity) IsKnown() bool {
	switch r {
	case IndividualIndividualResponseBodyEthnicityAsian, IndividualIndividualResponseBodyEthnicityWhite, IndividualIndividualResponseBodyEthnicityBlackOrAfricanAmerican, IndividualIndividualResponseBodyEthnicityNativeHawaiianOrPacificIslander, IndividualIndividualResponseBodyEthnicityAmericanIndianOrAlaskaNative, IndividualIndividualResponseBodyEthnicityHispanicOrLatino, IndividualIndividualResponseBodyEthnicityTwoOrMoreRaces, IndividualIndividualResponseBodyEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type IndividualIndividualResponseBodyGender string

const (
	IndividualIndividualResponseBodyGenderFemale           IndividualIndividualResponseBodyGender = "female"
	IndividualIndividualResponseBodyGenderMale             IndividualIndividualResponseBodyGender = "male"
	IndividualIndividualResponseBodyGenderOther            IndividualIndividualResponseBodyGender = "other"
	IndividualIndividualResponseBodyGenderDeclineToSpecify IndividualIndividualResponseBodyGender = "decline_to_specify"
)

func (r IndividualIndividualResponseBodyGender) IsKnown() bool {
	switch r {
	case IndividualIndividualResponseBodyGenderFemale, IndividualIndividualResponseBodyGenderMale, IndividualIndividualResponseBodyGenderOther, IndividualIndividualResponseBodyGenderDeclineToSpecify:
		return true
	}
	return false
}

type IndividualIndividualResponseBodyPhoneNumber struct {
	Data string                                           `json:"data" api:"required,nullable"`
	Type IndividualIndividualResponseBodyPhoneNumbersType `json:"type" api:"required,nullable"`
	JSON individualIndividualResponseBodyPhoneNumberJSON  `json:"-"`
}

// individualIndividualResponseBodyPhoneNumberJSON contains the JSON metadata for
// the struct [IndividualIndividualResponseBodyPhoneNumber]
type individualIndividualResponseBodyPhoneNumberJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualIndividualResponseBodyPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualIndividualResponseBodyPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type IndividualIndividualResponseBodyPhoneNumbersType string

const (
	IndividualIndividualResponseBodyPhoneNumbersTypeWork     IndividualIndividualResponseBodyPhoneNumbersType = "work"
	IndividualIndividualResponseBodyPhoneNumbersTypePersonal IndividualIndividualResponseBodyPhoneNumbersType = "personal"
)

func (r IndividualIndividualResponseBodyPhoneNumbersType) IsKnown() bool {
	switch r {
	case IndividualIndividualResponseBodyPhoneNumbersTypeWork, IndividualIndividualResponseBodyPhoneNumbersTypePersonal:
		return true
	}
	return false
}

type IndividualIndividualResponseBodyEmail struct {
	Data string                                     `json:"data" api:"required"`
	Type IndividualIndividualResponseBodyEmailsType `json:"type" api:"required,nullable"`
	JSON individualIndividualResponseBodyEmailJSON  `json:"-"`
}

// individualIndividualResponseBodyEmailJSON contains the JSON metadata for the
// struct [IndividualIndividualResponseBodyEmail]
type individualIndividualResponseBodyEmailJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualIndividualResponseBodyEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualIndividualResponseBodyEmailJSON) RawJSON() string {
	return r.raw
}

type IndividualIndividualResponseBodyEmailsType string

const (
	IndividualIndividualResponseBodyEmailsTypeWork     IndividualIndividualResponseBodyEmailsType = "work"
	IndividualIndividualResponseBodyEmailsTypePersonal IndividualIndividualResponseBodyEmailsType = "personal"
)

func (r IndividualIndividualResponseBodyEmailsType) IsKnown() bool {
	switch r {
	case IndividualIndividualResponseBodyEmailsTypeWork, IndividualIndividualResponseBodyEmailsTypePersonal:
		return true
	}
	return false
}

type IndividualBatchError struct {
	Code      float64                  `json:"code" api:"required"`
	Message   string                   `json:"message" api:"required"`
	Name      string                   `json:"name" api:"required"`
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
	Body         Individual             `json:"body" api:"required"`
	Code         int64                  `json:"code" api:"required"`
	IndividualID string                 `json:"individual_id" api:"required"`
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
	// The entity IDs to specify which entities' data to access.
	EntityIDs param.Field[[]string]                             `query:"entity_ids" format:"uuid"`
	Options   param.Field[HRISIndividualGetManyParamsOptions]   `json:"options"`
	Requests  param.Field[[]HRISIndividualGetManyParamsRequest] `json:"requests"`
}

func (r HRISIndividualGetManyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISIndividualGetManyParams]'s query parameters as
// `url.Values`.
func (r HRISIndividualGetManyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
