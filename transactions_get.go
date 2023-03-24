package gastrofix

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-gastrofix/utils"
)

func (c *Client) NewTransactionsGetRequest() TransactionsGetRequest {
	r := TransactionsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewTransactionsGetQueryParams()
	r.pathParams = r.NewTransactionsGetPathParams()
	r.requestBody = r.NewTransactionsGetRequestBody()
	return r
}

type TransactionsGetRequest struct {
	client      *Client
	queryParams *TransactionsGetQueryParams
	pathParams  *TransactionsGetPathParams
	method      string
	headers     http.Header
	requestBody TransactionsGetRequestBody
}

func (r TransactionsGetRequest) NewTransactionsGetQueryParams() *TransactionsGetQueryParams {
	return &TransactionsGetQueryParams{
		// Pagination: odata.NewPagination(),
	}
}

type TransactionsGetQueryParams struct {
}

func (p TransactionsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TransactionsGetRequest) QueryParams() *TransactionsGetQueryParams {
	return r.queryParams
}

func (r TransactionsGetRequest) NewTransactionsGetPathParams() *TransactionsGetPathParams {
	return &TransactionsGetPathParams{}
}

type TransactionsGetPathParams struct {
	PeriodID int `schema:"periodId,omitempty"`
}

func (p *TransactionsGetPathParams) Params() map[string]string {
	return map[string]string{
		"periodId": strconv.Itoa(p.PeriodID),
	}
}

func (r *TransactionsGetRequest) PathParams() *TransactionsGetPathParams {
	return r.pathParams
}

func (r *TransactionsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TransactionsGetRequest) Method() string {
	return r.method
}

func (r TransactionsGetRequest) NewTransactionsGetRequestBody() TransactionsGetRequestBody {
	return TransactionsGetRequestBody{}
}

type TransactionsGetRequestBody struct{}

func (r *TransactionsGetRequest) RequestBody() *TransactionsGetRequestBody {
	return &r.requestBody
}

func (r *TransactionsGetRequest) SetRequestBody(body TransactionsGetRequestBody) {
	r.requestBody = body
}

func (r *TransactionsGetRequest) NewResponseBody() *TransactionsGetResponseBody {
	return &TransactionsGetResponseBody{}
}

type TransactionsGetResponseBody struct {
	BaseCurrency     string       `json:"baseCurrency"`
	BusinessDay      string       `json:"businessDay"`
	CompanyID        int64        `json:"companyId"`
	CreatedTimestamp string       `json:"createdTimestamp"`
	LocationID       int64        `json:"locationId"`
	MasterData       interface{}  `json:"masterData"`
	PeriodID         int64        `json:"periodId"`
	SequenceNumber   int64        `json:"sequenceNumber"`
	TenantID         string       `json:"tenantId"`
	Transactions     Transactions `json:"transactions"`
	Type             string       `json:"type"`
	Version          string       `json:"version"`
}

func (r *TransactionsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/api/transaction/v3.0/transactions/{{.periodId}}", r.PathParams())
}

func (r *TransactionsGetRequest) Do() (TransactionsGetResponseBody, error) {
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
