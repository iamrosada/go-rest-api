package handler

import "fmt"

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s role (type: %s) is required", name, typ)
}

type CreateOportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOportunityRequest) Validate() error {

	if r.Role == "" {
		return errParamsIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamsIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamsIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamsIsRequired("link", "string")
	}

	if r.Remote == nil {
		return errParamsIsRequired("remote", "bool")

	}

	if r.Salary <= 0 {
		return errParamsIsRequired("salary", "int64")
	}
	return nil

}

//updated Oportunity

type UpdateOportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOportunityRequest) Validate() error {

	if r.Role != "" || r.Company != "" || r.Remote != nil || r.Location != "" || r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
