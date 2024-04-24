package rest

import (
	"encoding/json"
	"net/http"

	"github.com/itmo-education/delivery-backend/internal/domain"
)

func (s *Server) RegisterContract(resp http.ResponseWriter, req *http.Request) {
	var c domain.Contract
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		_, _ = resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.data.Save(c)
	if err != nil {
		_, _ = resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.WriteHeader(http.StatusCreated)
}
