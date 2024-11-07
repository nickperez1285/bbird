package api

import (
	"fmt"
)

type SportsArticle struct {
	Author	string	`json:"author"`
	Content	string	`json:"content"`
	PublishedDate	string	`json:"publishedDate"`
	Title	string	`json:"title"`
}

// Checks if all of the required fields for SportsArticle are set
// and validates all of the constraints for the object.
func (obj *SportsArticle) Validate() error {
	if obj == nil {
		return nil
	}
	fields := map[string]interface{}{
		"title": obj.Title,
		"author": obj.Author,
		"publishedDate": obj.PublishedDate,
		"content": obj.Content,
	}

	for field, value := range fields {
		if isEmpty := IsValEmpty(value); isEmpty{
			return fmt.Errorf("required field '%s' for object 'SportsArticle' is empty or unset", field)
		}
	}

	return nil
}

