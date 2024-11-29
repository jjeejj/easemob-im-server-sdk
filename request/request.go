package request

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	restyClient *resty.Client
}

func New(baseUrl string) *HttpClient {
	return &HttpClient{
		restyClient: resty.New().
			SetBaseURL(baseUrl).
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json"),
	}
}

func (h *HttpClient) Get(ctx context.Context, path string, queryParams map[string]string, result any) error {
	return nil
}

func (h *HttpClient) Post(ctx context.Context, path string, queryParams any, headers map[string]string, result any) (*resty.Response, error) {
	restyRequest := h.restyClient.R().SetContext(ctx).SetBody(queryParams).SetResult(result)
	if headers != nil {
		restyRequest.SetHeaders(headers)
	}
	return restyRequest.Post(path)
}
