package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Createweather

type CreateweatherRequest struct {
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
}

func (r *CreateweatherRequest) Validate() error {
	if r.Name == "" && r.Location == "" && r.Temperature > -100 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Temperature < -100 || r.Temperature > 100 {
		return errParamIsRequired("temperature", "float64")
	}
	return nil
}

// Updateweather

type UpdateweatherRequest struct {
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
}

func (r *UpdateweatherRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Name != "" || r.Location != "" || r.Temperature > -100 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
