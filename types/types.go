package types

type IPStore interface {
	GetCountryByIP(ip string) (*IPResponsePayload, error)
	ValidateCountryAccess(unconfirmedCountry string, allowedCountryList *[]string) bool
}

type IPRequestPayload struct {
	IP             string   `json:"ip" validate":"required,ip"`
	ValidCountries []string `json:"validCountries" validate:"required"`
}

type IPResponsePayload struct {
	IP          string `json:"ip"`
	GeoNameID   string `json:"geoNameId"`
	CountryName string `json:"countryName"`
}
