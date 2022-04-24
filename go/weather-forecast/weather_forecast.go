// Package weather provides tools to forecast
// weather condition of a given city.
package weather

// CurrentCondition represents the condition of a given city.
var CurrentCondition string

// CurrentLocation represents the name of a city.
var CurrentLocation string

// Forecast receive a city and condition arguments and
// returns the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
