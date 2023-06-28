// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// ATSCandidateService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewATSCandidateService] method instead.
type ATSCandidateService struct {
	Options []option.RequestOption
}

// NewATSCandidateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewATSCandidateService(opts ...option.RequestOption) (r *ATSCandidateService) {
	r = &ATSCandidateService{}
	r.Options = opts
	return
}

// Gets a candidate from an organization. A candidate represents an individual
// associated with one or more applications.
func (r *ATSCandidateService) Get(ctx context.Context, candidateID string, opts ...option.RequestOption) (res *Candidate, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ats/candidates/%s", candidateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all of an organization's candidates. A candidate represents an individual
// associated with one or more applications.
func (r *ATSCandidateService) List(ctx context.Context, query ATSCandidateListParams, opts ...option.RequestOption) (res *CandidatesPage, err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ats/candidates"
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

// Gets all of an organization's candidates. A candidate represents an individual
// associated with one or more applications.
func (r *ATSCandidateService) ListAutoPaging(ctx context.Context, query ATSCandidateListParams, opts ...option.RequestOption) *CandidatesPageAutoPager {
	return NewCandidatesPageAutoPager(r.List(ctx, query, opts...))
}

type CandidatesPage struct {
	Candidates []Candidate `json:"candidates,required"`
	Paging     Paging      `json:"paging,required"`
	JSON       candidatesPageJSON
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// candidatesPageJSON contains the JSON metadata for the struct [CandidatesPage]
type candidatesPageJSON struct {
	Candidates  apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CandidatesPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *CandidatesPage) GetNextPage() (res *CandidatesPage, err error) {
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
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

func (r *CandidatesPage) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type CandidatesPageAutoPager struct {
	page *CandidatesPage
	cur  Candidate
	idx  int
	run  int
	err  error
}

func NewCandidatesPageAutoPager(page *CandidatesPage, err error) *CandidatesPageAutoPager {
	return &CandidatesPageAutoPager{
		page: page,
		err:  err,
	}
}

func (r *CandidatesPageAutoPager) Next() bool {
	if r.page == nil || len(r.page.Candidates) == 0 {
		return false
	}
	if r.idx >= len(r.page.Candidates) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Candidates[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CandidatesPageAutoPager) Current() Candidate {
	return r.cur
}

func (r *CandidatesPageAutoPager) Err() error {
	return r.err
}

func (r *CandidatesPageAutoPager) Index() int {
	return r.run
}

// A candidate represents an individual associated with one or more applications.
type Candidate struct {
	ID string `json:"id,required"`
	// Array of Finch uuids corresponding to `application`s for this individual
	ApplicationIDs []string                `json:"application_ids,required"`
	CreatedAt      time.Time               `json:"created_at,required" format:"date-time"`
	Emails         []CandidateEmails       `json:"emails,required"`
	FirstName      string                  `json:"first_name,required,nullable"`
	FullName       string                  `json:"full_name,required,nullable"`
	LastActivityAt time.Time               `json:"last_activity_at,required" format:"date-time"`
	LastName       string                  `json:"last_name,required,nullable"`
	PhoneNumbers   []CandidatePhoneNumbers `json:"phone_numbers,required"`
	JSON           candidateJSON
}

// candidateJSON contains the JSON metadata for the struct [Candidate]
type candidateJSON struct {
	ID             apijson.Field
	ApplicationIDs apijson.Field
	CreatedAt      apijson.Field
	Emails         apijson.Field
	FirstName      apijson.Field
	FullName       apijson.Field
	LastActivityAt apijson.Field
	LastName       apijson.Field
	PhoneNumbers   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Candidate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CandidateEmails struct {
	Data string `json:"data,nullable"`
	Type string `json:"type,nullable"`
	JSON candidateEmailsJSON
}

// candidateEmailsJSON contains the JSON metadata for the struct [CandidateEmails]
type candidateEmailsJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CandidateEmails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CandidatePhoneNumbers struct {
	Data string `json:"data,nullable"`
	Type string `json:"type,nullable"`
	JSON candidatePhoneNumbersJSON
}

// candidatePhoneNumbersJSON contains the JSON metadata for the struct
// [CandidatePhoneNumbers]
type candidatePhoneNumbersJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CandidatePhoneNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ATSCandidateListParams struct {
	// Number of candidates to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [ATSCandidateListParams]'s query parameters as `url.Values`.
func (r ATSCandidateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
