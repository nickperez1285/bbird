package api

import (
	"fmt"
)

type WeatherResponse struct {
	Condition	string	`json:"condition"`
	Humidity	float64	`json:"humidity"`
	Location	string	`json:"location"`
	Temperature	float64	`json:"temperature"`
}

// Checks if all of the required fields for WeatherResponse are set
// and validates all of the constraints for the object.
func (obj *WeatherResponse) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"location": obj.Location,
		"temperature": obj.Temperature,
		"condition": obj.Condition,
		"humidity": obj.Humidity,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'WeatherResponse' is empty or unset", field)
		}
	}

	return nil
}

