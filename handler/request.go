package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is missing", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validade() error {
	if r == nil {
		return fmt.Errorf("request body está vazio")
	}

	if r.Role == "" {
		return errParamIsRequired(r.Role, "string")
	}
	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("Location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("Remote", "Boolean")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("Salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validade() error{
	if r.Role != "" || r.Company != "" || r.Link != "" || r.Location != ""|| r.Remote != nil || r.Link !="" || r.Salary > 0{
		return nil
	}

	return fmt.Errorf("at last one valid must be provided")
}
