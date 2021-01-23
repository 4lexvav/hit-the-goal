package requests

type TaskRequest struct {
	Name        string `json:"name" 	   validate:"required,min=1,max=500"`
	Description string `json:"description" validate:"max=5000"`
	Position    int16  `json:"position"    validate:"number,min=0"`
	NewListID   int64  `json:"new_list_id"  validate:"number"`
}
