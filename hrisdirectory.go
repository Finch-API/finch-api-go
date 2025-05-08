// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/shared"
)

// HRISDirectoryService contains methods and other services that help with
// interacting with the Finch API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHRISDirectoryService] method instead.
type HRISDirectoryService struct {
	Options []option.RequestOption
}

// NewHRISDirectoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHRISDirectoryService(opts ...option.RequestOption) (r *HRISDirectoryService) {
	r = &HRISDirectoryService{}
	r.Options = opts
	return
}

// Read company directory and organization structure
func (r *HRISDirectoryService) List(ctx context.Context, query HRISDirectoryListParams, opts ...option.RequestOption) (res *IndividualsPage, err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/directory"
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

// Read company directory and organization structure
func (r *HRISDirectoryService) ListAutoPaging(ctx context.Context, query HRISDirectoryListParams, opts ...option.RequestOption) *IndividualsPageAutoPager {
	return NewIndividualsPageAutoPager(r.List(ctx, query, opts...))
}

// Read company directory and organization structure
//
// Deprecated: use `List` instead
func (r *HRISDirectoryService) ListIndividuals(ctx context.Context, query HRISDirectoryListIndividualsParams, opts ...option.RequestOption) (res *IndividualsPage, err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "employer/directory"
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

// Read company directory and organization structure
//
// Deprecated: use `List` instead
func (r *HRISDirectoryService) ListIndividualsAutoPaging(ctx context.Context, query HRISDirectoryListIndividualsParams, opts ...option.RequestOption) *IndividualsPageAutoPager {
	return NewIndividualsPageAutoPager(r.ListIndividuals(ctx, query, opts...))
}

type IndividualsPage struct {
	// The array of employees.
	Individuals []IndividualInDirectory `json:"individuals"`
	Paging      shared.Paging           `json:"paging,required"`
	JSON        individualsPageJSON     `json:"-"`
	cfg         *requestconfig.RequestConfig
	res         *http.Response
}

// individualsPageJSON contains the JSON metadata for the struct [IndividualsPage]
type individualsPageJSON struct {
	Individuals apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualsPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *IndividualsPage) GetNextPage() (res *IndividualsPage, err error) {
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Paging.Offset

	if next < r.Paging.Count && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *IndividualsPage) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &IndividualsPage{}
	}
	r.cfg = cfg
	r.res = res
}

type IndividualsPageAutoPager struct {
	page *IndividualsPage
	cur  IndividualInDirectory
	idx  int
	run  int
	err  error
}

func NewIndividualsPageAutoPager(page *IndividualsPage, err error) *IndividualsPageAutoPager {
	return &IndividualsPageAutoPager{
		page: page,
		err:  err,
	}
}

func (r *IndividualsPageAutoPager) Next() bool {
	if r.page == nil || len(r.page.Individuals) == 0 {
		return false
	}
	if r.idx >= len(r.page.Individuals) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Individuals) == 0 {
			return false
		}
	}
	r.cur = r.page.Individuals[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *IndividualsPageAutoPager) Current() IndividualInDirectory {
	return r.cur
}

func (r *IndividualsPageAutoPager) Err() error {
	return r.err
}

func (r *IndividualsPageAutoPager) Index() int {
	return r.run
}

type IndividualInDirectory struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID string `json:"id,required" format:"uuid"`
	// The department object.
	Department IndividualInDirectoryDepartment `json:"department,required,nullable"`
	// The legal first name of the individual.
	FirstName string `json:"first_name,required,nullable"`
	// `true` if the individual is an active employee or contractor at the company.
	IsActive bool `json:"is_active,required,nullable"`
	// The legal last name of the individual.
	LastName string `json:"last_name,required,nullable"`
	// The manager object.
	Manager IndividualInDirectoryManager `json:"manager,required,nullable"`
	// The legal middle name of the individual.
	MiddleName string                    `json:"middle_name,required,nullable"`
	JSON       individualInDirectoryJSON `json:"-"`
}

// individualInDirectoryJSON contains the JSON metadata for the struct
// [IndividualInDirectory]
type individualInDirectoryJSON struct {
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

func (r *IndividualInDirectory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualInDirectoryJSON) RawJSON() string {
	return r.raw
}

// The department object.
type IndividualInDirectoryDepartment struct {
	// The name of the department.
	Name string                              `json:"name,nullable"`
	JSON individualInDirectoryDepartmentJSON `json:"-"`
}

// individualInDirectoryDepartmentJSON contains the JSON metadata for the struct
// [IndividualInDirectoryDepartment]
type individualInDirectoryDepartmentJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualInDirectoryDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualInDirectoryDepartmentJSON) RawJSON() string {
	return r.raw
}

// The manager object.
type IndividualInDirectoryManager struct {
	// A stable Finch `id` (UUID v4) for an individual in the company.
	ID   string                           `json:"id,required" format:"uuid"`
	JSON individualInDirectoryManagerJSON `json:"-"`
}

// individualInDirectoryManagerJSON contains the JSON metadata for the struct
// [IndividualInDirectoryManager]
type individualInDirectoryManagerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndividualInDirectoryManager) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r individualInDirectoryManagerJSON) RawJSON() string {
	return r.raw
}

type HRISDirectoryListParams struct {
	// Number of employees to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [HRISDirectoryListParams]'s query parameters as
// `url.Values`.
func (r HRISDirectoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HRISDirectoryListIndividualsParams struct {
	// Number of employees to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [HRISDirectoryListIndividualsParams]'s query parameters as
// `url.Values`.
func (r HRISDirectoryListIndividualsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
