package main

import (
	"Lesson/geocoder"
	"Lesson/suncalc"
	"fmt"
	"os"
	"time"
)

//В этом закомиченном main написана программа, котрую нужно протестировать. Тут есть некие зависимости, которые будут обращаться
//к API GOOGLE и это за деньги. естественно никто не будет платить деньги в тесте за обращение к стороннему сервису и поэтому
//нужно заменить это на mock. Но что бы заменить, нужно сразу переделеать код. Нужно отделить пользовательский интерфейс(stdin, stdout)
//от логики программы
//
//======================================================================================================================================

// func main() {

// 	placeName := os.Args[1]

// 	rise, set := Calc(placeName)

// 	fmt.Printf("Sunrise: %v; Sunset: %v\n", rise.Local().Format(time.TimeOnly), set.Local().Format(time.TimeOnly))
// }
// //теперь повилась функция, которую можно тестровать, но тут есть зависимость, которую нужно заменить как-то и тут помогут mock,
// //но для этого нужно сделать интерфейс - смотри в следующем блоке ==============/=================
// func Calc(placeName string) (rise, set time.Time){
// 	lat, long, err := getCoordsByName(placeName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	rise, set = sunrise.SunriseSunset(lat, long, 2000, time.January, 1)
// 	return
// }

// func getCoordsByName(city string) (latitude, longitude float64, err error) {
// 	geocoder.ApiKey = os.Getenv("GOOGLE_MAPS_API_KEY")

// 	address := geocoder.Address{
// 		City: city,
// 	}

// 	location, err := geocoder.Geocoding(address)
// 	return location.Latitude, location.Longitude, err
// }
//=============================================================================================================================
func main() {
	
	placeName := os.Args[1]
	g := &geocoder.GoogleGeocoder{ApiKey: os.Getenv("GOOGLE_MAPS_API_KEY")}
	rise, set := suncalc.Calc(g, placeName)

	fmt.Printf("Sunrise: %v; Sunset: %v\n", rise.Local().Format(time.TimeOnly), set.Local().Format(time.TimeOnly))
}


