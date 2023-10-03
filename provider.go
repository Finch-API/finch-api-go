// File generated from our OpenAPI spec by Stainless.

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

type BenefitSupportType struct {
	SupportedFeatures   BenefitSupportTypeSupportedFeatures   `json:"supported_features"`
	SupportedOperations BenefitSupportTypeSupportedOperations `json:"supported_operations"`
	JSON                benefitSupportTypeJSON
}

// benefitSupportTypeJSON contains the JSON metadata for the struct
// [BenefitSupportType]
type benefitSupportTypeJSON struct {
	SupportedFeatures   apijson.Field
	SupportedOperations apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BenefitSupportType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSupportTypeSupportedFeatures struct {
	// Whether the provider supports an annual maximum for this benefit.
	AnnualMaximum bool `json:"annual_maximum,nullable"`
	// Whether the provider supports catch up for this benefit. This field will only be
	// true for retirement benefits.
	CatchUp bool `json:"catch_up,nullable"`
	// Supported contribution types. An empty array indicates contributions are not
	// supported.
	CompanyContribution []BenefitSupportTypeSupportedFeaturesCompanyContribution `json:"company_contribution,nullable"`
	Description         string                                                   `json:"description,nullable"`
	// Supported deduction types. An empty array indicates deductions are not
	// supported.
	EmployeeDeduction []BenefitSupportTypeSupportedFeaturesEmployeeDeduction `json:"employee_deduction,nullable"`
	// The list of frequencies supported by the provider for this benefit
	Frequencies []BenefitFrequency `json:"frequencies"`
	// Whether the provider supports HSA contribution limits. Empty if this feature is
	// not supported for the benefit. This array only has values for HSA benefits.
	HsaContributionLimit []BenefitSupportTypeSupportedFeaturesHsaContributionLimit `json:"hsa_contribution_limit,nullable"`
	JSON                 benefitSupportTypeSupportedFeaturesJSON
}

// benefitSupportTypeSupportedFeaturesJSON contains the JSON metadata for the
// struct [BenefitSupportTypeSupportedFeatures]
type benefitSupportTypeSupportedFeaturesJSON struct {
	AnnualMaximum        apijson.Field
	CatchUp              apijson.Field
	CompanyContribution  apijson.Field
	Description          apijson.Field
	EmployeeDeduction    apijson.Field
	Frequencies          apijson.Field
	HsaContributionLimit apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BenefitSupportTypeSupportedFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSupportTypeSupportedFeaturesCompanyContribution string

const (
	BenefitSupportTypeSupportedFeaturesCompanyContributionFixed   BenefitSupportTypeSupportedFeaturesCompanyContribution = "fixed"
	BenefitSupportTypeSupportedFeaturesCompanyContributionPercent BenefitSupportTypeSupportedFeaturesCompanyContribution = "percent"
)

type BenefitSupportTypeSupportedFeaturesEmployeeDeduction string

const (
	BenefitSupportTypeSupportedFeaturesEmployeeDeductionFixed   BenefitSupportTypeSupportedFeaturesEmployeeDeduction = "fixed"
	BenefitSupportTypeSupportedFeaturesEmployeeDeductionPercent BenefitSupportTypeSupportedFeaturesEmployeeDeduction = "percent"
)

type BenefitSupportTypeSupportedFeaturesHsaContributionLimit string

const (
	BenefitSupportTypeSupportedFeaturesHsaContributionLimitIndividual BenefitSupportTypeSupportedFeaturesHsaContributionLimit = "individual"
	BenefitSupportTypeSupportedFeaturesHsaContributionLimitFamily     BenefitSupportTypeSupportedFeaturesHsaContributionLimit = "family"
)

type BenefitSupportTypeSupportedOperations struct {
	CompanyBenefits    BenefitSupportTypeSupportedOperationsCompanyBenefits    `json:"company_benefits"`
	IndividualBenefits BenefitSupportTypeSupportedOperationsIndividualBenefits `json:"individual_benefits"`
	JSON               benefitSupportTypeSupportedOperationsJSON
}

// benefitSupportTypeSupportedOperationsJSON contains the JSON metadata for the
// struct [BenefitSupportTypeSupportedOperations]
type benefitSupportTypeSupportedOperationsJSON struct {
	CompanyBenefits    apijson.Field
	IndividualBenefits apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BenefitSupportTypeSupportedOperations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BenefitSupportTypeSupportedOperationsCompanyBenefits struct {
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Create BenefitSupportTypeSupportedOperationsCompanyBenefitsCreate `json:"create"`
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Delete BenefitSupportTypeSupportedOperationsCompanyBenefitsDelete `json:"delete"`
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Read BenefitSupportTypeSupportedOperationsCompanyBenefitsRead `json:"read"`
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Update BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdate `json:"update"`
	JSON   benefitSupportTypeSupportedOperationsCompanyBenefitsJSON
}

// benefitSupportTypeSupportedOperationsCompanyBenefitsJSON contains the JSON
// metadata for the struct [BenefitSupportTypeSupportedOperationsCompanyBenefits]
type benefitSupportTypeSupportedOperationsCompanyBenefitsJSON struct {
	Create      apijson.Field
	Delete      apijson.Field
	Read        apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitSupportTypeSupportedOperationsCompanyBenefits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsCompanyBenefitsCreate string

const (
	BenefitSupportTypeSupportedOperationsCompanyBenefitsCreateSupported              BenefitSupportTypeSupportedOperationsCompanyBenefitsCreate = "supported"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsCreateNotSupportedByFinch    BenefitSupportTypeSupportedOperationsCompanyBenefitsCreate = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsCreateNotSupportedByProvider BenefitSupportTypeSupportedOperationsCompanyBenefitsCreate = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsCreateClientAccessOnly       BenefitSupportTypeSupportedOperationsCompanyBenefitsCreate = "client_access_only"
)

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsCompanyBenefitsDelete string

const (
	BenefitSupportTypeSupportedOperationsCompanyBenefitsDeleteSupported              BenefitSupportTypeSupportedOperationsCompanyBenefitsDelete = "supported"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsDeleteNotSupportedByFinch    BenefitSupportTypeSupportedOperationsCompanyBenefitsDelete = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsDeleteNotSupportedByProvider BenefitSupportTypeSupportedOperationsCompanyBenefitsDelete = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsDeleteClientAccessOnly       BenefitSupportTypeSupportedOperationsCompanyBenefitsDelete = "client_access_only"
)

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsCompanyBenefitsRead string

const (
	BenefitSupportTypeSupportedOperationsCompanyBenefitsReadSupported              BenefitSupportTypeSupportedOperationsCompanyBenefitsRead = "supported"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsReadNotSupportedByFinch    BenefitSupportTypeSupportedOperationsCompanyBenefitsRead = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsReadNotSupportedByProvider BenefitSupportTypeSupportedOperationsCompanyBenefitsRead = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsReadClientAccessOnly       BenefitSupportTypeSupportedOperationsCompanyBenefitsRead = "client_access_only"
)

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdate string

const (
	BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdateSupported              BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdate = "supported"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdateNotSupportedByFinch    BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdate = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdateNotSupportedByProvider BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdate = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdateClientAccessOnly       BenefitSupportTypeSupportedOperationsCompanyBenefitsUpdate = "client_access_only"
)

type BenefitSupportTypeSupportedOperationsIndividualBenefits struct {
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Create BenefitSupportTypeSupportedOperationsIndividualBenefitsCreate `json:"create"`
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Delete BenefitSupportTypeSupportedOperationsIndividualBenefitsDelete `json:"delete"`
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Read BenefitSupportTypeSupportedOperationsIndividualBenefitsRead `json:"read"`
	//   - `supported`: This operation is supported by both the provider and Finch <br>
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider <br>
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support <br>
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Update BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdate `json:"update"`
	JSON   benefitSupportTypeSupportedOperationsIndividualBenefitsJSON
}

// benefitSupportTypeSupportedOperationsIndividualBenefitsJSON contains the JSON
// metadata for the struct
// [BenefitSupportTypeSupportedOperationsIndividualBenefits]
type benefitSupportTypeSupportedOperationsIndividualBenefitsJSON struct {
	Create      apijson.Field
	Delete      apijson.Field
	Read        apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BenefitSupportTypeSupportedOperationsIndividualBenefits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsIndividualBenefitsCreate string

const (
	BenefitSupportTypeSupportedOperationsIndividualBenefitsCreateSupported              BenefitSupportTypeSupportedOperationsIndividualBenefitsCreate = "supported"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsCreateNotSupportedByFinch    BenefitSupportTypeSupportedOperationsIndividualBenefitsCreate = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsCreateNotSupportedByProvider BenefitSupportTypeSupportedOperationsIndividualBenefitsCreate = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsCreateClientAccessOnly       BenefitSupportTypeSupportedOperationsIndividualBenefitsCreate = "client_access_only"
)

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsIndividualBenefitsDelete string

const (
	BenefitSupportTypeSupportedOperationsIndividualBenefitsDeleteSupported              BenefitSupportTypeSupportedOperationsIndividualBenefitsDelete = "supported"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsDeleteNotSupportedByFinch    BenefitSupportTypeSupportedOperationsIndividualBenefitsDelete = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsDeleteNotSupportedByProvider BenefitSupportTypeSupportedOperationsIndividualBenefitsDelete = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsDeleteClientAccessOnly       BenefitSupportTypeSupportedOperationsIndividualBenefitsDelete = "client_access_only"
)

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsIndividualBenefitsRead string

const (
	BenefitSupportTypeSupportedOperationsIndividualBenefitsReadSupported              BenefitSupportTypeSupportedOperationsIndividualBenefitsRead = "supported"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsReadNotSupportedByFinch    BenefitSupportTypeSupportedOperationsIndividualBenefitsRead = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsReadNotSupportedByProvider BenefitSupportTypeSupportedOperationsIndividualBenefitsRead = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsReadClientAccessOnly       BenefitSupportTypeSupportedOperationsIndividualBenefitsRead = "client_access_only"
)

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdate string

const (
	BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdateSupported              BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdate = "supported"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdateNotSupportedByFinch    BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdate = "not_supported_by_finch"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdateNotSupportedByProvider BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdate = "not_supported_by_provider"
	BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdateClientAccessOnly       BenefitSupportTypeSupportedOperationsIndividualBenefitsUpdate = "client_access_only"
)

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
	Products []string `json:"products"`
	JSON     providerJSON
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

type ProviderAuthenticationMethod struct {
	// Each benefit type and their supported features. If the benefit type is not
	// supported, the property will be null
	BenefitsSupport ProviderAuthenticationMethodsBenefitsSupport `json:"benefits_support,nullable"`
	// The supported data fields returned by our HR and payroll endpoints
	SupportedFields ProviderAuthenticationMethodsSupportedFields `json:"supported_fields,nullable"`
	// The type of authentication method.
	Type ProviderAuthenticationMethodsType `json:"type"`
	JSON providerAuthenticationMethodJSON
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

// Each benefit type and their supported features. If the benefit type is not
// supported, the property will be null
type ProviderAuthenticationMethodsBenefitsSupport struct {
	Retirement401k     BenefitSupportType `json:"401k,nullable"`
	Retirement401kLoan BenefitSupportType `json:"401k_loan,nullable"`
	Retirement401kRoth BenefitSupportType `json:"401k_roth,nullable"`
	Retirement403B     BenefitSupportType `json:"403b,nullable"`
	Retirement403BRoth BenefitSupportType `json:"403b_roth,nullable"`
	Retirement457      BenefitSupportType `json:"457,nullable"`
	Retirement457Roth  BenefitSupportType `json:"457_roth,nullable"`
	Commuter           BenefitSupportType `json:"commuter,nullable"`
	CustomPostTax      BenefitSupportType `json:"custom_post_tax,nullable"`
	CustomPreTax       BenefitSupportType `json:"custom_pre_tax,nullable"`
	FsaDependentCare   BenefitSupportType `json:"fsa_dependent_care,nullable"`
	FsaMedical         BenefitSupportType `json:"fsa_medical,nullable"`
	HsaPost            BenefitSupportType `json:"hsa_post,nullable"`
	HsaPre             BenefitSupportType `json:"hsa_pre,nullable"`
	S125Dental         BenefitSupportType `json:"s125_dental,nullable"`
	S125Medical        BenefitSupportType `json:"s125_medical,nullable"`
	S125Vision         BenefitSupportType `json:"s125_vision,nullable"`
	Simple             BenefitSupportType `json:"simple,nullable"`
	SimpleIra          BenefitSupportType `json:"simple_ira,nullable"`
	JSON               providerAuthenticationMethodsBenefitsSupportJSON
}

// providerAuthenticationMethodsBenefitsSupportJSON contains the JSON metadata for
// the struct [ProviderAuthenticationMethodsBenefitsSupport]
type providerAuthenticationMethodsBenefitsSupportJSON struct {
	Retirement401k     apijson.Field
	Retirement401kLoan apijson.Field
	Retirement401kRoth apijson.Field
	Retirement403B     apijson.Field
	Retirement403BRoth apijson.Field
	Retirement457      apijson.Field
	Retirement457Roth  apijson.Field
	Commuter           apijson.Field
	CustomPostTax      apijson.Field
	CustomPreTax       apijson.Field
	FsaDependentCare   apijson.Field
	FsaMedical         apijson.Field
	HsaPost            apijson.Field
	HsaPre             apijson.Field
	S125Dental         apijson.Field
	S125Medical        apijson.Field
	S125Vision         apijson.Field
	Simple             apijson.Field
	SimpleIra          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsBenefitsSupport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The supported data fields returned by our HR and payroll endpoints
type ProviderAuthenticationMethodsSupportedFields struct {
	Company      ProviderAuthenticationMethodsSupportedFieldsCompany      `json:"company"`
	Directory    ProviderAuthenticationMethodsSupportedFieldsDirectory    `json:"directory"`
	Employment   ProviderAuthenticationMethodsSupportedFieldsEmployment   `json:"employment"`
	Individual   ProviderAuthenticationMethodsSupportedFieldsIndividual   `json:"individual"`
	PayStatement ProviderAuthenticationMethodsSupportedFieldsPayStatement `json:"pay_statement"`
	Payment      ProviderAuthenticationMethodsSupportedFieldsPayment      `json:"payment"`
	JSON         providerAuthenticationMethodsSupportedFieldsJSON
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
	JSON               providerAuthenticationMethodsSupportedFieldsCompanyJSON
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

type ProviderAuthenticationMethodsSupportedFieldsCompanyAccounts struct {
	AccountName     bool `json:"account_name"`
	AccountNumber   bool `json:"account_number"`
	AccountType     bool `json:"account_type"`
	InstitutionName bool `json:"institution_name"`
	RoutingNumber   bool `json:"routing_number"`
	JSON            providerAuthenticationMethodsSupportedFieldsCompanyAccountsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsCompanyDepartments struct {
	Name   bool                                                                 `json:"name"`
	Parent ProviderAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent `json:"parent"`
	JSON   providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsCompanyDepartmentsParent struct {
	Name bool `json:"name"`
	JSON providerAuthenticationMethodsSupportedFieldsCompanyDepartmentsParentJSON
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

type ProviderAuthenticationMethodsSupportedFieldsCompanyEntity struct {
	Subtype bool `json:"subtype"`
	Type    bool `json:"type"`
	JSON    providerAuthenticationMethodsSupportedFieldsCompanyEntityJSON
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

type ProviderAuthenticationMethodsSupportedFieldsCompanyLocations struct {
	City       bool `json:"city"`
	Country    bool `json:"country"`
	Line1      bool `json:"line1"`
	Line2      bool `json:"line2"`
	PostalCode bool `json:"postal_code"`
	State      bool `json:"state"`
	JSON       providerAuthenticationMethodsSupportedFieldsCompanyLocationsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsDirectory struct {
	Individuals ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividuals `json:"individuals"`
	Paging      ProviderAuthenticationMethodsSupportedFieldsDirectoryPaging      `json:"paging"`
	JSON        providerAuthenticationMethodsSupportedFieldsDirectoryJSON
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

type ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividuals struct {
	ID         bool                                                                    `json:"id"`
	Department interface{}                                                             `json:"department"`
	FirstName  bool                                                                    `json:"first_name"`
	IsActive   bool                                                                    `json:"is_active"`
	LastName   bool                                                                    `json:"last_name"`
	Manager    ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager `json:"manager"`
	MiddleName bool                                                                    `json:"middle_name"`
	JSON       providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsDirectoryIndividualsManager struct {
	ID   bool `json:"id"`
	JSON providerAuthenticationMethodsSupportedFieldsDirectoryIndividualsManagerJSON
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

type ProviderAuthenticationMethodsSupportedFieldsDirectoryPaging struct {
	Count  bool `json:"count"`
	Offset bool `json:"offset"`
	JSON   providerAuthenticationMethodsSupportedFieldsDirectoryPagingJSON
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
	Manager       interface{}                                                      `json:"manager"`
	MiddleName    bool                                                             `json:"middle_name"`
	StartDate     bool                                                             `json:"start_date"`
	Title         bool                                                             `json:"title"`
	JSON          providerAuthenticationMethodsSupportedFieldsEmploymentJSON
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

type ProviderAuthenticationMethodsSupportedFieldsEmploymentDepartment struct {
	Name bool `json:"name"`
	JSON providerAuthenticationMethodsSupportedFieldsEmploymentDepartmentJSON
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

type ProviderAuthenticationMethodsSupportedFieldsEmploymentEmployment struct {
	Subtype bool `json:"subtype"`
	Type    bool `json:"type"`
	JSON    providerAuthenticationMethodsSupportedFieldsEmploymentEmploymentJSON
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

type ProviderAuthenticationMethodsSupportedFieldsEmploymentIncome struct {
	Amount   bool `json:"amount"`
	Currency bool `json:"currency"`
	Unit     bool `json:"unit"`
	JSON     providerAuthenticationMethodsSupportedFieldsEmploymentIncomeJSON
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

type ProviderAuthenticationMethodsSupportedFieldsEmploymentLocation struct {
	City       bool `json:"city"`
	Country    bool `json:"country"`
	Line1      bool `json:"line1"`
	Line2      bool `json:"line2"`
	PostalCode bool `json:"postal_code"`
	State      bool `json:"state"`
	JSON       providerAuthenticationMethodsSupportedFieldsEmploymentLocationJSON
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

type ProviderAuthenticationMethodsSupportedFieldsIndividual struct {
	ID            bool                                                               `json:"id"`
	Dob           bool                                                               `json:"dob"`
	Emails        ProviderAuthenticationMethodsSupportedFieldsIndividualEmails       `json:"emails"`
	Ethnicity     bool                                                               `json:"ethnicity"`
	FirstName     bool                                                               `json:"first_name"`
	Gender        bool                                                               `json:"gender"`
	LastName      bool                                                               `json:"last_name"`
	MiddleName    bool                                                               `json:"middle_name"`
	PhoneNumbers  ProviderAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers `json:"phone_numbers"`
	PreferredName bool                                                               `json:"preferred_name"`
	Residence     ProviderAuthenticationMethodsSupportedFieldsIndividualResidence    `json:"residence"`
	Ssn           bool                                                               `json:"ssn"`
	JSON          providerAuthenticationMethodsSupportedFieldsIndividualJSON
}

// providerAuthenticationMethodsSupportedFieldsIndividualJSON contains the JSON
// metadata for the struct [ProviderAuthenticationMethodsSupportedFieldsIndividual]
type providerAuthenticationMethodsSupportedFieldsIndividualJSON struct {
	ID            apijson.Field
	Dob           apijson.Field
	Emails        apijson.Field
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

type ProviderAuthenticationMethodsSupportedFieldsIndividualEmails struct {
	Data bool `json:"data"`
	Type bool `json:"type"`
	JSON providerAuthenticationMethodsSupportedFieldsIndividualEmailsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsIndividualPhoneNumbers struct {
	Data bool `json:"data"`
	Type bool `json:"type"`
	JSON providerAuthenticationMethodsSupportedFieldsIndividualPhoneNumbersJSON
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

type ProviderAuthenticationMethodsSupportedFieldsIndividualResidence struct {
	City       bool `json:"city"`
	Country    bool `json:"country"`
	Line1      bool `json:"line1"`
	Line2      bool `json:"line2"`
	PostalCode bool `json:"postal_code"`
	State      bool `json:"state"`
	JSON       providerAuthenticationMethodsSupportedFieldsIndividualResidenceJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPayStatement struct {
	Paging        ProviderAuthenticationMethodsSupportedFieldsPayStatementPaging        `json:"paging"`
	PayStatements ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements `json:"pay_statements"`
	JSON          providerAuthenticationMethodsSupportedFieldsPayStatementJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPaging struct {
	Count  bool `json:"count,required"`
	Offset bool `json:"offset,required"`
	JSON   providerAuthenticationMethodsSupportedFieldsPayStatementPagingJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements struct {
	Earnings           ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings           `json:"earnings"`
	EmployeeDeductions ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions `json:"employee_deductions"`
	EmployerDeductions ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductions `json:"employer_deductions"`
	GrossPay           bool                                                                                    `json:"gross_pay"`
	IndividualID       bool                                                                                    `json:"individual_id"`
	NetPay             bool                                                                                    `json:"net_pay"`
	PaymentMethod      bool                                                                                    `json:"payment_method"`
	Taxes              ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes              `json:"taxes"`
	TotalHours         bool                                                                                    `json:"total_hours"`
	Type               bool                                                                                    `json:"type"`
	JSON               providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON
}

// providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON
// contains the JSON metadata for the struct
// [ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements]
type providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsJSON struct {
	Earnings           apijson.Field
	EmployeeDeductions apijson.Field
	EmployerDeductions apijson.Field
	GrossPay           apijson.Field
	IndividualID       apijson.Field
	NetPay             apijson.Field
	PaymentMethod      apijson.Field
	Taxes              apijson.Field
	TotalHours         apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarnings struct {
	Amount   bool `json:"amount"`
	Currency bool `json:"currency"`
	Name     bool `json:"name"`
	Type     bool `json:"type"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEarningsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductions struct {
	Amount   bool `json:"amount"`
	Currency bool `json:"currency"`
	Name     bool `json:"name"`
	PreTax   bool `json:"pre_tax"`
	Type     bool `json:"type"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployeeDeductionsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductions struct {
	Amount   bool `json:"amount"`
	Currency bool `json:"currency"`
	Name     bool `json:"name"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsEmployerDeductionsJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxes struct {
	Amount   bool `json:"amount"`
	Currency bool `json:"currency"`
	Employer bool `json:"employer"`
	Name     bool `json:"name"`
	Type     bool `json:"type"`
	JSON     providerAuthenticationMethodsSupportedFieldsPayStatementPayStatementsTaxesJSON
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
	JSON          providerAuthenticationMethodsSupportedFieldsPaymentJSON
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

type ProviderAuthenticationMethodsSupportedFieldsPaymentPayPeriod struct {
	EndDate   bool `json:"end_date"`
	StartDate bool `json:"start_date"`
	JSON      providerAuthenticationMethodsSupportedFieldsPaymentPayPeriodJSON
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

// The type of authentication method.
type ProviderAuthenticationMethodsType string

const (
	ProviderAuthenticationMethodsTypeAssisted   ProviderAuthenticationMethodsType = "assisted"
	ProviderAuthenticationMethodsTypeCredential ProviderAuthenticationMethodsType = "credential"
	ProviderAuthenticationMethodsTypeAPIToken   ProviderAuthenticationMethodsType = "api_token"
	ProviderAuthenticationMethodsTypeOauth      ProviderAuthenticationMethodsType = "oauth"
)
