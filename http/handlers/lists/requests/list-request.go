package requests

type ListRequest struct {
	Name     string `json:"name" 	 validate:"required,min=1,max=255"`
	Status   string `json:"status"   validate:"required,oneof=ACTIVE INACTIVE"`
	Position int16  `json:"position" validate:"number,min=0"`
}
