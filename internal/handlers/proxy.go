package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/4lerman/proxy_server/internal/services"
	"github.com/4lerman/proxy_server/internal/store"
	"github.com/4lerman/proxy_server/pkg/models"
	"github.com/4lerman/proxy_server/pkg/validators"
	"github.com/google/uuid"
)

func HandlerProxy(w http.ResponseWriter, r *http.Request) {
	var proxyReq models.ProxyRequest

	if err := json.NewDecoder(r.Body).Decode(&proxyReq); err != nil {
		http.Error(w, "Invalid Bad Request", http.StatusBadRequest)
		return
	}

	if err := utils.ValidatorProxyRequest(proxyReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqID := uuid.New().String()
	store.SaveRequest(reqID, proxyReq)

	proxyRes, err := services.ProxyRequest(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proxyRes.ID = reqID
	store.SaveResponse(reqID, proxyRes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxyRes)
}
