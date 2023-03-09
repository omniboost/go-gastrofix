package gastrofix

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-gastrofix/utils"
)

func (c *Client) NewBusinessPeriodsGetRequest() BusinessPeriodsGetRequest {
	r := BusinessPeriodsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewBusinessPeriodsGetQueryParams()
	r.pathParams = r.NewBusinessPeriodsGetPathParams()
	r.requestBody = r.NewBusinessPeriodsGetRequestBody()
	return r
}

type BusinessPeriodsGetRequest struct {
	client      *Client
	queryParams *BusinessPeriodsGetQueryParams
	pathParams  *BusinessPeriodsGetPathParams
	method      string
	headers     http.Header
	requestBody BusinessPeriodsGetRequestBody
}

func (r BusinessPeriodsGetRequest) NewBusinessPeriodsGetQueryParams() *BusinessPeriodsGetQueryParams {
	return &BusinessPeriodsGetQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type BusinessPeriodsGetQueryParams struct {
}

func (p BusinessPeriodsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *BusinessPeriodsGetRequest) QueryParams() *BusinessPeriodsGetQueryParams {
	return r.queryParams
}

func (r BusinessPeriodsGetRequest) NewBusinessPeriodsGetPathParams() *BusinessPeriodsGetPathParams {
	return &BusinessPeriodsGetPathParams{}
}

type BusinessPeriodsGetPathParams struct {
}

func (p *BusinessPeriodsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BusinessPeriodsGetRequest) PathParams() *BusinessPeriodsGetPathParams {
	return r.pathParams
}

func (r *BusinessPeriodsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BusinessPeriodsGetRequest) Method() string {
	return r.method
}

func (r BusinessPeriodsGetRequest) NewBusinessPeriodsGetRequestBody() BusinessPeriodsGetRequestBody {
	return BusinessPeriodsGetRequestBody{}
}

type BusinessPeriodsGetRequestBody struct{}

func (r *BusinessPeriodsGetRequest) RequestBody() *BusinessPeriodsGetRequestBody {
	return &r.requestBody
}

func (r *BusinessPeriodsGetRequest) SetRequestBody(body BusinessPeriodsGetRequestBody) {
	r.requestBody = body
}

func (r *BusinessPeriodsGetRequest) NewResponseBody() *BusinessPeriodsGetResponseBody {
	return &BusinessPeriodsGetResponseBody{}
}

type BusinessPeriodsGetResponseBody struct {
	BusinessPeriods []struct {
		BusinessDay           string    `json:"businessDay"`
		PeriodID              int       `json:"periodId"`
		StartPeriodTimestamp  time.Time `json:"startPeriodTimestamp"`
		FinishPeriodTimestamp time.Time `json:"finishPeriodTimestamp"`
	} `json:"businessPeriods"`
}

func (r *BusinessPeriodsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/transaction/v3.0/business_periods", r.PathParams())
}

func (r *BusinessPeriodsGetRequest) Do() (BusinessPeriodsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
