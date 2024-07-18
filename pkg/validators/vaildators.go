package utils

import (
	"fmt"
	"github.com/4lerman/proxy_server/pkg/models"
)

func ValidatorProxyRequest(proxyReq models.ProxyRequest) error {
	if proxyReq.Method == "" {
		return fmt.Errorf("method is required")
	}

	if proxyReq.URL == "" {
		return fmt.Errorf("url is required")
	}

	return nil
}
