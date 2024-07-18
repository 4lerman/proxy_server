package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/4lerman/proxy_server/pkg/models"
)

func ProxyRequest(proxyReq models.ProxyRequest) (models.ProxyResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest(proxyReq.Method, proxyReq.URL, nil)
	if err != nil {
		return models.ProxyResponse{}, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range proxyReq.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.ProxyResponse{}, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ProxyResponse{}, fmt.Errorf("failed to response body: %v", err)
	}

	headers := make(map[string][]string)
	for key, values := range resp.Header {
		headers[key] = values
	}

	proxyRes := models.ProxyResponse{
		Status: resp.StatusCode,
		Headers: headers,
		Length: len(body),
	}

	return proxyRes, nil
}