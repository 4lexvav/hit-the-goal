package requests

type CommentRequest struct {
	Text string `json:"text" validate:"required,min=1,max=5000"`
}
