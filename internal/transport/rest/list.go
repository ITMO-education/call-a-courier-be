package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/itmo-education/delivery-backend/internal/domain"
)

func (s *Server) ListContract(resp http.ResponseWriter, req *http.Request) {
	var c domain.ListContractRequest

	q := req.URL.Query()

	c.Limit, _ = strconv.Atoi(q.Get("limit"))
	c.Offset, _ = strconv.Atoi(q.Get("offset"))

	contracts, err := s.data.List(c)
	if err != nil {
		_, _ = resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(resp).Encode(contracts)
	if err != nil {
		_, _ = resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
}
