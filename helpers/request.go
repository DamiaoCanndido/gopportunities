package helpers

import "fmt"

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `gorm:"type:text;not null" json:"role"`
	Company  string `gorm:"type:text;not null" json:"company"`
	Location string `gorm:"type:text;not null" json:"location"`
	Remote   *bool  `gorm:"type:bool;not null" json:"remote"`
	Link     string `gorm:"type:text;not null" json:"link"`
	Salary   int64  `gorm:"type:int;not null" json:"salary"`
}

type UpdateOpeningRequest struct {
	Role     string `gorm:"type:text;not null" json:"role"`
	Company  string `gorm:"type:text;not null" json:"company"`
	Location string `gorm:"type:text;not null" json:"location"`
	Remote   *bool  `gorm:"type:bool;not null" json:"remote"`
	Link     string `gorm:"type:text;not null" json:"link"`
	Salary   int64  `gorm:"type:int;not null" json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Link == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return ErrParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return ErrParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return ErrParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return ErrParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return ErrParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return ErrParamIsRequired("salary", "int")
	}
	return nil
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Link != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
