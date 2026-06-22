// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Finch-API/finch-api-go/v2/internal/apijson"
	"github.com/Finch-API/finch-api-go/v2/internal/apiquery"
	"github.com/Finch-API/finch-api-go/v2/internal/param"
	"github.com/Finch-API/finch-api-go/v2/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/v2/option"
	"github.com/Finch-API/finch-api-go/v2/packages/pagination"
)

// HRISPayStatementItemRuleService contains methods and other services that help
// with interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISPayStatementItemRuleService] method instead.
type HRISPayStatementItemRuleService struct {
	Options []option.RequestOption
}

// NewHRISPayStatementItemRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHRISPayStatementItemRuleService(opts ...option.RequestOption) (r *HRISPayStatementItemRuleService) {
	r = &HRISPayStatementItemRuleService{}
	r.Options = opts
	return
}

// Custom rules can be created to associate specific attributes to pay statement
// items depending on the use case. For example, pay statement items that meet
// certain conditions can be labeled as a pre-tax 401k. This metadata can be
// retrieved where pay statement item information is available.
func (r *HRISPayStatementItemRuleService) New(ctx context.Context, params HRISPayStatementItemRuleNewParams, opts ...option.RequestOption) (res *HRISPayStatementItemRuleNewResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	path := "employer/pay-statement-item/rule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a rule for a pay statement item.
func (r *HRISPayStatementItemRuleService) Update(ctx context.Context, ruleID string, params HRISPayStatementItemRuleUpdateParams, opts ...option.RequestOption) (res *HRISPayStatementItemRuleUpdateResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("employer/pay-statement-item/rule/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List all rules of a connection account.
func (r *HRISPayStatementItemRuleService) List(ctx context.Context, query HRISPayStatementItemRuleListParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[HRISPayStatementItemRuleListResponse], err error) {
	var raw *http.Response
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/pay-statement-item/rule"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List all rules of a connection account.
func (r *HRISPayStatementItemRuleService) ListAutoPaging(ctx context.Context, query HRISPayStatementItemRuleListParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[HRISPayStatementItemRuleListResponse] {
	return pagination.NewResponsesPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a rule for a pay statement item.
func (r *HRISPayStatementItemRuleService) Delete(ctx context.Context, ruleID string, body HRISPayStatementItemRuleDeleteParams, opts ...option.RequestOption) (res *HRISPayStatementItemRuleDeleteResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("employer/pay-statement-item/rule/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type HRISPayStatementItemRuleNewResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISPayStatementItemRuleNewResponseAttributes  `json:"attributes"`
	Conditions []HRISPayStatementItemRuleNewResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date" api:"nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date" api:"nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISPayStatementItemRuleNewResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      hrisPayStatementItemRuleNewResponseJSON `json:"-"`
}

// hrisPayStatementItemRuleNewResponseJSON contains the JSON metadata for the
// struct [HRISPayStatementItemRuleNewResponse]
type hrisPayStatementItemRuleNewResponseJSON struct {
	ID                 apijson.Field
	Attributes         apijson.Field
	Conditions         apijson.Field
	CreatedAt          apijson.Field
	EffectiveEndDate   apijson.Field
	EffectiveStartDate apijson.Field
	EntityType         apijson.Field
	Priority           apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISPayStatementItemRuleNewResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                            `json:"metadata"`
	JSON     hrisPayStatementItemRuleNewResponseAttributesJSON `json:"-"`
}

// hrisPayStatementItemRuleNewResponseAttributesJSON contains the JSON metadata for
// the struct [HRISPayStatementItemRuleNewResponseAttributes]
type hrisPayStatementItemRuleNewResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleNewResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleNewResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISPayStatementItemRuleNewResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISPayStatementItemRuleNewResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                           `json:"value"`
	JSON  hrisPayStatementItemRuleNewResponseConditionJSON `json:"-"`
}

// hrisPayStatementItemRuleNewResponseConditionJSON contains the JSON metadata for
// the struct [HRISPayStatementItemRuleNewResponseCondition]
type hrisPayStatementItemRuleNewResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleNewResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleNewResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISPayStatementItemRuleNewResponseConditionsOperator string

const (
	HRISPayStatementItemRuleNewResponseConditionsOperatorEquals HRISPayStatementItemRuleNewResponseConditionsOperator = "equals"
)

func (r HRISPayStatementItemRuleNewResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleNewResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISPayStatementItemRuleNewResponseEntityType string

const (
	HRISPayStatementItemRuleNewResponseEntityTypePayStatementItem HRISPayStatementItemRuleNewResponseEntityType = "pay_statement_item"
)

func (r HRISPayStatementItemRuleNewResponseEntityType) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleNewResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISPayStatementItemRuleUpdateResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id" format:"uuid"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISPayStatementItemRuleUpdateResponseAttributes  `json:"attributes"`
	Conditions []HRISPayStatementItemRuleUpdateResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date" api:"nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date" api:"nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISPayStatementItemRuleUpdateResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      hrisPayStatementItemRuleUpdateResponseJSON `json:"-"`
}

// hrisPayStatementItemRuleUpdateResponseJSON contains the JSON metadata for the
// struct [HRISPayStatementItemRuleUpdateResponse]
type hrisPayStatementItemRuleUpdateResponseJSON struct {
	ID                 apijson.Field
	Attributes         apijson.Field
	Conditions         apijson.Field
	CreatedAt          apijson.Field
	EffectiveEndDate   apijson.Field
	EffectiveStartDate apijson.Field
	EntityType         apijson.Field
	Priority           apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISPayStatementItemRuleUpdateResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                               `json:"metadata"`
	JSON     hrisPayStatementItemRuleUpdateResponseAttributesJSON `json:"-"`
}

// hrisPayStatementItemRuleUpdateResponseAttributesJSON contains the JSON metadata
// for the struct [HRISPayStatementItemRuleUpdateResponseAttributes]
type hrisPayStatementItemRuleUpdateResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleUpdateResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleUpdateResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISPayStatementItemRuleUpdateResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISPayStatementItemRuleUpdateResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                              `json:"value"`
	JSON  hrisPayStatementItemRuleUpdateResponseConditionJSON `json:"-"`
}

// hrisPayStatementItemRuleUpdateResponseConditionJSON contains the JSON metadata
// for the struct [HRISPayStatementItemRuleUpdateResponseCondition]
type hrisPayStatementItemRuleUpdateResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleUpdateResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleUpdateResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISPayStatementItemRuleUpdateResponseConditionsOperator string

const (
	HRISPayStatementItemRuleUpdateResponseConditionsOperatorEquals HRISPayStatementItemRuleUpdateResponseConditionsOperator = "equals"
)

func (r HRISPayStatementItemRuleUpdateResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleUpdateResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISPayStatementItemRuleUpdateResponseEntityType string

const (
	HRISPayStatementItemRuleUpdateResponseEntityTypePayStatementItem HRISPayStatementItemRuleUpdateResponseEntityType = "pay_statement_item"
)

func (r HRISPayStatementItemRuleUpdateResponseEntityType) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleUpdateResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISPayStatementItemRuleListResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISPayStatementItemRuleListResponseAttributes  `json:"attributes"`
	Conditions []HRISPayStatementItemRuleListResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date" api:"nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date" api:"nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISPayStatementItemRuleListResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      hrisPayStatementItemRuleListResponseJSON `json:"-"`
}

// hrisPayStatementItemRuleListResponseJSON contains the JSON metadata for the
// struct [HRISPayStatementItemRuleListResponse]
type hrisPayStatementItemRuleListResponseJSON struct {
	ID                 apijson.Field
	Attributes         apijson.Field
	Conditions         apijson.Field
	CreatedAt          apijson.Field
	EffectiveEndDate   apijson.Field
	EffectiveStartDate apijson.Field
	EntityType         apijson.Field
	Priority           apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISPayStatementItemRuleListResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                             `json:"metadata"`
	JSON     hrisPayStatementItemRuleListResponseAttributesJSON `json:"-"`
}

// hrisPayStatementItemRuleListResponseAttributesJSON contains the JSON metadata
// for the struct [HRISPayStatementItemRuleListResponseAttributes]
type hrisPayStatementItemRuleListResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleListResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleListResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISPayStatementItemRuleListResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISPayStatementItemRuleListResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                            `json:"value"`
	JSON  hrisPayStatementItemRuleListResponseConditionJSON `json:"-"`
}

// hrisPayStatementItemRuleListResponseConditionJSON contains the JSON metadata for
// the struct [HRISPayStatementItemRuleListResponseCondition]
type hrisPayStatementItemRuleListResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleListResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleListResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISPayStatementItemRuleListResponseConditionsOperator string

const (
	HRISPayStatementItemRuleListResponseConditionsOperatorEquals HRISPayStatementItemRuleListResponseConditionsOperator = "equals"
)

func (r HRISPayStatementItemRuleListResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleListResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISPayStatementItemRuleListResponseEntityType string

const (
	HRISPayStatementItemRuleListResponseEntityTypePayStatementItem HRISPayStatementItemRuleListResponseEntityType = "pay_statement_item"
)

func (r HRISPayStatementItemRuleListResponseEntityType) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleListResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISPayStatementItemRuleDeleteResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id" format:"uuid"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISPayStatementItemRuleDeleteResponseAttributes  `json:"attributes"`
	Conditions []HRISPayStatementItemRuleDeleteResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The datetime when the rule was deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date" api:"nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date" api:"nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISPayStatementItemRuleDeleteResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      hrisPayStatementItemRuleDeleteResponseJSON `json:"-"`
}

// hrisPayStatementItemRuleDeleteResponseJSON contains the JSON metadata for the
// struct [HRISPayStatementItemRuleDeleteResponse]
type hrisPayStatementItemRuleDeleteResponseJSON struct {
	ID                 apijson.Field
	Attributes         apijson.Field
	Conditions         apijson.Field
	CreatedAt          apijson.Field
	DeletedAt          apijson.Field
	EffectiveEndDate   apijson.Field
	EffectiveStartDate apijson.Field
	EntityType         apijson.Field
	Priority           apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISPayStatementItemRuleDeleteResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                               `json:"metadata"`
	JSON     hrisPayStatementItemRuleDeleteResponseAttributesJSON `json:"-"`
}

// hrisPayStatementItemRuleDeleteResponseAttributesJSON contains the JSON metadata
// for the struct [HRISPayStatementItemRuleDeleteResponseAttributes]
type hrisPayStatementItemRuleDeleteResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleDeleteResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleDeleteResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISPayStatementItemRuleDeleteResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISPayStatementItemRuleDeleteResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                              `json:"value"`
	JSON  hrisPayStatementItemRuleDeleteResponseConditionJSON `json:"-"`
}

// hrisPayStatementItemRuleDeleteResponseConditionJSON contains the JSON metadata
// for the struct [HRISPayStatementItemRuleDeleteResponseCondition]
type hrisPayStatementItemRuleDeleteResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISPayStatementItemRuleDeleteResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisPayStatementItemRuleDeleteResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISPayStatementItemRuleDeleteResponseConditionsOperator string

const (
	HRISPayStatementItemRuleDeleteResponseConditionsOperatorEquals HRISPayStatementItemRuleDeleteResponseConditionsOperator = "equals"
)

func (r HRISPayStatementItemRuleDeleteResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleDeleteResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISPayStatementItemRuleDeleteResponseEntityType string

const (
	HRISPayStatementItemRuleDeleteResponseEntityTypePayStatementItem HRISPayStatementItemRuleDeleteResponseEntityType = "pay_statement_item"
)

func (r HRISPayStatementItemRuleDeleteResponseEntityType) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleDeleteResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISPayStatementItemRuleNewParams struct {
	// The entity IDs to create the rule for.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
	// Specifies the fields to be applied when the condition is met.
	Attributes param.Field[HRISPayStatementItemRuleNewParamsAttributes]  `json:"attributes"`
	Conditions param.Field[[]HRISPayStatementItemRuleNewParamsCondition] `json:"conditions"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate param.Field[string] `json:"effective_end_date"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate param.Field[string] `json:"effective_start_date"`
	// The entity type to which the rule is applied.
	EntityType param.Field[HRISPayStatementItemRuleNewParamsEntityType] `json:"entity_type"`
}

func (r HRISPayStatementItemRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISPayStatementItemRuleNewParams]'s query parameters as
// `url.Values`.
func (r HRISPayStatementItemRuleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the fields to be applied when the condition is met.
type HRISPayStatementItemRuleNewParamsAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r HRISPayStatementItemRuleNewParamsAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISPayStatementItemRuleNewParamsCondition struct {
	// The field to be checked in the rule.
	Field param.Field[string] `json:"field"`
	// The operator to be used in the rule.
	Operator param.Field[HRISPayStatementItemRuleNewParamsConditionsOperator] `json:"operator"`
	// The value of the field to be checked in the rule.
	Value param.Field[string] `json:"value"`
}

func (r HRISPayStatementItemRuleNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator to be used in the rule.
type HRISPayStatementItemRuleNewParamsConditionsOperator string

const (
	HRISPayStatementItemRuleNewParamsConditionsOperatorEquals HRISPayStatementItemRuleNewParamsConditionsOperator = "equals"
)

func (r HRISPayStatementItemRuleNewParamsConditionsOperator) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleNewParamsConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISPayStatementItemRuleNewParamsEntityType string

const (
	HRISPayStatementItemRuleNewParamsEntityTypePayStatementItem HRISPayStatementItemRuleNewParamsEntityType = "pay_statement_item"
)

func (r HRISPayStatementItemRuleNewParamsEntityType) IsKnown() bool {
	switch r {
	case HRISPayStatementItemRuleNewParamsEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISPayStatementItemRuleUpdateParams struct {
	// The entity IDs to update the rule for.
	EntityIDs        param.Field[[]string]    `query:"entity_ids" format:"uuid"`
	OptionalProperty param.Field[interface{}] `json:"optionalProperty"`
}

func (r HRISPayStatementItemRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISPayStatementItemRuleUpdateParams]'s query parameters as
// `url.Values`.
func (r HRISPayStatementItemRuleUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISPayStatementItemRuleListParams struct {
	// The entity IDs to retrieve rules for.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
}

// URLQuery serializes [HRISPayStatementItemRuleListParams]'s query parameters as
// `url.Values`.
func (r HRISPayStatementItemRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISPayStatementItemRuleDeleteParams struct {
	// The entity IDs to delete the rule for.
	EntityIDs param.Field[[]string] `query:"entity_ids" format:"uuid"`
}

// URLQuery serializes [HRISPayStatementItemRuleDeleteParams]'s query parameters as
// `url.Values`.
func (r HRISPayStatementItemRuleDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
