// Package weather provides weather forecasts for Goblinocus.
package weather

// CurrentCondition is used to detail the Current Condition.
var CurrentCondition string

// CurrentLocation is used to set the city.
var CurrentLocation string

// Forecast returns a string describing the current conditions in a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
