// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/internal/shared"
	"github.com/Finch-API/finch-api-go/option"
)

// HRISIndividualService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHRISIndividualService] method instead.
type HRISIndividualService struct {
	Options        []option.RequestOption
	EmploymentData *HRISIndividualEmploymentDataService
}

// NewHRISIndividualService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISIndividualService(opts ...option.RequestOption) (r *HRISIndividualService) {
	r = &HRISIndividualService{}
	r.Options = opts
	r.EmploymentData = NewHRISIndividualEmploymentDataService(opts...)
	return
}

// Read individual data, excluding income and employment data
func (r *HRISIndividualService) GetMany(ctx context.Context, body HRISIndividualGetManyParams, opts ...option.RequestOption) (res *shared.ResponsesPage[IndividualResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *HRISIndividualService) GetManyAutoPaging(ctx context.Context, body HRISIndividualGetManyParams, opts ...option.RequestOption) *shared.ResponsesPageAutoPager[IndividualResponse] {
	return shared.NewResponsesPageAutoPager(r.GetMany(ctx, body, opts...))
}

type Individual struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID     string             `json:"id"`
	Dob    string             `json:"dob,nullable"`
	Emails []IndividualEmails `json:"emails,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,nullable"`
	// The gender of the individual.
	Gender IndividualGender `json:"gender,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,nullable"`
	// The legal middle name of the individual.
	MiddleName   string                   `json:"middle_name,nullable"`
	PhoneNumbers []IndividualPhoneNumbers `json:"phone_numbers,nullable"`
	// The preferred name of the individual.
	PreferredName string   `json:"preferred_name,nullable"`
	Residence     Location `json:"residence,nullable"`
	// Note: This property is only available if enabled for your account. Please reach
	// out to your Finch representative if you would like access.
	Ssn  string `json:"ssn,nullable"`
	JSON individualJSON
}

// individualJSON contains the JSON metadata for the struct [Individual]
type individualJSON struct {
	ID            apijson.Field
	Dob           apijson.Field
	Emails        apijson.Field
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

type IndividualEmails struct {
	Data string               `json:"data"`
	Type IndividualEmailsType `json:"type"`
	JSON individualEmailsJSON
}

// individualEmailsJSON contains the JSON metadata for the struct
// [IndividualEmails]
type individualEmailsJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualEmails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualEmailsType string

const (
	IndividualEmailsTypeWork     IndividualEmailsType = "work"
	IndividualEmailsTypePersonal IndividualEmailsType = "personal"
)

// The gender of the individual.
type IndividualGender string

const (
	IndividualGenderFemale           IndividualGender = "female"
	IndividualGenderMale             IndividualGender = "male"
	IndividualGenderOther            IndividualGender = "other"
	IndividualGenderDeclineToSpecify IndividualGender = "decline_to_specify"
)

type IndividualPhoneNumbers struct {
	Data string                     `json:"data,nullable"`
	Type IndividualPhoneNumbersType `json:"type,nullable"`
	JSON individualPhoneNumbersJSON
}

// individualPhoneNumbersJSON contains the JSON metadata for the struct
// [IndividualPhoneNumbers]
type individualPhoneNumbersJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualPhoneNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IndividualPhoneNumbersType string

const (
	IndividualPhoneNumbersTypeWork     IndividualPhoneNumbersType = "work"
	IndividualPhoneNumbersTypePersonal IndividualPhoneNumbersType = "personal"
)

type IndividualResponse struct {
	Body         Individual `json:"body"`
	Code         int64      `json:"code"`
	IndividualID string     `json:"individual_id"`
	JSON         individualResponseJSON
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

type HRISIndividualGetManyParams struct {
	Options  param.Field[HRISIndividualGetManyParamsOptions]    `json:"options"`
	Requests param.Field[[]HRISIndividualGetManyParamsRequests] `json:"requests"`
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

type HRISIndividualGetManyParamsRequests struct {
	IndividualID param.Field[string] `json:"individual_id"`
}

func (r HRISIndividualGetManyParamsRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
