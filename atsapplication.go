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

// ATSApplicationService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewATSApplicationService] method instead.
type ATSApplicationService struct {
	Options []option.RequestOption
}

// NewATSApplicationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewATSApplicationService(opts ...option.RequestOption) (r *ATSApplicationService) {
	r = &ATSApplicationService{}
	r.Options = opts
	return
}

// Gets an application from an organization.
func (r *ATSApplicationService) Get(ctx context.Context, applicationID string, opts ...option.RequestOption) (res *Application, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ats/applications/%s", applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all of an organization's applications.
func (r *ATSApplicationService) List(ctx context.Context, query ATSApplicationListParams, opts ...option.RequestOption) (res *ApplicationsPage, err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ats/applications"
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

// Gets all of an organization's applications.
func (r *ATSApplicationService) ListAutoPaging(ctx context.Context, query ATSApplicationListParams, opts ...option.RequestOption) *ApplicationsPageAutoPager {
	return NewApplicationsPageAutoPager(r.List(ctx, query, opts...))
}

type ApplicationsPage struct {
	Paging       Paging        `json:"paging,required"`
	Applications []Application `json:"applications,required"`
	JSON         applicationsPageJSON
	cfg          *requestconfig.RequestConfig
	res          *http.Response
}

// applicationsPageJSON contains the JSON metadata for the struct
// [ApplicationsPage]
type applicationsPageJSON struct {
	Paging       apijson.Field
	Applications apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *ApplicationsPage) GetNextPage() (res *ApplicationsPage, err error) {
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

func (r *ApplicationsPage) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type ApplicationsPageAutoPager struct {
	page *ApplicationsPage
	cur  Application
	idx  int
	run  int
	err  error
}

func NewApplicationsPageAutoPager(page *ApplicationsPage, err error) *ApplicationsPageAutoPager {
	return &ApplicationsPageAutoPager{
		page: page,
		err:  err,
	}
}

func (r *ApplicationsPageAutoPager) Next() bool {
	if r.page == nil || len(r.page.Applications) == 0 {
		return false
	}
	if r.idx >= len(r.page.Applications) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Applications[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ApplicationsPageAutoPager) Current() Application {
	return r.cur
}

func (r *ApplicationsPageAutoPager) Err() error {
	return r.err
}

func (r *ApplicationsPageAutoPager) Index() int {
	return r.run
}

type Application struct {
	ID             string                    `json:"id,required" format:"uuid"`
	CandidateID    string                    `json:"candidate_id,required" format:"uuid"`
	JobID          string                    `json:"job_id,required" format:"uuid"`
	OfferID        string                    `json:"offer_id,required,nullable" format:"uuid"`
	Stage          Stage                     `json:"stage,required,nullable"`
	RejectedAt     time.Time                 `json:"rejected_at,required,nullable" format:"date-time"`
	RejectedReason ApplicationRejectedReason `json:"rejected_reason,required,nullable"`
	JSON           applicationJSON
}

// applicationJSON contains the JSON metadata for the struct [Application]
type applicationJSON struct {
	ID             apijson.Field
	CandidateID    apijson.Field
	JobID          apijson.Field
	OfferID        apijson.Field
	Stage          apijson.Field
	RejectedAt     apijson.Field
	RejectedReason apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Application) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationRejectedReason struct {
	Text string `json:"text,nullable"`
	JSON applicationRejectedReasonJSON
}

// applicationRejectedReasonJSON contains the JSON metadata for the struct
// [ApplicationRejectedReason]
type applicationRejectedReasonJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRejectedReason) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ATSApplicationListParams struct {
	// Number of applications to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [ATSApplicationListParams]'s query parameters as
// `url.Values`.
func (r ATSApplicationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ATSApplicationListResponse struct {
	Paging       Paging        `json:"paging,required"`
	Applications []Application `json:"applications,required"`
	JSON         atsApplicationListResponseJSON
}

// atsApplicationListResponseJSON contains the JSON metadata for the struct
// [ATSApplicationListResponse]
type atsApplicationListResponseJSON struct {
	Paging       apijson.Field
	Applications apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ATSApplicationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
