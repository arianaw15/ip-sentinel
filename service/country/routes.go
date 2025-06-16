package country

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"

	"github.com/arianaw15/ip-sentinel/types"
	"github.com/arianaw15/ip-sentinel/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.IPStore
}

func NewHandler(store types.IPStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/country/validate", h.ValidateCountryByIP).Methods("GET")
}

func (h *Handler) ValidateCountryByIP(w http.ResponseWriter, r *http.Request) {
	var payload types.IPRequestPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	ipAddress, _, err := net.ParseCIDR(payload.IP)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid CIDR format: %s", payload.IP))
		return
	}
	ip, err := netip.ParseAddr(ipAddress.String())
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid IP address: %s", payload.IP))
		return
	}

	if !ip.Is4() {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("IP address is not ipv4: %s", payload.IP))
		return
	}

	response, err := h.store.GetCountryByIP(payload.IP)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get country by IP: %w", err))
		return
	}

	if h.store.ValidateCountryAccess(response.CountryName, &payload.ValidCountries) {
		utils.WriteJSON(w, http.StatusOK, response)
	} else {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("access denied for country: %s", response.CountryName))
		return
	}

}
