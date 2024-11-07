package api

import (
	"fmt"
)

type SportsNewsResponse struct {
	Articles	[]SportsArticle	`json:"articles"`
}

// Checks if all of the required fields for SportsNewsResponse are set
// and validates all of the constraints for the object.
func (obj *SportsNewsResponse) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"articles": obj.Articles,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'SportsNewsResponse' is empty or unset", field)
		}
	}

	return nil
}

