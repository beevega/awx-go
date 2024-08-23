package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type IAuth interface {
	SetAuthorizationHeader(req *http.Request)
}

// BasicAuth represents http basic auth.
type BasicAuth struct {
	Username string
	Password string
}

func (auth *BasicAuth) SetAuthorizationHeader(req *http.Request) {
	req.SetBasicAuth(auth.Username, auth.Password)
}

// TokenAuth
type TokenAuth struct {
	Token string
}

func (auth *TokenAuth) SetAuthorizationHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+auth.Token)
}

// APIRequest represents the http api communication way.
type APIRequest struct {
	Method   string
	Endpoint string
	Payload  interface{}
	Headers  http.Header
	Query    map[string]string
	Suffix   string
}

// SetHeader sets http header by passing k,v.
func (ar *APIRequest) SetHeader(key string, value string) *APIRequest {
	ar.Headers.Set(key, value)
	return ar
}

// NewAPIRequest news an APIRequest object.
func NewAPIRequest(method string, endpoint string, payload interface{}, query map[string]string) *APIRequest {
	var headers = http.Header{}
	var suffix string
	ar := &APIRequest{
		Method:   method,
		Endpoint: endpoint,
		Payload:  payload,
		Headers:  headers,
		Query:    query,
		Suffix:   suffix,
	}
	return ar
}

// Requester implemented a base http client.
// It supports do POST/GET via an human-readable way,
// in other word, all data is in `application/json` format.
// It also originally supports basic auth.
// For production usage, It would be better to wrapper
// an another rest client on this requester.
type Requester struct {
	Base   string
	Auth   IAuth
	Client *http.Client
}

// Get performs http get request.
func (r *Requester) Get(ctx context.Context, endpoint string, responseStruct any, query map[string]string) (*http.Response, error) {
	ar := NewAPIRequest(http.MethodGet, endpoint, nil, query)
	ar.Suffix = ""
	return r.Do(ctx, ar, responseStruct)
}

// Post performs http post request with json response.
func (r *Requester) Post(ctx context.Context, endpoint string, payload interface{}, responseStruct interface{}) (*http.Response, error) {
	ar := NewAPIRequest(http.MethodPost, endpoint, payload, map[string]string{})
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ctx, ar, &responseStruct)
}

// Put perform http PUT request with json response
func (r *Requester) Put(ctx context.Context, endpoint string, payload interface{}, responseStruct interface{}) (*http.Response, error) {
	ar := NewAPIRequest(http.MethodPut, endpoint, payload, map[string]string{})
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ctx, ar, &responseStruct)
}

// Patch perform http patch request with json response
func (r *Requester) Patch(ctx context.Context, endpoint string, payload interface{}, responseStruct interface{}) (*http.Response, error) {
	ar := NewAPIRequest(http.MethodPatch, endpoint, payload, map[string]string{})
	ar.SetHeader("Content-Type", "application/json")
	ar.Suffix = ""
	return r.Do(ctx, ar, &responseStruct)
}

// Delete performs http Delete request.
func (r *Requester) Delete(ctx context.Context, endpoint string) (*http.Response, error) {
	ar := NewAPIRequest(http.MethodDelete, endpoint, nil, map[string]string{})
	ar.Suffix = ""
	return r.Do(ctx, ar, nil)
}

// ValidateParams is to validate the input to use the services.
func ValidateParams(data map[string]interface{}, mandatoryFields []string) (notfound []string, status bool) {
	status = true
	for _, key := range mandatoryFields {
		_, exists := data[key]

		if !exists {
			notfound = append(notfound, key)
			status = false
		}
	}
	return notfound, status
}

// Do do the actual http request.
func (r *Requester) Do(ctx context.Context, ar *APIRequest, responseStruct interface{}) (*http.Response, error) {
	var body io.Reader

	if !strings.HasSuffix(ar.Endpoint, "/") && ar.Method != "POST" {
		ar.Endpoint += "/"
	}

	URL, err := url.Parse(r.Base + ar.Endpoint + ar.Suffix)
	if err != nil {
		return nil, err
	}

	if ar.Query != nil {
		querystring := make(url.Values)
		for key, val := range ar.Query {
			querystring.Set(key, val)
		}

		URL.RawQuery = querystring.Encode()
	}

	if ar.Payload != nil {
		rendered, err := json.Marshal(ar.Payload)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(rendered)
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, ar.Method, URL.String(), body)
	if err != nil {
		return nil, err
	}

	if r.Auth != nil {
		r.Auth.SetAuthorizationHeader(req)
	}

	for k := range ar.Headers {
		req.Header.Add(k, ar.Headers.Get(k))
	}

	response, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("response code %d, resp: %s", response.StatusCode, string(bodyBytes))
	}

	// В методе DELETE не возвращается тело ответа от сервера AWX
	// По этому необходимо проверить что мы ожидаем это тело получить
	if len(bodyBytes) > 0 && responseStruct != nil {
		if err := json.Unmarshal(bodyBytes, &responseStruct); err != nil {
			return nil, fmt.Errorf("error unmarshal: %v", err)
		}
	}

	return response, nil
}
