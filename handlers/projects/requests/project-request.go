package requests

type ProjectRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=500"`
	Description string `json:"description" validate:"max=1000"`
}
