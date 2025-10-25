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

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/packages/pagination"
)

// HRISCompanyPayStatementItemRuleService contains methods and other services that
// help with interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISCompanyPayStatementItemRuleService] method instead.
type HRISCompanyPayStatementItemRuleService struct {
	Options []option.RequestOption
}

// NewHRISCompanyPayStatementItemRuleService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHRISCompanyPayStatementItemRuleService(opts ...option.RequestOption) (r *HRISCompanyPayStatementItemRuleService) {
	r = &HRISCompanyPayStatementItemRuleService{}
	r.Options = opts
	return
}

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon Custom rules can be created to associate
// specific attributes to pay statement items depending on the use case. For
// example, pay statement items that meet certain conditions can be labeled as a
// pre-tax 401k. This metadata can be retrieved where pay statement item
// information is available.
func (r *HRISCompanyPayStatementItemRuleService) New(ctx context.Context, params HRISCompanyPayStatementItemRuleNewParams, opts ...option.RequestOption) (res *HRISCompanyPayStatementItemRuleNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "employer/pay-statement-item/rule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon Update a rule for a pay statement item.
func (r *HRISCompanyPayStatementItemRuleService) Update(ctx context.Context, ruleID string, params HRISCompanyPayStatementItemRuleUpdateParams, opts ...option.RequestOption) (res *HRISCompanyPayStatementItemRuleUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("employer/pay-statement-item/rule/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon List all rules of a connection account.
func (r *HRISCompanyPayStatementItemRuleService) List(ctx context.Context, query HRISCompanyPayStatementItemRuleListParams, opts ...option.RequestOption) (res *pagination.ResponsesPage[HRISCompanyPayStatementItemRuleListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon List all rules of a connection account.
func (r *HRISCompanyPayStatementItemRuleService) ListAutoPaging(ctx context.Context, query HRISCompanyPayStatementItemRuleListParams, opts ...option.RequestOption) *pagination.ResponsesPageAutoPager[HRISCompanyPayStatementItemRuleListResponse] {
	return pagination.NewResponsesPageAutoPager(r.List(ctx, query, opts...))
}

// **Beta:** this endpoint currently serves employers onboarded after March 4th and
// historical support will be added soon Delete a rule for a pay statement item.
func (r *HRISCompanyPayStatementItemRuleService) Delete(ctx context.Context, ruleID string, body HRISCompanyPayStatementItemRuleDeleteParams, opts ...option.RequestOption) (res *HRISCompanyPayStatementItemRuleDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("employer/pay-statement-item/rule/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type HRISCompanyPayStatementItemRuleNewResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISCompanyPayStatementItemRuleNewResponseAttributes  `json:"attributes"`
	Conditions []HRISCompanyPayStatementItemRuleNewResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date,nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date,nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISCompanyPayStatementItemRuleNewResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      hrisCompanyPayStatementItemRuleNewResponseJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleNewResponseJSON contains the JSON metadata for
// the struct [HRISCompanyPayStatementItemRuleNewResponse]
type hrisCompanyPayStatementItemRuleNewResponseJSON struct {
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

func (r *HRISCompanyPayStatementItemRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISCompanyPayStatementItemRuleNewResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                                   `json:"metadata"`
	JSON     hrisCompanyPayStatementItemRuleNewResponseAttributesJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleNewResponseAttributesJSON contains the JSON
// metadata for the struct [HRISCompanyPayStatementItemRuleNewResponseAttributes]
type hrisCompanyPayStatementItemRuleNewResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleNewResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleNewResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISCompanyPayStatementItemRuleNewResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISCompanyPayStatementItemRuleNewResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                                  `json:"value"`
	JSON  hrisCompanyPayStatementItemRuleNewResponseConditionJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleNewResponseConditionJSON contains the JSON
// metadata for the struct [HRISCompanyPayStatementItemRuleNewResponseCondition]
type hrisCompanyPayStatementItemRuleNewResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleNewResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleNewResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISCompanyPayStatementItemRuleNewResponseConditionsOperator string

const (
	HRISCompanyPayStatementItemRuleNewResponseConditionsOperatorEquals HRISCompanyPayStatementItemRuleNewResponseConditionsOperator = "equals"
)

func (r HRISCompanyPayStatementItemRuleNewResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleNewResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISCompanyPayStatementItemRuleNewResponseEntityType string

const (
	HRISCompanyPayStatementItemRuleNewResponseEntityTypePayStatementItem HRISCompanyPayStatementItemRuleNewResponseEntityType = "pay_statement_item"
)

func (r HRISCompanyPayStatementItemRuleNewResponseEntityType) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleNewResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISCompanyPayStatementItemRuleUpdateResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id" format:"uuid"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISCompanyPayStatementItemRuleUpdateResponseAttributes  `json:"attributes"`
	Conditions []HRISCompanyPayStatementItemRuleUpdateResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date,nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date,nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISCompanyPayStatementItemRuleUpdateResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      hrisCompanyPayStatementItemRuleUpdateResponseJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleUpdateResponseJSON contains the JSON metadata for
// the struct [HRISCompanyPayStatementItemRuleUpdateResponse]
type hrisCompanyPayStatementItemRuleUpdateResponseJSON struct {
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

func (r *HRISCompanyPayStatementItemRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISCompanyPayStatementItemRuleUpdateResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                                      `json:"metadata"`
	JSON     hrisCompanyPayStatementItemRuleUpdateResponseAttributesJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleUpdateResponseAttributesJSON contains the JSON
// metadata for the struct
// [HRISCompanyPayStatementItemRuleUpdateResponseAttributes]
type hrisCompanyPayStatementItemRuleUpdateResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleUpdateResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleUpdateResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISCompanyPayStatementItemRuleUpdateResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISCompanyPayStatementItemRuleUpdateResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                                     `json:"value"`
	JSON  hrisCompanyPayStatementItemRuleUpdateResponseConditionJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleUpdateResponseConditionJSON contains the JSON
// metadata for the struct [HRISCompanyPayStatementItemRuleUpdateResponseCondition]
type hrisCompanyPayStatementItemRuleUpdateResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleUpdateResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleUpdateResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISCompanyPayStatementItemRuleUpdateResponseConditionsOperator string

const (
	HRISCompanyPayStatementItemRuleUpdateResponseConditionsOperatorEquals HRISCompanyPayStatementItemRuleUpdateResponseConditionsOperator = "equals"
)

func (r HRISCompanyPayStatementItemRuleUpdateResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleUpdateResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISCompanyPayStatementItemRuleUpdateResponseEntityType string

const (
	HRISCompanyPayStatementItemRuleUpdateResponseEntityTypePayStatementItem HRISCompanyPayStatementItemRuleUpdateResponseEntityType = "pay_statement_item"
)

func (r HRISCompanyPayStatementItemRuleUpdateResponseEntityType) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleUpdateResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISCompanyPayStatementItemRuleListResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISCompanyPayStatementItemRuleListResponseAttributes  `json:"attributes"`
	Conditions []HRISCompanyPayStatementItemRuleListResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date,nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date,nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISCompanyPayStatementItemRuleListResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      hrisCompanyPayStatementItemRuleListResponseJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleListResponseJSON contains the JSON metadata for
// the struct [HRISCompanyPayStatementItemRuleListResponse]
type hrisCompanyPayStatementItemRuleListResponseJSON struct {
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

func (r *HRISCompanyPayStatementItemRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISCompanyPayStatementItemRuleListResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                                    `json:"metadata"`
	JSON     hrisCompanyPayStatementItemRuleListResponseAttributesJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleListResponseAttributesJSON contains the JSON
// metadata for the struct [HRISCompanyPayStatementItemRuleListResponseAttributes]
type hrisCompanyPayStatementItemRuleListResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleListResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleListResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISCompanyPayStatementItemRuleListResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISCompanyPayStatementItemRuleListResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                                   `json:"value"`
	JSON  hrisCompanyPayStatementItemRuleListResponseConditionJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleListResponseConditionJSON contains the JSON
// metadata for the struct [HRISCompanyPayStatementItemRuleListResponseCondition]
type hrisCompanyPayStatementItemRuleListResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleListResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleListResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISCompanyPayStatementItemRuleListResponseConditionsOperator string

const (
	HRISCompanyPayStatementItemRuleListResponseConditionsOperatorEquals HRISCompanyPayStatementItemRuleListResponseConditionsOperator = "equals"
)

func (r HRISCompanyPayStatementItemRuleListResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleListResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISCompanyPayStatementItemRuleListResponseEntityType string

const (
	HRISCompanyPayStatementItemRuleListResponseEntityTypePayStatementItem HRISCompanyPayStatementItemRuleListResponseEntityType = "pay_statement_item"
)

func (r HRISCompanyPayStatementItemRuleListResponseEntityType) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleListResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISCompanyPayStatementItemRuleDeleteResponse struct {
	// Finch id (uuidv4) for the rule.
	ID string `json:"id" format:"uuid"`
	// Specifies the fields to be applied when the condition is met.
	Attributes HRISCompanyPayStatementItemRuleDeleteResponseAttributes  `json:"attributes"`
	Conditions []HRISCompanyPayStatementItemRuleDeleteResponseCondition `json:"conditions"`
	// The datetime when the rule was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The datetime when the rule was deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate string `json:"effective_end_date,nullable"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate string `json:"effective_start_date,nullable"`
	// The entity type to which the rule is applied.
	EntityType HRISCompanyPayStatementItemRuleDeleteResponseEntityType `json:"entity_type"`
	// The priority of the rule.
	Priority int64 `json:"priority"`
	// The datetime when the rule was last updated.
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      hrisCompanyPayStatementItemRuleDeleteResponseJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleDeleteResponseJSON contains the JSON metadata for
// the struct [HRISCompanyPayStatementItemRuleDeleteResponse]
type hrisCompanyPayStatementItemRuleDeleteResponseJSON struct {
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

func (r *HRISCompanyPayStatementItemRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the fields to be applied when the condition is met.
type HRISCompanyPayStatementItemRuleDeleteResponseAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata map[string]interface{}                                      `json:"metadata"`
	JSON     hrisCompanyPayStatementItemRuleDeleteResponseAttributesJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleDeleteResponseAttributesJSON contains the JSON
// metadata for the struct
// [HRISCompanyPayStatementItemRuleDeleteResponseAttributes]
type hrisCompanyPayStatementItemRuleDeleteResponseAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleDeleteResponseAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleDeleteResponseAttributesJSON) RawJSON() string {
	return r.raw
}

type HRISCompanyPayStatementItemRuleDeleteResponseCondition struct {
	// The field to be checked in the rule.
	Field string `json:"field"`
	// The operator to be used in the rule.
	Operator HRISCompanyPayStatementItemRuleDeleteResponseConditionsOperator `json:"operator"`
	// The value of the field to be checked in the rule.
	Value string                                                     `json:"value"`
	JSON  hrisCompanyPayStatementItemRuleDeleteResponseConditionJSON `json:"-"`
}

// hrisCompanyPayStatementItemRuleDeleteResponseConditionJSON contains the JSON
// metadata for the struct [HRISCompanyPayStatementItemRuleDeleteResponseCondition]
type hrisCompanyPayStatementItemRuleDeleteResponseConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HRISCompanyPayStatementItemRuleDeleteResponseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hrisCompanyPayStatementItemRuleDeleteResponseConditionJSON) RawJSON() string {
	return r.raw
}

// The operator to be used in the rule.
type HRISCompanyPayStatementItemRuleDeleteResponseConditionsOperator string

const (
	HRISCompanyPayStatementItemRuleDeleteResponseConditionsOperatorEquals HRISCompanyPayStatementItemRuleDeleteResponseConditionsOperator = "equals"
)

func (r HRISCompanyPayStatementItemRuleDeleteResponseConditionsOperator) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleDeleteResponseConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISCompanyPayStatementItemRuleDeleteResponseEntityType string

const (
	HRISCompanyPayStatementItemRuleDeleteResponseEntityTypePayStatementItem HRISCompanyPayStatementItemRuleDeleteResponseEntityType = "pay_statement_item"
)

func (r HRISCompanyPayStatementItemRuleDeleteResponseEntityType) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleDeleteResponseEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISCompanyPayStatementItemRuleNewParams struct {
	// The entity IDs to create the rule for.
	EntityIDs param.Field[[]string] `query:"entity_ids,required" format:"uuid"`
	// Specifies the fields to be applied when the condition is met.
	Attributes param.Field[HRISCompanyPayStatementItemRuleNewParamsAttributes]  `json:"attributes"`
	Conditions param.Field[[]HRISCompanyPayStatementItemRuleNewParamsCondition] `json:"conditions"`
	// Specifies when the rules should stop applying rules based on the date.
	EffectiveEndDate param.Field[string] `json:"effective_end_date"`
	// Specifies when the rule should begin applying based on the date.
	EffectiveStartDate param.Field[string] `json:"effective_start_date"`
	// The entity type to which the rule is applied.
	EntityType param.Field[HRISCompanyPayStatementItemRuleNewParamsEntityType] `json:"entity_type"`
}

func (r HRISCompanyPayStatementItemRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISCompanyPayStatementItemRuleNewParams]'s query
// parameters as `url.Values`.
func (r HRISCompanyPayStatementItemRuleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the fields to be applied when the condition is met.
type HRISCompanyPayStatementItemRuleNewParamsAttributes struct {
	// The metadata to be attached in the entity. It is a key-value pairs where the
	// values can be of any type (string, number, boolean, object, array, etc.).
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
}

func (r HRISCompanyPayStatementItemRuleNewParamsAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HRISCompanyPayStatementItemRuleNewParamsCondition struct {
	// The field to be checked in the rule.
	Field param.Field[string] `json:"field"`
	// The operator to be used in the rule.
	Operator param.Field[HRISCompanyPayStatementItemRuleNewParamsConditionsOperator] `json:"operator"`
	// The value of the field to be checked in the rule.
	Value param.Field[string] `json:"value"`
}

func (r HRISCompanyPayStatementItemRuleNewParamsCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator to be used in the rule.
type HRISCompanyPayStatementItemRuleNewParamsConditionsOperator string

const (
	HRISCompanyPayStatementItemRuleNewParamsConditionsOperatorEquals HRISCompanyPayStatementItemRuleNewParamsConditionsOperator = "equals"
)

func (r HRISCompanyPayStatementItemRuleNewParamsConditionsOperator) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleNewParamsConditionsOperatorEquals:
		return true
	}
	return false
}

// The entity type to which the rule is applied.
type HRISCompanyPayStatementItemRuleNewParamsEntityType string

const (
	HRISCompanyPayStatementItemRuleNewParamsEntityTypePayStatementItem HRISCompanyPayStatementItemRuleNewParamsEntityType = "pay_statement_item"
)

func (r HRISCompanyPayStatementItemRuleNewParamsEntityType) IsKnown() bool {
	switch r {
	case HRISCompanyPayStatementItemRuleNewParamsEntityTypePayStatementItem:
		return true
	}
	return false
}

type HRISCompanyPayStatementItemRuleUpdateParams struct {
	// The entity IDs to update the rule for.
	EntityIDs        param.Field[[]string]    `query:"entity_ids,required" format:"uuid"`
	OptionalProperty param.Field[interface{}] `json:"optionalProperty"`
}

func (r HRISCompanyPayStatementItemRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [HRISCompanyPayStatementItemRuleUpdateParams]'s query
// parameters as `url.Values`.
func (r HRISCompanyPayStatementItemRuleUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISCompanyPayStatementItemRuleListParams struct {
	// The entity IDs to retrieve rules for.
	EntityIDs param.Field[[]string] `query:"entity_ids,required" format:"uuid"`
}

// URLQuery serializes [HRISCompanyPayStatementItemRuleListParams]'s query
// parameters as `url.Values`.
func (r HRISCompanyPayStatementItemRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISCompanyPayStatementItemRuleDeleteParams struct {
	// The entity IDs to delete the rule for.
	EntityIDs param.Field[[]string] `query:"entity_ids,required" format:"uuid"`
}

// URLQuery serializes [HRISCompanyPayStatementItemRuleDeleteParams]'s query
// parameters as `url.Values`.
func (r HRISCompanyPayStatementItemRuleDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
