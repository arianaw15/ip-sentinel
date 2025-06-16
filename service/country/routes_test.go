package country

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arianaw15/ip-sentinel/types"
	"github.com/gorilla/mux"
)

func TestValidateCountryByIP(t *testing.T) {
	countryStore := &MockCountryStore{}
	handler := NewHandler(countryStore)

	t.Run("GetCountryByIP, Success, Country has access", func(t *testing.T) {
		payload := types.IPRequestPayload{
			IP:             "1.178.64.0/23",
			ValidCountries: []string{"United States", "Jordan", "Peru"},
		}

		marshalledPayload, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodGet, "/country/validate", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/country/validate", handler.ValidateCountryByIP)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("GetCountryByIP, Success, Country does not have access", func(t *testing.T) {
		payload := types.IPRequestPayload{
			IP:             "1.178.64.0/23",
			ValidCountries: []string{"Jordan", "Peru", "Canada"},
		}
		marshalledPayload, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodGet, "/country/validate", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/country/validate", handler.ValidateCountryByIP)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusForbidden {
			t.Errorf("expected status code %d, got %d", http.StatusForbidden, rr.Code)
		}
	})

	t.Run("GetCountryByIP, Bad Request, Invalid IP, Ipv6", func(t *testing.T) {
		payload := types.IPRequestPayload{
			IP:             "2001:db8:a0b:12f0::1/32",
			ValidCountries: []string{"United States", "Jordan", "Peru"},
		}
		marshalledPayload, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodGet, "/country/validate", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/country/validate", handler.ValidateCountryByIP)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("GetCountryByIP, Bad Request, Invalid IP, not a valid ip type", func(t *testing.T) {
		payload := types.IPRequestPayload{
			IP:             "invalid-ip",
			ValidCountries: []string{"United States", "Jordan", "Peru"},
		}
		marshalledPayload, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodGet, "/country/validate", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/country/validate", handler.ValidateCountryByIP)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type MockCountryStore struct{}

func (m *MockCountryStore) GetCountryByIP(ip string) (*types.IPResponsePayload, error) {
	response := &types.IPResponsePayload{
		IP:          ip,
		GeoNameID:   "1.178.64.0/23",
		CountryName: "United States",
	}
	return response, nil
}

func (m *MockCountryStore) ValidateCountryAccess(unconfirmedCountry string, allowedCountryList *[]string) bool {
	if allowedCountryList == nil {
		return false
	}
	for _, country := range *allowedCountryList {
		if country == unconfirmedCountry {
			return true
		}
	}
	return false
}
