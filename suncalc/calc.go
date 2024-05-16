package suncalc

import (
	"time"

	"github.com/nathan-osman/go-sunrise"
)

//теперь повилась функция, которую можно тестровать, но тут есть зависимость, которую нужно заменить как-то и тут помогут mock,
//но для этого нужно сделать интерфейс. И эта функция теперь зависит от этого интерфейса
func Calc(geocoder Geocoder, placeName string) (rise, set time.Time){
	lat, long, err := geocoder.GetCoordsByName(placeName)
	if err != nil {
		panic(err)
	}

	rise, set = sunrise.SunriseSunset(lat, long, 2000, time.January, 1)
	return
}

type Geocoder interface{
	GetCoordsByName(city string) (latitude, longitude float64, err error)
}
