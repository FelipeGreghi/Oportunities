package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("parameter %s (type : %s) is required", name, typ)
}

//CreateOpening

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if c.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if c.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if c.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if c.Salary == 0 {
		return errParamIsRequired("salary", "float64")
	}
	return nil
}
