package request

type TaskCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=199"`
	Description string `json:"description" validate:"required,min=3"`
}
