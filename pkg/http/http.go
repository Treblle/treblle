package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PendingRequest struct {
	BaseURL string
	Headers map[string]string
}

type Response struct {
	*http.Response
}

// JSON decodes the response body into the provided interface
func (r *Response) JSON(v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func New() *PendingRequest {
	return &PendingRequest{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}
}

func (pr *PendingRequest) WithHeaders(headers map[string]string) *PendingRequest {
	for key, value := range headers {
		pr.Headers[key] = value
	}
	return pr
}

func (pr *PendingRequest) Base(url string) *PendingRequest {
	pr.BaseURL = url
	return pr
}

func (pr *PendingRequest) WithToken(token string, prefix ...string) *PendingRequest {
	authPrefix := "Bearer" // Default prefix

	// Override the prefix if provided
	if len(prefix) > 0 && prefix[0] != "" {
		authPrefix = prefix[0]
	}

	pr.Headers["Authorization"] = authPrefix + " " + token
	return pr
}

func (pr *PendingRequest) Get(endpoint string) (*Response, error) {
	resp, err := pr.Send("GET", endpoint, nil)
	return &Response{resp}, err
}

func (pr *PendingRequest) Post(endpoint string, body interface{}) (*Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := pr.Send("POST", endpoint, jsonBody)
	return &Response{resp}, err
}

func (pr *PendingRequest) Put(endpoint string, body interface{}, out interface{}) (*Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := pr.Send("PUT", endpoint, jsonBody)
	return &Response{resp}, err
}

func (pr *PendingRequest) Patch(endpoint string, body interface{}, out interface{}) (*Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := pr.Send("PATCH", endpoint, jsonBody)
	return &Response{resp}, err
}

func (pr *PendingRequest) Delete(endpoint string, out interface{}) (*Response, error) {
	resp, err := pr.Send("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return &Response{resp}, err
}

func (pr *PendingRequest) Send(method, endpoint string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, pr.BaseURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for key, value := range pr.Headers {
		req.Header.Add(key, value)
	}
	return client.Do(req)
}
