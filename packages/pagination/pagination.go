// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/Finch-API/finch-api-go/shared"
)

type SinglePage[T any] struct {
	Items []T            `json:"-,inline"`
	JSON  singlePageJSON `json:"-"`
	cfg   *requestconfig.RequestConfig
	res   *http.Response
}

// singlePageJSON contains the JSON metadata for the struct [SinglePage[T]]
type singlePageJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinglePage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singlePageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SinglePage[T]) GetNextPage() (res *SinglePage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
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

func (r *SinglePage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SinglePage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SinglePageAutoPager[T any] struct {
	page *SinglePage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSinglePageAutoPager[T any](page *SinglePage[T], err error) *SinglePageAutoPager[T] {
	return &SinglePageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SinglePageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SinglePageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SinglePageAutoPager[T]) Err() error {
	return r.err
}

func (r *SinglePageAutoPager[T]) Index() int {
	return r.run
}

type ResponsesPage[T any] struct {
	Responses []T               `json:"responses"`
	JSON      responsesPageJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// responsesPageJSON contains the JSON metadata for the struct [ResponsesPage[T]]
type responsesPageJSON struct {
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponsesPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responsesPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ResponsesPage[T]) GetNextPage() (res *ResponsesPage[T], err error) {
	if len(r.Responses) == 0 {
		return nil, nil
	}
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

func (r *ResponsesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ResponsesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ResponsesPageAutoPager[T any] struct {
	page *ResponsesPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewResponsesPageAutoPager[T any](page *ResponsesPage[T], err error) *ResponsesPageAutoPager[T] {
	return &ResponsesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ResponsesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Responses) == 0 {
		return false
	}
	if r.idx >= len(r.page.Responses) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Responses) == 0 {
			return false
		}
	}
	r.cur = r.page.Responses[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ResponsesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ResponsesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ResponsesPageAutoPager[T]) Index() int {
	return r.run
}

type Page[T any] struct {
	Data   []T           `json:"data"`
	Paging shared.Paging `json:"paging,required"`
	JSON   pageJSON      `json:"-"`
	cfg    *requestconfig.RequestConfig
	res    *http.Response
}

// pageJSON contains the JSON metadata for the struct [Page[T]]
type pageJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
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

func (r *Page[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Page[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}
