package country

import (
	"database/sql"
	"fmt"

	"github.com/arianaw15/ip-sentinel/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetCountryByIP(ip string) (*types.IPResponsePayload, error) {
	// Get the GeoNameID for the given IP address
	query := "SELECT geoname_id FROM geo_lite_country_blocks WHERE network = ?"
	rows, err := s.db.Query(query, ip)
	if err != nil {
		return nil, err
	}

	response := new(types.IPResponsePayload)
	response.IP = ip
	for rows.Next() {
		err := rows.Scan(&response.GeoNameID)
		if err != nil {
			return nil, err
		}
		if response.GeoNameID == "" {
			return nil, fmt.Errorf("no geoNameID found for IP: %s", ip)
		}
	}

	// Get the country name using the GeoNameID
	countryNameQuery := "SELECT country_name FROM geo_lite_country_locations WHERE geoname_id = ?"
	countryRows, err := s.db.Query(countryNameQuery, response.GeoNameID)
	if err != nil {
		return nil, err
	}

	for countryRows.Next() {
		err := countryRows.Scan(&response.CountryName)
		if err != nil {
			return nil, err
		}
		if response.CountryName == "" {
			return nil, fmt.Errorf("no country name found for GeoNameID: %s", response.GeoNameID)
		}
	}

	return response, nil

}

func (s *Store) ValidateCountryAccess(countryRequest string, allowedCountryList *[]string) bool {
	// Check if the requested IP country is in the list of allowed countries
	hasAccess := false
	for _, country := range *allowedCountryList {
		if country == countryRequest {
			hasAccess = true
			break
		}
	}
	return hasAccess
}
