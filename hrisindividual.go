// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
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
	// This field can have the runtime type of [[]IndividualIndividualEmail].
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
	// This field can have the runtime type of [[]IndividualIndividualPhoneNumber].
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
// Possible runtime types of the union are [IndividualIndividual],
// [IndividualBatchError].
func (r Individual) AsUnion() IndividualUnion {
	return r.union
}

// Union satisfied by [IndividualIndividual] or [IndividualBatchError].
type IndividualUnion interface {
	implementsIndividual()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndividualUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualIndividual{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IndividualBatchError{}),
		},
	)
}

type IndividualIndividual struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID  string `json:"id" api:"required" format:"uuid"`
	Dob string `json:"dob" api:"required,nullable"`
	// The EEOC-defined ethnicity of the individual.
	Ethnicity IndividualIndividualEthnicity `json:"ethnicity" api:"required,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name" api:"required,nullable"`
	// The gender of the individual.
	Gender IndividualIndividualGender `json:"gender" api:"required,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name" api:"required,nullable"`
	// The legal middle name of the individual.
	MiddleName   string                            `json:"middle_name" api:"required,nullable"`
	PhoneNumbers []IndividualIndividualPhoneNumber `json:"phone_numbers" api:"required,nullable"`
	// The preferred name of the individual.
	PreferredName string                      `json:"preferred_name" api:"required,nullable"`
	Residence     Location                    `json:"residence" api:"required,nullable"`
	Emails        []IndividualIndividualEmail `json:"emails" api:"nullable"`
	// Social Security Number of the individual in **encrypted** format. This field is
	// only available with the `ssn` scope enabled and the
	// `options: { include: ['ssn'] }` param set in the body.
	EncryptedSsn string `json:"encrypted_ssn" api:"nullable"`
	// Social Security Number of the individual. This field is only available with the
	// `ssn` scope enabled and the `options: { include: ['ssn'] }` param set in the
	// body.
	// [Click here to learn more about enabling the SSN field](/developer-resources/Enable-SSN-Field).
	Ssn  string                   `json:"ssn" api:"nullable"`
	JSON individualIndividualJSON `json:"-"`
}

// individualIndividualJSON contains the JSON metadata for the struct
// [IndividualIndividual]
type individualIndividualJSON struct {
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

func (r *IndividualIndividual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualIndividualJSON) RawJSON() string {
	return r.raw
}

func (r IndividualIndividual) implementsIndividual() {}

// The EEOC-defined ethnicity of the individual.
type IndividualIndividualEthnicity string

const (
	IndividualIndividualEthnicityAsian                           IndividualIndividualEthnicity = "asian"
	IndividualIndividualEthnicityWhite                           IndividualIndividualEthnicity = "white"
	IndividualIndividualEthnicityBlackOrAfricanAmerican          IndividualIndividualEthnicity = "black_or_african_american"
	IndividualIndividualEthnicityNativeHawaiianOrPacificIslander IndividualIndividualEthnicity = "native_hawaiian_or_pacific_islander"
	IndividualIndividualEthnicityAmericanIndianOrAlaskaNative    IndividualIndividualEthnicity = "american_indian_or_alaska_native"
	IndividualIndividualEthnicityHispanicOrLatino                IndividualIndividualEthnicity = "hispanic_or_latino"
	IndividualIndividualEthnicityTwoOrMoreRaces                  IndividualIndividualEthnicity = "two_or_more_races"
	IndividualIndividualEthnicityDeclineToSpecify                IndividualIndividualEthnicity = "decline_to_specify"
)

func (r IndividualIndividualEthnicity) IsKnown() bool {
	switch r {
	case IndividualIndividualEthnicityAsian, IndividualIndividualEthnicityWhite, IndividualIndividualEthnicityBlackOrAfricanAmerican, IndividualIndividualEthnicityNativeHawaiianOrPacificIslander, IndividualIndividualEthnicityAmericanIndianOrAlaskaNative, IndividualIndividualEthnicityHispanicOrLatino, IndividualIndividualEthnicityTwoOrMoreRaces, IndividualIndividualEthnicityDeclineToSpecify:
		return true
	}
	return false
}

// The gender of the individual.
type IndividualIndividualGender string

const (
	IndividualIndividualGenderFemale           IndividualIndividualGender = "female"
	IndividualIndividualGenderMale             IndividualIndividualGender = "male"
	IndividualIndividualGenderOther            IndividualIndividualGender = "other"
	IndividualIndividualGenderDeclineToSpecify IndividualIndividualGender = "decline_to_specify"
)

func (r IndividualIndividualGender) IsKnown() bool {
	switch r {
	case IndividualIndividualGenderFemale, IndividualIndividualGenderMale, IndividualIndividualGenderOther, IndividualIndividualGenderDeclineToSpecify:
		return true
	}
	return false
}

type IndividualIndividualPhoneNumber struct {
	Data string                               `json:"data" api:"required,nullable"`
	Type IndividualIndividualPhoneNumbersType `json:"type" api:"required,nullable"`
	JSON individualIndividualPhoneNumberJSON  `json:"-"`
}

// individualIndividualPhoneNumberJSON contains the JSON metadata for the struct
// [IndividualIndividualPhoneNumber]
type individualIndividualPhoneNumberJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualIndividualPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualIndividualPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type IndividualIndividualPhoneNumbersType string

const (
	IndividualIndividualPhoneNumbersTypeWork     IndividualIndividualPhoneNumbersType = "work"
	IndividualIndividualPhoneNumbersTypePersonal IndividualIndividualPhoneNumbersType = "personal"
)

func (r IndividualIndividualPhoneNumbersType) IsKnown() bool {
	switch r {
	case IndividualIndividualPhoneNumbersTypeWork, IndividualIndividualPhoneNumbersTypePersonal:
		return true
	}
	return false
}

type IndividualIndividualEmail struct {
	Data string                         `json:"data" api:"required"`
	Type IndividualIndividualEmailsType `json:"type" api:"required,nullable"`
	JSON individualIndividualEmailJSON  `json:"-"`
}

// individualIndividualEmailJSON contains the JSON metadata for the struct
// [IndividualIndividualEmail]
type individualIndividualEmailJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualIndividualEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualIndividualEmailJSON) RawJSON() string {
	return r.raw
}

type IndividualIndividualEmailsType string

const (
	IndividualIndividualEmailsTypeWork     IndividualIndividualEmailsType = "work"
	IndividualIndividualEmailsTypePersonal IndividualIndividualEmailsType = "personal"
)

func (r IndividualIndividualEmailsType) IsKnown() bool {
	switch r {
	case IndividualIndividualEmailsTypeWork, IndividualIndividualEmailsTypePersonal:
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
