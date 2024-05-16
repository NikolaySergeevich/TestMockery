package geocoder

import (
	"github.com/kelvins/geocoder"
)

type GoogleGeocoder struct {
	ApiKey string
}

func (gg *GoogleGeocoder) GetCoordsByName(city string) (latitude, longitude float64, err error) {
	geocoder.ApiKey = gg.ApiKey

	address := geocoder.Address{
		City: city,
	}

	location, err := geocoder.Geocoding(address)
	return location.Latitude, location.Longitude, err
}